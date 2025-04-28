/*
Gapstone is a Go binding for the Capstone disassembly library.
For examples,try reading the *_test.go files.

	Library Author: Nguyen Anh Quynh
	Binding Author: Ben Nagy
	License: BSD style - see LICENSE file for details
    (c) 2013 COSEINC. All Rights Reserved.

    THIS FILE WAS AUTO-GENERATED -- DO NOT EDIT!
    Command: ./genresorces.sh capstone/tests/test_sh.c
    Created at: 2025-04-25 12:56:14

*/

package gapstone

//=== Found #define values in capstone/tests/test_sh.c ===

var shCodes = map[string]string{
	// SH4A_CODE="\xc\x31\x10\x20\x22\x21\x36\x64\x46\x25\x12\x12\x1c\x2\x8\xc1\x5\xc7\xc\x71\x1f\x2\x22\xcf\x6\x89\x23\x0\x2b\x41\xb\x0\xe\x40\x32\x0\xa\xf1\x9\x0"
	"SH4A_CODE": "\xc0\x31\x10\x20\x22\x21\x36\x64\x46\x25\x12\x12\x1c\x20\x80\xc1\x50\xc7\xc0\x71\x1f\x20\x22\xcf\x60\x89\x23\x00\x2b\x41\xb0\x00\xe0\x40\x32\x00\xa0\xf1\x90\x00",
	// SH2A_CODE="\x32\x11\x92\x0\x32\x49\x31\x0"
	"SH2A_CODE": "\x32\x11\x92\x00\x32\x49\x31\x00",
}

//=== Found platform entries in capstone/tests/test_sh.c ===

var shPlatforms = []platform{
	{
		// -------------------------------------------
		// Platform Entry (line 176):
		// Architecture: CS_ARCH_SH
		// Mode: (CS_MODE_SH4A | CS_MODE_SHFPU)
		// Code: SH4A_CODE
		// Comment: "SH_SH4A"
		arch:    CS_ARCH_SH,
		mode:    (CS_MODE_SH4A | CS_MODE_SHFPU),
		code:    shCodes["SH4A_CODE"],
		comment: "SH_SH4A",
		options: []option{
			{CS_OPT_DETAIL, CS_OPT_ON},
		},
	},
	{
		// -------------------------------------------
		// Platform Entry (line 183):
		// Architecture: CS_ARCH_SH
		// Mode: (CS_MODE_SH2A | CS_MODE_SHFPU | CS_MODE_BIG_ENDIAN)
		// Code: SH2A_CODE
		// Comment: "SH_SH2A"
		arch:    CS_ARCH_SH,
		mode:    (CS_MODE_SH2A | CS_MODE_SHFPU | CS_MODE_BIG_ENDIAN),
		code:    shCodes["SH2A_CODE"],
		comment: "SH_SH2A",
		options: []option{
			{CS_OPT_DETAIL, CS_OPT_ON},
		},
	},
}
