/*
Gapstone is a Go binding for the Capstone disassembly library.
For examples,try reading the *_test.go files.

	Library Author: Nguyen Anh Quynh
	Binding Author: Ben Nagy
	License: BSD style - see LICENSE file for details
    (c) 2013 COSEINC. All Rights Reserved.

    THIS FILE WAS AUTO-GENERATED -- DO NOT EDIT!
    Command: ./genresorces.sh capstone/tests/test_detail.c
    Created at: 2025-04-25 12:56:14

*/

package gapstone

//=== Found #define values in capstone/tests/test_detail.c ===

var detailCodes = map[string]string{
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
	// THUMB_MCLASS="\xef\xf3\x02\x80"
	"THUMB_MCLASS": "\xef\xf3\x02\x80",
	// ARMV8="\xe0\x3b\xb2\xee\x42\x00\x01\xe1\x51\xf0\x7f\xf5"
	"ARMV8": "\xe0\x3b\xb2\xee\x42\x00\x01\xe1\x51\xf0\x7f\xf5",
	// MIPS_CODE="\x0C\x10\x00\x97\x00\x00\x00\x00\x24\x02\x00\x0c\x8f\xa2\x00\x00\x34\x21\x34\x56\x00\x80\x04\x08"
	"MIPS_CODE": "\x0C\x10\x00\x97\x00\x00\x00\x00\x24\x02\x00\x0c\x8f\xa2\x00\x00\x34\x21\x34\x56\x00\x80\x04\x08",
	// MIPS_CODE2="\x56\x34\x21\x34\xc2\x17\x01\x00"
	"MIPS_CODE2": "\x56\x34\x21\x34\xc2\x17\x01\x00",
	// MIPS_32R6M="\x00\x07\x00\x07\x00\x11\x93\x7c\x01\x8c\x8b\x7c\x00\xc7\x48\xd0"
	"MIPS_32R6M": "\x00\x07\x00\x07\x00\x11\x93\x7c\x01\x8c\x8b\x7c\x00\xc7\x48\xd0",
	// MIPS_32R6="\xec\x80\x00\x19\x7c\x43\x22\xa0"
	"MIPS_32R6": "\xec\x80\x00\x19\x7c\x43\x22\xa0",
	// ARM64_CODE="\x09\x00\x38\xd5\xbf\x40\x00\xd5\x0c\x05\x13\xd5\x20\x50\x02\x0e\x20\xe4\x3d\x0f\x00\x18\xa0\x5f\xa2\x00\xae\x9e\x9f\x37\x03\xd5\xbf\x33\x03\xd5\xdf\x3f\x03\xd5\x21\x7c\x02\x9b\x21\x7c\x00\x53\x00\x40\x21\x4b\xe1\x0b\x40\xb9\x20\x04\x81\xda\x20\x08\x02\x8b\x10\x5b\xe8\x3c"
	"ARM64_CODE": "\x09\x00\x38\xd5\xbf\x40\x00\xd5\x0c\x05\x13\xd5\x20\x50\x02\x0e\x20\xe4\x3d\x0f\x00\x18\xa0\x5f\xa2\x00\xae\x9e\x9f\x37\x03\xd5\xbf\x33\x03\xd5\xdf\x3f\x03\xd5\x21\x7c\x02\x9b\x21\x7c\x00\x53\x00\x40\x21\x4b\xe1\x0b\x40\xb9\x20\x04\x81\xda\x20\x08\x02\x8b\x10\x5b\xe8\x3c",
	// PPC_CODE="\x80\x20\x00\x00\x80\x3f\x00\x00\x10\x43\x23\x0e\xd0\x44\x00\x80\x4c\x43\x22\x02\x2d\x03\x00\x80\x7c\x43\x20\x14\x7c\x43\x20\x93\x4f\x20\x00\x21\x4c\xc8\x00\x21\x40\x82\x00\x14"
	"PPC_CODE": "\x80\x20\x00\x00\x80\x3f\x00\x00\x10\x43\x23\x0e\xd0\x44\x00\x80\x4c\x43\x22\x02\x2d\x03\x00\x80\x7c\x43\x20\x14\x7c\x43\x20\x93\x4f\x20\x00\x21\x4c\xc8\x00\x21\x40\x82\x00\x14",
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
	// M680X_CODE="\x06\x10\x19\x1a\x55\x1e\x01\x23\xe9\x31\x06\x34\x55\xa6\x81\xa7\x89\x7f\xff\xa6\x9d\x10\x00\xa7\x91\xa6\x9f\x10\x00\x11\xac\x99\x10\x00\x39"
	"M680X_CODE": "\x06\x10\x19\x1a\x55\x1e\x01\x23\xe9\x31\x06\x34\x55\xa6\x81\xa7\x89\x7f\xff\xa6\x9d\x10\x00\xa7\x91\xa6\x9f\x10\x00\x11\xac\x99\x10\x00\x39",
	// MOS65XX_CODE="\x0A\x00\xFE\x34\x12\xD0\xFF\xEA\x19\x56\x34\x46\x80"
	"MOS65XX_CODE": "\x0A\x00\xFE\x34\x12\xD0\xFF\xEA\x19\x56\x34\x46\x80",
	// EBPF_CODE="\x97\x09\x00\x00\x37\x13\x03\x00\xdc\x02\x00\x00\x20\x00\x00\x00\x30\x00\x00\x00\x00\x00\x00\x00\xdb\x3a\x00\x01\x00\x00\x00\x00\x84\x02\x00\x00\x00\x00\x00\x00\x6d\x33\x17\x02\x00\x00\x00\x00"
	"EBPF_CODE": "\x97\x09\x00\x00\x37\x13\x03\x00\xdc\x02\x00\x00\x20\x00\x00\x00\x30\x00\x00\x00\x00\x00\x00\x00\xdb\x3a\x00\x01\x00\x00\x00\x00\x84\x02\x00\x00\x00\x00\x00\x00\x6d\x33\x17\x02\x00\x00\x00\x00",
}

