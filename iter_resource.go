/*
Gapstone is a Go binding for the Capstone disassembly library.
For examples,try reading the *_test.go files.

	Library Author: Nguyen Anh Quynh
	Binding Author: Ben Nagy
	License: BSD style - see LICENSE file for details
    (c) 2013 COSEINC. All Rights Reserved.

    THIS FILE WAS AUTO-GENERATED -- DO NOT EDIT!
    Command: ./genresorces.sh capstone/tests/test_iter.c
    Created at: 2025-04-25 12:56:14

*/

package gapstone

//=== Found #define values in capstone/tests/test_iter.c ===

var iterCodes = map[string]string{
	// X86_CODE16="\x8d\x4c\x32\x08\x01\xd8\x81\xc6\x34\x12\x00\x00"
	"X86_CODE16": "\x8d\x4c\x32\x08\x01\xd8\x81\xc6\x34\x12\x00\x00",
	// X86_CODE32="\x8d\x4c\x32\x08\x01\xd8\x81\xc6\x34\x12\x00\x00"
	"X86_CODE32": "\x8d\x4c\x32\x08\x01\xd8\x81\xc6\x34\x12\x00\x00",
	// X86_CODE64="\x55\x48\x8b\x05\xb8\x13\x00\x00"
	"X86_CODE64": "\x55\x48\x8b\x05\xb8\x13\x00\x00",
	// ARM_CODE="\xED\xFF\xFF\xEB\x04\xe0\x2d\xe5\x00\x00\x00\x00\xe0\x83\x22\xe5\xf1\x02\x03\x0e\x00\x00\xa0\xe3\x02\x30\xc1\xe7\x00\x00\x53\xe3"
	"ARM_CODE": "\xED\xFF\xFF\xEB\x04\xe0\x2d\xe5\x00\x00\x00\x00\xe0\x83\x22\xe5\xf1\x02\x03\x0e\x00\x00\xa0\xe3\x02\x30\xc1\xe7\x00\x00\x53\xe3",
	// ARM_CODE2="\x10\xf1\x10\xe7\x11\xf2\x31\xe7\xdc\xa1\x2e\xf3\xe8\x4e\x62\xf3"
	"ARM_CODE2": "\x10\xf1\x10\xe7\x11\xf2\x31\xe7\xdc\xa1\x2e\xf3\xe8\x4e\x62\xf3",
	// THUMB_CODE="\x70\x47\xeb\x46\x83\xb0\xc9\x68"
	"THUMB_CODE": "\x70\x47\xeb\x46\x83\xb0\xc9\x68",
	// THUMB_CODE2="\x4f\xf0\x00\x01\xbd\xe8\x00\x88\xd1\xe8\x00\xf0"
	"THUMB_CODE2": "\x4f\xf0\x00\x01\xbd\xe8\x00\x88\xd1\xe8\x00\xf0",
	// MIPS_CODE="\x0C\x10\x00\x97\x00\x00\x00\x00\x24\x02\x00\x0c\x8f\xa2\x00\x00\x34\x21\x34\x56\x00\x80\x04\x08"
	"MIPS_CODE": "\x0C\x10\x00\x97\x00\x00\x00\x00\x24\x02\x00\x0c\x8f\xa2\x00\x00\x34\x21\x34\x56\x00\x80\x04\x08",
	// MIPS_CODE2="\x56\x34\x21\x34\xc2\x17\x01\x00"
	"MIPS_CODE2": "\x56\x34\x21\x34\xc2\x17\x01\x00",
	// ARM64_CODE="\x09\x00\x38\xd5\xbf\x40\x00\xd5\x0c\x05\x13\xd5\x20\x50\x02\x0e\x20\xe4\x3d\x0f\x00\x18\xa0\x5f\xa2\x00\xae\x9e\x9f\x37\x03\xd5\xbf\x33\x03\xd5\xdf\x3f\x03\xd5\x21\x7c\x02\x9b\x21\x7c\x00\x53\x00\x40\x21\x4b\xe1\x0b\x40\xb9\x20\x04\x81\xda\x20\x08\x02\x8b\x10\x5b\xe8\x3c"
	"ARM64_CODE": "\x09\x00\x38\xd5\xbf\x40\x00\xd5\x0c\x05\x13\xd5\x20\x50\x02\x0e\x20\xe4\x3d\x0f\x00\x18\xa0\x5f\xa2\x00\xae\x9e\x9f\x37\x03\xd5\xbf\x33\x03\xd5\xdf\x3f\x03\xd5\x21\x7c\x02\x9b\x21\x7c\x00\x53\x00\x40\x21\x4b\xe1\x0b\x40\xb9\x20\x04\x81\xda\x20\x08\x02\x8b\x10\x5b\xe8\x3c",
	// PPC_CODE="\x80\x20\x00\x00\x80\x3f\x00\x00\x10\x43\x23\x0e\xd0\x44\x00\x80\x4c\x43\x22\x02\x2d\x03\x00\x80\x7c\x43\x20\x14\x7c\x43\x20\x93\x4f\x20\x00\x21\x4c\xc8\x00\x21\x40\x82\x00\x14"
	"PPC_CODE": "\x80\x20\x00\x00\x80\x3f\x00\x00\x10\x43\x23\x0e\xd0\x44\x00\x80\x4c\x43\x22\x02\x2d\x03\x00\x80\x7c\x43\x20\x14\x7c\x43\x20\x93\x4f\x20\x00\x21\x4c\xc8\x00\x21\x40\x82\x00\x14",
	// SPARC_CODE="\x80\xa0\x40\x02\x85\xc2\x60\x08\x85\xe8\x20\x01\x81\xe8\x00\x00\x90\x10\x20\x01\xd5\xf6\x10\x16\x21\x00\x00\x0a\x86\x00\x40\x02\x01\x00\x00\x00\x12\xbf\xff\xff\x10\xbf\xff\xff\xa0\x02\x00\x09\x0d\xbf\xff\xff\xd4\x20\x60\x00\xd4\x4e\x00\x16\x2a\xc2\x80\x03"
	"SPARC_CODE": "\x80\xa0\x40\x02\x85\xc2\x60\x08\x85\xe8\x20\x01\x81\xe8\x00\x00\x90\x10\x20\x01\xd5\xf6\x10\x16\x21\x00\x00\x0a\x86\x00\x40\x02\x01\x00\x00\x00\x12\xbf\xff\xff\x10\xbf\xff\xff\xa0\x02\x00\x09\x0d\xbf\xff\xff\xd4\x20\x60\x00\xd4\x4e\x00\x16\x2a\xc2\x80\x03",
	// SPARCV9_CODE="\x81\xa8\x0a\x24\x89\xa0\x10\x20\x89\xa0\x1a\x60\x89\xa0\x00\xe0"
	"SPARCV9_CODE": "\x81\xa8\x0a\x24\x89\xa0\x10\x20\x89\xa0\x1a\x60\x89\xa0\x00\xe0",
	// SYSZ_CODE="\xed\x00\x00\x00\x00\x1a\x5a\x0f\x1f\xff\xc2\x09\x80\x00\x00\x00\x07\xf7\xeb\x2a\xff\xff\x7f\x57\xe3\x01\xff\xff\x7f\x57\xeb\x00\xf0\x00\x00\x24\xb2\x4f\x00\x78"
	"SYSZ_CODE": "\xed\x00\x00\x00\x00\x1a\x5a\x0f\x1f\xff\xc2\x09\x80\x00\x00\x00\x07\xf7\xeb\x2a\xff\xff\x7f\x57\xe3\x01\xff\xff\x7f\x57\xeb\x00\xf0\x00\x00\x24\xb2\x4f\x00\x78",
	// XCORE_CODE="\xfe\x0f\xfe\x17\x13\x17\xc6\xfe\xec\x17\x97\xf8\xec\x4f\x1f\xfd\xec\x37\x07\xf2\x45\x5b\xf9\xfa\x02\x06\x1b\x10"
	"XCORE_CODE": "\xfe\x0f\xfe\x17\x13\x17\xc6\xfe\xec\x17\x97\xf8\xec\x4f\x1f\xfd\xec\x37\x07\xf2\x45\x5b\xf9\xfa\x02\x06\x1b\x10",
	// M680X_CODE="\x06\x10\x19\x1a\x55\x1e\x01\x23\xe9\x31\x06\x34\x55\xa6\x81\xa7\x89\x7f\xff\xa6\x9d\x10\x00\xa7\x91\xa6\x9f\x10\x00\x11\xac\x99\x10\x00\x39"
	"M680X_CODE": "\x06\x10\x19\x1a\x55\x1e\x01\x23\xe9\x31\x06\x34\x55\xa6\x81\xa7\x89\x7f\xff\xa6\x9d\x10\x00\xa7\x91\xa6\x9f\x10\x00\x11\xac\x99\x10\x00\x39",
	// MOS65XX_CODE="\x0d\x34\x12\x08\x09\xFF\x10\x80\x20\x00\x00\x98"
	"MOS65XX_CODE": "\x0d\x34\x12\x08\x09\xFF\x10\x80\x20\x00\x00\x98",
	// EBPF_CODE="\x97\x09\x00\x00\x37\x13\x03\x00\xdc\x02\x00\x00\x20\x00\x00\x00\x30\x00\x00\x00\x00\x00\x00\x00\xdb\x3a\x00\x01\x00\x00\x00\x00\x84\x02\x00\x00\x00\x00\x00\x00\x6d\x33\x17\x02\x00\x00\x00\x00"
	"EBPF_CODE": "\x97\x09\x00\x00\x37\x13\x03\x00\xdc\x02\x00\x00\x20\x00\x00\x00\x30\x00\x00\x00\x00\x00\x00\x00\xdb\x3a\x00\x01\x00\x00\x00\x00\x84\x02\x00\x00\x00\x00\x00\x00\x6d\x33\x17\x02\x00\x00\x00\x00",
	// RISCV_CODE32="\x37\x34\x00\x00\x97\x82\x00\x00\xef\x00\x80\x00\xef\xf0\x1f\xff\xe7\x00\x45\x00\xe7\x00\xc0\xff\x63\x05\x41\x00\xe3\x9d\x61\xfe\x63\xca\x93\x00\x63\x53\xb5\x00\x63\x65\xd6\x00\x63\x76\xf7\x00\x03\x88\x18\x00\x03\x99\x49\x00\x03\xaa\x6a\x00\x03\xcb\x2b\x01\x03\xdc\x8c\x01\x23\x86\xad\x03\x23\x9a\xce\x03\x23\x8f\xef\x01\x93\x00\xe0\x00\x13\xa1\x01\x01\x13\xb2\x02\x7d\x13\xc3\x03\xdd\x13\xe4\xc4\x12\x13\xf5\x85\x0c\x13\x96\xe6\x01\x13\xd7\x97\x01\x13\xd8\xf8\x40\x33\x89\x49\x01\xb3\x0a\x7b\x41\x33\xac\xac\x01\xb3\x3d\xde\x01\x33\xd2\x62\x40\xb3\x43\x94\x00\x33\xe5\xc5\x00\xb3\x76\xf7\x00\xb3\x54\x39\x01\xb3\x50\x31\x00\x33\x9f\x0f\x00"
	"RISCV_CODE32": "\x37\x34\x00\x00\x97\x82\x00\x00\xef\x00\x80\x00\xef\xf0\x1f\xff\xe7\x00\x45\x00\xe7\x00\xc0\xff\x63\x05\x41\x00\xe3\x9d\x61\xfe\x63\xca\x93\x00\x63\x53\xb5\x00\x63\x65\xd6\x00\x63\x76\xf7\x00\x03\x88\x18\x00\x03\x99\x49\x00\x03\xaa\x6a\x00\x03\xcb\x2b\x01\x03\xdc\x8c\x01\x23\x86\xad\x03\x23\x9a\xce\x03\x23\x8f\xef\x01\x93\x00\xe0\x00\x13\xa1\x01\x01\x13\xb2\x02\x7d\x13\xc3\x03\xdd\x13\xe4\xc4\x12\x13\xf5\x85\x0c\x13\x96\xe6\x01\x13\xd7\x97\x01\x13\xd8\xf8\x40\x33\x89\x49\x01\xb3\x0a\x7b\x41\x33\xac\xac\x01\xb3\x3d\xde\x01\x33\xd2\x62\x40\xb3\x43\x94\x00\x33\xe5\xc5\x00\xb3\x76\xf7\x00\xb3\x54\x39\x01\xb3\x50\x31\x00\x33\x9f\x0f\x00",
	// RISCV_CODE64="\x13\x04\xa8\x7a"
	"RISCV_CODE64": "\x13\x04\xa8\x7a",
	// TRICORE_CODE="\x16\x01\x20\x01\x1d\x00\x02\x00\x8f\x70\x00\x11\x40\xae\x89\xee\x04\x09\x42\xf2\xe2\xf2\xc2\x11\x19\xff\xc0\x70\x19\xff\x20\x10"
	"TRICORE_CODE": "\x16\x01\x20\x01\x1d\x00\x02\x00\x8f\x70\x00\x11\x40\xae\x89\xee\x04\x09\x42\xf2\xe2\xf2\xc2\x11\x19\xff\xc0\x70\x19\xff\x20\x10",
}

