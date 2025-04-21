/*
Gapstone is a Go binding for the Capstone disassembly library.
For examples,try reading the *_test.go files.

    Library Author: Nguyen Anh Quynh
    Binding Author: Ben Nagy
    License: BSD style - see LICENSE file for details
    (c) 2013 COSEINC. All Rights Reserved.

    THIS FILE WAS AUTO-GENERATED -- DO NOT EDIT!
    Command: ./gensanity.py
    Created at: 2025-04-19 08:40:30

*/

package gapstone

// !WARN
// Maintain the expected version and sanity checks manually, so we can verify
// against the installed C lib. Not foolproof, but should save 90% of accidents
const expectedMaj = 5
const expectedMin = 0

type sanityCheck struct {
	insMax int
	regMax int
	grpMax int
}

type sanityChecks map[int]sanityCheck

func (s *sanityChecks) Maj() int { return expectedMaj }
func (s *sanityChecks) Min() int { return expectedMin }

// !WARN
// Remember the all the constants CONST are direct refs to C.CONST, so in
// combination with these we should be _fairly_ sure we're getting the
// disassembly capstone expects to provide.
// * <ARCH>_GRP_ENDING || <ARCH>_INS_ENDING || <ARCH>_REG_ENDING
var checks = sanityChecks{
	CS_ARCH_ARM: sanityCheck{
		regMax: ARM_REG_ENDING,
		insMax: ARM_INS_ENDING,
		grpMax: ARM_GRP_ENDING,
	},
	CS_ARCH_ARM64: sanityCheck{
		regMax: ARM64_REG_ENDING,
		insMax: ARM64_INS_ENDING,
		grpMax: ARM64_GRP_ENDING,
	},
	CS_ARCH_BPF: sanityCheck{
		regMax: BPF_REG_ENDING,
		insMax: BPF_INS_ENDING,
		grpMax: BPF_GRP_ENDING,
	},
	CS_ARCH_M680X: sanityCheck{
		regMax: M680X_REG_ENDING,
		insMax: M680X_INS_ENDING,
		grpMax: M680X_GRP_ENDING,
	},
	CS_ARCH_M68K: sanityCheck{
		regMax: M68K_REG_ENDING,
		insMax: M68K_INS_ENDING,
		grpMax: M68K_GRP_ENDING,
	},
	CS_ARCH_MIPS: sanityCheck{
		regMax: MIPS_REG_ENDING,
		insMax: MIPS_INS_ENDING,
		grpMax: MIPS_GRP_ENDING,
	},
	CS_ARCH_MOS65XX: sanityCheck{
		regMax: MOS65XX_REG_ENDING,
		insMax: MOS65XX_INS_ENDING,
		grpMax: MOS65XX_GRP_ENDING,
	},
	CS_ARCH_PPC: sanityCheck{
		regMax: PPC_REG_ENDING,
		insMax: PPC_INS_ENDING,
		grpMax: PPC_GRP_ENDING,
	},
	CS_ARCH_RISCV: sanityCheck{
		regMax: RISCV_REG_ENDING,
		insMax: RISCV_INS_ENDING,
		grpMax: RISCV_GRP_ENDING,
	},
	CS_ARCH_SPARC: sanityCheck{
		regMax: SPARC_REG_ENDING,
		insMax: SPARC_INS_ENDING,
		grpMax: SPARC_GRP_ENDING,
	},
	CS_ARCH_SYSZ: sanityCheck{
		regMax: SYSZ_REG_ENDING,
		insMax: SYSZ_INS_ENDING,
		grpMax: SYSZ_GRP_ENDING,
	},
	CS_ARCH_TMS320C64X: sanityCheck{
		regMax: TMS320C64X_REG_ENDING,
		insMax: TMS320C64X_INS_ENDING,
		grpMax: TMS320C64X_GRP_ENDING,
	},
	CS_ARCH_X86: sanityCheck{
		regMax: X86_REG_ENDING,
		insMax: X86_INS_ENDING,
		grpMax: X86_GRP_ENDING,
	},
	CS_ARCH_XCORE: sanityCheck{
		regMax: XCORE_REG_ENDING,
		insMax: XCORE_INS_ENDING,
		grpMax: XCORE_GRP_ENDING,
	},
}
