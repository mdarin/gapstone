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

pyfiles = Dir.glob(File.join(ARGV[0], "*_const.py"))
if pyfiles.empty?
	fail "No *_const.py files found in #{ARGV[0]}"
end

pyfiles.each {|pyfn|

	gofn = "#{File.basename(pyfn).split('_').first}_constants.go"

	File.open(gofn, 'w+') {|gofh|

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
				print "Weird line #{l}"
			end
		}

		gofh.write(")\n\n") if @clumping
		
	}

	# if all is well, this will nicely format all the files we just wrote. <3
	#`go fmt`
}