//=== Found platform entries in capstone/tests/test_iter.c ===

var iterPlatforms = []platform{
	{
		// -------------------------------------------
		// Platform Entry (line 84):
		// Architecture: CS_ARCH_X86
		// Mode: CS_MODE_16
		// Code: X86_CODE16
		// Comment: "X86 16bit (Intel syntax)"
		arch:    CS_ARCH_X86,
		mode:    CS_MODE_16,
		code:    iterCodes["X86_CODE16"],
		comment: "X86 16bit (Intel syntax)",
		options: []option{
			{CS_OPT_DETAIL, CS_OPT_ON},
		},
	},
	{
		// -------------------------------------------
		// Platform Entry (line 91):
		// Architecture: CS_ARCH_X86
		// Mode: CS_MODE_32
		// Code: X86_CODE32
		// Comment: "X86 32bit (ATT syntax)"
		//  Options:
		//    CS_OPT_SYNTAX: CS_OPT_SYNTAX_ATT
		arch:    CS_ARCH_X86,
		mode:    CS_MODE_32,
		code:    iterCodes["X86_CODE32"],
		comment: "X86 32bit (ATT syntax)",
		options: []option{
			{CS_OPT_DETAIL, CS_OPT_ON},
			{CS_OPT_SYNTAX, CS_OPT_SYNTAX_ATT},
		},
	},
	{
		// -------------------------------------------
		// Platform Entry (line 100):
		// Architecture: CS_ARCH_X86
		// Mode: CS_MODE_32
		// Code: X86_CODE32
		// Comment: "X86 32 (Intel syntax)"
		arch:    CS_ARCH_X86,
		mode:    CS_MODE_32,
		code:    iterCodes["X86_CODE32"],
		comment: "X86 32 (Intel syntax)",
		options: []option{
			{CS_OPT_DETAIL, CS_OPT_ON},
		},
	},
	{
		// -------------------------------------------
		// Platform Entry (line 107):
		// Architecture: CS_ARCH_X86
		// Mode: CS_MODE_64
		// Code: X86_CODE64
		// Comment: "X86 64 (Intel syntax)"
		arch:    CS_ARCH_X86,
		mode:    CS_MODE_64,
		code:    iterCodes["X86_CODE64"],
		comment: "X86 64 (Intel syntax)",
		options: []option{
			{CS_OPT_DETAIL, CS_OPT_ON},
		},
	},
	{
		// -------------------------------------------
		// Platform Entry (line 116):
		// Architecture: CS_ARCH_ARM
		// Mode: CS_MODE_ARM
		// Code: ARM_CODE
		// Comment: "ARM"
		arch:    CS_ARCH_ARM,
		mode:    CS_MODE_ARM,
		code:    iterCodes["ARM_CODE"],
		comment: "ARM",
		options: []option{
			{CS_OPT_DETAIL, CS_OPT_ON},
		},
	},
	{
		// -------------------------------------------
		// Platform Entry (line 123):
		// Architecture: CS_ARCH_ARM
		// Mode: CS_MODE_THUMB
		// Code: THUMB_CODE2
		// Comment: "THUMB-2"
		arch:    CS_ARCH_ARM,
		mode:    CS_MODE_THUMB,
		code:    iterCodes["THUMB_CODE2"],
		comment: "THUMB-2",
		options: []option{
			{CS_OPT_DETAIL, CS_OPT_ON},
		},
	},
	{
		// -------------------------------------------
		// Platform Entry (line 130):
		// Architecture: CS_ARCH_ARM
		// Mode: CS_MODE_ARM
		// Code: ARM_CODE2
		// Comment: "ARM: Cortex-A15 + NEON"
		arch:    CS_ARCH_ARM,
		mode:    CS_MODE_ARM,
		code:    iterCodes["ARM_CODE2"],
		comment: "ARM: Cortex-A15 + NEON",
		options: []option{
			{CS_OPT_DETAIL, CS_OPT_ON},
		},
	},
	{
		// -------------------------------------------
		// Platform Entry (line 137):
		// Architecture: CS_ARCH_ARM
		// Mode: CS_MODE_THUMB
		// Code: THUMB_CODE
		// Comment: "THUMB"
		arch:    CS_ARCH_ARM,
		mode:    CS_MODE_THUMB,
		code:    iterCodes["THUMB_CODE"],
		comment: "THUMB",
		options: []option{
			{CS_OPT_DETAIL, CS_OPT_ON},
		},
	},
	{
		// -------------------------------------------
		// Platform Entry (line 146):
		// Architecture: CS_ARCH_MIPS
		// Mode: (CS_MODE_MIPS32 + CS_MODE_BIG_ENDIAN)
		// Code: MIPS_CODE
		// Comment: "MIPS-32 (Big-endian)"
		arch:    CS_ARCH_MIPS,
		mode:    (CS_MODE_MIPS32 + CS_MODE_BIG_ENDIAN),
		code:    iterCodes["MIPS_CODE"],
		comment: "MIPS-32 (Big-endian)",
		options: []option{
			{CS_OPT_DETAIL, CS_OPT_ON},
		},
	},
	{
		// -------------------------------------------
		// Platform Entry (line 153):
		// Architecture: CS_ARCH_MIPS
		// Mode: (CS_MODE_MIPS64 + CS_MODE_LITTLE_ENDIAN)
		// Code: MIPS_CODE2
		// Comment: "MIPS-64-EL (Little-endian)"
		arch:    CS_ARCH_MIPS,
		mode:    (CS_MODE_MIPS64 + CS_MODE_LITTLE_ENDIAN),
		code:    iterCodes["MIPS_CODE2"],
		comment: "MIPS-64-EL (Little-endian)",
		options: []option{
			{CS_OPT_DETAIL, CS_OPT_ON},
		},
	},
	{
		// -------------------------------------------
		// Platform Entry (line 162):
		// Architecture: CS_ARCH_ARM64
		// Mode: CS_MODE_ARM
		// Code: ARM64_CODE
		// Comment: "ARM-64"
		arch:    CS_ARCH_ARM64,
		mode:    CS_MODE_ARM,
		code:    iterCodes["ARM64_CODE"],
		comment: "ARM-64",
		options: []option{
			{CS_OPT_DETAIL, CS_OPT_ON},
		},
	},
	{
		// -------------------------------------------
		// Platform Entry (line 171):
		// Architecture: CS_ARCH_PPC
		// Mode: CS_MODE_BIG_ENDIAN
		// Code: PPC_CODE
		// Comment: "PPC-64"
		arch:    CS_ARCH_PPC,
		mode:    CS_MODE_BIG_ENDIAN,
		code:    iterCodes["PPC_CODE"],
		comment: "PPC-64",
		options: []option{
			{CS_OPT_DETAIL, CS_OPT_ON},
		},
	},
	{
		// -------------------------------------------
		// Platform Entry (line 180):
		// Architecture: CS_ARCH_SPARC
		// Mode: CS_MODE_BIG_ENDIAN
		// Code: SPARC_CODE
		// Comment: "Sparc"
		arch:    CS_ARCH_SPARC,
		mode:    CS_MODE_BIG_ENDIAN,
		code:    iterCodes["SPARC_CODE"],
		comment: "Sparc",
		options: []option{
			{CS_OPT_DETAIL, CS_OPT_ON},
		},
	},
	{
		// -------------------------------------------
		// Platform Entry (line 187):
		// Architecture: CS_ARCH_SPARC
		// Mode: (CS_MODE_BIG_ENDIAN + CS_MODE_V9)
		// Code: SPARCV9_CODE
		// Comment: "SparcV9"
		arch:    CS_ARCH_SPARC,
		mode:    (CS_MODE_BIG_ENDIAN + CS_MODE_V9),
		code:    iterCodes["SPARCV9_CODE"],
		comment: "SparcV9",
		options: []option{
			{CS_OPT_DETAIL, CS_OPT_ON},
		},
	},
	{
		// -------------------------------------------
		// Platform Entry (line 196):
		// Architecture: CS_ARCH_SYSZ
		// Mode: 0
		// Code: SYSZ_CODE
		// Comment: "SystemZ"
		arch:    CS_ARCH_SYSZ,
		mode:    0,
		code:    iterCodes["SYSZ_CODE"],
		comment: "SystemZ",
		options: []option{
			{CS_OPT_DETAIL, CS_OPT_ON},
		},
	},
	{
		// -------------------------------------------
		// Platform Entry (line 205):
		// Architecture: CS_ARCH_XCORE
		// Mode: 0
		// Code: XCORE_CODE
		// Comment: "XCore"
		arch:    CS_ARCH_XCORE,
		mode:    0,
		code:    iterCodes["XCORE_CODE"],
		comment: "XCore",
		options: []option{
			{CS_OPT_DETAIL, CS_OPT_ON},
		},
	},
	{
		// -------------------------------------------
		// Platform Entry (line 214):
		// Architecture: CS_ARCH_M680X
		// Mode: CS_MODE_M680X_6809
		// Code: M680X_CODE
		// Comment: "M680X_6809"
		arch:    CS_ARCH_M680X,
		mode:    CS_MODE_M680X_6809,
		code:    iterCodes["M680X_CODE"],
		comment: "M680X_6809",
		options: []option{
			{CS_OPT_DETAIL, CS_OPT_ON},
		},
	},
	{
		// -------------------------------------------
		// Platform Entry (line 223):
		// Architecture: CS_ARCH_MOS65XX
		// Mode: CS_MODE_LITTLE_ENDIAN
		// Code: MOS65XX_CODE
		// Comment: "MOS65XX"
		arch:    CS_ARCH_MOS65XX,
		mode:    CS_MODE_LITTLE_ENDIAN,
		code:    iterCodes["MOS65XX_CODE"],
		comment: "MOS65XX",
		options: []option{
			{CS_OPT_DETAIL, CS_OPT_ON},
		},
	},
	{
		// -------------------------------------------
		// Platform Entry (line 232):
		// Architecture: CS_ARCH_BPF
		// Mode: CS_MODE_LITTLE_ENDIAN | CS_MODE_BPF_EXTENDED
		// Code: EBPF_CODE
		// Comment: "eBPF"
		arch:    CS_ARCH_BPF,
		mode:    CS_MODE_LITTLE_ENDIAN | CS_MODE_BPF_EXTENDED,
		code:    iterCodes["EBPF_CODE"],
		comment: "eBPF",
		options: []option{
			{CS_OPT_DETAIL, CS_OPT_ON},
		},
	},
	{
		// -------------------------------------------
		// Platform Entry (line 241):
		// Architecture: CS_ARCH_RISCV
		// Mode: CS_MODE_RISCV32
		// Code: RISCV_CODE32
		// Comment: "RISCV32"
		arch:    CS_ARCH_RISCV,
		mode:    CS_MODE_RISCV32,
		code:    iterCodes["RISCV_CODE32"],
		comment: "RISCV32",
		options: []option{
			{CS_OPT_DETAIL, CS_OPT_ON},
		},
	},
	{
		// -------------------------------------------
		// Platform Entry (line 248):
		// Architecture: CS_ARCH_RISCV
		// Mode: CS_MODE_RISCV64
		// Code: RISCV_CODE64
		// Comment: "RISCV64"
		arch:    CS_ARCH_RISCV,
		mode:    CS_MODE_RISCV64,
		code:    iterCodes["RISCV_CODE64"],
		comment: "RISCV64",
		options: []option{
			{CS_OPT_DETAIL, CS_OPT_ON},
		},
	},
	{
		// -------------------------------------------
		// Platform Entry (line 257):
		// Architecture: CS_ARCH_TRICORE
		// Mode: CS_MODE_TRICORE_162
		// Code: TRICORE_CODE
		// Comment: "TriCore"
		arch:    CS_ARCH_TRICORE,
		mode:    CS_MODE_TRICORE_162,
		code:    iterCodes["TRICORE_CODE"],
		comment: "TriCore",
		options: []option{
			{CS_OPT_DETAIL, CS_OPT_ON},
		},
	},
}
