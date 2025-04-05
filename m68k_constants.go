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

// For Capstone Engine. AUTO-GENERATED FILE, DO NOT EDIT [m68k_const.py]
const (
	M68K_OPERAND_COUNT = C.M68K_OPERAND_COUNT
)

const (
	M68K_REG_INVALID = C.M68K_REG_INVALID
	M68K_REG_D0 = C.M68K_REG_D0
	M68K_REG_D1 = C.M68K_REG_D1
	M68K_REG_D2 = C.M68K_REG_D2
	M68K_REG_D3 = C.M68K_REG_D3
	M68K_REG_D4 = C.M68K_REG_D4
	M68K_REG_D5 = C.M68K_REG_D5
	M68K_REG_D6 = C.M68K_REG_D6
	M68K_REG_D7 = C.M68K_REG_D7
	M68K_REG_A0 = C.M68K_REG_A0
	M68K_REG_A1 = C.M68K_REG_A1
	M68K_REG_A2 = C.M68K_REG_A2
	M68K_REG_A3 = C.M68K_REG_A3
	M68K_REG_A4 = C.M68K_REG_A4
	M68K_REG_A5 = C.M68K_REG_A5
	M68K_REG_A6 = C.M68K_REG_A6
	M68K_REG_A7 = C.M68K_REG_A7
	M68K_REG_FP0 = C.M68K_REG_FP0
	M68K_REG_FP1 = C.M68K_REG_FP1
	M68K_REG_FP2 = C.M68K_REG_FP2
	M68K_REG_FP3 = C.M68K_REG_FP3
	M68K_REG_FP4 = C.M68K_REG_FP4
	M68K_REG_FP5 = C.M68K_REG_FP5
	M68K_REG_FP6 = C.M68K_REG_FP6
	M68K_REG_FP7 = C.M68K_REG_FP7
	M68K_REG_PC = C.M68K_REG_PC
	M68K_REG_SR = C.M68K_REG_SR
	M68K_REG_CCR = C.M68K_REG_CCR
	M68K_REG_SFC = C.M68K_REG_SFC
	M68K_REG_DFC = C.M68K_REG_DFC
	M68K_REG_USP = C.M68K_REG_USP
	M68K_REG_VBR = C.M68K_REG_VBR
	M68K_REG_CACR = C.M68K_REG_CACR
	M68K_REG_CAAR = C.M68K_REG_CAAR
	M68K_REG_MSP = C.M68K_REG_MSP
	M68K_REG_ISP = C.M68K_REG_ISP
	M68K_REG_TC = C.M68K_REG_TC
	M68K_REG_ITT0 = C.M68K_REG_ITT0
	M68K_REG_ITT1 = C.M68K_REG_ITT1
	M68K_REG_DTT0 = C.M68K_REG_DTT0
	M68K_REG_DTT1 = C.M68K_REG_DTT1
	M68K_REG_MMUSR = C.M68K_REG_MMUSR
	M68K_REG_URP = C.M68K_REG_URP
	M68K_REG_SRP = C.M68K_REG_SRP
	M68K_REG_FPCR = C.M68K_REG_FPCR
	M68K_REG_FPSR = C.M68K_REG_FPSR
	M68K_REG_FPIAR = C.M68K_REG_FPIAR
	M68K_REG_ENDING = C.M68K_REG_ENDING
)

