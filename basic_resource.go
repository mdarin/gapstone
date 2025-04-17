/*
Gapstone is a Go binding for the Capstone disassembly library.
For examples,try reading the *_test.go files.

	Library Author: Nguyen Anh Quynh
	Binding Author: Ben Nagy
	License: BSD style - see LICENSE file for details
    (c) 2013 COSEINC. All Rights Reserved.

    THIS FILE WAS AUTO-GENERATED -- DO NOT EDIT!
    Command: ./genresorces.sh capstone/tests/test_basic.c
    Created at: 2025-04-17 13:06:57

*/

package gapstone

//=== Found #define values in capstone/tests/test_basic.c ===

var basicCodes = map[string]string{
	// X86_CODE16="\x8d\x4c\x32\x08\x01\xd8\x81\xc6\x34\x12\x00\x00"
	"X86_CODE16": "\x8d\x4c\x32\x08\x01\xd8\x81\xc6\x34\x12\x00\x00",
	// X86_CODE32="\xba\xcd\xab\x00\x00\x8d\x4c\x32\x08\x01\xd8\x81\xc6\x34\x12\x00\x00"
	"X86_CODE32": "\xba\xcd\xab\x00\x00\x8d\x4c\x32\x08\x01\xd8\x81\xc6\x34\x12\x00\x00",
	// X86_CODE64="\x55\x48\x8b\x05\xb8\x13\x00\x00"
	"X86_CODE64": "\x55\x48\x8b\x05\xb8\x13\x00\x00",
	// ARM_CODE="\xED\xFF\xFF\xEB\x04\xe0\x2d\xe5\x00\x00\x00\x00\xe0\x83\x22\xe5\xf1\x02\x03\x0e\x00\x00\xa0\xe3\x02\x30\xc1\xe7\x00\x00\x53\xe3"
	"ARM_CODE": "\xED\xFF\xFF\xEB\x04\xe0\x2d\xe5\x00\x00\x00\x00\xe0\x83\x22\xe5\xf1\x02\x03\x0e\x00\x00\xa0\xe3\x02\x30\xc1\xe7\x00\x00\x53\xe3",
	// ARM_CODE2="\x10\xf1\x10\xe7\x11\xf2\x31\xe7\xdc\xa1\x2e\xf3\xe8\x4e\x62\xf3"
	"ARM_CODE2": "\x10\xf1\x10\xe7\x11\xf2\x31\xe7\xdc\xa1\x2e\xf3\xe8\x4e\x62\xf3",
	// ARMV8="\xe0\x3b\xb2\xee\x42\x00\x01\xe1\x51\xf0\x7f\xf5"
	"ARMV8": "\xe0\x3b\xb2\xee\x42\x00\x01\xe1\x51\xf0\x7f\xf5",
	// THUMB_MCLASS="\xef\xf3\x02\x80"
	"THUMB_MCLASS": "\xef\xf3\x02\x80",
	// THUMB_CODE="\x70\x47\xeb\x46\x83\xb0\xc9\x68"
	"THUMB_CODE": "\x70\x47\xeb\x46\x83\xb0\xc9\x68",
	// THUMB_CODE2="\x4f\xf0\x00\x01\xbd\xe8\x00\x88\xd1\xe8\x00\xf0"
	"THUMB_CODE2": "\x4f\xf0\x00\x01\xbd\xe8\x00\x88\xd1\xe8\x00\xf0",
	// MIPS_CODE="\x0C\x10\x00\x97\x00\x00\x00\x00\x24\x02\x00\x0c\x8f\xa2\x00\x00\x34\x21\x34\x56"
	"MIPS_CODE": "\x0C\x10\x00\x97\x00\x00\x00\x00\x24\x02\x00\x0c\x8f\xa2\x00\x00\x34\x21\x34\x56",
	// MIPS_CODE2="\x56\x34\x21\x34\xc2\x17\x01\x00"
	"MIPS_CODE2": "\x56\x34\x21\x34\xc2\x17\x01\x00",
	// MIPS_32R6M="\x00\x07\x00\x07\x00\x11\x93\x7c\x01\x8c\x8b\x7c\x00\xc7\x48\xd0"
	"MIPS_32R6M": "\x00\x07\x00\x07\x00\x11\x93\x7c\x01\x8c\x8b\x7c\x00\xc7\x48\xd0",
	// MIPS_32R6="\xec\x80\x00\x19\x7c\x43\x22\xa0"
	"MIPS_32R6": "\xec\x80\x00\x19\x7c\x43\x22\xa0",
	// ARM64_CODE="\x21\x7c\x02\x9b\x21\x7c\x00\x53\x00\x40\x21\x4b\xe1\x0b\x40\xb9"
	"ARM64_CODE": "\x21\x7c\x02\x9b\x21\x7c\x00\x53\x00\x40\x21\x4b\xe1\x0b\x40\xb9",
	// PPC_CODE="\x80\x20\x00\x00\x80\x3f\x00\x00\x10\x43\x23\x0e\xd0\x44\x00\x80\x4c\x43\x22\x02\x2d\x03\x00\x80\x7c\x43\x20\x14\x7c\x43\x20\x93\x4f\x20\x00\x21\x4c\xc8\x00\x21"
	"PPC_CODE": "\x80\x20\x00\x00\x80\x3f\x00\x00\x10\x43\x23\x0e\xd0\x44\x00\x80\x4c\x43\x22\x02\x2d\x03\x00\x80\x7c\x43\x20\x14\x7c\x43\x20\x93\x4f\x20\x00\x21\x4c\xc8\x00\x21",
	// PPC_CODE2="\x10\x60\x2a\x10\x10\x64\x28\x88\x7c\x4a\x5d\x0f"
	"PPC_CODE2": "\x10\x60\x2a\x10\x10\x64\x28\x88\x7c\x4a\x5d\x0f",
	// SPARC_CODE="\x80\xa0\x40\x02\x85\xc2\x60\x08\x85\xe8\x20\x01\x81\xe8\x00\x00\x90\x10\x20\x01\xd5\xf6\x10\x16\x21\x00\x00\x0a\x86\x00\x40\x02\x01\x00\x00\x00\x12\xbf\xff\xff\x10\xbf\xff\xff\xa0\x02\x00\x09\x0d\xbf\xff\xff\xd4\x20\x60\x00\xd4\x4e\x00\x16\x2a\xc2\x80\x03"
	"SPARC_CODE": "\x80\xa0\x40\x02\x85\xc2\x60\x08\x85\xe8\x20\x01\x81\xe8\x00\x00\x90\x10\x20\x01\xd5\xf6\x10\x16\x21\x00\x00\x0a\x86\x00\x40\x02\x01\x00\x00\x00\x12\xbf\xff\xff\x10\xbf\xff\xff\xa0\x02\x00\x09\x0d\xbf\xff\xff\xd4\x20\x60\x00\xd4\x4e\x00\x16\x2a\xc2\x80\x03",
	// SPARCV9_CODE="\x81\xa8\x0a\x24\x89\xa0\x10\x20\x89\xa0\x1a\x60\x89\xa0\x00\xe0"
	"SPARCV9_CODE": "\x81\xa8\x0a\x24\x89\xa0\x10\x20\x89\xa0\x1a\x60\x89\xa0\x00\xe0",
	// SYSZ_CODE="\xed\x00\x00\x00\x00\x1a\x5a\x0f\x1f\xff\xc2\x09\x80\x00\x00\x00\x07\xf7\xeb\x2a\xff\xff\x7f\x57\xe3\x01\xff\xff\x7f\x57\xeb\x00\xf0\x00\x00\x24\xb2\x4f\x00\x78"
	"SYSZ_CODE": "\xed\x00\x00\x00\x00\x1a\x5a\x0f\x1f\xff\xc2\x09\x80\x00\x00\x00\x07\xf7\xeb\x2a\xff\xff\x7f\x57\xe3\x01\xff\xff\x7f\x57\xeb\x00\xf0\x00\x00\x24\xb2\x4f\x00\x78",
	// XCORE_CODE="\xfe\x0f\xfe\x17\x13\x17\xc6\xfe\xec\x17\x97\xf8\xec\x4f\x1f\xfd\xec\x37\x07\xf2\x45\x5b\xf9\xfa\x02\x06\x1b\x10"
	"XCORE_CODE": "\xfe\x0f\xfe\x17\x13\x17\xc6\xfe\xec\x17\x97\xf8\xec\x4f\x1f\xfd\xec\x37\x07\xf2\x45\x5b\xf9\xfa\x02\x06\x1b\x10",
	// M68K_CODE="\xd4\x40\x87\x5a\x4e\x71\x02\xb4\xc0\xde\xc0\xde\x5c\x00\x1d\x80\x71\x12\x01\x23\xf2\x3c\x44\x22\x40\x49\x0e\x56\x54\xc5\xf2\x3c\x44\x00\x44\x7a\x00\x00\xf2\x00\x0a\x28"
	"M68K_CODE": "\xd4\x40\x87\x5a\x4e\x71\x02\xb4\xc0\xde\xc0\xde\x5c\x00\x1d\x80\x71\x12\x01\x23\xf2\x3c\x44\x22\x40\x49\x0e\x56\x54\xc5\xf2\x3c\x44\x00\x44\x7a\x00\x00\xf2\x00\x0a\x28",
	// TMS320C64X_CODE="\x01\xac\x88\x40\x81\xac\x88\x43\x00\x00\x00\x00\x02\x90\x32\x96\x02\x80\x46\x9e\x05\x3c\x83\xe6\x0b\x0c\x8b\x24"
	"TMS320C64X_CODE": "\x01\xac\x88\x40\x81\xac\x88\x43\x00\x00\x00\x00\x02\x90\x32\x96\x02\x80\x46\x9e\x05\x3c\x83\xe6\x0b\x0c\x8b\x24",
	// M680X_CODE="\x06\x10\x19\x1a\x55\x1e\x01\x23\xe9\x31\x06\x34\x55\xa6\x81\xa7\x89\x7f\xff\xa6\x9d\x10\x00\xa7\x91\xa6\x9f\x10\x00\x11\xac\x99\x10\x00\x39"
	"M680X_CODE": "\x06\x10\x19\x1a\x55\x1e\x01\x23\xe9\x31\x06\x34\x55\xa6\x81\xa7\x89\x7f\xff\xa6\x9d\x10\x00\xa7\x91\xa6\x9f\x10\x00\x11\xac\x99\x10\x00\x39",
	// EVM_CODE="\x60\x61"
	"EVM_CODE": "\x60\x61",
	// WASM_CODE="\x20\x00\x20\x01\x41\x20\x10\xc9\x01\x45\x0b"
	"WASM_CODE": "\x20\x00\x20\x01\x41\x20\x10\xc9\x01\x45\x0b",
	// MOS65XX_CODE="\x0d\x34\x12\x00\x81\x65\x6c\x01\x00\x85\xFF\x10\x00\x19\x42\x42\x00\x49\x42"
	"MOS65XX_CODE": "\x0d\x34\x12\x00\x81\x65\x6c\x01\x00\x85\xFF\x10\x00\x19\x42\x42\x00\x49\x42",
	// EBPF_CODE="\x97\x09\x00\x00\x37\x13\x03\x00\xdc\x02\x00\x00\x20\x00\x00\x00\x30\x00\x00\x00\x00\x00\x00\x00\xdb\x3a\x00\x01\x00\x00\x00\x00\x84\x02\x00\x00\x00\x00\x00\x00\x6d\x33\x17\x02\x00\x00\x00\x00"
	"EBPF_CODE": "\x97\x09\x00\x00\x37\x13\x03\x00\xdc\x02\x00\x00\x20\x00\x00\x00\x30\x00\x00\x00\x00\x00\x00\x00\xdb\x3a\x00\x01\x00\x00\x00\x00\x84\x02\x00\x00\x00\x00\x00\x00\x6d\x33\x17\x02\x00\x00\x00\x00",
	// RISCV_CODE32="\x37\x34\x00\x00\x97\x82\x00\x00\xef\x00\x80\x00\xef\xf0\x1f\xff\xe7\x00\x45\x00\xe7\x00\xc0\xff\x63\x05\x41\x00\xe3\x9d\x61\xfe\x63\xca\x93\x00\x63\x53\xb5\x00\x63\x65\xd6\x00\x63\x76\xf7\x00\x03\x88\x18\x00\x03\x99\x49\x00\x03\xaa\x6a\x00\x03\xcb\x2b\x01\x03\xdc\x8c\x01\x23\x86\xad\x03\x23\x9a\xce\x03\x23\x8f\xef\x01\x93\x00\xe0\x00\x13\xa1\x01\x01\x13\xb2\x02\x7d\x13\xc3\x03\xdd\x13\xe4\xc4\x12\x13\xf5\x85\x0c\x13\x96\xe6\x01\x13\xd7\x97\x01\x13\xd8\xf8\x40\x33\x89\x49\x01\xb3\x0a\x7b\x41\x33\xac\xac\x01\xb3\x3d\xde\x01\x33\xd2\x62\x40\xb3\x43\x94\x00\x33\xe5\xc5\x00\xb3\x76\xf7\x00\xb3\x54\x39\x01\xb3\x50\x31\x00\x33\x9f\x0f\x00"
	"RISCV_CODE32": "\x37\x34\x00\x00\x97\x82\x00\x00\xef\x00\x80\x00\xef\xf0\x1f\xff\xe7\x00\x45\x00\xe7\x00\xc0\xff\x63\x05\x41\x00\xe3\x9d\x61\xfe\x63\xca\x93\x00\x63\x53\xb5\x00\x63\x65\xd6\x00\x63\x76\xf7\x00\x03\x88\x18\x00\x03\x99\x49\x00\x03\xaa\x6a\x00\x03\xcb\x2b\x01\x03\xdc\x8c\x01\x23\x86\xad\x03\x23\x9a\xce\x03\x23\x8f\xef\x01\x93\x00\xe0\x00\x13\xa1\x01\x01\x13\xb2\x02\x7d\x13\xc3\x03\xdd\x13\xe4\xc4\x12\x13\xf5\x85\x0c\x13\x96\xe6\x01\x13\xd7\x97\x01\x13\xd8\xf8\x40\x33\x89\x49\x01\xb3\x0a\x7b\x41\x33\xac\xac\x01\xb3\x3d\xde\x01\x33\xd2\x62\x40\xb3\x43\x94\x00\x33\xe5\xc5\x00\xb3\x76\xf7\x00\xb3\x54\x39\x01\xb3\x50\x31\x00\x33\x9f\x0f\x00",
	// RISCV_CODE64="\x13\x04\xa8\x7a"
	"RISCV_CODE64": "\x13\x04\xa8\x7a",
}

