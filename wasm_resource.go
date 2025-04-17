/*
Gapstone is a Go binding for the Capstone disassembly library.
For examples,try reading the *_test.go files.

	Library Author: Nguyen Anh Quynh
	Binding Author: Ben Nagy
	License: BSD style - see LICENSE file for details
    (c) 2013 COSEINC. All Rights Reserved.

    THIS FILE WAS AUTO-GENERATED -- DO NOT EDIT!
    Command: ./genresorces.sh capstone/tests/test_wasm.c
    Created at: 2025-04-17 13:06:58

*/

package gapstone

//=== Found #define values in capstone/tests/test_wasm.c ===

var wasmCodes = map[string]string{
	// WASM_CODE="\x20\x00\x20\x01\x41\x20\x10\xc9\x01\x45\x0b"
	"WASM_CODE": "\x20\x00\x20\x01\x41\x20\x10\xc9\x01\x45\x0b",
}

//=== Found platform entries in capstone/tests/test_wasm.c ===

var wasmPlatforms = []platform{
	{
		// -------------------------------------------
		// Platform Entry (line 91):
		// Architecture: CS_ARCH_WASM
		// Mode: 0
		// Code: WASM_CODE
		// Comment: "WASM"
		arch:    CS_ARCH_WASM,
		mode:    0,
		code:    wasmCodes["WASM_CODE"],
		comment: "WASM",
		options: []option{
			{CS_OPT_DETAIL, CS_OPT_ON},
		},
	},
}