const (
	M68K_AM_NONE = C.M68K_AM_NONE
	M68K_AM_REG_DIRECT_DATA = C.M68K_AM_REG_DIRECT_DATA
	M68K_AM_REG_DIRECT_ADDR = C.M68K_AM_REG_DIRECT_ADDR
	M68K_AM_REGI_ADDR = C.M68K_AM_REGI_ADDR
	M68K_AM_REGI_ADDR_POST_INC = C.M68K_AM_REGI_ADDR_POST_INC
	M68K_AM_REGI_ADDR_PRE_DEC = C.M68K_AM_REGI_ADDR_PRE_DEC
	M68K_AM_REGI_ADDR_DISP = C.M68K_AM_REGI_ADDR_DISP
	M68K_AM_AREGI_INDEX_8_BIT_DISP = C.M68K_AM_AREGI_INDEX_8_BIT_DISP
	M68K_AM_AREGI_INDEX_BASE_DISP = C.M68K_AM_AREGI_INDEX_BASE_DISP
	M68K_AM_MEMI_POST_INDEX = C.M68K_AM_MEMI_POST_INDEX
	M68K_AM_MEMI_PRE_INDEX = C.M68K_AM_MEMI_PRE_INDEX
	M68K_AM_PCI_DISP = C.M68K_AM_PCI_DISP
	M68K_AM_PCI_INDEX_8_BIT_DISP = C.M68K_AM_PCI_INDEX_8_BIT_DISP
	M68K_AM_PCI_INDEX_BASE_DISP = C.M68K_AM_PCI_INDEX_BASE_DISP
	M68K_AM_PC_MEMI_POST_INDEX = C.M68K_AM_PC_MEMI_POST_INDEX
	M68K_AM_PC_MEMI_PRE_INDEX = C.M68K_AM_PC_MEMI_PRE_INDEX
	M68K_AM_ABSOLUTE_DATA_SHORT = C.M68K_AM_ABSOLUTE_DATA_SHORT
	M68K_AM_ABSOLUTE_DATA_LONG = C.M68K_AM_ABSOLUTE_DATA_LONG
	M68K_AM_IMMEDIATE = C.M68K_AM_IMMEDIATE
	M68K_AM_BRANCH_DISPLACEMENT = C.M68K_AM_BRANCH_DISPLACEMENT
)

const (
	M68K_OP_INVALID = C.M68K_OP_INVALID
	M68K_OP_REG = C.M68K_OP_REG
	M68K_OP_IMM = C.M68K_OP_IMM
	M68K_OP_MEM = C.M68K_OP_MEM
	M68K_OP_FP_SINGLE = C.M68K_OP_FP_SINGLE
	M68K_OP_FP_DOUBLE = C.M68K_OP_FP_DOUBLE
	M68K_OP_REG_BITS = C.M68K_OP_REG_BITS
	M68K_OP_REG_PAIR = C.M68K_OP_REG_PAIR
	M68K_OP_BR_DISP = C.M68K_OP_BR_DISP
)

const (
	M68K_OP_BR_DISP_SIZE_INVALID = C.M68K_OP_BR_DISP_SIZE_INVALID
	M68K_OP_BR_DISP_SIZE_BYTE = C.M68K_OP_BR_DISP_SIZE_BYTE
	M68K_OP_BR_DISP_SIZE_WORD = C.M68K_OP_BR_DISP_SIZE_WORD
	M68K_OP_BR_DISP_SIZE_LONG = C.M68K_OP_BR_DISP_SIZE_LONG
)

const (
	M68K_CPU_SIZE_NONE = C.M68K_CPU_SIZE_NONE
	M68K_CPU_SIZE_BYTE = C.M68K_CPU_SIZE_BYTE
	M68K_CPU_SIZE_WORD = C.M68K_CPU_SIZE_WORD
	M68K_CPU_SIZE_LONG = C.M68K_CPU_SIZE_LONG
)

const (
	M68K_FPU_SIZE_NONE = C.M68K_FPU_SIZE_NONE
	M68K_FPU_SIZE_SINGLE = C.M68K_FPU_SIZE_SINGLE
	M68K_FPU_SIZE_DOUBLE = C.M68K_FPU_SIZE_DOUBLE
	M68K_FPU_SIZE_EXTENDED = C.M68K_FPU_SIZE_EXTENDED
)

const (
	M68K_SIZE_TYPE_INVALID = C.M68K_SIZE_TYPE_INVALID
	M68K_SIZE_TYPE_CPU = C.M68K_SIZE_TYPE_CPU
	M68K_SIZE_TYPE_FPU = C.M68K_SIZE_TYPE_FPU
)

