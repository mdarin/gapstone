#!/usr/bin/env ruby
# frozen_string_literal: true

# Для одного файла
# ruby parser.rb test.c

# Для всей директории (рекурсивно)
# ruby parser.rb /path/to/code/

#AI: почему зависает скрипт, где может быть ошибка? AI!

require 'strscan'
require 'pathname'

class CodeParser
  class ParseError < StandardError; end

  def initialize(file_path)
    @file_path = file_path
    @defines = {}
    @platforms = []
    @conditional_stack = []
    @current_conditional = true
    @skip_mode = false  # Флаг для пропуска блоков кода
  end

  def parse
    content = File.read(@file_path)
    scanner = StringScanner.new(content)

    until scanner.eos?
      if scanner.scan(/\s+/)
        next
      elsif parse_conditional(scanner)
        next
      elsif @skip_mode
        skip_irrelevant_code(scanner)
        next
      elsif parse_define(scanner)
        next
      elsif parse_platform_struct(scanner)
        next
      else
        # Если не распознано как значимая конструкция, пропускаем
        skip_irrelevant_code(scanner)
      end
    end

    { defines: @defines, platforms: @platforms }
  rescue ParseError => e
    { error: "Parse error in #{@file_path}: #{e.message}" }
  rescue => e
    { error: "Unexpected error in #{@file_path}: #{e.message}\n#{e.backtrace.join("\n")}" }
  end

  private

  def skip_irrelevant_code(scanner)
    # Пропускаем до следующей значимой конструкции
    scanner.scan_until(/(?=#define|struct\s+platform|#if|#endif|\z)/)
  end

  # ... (остальные методы остаются без изменений)
  def parse_conditional(scanner)
    if scanner.scan(/#ifdef\s+(\w+)/)
      @conditional_stack.push(@current_conditional)
      @current_conditional &&= @defines.key?(scanner[1])
      true
    elsif scanner.scan(/#ifndef\s+(\w+)/)
      @conditional_stack.push(@current_conditional)
      @current_conditional &&= !@defines.key?(scanner[1])
      true
    elsif scanner.scan(/#endif/)
      raise ParseError, "Unmatched #endif" if @conditional_stack.empty?
      @current_conditional = @conditional_stack.pop
      true
    else
      false
    end
  end

  def parse_define(scanner)
    return false unless @current_conditional
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
      error_msg = "While processing define #{name}: #{e.message}"
      error_msg += " (unclosed quote)" if open_quote
      raise ParseError, error_msg
    rescue => e
      raise ParseError, "Unexpected error processing define #{name}: #{e.message}"
    end
  end

  def parse_platform_struct(scanner)
    return false unless @current_conditional
    return false unless scanner.scan(/struct\s+platform\s+platforms\s*\[\s*\]\s*=\s*\{\s*/)

    begin
      until scanner.scan(/\s*\};/)
        next if scanner.scan(/\s*\/\/.*?\n|\s*\/\*.*?\*\/|\s+/)

        if scanner.scan(/\{/)
          platform = parse_platform_entry(scanner)
          @platforms << platform if platform
          scanner.scan(/\s*,\s*/)
        else
          raise ParseError, "Expected platform entry"
        end
      end

      true
    rescue => e
      raise ParseError, "Error parsing platform struct: #{e.message}"
    end
  end

  def parse_platform_entry(scanner)
    fields = []
    comments = []

    begin
      until scanner.scan(/\s*\}/)
        if scanner.scan(/\s*\/\/(.*?)\n/)
          comments << scanner[1].strip
        elsif scanner.scan(/\s*\/\*.*?\*\//)
          next
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
    rescue => e
      raise ParseError, "Error parsing platform entry: #{e.message}"
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
end

def process_directory(directory)
  results = {}
  Dir.glob(File.join(directory, '**/*.c')).each do |file|
    puts "Processing #{file}..."
    parser = CodeParser.new(file)
    result = parser.parse
    
    if result[:error]
      puts "  Error: #{result[:error]}"
      results[file] = { error: result[:error] }
    else
      results[file] = result
      puts "  Found #{result[:defines].size} defines and #{result[:platforms].size} platforms"
    end
  end
  results
end

def print_results(results)
  results.each do |file, result|
    puts "\nFile: #{file}"
    if result[:error]
      puts "  ERROR: #{result[:error]}"
      next
    end

    puts "  Defines:"
    result[:defines].each do |name, value|
      puts "    #{name} = #{value.inspect}"
    end

    puts "\n  Platforms:"
    result[:platforms].each_with_index do |platform, i|
      puts "    Platform #{i + 1}:"
      platform.each do |key, value|
        puts "      #{key}: #{value.inspect}"
      end
    end
  end
end

if ARGV.empty?
  puts "Usage: #{$PROGRAM_NAME} <input_file.c or directory>"
  exit 1
end

input_path = ARGV[0]

if File.directory?(input_path)
  results = process_directory(input_path)
  print_results(results)
else
  #TODO: 
  # begin
  #   parser.parse
  # rescue CodeParser::ParseError => e
  #   puts "Parse error in #{file_path}: #{e.message}"
  #   next # или другая логика обработки ошибки
  # end
  parser = CodeParser.new(input_path)
  result = parser.parse

  if result[:error]
    puts result[:error]
    exit 1
  else
    print_results({ input_path => result })
  end
end