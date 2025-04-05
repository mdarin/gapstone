/*
Gapstone is a Go binding for the Capstone disassembly library.
For examples, try reading the *_test.go files.

	Library Author: Nguyen Anh Quynh
	Binding Author: Ben Nagy
	License: BSD style - see LICENSE file for details
    (c) 2013 COSEINC. All Rights Reserved.

    THIS FILE WAS AUTO-GENERATED -- DO NOT EDIT!
	Command: ./genconst capstone/bindings/python/capstone/
	Created at: 2025-04-05T13:11:18+00:00

*/

package gapstone

// #cgo LDFLAGS: -lcapstone
// #cgo freebsd CFLAGS: -I/usr/local/include
// #cgo freebsd LDFLAGS: -L/usr/local/lib
// #include <stdlib.h>
// #include <capstone/capstone.h>
import "C"

// For Capstone Engine. AUTO-GENERATED FILE, DO NOT EDIT [arm_const.py]
const (
	ARM_SFT_INVALID = C.ARM_SFT_INVALID
	ARM_SFT_ASR = C.ARM_SFT_ASR
	ARM_SFT_LSL = C.ARM_SFT_LSL
	ARM_SFT_LSR = C.ARM_SFT_LSR
	ARM_SFT_ROR = C.ARM_SFT_ROR
	ARM_SFT_RRX = C.ARM_SFT_RRX
	ARM_SFT_ASR_REG = C.ARM_SFT_ASR_REG
	ARM_SFT_LSL_REG = C.ARM_SFT_LSL_REG
	ARM_SFT_LSR_REG = C.ARM_SFT_LSR_REG
	ARM_SFT_ROR_REG = C.ARM_SFT_ROR_REG
	ARM_SFT_RRX_REG = C.ARM_SFT_RRX_REG
)

const (
	ARM_CC_INVALID = C.ARM_CC_INVALID
	ARM_CC_EQ = C.ARM_CC_EQ
	ARM_CC_NE = C.ARM_CC_NE
	ARM_CC_HS = C.ARM_CC_HS
	ARM_CC_LO = C.ARM_CC_LO
	ARM_CC_MI = C.ARM_CC_MI
	ARM_CC_PL = C.ARM_CC_PL
	ARM_CC_VS = C.ARM_CC_VS
	ARM_CC_VC = C.ARM_CC_VC
	ARM_CC_HI = C.ARM_CC_HI
	ARM_CC_LS = C.ARM_CC_LS
	ARM_CC_GE = C.ARM_CC_GE
	ARM_CC_LT = C.ARM_CC_LT
	ARM_CC_GT = C.ARM_CC_GT
	ARM_CC_LE = C.ARM_CC_LE
	ARM_CC_AL = C.ARM_CC_AL
)

const (
	ARM_SYSREG_INVALID = C.ARM_SYSREG_INVALID
	ARM_SYSREG_SPSR_C = C.ARM_SYSREG_SPSR_C
	ARM_SYSREG_SPSR_X = C.ARM_SYSREG_SPSR_X
	ARM_SYSREG_SPSR_S = C.ARM_SYSREG_SPSR_S
	ARM_SYSREG_SPSR_F = C.ARM_SYSREG_SPSR_F
	ARM_SYSREG_CPSR_C = C.ARM_SYSREG_CPSR_C
	ARM_SYSREG_CPSR_X = C.ARM_SYSREG_CPSR_X
	ARM_SYSREG_CPSR_S = C.ARM_SYSREG_CPSR_S
	ARM_SYSREG_CPSR_F = C.ARM_SYSREG_CPSR_F
	ARM_SYSREG_APSR = C.ARM_SYSREG_APSR
	ARM_SYSREG_APSR_G = C.ARM_SYSREG_APSR_G
	ARM_SYSREG_APSR_NZCVQ = C.ARM_SYSREG_APSR_NZCVQ
	ARM_SYSREG_APSR_NZCVQG = C.ARM_SYSREG_APSR_NZCVQG
	ARM_SYSREG_IAPSR = C.ARM_SYSREG_IAPSR
	ARM_SYSREG_IAPSR_G = C.ARM_SYSREG_IAPSR_G
	ARM_SYSREG_IAPSR_NZCVQG = C.ARM_SYSREG_IAPSR_NZCVQG
	ARM_SYSREG_IAPSR_NZCVQ = C.ARM_SYSREG_IAPSR_NZCVQ
	ARM_SYSREG_EAPSR = C.ARM_SYSREG_EAPSR
	ARM_SYSREG_EAPSR_G = C.ARM_SYSREG_EAPSR_G
	ARM_SYSREG_EAPSR_NZCVQG = C.ARM_SYSREG_EAPSR_NZCVQG
	ARM_SYSREG_EAPSR_NZCVQ = C.ARM_SYSREG_EAPSR_NZCVQ
	ARM_SYSREG_XPSR = C.ARM_SYSREG_XPSR
	ARM_SYSREG_XPSR_G = C.ARM_SYSREG_XPSR_G
	ARM_SYSREG_XPSR_NZCVQG = C.ARM_SYSREG_XPSR_NZCVQG
	ARM_SYSREG_XPSR_NZCVQ = C.ARM_SYSREG_XPSR_NZCVQ
	ARM_SYSREG_IPSR = C.ARM_SYSREG_IPSR
	ARM_SYSREG_EPSR = C.ARM_SYSREG_EPSR
	ARM_SYSREG_IEPSR = C.ARM_SYSREG_IEPSR
	ARM_SYSREG_MSP = C.ARM_SYSREG_MSP
	ARM_SYSREG_PSP = C.ARM_SYSREG_PSP
	ARM_SYSREG_PRIMASK = C.ARM_SYSREG_PRIMASK
	ARM_SYSREG_BASEPRI = C.ARM_SYSREG_BASEPRI
	ARM_SYSREG_BASEPRI_MAX = C.ARM_SYSREG_BASEPRI_MAX
	ARM_SYSREG_FAULTMASK = C.ARM_SYSREG_FAULTMASK
	ARM_SYSREG_CONTROL = C.ARM_SYSREG_CONTROL
	ARM_SYSREG_MSPLIM = C.ARM_SYSREG_MSPLIM
	ARM_SYSREG_PSPLIM = C.ARM_SYSREG_PSPLIM
	ARM_SYSREG_MSP_NS = C.ARM_SYSREG_MSP_NS
	ARM_SYSREG_PSP_NS = C.ARM_SYSREG_PSP_NS
	ARM_SYSREG_MSPLIM_NS = C.ARM_SYSREG_MSPLIM_NS
	ARM_SYSREG_PSPLIM_NS = C.ARM_SYSREG_PSPLIM_NS
	ARM_SYSREG_PRIMASK_NS = C.ARM_SYSREG_PRIMASK_NS
	ARM_SYSREG_BASEPRI_NS = C.ARM_SYSREG_BASEPRI_NS
	ARM_SYSREG_FAULTMASK_NS = C.ARM_SYSREG_FAULTMASK_NS
	ARM_SYSREG_CONTROL_NS = C.ARM_SYSREG_CONTROL_NS
	ARM_SYSREG_SP_NS = C.ARM_SYSREG_SP_NS
	ARM_SYSREG_R8_USR = C.ARM_SYSREG_R8_USR
	ARM_SYSREG_R9_USR = C.ARM_SYSREG_R9_USR
	ARM_SYSREG_R10_USR = C.ARM_SYSREG_R10_USR
	ARM_SYSREG_R11_USR = C.ARM_SYSREG_R11_USR
	ARM_SYSREG_R12_USR = C.ARM_SYSREG_R12_USR
	ARM_SYSREG_SP_USR = C.ARM_SYSREG_SP_USR
	ARM_SYSREG_LR_USR = C.ARM_SYSREG_LR_USR
	ARM_SYSREG_R8_FIQ = C.ARM_SYSREG_R8_FIQ
	ARM_SYSREG_R9_FIQ = C.ARM_SYSREG_R9_FIQ
	ARM_SYSREG_R10_FIQ = C.ARM_SYSREG_R10_FIQ
	ARM_SYSREG_R11_FIQ = C.ARM_SYSREG_R11_FIQ
	ARM_SYSREG_R12_FIQ = C.ARM_SYSREG_R12_FIQ
	ARM_SYSREG_SP_FIQ = C.ARM_SYSREG_SP_FIQ
	ARM_SYSREG_LR_FIQ = C.ARM_SYSREG_LR_FIQ
	ARM_SYSREG_LR_IRQ = C.ARM_SYSREG_LR_IRQ
	ARM_SYSREG_SP_IRQ = C.ARM_SYSREG_SP_IRQ
	ARM_SYSREG_LR_SVC = C.ARM_SYSREG_LR_SVC
	ARM_SYSREG_SP_SVC = C.ARM_SYSREG_SP_SVC
	ARM_SYSREG_LR_ABT = C.ARM_SYSREG_LR_ABT
	ARM_SYSREG_SP_ABT = C.ARM_SYSREG_SP_ABT
	ARM_SYSREG_LR_UND = C.ARM_SYSREG_LR_UND
	ARM_SYSREG_SP_UND = C.ARM_SYSREG_SP_UND
	ARM_SYSREG_LR_MON = C.ARM_SYSREG_LR_MON
	ARM_SYSREG_SP_MON = C.ARM_SYSREG_SP_MON
	ARM_SYSREG_ELR_HYP = C.ARM_SYSREG_ELR_HYP
	ARM_SYSREG_SP_HYP = C.ARM_SYSREG_SP_HYP
	ARM_SYSREG_SPSR_FIQ = C.ARM_SYSREG_SPSR_FIQ
	ARM_SYSREG_SPSR_IRQ = C.ARM_SYSREG_SPSR_IRQ
	ARM_SYSREG_SPSR_SVC = C.ARM_SYSREG_SPSR_SVC
	ARM_SYSREG_SPSR_ABT = C.ARM_SYSREG_SPSR_ABT
	ARM_SYSREG_SPSR_UND = C.ARM_SYSREG_SPSR_UND
	ARM_SYSREG_SPSR_MON = C.ARM_SYSREG_SPSR_MON
	ARM_SYSREG_SPSR_HYP = C.ARM_SYSREG_SPSR_HYP
)

