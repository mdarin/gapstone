/*
Gapstone is a Go binding for the Capstone disassembly library.
For examples,try reading the *_test.go files.

	Library Author: Nguyen Anh Quynh
	Binding Author: Ben Nagy
	License: BSD style - see LICENSE file for details
    (c) 2013 COSEINC. All Rights Reserved.

    THIS FILE WAS AUTO-GENERATED -- DO NOT EDIT!
    Command: ./genresorces.sh capstone/tests/test_ppc.c
    Created at: 2025-04-17 13:06:57

*/

package gapstone

//=== Found #define values in capstone/tests/test_ppc.c ===

var ppcCodes = map[string]string{
	// PPC_CODE="\x43\x20\x0c\x07\x41\x56\xff\x17\x80\x20\x00\x00\x80\x3f\x00\x00\x10\x43\x23\x0e\xd0\x44\x00\x80\x4c\x43\x22\x02\x2d\x03\x00\x80\x7c\x43\x20\x14\x7c\x43\x20\x93\x4f\x20\x00\x21\x4c\xc8\x00\x21\x40\x82\x00\x14"
	"PPC_CODE": "\x43\x20\x0c\x07\x41\x56\xff\x17\x80\x20\x00\x00\x80\x3f\x00\x00\x10\x43\x23\x0e\xd0\x44\x00\x80\x4c\x43\x22\x02\x2d\x03\x00\x80\x7c\x43\x20\x14\x7c\x43\x20\x93\x4f\x20\x00\x21\x4c\xc8\x00\x21\x40\x82\x00\x14",
	// PPC_CODE2="\x10\x60\x2a\x10\x10\x64\x28\x88\x7c\x4a\x5d\x0f"
	"PPC_CODE2": "\x10\x60\x2a\x10\x10\x64\x28\x88\x7c\x4a\x5d\x0f",
}

//=== Found platform entries in capstone/tests/test_ppc.c ===

var ppcPlatforms = []platform{
	{
		// -------------------------------------------
		// Platform Entry (line 120):
		// Architecture: CS_ARCH_PPC
		// Mode: CS_MODE_BIG_ENDIAN
		// Code: PPC_CODE
		// Comment: "PPC-64"
		arch:    CS_ARCH_PPC,
		mode:    CS_MODE_BIG_ENDIAN,
		code:    ppcCodes["PPC_CODE"],
		comment: "PPC-64",
		options: []option{
			{CS_OPT_DETAIL, CS_OPT_ON},
		},
	},
	{
		// -------------------------------------------
		// Platform Entry (line 127):
		// Architecture: CS_ARCH_PPC
		// Mode: (CS_MODE_BIG_ENDIAN + CS_MODE_QPX)
		// Code: PPC_CODE2
		// Comment: "PPC-64 + QPX"
		arch:    CS_ARCH_PPC,
		mode:    (CS_MODE_BIG_ENDIAN + CS_MODE_QPX),
		code:    ppcCodes["PPC_CODE2"],
		comment: "PPC-64 + QPX",
		options: []option{
			{CS_OPT_DETAIL, CS_OPT_ON},
		},
	},
}
