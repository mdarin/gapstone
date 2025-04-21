#!/usr/bin/env ruby

require 'fileutils'
require 'open3'
require 'date'

# Проверка аргументов
unless ARGV.length == 1 && File.directory?(ARGV[0])
  $stderr.puts "Usage: #{$0} path/to/capstone/tests"
  exit 1
end


prefix = <<END
/*
Gapstone is a Go binding for the Capstone disassembly library.
For examples, try reading the *_test.go files.

	Library Author: Nguyen Anh Quynh
	Binding Author: Ben Nagy
	License: BSD style - see LICENSE file for details
    (c) 2013 COSEINC. All Rights Reserved.

    THIS FILE WAS AUTO-GENERATED -- DO NOT EDIT!
	Command: #{$0} #{ARGV[0]}
	Created at: #{DateTime.now}

*/

package gapstone

// List of generated specs to use in tests.

END


tests_dir = ARGV[0].chomp('/')

# Функция для преобразования строки в PascalCase
# Примеры
# puts to_pascal_case("xml_http_request")          # => "XmlHttpRequest"
# puts to_pascal_case("xml_http_request", preserve_acronyms: false) # => "XMLHTTPRequest"
# puts to_pascal_case("user_id")                   # => "UserId"
def to_pascal_case(str, preserve_acronyms: true)
  return "" if str.nil? || str.empty?

  # Разделяем строку на слова
  words = str.to_s.split(/[\s_\.-]/)
  
  words.map! do |word|
    if preserve_acronyms && word == word.upcase
      word
    else
      word.empty? ? "" : word[0].upcase + word[1..-1].downcase
    end
  end
  
  words.join
end

# Функция для переименования test_* в *.SPEC
def rename_test_to_spec(filename)
  # Удаляем путь и префикс test_, оставляя только имя
  base_name = File.basename(filename).gsub(/^test_/, '')
  
  # Заменяем расширение на .SPEC
  new_name = if base_name.include?('_')
               # Для файлов типа test_detail -> detail.SPEC
               "#{base_name}.SPEC"
             else
               # Для файлов типа test_x86 -> x86.SPEC
               "#{base_name}.SPEC"
             end
             
  #new_name.downcase!  # Приводим к нижнему регистру для единообразия
  new_name
end

# Получаем только тестовые файлы без расширения
test_files = Dir.glob("#{tests_dir}/test_*").select do |f|
  # Проверяем что это файл (не директория), исполняемый и без точки в имени
  File.file?(f) && 
  File.executable?(f) && 
  !f.include?('.') &&
  f.match?(%r{/test_[a-z0-9]+$}i)
end.sort

if test_files.empty?
  $stderr.puts "No test_* files found in #{tests_dir}"
  exit 1
end

gofn = "spec_list_constants.go"

File.open(gofn, 'w+') {|gofh|

  gofh.write prefix
  gofh.puts "const(\n"

  # Выполняем команды для каждого файла
  test_files.each do |test_file|
    output_file = rename_test_to_spec(test_file)
    command = "#{test_file} > #{output_file}"
    
    _, err, status = Open3.capture3(command)
    
    unless status.success?
      $stderr.puts "Error running #{command}: #{err}"
      next  # Продолжаем со следующим файлом вместо выхода
    end

    gofh.puts "\t#{to_pascal_case(output_file, preserve_acronyms: false)} = \"#{output_file}\""
    
    puts "✓ Generated: #{output_file}"
  end

  gofh.write(")\n") 
}

puts "\033[32mDone!\033[0m Generated #{test_files.size} .SPEC files"
