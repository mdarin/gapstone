/*
Gapstone is a Go binding for the Capstone disassembly library.
For examples,try reading the *_test.go files.

	Library Author: Nguyen Anh Quynh
	Binding Author: Ben Nagy
	License: BSD style - see LICENSE file for details
    (c) 2013 COSEINC. All Rights Reserved.

    THIS FILE WAS AUTO-GENERATED -- DO NOT EDIT!
    Command: ./genresorces.sh capstone/tests/test_xcore.c
    Created at: 2025-04-17 13:06:57

*/

package gapstone

//=== Found #define values in capstone/tests/test_xcore.c ===

var xcoreCodes = map[string]string{
	// XCORE_CODE="\xfe\x0f\xfe\x17\x13\x17\xc6\xfe\xec\x17\x97\xf8\xec\x4f\x1f\xfd\xec\x37\x07\xf2\x45\x5b\xf9\xfa\x02\x06\x1b\x10\x09\xfd\xec\xa7"
	"XCORE_CODE": "\xfe\x0f\xfe\x17\x13\x17\xc6\xfe\xec\x17\x97\xf8\xec\x4f\x1f\xfd\xec\x37\x07\xf2\x45\x5b\xf9\xfa\x02\x06\x1b\x10\x09\xfd\xec\xa7",
}

//=== Found platform entries in capstone/tests/test_xcore.c ===

var xcorePlatforms = []platform{
	{
		// -------------------------------------------
		// Platform Entry (line 81):
		// Architecture: CS_ARCH_XCORE
		// Mode: CS_MODE_BIG_ENDIAN
		// Code: XCORE_CODE
		// Comment: "XCore"
		arch:    CS_ARCH_XCORE,
		mode:    CS_MODE_BIG_ENDIAN,
		code:    xcoreCodes["XCORE_CODE"],
		comment: "XCore",
		options: []option{
			{CS_OPT_DETAIL, CS_OPT_ON},
		},
	},
}
