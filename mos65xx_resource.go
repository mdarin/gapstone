/*
Gapstone is a Go binding for the Capstone disassembly library.
For examples,try reading the *_test.go files.

	Library Author: Nguyen Anh Quynh
	Binding Author: Ben Nagy
	License: BSD style - see LICENSE file for details
    (c) 2013 COSEINC. All Rights Reserved.

    THIS FILE WAS AUTO-GENERATED -- DO NOT EDIT!
    Command: ./genresorces.sh capstone/tests/test_mos65xx.c
    Created at: 2025-04-17 13:06:57

*/

package gapstone

//=== Found #define values in capstone/tests/test_mos65xx.c ===

var mos65xxCodes = map[string]string{
	// M6502_CODE="\xa1\x12\xa5\x12\xa9\x12\xad\x34\x12\xb1\x12\xb5\x12\xb9\x34\x12\xbd\x34\x12\x0d\x34\x12\x00\x81\x87\x6c\x01\x00\x85\xFF\x10\x00\x19\x42\x42\x00\x49\x42"
	"M6502_CODE": "\xa1\x12\xa5\x12\xa9\x12\xad\x34\x12\xb1\x12\xb5\x12\xb9\x34\x12\xbd\x34\x12\x0d\x34\x12\x00\x81\x87\x6c\x01\x00\x85\xFF\x10\x00\x19\x42\x42\x00\x49\x42",
	// M65C02_CODE="\x1a\x3a\x02\x12\x03\x5c\x34\x12"
	"M65C02_CODE": "\x1a\x3a\x02\x12\x03\x5c\x34\x12",
	// MW65C02_CODE="\x07\x12\x27\x12\x47\x12\x67\x12\x87\x12\xa7\x12\xc7\x12\xe7\x12\x10\xfe\x0f\x12\xfd\x4f\x12\xfd\x8f\x12\xfd\xcf\x12\xfd"
	"MW65C02_CODE": "\x07\x12\x27\x12\x47\x12\x67\x12\x87\x12\xa7\x12\xc7\x12\xe7\x12\x10\xfe\x0f\x12\xfd\x4f\x12\xfd\x8f\x12\xfd\xcf\x12\xfd",
	// M65816_CODE="\xa9\x34\x12\xad\x34\x12\xbd\x34\x12\xb9\x34\x12\xaf\x56\x34\x12\xbf\x56\x34\x12\xa5\x12\xb5\x12\xb2\x12\xa1\x12\xb1\x12\xa7\x12\xb7\x12\xa3\x12\xb3\x12\xc2\x00\xe2\x00\x54\x34\x12\x44\x34\x12\x02\x12"
	"M65816_CODE": "\xa9\x34\x12\xad\x34\x12\xbd\x34\x12\xb9\x34\x12\xaf\x56\x34\x12\xbf\x56\x34\x12\xa5\x12\xb5\x12\xb2\x12\xa1\x12\xb1\x12\xa7\x12\xb7\x12\xa3\x12\xb3\x12\xc2\x00\xe2\x00\x54\x34\x12\x44\x34\x12\x02\x12",
}

//=== Found platform entries in capstone/tests/test_mos65xx.c ===

var mos65xxPlatforms = []platform{
	{
		// -------------------------------------------
		// Platform Entry (line 148):
		// Architecture: CS_ARCH_MOS65XX
		// Mode: (CS_MODE_MOS65XX_6502)
		// Code: M6502_CODE
		// Comment: "MOS65XX_6502"
		arch:    CS_ARCH_MOS65XX,
		mode:    (CS_MODE_MOS65XX_6502),
		code:    mos65xxCodes["M6502_CODE"],
		comment: "MOS65XX_6502",
		options: []option{
			{CS_OPT_DETAIL, CS_OPT_ON},
		},
	},
	{
		// -------------------------------------------
		// Platform Entry (line 155):
		// Architecture: CS_ARCH_MOS65XX
		// Mode: (CS_MODE_MOS65XX_65C02)
		// Code: M65C02_CODE
		// Comment: "MOS65XX_65C02"
		arch:    CS_ARCH_MOS65XX,
		mode:    (CS_MODE_MOS65XX_65C02),
		code:    mos65xxCodes["M65C02_CODE"],
		comment: "MOS65XX_65C02",
		options: []option{
			{CS_OPT_DETAIL, CS_OPT_ON},
		},
	},
	{
		// -------------------------------------------
		// Platform Entry (line 162):
		// Architecture: CS_ARCH_MOS65XX
		// Mode: (CS_MODE_MOS65XX_W65C02)
		// Code: MW65C02_CODE
		// Comment: "MOS65XX_W65C02"
		arch:    CS_ARCH_MOS65XX,
		mode:    (CS_MODE_MOS65XX_W65C02),
		code:    mos65xxCodes["MW65C02_CODE"],
		comment: "MOS65XX_W65C02",
		options: []option{
			{CS_OPT_DETAIL, CS_OPT_ON},
		},
	},
	{
		// -------------------------------------------
		// Platform Entry (line 169):
		// Architecture: CS_ARCH_MOS65XX
		// Mode: (CS_MODE_MOS65XX_65816_LONG_MX)
		// Code: M65816_CODE
		// Comment: "MOS65XX_65816 (long m/x)"
		arch:    CS_ARCH_MOS65XX,
		mode:    (CS_MODE_MOS65XX_65816_LONG_MX),
		code:    mos65xxCodes["M65816_CODE"],
		comment: "MOS65XX_65816 (long m/x)",
		options: []option{
			{CS_OPT_DETAIL, CS_OPT_ON},
		},
	},
}