const (
	ARM_MB_INVALID = C.ARM_MB_INVALID
	ARM_MB_RESERVED_0 = C.ARM_MB_RESERVED_0
	ARM_MB_OSHLD = C.ARM_MB_OSHLD
	ARM_MB_OSHST = C.ARM_MB_OSHST
	ARM_MB_OSH = C.ARM_MB_OSH
	ARM_MB_RESERVED_4 = C.ARM_MB_RESERVED_4
	ARM_MB_NSHLD = C.ARM_MB_NSHLD
	ARM_MB_NSHST = C.ARM_MB_NSHST
	ARM_MB_NSH = C.ARM_MB_NSH
	ARM_MB_RESERVED_8 = C.ARM_MB_RESERVED_8
	ARM_MB_ISHLD = C.ARM_MB_ISHLD
	ARM_MB_ISHST = C.ARM_MB_ISHST
	ARM_MB_ISH = C.ARM_MB_ISH
	ARM_MB_RESERVED_12 = C.ARM_MB_RESERVED_12
	ARM_MB_LD = C.ARM_MB_LD
	ARM_MB_ST = C.ARM_MB_ST
	ARM_MB_SY = C.ARM_MB_SY
)

const (
	ARM_OP_INVALID = C.ARM_OP_INVALID
	ARM_OP_REG = C.ARM_OP_REG
	ARM_OP_IMM = C.ARM_OP_IMM
	ARM_OP_MEM = C.ARM_OP_MEM
	ARM_OP_FP = C.ARM_OP_FP
	ARM_OP_CIMM = C.ARM_OP_CIMM
	ARM_OP_PIMM = C.ARM_OP_PIMM
	ARM_OP_SETEND = C.ARM_OP_SETEND
	ARM_OP_SYSREG = C.ARM_OP_SYSREG
)

const (
	ARM_SETEND_INVALID = C.ARM_SETEND_INVALID
	ARM_SETEND_BE = C.ARM_SETEND_BE
	ARM_SETEND_LE = C.ARM_SETEND_LE
)

const (
	ARM_CPSMODE_INVALID = C.ARM_CPSMODE_INVALID
	ARM_CPSMODE_IE = C.ARM_CPSMODE_IE
	ARM_CPSMODE_ID = C.ARM_CPSMODE_ID
)

const (
	ARM_CPSFLAG_INVALID = C.ARM_CPSFLAG_INVALID
	ARM_CPSFLAG_F = C.ARM_CPSFLAG_F
	ARM_CPSFLAG_I = C.ARM_CPSFLAG_I
	ARM_CPSFLAG_A = C.ARM_CPSFLAG_A
	ARM_CPSFLAG_NONE = C.ARM_CPSFLAG_NONE
)

const (
	ARM_VECTORDATA_INVALID = C.ARM_VECTORDATA_INVALID
	ARM_VECTORDATA_I8 = C.ARM_VECTORDATA_I8
	ARM_VECTORDATA_I16 = C.ARM_VECTORDATA_I16
	ARM_VECTORDATA_I32 = C.ARM_VECTORDATA_I32
	ARM_VECTORDATA_I64 = C.ARM_VECTORDATA_I64
	ARM_VECTORDATA_S8 = C.ARM_VECTORDATA_S8
	ARM_VECTORDATA_S16 = C.ARM_VECTORDATA_S16
	ARM_VECTORDATA_S32 = C.ARM_VECTORDATA_S32
	ARM_VECTORDATA_S64 = C.ARM_VECTORDATA_S64
	ARM_VECTORDATA_U8 = C.ARM_VECTORDATA_U8
	ARM_VECTORDATA_U16 = C.ARM_VECTORDATA_U16
	ARM_VECTORDATA_U32 = C.ARM_VECTORDATA_U32
	ARM_VECTORDATA_U64 = C.ARM_VECTORDATA_U64
	ARM_VECTORDATA_P8 = C.ARM_VECTORDATA_P8
	ARM_VECTORDATA_F16 = C.ARM_VECTORDATA_F16
	ARM_VECTORDATA_F32 = C.ARM_VECTORDATA_F32
	ARM_VECTORDATA_F64 = C.ARM_VECTORDATA_F64
	ARM_VECTORDATA_F16F64 = C.ARM_VECTORDATA_F16F64
	ARM_VECTORDATA_F64F16 = C.ARM_VECTORDATA_F64F16
	ARM_VECTORDATA_F32F16 = C.ARM_VECTORDATA_F32F16
	ARM_VECTORDATA_F16F32 = C.ARM_VECTORDATA_F16F32
	ARM_VECTORDATA_F64F32 = C.ARM_VECTORDATA_F64F32
	ARM_VECTORDATA_F32F64 = C.ARM_VECTORDATA_F32F64
	ARM_VECTORDATA_S32F32 = C.ARM_VECTORDATA_S32F32
	ARM_VECTORDATA_U32F32 = C.ARM_VECTORDATA_U32F32
	ARM_VECTORDATA_F32S32 = C.ARM_VECTORDATA_F32S32
	ARM_VECTORDATA_F32U32 = C.ARM_VECTORDATA_F32U32
	ARM_VECTORDATA_F64S16 = C.ARM_VECTORDATA_F64S16
	ARM_VECTORDATA_F32S16 = C.ARM_VECTORDATA_F32S16
	ARM_VECTORDATA_F64S32 = C.ARM_VECTORDATA_F64S32
	ARM_VECTORDATA_S16F64 = C.ARM_VECTORDATA_S16F64
	ARM_VECTORDATA_S16F32 = C.ARM_VECTORDATA_S16F32
	ARM_VECTORDATA_S32F64 = C.ARM_VECTORDATA_S32F64
	ARM_VECTORDATA_U16F64 = C.ARM_VECTORDATA_U16F64
	ARM_VECTORDATA_U16F32 = C.ARM_VECTORDATA_U16F32
	ARM_VECTORDATA_U32F64 = C.ARM_VECTORDATA_U32F64
	ARM_VECTORDATA_F64U16 = C.ARM_VECTORDATA_F64U16
	ARM_VECTORDATA_F32U16 = C.ARM_VECTORDATA_F32U16
	ARM_VECTORDATA_F64U32 = C.ARM_VECTORDATA_F64U32
	ARM_VECTORDATA_F16U16 = C.ARM_VECTORDATA_F16U16
	ARM_VECTORDATA_U16F16 = C.ARM_VECTORDATA_U16F16
	ARM_VECTORDATA_F16U32 = C.ARM_VECTORDATA_F16U32
	ARM_VECTORDATA_U32F16 = C.ARM_VECTORDATA_U32F16
)

