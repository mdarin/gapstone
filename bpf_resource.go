/*
Gapstone is a Go binding for the Capstone disassembly library.
For examples,try reading the *_test.go files.

	Library Author: Nguyen Anh Quynh
	Binding Author: Ben Nagy
	License: BSD style - see LICENSE file for details
    (c) 2013 COSEINC. All Rights Reserved.

    THIS FILE WAS AUTO-GENERATED -- DO NOT EDIT!
    Command: ./genresorces.sh capstone/tests/test_bpf.c
    Created at: 2025-04-25 12:56:14

*/

package gapstone

//=== Found #define values in capstone/tests/test_bpf.c ===

var bpfCodes = map[string]string{
	// CBPF_CODE="\x94\x09\x00\x00\x37\x13\x03\x00\x87\x00\x00\x00\x00\x00\x00\x00\x07\x00\x00\x00\x00\x00\x00\x00\x16\x00\x00\x00\x00\x00\x00\x00\x80\x00\x00\x00\x00\x00\x00\x00"
	"CBPF_CODE": "\x94\x09\x00\x00\x37\x13\x03\x00\x87\x00\x00\x00\x00\x00\x00\x00\x07\x00\x00\x00\x00\x00\x00\x00\x16\x00\x00\x00\x00\x00\x00\x00\x80\x00\x00\x00\x00\x00\x00\x00",
	// EBPF_CODE="\x97\x09\x00\x00\x37\x13\x03\x00\xdc\x02\x00\x00\x20\x00\x00\x00\x30\x00\x00\x00\x00\x00\x00\x00\xdb\x3a\x00\x01\x00\x00\x00\x00\x84\x02\x00\x00\x00\x00\x00\x00\x6d\x33\x17\x02\x00\x00\x00\x00"
	"EBPF_CODE": "\x97\x09\x00\x00\x37\x13\x03\x00\xdc\x02\x00\x00\x20\x00\x00\x00\x30\x00\x00\x00\x00\x00\x00\x00\xdb\x3a\x00\x01\x00\x00\x00\x00\x84\x02\x00\x00\x00\x00\x00\x00\x6d\x33\x17\x02\x00\x00\x00\x00",
}

//=== Found platform entries in capstone/tests/test_bpf.c ===

var bpfPlatforms = []platform{
	{
		// -------------------------------------------
		// Platform Entry (line 127):
		// Architecture: CS_ARCH_BPF
		// Mode: CS_MODE_LITTLE_ENDIAN | CS_MODE_BPF_CLASSIC
		// Code: CBPF_CODE
		// Comment: "cBPF Le"
		arch:    CS_ARCH_BPF,
		mode:    CS_MODE_LITTLE_ENDIAN | CS_MODE_BPF_CLASSIC,
		code:    bpfCodes["CBPF_CODE"],
		comment: "cBPF Le",
		options: []option{
			{CS_OPT_DETAIL, CS_OPT_ON},
		},
	},
	{
		// -------------------------------------------
		// Platform Entry (line 134):
		// Architecture: CS_ARCH_BPF
		// Mode: CS_MODE_LITTLE_ENDIAN | CS_MODE_BPF_EXTENDED
		// Code: EBPF_CODE
		// Comment: "eBPF Le"
		arch:    CS_ARCH_BPF,
		mode:    CS_MODE_LITTLE_ENDIAN | CS_MODE_BPF_EXTENDED,
		code:    bpfCodes["EBPF_CODE"],
		comment: "eBPF Le",
		options: []option{
			{CS_OPT_DETAIL, CS_OPT_ON},
		},
	},
}
