/*
Gapstone is a Go binding for the Capstone disassembly library.
For examples,try reading the *_test.go files.

	Library Author: Nguyen Anh Quynh
	Binding Author: Ben Nagy
	License: BSD style - see LICENSE file for details
    (c) 2013 COSEINC. All Rights Reserved.

    THIS FILE WAS AUTO-GENERATED -- DO NOT EDIT!
    Command: ./genresorces.sh capstone/tests/test_skipdata.c
    Created at: 2025-04-17 13:06:58

*/

package gapstone

//=== Found #define values in capstone/tests/test_skipdata.c ===

var skipdataCodes = map[string]string{
	// X86_CODE32="\x8d\x4c\x32\x08\x01\xd8\x81\xc6\x34\x12\x00\x00\x00\x91\x92"
	"X86_CODE32": "\x8d\x4c\x32\x08\x01\xd8\x81\xc6\x34\x12\x00\x00\x00\x91\x92",
	// RANDOM_CODE="\xed\x00\x00\x00\x00\x1a\x5a\x0f\x1f\xff\xc2\x09\x80\x00\x00\x00\x07\xf7\xeb\x2a\xff\xff\x7f\x57\xe3\x01\xff\xff\x7f\x57\xeb\x00\xf0\x00\x00\x24\xb2\x4f\x00\x78"
	"RANDOM_CODE": "\xed\x00\x00\x00\x00\x1a\x5a\x0f\x1f\xff\xc2\x09\x80\x00\x00\x00\x07\xf7\xeb\x2a\xff\xff\x7f\x57\xe3\x01\xff\xff\x7f\x57\xeb\x00\xf0\x00\x00\x24\xb2\x4f\x00\x78",
}

//=== Found platform entries in capstone/tests/test_skipdata.c ===

var skipdataPlatforms = []platform{
	{
		// -------------------------------------------
		// Platform Entry (line 64):
		// Architecture: CS_ARCH_X86
		// Mode: CS_MODE_32
		// Code: X86_CODE32
		// Comment: "X86 32 (Intel syntax) - Skip data"
		arch:    CS_ARCH_X86,
		mode:    CS_MODE_32,
		code:    skipdataCodes["X86_CODE32"],
		comment: "X86 32 (Intel syntax) - Skip data",
		options: []option{
			{CS_OPT_DETAIL, CS_OPT_ON},
		},
	},
	{
		// -------------------------------------------
		// Platform Entry (line 71):
		// Architecture: CS_ARCH_X86
		// Mode: CS_MODE_32
		// Code: X86_CODE32
		// Comment: "X86 32 (Intel syntax) - Skip data with custom mnemonic"
		//  Options:
		//    CS_OPT_INVALID: CS_OPT_OFF
		//    CS_OPT_SKIPDATA_SETUP: &skipdata
		arch:    CS_ARCH_X86,
		mode:    CS_MODE_32,
		code:    skipdataCodes["X86_CODE32"],
		comment: "X86 32 (Intel syntax) - Skip data with custom mnemonic",
		options: []option{
			{CS_OPT_DETAIL, CS_OPT_ON},
			// {CS_OPT_INVALID, CS_OPT_OFF},
			// {CS_OPT_SKIPDATA_SETUP, &skipdata},
		},
	},
	{
		// -------------------------------------------
		// Platform Entry (line 84):
		// Architecture: CS_ARCH_ARM
		// Mode: CS_MODE_ARM
		// Code: RANDOM_CODE
		// Comment: "Arm - Skip data"
		arch:    CS_ARCH_ARM,
		mode:    CS_MODE_ARM,
		code:    skipdataCodes["RANDOM_CODE"],
		comment: "Arm - Skip data",
		options: []option{
			{CS_OPT_DETAIL, CS_OPT_ON},
		},
	},
	{
		// -------------------------------------------
		// Platform Entry (line 91):
		// Architecture: CS_ARCH_ARM
		// Mode: CS_MODE_ARM
		// Code: RANDOM_CODE
		// Comment: "Arm - Skip data with callback"
		//  Options:
		//    CS_OPT_INVALID: CS_OPT_OFF
		//    CS_OPT_SKIPDATA_SETUP: &skipdata_callback
		arch:    CS_ARCH_ARM,
		mode:    CS_MODE_ARM,
		code:    skipdataCodes["RANDOM_CODE"],
		comment: "Arm - Skip data with callback",
		options: []option{
			{CS_OPT_DETAIL, CS_OPT_ON},
			// {CS_OPT_INVALID, CS_OPT_OFF},
			// {CS_OPT_SKIPDATA_SETUP, &skipdata_callback},
		},
	},
}
