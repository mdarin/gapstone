/*
Gapstone is a Go binding for the Capstone disassembly library.
For examples,try reading the *_test.go files.

	Library Author: Nguyen Anh Quynh
	Binding Author: Ben Nagy
	License: BSD style - see LICENSE file for details
    (c) 2013 COSEINC. All Rights Reserved.

    THIS FILE WAS AUTO-GENERATED -- DO NOT EDIT!
    Command: ./genresorces.sh capstone/tests/test_m68k.c
    Created at: 2025-04-25 12:56:14

*/

package gapstone

//=== Found #define values in capstone/tests/test_m68k.c ===

var m68kCodes = map[string]string{
	// M68K_CODE="\xf0\x10\xf0\x00\x48\xaf\xff\xff\x7f\xff\x11\xb0\x01\x37\x7f\xff\xff\xff\x12\x34\x56\x78\x01\x33\x10\x10\x10\x10\x32\x32\x32\x32\x4C\x00\x54\x04\x48\xe7\xe0\x30\x4C\xDF\x0C\x07\xd4\x40\x87\x5a\x4e\x71\x02\xb4\xc0\xde\xc0\xde\x5c\x00\x1d\x80\x71\x12\x01\x23\xf2\x3c\x44\x22\x40\x49\x0e\x56\x54\xc5\xf2\x3c\x44\x00\x44\x7a\x00\x00\xf2\x00\x0a\x28\x4E\xB9\x00\x00\x00\x12\x4E\x75"
	"M68K_CODE": "\xf0\x10\xf0\x00\x48\xaf\xff\xff\x7f\xff\x11\xb0\x01\x37\x7f\xff\xff\xff\x12\x34\x56\x78\x01\x33\x10\x10\x10\x10\x32\x32\x32\x32\x4C\x00\x54\x04\x48\xe7\xe0\x30\x4C\xDF\x0C\x07\xd4\x40\x87\x5a\x4e\x71\x02\xb4\xc0\xde\xc0\xde\x5c\x00\x1d\x80\x71\x12\x01\x23\xf2\x3c\x44\x22\x40\x49\x0e\x56\x54\xc5\xf2\x3c\x44\x00\x44\x7a\x00\x00\xf2\x00\x0a\x28\x4E\xB9\x00\x00\x00\x12\x4E\x75",
}

//=== Found platform entries in capstone/tests/test_m68k.c ===

var m68kPlatforms = []platform{
	{
		// -------------------------------------------
		// Platform Entry (line 156):
		// Architecture: CS_ARCH_M68K
		// Mode: (CS_MODE_BIG_ENDIAN | CS_MODE_M68K_040)
		// Code: M68K_CODE
		// Comment: "M68K"
		arch:    CS_ARCH_M68K,
		mode:    (CS_MODE_BIG_ENDIAN | CS_MODE_M68K_040),
		code:    m68kCodes["M68K_CODE"],
		comment: "M68K",
		options: []option{
			{CS_OPT_DETAIL, CS_OPT_ON},
		},
	},
}
