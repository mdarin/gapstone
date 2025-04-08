#!/usr/bin/env ruby

require 'open3'

# I hate these little scripts :(

unless ARGV.length == 1 && File.directory?(ARGV[0])
	$stderr.puts "Usage: #{$0} path/to/capstone/tests"
	exit
end
stub = ARGV[0].chomp('/')

commands = [
	"#{stub}/test_basic > test.SPEC",
	"#{stub}/test_detail > test_detail.SPEC",
	"#{stub}/test_arm > arm.SPEC",
	"#{stub}/test_arm64 > arm64.SPEC",
    "#{stub}/test_evm > evm.SPEC",
	"#{stub}/test_x86 > x86.SPEC",
    "#{stub}/test_m68k > m68k.SPEC",
    "#{stub}/test_m680x > m680x.SPEC",
	"#{stub}/test_mips > mips.SPEC",
	"#{stub}/test_ppc > ppc.SPEC",
	"#{stub}/test_systemz > sysZ.SPEC",
	"#{stub}/test_sparc > sparc.SPEC",
    "#{stub}/test_tms320c64x > tms320c64x.SPEC",
	"#{stub}/test_xcore > xcore.SPEC"
]

commands.each {|c|
	_, err, status = Open3.capture3 c
	if not status.success?
		$stderr.puts "#{$0}: error running command: #{err}"
		exit
	end
	puts "Wrote: #{c.split.last}"
}
