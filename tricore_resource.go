/*
Gapstone is a Go binding for the Capstone disassembly library.
For examples,try reading the *_test.go files.

	Library Author: Nguyen Anh Quynh
	Binding Author: Ben Nagy
	License: BSD style - see LICENSE file for details
    (c) 2013 COSEINC. All Rights Reserved.

    THIS FILE WAS AUTO-GENERATED -- DO NOT EDIT!
    Command: ./genresorces.sh capstone/tests/test_tricore.c
    Created at: 2025-04-25 12:56:14

*/

package gapstone

//=== Found #define values in capstone/tests/test_tricore.c ===

var tricoreCodes = map[string]string{
	// TRICORE_CODE="\x09\xcf\xbc\xf5\x09\xf4\x01\x00\x89\xfb\x8f\x74\x89\xfe\x48\x01\x29\x00\x19\x25\x29\x03\x09\xf4\x85\xf9\x68\x0f\x16\x01"
	"TRICORE_CODE": "\x09\xcf\xbc\xf5\x09\xf4\x01\x00\x89\xfb\x8f\x74\x89\xfe\x48\x01\x29\x00\x19\x25\x29\x03\x09\xf4\x85\xf9\x68\x0f\x16\x01",
}

//=== Found platform entries in capstone/tests/test_tricore.c ===

var tricorePlatforms = []platform{
	{
		// -------------------------------------------
		// Platform Entry (line 80):
		// Architecture: CS_ARCH_TRICORE
		// Mode: CS_MODE_TRICORE_162
		// Code: TRICORE_CODE
		// Comment: "TriCore"
		arch:    CS_ARCH_TRICORE,
		mode:    CS_MODE_TRICORE_162,
		code:    tricoreCodes["TRICORE_CODE"],
		comment: "TriCore",
		options: []option{
			{CS_OPT_DETAIL, CS_OPT_ON},
		},
	},
}