const (
	M68K_INS_INVALID = C.M68K_INS_INVALID
	M68K_INS_ABCD = C.M68K_INS_ABCD
	M68K_INS_ADD = C.M68K_INS_ADD
	M68K_INS_ADDA = C.M68K_INS_ADDA
	M68K_INS_ADDI = C.M68K_INS_ADDI
	M68K_INS_ADDQ = C.M68K_INS_ADDQ
	M68K_INS_ADDX = C.M68K_INS_ADDX
	M68K_INS_AND = C.M68K_INS_AND
	M68K_INS_ANDI = C.M68K_INS_ANDI
	M68K_INS_ASL = C.M68K_INS_ASL
	M68K_INS_ASR = C.M68K_INS_ASR
	M68K_INS_BHS = C.M68K_INS_BHS
	M68K_INS_BLO = C.M68K_INS_BLO
	M68K_INS_BHI = C.M68K_INS_BHI
	M68K_INS_BLS = C.M68K_INS_BLS
	M68K_INS_BCC = C.M68K_INS_BCC
	M68K_INS_BCS = C.M68K_INS_BCS
	M68K_INS_BNE = C.M68K_INS_BNE
	M68K_INS_BEQ = C.M68K_INS_BEQ
	M68K_INS_BVC = C.M68K_INS_BVC
	M68K_INS_BVS = C.M68K_INS_BVS
	M68K_INS_BPL = C.M68K_INS_BPL
	M68K_INS_BMI = C.M68K_INS_BMI
	M68K_INS_BGE = C.M68K_INS_BGE
	M68K_INS_BLT = C.M68K_INS_BLT
	M68K_INS_BGT = C.M68K_INS_BGT
	M68K_INS_BLE = C.M68K_INS_BLE
	M68K_INS_BRA = C.M68K_INS_BRA
	M68K_INS_BSR = C.M68K_INS_BSR
	M68K_INS_BCHG = C.M68K_INS_BCHG
	M68K_INS_BCLR = C.M68K_INS_BCLR
	M68K_INS_BSET = C.M68K_INS_BSET
	M68K_INS_BTST = C.M68K_INS_BTST
	M68K_INS_BFCHG = C.M68K_INS_BFCHG
	M68K_INS_BFCLR = C.M68K_INS_BFCLR
	M68K_INS_BFEXTS = C.M68K_INS_BFEXTS
	M68K_INS_BFEXTU = C.M68K_INS_BFEXTU
	M68K_INS_BFFFO = C.M68K_INS_BFFFO
	M68K_INS_BFINS = C.M68K_INS_BFINS
	M68K_INS_BFSET = C.M68K_INS_BFSET
	M68K_INS_BFTST = C.M68K_INS_BFTST
	M68K_INS_BKPT = C.M68K_INS_BKPT
	M68K_INS_CALLM = C.M68K_INS_CALLM
	M68K_INS_CAS = C.M68K_INS_CAS
	M68K_INS_CAS2 = C.M68K_INS_CAS2
	M68K_INS_CHK = C.M68K_INS_CHK
	M68K_INS_CHK2 = C.M68K_INS_CHK2
	M68K_INS_CLR = C.M68K_INS_CLR
	M68K_INS_CMP = C.M68K_INS_CMP
	M68K_INS_CMPA = C.M68K_INS_CMPA
	M68K_INS_CMPI = C.M68K_INS_CMPI
	M68K_INS_CMPM = C.M68K_INS_CMPM
	M68K_INS_CMP2 = C.M68K_INS_CMP2
	M68K_INS_CINVL = C.M68K_INS_CINVL
	M68K_INS_CINVP = C.M68K_INS_CINVP
	M68K_INS_CINVA = C.M68K_INS_CINVA
	M68K_INS_CPUSHL = C.M68K_INS_CPUSHL
	M68K_INS_CPUSHP = C.M68K_INS_CPUSHP
	M68K_INS_CPUSHA = C.M68K_INS_CPUSHA
	M68K_INS_DBT = C.M68K_INS_DBT
	M68K_INS_DBF = C.M68K_INS_DBF
	M68K_INS_DBHI = C.M68K_INS_DBHI
	M68K_INS_DBLS = C.M68K_INS_DBLS
	M68K_INS_DBCC = C.M68K_INS_DBCC
	M68K_INS_DBCS = C.M68K_INS_DBCS
	M68K_INS_DBNE = C.M68K_INS_DBNE
	M68K_INS_DBEQ = C.M68K_INS_DBEQ
	M68K_INS_DBVC = C.M68K_INS_DBVC
	M68K_INS_DBVS = C.M68K_INS_DBVS
	M68K_INS_DBPL = C.M68K_INS_DBPL
	M68K_INS_DBMI = C.M68K_INS_DBMI
	M68K_INS_DBGE = C.M68K_INS_DBGE
	M68K_INS_DBLT = C.M68K_INS_DBLT
	M68K_INS_DBGT = C.M68K_INS_DBGT
	M68K_INS_DBLE = C.M68K_INS_DBLE
	M68K_INS_DBRA = C.M68K_INS_DBRA
	M68K_INS_DIVS = C.M68K_INS_DIVS
	M68K_INS_DIVSL = C.M68K_INS_DIVSL
	M68K_INS_DIVU = C.M68K_INS_DIVU
	M68K_INS_DIVUL = C.M68K_INS_DIVUL
	M68K_INS_EOR = C.M68K_INS_EOR
	M68K_INS_EORI = C.M68K_INS_EORI
	M68K_INS_EXG = C.M68K_INS_EXG
	M68K_INS_EXT = C.M68K_INS_EXT
	M68K_INS_EXTB = C.M68K_INS_EXTB
	M68K_INS_FABS = C.M68K_INS_FABS
	M68K_INS_FSABS = C.M68K_INS_FSABS
	M68K_INS_FDABS = C.M68K_INS_FDABS
	M68K_INS_FACOS = C.M68K_INS_FACOS
	M68K_INS_FADD = C.M68K_INS_FADD
	M68K_INS_FSADD = C.M68K_INS_FSADD
	M68K_INS_FDADD = C.M68K_INS_FDADD
	M68K_INS_FASIN = C.M68K_INS_FASIN
	M68K_INS_FATAN = C.M68K_INS_FATAN
	M68K_INS_FATANH = C.M68K_INS_FATANH
	M68K_INS_FBF = C.M68K_INS_FBF
	M68K_INS_FBEQ = C.M68K_INS_FBEQ
	M68K_INS_FBOGT = C.M68K_INS_FBOGT
	M68K_INS_FBOGE = C.M68K_INS_FBOGE
	M68K_INS_FBOLT = C.M68K_INS_FBOLT
	M68K_INS_FBOLE = C.M68K_INS_FBOLE
	M68K_INS_FBOGL = C.M68K_INS_FBOGL
	M68K_INS_FBOR = C.M68K_INS_FBOR
	M68K_INS_FBUN = C.M68K_INS_FBUN
	M68K_INS_FBUEQ = C.M68K_INS_FBUEQ
	M68K_INS_FBUGT = C.M68K_INS_FBUGT
	M68K_INS_FBUGE = C.M68K_INS_FBUGE
	M68K_INS_FBULT = C.M68K_INS_FBULT
	M68K_INS_FBULE = C.M68K_INS_FBULE
	M68K_INS_FBNE = C.M68K_INS_FBNE
	M68K_INS_FBT = C.M68K_INS_FBT
	M68K_INS_FBSF = C.M68K_INS_FBSF
	M68K_INS_FBSEQ = C.M68K_INS_FBSEQ
	M68K_INS_FBGT = C.M68K_INS_FBGT
	M68K_INS_FBGE = C.M68K_INS_FBGE
	M68K_INS_FBLT = C.M68K_INS_FBLT
	M68K_INS_FBLE = C.M68K_INS_FBLE
	M68K_INS_FBGL = C.M68K_INS_FBGL
	M68K_INS_FBGLE = C.M68K_INS_FBGLE
	M68K_INS_FBNGLE = C.M68K_INS_FBNGLE
	M68K_INS_FBNGL = C.M68K_INS_FBNGL
	M68K_INS_FBNLE = C.M68K_INS_FBNLE
	M68K_INS_FBNLT = C.M68K_INS_FBNLT
	M68K_INS_FBNGE = C.M68K_INS_FBNGE
	M68K_INS_FBNGT = C.M68K_INS_FBNGT
	M68K_INS_FBSNE = C.M68K_INS_FBSNE
	M68K_INS_FBST = C.M68K_INS_FBST
	M68K_INS_FCMP = C.M68K_INS_FCMP
	M68K_INS_FCOS = C.M68K_INS_FCOS
	M68K_INS_FCOSH = C.M68K_INS_FCOSH
	M68K_INS_FDBF = C.M68K_INS_FDBF
	M68K_INS_FDBEQ = C.M68K_INS_FDBEQ
	M68K_INS_FDBOGT = C.M68K_INS_FDBOGT
	M68K_INS_FDBOGE = C.M68K_INS_FDBOGE
	M68K_INS_FDBOLT = C.M68K_INS_FDBOLT
	M68K_INS_FDBOLE = C.M68K_INS_FDBOLE
	M68K_INS_FDBOGL = C.M68K_INS_FDBOGL
	M68K_INS_FDBOR = C.M68K_INS_FDBOR
	M68K_INS_FDBUN = C.M68K_INS_FDBUN
	M68K_INS_FDBUEQ = C.M68K_INS_FDBUEQ
	M68K_INS_FDBUGT = C.M68K_INS_FDBUGT
	M68K_INS_FDBUGE = C.M68K_INS_FDBUGE
	M68K_INS_FDBULT = C.M68K_INS_FDBULT
	M68K_INS_FDBULE = C.M68K_INS_FDBULE
	M68K_INS_FDBNE = C.M68K_INS_FDBNE
	M68K_INS_FDBT = C.M68K_INS_FDBT
	M68K_INS_FDBSF = C.M68K_INS_FDBSF
	M68K_INS_FDBSEQ = C.M68K_INS_FDBSEQ
	M68K_INS_FDBGT = C.M68K_INS_FDBGT
	M68K_INS_FDBGE = C.M68K_INS_FDBGE
	M68K_INS_FDBLT = C.M68K_INS_FDBLT
	M68K_INS_FDBLE = C.M68K_INS_FDBLE
	M68K_INS_FDBGL = C.M68K_INS_FDBGL
	M68K_INS_FDBGLE = C.M68K_INS_FDBGLE
	M68K_INS_FDBNGLE = C.M68K_INS_FDBNGLE
	M68K_INS_FDBNGL = C.M68K_INS_FDBNGL
	M68K_INS_FDBNLE = C.M68K_INS_FDBNLE
	M68K_INS_FDBNLT = C.M68K_INS_FDBNLT
	M68K_INS_FDBNGE = C.M68K_INS_FDBNGE
	M68K_INS_FDBNGT = C.M68K_INS_FDBNGT
	M68K_INS_FDBSNE = C.M68K_INS_FDBSNE
	M68K_INS_FDBST = C.M68K_INS_FDBST
	M68K_INS_FDIV = C.M68K_INS_FDIV
	M68K_INS_FSDIV = C.M68K_INS_FSDIV
	M68K_INS_FDDIV = C.M68K_INS_FDDIV
	M68K_INS_FETOX = C.M68K_INS_FETOX
	M68K_INS_FETOXM1 = C.M68K_INS_FETOXM1
	M68K_INS_FGETEXP = C.M68K_INS_FGETEXP
	M68K_INS_FGETMAN = C.M68K_INS_FGETMAN
	M68K_INS_FINT = C.M68K_INS_FINT
	M68K_INS_FINTRZ = C.M68K_INS_FINTRZ
	M68K_INS_FLOG10 = C.M68K_INS_FLOG10
	M68K_INS_FLOG2 = C.M68K_INS_FLOG2
	M68K_INS_FLOGN = C.M68K_INS_FLOGN
	M68K_INS_FLOGNP1 = C.M68K_INS_FLOGNP1
	M68K_INS_FMOD = C.M68K_INS_FMOD
	M68K_INS_FMOVE = C.M68K_INS_FMOVE
	M68K_INS_FSMOVE = C.M68K_INS_FSMOVE
	M68K_INS_FDMOVE = C.M68K_INS_FDMOVE
	M68K_INS_FMOVECR = C.M68K_INS_FMOVECR
	M68K_INS_FMOVEM = C.M68K_INS_FMOVEM
	M68K_INS_FMUL = C.M68K_INS_FMUL
	M68K_INS_FSMUL = C.M68K_INS_FSMUL
	M68K_INS_FDMUL = C.M68K_INS_FDMUL
	M68K_INS_FNEG = C.M68K_INS_FNEG
	M68K_INS_FSNEG = C.M68K_INS_FSNEG
	M68K_INS_FDNEG = C.M68K_INS_FDNEG
	M68K_INS_FNOP = C.M68K_INS_FNOP
	M68K_INS_FREM = C.M68K_INS_FREM
	M68K_INS_FRESTORE = C.M68K_INS_FRESTORE
	M68K_INS_FSAVE = C.M68K_INS_FSAVE
	M68K_INS_FSCALE = C.M68K_INS_FSCALE
	M68K_INS_FSGLDIV = C.M68K_INS_FSGLDIV
	M68K_INS_FSGLMUL = C.M68K_INS_FSGLMUL
	M68K_INS_FSIN = C.M68K_INS_FSIN
	M68K_INS_FSINCOS = C.M68K_INS_FSINCOS
	M68K_INS_FSINH = C.M68K_INS_FSINH
	M68K_INS_FSQRT = C.M68K_INS_FSQRT
	M68K_INS_FSSQRT = C.M68K_INS_FSSQRT
	M68K_INS_FDSQRT = C.M68K_INS_FDSQRT
	M68K_INS_FSF = C.M68K_INS_FSF
	M68K_INS_FSBEQ = C.M68K_INS_FSBEQ
	M68K_INS_FSOGT = C.M68K_INS_FSOGT
	M68K_INS_FSOGE = C.M68K_INS_FSOGE
	M68K_INS_FSOLT = C.M68K_INS_FSOLT
	M68K_INS_FSOLE = C.M68K_INS_FSOLE
	M68K_INS_FSOGL = C.M68K_INS_FSOGL
	M68K_INS_FSOR = C.M68K_INS_FSOR
	M68K_INS_FSUN = C.M68K_INS_FSUN
	M68K_INS_FSUEQ = C.M68K_INS_FSUEQ
	M68K_INS_FSUGT = C.M68K_INS_FSUGT
	M68K_INS_FSUGE = C.M68K_INS_FSUGE
	M68K_INS_FSULT = C.M68K_INS_FSULT
	M68K_INS_FSULE = C.M68K_INS_FSULE
	M68K_INS_FSNE = C.M68K_INS_FSNE
	M68K_INS_FST = C.M68K_INS_FST
	M68K_INS_FSSF = C.M68K_INS_FSSF
	M68K_INS_FSSEQ = C.M68K_INS_FSSEQ
	M68K_INS_FSGT = C.M68K_INS_FSGT
	M68K_INS_FSGE = C.M68K_INS_FSGE
	M68K_INS_FSLT = C.M68K_INS_FSLT
	M68K_INS_FSLE = C.M68K_INS_FSLE
	M68K_INS_FSGL = C.M68K_INS_FSGL
	M68K_INS_FSGLE = C.M68K_INS_FSGLE
	M68K_INS_FSNGLE = C.M68K_INS_FSNGLE
	M68K_INS_FSNGL = C.M68K_INS_FSNGL
	M68K_INS_FSNLE = C.M68K_INS_FSNLE
	M68K_INS_FSNLT = C.M68K_INS_FSNLT
	M68K_INS_FSNGE = C.M68K_INS_FSNGE
	M68K_INS_FSNGT = C.M68K_INS_FSNGT
	M68K_INS_FSSNE = C.M68K_INS_FSSNE
	M68K_INS_FSST = C.M68K_INS_FSST
	M68K_INS_FSUB = C.M68K_INS_FSUB
	M68K_INS_FSSUB = C.M68K_INS_FSSUB
	M68K_INS_FDSUB = C.M68K_INS_FDSUB
	M68K_INS_FTAN = C.M68K_INS_FTAN
	M68K_INS_FTANH = C.M68K_INS_FTANH
	M68K_INS_FTENTOX = C.M68K_INS_FTENTOX
	M68K_INS_FTRAPF = C.M68K_INS_FTRAPF
	M68K_INS_FTRAPEQ = C.M68K_INS_FTRAPEQ
	M68K_INS_FTRAPOGT = C.M68K_INS_FTRAPOGT
	M68K_INS_FTRAPOGE = C.M68K_INS_FTRAPOGE
	M68K_INS_FTRAPOLT = C.M68K_INS_FTRAPOLT
	M68K_INS_FTRAPOLE = C.M68K_INS_FTRAPOLE
	M68K_INS_FTRAPOGL = C.M68K_INS_FTRAPOGL
	M68K_INS_FTRAPOR = C.M68K_INS_FTRAPOR
	M68K_INS_FTRAPUN = C.M68K_INS_FTRAPUN
	M68K_INS_FTRAPUEQ = C.M68K_INS_FTRAPUEQ
	M68K_INS_FTRAPUGT = C.M68K_INS_FTRAPUGT
	M68K_INS_FTRAPUGE = C.M68K_INS_FTRAPUGE
	M68K_INS_FTRAPULT = C.M68K_INS_FTRAPULT
	M68K_INS_FTRAPULE = C.M68K_INS_FTRAPULE
	M68K_INS_FTRAPNE = C.M68K_INS_FTRAPNE
	M68K_INS_FTRAPT = C.M68K_INS_FTRAPT
	M68K_INS_FTRAPSF = C.M68K_INS_FTRAPSF
	M68K_INS_FTRAPSEQ = C.M68K_INS_FTRAPSEQ
	M68K_INS_FTRAPGT = C.M68K_INS_FTRAPGT
	M68K_INS_FTRAPGE = C.M68K_INS_FTRAPGE
	M68K_INS_FTRAPLT = C.M68K_INS_FTRAPLT
	M68K_INS_FTRAPLE = C.M68K_INS_FTRAPLE
	M68K_INS_FTRAPGL = C.M68K_INS_FTRAPGL
	M68K_INS_FTRAPGLE = C.M68K_INS_FTRAPGLE
	M68K_INS_FTRAPNGLE = C.M68K_INS_FTRAPNGLE
	M68K_INS_FTRAPNGL = C.M68K_INS_FTRAPNGL
	M68K_INS_FTRAPNLE = C.M68K_INS_FTRAPNLE
	M68K_INS_FTRAPNLT = C.M68K_INS_FTRAPNLT
	M68K_INS_FTRAPNGE = C.M68K_INS_FTRAPNGE
	M68K_INS_FTRAPNGT = C.M68K_INS_FTRAPNGT
	M68K_INS_FTRAPSNE = C.M68K_INS_FTRAPSNE
	M68K_INS_FTRAPST = C.M68K_INS_FTRAPST
	M68K_INS_FTST = C.M68K_INS_FTST
	M68K_INS_FTWOTOX = C.M68K_INS_FTWOTOX
	M68K_INS_HALT = C.M68K_INS_HALT
	M68K_INS_ILLEGAL = C.M68K_INS_ILLEGAL
	M68K_INS_JMP = C.M68K_INS_JMP
	M68K_INS_JSR = C.M68K_INS_JSR
	M68K_INS_LEA = C.M68K_INS_LEA
	M68K_INS_LINK = C.M68K_INS_LINK
	M68K_INS_LPSTOP = C.M68K_INS_LPSTOP
	M68K_INS_LSL = C.M68K_INS_LSL
	M68K_INS_LSR = C.M68K_INS_LSR
	M68K_INS_MOVE = C.M68K_INS_MOVE
	M68K_INS_MOVEA = C.M68K_INS_MOVEA
	M68K_INS_MOVEC = C.M68K_INS_MOVEC
	M68K_INS_MOVEM = C.M68K_INS_MOVEM
	M68K_INS_MOVEP = C.M68K_INS_MOVEP
	M68K_INS_MOVEQ = C.M68K_INS_MOVEQ
	M68K_INS_MOVES = C.M68K_INS_MOVES
	M68K_INS_MOVE16 = C.M68K_INS_MOVE16
	M68K_INS_MULS = C.M68K_INS_MULS
	M68K_INS_MULU = C.M68K_INS_MULU
	M68K_INS_NBCD = C.M68K_INS_NBCD
	M68K_INS_NEG = C.M68K_INS_NEG
	M68K_INS_NEGX = C.M68K_INS_NEGX
	M68K_INS_NOP = C.M68K_INS_NOP
	M68K_INS_NOT = C.M68K_INS_NOT
	M68K_INS_OR = C.M68K_INS_OR
	M68K_INS_ORI = C.M68K_INS_ORI
	M68K_INS_PACK = C.M68K_INS_PACK
	M68K_INS_PEA = C.M68K_INS_PEA
	M68K_INS_PFLUSH = C.M68K_INS_PFLUSH
	M68K_INS_PFLUSHA = C.M68K_INS_PFLUSHA
	M68K_INS_PFLUSHAN = C.M68K_INS_PFLUSHAN
	M68K_INS_PFLUSHN = C.M68K_INS_PFLUSHN
	M68K_INS_PLOADR = C.M68K_INS_PLOADR
	M68K_INS_PLOADW = C.M68K_INS_PLOADW
	M68K_INS_PLPAR = C.M68K_INS_PLPAR
	M68K_INS_PLPAW = C.M68K_INS_PLPAW
	M68K_INS_PMOVE = C.M68K_INS_PMOVE
	M68K_INS_PMOVEFD = C.M68K_INS_PMOVEFD
	M68K_INS_PTESTR = C.M68K_INS_PTESTR
	M68K_INS_PTESTW = C.M68K_INS_PTESTW
	M68K_INS_PULSE = C.M68K_INS_PULSE
	M68K_INS_REMS = C.M68K_INS_REMS
	M68K_INS_REMU = C.M68K_INS_REMU
	M68K_INS_RESET = C.M68K_INS_RESET
	M68K_INS_ROL = C.M68K_INS_ROL
	M68K_INS_ROR = C.M68K_INS_ROR
	M68K_INS_ROXL = C.M68K_INS_ROXL
	M68K_INS_ROXR = C.M68K_INS_ROXR
	M68K_INS_RTD = C.M68K_INS_RTD
	M68K_INS_RTE = C.M68K_INS_RTE
	M68K_INS_RTM = C.M68K_INS_RTM
	M68K_INS_RTR = C.M68K_INS_RTR
	M68K_INS_RTS = C.M68K_INS_RTS
	M68K_INS_SBCD = C.M68K_INS_SBCD
	M68K_INS_ST = C.M68K_INS_ST
	M68K_INS_SF = C.M68K_INS_SF
	M68K_INS_SHI = C.M68K_INS_SHI
	M68K_INS_SLS = C.M68K_INS_SLS
	M68K_INS_SCC = C.M68K_INS_SCC
	M68K_INS_SHS = C.M68K_INS_SHS
	M68K_INS_SCS = C.M68K_INS_SCS
	M68K_INS_SLO = C.M68K_INS_SLO
	M68K_INS_SNE = C.M68K_INS_SNE
	M68K_INS_SEQ = C.M68K_INS_SEQ
	M68K_INS_SVC = C.M68K_INS_SVC
	M68K_INS_SVS = C.M68K_INS_SVS
	M68K_INS_SPL = C.M68K_INS_SPL
	M68K_INS_SMI = C.M68K_INS_SMI
	M68K_INS_SGE = C.M68K_INS_SGE
	M68K_INS_SLT = C.M68K_INS_SLT
	M68K_INS_SGT = C.M68K_INS_SGT
	M68K_INS_SLE = C.M68K_INS_SLE
	M68K_INS_STOP = C.M68K_INS_STOP
	M68K_INS_SUB = C.M68K_INS_SUB
	M68K_INS_SUBA = C.M68K_INS_SUBA
	M68K_INS_SUBI = C.M68K_INS_SUBI
	M68K_INS_SUBQ = C.M68K_INS_SUBQ
	M68K_INS_SUBX = C.M68K_INS_SUBX
	M68K_INS_SWAP = C.M68K_INS_SWAP
	M68K_INS_TAS = C.M68K_INS_TAS
	M68K_INS_TRAP = C.M68K_INS_TRAP
	M68K_INS_TRAPV = C.M68K_INS_TRAPV
	M68K_INS_TRAPT = C.M68K_INS_TRAPT
	M68K_INS_TRAPF = C.M68K_INS_TRAPF
	M68K_INS_TRAPHI = C.M68K_INS_TRAPHI
	M68K_INS_TRAPLS = C.M68K_INS_TRAPLS
	M68K_INS_TRAPCC = C.M68K_INS_TRAPCC
	M68K_INS_TRAPHS = C.M68K_INS_TRAPHS
	M68K_INS_TRAPCS = C.M68K_INS_TRAPCS
	M68K_INS_TRAPLO = C.M68K_INS_TRAPLO
	M68K_INS_TRAPNE = C.M68K_INS_TRAPNE
	M68K_INS_TRAPEQ = C.M68K_INS_TRAPEQ
	M68K_INS_TRAPVC = C.M68K_INS_TRAPVC
	M68K_INS_TRAPVS = C.M68K_INS_TRAPVS
	M68K_INS_TRAPPL = C.M68K_INS_TRAPPL
	M68K_INS_TRAPMI = C.M68K_INS_TRAPMI
	M68K_INS_TRAPGE = C.M68K_INS_TRAPGE
	M68K_INS_TRAPLT = C.M68K_INS_TRAPLT
	M68K_INS_TRAPGT = C.M68K_INS_TRAPGT
	M68K_INS_TRAPLE = C.M68K_INS_TRAPLE
	M68K_INS_TST = C.M68K_INS_TST
	M68K_INS_UNLK = C.M68K_INS_UNLK
	M68K_INS_UNPK = C.M68K_INS_UNPK
	M68K_INS_ENDING = C.M68K_INS_ENDING
)

const (
	M68K_GRP_INVALID = C.M68K_GRP_INVALID
	M68K_GRP_JUMP = C.M68K_GRP_JUMP
	M68K_GRP_RET = C.M68K_GRP_RET
	M68K_GRP_IRET = C.M68K_GRP_IRET
	M68K_GRP_BRANCH_RELATIVE = C.M68K_GRP_BRANCH_RELATIVE
	M68K_GRP_ENDING = C.M68K_GRP_ENDING
)