const (
	ARM_REG_INVALID = C.ARM_REG_INVALID
	ARM_REG_APSR = C.ARM_REG_APSR
	ARM_REG_APSR_NZCV = C.ARM_REG_APSR_NZCV
	ARM_REG_CPSR = C.ARM_REG_CPSR
	ARM_REG_FPEXC = C.ARM_REG_FPEXC
	ARM_REG_FPINST = C.ARM_REG_FPINST
	ARM_REG_FPSCR = C.ARM_REG_FPSCR
	ARM_REG_FPSCR_NZCV = C.ARM_REG_FPSCR_NZCV
	ARM_REG_FPSID = C.ARM_REG_FPSID
	ARM_REG_ITSTATE = C.ARM_REG_ITSTATE
	ARM_REG_LR = C.ARM_REG_LR
	ARM_REG_PC = C.ARM_REG_PC
	ARM_REG_SP = C.ARM_REG_SP
	ARM_REG_SPSR = C.ARM_REG_SPSR
	ARM_REG_D0 = C.ARM_REG_D0
	ARM_REG_D1 = C.ARM_REG_D1
	ARM_REG_D2 = C.ARM_REG_D2
	ARM_REG_D3 = C.ARM_REG_D3
	ARM_REG_D4 = C.ARM_REG_D4
	ARM_REG_D5 = C.ARM_REG_D5
	ARM_REG_D6 = C.ARM_REG_D6
	ARM_REG_D7 = C.ARM_REG_D7
	ARM_REG_D8 = C.ARM_REG_D8
	ARM_REG_D9 = C.ARM_REG_D9
	ARM_REG_D10 = C.ARM_REG_D10
	ARM_REG_D11 = C.ARM_REG_D11
	ARM_REG_D12 = C.ARM_REG_D12
	ARM_REG_D13 = C.ARM_REG_D13
	ARM_REG_D14 = C.ARM_REG_D14
	ARM_REG_D15 = C.ARM_REG_D15
	ARM_REG_D16 = C.ARM_REG_D16
	ARM_REG_D17 = C.ARM_REG_D17
	ARM_REG_D18 = C.ARM_REG_D18
	ARM_REG_D19 = C.ARM_REG_D19
	ARM_REG_D20 = C.ARM_REG_D20
	ARM_REG_D21 = C.ARM_REG_D21
	ARM_REG_D22 = C.ARM_REG_D22
	ARM_REG_D23 = C.ARM_REG_D23
	ARM_REG_D24 = C.ARM_REG_D24
	ARM_REG_D25 = C.ARM_REG_D25
	ARM_REG_D26 = C.ARM_REG_D26
	ARM_REG_D27 = C.ARM_REG_D27
	ARM_REG_D28 = C.ARM_REG_D28
	ARM_REG_D29 = C.ARM_REG_D29
	ARM_REG_D30 = C.ARM_REG_D30
	ARM_REG_D31 = C.ARM_REG_D31
	ARM_REG_FPINST2 = C.ARM_REG_FPINST2
	ARM_REG_MVFR0 = C.ARM_REG_MVFR0
	ARM_REG_MVFR1 = C.ARM_REG_MVFR1
	ARM_REG_MVFR2 = C.ARM_REG_MVFR2
	ARM_REG_Q0 = C.ARM_REG_Q0
	ARM_REG_Q1 = C.ARM_REG_Q1
	ARM_REG_Q2 = C.ARM_REG_Q2
	ARM_REG_Q3 = C.ARM_REG_Q3
	ARM_REG_Q4 = C.ARM_REG_Q4
	ARM_REG_Q5 = C.ARM_REG_Q5
	ARM_REG_Q6 = C.ARM_REG_Q6
	ARM_REG_Q7 = C.ARM_REG_Q7
	ARM_REG_Q8 = C.ARM_REG_Q8
	ARM_REG_Q9 = C.ARM_REG_Q9
	ARM_REG_Q10 = C.ARM_REG_Q10
	ARM_REG_Q11 = C.ARM_REG_Q11
	ARM_REG_Q12 = C.ARM_REG_Q12
	ARM_REG_Q13 = C.ARM_REG_Q13
	ARM_REG_Q14 = C.ARM_REG_Q14
	ARM_REG_Q15 = C.ARM_REG_Q15
	ARM_REG_R0 = C.ARM_REG_R0
	ARM_REG_R1 = C.ARM_REG_R1
	ARM_REG_R2 = C.ARM_REG_R2
	ARM_REG_R3 = C.ARM_REG_R3
	ARM_REG_R4 = C.ARM_REG_R4
	ARM_REG_R5 = C.ARM_REG_R5
	ARM_REG_R6 = C.ARM_REG_R6
	ARM_REG_R7 = C.ARM_REG_R7
	ARM_REG_R8 = C.ARM_REG_R8
	ARM_REG_R9 = C.ARM_REG_R9
	ARM_REG_R10 = C.ARM_REG_R10
	ARM_REG_R11 = C.ARM_REG_R11
	ARM_REG_R12 = C.ARM_REG_R12
	ARM_REG_S0 = C.ARM_REG_S0
	ARM_REG_S1 = C.ARM_REG_S1
	ARM_REG_S2 = C.ARM_REG_S2
	ARM_REG_S3 = C.ARM_REG_S3
	ARM_REG_S4 = C.ARM_REG_S4
	ARM_REG_S5 = C.ARM_REG_S5
	ARM_REG_S6 = C.ARM_REG_S6
	ARM_REG_S7 = C.ARM_REG_S7
	ARM_REG_S8 = C.ARM_REG_S8
	ARM_REG_S9 = C.ARM_REG_S9
	ARM_REG_S10 = C.ARM_REG_S10
	ARM_REG_S11 = C.ARM_REG_S11
	ARM_REG_S12 = C.ARM_REG_S12
	ARM_REG_S13 = C.ARM_REG_S13
	ARM_REG_S14 = C.ARM_REG_S14
	ARM_REG_S15 = C.ARM_REG_S15
	ARM_REG_S16 = C.ARM_REG_S16
	ARM_REG_S17 = C.ARM_REG_S17
	ARM_REG_S18 = C.ARM_REG_S18
	ARM_REG_S19 = C.ARM_REG_S19
	ARM_REG_S20 = C.ARM_REG_S20
	ARM_REG_S21 = C.ARM_REG_S21
	ARM_REG_S22 = C.ARM_REG_S22
	ARM_REG_S23 = C.ARM_REG_S23
	ARM_REG_S24 = C.ARM_REG_S24
	ARM_REG_S25 = C.ARM_REG_S25
	ARM_REG_S26 = C.ARM_REG_S26
	ARM_REG_S27 = C.ARM_REG_S27
	ARM_REG_S28 = C.ARM_REG_S28
	ARM_REG_S29 = C.ARM_REG_S29
	ARM_REG_S30 = C.ARM_REG_S30
	ARM_REG_S31 = C.ARM_REG_S31
	ARM_REG_ENDING = C.ARM_REG_ENDING
	ARM_REG_R13 = C.ARM_REG_R13
	ARM_REG_R14 = C.ARM_REG_R14
	ARM_REG_R15 = C.ARM_REG_R15
	ARM_REG_SB = C.ARM_REG_SB
	ARM_REG_SL = C.ARM_REG_SL
	ARM_REG_FP = C.ARM_REG_FP
	ARM_REG_IP = C.ARM_REG_IP
)