//=== Found platform entries in capstone/tests/test_detail.c ===

var detailPlatforms = []platform{
	{
		// -------------------------------------------
		// Platform Entry (line 82):
		// Architecture: CS_ARCH_X86
		// Mode: CS_MODE_16
		// Code: X86_CODE16
		// Comment: "X86 16bit (Intel syntax)"
		arch:    CS_ARCH_X86,
		mode:    CS_MODE_16,
		code:    detailCodes["X86_CODE16"],
		comment: "X86 16bit (Intel syntax)",
		options: []option{
			{CS_OPT_DETAIL, CS_OPT_ON},
		},
	},
	{
		// -------------------------------------------
		// Platform Entry (line 89):
		// Architecture: CS_ARCH_X86
		// Mode: CS_MODE_32
		// Code: X86_CODE32
		// Comment: "X86 32bit (ATT syntax)"
		//  Options:
		//    CS_OPT_SYNTAX: CS_OPT_SYNTAX_ATT
		arch:    CS_ARCH_X86,
		mode:    CS_MODE_32,
		code:    detailCodes["X86_CODE32"],
		comment: "X86 32bit (ATT syntax)",
		options: []option{
			{CS_OPT_DETAIL, CS_OPT_ON},
			{CS_OPT_SYNTAX, CS_OPT_SYNTAX_ATT},
		},
	},
	{
		// -------------------------------------------
		// Platform Entry (line 98):
		// Architecture: CS_ARCH_X86
		// Mode: CS_MODE_32
		// Code: X86_CODE32
		// Comment: "X86 32 (Intel syntax)"
		arch:    CS_ARCH_X86,
		mode:    CS_MODE_32,
		code:    detailCodes["X86_CODE32"],
		comment: "X86 32 (Intel syntax)",
		options: []option{
			{CS_OPT_DETAIL, CS_OPT_ON},
		},
	},
	{
		// -------------------------------------------
		// Platform Entry (line 105):
		// Architecture: CS_ARCH_X86
		// Mode: CS_MODE_64
		// Code: X86_CODE64
		// Comment: "X86 64 (Intel syntax)"
		arch:    CS_ARCH_X86,
		mode:    CS_MODE_64,
		code:    detailCodes["X86_CODE64"],
		comment: "X86 64 (Intel syntax)",
		options: []option{
			{CS_OPT_DETAIL, CS_OPT_ON},
		},
	},
	{
		// -------------------------------------------
		// Platform Entry (line 114):
		// Architecture: CS_ARCH_ARM
		// Mode: CS_MODE_ARM
		// Code: ARM_CODE
		// Comment: "ARM"
		arch:    CS_ARCH_ARM,
		mode:    CS_MODE_ARM,
		code:    detailCodes["ARM_CODE"],
		comment: "ARM",
		options: []option{
			{CS_OPT_DETAIL, CS_OPT_ON},
		},
	},
	{
		// -------------------------------------------
		// Platform Entry (line 121):
		// Architecture: CS_ARCH_ARM
		// Mode: CS_MODE_THUMB
		// Code: THUMB_CODE2
		// Comment: "THUMB-2"
		arch:    CS_ARCH_ARM,
		mode:    CS_MODE_THUMB,
		code:    detailCodes["THUMB_CODE2"],
		comment: "THUMB-2",
		options: []option{
			{CS_OPT_DETAIL, CS_OPT_ON},
		},
	},
	{
		// -------------------------------------------
		// Platform Entry (line 128):
		// Architecture: CS_ARCH_ARM
		// Mode: CS_MODE_ARM
		// Code: ARM_CODE2
		// Comment: "ARM: Cortex-A15 + NEON"
		arch:    CS_ARCH_ARM,
		mode:    CS_MODE_ARM,
		code:    detailCodes["ARM_CODE2"],
		comment: "ARM: Cortex-A15 + NEON",
		options: []option{
			{CS_OPT_DETAIL, CS_OPT_ON},
		},
	},
	{
		// -------------------------------------------
		// Platform Entry (line 135):
		// Architecture: CS_ARCH_ARM
		// Mode: CS_MODE_THUMB
		// Code: THUMB_CODE
		// Comment: "THUMB"
		arch:    CS_ARCH_ARM,
		mode:    CS_MODE_THUMB,
		code:    detailCodes["THUMB_CODE"],
		comment: "THUMB",
		options: []option{
			{CS_OPT_DETAIL, CS_OPT_ON},
		},
	},
	{
		// -------------------------------------------
		// Platform Entry (line 142):
		// Architecture: CS_ARCH_ARM
		// Mode: (CS_MODE_THUMB + CS_MODE_MCLASS)
		// Code: THUMB_MCLASS
		// Comment: "Thumb-MClass"
		arch:    CS_ARCH_ARM,
		mode:    (CS_MODE_THUMB + CS_MODE_MCLASS),
		code:    detailCodes["THUMB_MCLASS"],
		comment: "Thumb-MClass",
		options: []option{
			{CS_OPT_DETAIL, CS_OPT_ON},
		},
	},
	{
		// -------------------------------------------
		// Platform Entry (line 149):
		// Architecture: CS_ARCH_ARM
		// Mode: (CS_MODE_ARM + CS_MODE_V8)
		// Code: ARMV8
		// Comment: "Arm-V8"
		arch:    CS_ARCH_ARM,
		mode:    (CS_MODE_ARM + CS_MODE_V8),
		code:    detailCodes["ARMV8"],
		comment: "Arm-V8",
		options: []option{
			{CS_OPT_DETAIL, CS_OPT_ON},
		},
	},
	{
		// -------------------------------------------
		// Platform Entry (line 158):
		// Architecture: CS_ARCH_MIPS
		// Mode: (CS_MODE_MIPS32 + CS_MODE_BIG_ENDIAN)
		// Code: MIPS_CODE
		// Comment: "MIPS-32 (Big-endian)"
		arch:    CS_ARCH_MIPS,
		mode:    (CS_MODE_MIPS32 + CS_MODE_BIG_ENDIAN),
		code:    detailCodes["MIPS_CODE"],
		comment: "MIPS-32 (Big-endian)",
		options: []option{
			{CS_OPT_DETAIL, CS_OPT_ON},
		},
	},
	{
		// -------------------------------------------
		// Platform Entry (line 165):
		// Architecture: CS_ARCH_MIPS
		// Mode: (CS_MODE_MIPS64 + CS_MODE_LITTLE_ENDIAN)
		// Code: MIPS_CODE2
		// Comment: "MIPS-64-EL (Little-endian)"
		arch:    CS_ARCH_MIPS,
		mode:    (CS_MODE_MIPS64 + CS_MODE_LITTLE_ENDIAN),
		code:    detailCodes["MIPS_CODE2"],
		comment: "MIPS-64-EL (Little-endian)",
		options: []option{
			{CS_OPT_DETAIL, CS_OPT_ON},
		},
	},
	{
		// -------------------------------------------
		// Platform Entry (line 172):
		// Architecture: CS_ARCH_MIPS
		// Mode: (CS_MODE_MIPS32R6 + CS_MODE_MICRO + CS_MODE_BIG_ENDIAN)
		// Code: MIPS_32R6M
		// Comment: "MIPS-32R6 | Micro (Big-endian)"
		arch:    CS_ARCH_MIPS,
		mode:    (CS_MODE_MIPS32R6 + CS_MODE_MICRO + CS_MODE_BIG_ENDIAN),
		code:    detailCodes["MIPS_32R6M"],
		comment: "MIPS-32R6 | Micro (Big-endian)",
		options: []option{
			{CS_OPT_DETAIL, CS_OPT_ON},
		},
	},
	{
		// -------------------------------------------
		// Platform Entry (line 179):
		// Architecture: CS_ARCH_MIPS
		// Mode: (CS_MODE_MIPS32R6 + CS_MODE_BIG_ENDIAN)
		// Code: MIPS_32R6
		// Comment: "MIPS-32R6 (Big-endian)"
		arch:    CS_ARCH_MIPS,
		mode:    (CS_MODE_MIPS32R6 + CS_MODE_BIG_ENDIAN),
		code:    detailCodes["MIPS_32R6"],
		comment: "MIPS-32R6 (Big-endian)",
		options: []option{
			{CS_OPT_DETAIL, CS_OPT_ON},
		},
	},
	{
		// -------------------------------------------
		// Platform Entry (line 188):
		// Architecture: CS_ARCH_ARM64
		// Mode: CS_MODE_ARM
		// Code: ARM64_CODE
		// Comment: "ARM-64"
		arch:    CS_ARCH_ARM64,
		mode:    CS_MODE_ARM,
		code:    detailCodes["ARM64_CODE"],
		comment: "ARM-64",
		options: []option{
			{CS_OPT_DETAIL, CS_OPT_ON},
		},
	},
	{
		// -------------------------------------------
		// Platform Entry (line 197):
		// Architecture: CS_ARCH_PPC
		// Mode: CS_MODE_BIG_ENDIAN
		// Code: PPC_CODE
		// Comment: "PPC-64"
		arch:    CS_ARCH_PPC,
		mode:    CS_MODE_BIG_ENDIAN,
		code:    detailCodes["PPC_CODE"],
		comment: "PPC-64",
		options: []option{
			{CS_OPT_DETAIL, CS_OPT_ON},
		},
	},
	{
		// -------------------------------------------
		// Platform Entry (line 204):
		// Architecture: CS_ARCH_PPC
		// Mode: CS_MODE_BIG_ENDIAN + CS_MODE_QPX
		// Code: PPC_CODE2
		// Comment: "PPC-64 + QPX"
		arch:    CS_ARCH_PPC,
		mode:    CS_MODE_BIG_ENDIAN + CS_MODE_QPX,
		code:    detailCodes["PPC_CODE2"],
		comment: "PPC-64 + QPX",
		options: []option{
			{CS_OPT_DETAIL, CS_OPT_ON},
		},
	},
	{
		// -------------------------------------------
		// Platform Entry (line 213):
		// Architecture: CS_ARCH_SPARC
		// Mode: CS_MODE_BIG_ENDIAN
		// Code: SPARC_CODE
		// Comment: "Sparc"
		arch:    CS_ARCH_SPARC,
		mode:    CS_MODE_BIG_ENDIAN,
		code:    detailCodes["SPARC_CODE"],
		comment: "Sparc",
		options: []option{
			{CS_OPT_DETAIL, CS_OPT_ON},
		},
	},
	{
		// -------------------------------------------
		// Platform Entry (line 220):
		// Architecture: CS_ARCH_SPARC
		// Mode: (CS_MODE_BIG_ENDIAN + CS_MODE_V9)
		// Code: SPARCV9_CODE
		// Comment: "SparcV9"
		arch:    CS_ARCH_SPARC,
		mode:    (CS_MODE_BIG_ENDIAN + CS_MODE_V9),
		code:    detailCodes["SPARCV9_CODE"],
		comment: "SparcV9",
		options: []option{
			{CS_OPT_DETAIL, CS_OPT_ON},
		},
	},
	{
		// -------------------------------------------
		// Platform Entry (line 229):
		// Architecture: CS_ARCH_SYSZ
		// Mode: 0
		// Code: SYSZ_CODE
		// Comment: "SystemZ"
		arch:    CS_ARCH_SYSZ,
		mode:    0,
		code:    detailCodes["SYSZ_CODE"],
		comment: "SystemZ",
		options: []option{
			{CS_OPT_DETAIL, CS_OPT_ON},
		},
	},
	{
		// -------------------------------------------
		// Platform Entry (line 238):
		// Architecture: CS_ARCH_XCORE
		// Mode: 0
		// Code: XCORE_CODE
		// Comment: "XCore"
		arch:    CS_ARCH_XCORE,
		mode:    0,
		code:    detailCodes["XCORE_CODE"],
		comment: "XCore",
		options: []option{
			{CS_OPT_DETAIL, CS_OPT_ON},
		},
	},
	{
		// -------------------------------------------
		// Platform Entry (line 247):
		// Architecture: CS_ARCH_M68K
		// Mode: (CS_MODE_BIG_ENDIAN | CS_MODE_M68K_040)
		// Code: M68K_CODE
		// Comment: "M68K"
		arch:    CS_ARCH_M68K,
		mode:    (CS_MODE_BIG_ENDIAN | CS_MODE_M68K_040),
		code:    detailCodes["M68K_CODE"],
		comment: "M68K",
		options: []option{
			{CS_OPT_DETAIL, CS_OPT_ON},
		},
	},
	{
		// -------------------------------------------
		// Platform Entry (line 256):
		// Architecture: CS_ARCH_M680X
		// Mode: (CS_MODE_M680X_6809)
		// Code: M680X_CODE
		// Comment: "M680X_M6809"
		arch:    CS_ARCH_M680X,
		mode:    (CS_MODE_M680X_6809),
		code:    detailCodes["M680X_CODE"],
		comment: "M680X_M6809",
		options: []option{
			{CS_OPT_DETAIL, CS_OPT_ON},
		},
	},
	{
		// -------------------------------------------
		// Platform Entry (line 265):
		// Architecture: CS_ARCH_MOS65XX
		// Mode: 0
		// Code: MOS65XX_CODE
		// Comment: "MOS65XX"
		arch:    CS_ARCH_MOS65XX,
		mode:    0,
		code:    detailCodes["MOS65XX_CODE"],
		comment: "MOS65XX",
		options: []option{
			{CS_OPT_DETAIL, CS_OPT_ON},
		},
	},
	{
		// -------------------------------------------
		// Platform Entry (line 274):
		// Architecture: CS_ARCH_BPF
		// Mode: CS_MODE_LITTLE_ENDIAN | CS_MODE_BPF_EXTENDED
		// Code: EBPF_CODE
		// Comment: "eBPF"
		arch:    CS_ARCH_BPF,
		mode:    CS_MODE_LITTLE_ENDIAN | CS_MODE_BPF_EXTENDED,
		code:    detailCodes["EBPF_CODE"],
		comment: "eBPF",
		options: []option{
			{CS_OPT_DETAIL, CS_OPT_ON},
		},
	},
}
