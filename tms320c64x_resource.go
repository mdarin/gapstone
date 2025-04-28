/*
Gapstone is a Go binding for the Capstone disassembly library.
For examples,try reading the *_test.go files.

	Library Author: Nguyen Anh Quynh
	Binding Author: Ben Nagy
	License: BSD style - see LICENSE file for details
    (c) 2013 COSEINC. All Rights Reserved.

    THIS FILE WAS AUTO-GENERATED -- DO NOT EDIT!
    Command: ./genresorces.sh capstone/tests/test_tms320c64x.c
    Created at: 2025-04-25 12:56:14

*/

package gapstone

//=== Found #define values in capstone/tests/test_tms320c64x.c ===

var tms320c64xCodes = map[string]string{
	// TMS320C64X_CODE="\x01\xac\x88\x40\x81\xac\x88\x43\x00\x00\x00\x00\x02\x90\x32\x96\x02\x80\x46\x9e\x05\x3c\x83\xe6\x0b\x0c\x8b\x24"
	"TMS320C64X_CODE": "\x01\xac\x88\x40\x81\xac\x88\x43\x00\x00\x00\x00\x02\x90\x32\x96\x02\x80\x46\x9e\x05\x3c\x83\xe6\x0b\x0c\x8b\x24",
}

//=== Found platform entries in capstone/tests/test_tms320c64x.c ===

var tms320c64xPlatforms = []platform{
	{
		// -------------------------------------------
		// Platform Entry (line 135):
		// Architecture: CS_ARCH_TMS320C64X
		// Mode: CS_MODE_BIG_ENDIAN
		// Code: TMS320C64X_CODE
		// Comment: "TMS320C64x"
		arch:    CS_ARCH_TMS320C64X,
		mode:    CS_MODE_BIG_ENDIAN,
		code:    tms320c64xCodes["TMS320C64X_CODE"],
		comment: "TMS320C64x",
		options: []option{
			{CS_OPT_DETAIL, CS_OPT_ON},
		},
	},
}
