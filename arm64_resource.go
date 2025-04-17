/*
Gapstone is a Go binding for the Capstone disassembly library.
For examples,try reading the *_test.go files.

	Library Author: Nguyen Anh Quynh
	Binding Author: Ben Nagy
	License: BSD style - see LICENSE file for details
    (c) 2013 COSEINC. All Rights Reserved.

    THIS FILE WAS AUTO-GENERATED -- DO NOT EDIT!
    Command: ./genresorces.sh capstone/tests/test_arm64.c
    Created at: 2025-04-17 13:06:57

*/

package gapstone

//=== Found #define values in capstone/tests/test_arm64.c ===

var arm64Codes = map[string]string{
	// ARM64_CODE="\x09\x00\x38\xd5\xbf\x40\x00\xd5\x0c\x05\x13\xd5\x20\x50\x02\x0e\x20\xe4\x3d\x0f\x00\x18\xa0\x5f\xa2\x00\xae\x9e\x9f\x37\x03\xd5\xbf\x33\x03\xd5\xdf\x3f\x03\xd5\x21\x7c\x02\x9b\x21\x7c\x00\x53\x00\x40\x21\x4b\xe1\x0b\x40\xb9\x20\x04\x81\xda\x20\x08\x02\x8b\x10\x5b\xe8\x3c"
	"ARM64_CODE": "\x09\x00\x38\xd5\xbf\x40\x00\xd5\x0c\x05\x13\xd5\x20\x50\x02\x0e\x20\xe4\x3d\x0f\x00\x18\xa0\x5f\xa2\x00\xae\x9e\x9f\x37\x03\xd5\xbf\x33\x03\xd5\xdf\x3f\x03\xd5\x21\x7c\x02\x9b\x21\x7c\x00\x53\x00\x40\x21\x4b\xe1\x0b\x40\xb9\x20\x04\x81\xda\x20\x08\x02\x8b\x10\x5b\xe8\x3c",
}

//=== Found platform entries in capstone/tests/test_arm64.c ===

var arm64Platforms = []platform{
	{
		// -------------------------------------------
		// Platform Entry (line 168):
		// Architecture: CS_ARCH_ARM64
		// Mode: CS_MODE_ARM
		// Code: ARM64_CODE
		// Comment: "ARM-64"
		arch:    CS_ARCH_ARM64,
		mode:    CS_MODE_ARM,
		code:    arm64Codes["ARM64_CODE"],
		comment: "ARM-64",
		options: []option{
			{CS_OPT_DETAIL, CS_OPT_ON},
		},
	},
}
