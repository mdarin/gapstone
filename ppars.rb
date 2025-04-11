#!/usr/bin/env ruby
# frozen_string_literal: true

require 'strscan'
require 'pathname'
require 'timeout'

class CodeParser
  class ParseError < StandardError; end

  MAX_SCAN_LENGTH = 10_000
  TIMEOUT_SECONDS = 5
  DEBUG = ENV['DEBUG']

  def initialize(file_path)
    @file_path = file_path
    validate_file!
    reset_state
  end

  def parse
    Timeout.timeout(TIMEOUT_SECONDS) { parse_file_content }
  rescue Timeout::Error
    raise ParseError, timeout_error(@scanner)
  rescue ParseError => e
    raise e
  rescue => e
    raise ParseError, format_unexpected_error(e, @scanner)
  end

  private


  def format_unexpected_error(error, scanner)
    pos = scanner ? scanner.pos : 0
    backtrace = error.backtrace ? error.backtrace[0..3].join("\n") : "no backtrace"
    "Unexpected #{error.class.name} at position #{pos}: #{error.message[0..200]}\n" +
    "Context: #{current_file_context(scanner)}\n#{backtrace}"
  end

  def handle_unexpected_eof(scanner)
    return "Unexpected end of file" unless scanner

    error_details = detect_eof_cause(scanner)
    context = build_error_context(scanner)

    # Защита от пустого контекста
    if context.strip.empty?
      "Unexpected end of file#{error_details} at position #{scanner.pos}"
    else
      # Удаляем лишние переносы строк
      "Unexpected end of file#{error_details}#{context.gsub(/\n{3,}/, "\n\n")}"
    end
  end

  def detect_eof_cause(scanner)
    return "" unless scanner&.string
    
    if @conditional_stack.any?
      " (unclosed conditionals: #{@conditional_stack.size})"
    elsif scanner.string.match?(/\/\*[^*]*\z/)
      " (unclosed multi-line comment)"
    elsif scanner.string.match?(/\\\s*\z/)
      " (unterminated line continuation)"
    elsif scanner.string.match?(/"[^"]*\z/)
      " (unclosed string literal)"
    else
      ""
    end
  end

  def build_error_context(scanner)
    return "" unless scanner && scanner.pos.positive?
    
    pos = scanner.pos
    content = scanner.string
    return "" if content.nil? || content.empty?

    # Находим начало текущей строки
    line_start = content[0..pos].rindex("\n") || 0
    line_start += 1 if line_start < pos && content[line_start] == "\n"
    
    # Находим конец текущей строки
    line_end = content.index("\n", pos) || content.length
    
    # Извлекаем строку с ошибкой (защищаемся от отрицательных значений)
    error_line = if line_end > line_start
                  content[line_start...[line_end, content.length].min].chomp
                else
                  ""
                end

    # Пропускаем пустые строки
    return "" if error_line.strip.empty?

    # Вычисляем номер строки и позицию
    line_number = content[0..pos].count("\n") + 1
    column = pos - line_start + 1

    # Формируем контекст (защита от отрицательной колонки)
    if column.positive?
      "\nAt line #{line_number}, column #{column}:\n#{error_line}\n#{' ' * (column - 1)}^"
    else
      "\nAt line #{line_number} (position #{pos})"
    end
  rescue => e
    # Резервный вариант если что-то пошло не так
    "\n[Error building context: #{e.message[0..50]}]"
  end

  def current_file_context(scanner)
    return "" unless scanner && scanner.pos.positive?

    pos = scanner.pos
    content = scanner.string
    
    line = content[0..pos].count("\n") + 1
    last_newline = content[0..pos].rindex("\n") || -1
    column = pos - last_newline

    context_start = [last_newline + 1, 0].max
    context_end = [content.index("\n", pos) || content.length, content.length].min
    
    error_line = content[context_start...context_end].chomp
    indicator = "#{' ' * (column - 1)}^"

    "\nAt line #{line}, column #{column}:\n#{error_line}\n#{indicator}"
  end

  def infinite_loop_error(scanner)
    "Infinite loop detected at position #{scanner.pos}. Context: '#{scanner.peek(20)}...' #{current_file_context(scanner)}"
  end

  def scanning_too_long_error(scanner, pattern)
    "Scanning too far without match (over #{MAX_SCAN_LENGTH} chars). " +
    "Last pattern: #{pattern} #{current_file_context(scanner)}"
  end

  def timeout_error(scanner)
    pos = scanner ? scanner.pos : 0
    "Timeout while processing file. Last position: #{pos} #{current_file_context(scanner)}"
  end



  def skip_comments_and_whitespace
    loop do
      @scanner.skip(/[ \t\r\n]+/)
      
      if @scanner.skip(/\/\/.*?(\n|\z)/)
        next
      elsif @scanner.skip(/\/\*/)
        skip_multiline_comment
        next
      end
      
      break
    end
  end

  def skip_multiline_comment
    until @scanner.skip(/\*\//) || @scanner.eos?
      @scanner.getch
    end
  end

  def log_scan_position
    puts "[DEBUG] Pos: #{@scanner.pos}, Char: #{@scanner.peek(1)}"
  end

  def skip_irrelevant_code(scanner)
    # Пропускаем до следующей значимой конструкции
    scanner.scan_until(/(?=#define|struct\s+platform|#if|#endif|\z)/)
  end

  def reset_state
    @defines = {}
    @platforms = []
    @conditional_stack = []
    @current_conditional = true
    @scanner = nil
  end

  def validate_file!
    raise ParseError, "File not found: #{@file_path}" unless File.exist?(@file_path)
    raise ParseError, "Cannot read file: #{@file_path}" unless File.readable?(@file_path)
  end

  def parse_file_content
    content = File.read(@file_path)
    @scanner = StringScanner.new(content)
    scan_position = 0
    last_position = -1
    consecutive_failures = 0

    until @scanner.eos?
      log_scan_position if DEBUG

      skip_irrelevant_code(@scanner)

      if infinite_loop_detected?(last_position, consecutive_failures)
        raise ParseError, infinite_loop_error(@scanner)
      end
      last_position = @scanner.pos

      if scanning_too_long?(scan_position)
        raise ParseError, scanning_too_long_error(@scanner, @last_success_pattern)
      end

      if parse_relevant_line
        scan_position = @scanner.pos
        next
      end

      advance_scanner(scan_position)
    end

    { defines: @defines, platforms: @platforms }
  end

  def infinite_loop_detected?(last_position, consecutive_failures)
    if @scanner.pos == last_position
      consecutive_failures += 1
      consecutive_failures > 5
    else
      consecutive_failures = 0
      false
    end
  end

  def scanning_too_long?(scan_position)
    @scanner.pos - scan_position > MAX_SCAN_LENGTH
  end

  def parse_relevant_line                                                                                                                                                                         
    initial_pos = @scanner.pos                                                                                                                                                                    
    skip_comments_and_whitespace                                                                                                                                                                  
                                                                                                                                                                                                  
    if parse_conditional || parse_define(@scanner) || parse_platform_struct                                                                                                                       
      true                                                                                                                                                                                        
    else                                                                                                                                                                                          
      @scanner.pos = initial_pos                                                                                                                                                                  
      false                                                                                                                                                                                       
    end                                                                                                                                                                                           
  end 

  def parse_conditional
    return false unless @scanner.check(/#(ifdef|ifndef)\s+\w+/)

    start_pos = @scanner.pos
    conditional_depth = 0
    defines_in_block = []

    begin
      loop do
        skip_comments_and_whitespace

        if @scanner.scan(/#ifdef\s+(\w+)/)
          handle_ifdef(scanner[1])
          conditional_depth += 1
        elsif @scanner.scan(/#ifndef\s+(\w+)/)
          handle_ifndef(scanner[1])
          conditional_depth += 1
        elsif @scanner.scan(/#endif/)
          handle_endif
          conditional_depth -= 1
          break if conditional_depth.zero?
        elsif @current_conditional && @scanner.scan(/#define\s+(\w+)\s+/)
          parse_and_store_define
        else
          @scanner.getch || break
        end
      end

      true
    rescue => e
      @scanner.pos = start_pos
      raise ParseError, "Error parsing conditional block: #{e.message}"
    end
  end

  def parse_and_store_define
    name = @scanner[1]
    value = parse_define_value
    @defines[name] = value
  end

  def parse_define_value
    if @scanner.scan(/"/)
      parse_quoted_define_value
    else
      parse_unquoted_define_value
    end
  end

  def parse_quoted_define_value
    value = +""
    until @scanner.scan(/"/)
      value << (@scanner.scan(/[^\\"]+/) || "")
      if @scanner.scan(/\\/)
        escaped_char = @scanner.getch
        raise ParseError, "Unterminated escape sequence" unless escaped_char
        value << "\\" << escaped_char
      end
      raise ParseError, "Unterminated string literal" if @scanner.eos?
    end
    value
  end

  def parse_unquoted_define_value
    value = +""
    loop do
      value << (@scanner.scan(/[^\\\n]+/) || "")
      break unless @scanner.scan(/\\\s*\n\s*/)
      raise ParseError, "Unterminated line continuation" if @scanner.eos?
    end
    value.strip
  end

#   AI: исправь ошибку в этой функции
#   File: capstone/tests/test_wasm.c
#   ERROR: Unexpected ArgumentError at position 205: wrong number of arguments (given 0, expected 1)
# Context: 
# At line 12, column 1:
# struct platform {
# ^
# ./ppars.rb:320:in 'parse_define'
# ./ppars.rb:239:in 'CodeParser#parse_relevant_line'
# ./ppars.rb:210:in 'CodeParser#parse_file_content'
# ./ppars.rb:22:in 'block in CodeParser#parse' AI?
def parse_define(scanner)                                                                                                                                                                       
  return false unless @current_conditional                                                                                                                                                      
                                                                                                                                                                                                
  start_pos = scanner.pos                                                                                                                                                                       
  return false unless scanner.scan(/#define\s+(\w+)\s+/)                                                                                                                                        
                                                                                                                                                                                                
  name = scanner[1]                                                                                                                                                                             
  value = +""                                                                                                                                                                                   
  open_quote = false                                                                                                                                                                            
                                                                                                                                                                                                
  begin                                                                                                                                                                                         
    if scanner.scan(/"/)                                                                                                                                                                        
      open_quote = true                                                                                                                                                                         
      until scanner.scan(/"/)                                                                                                                                                                   
        chunk = scanner.scan(/[^\\"]+/) || ""                                                                                                                                                   
        value << chunk                                                                                                                                                                          
                                                                                                                                                                                                
        if scanner.scan(/\\/)                                                                                                                                                                   
          escaped_char = scanner.getch                                                                                                                                                          
          raise ParseError, "Unterminated escape sequence in define #{name}" unless escaped_char                                                                                                
          value << "\\" << escaped_char                                                                                                                                                         
        end                                                                                                                                                                                     
                                                                                                                                                                                                
        raise ParseError, "Unterminated string literal in define #{name}" if scanner.eos?                                                                                                       
      end                                                                                                                                                                                       
      open_quote = false                                                                                                                                                                        
    else                                                                                                                                                                                        
      loop do                                                                                                                                                                                   
        chunk = scanner.scan(/[^\\\n]+/) || ""                                                                                                                                                  
        value << chunk                                                                                                                                                                          
                                                                                                                                                                                                
        if scanner.check(/#(if|endif|define|undef)/)                                                                                                                                            
          raise ParseError, "Unexpected preprocessor directive inside define #{name}"                                                                                                           
        end                                                                                                                                                                                     
                                                                                                                                                                                                
        break unless scanner.scan(/\\\s*\n\s*/)                                                                                                                                                 
        raise ParseError, "Unterminated line continuation in define #{name}" if scanner.eos?                                                                                                    
      end                                                                                                                                                                                       
    end                                                                                                                                                                                         
                                                                                                                                                                                                
    @defines[name] = value.strip                                                                                                                                                                
    true                                                                                                                                                                                        
  rescue ParseError => e                                                                                                                                                                        
    scanner.pos = start_pos # Возвращаем позицию при ошибке                                                                                                                                     
    raise                                                                                                                                                                                       
  rescue => e                                                                                                                                                                                   
    scanner.pos = start_pos                                                                                                                                                                     
    raise ParseError, "Unexpected error processing define #{name}: #{e.message[0..100]}"                                                                                                        
  end                                                                                                                                                                                           
end  

# AI: исправь ошибку в этой функции
# File: capstone/tests/test_riscv.c
#   ERROR: Unexpected ArgumentError at position 103: wrong number of arguments (given 0, expected 1)
# Context: 
# At line 7, column 1:
# struct platform {
# ^
# ./ppars.rb:384:in 'parse_platform_struct'
# ./ppars.rb:239:in 'CodeParser#parse_relevant_line'
# ./ppars.rb:210:in 'CodeParser#parse_file_content'
# ./ppars.rb:22:in 'block in CodeParser#parse' AI!
  

  def parse_platform_struct(scanner)
    return false unless @current_conditional
    
    start_pos = scanner.pos
    return false unless scanner.scan(/struct\s+platform\s+platforms\s*\[\s*\]\s*=\s*\{\s*/)

    begin
      platforms_count = 0
      until scanner.scan(/\s*\};/)
        if scanner.pos - start_pos > MAX_SCAN_LENGTH
          raise ParseError, "Platform struct parsing too long at position #{scanner.pos}"
        end

        skip_comments_and_whitespace(scanner)

        if scanner.scan(/\{/)
          platform = parse_platform_entry(scanner)
          if platform
            @platforms << platform
            platforms_count += 1
          end
          scanner.scan(/\s*,\s*/)
        elsif scanner.scan(/\s*\}/)
          break
        else
          raise ParseError, "Expected platform entry or closing brace #{current_file_context(scanner)}"
        end
      end

      platforms_count.positive?
    rescue => e
      scanner&.pos = start_pos if scanner && start_pos
      raise ParseError, "Error parsing platform struct: #{e.message}"
    end
  end

  def parse_value(scanner)
    if scanner.scan(/\(unsigned char \*\)(\w+)/)
      scanner[1]
    elsif scanner.scan(/sizeof\((\w+)\)\s*-\s*1/)
      "#{scanner[1]}.size"
    elsif scanner.scan(/(?:\w+)\s*\|\s*(?:\w+)/)
      scanner.matched
    elsif scanner.scan(/\w+/)
      scanner.matched
    elsif scanner.scan(/".*?"/)
      scanner.matched.gsub(/"/, '')
    else
      nil
    end
  end

  def resolve_code(code_ref)
    return code_ref unless code_ref.is_a?(String)
    code_ref = code_ref.gsub(/\(unsigned char \*\)/, '')
    @defines.fetch(code_ref) { code_ref }
  end

  def parse_platform_entry(scanner)
    fields = []
    comments = []

    until scanner.scan(/\s*\}/)
      skip_comments_and_whitespace(scanner)
      
      if scanner.scan(/\s*\/\/(.*?)\n/)
        comments << scanner[1].strip
      elsif scanner.scan(/\/\*/)
        # Пропускаем весь блок комментария
        until scanner.scan(/\*\//) || scanner.eos?
          scanner.getch
        end
      elsif value = parse_value(scanner)
        fields << value
        scanner.scan(/\s*,\s*/)
      else
        raise ParseError, "Unexpected token in platform entry"
      end
    end

    {
      arch: fields[0],
      mode: fields[1],
      code: resolve_code(fields[2]),
      size: fields[3],
      comment: fields[4],
      comments: comments
    }
  end

  
end

def process_directory(directory)
  results = {}
  processed_files = 0
  failed_files = 0

  Dir.glob(File.join(directory, '**/*.c')).each do |file|
    puts "Processing #{file}..."
    processed_files += 1

    begin
      parser = CodeParser.new(file)
      result = parser.parse
      results[file] = result
      puts "  Success: #{result[:defines].size} defines, #{result[:platforms].size} platforms"
    rescue CodeParser::ParseError => e
      puts "  Parse Error: #{e.message}"
      results[file] = { error: e.message }
      failed_files += 1
    rescue => e
      error_msg = "Fatal Error: #{e.message[0..200]}"
      puts "  #{error_msg}"
      results[file] = { error: error_msg }
      failed_files += 1
    end
  end

  puts "\nProcessing complete. Files: #{processed_files}, Success: #{processed_files - failed_files}, Failed: #{failed_files}"
  results
end

def print_results(results)
  results.each do |file, result|
    puts "\nFile: #{file}"
    
    if result[:error]
      puts "  ERROR: #{result[:error]}"
      next
    end

    if result[:defines].any?
      puts "  Defines:"
      result[:defines].each { |name, value| puts "    #{name} = #{value.inspect}" }
    else
      puts "  No defines found"
    end

    if result[:platforms].any?
      puts "  Platforms:"
      result[:platforms].each_with_index do |platform, i|
        puts "    Platform #{i + 1}:"
        platform.each { |key, value| puts "      #{key}: #{value.inspect}" }
      end
    else
      puts "  No platforms found"
    end
  end
end

if ARGV.empty?
  puts "Usage: #{$PROGRAM_NAME} <input_file.c or directory>"
  exit 1
end

input_path = ARGV[0]

begin
  if File.directory?(input_path)
    results = process_directory(input_path)
    print_results(results)
  else
    parser = CodeParser.new(input_path)
    result = parser.parse
    print_results({ input_path => result })
  end
rescue CodeParser::ParseError => e
  puts "Error: #{e.message}"
  exit 1
rescue => e
  puts "Fatal error: #{e.message}"
  exit 2
end
