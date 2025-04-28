/*
Gapstone is a Go binding for the Capstone disassembly library.
For examples,try reading the *_test.go files.

	Library Author: Nguyen Anh Quynh
	Binding Author: Ben Nagy
	License: BSD style - see LICENSE file for details
    (c) 2013 COSEINC. All Rights Reserved.

    THIS FILE WAS AUTO-GENERATED -- DO NOT EDIT!
    Command: ./genresorces.sh capstone/tests/test_evm.c
    Created at: 2025-04-25 12:56:14

*/

package gapstone

//=== Found #define values in capstone/tests/test_evm.c ===

var evmCodes = map[string]string{
	// EVM_CODE="\x60\x61\x50"
	"EVM_CODE": "\x60\x61\x50",
}

//=== Found platform entries in capstone/tests/test_evm.c ===

var evmPlatforms = []platform{
	{
		// -------------------------------------------
		// Platform Entry (line 67):
		// Architecture: CS_ARCH_EVM
		// Mode: 0
		// Code: EVM_CODE
		// Comment: "EVM"
		arch:    CS_ARCH_EVM,
		mode:    0,
		code:    evmCodes["EVM_CODE"],
		comment: "EVM",
		options: []option{
			{CS_OPT_DETAIL, CS_OPT_ON},
		},
	},
}
