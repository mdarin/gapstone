/*
Gapstone is a Go binding for the Capstone disassembly library.
For examples,try reading the *_test.go files.

	Library Author: Nguyen Anh Quynh
	Binding Author: Ben Nagy
	License: BSD style - see LICENSE file for details
    (c) 2013 COSEINC. All Rights Reserved.

    THIS FILE WAS AUTO-GENERATED -- DO NOT EDIT!
    Command: ./genresorces.sh capstone/tests/test_sparc.c
    Created at: 2025-04-17 13:06:58

*/

package gapstone

//=== Found #define values in capstone/tests/test_sparc.c ===

var sparcCodes = map[string]string{
	// SPARC_CODE="\x80\xa0\x40\x02\x85\xc2\x60\x08\x85\xe8\x20\x01\x81\xe8\x00\x00\x90\x10\x20\x01\xd5\xf6\x10\x16\x21\x00\x00\x0a\x86\x00\x40\x02\x01\x00\x00\x00\x12\xbf\xff\xff\x10\xbf\xff\xff\xa0\x02\x00\x09\x0d\xbf\xff\xff\xd4\x20\x60\x00\xd4\x4e\x00\x16\x2a\xc2\x80\x03"
	"SPARC_CODE": "\x80\xa0\x40\x02\x85\xc2\x60\x08\x85\xe8\x20\x01\x81\xe8\x00\x00\x90\x10\x20\x01\xd5\xf6\x10\x16\x21\x00\x00\x0a\x86\x00\x40\x02\x01\x00\x00\x00\x12\xbf\xff\xff\x10\xbf\xff\xff\xa0\x02\x00\x09\x0d\xbf\xff\xff\xd4\x20\x60\x00\xd4\x4e\x00\x16\x2a\xc2\x80\x03",
	// SPARCV9_CODE="\x81\xa8\x0a\x24\x89\xa0\x10\x20\x89\xa0\x1a\x60\x89\xa0\x00\xe0"
	"SPARCV9_CODE": "\x81\xa8\x0a\x24\x89\xa0\x10\x20\x89\xa0\x1a\x60\x89\xa0\x00\xe0",
}

//=== Found platform entries in capstone/tests/test_sparc.c ===

var sparcPlatforms = []platform{
	{
		// -------------------------------------------
		// Platform Entry (line 86):
		// Architecture: CS_ARCH_SPARC
		// Mode: CS_MODE_BIG_ENDIAN
		// Code: SPARC_CODE
		// Comment: "Sparc"
		arch:    CS_ARCH_SPARC,
		mode:    CS_MODE_BIG_ENDIAN,
		code:    sparcCodes["SPARC_CODE"],
		comment: "Sparc",
		options: []option{
			{CS_OPT_DETAIL, CS_OPT_ON},
		},
	},
	{
		// -------------------------------------------
		// Platform Entry (line 93):
		// Architecture: CS_ARCH_SPARC
		// Mode: (CS_MODE_BIG_ENDIAN + CS_MODE_V9)
		// Code: SPARCV9_CODE
		// Comment: "SparcV9"
		arch:    CS_ARCH_SPARC,
		mode:    (CS_MODE_BIG_ENDIAN + CS_MODE_V9),
		code:    sparcCodes["SPARCV9_CODE"],
		comment: "SparcV9",
		options: []option{
			{CS_OPT_DETAIL, CS_OPT_ON},
		},
	},
}