//=== Found platform entries in capstone/tests/test_basic.c ===

var basicPlatforms = []platform{
	{
		// -------------------------------------------
		// Platform Entry (line 105):
		// Architecture: CS_ARCH_X86
		// Mode: CS_MODE_16
		// Code: X86_CODE16
		// Comment: "X86 16bit (Intel syntax)"
		arch:    CS_ARCH_X86,
		mode:    CS_MODE_16,
		code:    basicCodes["X86_CODE16"],
		comment: "X86 16bit (Intel syntax)",
		options: []option{
			{CS_OPT_DETAIL, CS_OPT_ON},
		},
	},
	{
		// -------------------------------------------
		// Platform Entry (line 112):
		// Architecture: CS_ARCH_X86
		// Mode: CS_MODE_32
		// Code: X86_CODE32
		// Comment: "X86 32bit (ATT syntax)"
		//  Options:
		//    CS_OPT_SYNTAX: CS_OPT_SYNTAX_ATT
		arch:    CS_ARCH_X86,
		mode:    CS_MODE_32,
		code:    basicCodes["X86_CODE32"],
		comment: "X86 32bit (ATT syntax)",
		options: []option{
			{CS_OPT_DETAIL, CS_OPT_ON},
			{CS_OPT_SYNTAX, CS_OPT_SYNTAX_ATT},
		},
	},
	{
		// -------------------------------------------
		// Platform Entry (line 121):
		// Architecture: CS_ARCH_X86
		// Mode: CS_MODE_32
		// Code: X86_CODE32
		// Comment: "X86 32 (Intel syntax)"
		arch:    CS_ARCH_X86,
		mode:    CS_MODE_32,
		code:    basicCodes["X86_CODE32"],
		comment: "X86 32 (Intel syntax)",
		options: []option{
			{CS_OPT_DETAIL, CS_OPT_ON},
		},
	},
	{
		// -------------------------------------------
		// Platform Entry (line 128):
		// Architecture: CS_ARCH_X86
		// Mode: CS_MODE_32
		// Code: X86_CODE32
		// Comment: "X86 32 (MASM syntax)"
		//  Options:
		//    CS_OPT_SYNTAX: CS_OPT_SYNTAX_MASM
		arch:    CS_ARCH_X86,
		mode:    CS_MODE_32,
		code:    basicCodes["X86_CODE32"],
		comment: "X86 32 (MASM syntax)",
		options: []option{
			{CS_OPT_DETAIL, CS_OPT_ON},
			{CS_OPT_SYNTAX, CS_OPT_SYNTAX_MASM},
		},
	},
	{
		// -------------------------------------------
		// Platform Entry (line 137):
		// Architecture: CS_ARCH_X86
		// Mode: CS_MODE_64
		// Code: X86_CODE64
		// Comment: "X86 64 (Intel syntax)"
		arch:    CS_ARCH_X86,
		mode:    CS_MODE_64,
		code:    basicCodes["X86_CODE64"],
		comment: "X86 64 (Intel syntax)",
		options: []option{
			{CS_OPT_DETAIL, CS_OPT_ON},
		},
	},
	{
		// -------------------------------------------
		// Platform Entry (line 146):
		// Architecture: CS_ARCH_ARM
		// Mode: CS_MODE_ARM
		// Code: ARM_CODE
		// Comment: "ARM"
		arch:    CS_ARCH_ARM,
		mode:    CS_MODE_ARM,
		code:    basicCodes["ARM_CODE"],
		comment: "ARM",
		options: []option{
			{CS_OPT_DETAIL, CS_OPT_ON},
		},
	},
	{
		// -------------------------------------------
		// Platform Entry (line 153):
		// Architecture: CS_ARCH_ARM
		// Mode: CS_MODE_THUMB
		// Code: THUMB_CODE2
		// Comment: "THUMB-2"
		arch:    CS_ARCH_ARM,
		mode:    CS_MODE_THUMB,
		code:    basicCodes["THUMB_CODE2"],
		comment: "THUMB-2",
		options: []option{
			{CS_OPT_DETAIL, CS_OPT_ON},
		},
	},
	{
		// -------------------------------------------
		// Platform Entry (line 160):
		// Architecture: CS_ARCH_ARM
		// Mode: CS_MODE_ARM
		// Code: ARM_CODE2
		// Comment: "ARM: Cortex-A15 + NEON"
		arch:    CS_ARCH_ARM,
		mode:    CS_MODE_ARM,
		code:    basicCodes["ARM_CODE2"],
		comment: "ARM: Cortex-A15 + NEON",
		options: []option{
			{CS_OPT_DETAIL, CS_OPT_ON},
		},
	},
	{
		// -------------------------------------------
		// Platform Entry (line 167):
		// Architecture: CS_ARCH_ARM
		// Mode: CS_MODE_THUMB
		// Code: THUMB_CODE
		// Comment: "THUMB"
		arch:    CS_ARCH_ARM,
		mode:    CS_MODE_THUMB,
		code:    basicCodes["THUMB_CODE"],
		comment: "THUMB",
		options: []option{
			{CS_OPT_DETAIL, CS_OPT_ON},
		},
	},
	{
		// -------------------------------------------
		// Platform Entry (line 174):
		// Architecture: CS_ARCH_ARM
		// Mode: (CS_MODE_THUMB + CS_MODE_MCLASS)
		// Code: THUMB_MCLASS
		// Comment: "Thumb-MClass"
		arch:    CS_ARCH_ARM,
		mode:    (CS_MODE_THUMB + CS_MODE_MCLASS),
		code:    basicCodes["THUMB_MCLASS"],
		comment: "Thumb-MClass",
		options: []option{
			{CS_OPT_DETAIL, CS_OPT_ON},
		},
	},
	{
		// -------------------------------------------
		// Platform Entry (line 181):
		// Architecture: CS_ARCH_ARM
		// Mode: (CS_MODE_ARM + CS_MODE_V8)
		// Code: ARMV8
		// Comment: "Arm-V8"
		arch:    CS_ARCH_ARM,
		mode:    (CS_MODE_ARM + CS_MODE_V8),
		code:    basicCodes["ARMV8"],
		comment: "Arm-V8",
		options: []option{
			{CS_OPT_DETAIL, CS_OPT_ON},
		},
	},
	{
		// -------------------------------------------
		// Platform Entry (line 190):
		// Architecture: CS_ARCH_MIPS
		// Mode: (CS_MODE_MIPS32 + CS_MODE_BIG_ENDIAN)
		// Code: MIPS_CODE
		// Comment: "MIPS-32 (Big-endian)"
		arch:    CS_ARCH_MIPS,
		mode:    (CS_MODE_MIPS32 + CS_MODE_BIG_ENDIAN),
		code:    basicCodes["MIPS_CODE"],
		comment: "MIPS-32 (Big-endian)",
		options: []option{
			{CS_OPT_DETAIL, CS_OPT_ON},
		},
	},
	{
		// -------------------------------------------
		// Platform Entry (line 197):
		// Architecture: CS_ARCH_MIPS
		// Mode: (CS_MODE_MIPS64 + CS_MODE_LITTLE_ENDIAN)
		// Code: MIPS_CODE2
		// Comment: "MIPS-64-EL (Little-endian)"
		arch:    CS_ARCH_MIPS,
		mode:    (CS_MODE_MIPS64 + CS_MODE_LITTLE_ENDIAN),
		code:    basicCodes["MIPS_CODE2"],
		comment: "MIPS-64-EL (Little-endian)",
		options: []option{
			{CS_OPT_DETAIL, CS_OPT_ON},
		},
	},
	{
		// -------------------------------------------
		// Platform Entry (line 204):
		// Architecture: CS_ARCH_MIPS
		// Mode: (CS_MODE_MIPS32R6 + CS_MODE_MICRO + CS_MODE_BIG_ENDIAN)
		// Code: MIPS_32R6M
		// Comment: "MIPS-32R6 | Micro (Big-endian)"
		arch:    CS_ARCH_MIPS,
		mode:    (CS_MODE_MIPS32R6 + CS_MODE_MICRO + CS_MODE_BIG_ENDIAN),
		code:    basicCodes["MIPS_32R6M"],
		comment: "MIPS-32R6 | Micro (Big-endian)",
		options: []option{
			{CS_OPT_DETAIL, CS_OPT_ON},
		},
	},
	{
		// -------------------------------------------
		// Platform Entry (line 211):
		// Architecture: CS_ARCH_MIPS
		// Mode: (CS_MODE_MIPS32R6 + CS_MODE_BIG_ENDIAN)
		// Code: MIPS_32R6
		// Comment: "MIPS-32R6 (Big-endian)"
		arch:    CS_ARCH_MIPS,
		mode:    (CS_MODE_MIPS32R6 + CS_MODE_BIG_ENDIAN),
		code:    basicCodes["MIPS_32R6"],
		comment: "MIPS-32R6 (Big-endian)",
		options: []option{
			{CS_OPT_DETAIL, CS_OPT_ON},
		},
	},
	{
		// -------------------------------------------
		// Platform Entry (line 220):
		// Architecture: CS_ARCH_ARM64
		// Mode: CS_MODE_ARM
		// Code: ARM64_CODE
		// Comment: "ARM-64"
		arch:    CS_ARCH_ARM64,
		mode:    CS_MODE_ARM,
		code:    basicCodes["ARM64_CODE"],
		comment: "ARM-64",
		options: []option{
			{CS_OPT_DETAIL, CS_OPT_ON},
		},
	},
	{
		// -------------------------------------------
		// Platform Entry (line 229):
		// Architecture: CS_ARCH_PPC
		// Mode: CS_MODE_BIG_ENDIAN
		// Code: PPC_CODE
		// Comment: "PPC-64"
		arch:    CS_ARCH_PPC,
		mode:    CS_MODE_BIG_ENDIAN,
		code:    basicCodes["PPC_CODE"],
		comment: "PPC-64",
		options: []option{
			{CS_OPT_DETAIL, CS_OPT_ON},
		},
	},
	{
		// -------------------------------------------
		// Platform Entry (line 236):
		// Architecture: CS_ARCH_PPC
		// Mode: CS_MODE_BIG_ENDIAN
		// Code: PPC_CODE
		// Comment: "PPC-64
		//  Options:
		//    print register with number only": CS_OPT_SYNTAX
		arch:    CS_ARCH_PPC,
		mode:    CS_MODE_BIG_ENDIAN,
		code:    basicCodes["PPC_CODE"],
		comment: "PPC-64",
		options: []option{
			{CS_OPT_DETAIL, CS_OPT_ON},
			{CS_OPT_SYNTAX, CS_OPT_SYNTAX_NOREGNAME},
		},
	},
	{
		// -------------------------------------------
		// Platform Entry (line 245):
		// Architecture: CS_ARCH_PPC
		// Mode: CS_MODE_BIG_ENDIAN + CS_MODE_QPX
		// Code: PPC_CODE2
		// Comment: "PPC-64 + QPX"
		arch:    CS_ARCH_PPC,
		mode:    CS_MODE_BIG_ENDIAN + CS_MODE_QPX,
		code:    basicCodes["PPC_CODE2"],
		comment: "PPC-64 + QPX",
		options: []option{
			{CS_OPT_DETAIL, CS_OPT_ON},
		},
	},
	{
		// -------------------------------------------
		// Platform Entry (line 254):
		// Architecture: CS_ARCH_SPARC
		// Mode: CS_MODE_BIG_ENDIAN
		// Code: SPARC_CODE
		// Comment: "Sparc"
		arch:    CS_ARCH_SPARC,
		mode:    CS_MODE_BIG_ENDIAN,
		code:    basicCodes["SPARC_CODE"],
		comment: "Sparc",
		options: []option{
			{CS_OPT_DETAIL, CS_OPT_ON},
		},
	},
	{
		// -------------------------------------------
		// Platform Entry (line 261):
		// Architecture: CS_ARCH_SPARC
		// Mode: (CS_MODE_BIG_ENDIAN + CS_MODE_V9)
		// Code: SPARCV9_CODE
		// Comment: "SparcV9"
		arch:    CS_ARCH_SPARC,
		mode:    (CS_MODE_BIG_ENDIAN + CS_MODE_V9),
		code:    basicCodes["SPARCV9_CODE"],
		comment: "SparcV9",
		options: []option{
			{CS_OPT_DETAIL, CS_OPT_ON},
		},
	},
	{
		// -------------------------------------------
		// Platform Entry (line 270):
		// Architecture: CS_ARCH_SYSZ
		// Mode: 0
		// Code: SYSZ_CODE
		// Comment: "SystemZ"
		arch:    CS_ARCH_SYSZ,
		mode:    0,
		code:    basicCodes["SYSZ_CODE"],
		comment: "SystemZ",
		options: []option{
			{CS_OPT_DETAIL, CS_OPT_ON},
		},
	},
	{
		// -------------------------------------------
		// Platform Entry (line 279):
		// Architecture: CS_ARCH_XCORE
		// Mode: 0
		// Code: XCORE_CODE
		// Comment: "XCore"
		arch:    CS_ARCH_XCORE,
		mode:    0,
		code:    basicCodes["XCORE_CODE"],
		comment: "XCore",
		options: []option{
			{CS_OPT_DETAIL, CS_OPT_ON},
		},
	},
	{
		// -------------------------------------------
		// Platform Entry (line 288):
		// Architecture: CS_ARCH_M68K
		// Mode: (CS_MODE_BIG_ENDIAN | CS_MODE_M68K_040)
		// Code: M68K_CODE
		// Comment: "M68K"
		arch:    CS_ARCH_M68K,
		mode:    (CS_MODE_BIG_ENDIAN | CS_MODE_M68K_040),
		code:    basicCodes["M68K_CODE"],
		comment: "M68K",
		options: []option{
			{CS_OPT_DETAIL, CS_OPT_ON},
		},
	},
	{
		// -------------------------------------------
		// Platform Entry (line 297):
		// Architecture: CS_ARCH_TMS320C64X
		// Mode: 0
		// Code: TMS320C64X_CODE
		// Comment: "TMS320C64x"
		arch:    CS_ARCH_TMS320C64X,
		mode:    0,
		code:    basicCodes["TMS320C64X_CODE"],
		comment: "TMS320C64x",
		options: []option{
			{CS_OPT_DETAIL, CS_OPT_ON},
		},
	},
	{
		// -------------------------------------------
		// Platform Entry (line 306):
		// Architecture: CS_ARCH_M680X
		// Mode: (CS_MODE_M680X_6809)
		// Code: M680X_CODE
		// Comment: "M680X_M6809"
		arch:    CS_ARCH_M680X,
		mode:    (CS_MODE_M680X_6809),
		code:    basicCodes["M680X_CODE"],
		comment: "M680X_M6809",
		options: []option{
			{CS_OPT_DETAIL, CS_OPT_ON},
		},
	},
	{
		// -------------------------------------------
		// Platform Entry (line 315):
		// Architecture: CS_ARCH_EVM
		// Mode: 0
		// Code: EVM_CODE
		// Comment: "EVM"
		arch:    CS_ARCH_EVM,
		mode:    0,
		code:    basicCodes["EVM_CODE"],
		comment: "EVM",
		options: []option{
			{CS_OPT_DETAIL, CS_OPT_ON},
		},
	},
	{
		// -------------------------------------------
		// Platform Entry (line 324):
		// Architecture: CS_ARCH_WASM
		// Mode: 0
		// Code: WASM_CODE
		// Comment: "WASM"
		arch:    CS_ARCH_WASM,
		mode:    0,
		code:    basicCodes["WASM_CODE"],
		comment: "WASM",
		options: []option{
			{CS_OPT_DETAIL, CS_OPT_ON},
		},
	},
	{
		// -------------------------------------------
		// Platform Entry (line 333):
		// Architecture: CS_ARCH_MOS65XX
		// Mode: 0
		// Code: MOS65XX_CODE
		// Comment: "MOS65XX"
		arch:    CS_ARCH_MOS65XX,
		mode:    0,
		code:    basicCodes["MOS65XX_CODE"],
		comment: "MOS65XX",
		options: []option{
			{CS_OPT_DETAIL, CS_OPT_ON},
		},
	},
	{
		// -------------------------------------------
		// Platform Entry (line 342):
		// Architecture: CS_ARCH_BPF
		// Mode: CS_MODE_LITTLE_ENDIAN | CS_MODE_BPF_EXTENDED
		// Code: EBPF_CODE
		// Comment: "eBPF"
		arch:    CS_ARCH_BPF,
		mode:    CS_MODE_LITTLE_ENDIAN | CS_MODE_BPF_EXTENDED,
		code:    basicCodes["EBPF_CODE"],
		comment: "eBPF",
		options: []option{
			{CS_OPT_DETAIL, CS_OPT_ON},
		},
	},
	{
		// -------------------------------------------
		// Platform Entry (line 351):
		// Architecture: CS_ARCH_RISCV
		// Mode: CS_MODE_RISCV32
		// Code: RISCV_CODE32
		// Comment: "RISCV32"
		arch:    CS_ARCH_RISCV,
		mode:    CS_MODE_RISCV32,
		code:    basicCodes["RISCV_CODE32"],
		comment: "RISCV32",
		options: []option{
			{CS_OPT_DETAIL, CS_OPT_ON},
		},
	},
	{
		// -------------------------------------------
		// Platform Entry (line 358):
		// Architecture: CS_ARCH_RISCV
		// Mode: CS_MODE_RISCV64
		// Code: RISCV_CODE64
		// Comment: "RISCV64"
		arch:    CS_ARCH_RISCV,
		mode:    CS_MODE_RISCV64,
		code:    basicCodes["RISCV_CODE64"],
		comment: "RISCV64",
		options: []option{
			{CS_OPT_DETAIL, CS_OPT_ON},
		},
	},
}
