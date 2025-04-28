/*
Gapstone is a Go binding for the Capstone disassembly library.
For examples,try reading the *_test.go files.

	Library Author: Nguyen Anh Quynh
	Binding Author: Ben Nagy
	License: BSD style - see LICENSE file for details
    (c) 2013 COSEINC. All Rights Reserved.

    THIS FILE WAS AUTO-GENERATED -- DO NOT EDIT!
    Command: ./genresorces.sh capstone/tests/test_systemz.c
    Created at: 2025-04-25 12:56:14

*/

package gapstone

//=== Found #define values in capstone/tests/test_systemz.c ===

var systemzCodes = map[string]string{
	// SYSZ_CODE="\xed\x00\x00\x00\x00\x1a\x5a\x0f\x1f\xff\xc2\x09\x80\x00\x00\x00\x07\xf7\xeb\x2a\xff\xff\x7f\x57\xe3\x01\xff\xff\x7f\x57\xeb\x00\xf0\x00\x00\x24\xb2\x4f\x00\x78\xec\x18\x00\x00\xc1\x7f"
	"SYSZ_CODE": "\xed\x00\x00\x00\x00\x1a\x5a\x0f\x1f\xff\xc2\x09\x80\x00\x00\x00\x07\xf7\xeb\x2a\xff\xff\x7f\x57\xe3\x01\xff\xff\x7f\x57\xeb\x00\xf0\x00\x00\x24\xb2\x4f\x00\x78\xec\x18\x00\x00\xc1\x7f",
}

//=== Found platform entries in capstone/tests/test_systemz.c ===

var systemzPlatforms = []platform{
	{
		// -------------------------------------------
		// Platform Entry (line 86):
		// Architecture: CS_ARCH_SYSZ
		// Mode: CS_MODE_BIG_ENDIAN
		// Code: SYSZ_CODE
		// Comment: "SystemZ"
		arch:    CS_ARCH_SYSZ,
		mode:    CS_MODE_BIG_ENDIAN,
		code:    systemzCodes["SYSZ_CODE"],
		comment: "SystemZ",
		options: []option{
			{CS_OPT_DETAIL, CS_OPT_ON},
		},
	},
}
