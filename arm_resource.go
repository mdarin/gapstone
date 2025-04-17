/*
Gapstone is a Go binding for the Capstone disassembly library.
For examples,try reading the *_test.go files.

	Library Author: Nguyen Anh Quynh
	Binding Author: Ben Nagy
	License: BSD style - see LICENSE file for details
    (c) 2013 COSEINC. All Rights Reserved.

    THIS FILE WAS AUTO-GENERATED -- DO NOT EDIT!
    Command: ./genresorces.sh capstone/tests/test_arm.c
    Created at: 2025-04-17 13:06:58

*/

package gapstone

//=== Found #define values in capstone/tests/test_arm.c ===

var armCodes = map[string]string{
	// ARM_CODE="\x86\x48\x60\xf4\x4d\x0f\xe2\xf4\xED\xFF\xFF\xEB\x04\xe0\x2d\xe5\x00\x00\x00\x00\xe0\x83\x22\xe5\xf1\x02\x03\x0e\x00\x00\xa0\xe3\x02\x30\xc1\xe7\x00\x00\x53\xe3\x00\x02\x01\xf1\x05\x40\xd0\xe8\xf4\x80\x00\x00"
	"ARM_CODE": "\x86\x48\x60\xf4\x4d\x0f\xe2\xf4\xED\xFF\xFF\xEB\x04\xe0\x2d\xe5\x00\x00\x00\x00\xe0\x83\x22\xe5\xf1\x02\x03\x0e\x00\x00\xa0\xe3\x02\x30\xc1\xe7\x00\x00\x53\xe3\x00\x02\x01\xf1\x05\x40\xd0\xe8\xf4\x80\x00\x00",
	// ARM_CODE2="\xd1\xe8\x00\xf0\xf0\x24\x04\x07\x1f\x3c\xf2\xc0\x00\x00\x4f\xf0\x00\x01\x46\x6c"
	"ARM_CODE2": "\xd1\xe8\x00\xf0\xf0\x24\x04\x07\x1f\x3c\xf2\xc0\x00\x00\x4f\xf0\x00\x01\x46\x6c",
	// THUMB_CODE="\x60\xf9\x1f\x04\xe0\xf9\x4f\x07\x70\x47\x00\xf0\x10\xe8\xeb\x46\x83\xb0\xc9\x68\x1f\xb1\x30\xbf\xaf\xf3\x20\x84\x52\xf8\x23\xf0"
	"THUMB_CODE": "\x60\xf9\x1f\x04\xe0\xf9\x4f\x07\x70\x47\x00\xf0\x10\xe8\xeb\x46\x83\xb0\xc9\x68\x1f\xb1\x30\xbf\xaf\xf3\x20\x84\x52\xf8\x23\xf0",
	// THUMB_CODE2="\x4f\xf0\x00\x01\xbd\xe8\x00\x88\xd1\xe8\x00\xf0\x18\xbf\xad\xbf\xf3\xff\x0b\x0c\x86\xf3\x00\x89\x80\xf3\x00\x8c\x4f\xfa\x99\xf6\xd0\xff\xa2\x01"
	"THUMB_CODE2": "\x4f\xf0\x00\x01\xbd\xe8\x00\x88\xd1\xe8\x00\xf0\x18\xbf\xad\xbf\xf3\xff\x0b\x0c\x86\xf3\x00\x89\x80\xf3\x00\x8c\x4f\xfa\x99\xf6\xd0\xff\xa2\x01",
	// THUMB_MCLASS="\xef\xf3\x02\x80"
	"THUMB_MCLASS": "\xef\xf3\x02\x80",
	// ARMV8="\xe0\x3b\xb2\xee\x42\x00\x01\xe1\x51\xf0\x7f\xf5"
	"ARMV8": "\xe0\x3b\xb2\xee\x42\x00\x01\xe1\x51\xf0\x7f\xf5",
}

//=== Found platform entries in capstone/tests/test_arm.c ===

var armPlatforms = []platform{
	{
		// -------------------------------------------
		// Platform Entry (line 195):
		// Architecture: CS_ARCH_ARM
		// Mode: CS_MODE_ARM
		// Code: ARM_CODE
		// Comment: "ARM"
		arch:    CS_ARCH_ARM,
		mode:    CS_MODE_ARM,
		code:    armCodes["ARM_CODE"],
		comment: "ARM",
		options: []option{
			{CS_OPT_DETAIL, CS_OPT_ON},
		},
	},
	{
		// -------------------------------------------
		// Platform Entry (line 202):
		// Architecture: CS_ARCH_ARM
		// Mode: CS_MODE_THUMB
		// Code: THUMB_CODE
		// Comment: "Thumb"
		arch:    CS_ARCH_ARM,
		mode:    CS_MODE_THUMB,
		code:    armCodes["THUMB_CODE"],
		comment: "Thumb",
		options: []option{
			{CS_OPT_DETAIL, CS_OPT_ON},
		},
	},
	{
		// -------------------------------------------
		// Platform Entry (line 209):
		// Architecture: CS_ARCH_ARM
		// Mode: CS_MODE_THUMB
		// Code: ARM_CODE2
		// Comment: "Thumb-mixed"
		arch:    CS_ARCH_ARM,
		mode:    CS_MODE_THUMB,
		code:    armCodes["ARM_CODE2"],
		comment: "Thumb-mixed",
		options: []option{
			{CS_OPT_DETAIL, CS_OPT_ON},
		},
	},
	{
		// -------------------------------------------
		// Platform Entry (line 216):
		// Architecture: CS_ARCH_ARM
		// Mode: CS_MODE_THUMB
		// Code: THUMB_CODE2
		// Comment: "Thumb-2 & register named with numbers"
		arch:    CS_ARCH_ARM,
		mode:    CS_MODE_THUMB,
		code:    armCodes["THUMB_CODE2"],
		comment: "Thumb-2 & register named with numbers",
		options: []option{
			{CS_OPT_DETAIL, CS_OPT_ON},
		},
	},
	{
		// -------------------------------------------
		// Platform Entry (line 224):
		// Architecture: CS_ARCH_ARM
		// Mode: (CS_MODE_THUMB + CS_MODE_MCLASS)
		// Code: THUMB_MCLASS
		// Comment: "Thumb-MClass"
		arch:    CS_ARCH_ARM,
		mode:    (CS_MODE_THUMB + CS_MODE_MCLASS),
		code:    armCodes["THUMB_MCLASS"],
		comment: "Thumb-MClass",
		options: []option{
			{CS_OPT_DETAIL, CS_OPT_ON},
		},
	},
	{
		// -------------------------------------------
		// Platform Entry (line 231):
		// Architecture: CS_ARCH_ARM
		// Mode: (CS_MODE_ARM + CS_MODE_V8)
		// Code: ARMV8
		// Comment: "Arm-V8"
		arch:    CS_ARCH_ARM,
		mode:    (CS_MODE_ARM + CS_MODE_V8),
		code:    armCodes["ARMV8"],
		comment: "Arm-V8",
		options: []option{
			{CS_OPT_DETAIL, CS_OPT_ON},
		},
	},
}