const (
	ARM_INS_INVALID = C.ARM_INS_INVALID
	ARM_INS_ADC = C.ARM_INS_ADC
	ARM_INS_ADD = C.ARM_INS_ADD
	ARM_INS_ADDW = C.ARM_INS_ADDW
	ARM_INS_ADR = C.ARM_INS_ADR
	ARM_INS_AESD = C.ARM_INS_AESD
	ARM_INS_AESE = C.ARM_INS_AESE
	ARM_INS_AESIMC = C.ARM_INS_AESIMC
	ARM_INS_AESMC = C.ARM_INS_AESMC
	ARM_INS_AND = C.ARM_INS_AND
	ARM_INS_ASR = C.ARM_INS_ASR
	ARM_INS_B = C.ARM_INS_B
	ARM_INS_BFC = C.ARM_INS_BFC
	ARM_INS_BFI = C.ARM_INS_BFI
	ARM_INS_BIC = C.ARM_INS_BIC
	ARM_INS_BKPT = C.ARM_INS_BKPT
	ARM_INS_BL = C.ARM_INS_BL
	ARM_INS_BLX = C.ARM_INS_BLX
	ARM_INS_BLXNS = C.ARM_INS_BLXNS
	ARM_INS_BX = C.ARM_INS_BX
	ARM_INS_BXJ = C.ARM_INS_BXJ
	ARM_INS_BXNS = C.ARM_INS_BXNS
	ARM_INS_CBNZ = C.ARM_INS_CBNZ
	ARM_INS_CBZ = C.ARM_INS_CBZ
	ARM_INS_CDP = C.ARM_INS_CDP
	ARM_INS_CDP2 = C.ARM_INS_CDP2
	ARM_INS_CLREX = C.ARM_INS_CLREX
	ARM_INS_CLZ = C.ARM_INS_CLZ
	ARM_INS_CMN = C.ARM_INS_CMN
	ARM_INS_CMP = C.ARM_INS_CMP
	ARM_INS_CPS = C.ARM_INS_CPS
	ARM_INS_CRC32B = C.ARM_INS_CRC32B
	ARM_INS_CRC32CB = C.ARM_INS_CRC32CB
	ARM_INS_CRC32CH = C.ARM_INS_CRC32CH
	ARM_INS_CRC32CW = C.ARM_INS_CRC32CW
	ARM_INS_CRC32H = C.ARM_INS_CRC32H
	ARM_INS_CRC32W = C.ARM_INS_CRC32W
	ARM_INS_CSDB = C.ARM_INS_CSDB
	ARM_INS_DBG = C.ARM_INS_DBG
	ARM_INS_DCPS1 = C.ARM_INS_DCPS1
	ARM_INS_DCPS2 = C.ARM_INS_DCPS2
	ARM_INS_DCPS3 = C.ARM_INS_DCPS3
	ARM_INS_DFB = C.ARM_INS_DFB
	ARM_INS_DMB = C.ARM_INS_DMB
	ARM_INS_DSB = C.ARM_INS_DSB
	ARM_INS_EOR = C.ARM_INS_EOR
	ARM_INS_ERET = C.ARM_INS_ERET
	ARM_INS_ESB = C.ARM_INS_ESB
	ARM_INS_FADDD = C.ARM_INS_FADDD
	ARM_INS_FADDS = C.ARM_INS_FADDS
	ARM_INS_FCMPZD = C.ARM_INS_FCMPZD
	ARM_INS_FCMPZS = C.ARM_INS_FCMPZS
	ARM_INS_FCONSTD = C.ARM_INS_FCONSTD
	ARM_INS_FCONSTS = C.ARM_INS_FCONSTS
	ARM_INS_FLDMDBX = C.ARM_INS_FLDMDBX
	ARM_INS_FLDMIAX = C.ARM_INS_FLDMIAX
	ARM_INS_FMDHR = C.ARM_INS_FMDHR
	ARM_INS_FMDLR = C.ARM_INS_FMDLR
	ARM_INS_FMSTAT = C.ARM_INS_FMSTAT
	ARM_INS_FSTMDBX = C.ARM_INS_FSTMDBX
	ARM_INS_FSTMIAX = C.ARM_INS_FSTMIAX
	ARM_INS_FSUBD = C.ARM_INS_FSUBD
	ARM_INS_FSUBS = C.ARM_INS_FSUBS
	ARM_INS_HINT = C.ARM_INS_HINT
	ARM_INS_HLT = C.ARM_INS_HLT
	ARM_INS_HVC = C.ARM_INS_HVC
	ARM_INS_ISB = C.ARM_INS_ISB
	ARM_INS_IT = C.ARM_INS_IT
	ARM_INS_LDA = C.ARM_INS_LDA
	ARM_INS_LDAB = C.ARM_INS_LDAB
	ARM_INS_LDAEX = C.ARM_INS_LDAEX
	ARM_INS_LDAEXB = C.ARM_INS_LDAEXB
	ARM_INS_LDAEXD = C.ARM_INS_LDAEXD
	ARM_INS_LDAEXH = C.ARM_INS_LDAEXH
	ARM_INS_LDAH = C.ARM_INS_LDAH
	ARM_INS_LDC = C.ARM_INS_LDC
	ARM_INS_LDC2 = C.ARM_INS_LDC2
	ARM_INS_LDC2L = C.ARM_INS_LDC2L
	ARM_INS_LDCL = C.ARM_INS_LDCL
	ARM_INS_LDM = C.ARM_INS_LDM
	ARM_INS_LDMDA = C.ARM_INS_LDMDA
	ARM_INS_LDMDB = C.ARM_INS_LDMDB
	ARM_INS_LDMIB = C.ARM_INS_LDMIB
	ARM_INS_LDR = C.ARM_INS_LDR
	ARM_INS_LDRB = C.ARM_INS_LDRB
	ARM_INS_LDRBT = C.ARM_INS_LDRBT
	ARM_INS_LDRD = C.ARM_INS_LDRD
	ARM_INS_LDREX = C.ARM_INS_LDREX
	ARM_INS_LDREXB = C.ARM_INS_LDREXB
	ARM_INS_LDREXD = C.ARM_INS_LDREXD
	ARM_INS_LDREXH = C.ARM_INS_LDREXH
	ARM_INS_LDRH = C.ARM_INS_LDRH
	ARM_INS_LDRHT = C.ARM_INS_LDRHT
	ARM_INS_LDRSB = C.ARM_INS_LDRSB
	ARM_INS_LDRSBT = C.ARM_INS_LDRSBT
	ARM_INS_LDRSH = C.ARM_INS_LDRSH
	ARM_INS_LDRSHT = C.ARM_INS_LDRSHT
	ARM_INS_LDRT = C.ARM_INS_LDRT
	ARM_INS_LSL = C.ARM_INS_LSL
	ARM_INS_LSR = C.ARM_INS_LSR
	ARM_INS_MCR = C.ARM_INS_MCR
	ARM_INS_MCR2 = C.ARM_INS_MCR2
	ARM_INS_MCRR = C.ARM_INS_MCRR
	ARM_INS_MCRR2 = C.ARM_INS_MCRR2
	ARM_INS_MLA = C.ARM_INS_MLA
	ARM_INS_MLS = C.ARM_INS_MLS
	ARM_INS_MOV = C.ARM_INS_MOV
	ARM_INS_MOVS = C.ARM_INS_MOVS
	ARM_INS_MOVT = C.ARM_INS_MOVT
	ARM_INS_MOVW = C.ARM_INS_MOVW
	ARM_INS_MRC = C.ARM_INS_MRC
	ARM_INS_MRC2 = C.ARM_INS_MRC2
	ARM_INS_MRRC = C.ARM_INS_MRRC
	ARM_INS_MRRC2 = C.ARM_INS_MRRC2
	ARM_INS_MRS = C.ARM_INS_MRS
	ARM_INS_MSR = C.ARM_INS_MSR
	ARM_INS_MUL = C.ARM_INS_MUL
	ARM_INS_MVN = C.ARM_INS_MVN
	ARM_INS_NEG = C.ARM_INS_NEG
	ARM_INS_NOP = C.ARM_INS_NOP
	ARM_INS_ORN = C.ARM_INS_ORN
	ARM_INS_ORR = C.ARM_INS_ORR
	ARM_INS_PKHBT = C.ARM_INS_PKHBT
	ARM_INS_PKHTB = C.ARM_INS_PKHTB
	ARM_INS_PLD = C.ARM_INS_PLD
	ARM_INS_PLDW = C.ARM_INS_PLDW
	ARM_INS_PLI = C.ARM_INS_PLI
	ARM_INS_POP = C.ARM_INS_POP
	ARM_INS_PUSH = C.ARM_INS_PUSH
	ARM_INS_QADD = C.ARM_INS_QADD
	ARM_INS_QADD16 = C.ARM_INS_QADD16
	ARM_INS_QADD8 = C.ARM_INS_QADD8
	ARM_INS_QASX = C.ARM_INS_QASX
	ARM_INS_QDADD = C.ARM_INS_QDADD
	ARM_INS_QDSUB = C.ARM_INS_QDSUB
	ARM_INS_QSAX = C.ARM_INS_QSAX
	ARM_INS_QSUB = C.ARM_INS_QSUB
	ARM_INS_QSUB16 = C.ARM_INS_QSUB16
	ARM_INS_QSUB8 = C.ARM_INS_QSUB8
	ARM_INS_RBIT = C.ARM_INS_RBIT
	ARM_INS_REV = C.ARM_INS_REV
	ARM_INS_REV16 = C.ARM_INS_REV16
	ARM_INS_REVSH = C.ARM_INS_REVSH
	ARM_INS_RFEDA = C.ARM_INS_RFEDA
	ARM_INS_RFEDB = C.ARM_INS_RFEDB
	ARM_INS_RFEIA = C.ARM_INS_RFEIA
	ARM_INS_RFEIB = C.ARM_INS_RFEIB
	ARM_INS_ROR = C.ARM_INS_ROR
	ARM_INS_RRX = C.ARM_INS_RRX
	ARM_INS_RSB = C.ARM_INS_RSB
	ARM_INS_RSC = C.ARM_INS_RSC
	ARM_INS_SADD16 = C.ARM_INS_SADD16
	ARM_INS_SADD8 = C.ARM_INS_SADD8
	ARM_INS_SASX = C.ARM_INS_SASX
	ARM_INS_SBC = C.ARM_INS_SBC
	ARM_INS_SBFX = C.ARM_INS_SBFX
	ARM_INS_SDIV = C.ARM_INS_SDIV
	ARM_INS_SEL = C.ARM_INS_SEL
	ARM_INS_SETEND = C.ARM_INS_SETEND
	ARM_INS_SETPAN = C.ARM_INS_SETPAN
	ARM_INS_SEV = C.ARM_INS_SEV
	ARM_INS_SEVL = C.ARM_INS_SEVL
	ARM_INS_SG = C.ARM_INS_SG
	ARM_INS_SHA1C = C.ARM_INS_SHA1C
	ARM_INS_SHA1H = C.ARM_INS_SHA1H
	ARM_INS_SHA1M = C.ARM_INS_SHA1M
	ARM_INS_SHA1P = C.ARM_INS_SHA1P
	ARM_INS_SHA1SU0 = C.ARM_INS_SHA1SU0
	ARM_INS_SHA1SU1 = C.ARM_INS_SHA1SU1
	ARM_INS_SHA256H = C.ARM_INS_SHA256H
	ARM_INS_SHA256H2 = C.ARM_INS_SHA256H2
	ARM_INS_SHA256SU0 = C.ARM_INS_SHA256SU0
	ARM_INS_SHA256SU1 = C.ARM_INS_SHA256SU1
	ARM_INS_SHADD16 = C.ARM_INS_SHADD16
	ARM_INS_SHADD8 = C.ARM_INS_SHADD8
	ARM_INS_SHASX = C.ARM_INS_SHASX
	ARM_INS_SHSAX = C.ARM_INS_SHSAX
	ARM_INS_SHSUB16 = C.ARM_INS_SHSUB16
	ARM_INS_SHSUB8 = C.ARM_INS_SHSUB8
	ARM_INS_SMC = C.ARM_INS_SMC
	ARM_INS_SMLABB = C.ARM_INS_SMLABB
	ARM_INS_SMLABT = C.ARM_INS_SMLABT
	ARM_INS_SMLAD = C.ARM_INS_SMLAD
	ARM_INS_SMLADX = C.ARM_INS_SMLADX
	ARM_INS_SMLAL = C.ARM_INS_SMLAL
	ARM_INS_SMLALBB = C.ARM_INS_SMLALBB
	ARM_INS_SMLALBT = C.ARM_INS_SMLALBT
	ARM_INS_SMLALD = C.ARM_INS_SMLALD
	ARM_INS_SMLALDX = C.ARM_INS_SMLALDX
	ARM_INS_SMLALTB = C.ARM_INS_SMLALTB
	ARM_INS_SMLALTT = C.ARM_INS_SMLALTT
	ARM_INS_SMLATB = C.ARM_INS_SMLATB
	ARM_INS_SMLATT = C.ARM_INS_SMLATT
	ARM_INS_SMLAWB = C.ARM_INS_SMLAWB
	ARM_INS_SMLAWT = C.ARM_INS_SMLAWT
	ARM_INS_SMLSD = C.ARM_INS_SMLSD
	ARM_INS_SMLSDX = C.ARM_INS_SMLSDX
	ARM_INS_SMLSLD = C.ARM_INS_SMLSLD
	ARM_INS_SMLSLDX = C.ARM_INS_SMLSLDX
	ARM_INS_SMMLA = C.ARM_INS_SMMLA
	ARM_INS_SMMLAR = C.ARM_INS_SMMLAR
	ARM_INS_SMMLS = C.ARM_INS_SMMLS
	ARM_INS_SMMLSR = C.ARM_INS_SMMLSR
	ARM_INS_SMMUL = C.ARM_INS_SMMUL
	ARM_INS_SMMULR = C.ARM_INS_SMMULR
	ARM_INS_SMUAD = C.ARM_INS_SMUAD
	ARM_INS_SMUADX = C.ARM_INS_SMUADX
	ARM_INS_SMULBB = C.ARM_INS_SMULBB
	ARM_INS_SMULBT = C.ARM_INS_SMULBT
	ARM_INS_SMULL = C.ARM_INS_SMULL
	ARM_INS_SMULTB = C.ARM_INS_SMULTB
	ARM_INS_SMULTT = C.ARM_INS_SMULTT
	ARM_INS_SMULWB = C.ARM_INS_SMULWB
	ARM_INS_SMULWT = C.ARM_INS_SMULWT
	ARM_INS_SMUSD = C.ARM_INS_SMUSD
	ARM_INS_SMUSDX = C.ARM_INS_SMUSDX
	ARM_INS_SRSDA = C.ARM_INS_SRSDA
	ARM_INS_SRSDB = C.ARM_INS_SRSDB
	ARM_INS_SRSIA = C.ARM_INS_SRSIA
	ARM_INS_SRSIB = C.ARM_INS_SRSIB
	ARM_INS_SSAT = C.ARM_INS_SSAT
	ARM_INS_SSAT16 = C.ARM_INS_SSAT16
	ARM_INS_SSAX = C.ARM_INS_SSAX
	ARM_INS_SSUB16 = C.ARM_INS_SSUB16
	ARM_INS_SSUB8 = C.ARM_INS_SSUB8
	ARM_INS_STC = C.ARM_INS_STC
	ARM_INS_STC2 = C.ARM_INS_STC2
	ARM_INS_STC2L = C.ARM_INS_STC2L
	ARM_INS_STCL = C.ARM_INS_STCL
	ARM_INS_STL = C.ARM_INS_STL
	ARM_INS_STLB = C.ARM_INS_STLB
	ARM_INS_STLEX = C.ARM_INS_STLEX
	ARM_INS_STLEXB = C.ARM_INS_STLEXB
	ARM_INS_STLEXD = C.ARM_INS_STLEXD
	ARM_INS_STLEXH = C.ARM_INS_STLEXH
	ARM_INS_STLH = C.ARM_INS_STLH
	ARM_INS_STM = C.ARM_INS_STM
	ARM_INS_STMDA = C.ARM_INS_STMDA
	ARM_INS_STMDB = C.ARM_INS_STMDB
	ARM_INS_STMIB = C.ARM_INS_STMIB
	ARM_INS_STR = C.ARM_INS_STR
	ARM_INS_STRB = C.ARM_INS_STRB
	ARM_INS_STRBT = C.ARM_INS_STRBT
	ARM_INS_STRD = C.ARM_INS_STRD
	ARM_INS_STREX = C.ARM_INS_STREX
	ARM_INS_STREXB = C.ARM_INS_STREXB
	ARM_INS_STREXD = C.ARM_INS_STREXD
	ARM_INS_STREXH = C.ARM_INS_STREXH
	ARM_INS_STRH = C.ARM_INS_STRH
	ARM_INS_STRHT = C.ARM_INS_STRHT
	ARM_INS_STRT = C.ARM_INS_STRT
	ARM_INS_SUB = C.ARM_INS_SUB
	ARM_INS_SUBS = C.ARM_INS_SUBS
	ARM_INS_SUBW = C.ARM_INS_SUBW
	ARM_INS_SVC = C.ARM_INS_SVC
	ARM_INS_SWP = C.ARM_INS_SWP
	ARM_INS_SWPB = C.ARM_INS_SWPB
	ARM_INS_SXTAB = C.ARM_INS_SXTAB
	ARM_INS_SXTAB16 = C.ARM_INS_SXTAB16
	ARM_INS_SXTAH = C.ARM_INS_SXTAH
	ARM_INS_SXTB = C.ARM_INS_SXTB
	ARM_INS_SXTB16 = C.ARM_INS_SXTB16
	ARM_INS_SXTH = C.ARM_INS_SXTH
	ARM_INS_TBB = C.ARM_INS_TBB
	ARM_INS_TBH = C.ARM_INS_TBH
	ARM_INS_TEQ = C.ARM_INS_TEQ
	ARM_INS_TRAP = C.ARM_INS_TRAP
	ARM_INS_TSB = C.ARM_INS_TSB
	ARM_INS_TST = C.ARM_INS_TST
	ARM_INS_TT = C.ARM_INS_TT
	ARM_INS_TTA = C.ARM_INS_TTA
	ARM_INS_TTAT = C.ARM_INS_TTAT
	ARM_INS_TTT = C.ARM_INS_TTT
	ARM_INS_UADD16 = C.ARM_INS_UADD16
	ARM_INS_UADD8 = C.ARM_INS_UADD8
	ARM_INS_UASX = C.ARM_INS_UASX
	ARM_INS_UBFX = C.ARM_INS_UBFX
	ARM_INS_UDF = C.ARM_INS_UDF
	ARM_INS_UDIV = C.ARM_INS_UDIV
	ARM_INS_UHADD16 = C.ARM_INS_UHADD16
	ARM_INS_UHADD8 = C.ARM_INS_UHADD8
	ARM_INS_UHASX = C.ARM_INS_UHASX
	ARM_INS_UHSAX = C.ARM_INS_UHSAX
	ARM_INS_UHSUB16 = C.ARM_INS_UHSUB16
	ARM_INS_UHSUB8 = C.ARM_INS_UHSUB8
	ARM_INS_UMAAL = C.ARM_INS_UMAAL
	ARM_INS_UMLAL = C.ARM_INS_UMLAL
	ARM_INS_UMULL = C.ARM_INS_UMULL
	ARM_INS_UQADD16 = C.ARM_INS_UQADD16
	ARM_INS_UQADD8 = C.ARM_INS_UQADD8
	ARM_INS_UQASX = C.ARM_INS_UQASX
	ARM_INS_UQSAX = C.ARM_INS_UQSAX
	ARM_INS_UQSUB16 = C.ARM_INS_UQSUB16
	ARM_INS_UQSUB8 = C.ARM_INS_UQSUB8
	ARM_INS_USAD8 = C.ARM_INS_USAD8
	ARM_INS_USADA8 = C.ARM_INS_USADA8
	ARM_INS_USAT = C.ARM_INS_USAT
	ARM_INS_USAT16 = C.ARM_INS_USAT16
	ARM_INS_USAX = C.ARM_INS_USAX
	ARM_INS_USUB16 = C.ARM_INS_USUB16
	ARM_INS_USUB8 = C.ARM_INS_USUB8
	ARM_INS_UXTAB = C.ARM_INS_UXTAB
	ARM_INS_UXTAB16 = C.ARM_INS_UXTAB16
	ARM_INS_UXTAH = C.ARM_INS_UXTAH
	ARM_INS_UXTB = C.ARM_INS_UXTB
	ARM_INS_UXTB16 = C.ARM_INS_UXTB16
	ARM_INS_UXTH = C.ARM_INS_UXTH
	ARM_INS_VABA = C.ARM_INS_VABA
	ARM_INS_VABAL = C.ARM_INS_VABAL
	ARM_INS_VABD = C.ARM_INS_VABD
	ARM_INS_VABDL = C.ARM_INS_VABDL
	ARM_INS_VABS = C.ARM_INS_VABS
	ARM_INS_VACGE = C.ARM_INS_VACGE
	ARM_INS_VACGT = C.ARM_INS_VACGT
	ARM_INS_VACLE = C.ARM_INS_VACLE
	ARM_INS_VACLT = C.ARM_INS_VACLT
	ARM_INS_VADD = C.ARM_INS_VADD
	ARM_INS_VADDHN = C.ARM_INS_VADDHN
	ARM_INS_VADDL = C.ARM_INS_VADDL
	ARM_INS_VADDW = C.ARM_INS_VADDW
	ARM_INS_VAND = C.ARM_INS_VAND
	ARM_INS_VBIC = C.ARM_INS_VBIC
	ARM_INS_VBIF = C.ARM_INS_VBIF
	ARM_INS_VBIT = C.ARM_INS_VBIT
	ARM_INS_VBSL = C.ARM_INS_VBSL
	ARM_INS_VCADD = C.ARM_INS_VCADD
	ARM_INS_VCEQ = C.ARM_INS_VCEQ
	ARM_INS_VCGE = C.ARM_INS_VCGE
	ARM_INS_VCGT = C.ARM_INS_VCGT
	ARM_INS_VCLE = C.ARM_INS_VCLE
	ARM_INS_VCLS = C.ARM_INS_VCLS
	ARM_INS_VCLT = C.ARM_INS_VCLT
	ARM_INS_VCLZ = C.ARM_INS_VCLZ
	ARM_INS_VCMLA = C.ARM_INS_VCMLA
	ARM_INS_VCMP = C.ARM_INS_VCMP
	ARM_INS_VCMPE = C.ARM_INS_VCMPE
	ARM_INS_VCNT = C.ARM_INS_VCNT
	ARM_INS_VCVT = C.ARM_INS_VCVT
	ARM_INS_VCVTA = C.ARM_INS_VCVTA
	ARM_INS_VCVTB = C.ARM_INS_VCVTB
	ARM_INS_VCVTM = C.ARM_INS_VCVTM
	ARM_INS_VCVTN = C.ARM_INS_VCVTN
	ARM_INS_VCVTP = C.ARM_INS_VCVTP
	ARM_INS_VCVTR = C.ARM_INS_VCVTR
	ARM_INS_VCVTT = C.ARM_INS_VCVTT
	ARM_INS_VDIV = C.ARM_INS_VDIV
	ARM_INS_VDUP = C.ARM_INS_VDUP
	ARM_INS_VEOR = C.ARM_INS_VEOR
	ARM_INS_VEXT = C.ARM_INS_VEXT
	ARM_INS_VFMA = C.ARM_INS_VFMA
	ARM_INS_VFMS = C.ARM_INS_VFMS
	ARM_INS_VFNMA = C.ARM_INS_VFNMA
	ARM_INS_VFNMS = C.ARM_INS_VFNMS
	ARM_INS_VHADD = C.ARM_INS_VHADD
	ARM_INS_VHSUB = C.ARM_INS_VHSUB
	ARM_INS_VINS = C.ARM_INS_VINS
	ARM_INS_VJCVT = C.ARM_INS_VJCVT
	ARM_INS_VLD1 = C.ARM_INS_VLD1
	ARM_INS_VLD2 = C.ARM_INS_VLD2
	ARM_INS_VLD3 = C.ARM_INS_VLD3
	ARM_INS_VLD4 = C.ARM_INS_VLD4
	ARM_INS_VLDMDB = C.ARM_INS_VLDMDB
	ARM_INS_VLDMIA = C.ARM_INS_VLDMIA
	ARM_INS_VLDR = C.ARM_INS_VLDR
	ARM_INS_VLLDM = C.ARM_INS_VLLDM
	ARM_INS_VLSTM = C.ARM_INS_VLSTM
	ARM_INS_VMAX = C.ARM_INS_VMAX
	ARM_INS_VMAXNM = C.ARM_INS_VMAXNM
	ARM_INS_VMIN = C.ARM_INS_VMIN
	ARM_INS_VMINNM = C.ARM_INS_VMINNM
	ARM_INS_VMLA = C.ARM_INS_VMLA
	ARM_INS_VMLAL = C.ARM_INS_VMLAL
	ARM_INS_VMLS = C.ARM_INS_VMLS
	ARM_INS_VMLSL = C.ARM_INS_VMLSL
	ARM_INS_VMOV = C.ARM_INS_VMOV
	ARM_INS_VMOVL = C.ARM_INS_VMOVL
	ARM_INS_VMOVN = C.ARM_INS_VMOVN
	ARM_INS_VMOVX = C.ARM_INS_VMOVX
	ARM_INS_VMRS = C.ARM_INS_VMRS
	ARM_INS_VMSR = C.ARM_INS_VMSR
	ARM_INS_VMUL = C.ARM_INS_VMUL
	ARM_INS_VMULL = C.ARM_INS_VMULL
	ARM_INS_VMVN = C.ARM_INS_VMVN
	ARM_INS_VNEG = C.ARM_INS_VNEG
	ARM_INS_VNMLA = C.ARM_INS_VNMLA
	ARM_INS_VNMLS = C.ARM_INS_VNMLS
	ARM_INS_VNMUL = C.ARM_INS_VNMUL
	ARM_INS_VORN = C.ARM_INS_VORN
	ARM_INS_VORR = C.ARM_INS_VORR
	ARM_INS_VPADAL = C.ARM_INS_VPADAL
	ARM_INS_VPADD = C.ARM_INS_VPADD
	ARM_INS_VPADDL = C.ARM_INS_VPADDL
	ARM_INS_VPMAX = C.ARM_INS_VPMAX
	ARM_INS_VPMIN = C.ARM_INS_VPMIN
	ARM_INS_VPOP = C.ARM_INS_VPOP
	ARM_INS_VPUSH = C.ARM_INS_VPUSH
	ARM_INS_VQABS = C.ARM_INS_VQABS
	ARM_INS_VQADD = C.ARM_INS_VQADD
	ARM_INS_VQDMLAL = C.ARM_INS_VQDMLAL
	ARM_INS_VQDMLSL = C.ARM_INS_VQDMLSL
	ARM_INS_VQDMULH = C.ARM_INS_VQDMULH
	ARM_INS_VQDMULL = C.ARM_INS_VQDMULL
	ARM_INS_VQMOVN = C.ARM_INS_VQMOVN
	ARM_INS_VQMOVUN = C.ARM_INS_VQMOVUN
	ARM_INS_VQNEG = C.ARM_INS_VQNEG
	ARM_INS_VQRDMLAH = C.ARM_INS_VQRDMLAH
	ARM_INS_VQRDMLSH = C.ARM_INS_VQRDMLSH
	ARM_INS_VQRDMULH = C.ARM_INS_VQRDMULH
	ARM_INS_VQRSHL = C.ARM_INS_VQRSHL
	ARM_INS_VQRSHRN = C.ARM_INS_VQRSHRN
	ARM_INS_VQRSHRUN = C.ARM_INS_VQRSHRUN
	ARM_INS_VQSHL = C.ARM_INS_VQSHL
	ARM_INS_VQSHLU = C.ARM_INS_VQSHLU
	ARM_INS_VQSHRN = C.ARM_INS_VQSHRN
	ARM_INS_VQSHRUN = C.ARM_INS_VQSHRUN
	ARM_INS_VQSUB = C.ARM_INS_VQSUB
	ARM_INS_VRADDHN = C.ARM_INS_VRADDHN
	ARM_INS_VRECPE = C.ARM_INS_VRECPE
	ARM_INS_VRECPS = C.ARM_INS_VRECPS
	ARM_INS_VREV16 = C.ARM_INS_VREV16
	ARM_INS_VREV32 = C.ARM_INS_VREV32
	ARM_INS_VREV64 = C.ARM_INS_VREV64
	ARM_INS_VRHADD = C.ARM_INS_VRHADD
	ARM_INS_VRINTA = C.ARM_INS_VRINTA
	ARM_INS_VRINTM = C.ARM_INS_VRINTM
	ARM_INS_VRINTN = C.ARM_INS_VRINTN
	ARM_INS_VRINTP = C.ARM_INS_VRINTP
	ARM_INS_VRINTR = C.ARM_INS_VRINTR
	ARM_INS_VRINTX = C.ARM_INS_VRINTX
	ARM_INS_VRINTZ = C.ARM_INS_VRINTZ
	ARM_INS_VRSHL = C.ARM_INS_VRSHL
	ARM_INS_VRSHR = C.ARM_INS_VRSHR
	ARM_INS_VRSHRN = C.ARM_INS_VRSHRN
	ARM_INS_VRSQRTE = C.ARM_INS_VRSQRTE
	ARM_INS_VRSQRTS = C.ARM_INS_VRSQRTS
	ARM_INS_VRSRA = C.ARM_INS_VRSRA
	ARM_INS_VRSUBHN = C.ARM_INS_VRSUBHN
	ARM_INS_VSDOT = C.ARM_INS_VSDOT
	ARM_INS_VSELEQ = C.ARM_INS_VSELEQ
	ARM_INS_VSELGE = C.ARM_INS_VSELGE
	ARM_INS_VSELGT = C.ARM_INS_VSELGT
	ARM_INS_VSELVS = C.ARM_INS_VSELVS
	ARM_INS_VSHL = C.ARM_INS_VSHL
	ARM_INS_VSHLL = C.ARM_INS_VSHLL
	ARM_INS_VSHR = C.ARM_INS_VSHR
	ARM_INS_VSHRN = C.ARM_INS_VSHRN
	ARM_INS_VSLI = C.ARM_INS_VSLI
	ARM_INS_VSQRT = C.ARM_INS_VSQRT
	ARM_INS_VSRA = C.ARM_INS_VSRA
	ARM_INS_VSRI = C.ARM_INS_VSRI
	ARM_INS_VST1 = C.ARM_INS_VST1
	ARM_INS_VST2 = C.ARM_INS_VST2
	ARM_INS_VST3 = C.ARM_INS_VST3
	ARM_INS_VST4 = C.ARM_INS_VST4
	ARM_INS_VSTMDB = C.ARM_INS_VSTMDB
	ARM_INS_VSTMIA = C.ARM_INS_VSTMIA
	ARM_INS_VSTR = C.ARM_INS_VSTR
	ARM_INS_VSUB = C.ARM_INS_VSUB
	ARM_INS_VSUBHN = C.ARM_INS_VSUBHN
	ARM_INS_VSUBL = C.ARM_INS_VSUBL
	ARM_INS_VSUBW = C.ARM_INS_VSUBW
	ARM_INS_VSWP = C.ARM_INS_VSWP
	ARM_INS_VTBL = C.ARM_INS_VTBL
	ARM_INS_VTBX = C.ARM_INS_VTBX
	ARM_INS_VTRN = C.ARM_INS_VTRN
	ARM_INS_VTST = C.ARM_INS_VTST
	ARM_INS_VUDOT = C.ARM_INS_VUDOT
	ARM_INS_VUZP = C.ARM_INS_VUZP
	ARM_INS_VZIP = C.ARM_INS_VZIP
	ARM_INS_WFE = C.ARM_INS_WFE
	ARM_INS_WFI = C.ARM_INS_WFI
	ARM_INS_YIELD = C.ARM_INS_YIELD
	ARM_INS_ENDING = C.ARM_INS_ENDING
)

