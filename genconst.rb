#!/usr/bin/env ruby

require 'date'

# I hate these little scripts :(

unless ARGV.length == 1
	fail "Usage: #{$0} path/to/capstone/python/capstone"
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

// #cgo LDFLAGS: -lcapstone
// #cgo freebsd CFLAGS: -I/usr/local/include
// #cgo freebsd LDFLAGS: -L/usr/local/lib
// #include <stdlib.h>
// #include <capstone/capstone.h>
import "C"

END

def find_python_files(directory, recursive: false)
	# Определяем шаблон поиска в зависимости от режима
	pattern = if recursive
	  File.join(directory, "**", "{*_const.py,__init__.py}")
	else
	  File.join(directory, "{*_const.py,__init__.py}")
	end
  
	# Ищем файлы, исключая директории
	Dir.glob(pattern).select { |f| File.file?(f) }.uniq
  end


dir = ARGV[0] || "."
puts "Поиск в директории: #{dir}"

# Обычный поиск
pyfiles = find_python_files(dir)

# * INFO Рекурсивный поиск
#  puts find_python_files(dir, recursive: true)

if pyfiles.empty?
	fail "No *_const.py files found in #{ARGV[0]}"
end


def generate_go_filename(pyfn, default_prefix: 'capstone')
	# Извлекаем базовое имя до первого подчеркивания
	base = File.basename(pyfn, '.py').split('_').first
	
	# Формируем имя .go файла
	if base.nil? || base.empty? || base == pyfn
	  "#{default_prefix}_constants.go"
	else
	  "#{base}_constants.go"
	end
  end
  
# Тестовые случаи
#   {
# 	"module_const.py" => "module_constants.go",
# 	"_const.py"      => "capstone_constants.go",
# 	"const.py"       => "capstone_constants.go",
# 	"test.py"        => "capstone_constants.go",
# 	nil              => "capstone_constants.go"
#   }.each do |input, expected|
# 	output = generate_go_filename(input)
# 	puts "Input: #{input || 'nil'} => #{output} (Expected: #{expected}) #{output == expected ? '✓' : '✗'}"
#   end

pyfiles.each {|pyfn|
	# версия кода с проверкой и добавлением префикса capstone_, если результат _constants.go
	gofn = generate_go_filename(pyfn)

	File.open(gofn, 'w+') {|gofh|

  		print "✓ #{File.basename(pyfn)} -> #{gofn}\n"

		gofh.write prefix
		@clumping = false

		File.foreach(pyfn) {|l|
			case l
			when /^#|^$/
				# close any clump
				gofh.write(")\n\n") if @clumping
				@clumping = false
				if l =~ /^#/
					# Emit a go-style comment
					gofh.write "#{l.sub('#', %q(//))}"
				end
			when /^[A-Z]/
				gofh.puts("const (") unless @clumping
				@clumping = true
				const = l.split.first
		
				gofh.puts "\t#{const} = C.#{const}"
			else
				# print "Weird line #{l}"
				next
			end
		}

		gofh.write(")\n\n") if @clumping
		
	}


	#TODO: move it to another stage
	#  if all is well, this will nicely format all the files we just wrote. <3
	# `go fmt` 
}

