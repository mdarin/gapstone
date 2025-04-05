/*
Gapstone is a Go binding for the Capstone disassembly library.
For examples, try reading the *_test.go files.

	Library Author: Nguyen Anh Quynh
	Binding Author: Ben Nagy
	License: BSD style - see LICENSE file for details
    (c) 2013 COSEINC. All Rights Reserved.

    THIS FILE WAS AUTO-GENERATED -- DO NOT EDIT!
	Command: ./genconst capstone/bindings/python/capstone/
	Created at: 2025-04-05T12:41:57+00:00

*/

package gapstone

// #cgo LDFLAGS: -lcapstone
// #cgo freebsd CFLAGS: -I/usr/local/include
// #cgo freebsd LDFLAGS: -L/usr/local/lib
// #include <stdlib.h>
// #include <capstone/capstone.h>
import "C"

// For Capstone Engine. AUTO-GENERATED FILE, DO NOT EDIT [alpha_const.py]
// Operand type for instruction's operands
const (
	ALPHA_OP_INVALID = C.ALPHA_OP_INVALID
	ALPHA_OP_REG = C.ALPHA_OP_REG
	ALPHA_OP_IMM = C.ALPHA_OP_IMM
)

// Alpha registers
const (
	Alpha_REG_INVALID = C.Alpha_REG_INVALID
	Alpha_REG_F0 = C.Alpha_REG_F0
	Alpha_REG_F1 = C.Alpha_REG_F1
	Alpha_REG_F2 = C.Alpha_REG_F2
	Alpha_REG_F3 = C.Alpha_REG_F3
	Alpha_REG_F4 = C.Alpha_REG_F4
	Alpha_REG_F5 = C.Alpha_REG_F5
	Alpha_REG_F6 = C.Alpha_REG_F6
	Alpha_REG_F7 = C.Alpha_REG_F7
	Alpha_REG_F8 = C.Alpha_REG_F8
	Alpha_REG_F9 = C.Alpha_REG_F9
	Alpha_REG_F10 = C.Alpha_REG_F10
	Alpha_REG_F11 = C.Alpha_REG_F11
	Alpha_REG_F12 = C.Alpha_REG_F12
	Alpha_REG_F13 = C.Alpha_REG_F13
	Alpha_REG_F14 = C.Alpha_REG_F14
	Alpha_REG_F15 = C.Alpha_REG_F15
	Alpha_REG_F16 = C.Alpha_REG_F16
	Alpha_REG_F17 = C.Alpha_REG_F17
	Alpha_REG_F18 = C.Alpha_REG_F18
	Alpha_REG_F19 = C.Alpha_REG_F19
	Alpha_REG_F20 = C.Alpha_REG_F20
	Alpha_REG_F21 = C.Alpha_REG_F21
	Alpha_REG_F22 = C.Alpha_REG_F22
	Alpha_REG_F23 = C.Alpha_REG_F23
	Alpha_REG_F24 = C.Alpha_REG_F24
	Alpha_REG_F25 = C.Alpha_REG_F25
	Alpha_REG_F26 = C.Alpha_REG_F26
	Alpha_REG_F27 = C.Alpha_REG_F27
	Alpha_REG_F28 = C.Alpha_REG_F28
	Alpha_REG_F29 = C.Alpha_REG_F29
	Alpha_REG_F30 = C.Alpha_REG_F30
	Alpha_REG_F31 = C.Alpha_REG_F31
	Alpha_REG_R0 = C.Alpha_REG_R0
	Alpha_REG_R1 = C.Alpha_REG_R1
	Alpha_REG_R2 = C.Alpha_REG_R2
	Alpha_REG_R3 = C.Alpha_REG_R3
	Alpha_REG_R4 = C.Alpha_REG_R4
	Alpha_REG_R5 = C.Alpha_REG_R5
	Alpha_REG_R6 = C.Alpha_REG_R6
	Alpha_REG_R7 = C.Alpha_REG_R7
	Alpha_REG_R8 = C.Alpha_REG_R8
	Alpha_REG_R9 = C.Alpha_REG_R9
	Alpha_REG_R10 = C.Alpha_REG_R10
	Alpha_REG_R11 = C.Alpha_REG_R11
	Alpha_REG_R12 = C.Alpha_REG_R12
	Alpha_REG_R13 = C.Alpha_REG_R13
	Alpha_REG_R14 = C.Alpha_REG_R14
	Alpha_REG_R15 = C.Alpha_REG_R15
	Alpha_REG_R16 = C.Alpha_REG_R16
	Alpha_REG_R17 = C.Alpha_REG_R17
	Alpha_REG_R18 = C.Alpha_REG_R18
	Alpha_REG_R19 = C.Alpha_REG_R19
	Alpha_REG_R20 = C.Alpha_REG_R20
	Alpha_REG_R21 = C.Alpha_REG_R21
	Alpha_REG_R22 = C.Alpha_REG_R22
	Alpha_REG_R23 = C.Alpha_REG_R23
	Alpha_REG_R24 = C.Alpha_REG_R24
	Alpha_REG_R25 = C.Alpha_REG_R25
	Alpha_REG_R26 = C.Alpha_REG_R26
	Alpha_REG_R27 = C.Alpha_REG_R27
	Alpha_REG_R28 = C.Alpha_REG_R28
	Alpha_REG_R29 = C.Alpha_REG_R29
	Alpha_REG_R30 = C.Alpha_REG_R30
	Alpha_REG_R31 = C.Alpha_REG_R31
	Alpha_REG_ENDING = C.Alpha_REG_ENDING
)

