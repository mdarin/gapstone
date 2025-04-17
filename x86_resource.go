/*
Gapstone is a Go binding for the Capstone disassembly library.
For examples,try reading the *_test.go files.

	Library Author: Nguyen Anh Quynh
	Binding Author: Ben Nagy
	License: BSD style - see LICENSE file for details
    (c) 2013 COSEINC. All Rights Reserved.

    THIS FILE WAS AUTO-GENERATED -- DO NOT EDIT!
    Command: ./genresorces.sh capstone/tests/test_x86.c
    Created at: 2025-04-17 13:06:58

*/

package gapstone

//=== Found #define values in capstone/tests/test_x86.c ===

var x86Codes = map[string]string{
	// X86_CODE64="\x55\x48\x8b\x05\xb8\x13\x00\x00\xe9\xea\xbe\xad\xde\xff\x25\x23\x01\x00\x00\xe8\xdf\xbe\xad\xde\x74\xff"
	"X86_CODE64": "\x55\x48\x8b\x05\xb8\x13\x00\x00\xe9\xea\xbe\xad\xde\xff\x25\x23\x01\x00\x00\xe8\xdf\xbe\xad\xde\x74\xff",
	// X86_CODE16="\x8d\x4c\x32\x08\x01\xd8\x81\xc6\x34\x12\x00\x00\x05\x23\x01\x00\x00\x36\x8b\x84\x91\x23\x01\x00\x00\x41\x8d\x84\x39\x89\x67\x00\x00\x8d\x87\x89\x67\x00\x00\xb4\xc6\x66\xe9\xb8\x00\x00\x00\x67\xff\xa0\x23\x01\x00\x00\x66\xe8\xcb\x00\x00\x00\x74\xfc"
	"X86_CODE16": "\x8d\x4c\x32\x08\x01\xd8\x81\xc6\x34\x12\x00\x00\x05\x23\x01\x00\x00\x36\x8b\x84\x91\x23\x01\x00\x00\x41\x8d\x84\x39\x89\x67\x00\x00\x8d\x87\x89\x67\x00\x00\xb4\xc6\x66\xe9\xb8\x00\x00\x00\x67\xff\xa0\x23\x01\x00\x00\x66\xe8\xcb\x00\x00\x00\x74\xfc",
	// X86_CODE32="\x8d\x4c\x32\x08\x01\xd8\x81\xc6\x34\x12\x00\x00\x05\x23\x01\x00\x00\x36\x8b\x84\x91\x23\x01\x00\x00\x41\x8d\x84\x39\x89\x67\x00\x00\x8d\x87\x89\x67\x00\x00\xb4\xc6\xe9\xea\xbe\xad\xde\xff\xa0\x23\x01\x00\x00\xe8\xdf\xbe\xad\xde\x74\xff"
	"X86_CODE32": "\x8d\x4c\x32\x08\x01\xd8\x81\xc6\x34\x12\x00\x00\x05\x23\x01\x00\x00\x36\x8b\x84\x91\x23\x01\x00\x00\x41\x8d\x84\x39\x89\x67\x00\x00\x8d\x87\x89\x67\x00\x00\xb4\xc6\xe9\xea\xbe\xad\xde\xff\xa0\x23\x01\x00\x00\xe8\xdf\xbe\xad\xde\x74\xff",
}

//=== Found platform entries in capstone/tests/test_x86.c ===

var x86Platforms = []platform{
	{
		// -------------------------------------------
		// Platform Entry (line 379):
		// Architecture: CS_ARCH_X86
		// Mode: CS_MODE_16
		// Code: X86_CODE16
		// Comment: "X86 16bit (Intel syntax)"
		arch:    CS_ARCH_X86,
		mode:    CS_MODE_16,
		code:    x86Codes["X86_CODE16"],
		comment: "X86 16bit (Intel syntax)",
		options: []option{
			{CS_OPT_DETAIL, CS_OPT_ON},
		},
	},
	{
		// -------------------------------------------
		// Platform Entry (line 386):
		// Architecture: CS_ARCH_X86
		// Mode: CS_MODE_32
		// Code: X86_CODE32
		// Comment: "X86 32 (AT&T syntax)"
		//  Options:
		//    CS_OPT_SYNTAX: CS_OPT_SYNTAX_ATT
		arch:    CS_ARCH_X86,
		mode:    CS_MODE_32,
		code:    x86Codes["X86_CODE32"],
		comment: "X86 32 (AT&T syntax)",
		options: []option{
			{CS_OPT_DETAIL, CS_OPT_ON},
			{CS_OPT_SYNTAX, CS_OPT_SYNTAX_ATT},
		},
	},
	{
		// -------------------------------------------
		// Platform Entry (line 395):
		// Architecture: CS_ARCH_X86
		// Mode: CS_MODE_32
		// Code: X86_CODE32
		// Comment: "X86 32 (Intel syntax)"
		arch:    CS_ARCH_X86,
		mode:    CS_MODE_32,
		code:    x86Codes["X86_CODE32"],
		comment: "X86 32 (Intel syntax)",
		options: []option{
			{CS_OPT_DETAIL, CS_OPT_ON},
		},
	},
	{
		// -------------------------------------------
		// Platform Entry (line 402):
		// Architecture: CS_ARCH_X86
		// Mode: CS_MODE_64
		// Code: X86_CODE64
		// Comment: "X86 64 (Intel syntax)"
		arch:    CS_ARCH_X86,
		mode:    CS_MODE_64,
		code:    x86Codes["X86_CODE64"],
		comment: "X86 64 (Intel syntax)",
		options: []option{
			{CS_OPT_DETAIL, CS_OPT_ON},
		},
	},
}