const (
	ARM_GRP_INVALID = C.ARM_GRP_INVALID
	ARM_GRP_JUMP = C.ARM_GRP_JUMP
	ARM_GRP_CALL = C.ARM_GRP_CALL
	ARM_GRP_INT = C.ARM_GRP_INT
	ARM_GRP_PRIVILEGE = C.ARM_GRP_PRIVILEGE
	ARM_GRP_BRANCH_RELATIVE = C.ARM_GRP_BRANCH_RELATIVE
	ARM_GRP_CRYPTO = C.ARM_GRP_CRYPTO
	ARM_GRP_DATABARRIER = C.ARM_GRP_DATABARRIER
	ARM_GRP_DIVIDE = C.ARM_GRP_DIVIDE
	ARM_GRP_FPARMV8 = C.ARM_GRP_FPARMV8
	ARM_GRP_MULTPRO = C.ARM_GRP_MULTPRO
	ARM_GRP_NEON = C.ARM_GRP_NEON
	ARM_GRP_T2EXTRACTPACK = C.ARM_GRP_T2EXTRACTPACK
	ARM_GRP_THUMB2DSP = C.ARM_GRP_THUMB2DSP
	ARM_GRP_TRUSTZONE = C.ARM_GRP_TRUSTZONE
	ARM_GRP_V4T = C.ARM_GRP_V4T
	ARM_GRP_V5T = C.ARM_GRP_V5T
	ARM_GRP_V5TE = C.ARM_GRP_V5TE
	ARM_GRP_V6 = C.ARM_GRP_V6
	ARM_GRP_V6T2 = C.ARM_GRP_V6T2
	ARM_GRP_V7 = C.ARM_GRP_V7
	ARM_GRP_V8 = C.ARM_GRP_V8
	ARM_GRP_VFP2 = C.ARM_GRP_VFP2
	ARM_GRP_VFP3 = C.ARM_GRP_VFP3
	ARM_GRP_VFP4 = C.ARM_GRP_VFP4
	ARM_GRP_ARM = C.ARM_GRP_ARM
	ARM_GRP_MCLASS = C.ARM_GRP_MCLASS
	ARM_GRP_NOTMCLASS = C.ARM_GRP_NOTMCLASS
	ARM_GRP_THUMB = C.ARM_GRP_THUMB
	ARM_GRP_THUMB1ONLY = C.ARM_GRP_THUMB1ONLY
	ARM_GRP_THUMB2 = C.ARM_GRP_THUMB2
	ARM_GRP_PREV8 = C.ARM_GRP_PREV8
	ARM_GRP_FPVMLX = C.ARM_GRP_FPVMLX
	ARM_GRP_MULOPS = C.ARM_GRP_MULOPS
	ARM_GRP_CRC = C.ARM_GRP_CRC
	ARM_GRP_DPVFP = C.ARM_GRP_DPVFP
	ARM_GRP_V6M = C.ARM_GRP_V6M
	ARM_GRP_VIRTUALIZATION = C.ARM_GRP_VIRTUALIZATION
	ARM_GRP_ENDING = C.ARM_GRP_ENDING
)

