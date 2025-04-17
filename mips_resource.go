/*
Gapstone is a Go binding for the Capstone disassembly library.
For examples,try reading the *_test.go files.

	Library Author: Nguyen Anh Quynh
	Binding Author: Ben Nagy
	License: BSD style - see LICENSE file for details
    (c) 2013 COSEINC. All Rights Reserved.

    THIS FILE WAS AUTO-GENERATED -- DO NOT EDIT!
    Command: ./genresorces.sh capstone/tests/test_mips.c
    Created at: 2025-04-17 13:06:58

*/

package gapstone

//=== Found #define values in capstone/tests/test_mips.c ===

var mipsCodes = map[string]string{
	// MIPS_CODE="\x0C\x10\x00\x97\x00\x00\x00\x00\x24\x02\x00\x0c\x8f\xa2\x00\x00\x34\x21\x34\x56"
	"MIPS_CODE": "\x0C\x10\x00\x97\x00\x00\x00\x00\x24\x02\x00\x0c\x8f\xa2\x00\x00\x34\x21\x34\x56",
	// MIPS_CODE2="\x56\x34\x21\x34\xc2\x17\x01\x00"
	"MIPS_CODE2": "\x56\x34\x21\x34\xc2\x17\x01\x00",
	// MIPS_32R6M="\x00\x07\x00\x07\x00\x11\x93\x7c\x01\x8c\x8b\x7c\x00\xc7\x48\xd0"
	"MIPS_32R6M": "\x00\x07\x00\x07\x00\x11\x93\x7c\x01\x8c\x8b\x7c\x00\xc7\x48\xd0",
	// MIPS_32R6="\xec\x80\x00\x19\x7c\x43\x22\xa0"
	"MIPS_32R6": "\xec\x80\x00\x19\x7c\x43\x22\xa0",
	// MIPS_64SD="\x70\x00\xb2\xff"
	"MIPS_64SD": "\x70\x00\xb2\xff",
}

//=== Found platform entries in capstone/tests/test_mips.c ===

var mipsPlatforms = []platform{
	{
		// -------------------------------------------
		// Platform Entry (line 81):
		// Architecture: CS_ARCH_MIPS
		// Mode: (CS_MODE_MIPS32 | CS_MODE_BIG_ENDIAN)
		// Code: MIPS_CODE
		// Comment: "MIPS-32 (Big-endian)"
		arch:    CS_ARCH_MIPS,
		mode:    (CS_MODE_MIPS32 | CS_MODE_BIG_ENDIAN),
		code:    mipsCodes["MIPS_CODE"],
		comment: "MIPS-32 (Big-endian)",
		options: []option{
			{CS_OPT_DETAIL, CS_OPT_ON},
		},
	},
	{
		// -------------------------------------------
		// Platform Entry (line 88):
		// Architecture: CS_ARCH_MIPS
		// Mode: (CS_MODE_MIPS64 | CS_MODE_LITTLE_ENDIAN)
		// Code: MIPS_CODE2
		// Comment: "MIPS-64-EL (Little-endian)"
		arch:    CS_ARCH_MIPS,
		mode:    (CS_MODE_MIPS64 | CS_MODE_LITTLE_ENDIAN),
		code:    mipsCodes["MIPS_CODE2"],
		comment: "MIPS-64-EL (Little-endian)",
		options: []option{
			{CS_OPT_DETAIL, CS_OPT_ON},
		},
	},
	{
		// -------------------------------------------
		// Platform Entry (line 95):
		// Architecture: CS_ARCH_MIPS
		// Mode: (CS_MODE_MIPS32R6 | CS_MODE_MICRO | CS_MODE_BIG_ENDIAN)
		// Code: MIPS_32R6M
		// Comment: "MIPS-32R6 | Micro (Big-endian)"
		arch:    CS_ARCH_MIPS,
		mode:    (CS_MODE_MIPS32R6 | CS_MODE_MICRO | CS_MODE_BIG_ENDIAN),
		code:    mipsCodes["MIPS_32R6M"],
		comment: "MIPS-32R6 | Micro (Big-endian)",
		options: []option{
			{CS_OPT_DETAIL, CS_OPT_ON},
		},
	},
	{
		// -------------------------------------------
		// Platform Entry (line 102):
		// Architecture: CS_ARCH_MIPS
		// Mode: (CS_MODE_MIPS32R6 | CS_MODE_BIG_ENDIAN)
		// Code: MIPS_32R6
		// Comment: "MIPS-32R6 (Big-endian)"
		arch:    CS_ARCH_MIPS,
		mode:    (CS_MODE_MIPS32R6 | CS_MODE_BIG_ENDIAN),
		code:    mipsCodes["MIPS_32R6"],
		comment: "MIPS-32R6 (Big-endian)",
		options: []option{
			{CS_OPT_DETAIL, CS_OPT_ON},
		},
	},
	{
		// -------------------------------------------
		// Platform Entry (line 109):
		// Architecture: CS_ARCH_MIPS
		// Mode: (CS_MODE_MIPS64 | CS_MODE_MIPS2 | CS_MODE_LITTLE_ENDIAN)
		// Code: MIPS_64SD
		// Comment: "MIPS-64-EL + Mips II (Little-endian)"
		arch:    CS_ARCH_MIPS,
		mode:    (CS_MODE_MIPS64 | CS_MODE_MIPS2 | CS_MODE_LITTLE_ENDIAN),
		code:    mipsCodes["MIPS_64SD"],
		comment: "MIPS-64-EL + Mips II (Little-endian)",
		options: []option{
			{CS_OPT_DETAIL, CS_OPT_ON},
		},
	},
	{
		// -------------------------------------------
		// Platform Entry (line 116):
		// Architecture: CS_ARCH_MIPS
		// Mode: (CS_MODE_MIPS64 | CS_MODE_LITTLE_ENDIAN)
		// Code: MIPS_64SD
		// Comment: "MIPS-64-EL (Little-endian)"
		arch:    CS_ARCH_MIPS,
		mode:    (CS_MODE_MIPS64 | CS_MODE_LITTLE_ENDIAN),
		code:    mipsCodes["MIPS_64SD"],
		comment: "MIPS-64-EL (Little-endian)",
		options: []option{
			{CS_OPT_DETAIL, CS_OPT_ON},
		},
	},
}