// Alpha instruction
const (
	Alpha_INS_INVALID = C.Alpha_INS_INVALID
	Alpha_INS_ADDL = C.Alpha_INS_ADDL
	Alpha_INS_ADDQ = C.Alpha_INS_ADDQ
	Alpha_INS_ADDSsSU = C.Alpha_INS_ADDSsSU
	Alpha_INS_ADDTsSU = C.Alpha_INS_ADDTsSU
	Alpha_INS_AND = C.Alpha_INS_AND
	Alpha_INS_BEQ = C.Alpha_INS_BEQ
	Alpha_INS_BGE = C.Alpha_INS_BGE
	Alpha_INS_BGT = C.Alpha_INS_BGT
	Alpha_INS_BIC = C.Alpha_INS_BIC
	Alpha_INS_BIS = C.Alpha_INS_BIS
	Alpha_INS_BLBC = C.Alpha_INS_BLBC
	Alpha_INS_BLBS = C.Alpha_INS_BLBS
	Alpha_INS_BLE = C.Alpha_INS_BLE
	Alpha_INS_BLT = C.Alpha_INS_BLT
	Alpha_INS_BNE = C.Alpha_INS_BNE
	Alpha_INS_BR = C.Alpha_INS_BR
	Alpha_INS_BSR = C.Alpha_INS_BSR
	Alpha_INS_CMOVEQ = C.Alpha_INS_CMOVEQ
	Alpha_INS_CMOVGE = C.Alpha_INS_CMOVGE
	Alpha_INS_CMOVGT = C.Alpha_INS_CMOVGT
	Alpha_INS_CMOVLBC = C.Alpha_INS_CMOVLBC
	Alpha_INS_CMOVLBS = C.Alpha_INS_CMOVLBS
	Alpha_INS_CMOVLE = C.Alpha_INS_CMOVLE
	Alpha_INS_CMOVLT = C.Alpha_INS_CMOVLT
	Alpha_INS_CMOVNE = C.Alpha_INS_CMOVNE
	Alpha_INS_CMPBGE = C.Alpha_INS_CMPBGE
	Alpha_INS_CMPEQ = C.Alpha_INS_CMPEQ
	Alpha_INS_CMPLE = C.Alpha_INS_CMPLE
	Alpha_INS_CMPLT = C.Alpha_INS_CMPLT
	Alpha_INS_CMPTEQsSU = C.Alpha_INS_CMPTEQsSU
	Alpha_INS_CMPTLEsSU = C.Alpha_INS_CMPTLEsSU
	Alpha_INS_CMPTLTsSU = C.Alpha_INS_CMPTLTsSU
	Alpha_INS_CMPTUNsSU = C.Alpha_INS_CMPTUNsSU
	Alpha_INS_CMPULE = C.Alpha_INS_CMPULE
	Alpha_INS_CMPULT = C.Alpha_INS_CMPULT
	Alpha_INS_COND_BRANCH = C.Alpha_INS_COND_BRANCH
	Alpha_INS_CPYSE = C.Alpha_INS_CPYSE
	Alpha_INS_CPYSN = C.Alpha_INS_CPYSN
	Alpha_INS_CPYS = C.Alpha_INS_CPYS
	Alpha_INS_CTLZ = C.Alpha_INS_CTLZ
	Alpha_INS_CTPOP = C.Alpha_INS_CTPOP
	Alpha_INS_CTTZ = C.Alpha_INS_CTTZ
	Alpha_INS_CVTQSsSUI = C.Alpha_INS_CVTQSsSUI
	Alpha_INS_CVTQTsSUI = C.Alpha_INS_CVTQTsSUI
	Alpha_INS_CVTSTsS = C.Alpha_INS_CVTSTsS
	Alpha_INS_CVTTQsSVC = C.Alpha_INS_CVTTQsSVC
	Alpha_INS_CVTTSsSUI = C.Alpha_INS_CVTTSsSUI
	Alpha_INS_DIVSsSU = C.Alpha_INS_DIVSsSU
	Alpha_INS_DIVTsSU = C.Alpha_INS_DIVTsSU
	Alpha_INS_ECB = C.Alpha_INS_ECB
	Alpha_INS_EQV = C.Alpha_INS_EQV
	Alpha_INS_EXCB = C.Alpha_INS_EXCB
	Alpha_INS_EXTBL = C.Alpha_INS_EXTBL
	Alpha_INS_EXTLH = C.Alpha_INS_EXTLH
	Alpha_INS_EXTLL = C.Alpha_INS_EXTLL
	Alpha_INS_EXTQH = C.Alpha_INS_EXTQH
	Alpha_INS_EXTQL = C.Alpha_INS_EXTQL
	Alpha_INS_EXTWH = C.Alpha_INS_EXTWH
	Alpha_INS_EXTWL = C.Alpha_INS_EXTWL
	Alpha_INS_FBEQ = C.Alpha_INS_FBEQ
	Alpha_INS_FBGE = C.Alpha_INS_FBGE
	Alpha_INS_FBGT = C.Alpha_INS_FBGT
	Alpha_INS_FBLE = C.Alpha_INS_FBLE
	Alpha_INS_FBLT = C.Alpha_INS_FBLT
	Alpha_INS_FBNE = C.Alpha_INS_FBNE
	Alpha_INS_FCMOVEQ = C.Alpha_INS_FCMOVEQ
	Alpha_INS_FCMOVGE = C.Alpha_INS_FCMOVGE
	Alpha_INS_FCMOVGT = C.Alpha_INS_FCMOVGT
	Alpha_INS_FCMOVLE = C.Alpha_INS_FCMOVLE
	Alpha_INS_FCMOVLT = C.Alpha_INS_FCMOVLT
	Alpha_INS_FCMOVNE = C.Alpha_INS_FCMOVNE
	Alpha_INS_FETCH = C.Alpha_INS_FETCH
	Alpha_INS_FETCH_M = C.Alpha_INS_FETCH_M
	Alpha_INS_FTOIS = C.Alpha_INS_FTOIS
	Alpha_INS_FTOIT = C.Alpha_INS_FTOIT
	Alpha_INS_INSBL = C.Alpha_INS_INSBL
	Alpha_INS_INSLH = C.Alpha_INS_INSLH
	Alpha_INS_INSLL = C.Alpha_INS_INSLL
	Alpha_INS_INSQH = C.Alpha_INS_INSQH
	Alpha_INS_INSQL = C.Alpha_INS_INSQL
	Alpha_INS_INSWH = C.Alpha_INS_INSWH
	Alpha_INS_INSWL = C.Alpha_INS_INSWL
	Alpha_INS_ITOFS = C.Alpha_INS_ITOFS
	Alpha_INS_ITOFT = C.Alpha_INS_ITOFT
	Alpha_INS_JMP = C.Alpha_INS_JMP
	Alpha_INS_JSR = C.Alpha_INS_JSR
	Alpha_INS_JSR_COROUTINE = C.Alpha_INS_JSR_COROUTINE
	Alpha_INS_LDA = C.Alpha_INS_LDA
	Alpha_INS_LDAH = C.Alpha_INS_LDAH
	Alpha_INS_LDBU = C.Alpha_INS_LDBU
	Alpha_INS_LDL = C.Alpha_INS_LDL
	Alpha_INS_LDL_L = C.Alpha_INS_LDL_L
	Alpha_INS_LDQ = C.Alpha_INS_LDQ
	Alpha_INS_LDQ_L = C.Alpha_INS_LDQ_L
	Alpha_INS_LDQ_U = C.Alpha_INS_LDQ_U
	Alpha_INS_LDS = C.Alpha_INS_LDS
	Alpha_INS_LDT = C.Alpha_INS_LDT
	Alpha_INS_LDWU = C.Alpha_INS_LDWU
	Alpha_INS_MB = C.Alpha_INS_MB
	Alpha_INS_MSKBL = C.Alpha_INS_MSKBL
	Alpha_INS_MSKLH = C.Alpha_INS_MSKLH
	Alpha_INS_MSKLL = C.Alpha_INS_MSKLL
	Alpha_INS_MSKQH = C.Alpha_INS_MSKQH
	Alpha_INS_MSKQL = C.Alpha_INS_MSKQL
	Alpha_INS_MSKWH = C.Alpha_INS_MSKWH
	Alpha_INS_MSKWL = C.Alpha_INS_MSKWL
	Alpha_INS_MULL = C.Alpha_INS_MULL
	Alpha_INS_MULQ = C.Alpha_INS_MULQ
	Alpha_INS_MULSsSU = C.Alpha_INS_MULSsSU
	Alpha_INS_MULTsSU = C.Alpha_INS_MULTsSU
	Alpha_INS_ORNOT = C.Alpha_INS_ORNOT
	Alpha_INS_RC = C.Alpha_INS_RC
	Alpha_INS_RET = C.Alpha_INS_RET
	Alpha_INS_RPCC = C.Alpha_INS_RPCC
	Alpha_INS_RS = C.Alpha_INS_RS
	Alpha_INS_S4ADDL = C.Alpha_INS_S4ADDL
	Alpha_INS_S4ADDQ = C.Alpha_INS_S4ADDQ
	Alpha_INS_S4SUBL = C.Alpha_INS_S4SUBL
	Alpha_INS_S4SUBQ = C.Alpha_INS_S4SUBQ
	Alpha_INS_S8ADDL = C.Alpha_INS_S8ADDL
	Alpha_INS_S8ADDQ = C.Alpha_INS_S8ADDQ
	Alpha_INS_S8SUBL = C.Alpha_INS_S8SUBL
	Alpha_INS_S8SUBQ = C.Alpha_INS_S8SUBQ
	Alpha_INS_SEXTB = C.Alpha_INS_SEXTB
	Alpha_INS_SEXTW = C.Alpha_INS_SEXTW
	Alpha_INS_SLL = C.Alpha_INS_SLL
	Alpha_INS_SQRTSsSU = C.Alpha_INS_SQRTSsSU
	Alpha_INS_SQRTTsSU = C.Alpha_INS_SQRTTsSU
	Alpha_INS_SRA = C.Alpha_INS_SRA
	Alpha_INS_SRL = C.Alpha_INS_SRL
	Alpha_INS_STB = C.Alpha_INS_STB
	Alpha_INS_STL = C.Alpha_INS_STL
	Alpha_INS_STL_C = C.Alpha_INS_STL_C
	Alpha_INS_STQ = C.Alpha_INS_STQ
	Alpha_INS_STQ_C = C.Alpha_INS_STQ_C
	Alpha_INS_STQ_U = C.Alpha_INS_STQ_U
	Alpha_INS_STS = C.Alpha_INS_STS
	Alpha_INS_STT = C.Alpha_INS_STT
	Alpha_INS_STW = C.Alpha_INS_STW
	Alpha_INS_SUBL = C.Alpha_INS_SUBL
	Alpha_INS_SUBQ = C.Alpha_INS_SUBQ
	Alpha_INS_SUBSsSU = C.Alpha_INS_SUBSsSU
	Alpha_INS_SUBTsSU = C.Alpha_INS_SUBTsSU
	Alpha_INS_TRAPB = C.Alpha_INS_TRAPB
	Alpha_INS_UMULH = C.Alpha_INS_UMULH
	Alpha_INS_WH64 = C.Alpha_INS_WH64
	Alpha_INS_WH64EN = C.Alpha_INS_WH64EN
	Alpha_INS_WMB = C.Alpha_INS_WMB
	Alpha_INS_XOR = C.Alpha_INS_XOR
	Alpha_INS_ZAPNOT = C.Alpha_INS_ZAPNOT
	ALPHA_INS_ENDING = C.ALPHA_INS_ENDING
)

// Group of Alpha instructions
const (
	Alpha_GRP_INVALID = C.Alpha_GRP_INVALID
)

// Generic groups
const (
	Alpha_GRP_CALL = C.Alpha_GRP_CALL
	Alpha_GRP_JUMP = C.Alpha_GRP_JUMP
	Alpha_GRP_BRANCH_RELATIVE = C.Alpha_GRP_BRANCH_RELATIVE
	Alpha_GRP_ENDING = C.Alpha_GRP_ENDING
)

