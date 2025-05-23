/*
Gapstone is a Go binding for the Capstone disassembly library.
For examples, try reading the *_test.go files.

	Library Author: Nguyen Anh Quynh
	Binding Author: Ben Nagy
	License: BSD style - see LICENSE file for details
    (c) 2013 COSEINC. All Rights Reserved.

    THIS FILE WAS AUTO-GENERATED -- DO NOT EDIT!
	Command: ./genconst.rb capstone/bindings/python/capstone/
	Created at: 2025-04-25T09:04:27+00:00

*/

package gapstone

// #cgo LDFLAGS: -lcapstone
// #cgo freebsd CFLAGS: -I/usr/local/include
// #cgo freebsd LDFLAGS: -L/usr/local/lib
// #include <stdlib.h>
// #include <capstone/capstone.h>
import "C"

// For Capstone Engine. AUTO-GENERATED FILE, DO NOT EDIT [sparc_const.py]
const (
	SPARC_CC_INVALID = C.SPARC_CC_INVALID
	SPARC_CC_ICC_A   = C.SPARC_CC_ICC_A
	SPARC_CC_ICC_N   = C.SPARC_CC_ICC_N
	SPARC_CC_ICC_NE  = C.SPARC_CC_ICC_NE
	SPARC_CC_ICC_E   = C.SPARC_CC_ICC_E
	SPARC_CC_ICC_G   = C.SPARC_CC_ICC_G
	SPARC_CC_ICC_LE  = C.SPARC_CC_ICC_LE
	SPARC_CC_ICC_GE  = C.SPARC_CC_ICC_GE
	SPARC_CC_ICC_L   = C.SPARC_CC_ICC_L
	SPARC_CC_ICC_GU  = C.SPARC_CC_ICC_GU
	SPARC_CC_ICC_LEU = C.SPARC_CC_ICC_LEU
	SPARC_CC_ICC_CC  = C.SPARC_CC_ICC_CC
	SPARC_CC_ICC_CS  = C.SPARC_CC_ICC_CS
	SPARC_CC_ICC_POS = C.SPARC_CC_ICC_POS
	SPARC_CC_ICC_NEG = C.SPARC_CC_ICC_NEG
	SPARC_CC_ICC_VC  = C.SPARC_CC_ICC_VC
	SPARC_CC_ICC_VS  = C.SPARC_CC_ICC_VS
	SPARC_CC_FCC_A   = C.SPARC_CC_FCC_A
	SPARC_CC_FCC_N   = C.SPARC_CC_FCC_N
	SPARC_CC_FCC_U   = C.SPARC_CC_FCC_U
	SPARC_CC_FCC_G   = C.SPARC_CC_FCC_G
	SPARC_CC_FCC_UG  = C.SPARC_CC_FCC_UG
	SPARC_CC_FCC_L   = C.SPARC_CC_FCC_L
	SPARC_CC_FCC_UL  = C.SPARC_CC_FCC_UL
	SPARC_CC_FCC_LG  = C.SPARC_CC_FCC_LG
	SPARC_CC_FCC_NE  = C.SPARC_CC_FCC_NE
	SPARC_CC_FCC_E   = C.SPARC_CC_FCC_E
	SPARC_CC_FCC_UE  = C.SPARC_CC_FCC_UE
	SPARC_CC_FCC_GE  = C.SPARC_CC_FCC_GE
	SPARC_CC_FCC_UGE = C.SPARC_CC_FCC_UGE
	SPARC_CC_FCC_LE  = C.SPARC_CC_FCC_LE
	SPARC_CC_FCC_ULE = C.SPARC_CC_FCC_ULE
	SPARC_CC_FCC_O   = C.SPARC_CC_FCC_O
)

const (
	SPARC_HINT_INVALID = C.SPARC_HINT_INVALID
	SPARC_HINT_A       = C.SPARC_HINT_A
	SPARC_HINT_PT      = C.SPARC_HINT_PT
	SPARC_HINT_PN      = C.SPARC_HINT_PN
)

const (
	SPARC_OP_INVALID = C.SPARC_OP_INVALID
	SPARC_OP_REG     = C.SPARC_OP_REG
	SPARC_OP_IMM     = C.SPARC_OP_IMM
	SPARC_OP_MEM     = C.SPARC_OP_MEM
)

const (
	SPARC_REG_INVALID = C.SPARC_REG_INVALID
	SPARC_REG_F0      = C.SPARC_REG_F0
	SPARC_REG_F1      = C.SPARC_REG_F1
	SPARC_REG_F2      = C.SPARC_REG_F2
	SPARC_REG_F3      = C.SPARC_REG_F3
	SPARC_REG_F4      = C.SPARC_REG_F4
	SPARC_REG_F5      = C.SPARC_REG_F5
	SPARC_REG_F6      = C.SPARC_REG_F6
	SPARC_REG_F7      = C.SPARC_REG_F7
	SPARC_REG_F8      = C.SPARC_REG_F8
	SPARC_REG_F9      = C.SPARC_REG_F9
	SPARC_REG_F10     = C.SPARC_REG_F10
	SPARC_REG_F11     = C.SPARC_REG_F11
	SPARC_REG_F12     = C.SPARC_REG_F12
	SPARC_REG_F13     = C.SPARC_REG_F13
	SPARC_REG_F14     = C.SPARC_REG_F14
	SPARC_REG_F15     = C.SPARC_REG_F15
	SPARC_REG_F16     = C.SPARC_REG_F16
	SPARC_REG_F17     = C.SPARC_REG_F17
	SPARC_REG_F18     = C.SPARC_REG_F18
	SPARC_REG_F19     = C.SPARC_REG_F19
	SPARC_REG_F20     = C.SPARC_REG_F20
	SPARC_REG_F21     = C.SPARC_REG_F21
	SPARC_REG_F22     = C.SPARC_REG_F22
	SPARC_REG_F23     = C.SPARC_REG_F23
	SPARC_REG_F24     = C.SPARC_REG_F24
	SPARC_REG_F25     = C.SPARC_REG_F25
	SPARC_REG_F26     = C.SPARC_REG_F26
	SPARC_REG_F27     = C.SPARC_REG_F27
	SPARC_REG_F28     = C.SPARC_REG_F28
	SPARC_REG_F29     = C.SPARC_REG_F29
	SPARC_REG_F30     = C.SPARC_REG_F30
	SPARC_REG_F31     = C.SPARC_REG_F31
	SPARC_REG_F32     = C.SPARC_REG_F32
	SPARC_REG_F34     = C.SPARC_REG_F34
	SPARC_REG_F36     = C.SPARC_REG_F36
	SPARC_REG_F38     = C.SPARC_REG_F38
	SPARC_REG_F40     = C.SPARC_REG_F40
	SPARC_REG_F42     = C.SPARC_REG_F42
	SPARC_REG_F44     = C.SPARC_REG_F44
	SPARC_REG_F46     = C.SPARC_REG_F46
	SPARC_REG_F48     = C.SPARC_REG_F48
	SPARC_REG_F50     = C.SPARC_REG_F50
	SPARC_REG_F52     = C.SPARC_REG_F52
	SPARC_REG_F54     = C.SPARC_REG_F54
	SPARC_REG_F56     = C.SPARC_REG_F56
	SPARC_REG_F58     = C.SPARC_REG_F58
	SPARC_REG_F60     = C.SPARC_REG_F60
	SPARC_REG_F62     = C.SPARC_REG_F62
	SPARC_REG_FCC0    = C.SPARC_REG_FCC0
	SPARC_REG_FCC1    = C.SPARC_REG_FCC1
	SPARC_REG_FCC2    = C.SPARC_REG_FCC2
	SPARC_REG_FCC3    = C.SPARC_REG_FCC3
	SPARC_REG_FP      = C.SPARC_REG_FP
	SPARC_REG_G0      = C.SPARC_REG_G0
	SPARC_REG_G1      = C.SPARC_REG_G1
	SPARC_REG_G2      = C.SPARC_REG_G2
	SPARC_REG_G3      = C.SPARC_REG_G3
	SPARC_REG_G4      = C.SPARC_REG_G4
	SPARC_REG_G5      = C.SPARC_REG_G5
	SPARC_REG_G6      = C.SPARC_REG_G6
	SPARC_REG_G7      = C.SPARC_REG_G7
	SPARC_REG_I0      = C.SPARC_REG_I0
	SPARC_REG_I1      = C.SPARC_REG_I1
	SPARC_REG_I2      = C.SPARC_REG_I2
	SPARC_REG_I3      = C.SPARC_REG_I3
	SPARC_REG_I4      = C.SPARC_REG_I4
	SPARC_REG_I5      = C.SPARC_REG_I5
	SPARC_REG_I7      = C.SPARC_REG_I7
	SPARC_REG_ICC     = C.SPARC_REG_ICC
	SPARC_REG_L0      = C.SPARC_REG_L0
	SPARC_REG_L1      = C.SPARC_REG_L1
	SPARC_REG_L2      = C.SPARC_REG_L2
	SPARC_REG_L3      = C.SPARC_REG_L3
	SPARC_REG_L4      = C.SPARC_REG_L4
	SPARC_REG_L5      = C.SPARC_REG_L5
	SPARC_REG_L6      = C.SPARC_REG_L6
	SPARC_REG_L7      = C.SPARC_REG_L7
	SPARC_REG_O0      = C.SPARC_REG_O0
	SPARC_REG_O1      = C.SPARC_REG_O1
	SPARC_REG_O2      = C.SPARC_REG_O2
	SPARC_REG_O3      = C.SPARC_REG_O3
	SPARC_REG_O4      = C.SPARC_REG_O4
	SPARC_REG_O5      = C.SPARC_REG_O5
	SPARC_REG_O7      = C.SPARC_REG_O7
	SPARC_REG_SP      = C.SPARC_REG_SP
	SPARC_REG_Y       = C.SPARC_REG_Y
	SPARC_REG_XCC     = C.SPARC_REG_XCC
	SPARC_REG_ENDING  = C.SPARC_REG_ENDING
	SPARC_REG_O6      = C.SPARC_REG_O6
	SPARC_REG_I6      = C.SPARC_REG_I6
)

const (
	SPARC_INS_INVALID     = C.SPARC_INS_INVALID
	SPARC_INS_ADDCC       = C.SPARC_INS_ADDCC
	SPARC_INS_ADDX        = C.SPARC_INS_ADDX
	SPARC_INS_ADDXCC      = C.SPARC_INS_ADDXCC
	SPARC_INS_ADDXC       = C.SPARC_INS_ADDXC
	SPARC_INS_ADDXCCC     = C.SPARC_INS_ADDXCCC
	SPARC_INS_ADD         = C.SPARC_INS_ADD
	SPARC_INS_ALIGNADDR   = C.SPARC_INS_ALIGNADDR
	SPARC_INS_ALIGNADDRL  = C.SPARC_INS_ALIGNADDRL
	SPARC_INS_ANDCC       = C.SPARC_INS_ANDCC
	SPARC_INS_ANDNCC      = C.SPARC_INS_ANDNCC
	SPARC_INS_ANDN        = C.SPARC_INS_ANDN
	SPARC_INS_AND         = C.SPARC_INS_AND
	SPARC_INS_ARRAY16     = C.SPARC_INS_ARRAY16
	SPARC_INS_ARRAY32     = C.SPARC_INS_ARRAY32
	SPARC_INS_ARRAY8      = C.SPARC_INS_ARRAY8
	SPARC_INS_B           = C.SPARC_INS_B
	SPARC_INS_JMP         = C.SPARC_INS_JMP
	SPARC_INS_BMASK       = C.SPARC_INS_BMASK
	SPARC_INS_FB          = C.SPARC_INS_FB
	SPARC_INS_BRGEZ       = C.SPARC_INS_BRGEZ
	SPARC_INS_BRGZ        = C.SPARC_INS_BRGZ
	SPARC_INS_BRLEZ       = C.SPARC_INS_BRLEZ
	SPARC_INS_BRLZ        = C.SPARC_INS_BRLZ
	SPARC_INS_BRNZ        = C.SPARC_INS_BRNZ
	SPARC_INS_BRZ         = C.SPARC_INS_BRZ
	SPARC_INS_BSHUFFLE    = C.SPARC_INS_BSHUFFLE
	SPARC_INS_CALL        = C.SPARC_INS_CALL
	SPARC_INS_CASX        = C.SPARC_INS_CASX
	SPARC_INS_CAS         = C.SPARC_INS_CAS
	SPARC_INS_CMASK16     = C.SPARC_INS_CMASK16
	SPARC_INS_CMASK32     = C.SPARC_INS_CMASK32
	SPARC_INS_CMASK8      = C.SPARC_INS_CMASK8
	SPARC_INS_CMP         = C.SPARC_INS_CMP
	SPARC_INS_EDGE16      = C.SPARC_INS_EDGE16
	SPARC_INS_EDGE16L     = C.SPARC_INS_EDGE16L
	SPARC_INS_EDGE16LN    = C.SPARC_INS_EDGE16LN
	SPARC_INS_EDGE16N     = C.SPARC_INS_EDGE16N
	SPARC_INS_EDGE32      = C.SPARC_INS_EDGE32
	SPARC_INS_EDGE32L     = C.SPARC_INS_EDGE32L
	SPARC_INS_EDGE32LN    = C.SPARC_INS_EDGE32LN
	SPARC_INS_EDGE32N     = C.SPARC_INS_EDGE32N
	SPARC_INS_EDGE8       = C.SPARC_INS_EDGE8
	SPARC_INS_EDGE8L      = C.SPARC_INS_EDGE8L
	SPARC_INS_EDGE8LN     = C.SPARC_INS_EDGE8LN
	SPARC_INS_EDGE8N      = C.SPARC_INS_EDGE8N
	SPARC_INS_FABSD       = C.SPARC_INS_FABSD
	SPARC_INS_FABSQ       = C.SPARC_INS_FABSQ
	SPARC_INS_FABSS       = C.SPARC_INS_FABSS
	SPARC_INS_FADDD       = C.SPARC_INS_FADDD
	SPARC_INS_FADDQ       = C.SPARC_INS_FADDQ
	SPARC_INS_FADDS       = C.SPARC_INS_FADDS
	SPARC_INS_FALIGNDATA  = C.SPARC_INS_FALIGNDATA
	SPARC_INS_FAND        = C.SPARC_INS_FAND
	SPARC_INS_FANDNOT1    = C.SPARC_INS_FANDNOT1
	SPARC_INS_FANDNOT1S   = C.SPARC_INS_FANDNOT1S
	SPARC_INS_FANDNOT2    = C.SPARC_INS_FANDNOT2
	SPARC_INS_FANDNOT2S   = C.SPARC_INS_FANDNOT2S
	SPARC_INS_FANDS       = C.SPARC_INS_FANDS
	SPARC_INS_FCHKSM16    = C.SPARC_INS_FCHKSM16
	SPARC_INS_FCMPD       = C.SPARC_INS_FCMPD
	SPARC_INS_FCMPEQ16    = C.SPARC_INS_FCMPEQ16
	SPARC_INS_FCMPEQ32    = C.SPARC_INS_FCMPEQ32
	SPARC_INS_FCMPGT16    = C.SPARC_INS_FCMPGT16
	SPARC_INS_FCMPGT32    = C.SPARC_INS_FCMPGT32
	SPARC_INS_FCMPLE16    = C.SPARC_INS_FCMPLE16
	SPARC_INS_FCMPLE32    = C.SPARC_INS_FCMPLE32
	SPARC_INS_FCMPNE16    = C.SPARC_INS_FCMPNE16
	SPARC_INS_FCMPNE32    = C.SPARC_INS_FCMPNE32
	SPARC_INS_FCMPQ       = C.SPARC_INS_FCMPQ
	SPARC_INS_FCMPS       = C.SPARC_INS_FCMPS
	SPARC_INS_FDIVD       = C.SPARC_INS_FDIVD
	SPARC_INS_FDIVQ       = C.SPARC_INS_FDIVQ
	SPARC_INS_FDIVS       = C.SPARC_INS_FDIVS
	SPARC_INS_FDMULQ      = C.SPARC_INS_FDMULQ
	SPARC_INS_FDTOI       = C.SPARC_INS_FDTOI
	SPARC_INS_FDTOQ       = C.SPARC_INS_FDTOQ
	SPARC_INS_FDTOS       = C.SPARC_INS_FDTOS
	SPARC_INS_FDTOX       = C.SPARC_INS_FDTOX
	SPARC_INS_FEXPAND     = C.SPARC_INS_FEXPAND
	SPARC_INS_FHADDD      = C.SPARC_INS_FHADDD
	SPARC_INS_FHADDS      = C.SPARC_INS_FHADDS
	SPARC_INS_FHSUBD      = C.SPARC_INS_FHSUBD
	SPARC_INS_FHSUBS      = C.SPARC_INS_FHSUBS
	SPARC_INS_FITOD       = C.SPARC_INS_FITOD
	SPARC_INS_FITOQ       = C.SPARC_INS_FITOQ
	SPARC_INS_FITOS       = C.SPARC_INS_FITOS
	SPARC_INS_FLCMPD      = C.SPARC_INS_FLCMPD
	SPARC_INS_FLCMPS      = C.SPARC_INS_FLCMPS
	SPARC_INS_FLUSHW      = C.SPARC_INS_FLUSHW
	SPARC_INS_FMEAN16     = C.SPARC_INS_FMEAN16
	SPARC_INS_FMOVD       = C.SPARC_INS_FMOVD
	SPARC_INS_FMOVQ       = C.SPARC_INS_FMOVQ
	SPARC_INS_FMOVRDGEZ   = C.SPARC_INS_FMOVRDGEZ
	SPARC_INS_FMOVRQGEZ   = C.SPARC_INS_FMOVRQGEZ
	SPARC_INS_FMOVRSGEZ   = C.SPARC_INS_FMOVRSGEZ
	SPARC_INS_FMOVRDGZ    = C.SPARC_INS_FMOVRDGZ
	SPARC_INS_FMOVRQGZ    = C.SPARC_INS_FMOVRQGZ
	SPARC_INS_FMOVRSGZ    = C.SPARC_INS_FMOVRSGZ
	SPARC_INS_FMOVRDLEZ   = C.SPARC_INS_FMOVRDLEZ
	SPARC_INS_FMOVRQLEZ   = C.SPARC_INS_FMOVRQLEZ
	SPARC_INS_FMOVRSLEZ   = C.SPARC_INS_FMOVRSLEZ
	SPARC_INS_FMOVRDLZ    = C.SPARC_INS_FMOVRDLZ
	SPARC_INS_FMOVRQLZ    = C.SPARC_INS_FMOVRQLZ
	SPARC_INS_FMOVRSLZ    = C.SPARC_INS_FMOVRSLZ
	SPARC_INS_FMOVRDNZ    = C.SPARC_INS_FMOVRDNZ
	SPARC_INS_FMOVRQNZ    = C.SPARC_INS_FMOVRQNZ
	SPARC_INS_FMOVRSNZ    = C.SPARC_INS_FMOVRSNZ
	SPARC_INS_FMOVRDZ     = C.SPARC_INS_FMOVRDZ
	SPARC_INS_FMOVRQZ     = C.SPARC_INS_FMOVRQZ
	SPARC_INS_FMOVRSZ     = C.SPARC_INS_FMOVRSZ
	SPARC_INS_FMOVS       = C.SPARC_INS_FMOVS
	SPARC_INS_FMUL8SUX16  = C.SPARC_INS_FMUL8SUX16
	SPARC_INS_FMUL8ULX16  = C.SPARC_INS_FMUL8ULX16
	SPARC_INS_FMUL8X16    = C.SPARC_INS_FMUL8X16
	SPARC_INS_FMUL8X16AL  = C.SPARC_INS_FMUL8X16AL
	SPARC_INS_FMUL8X16AU  = C.SPARC_INS_FMUL8X16AU
	SPARC_INS_FMULD       = C.SPARC_INS_FMULD
	SPARC_INS_FMULD8SUX16 = C.SPARC_INS_FMULD8SUX16
	SPARC_INS_FMULD8ULX16 = C.SPARC_INS_FMULD8ULX16
	SPARC_INS_FMULQ       = C.SPARC_INS_FMULQ
	SPARC_INS_FMULS       = C.SPARC_INS_FMULS
	SPARC_INS_FNADDD      = C.SPARC_INS_FNADDD
	SPARC_INS_FNADDS      = C.SPARC_INS_FNADDS
	SPARC_INS_FNAND       = C.SPARC_INS_FNAND
	SPARC_INS_FNANDS      = C.SPARC_INS_FNANDS
	SPARC_INS_FNEGD       = C.SPARC_INS_FNEGD
	SPARC_INS_FNEGQ       = C.SPARC_INS_FNEGQ
	SPARC_INS_FNEGS       = C.SPARC_INS_FNEGS
	SPARC_INS_FNHADDD     = C.SPARC_INS_FNHADDD
	SPARC_INS_FNHADDS     = C.SPARC_INS_FNHADDS
	SPARC_INS_FNOR        = C.SPARC_INS_FNOR
	SPARC_INS_FNORS       = C.SPARC_INS_FNORS
	SPARC_INS_FNOT1       = C.SPARC_INS_FNOT1
	SPARC_INS_FNOT1S      = C.SPARC_INS_FNOT1S
	SPARC_INS_FNOT2       = C.SPARC_INS_FNOT2
	SPARC_INS_FNOT2S      = C.SPARC_INS_FNOT2S
	SPARC_INS_FONE        = C.SPARC_INS_FONE
	SPARC_INS_FONES       = C.SPARC_INS_FONES
	SPARC_INS_FOR         = C.SPARC_INS_FOR
	SPARC_INS_FORNOT1     = C.SPARC_INS_FORNOT1
	SPARC_INS_FORNOT1S    = C.SPARC_INS_FORNOT1S
	SPARC_INS_FORNOT2     = C.SPARC_INS_FORNOT2
	SPARC_INS_FORNOT2S    = C.SPARC_INS_FORNOT2S
	SPARC_INS_FORS        = C.SPARC_INS_FORS
	SPARC_INS_FPACK16     = C.SPARC_INS_FPACK16
	SPARC_INS_FPACK32     = C.SPARC_INS_FPACK32
	SPARC_INS_FPACKFIX    = C.SPARC_INS_FPACKFIX
	SPARC_INS_FPADD16     = C.SPARC_INS_FPADD16
	SPARC_INS_FPADD16S    = C.SPARC_INS_FPADD16S
	SPARC_INS_FPADD32     = C.SPARC_INS_FPADD32
	SPARC_INS_FPADD32S    = C.SPARC_INS_FPADD32S
	SPARC_INS_FPADD64     = C.SPARC_INS_FPADD64
	SPARC_INS_FPMERGE     = C.SPARC_INS_FPMERGE
	SPARC_INS_FPSUB16     = C.SPARC_INS_FPSUB16
	SPARC_INS_FPSUB16S    = C.SPARC_INS_FPSUB16S
	SPARC_INS_FPSUB32     = C.SPARC_INS_FPSUB32
	SPARC_INS_FPSUB32S    = C.SPARC_INS_FPSUB32S
	SPARC_INS_FQTOD       = C.SPARC_INS_FQTOD
	SPARC_INS_FQTOI       = C.SPARC_INS_FQTOI
	SPARC_INS_FQTOS       = C.SPARC_INS_FQTOS
	SPARC_INS_FQTOX       = C.SPARC_INS_FQTOX
	SPARC_INS_FSLAS16     = C.SPARC_INS_FSLAS16
	SPARC_INS_FSLAS32     = C.SPARC_INS_FSLAS32
	SPARC_INS_FSLL16      = C.SPARC_INS_FSLL16
	SPARC_INS_FSLL32      = C.SPARC_INS_FSLL32
	SPARC_INS_FSMULD      = C.SPARC_INS_FSMULD
	SPARC_INS_FSQRTD      = C.SPARC_INS_FSQRTD
	SPARC_INS_FSQRTQ      = C.SPARC_INS_FSQRTQ
	SPARC_INS_FSQRTS      = C.SPARC_INS_FSQRTS
	SPARC_INS_FSRA16      = C.SPARC_INS_FSRA16
	SPARC_INS_FSRA32      = C.SPARC_INS_FSRA32
	SPARC_INS_FSRC1       = C.SPARC_INS_FSRC1
	SPARC_INS_FSRC1S      = C.SPARC_INS_FSRC1S
	SPARC_INS_FSRC2       = C.SPARC_INS_FSRC2
	SPARC_INS_FSRC2S      = C.SPARC_INS_FSRC2S
	SPARC_INS_FSRL16      = C.SPARC_INS_FSRL16
	SPARC_INS_FSRL32      = C.SPARC_INS_FSRL32
	SPARC_INS_FSTOD       = C.SPARC_INS_FSTOD
	SPARC_INS_FSTOI       = C.SPARC_INS_FSTOI
	SPARC_INS_FSTOQ       = C.SPARC_INS_FSTOQ
	SPARC_INS_FSTOX       = C.SPARC_INS_FSTOX
	SPARC_INS_FSUBD       = C.SPARC_INS_FSUBD
	SPARC_INS_FSUBQ       = C.SPARC_INS_FSUBQ
	SPARC_INS_FSUBS       = C.SPARC_INS_FSUBS
	SPARC_INS_FXNOR       = C.SPARC_INS_FXNOR
	SPARC_INS_FXNORS      = C.SPARC_INS_FXNORS
	SPARC_INS_FXOR        = C.SPARC_INS_FXOR
	SPARC_INS_FXORS       = C.SPARC_INS_FXORS
	SPARC_INS_FXTOD       = C.SPARC_INS_FXTOD
	SPARC_INS_FXTOQ       = C.SPARC_INS_FXTOQ
	SPARC_INS_FXTOS       = C.SPARC_INS_FXTOS
	SPARC_INS_FZERO       = C.SPARC_INS_FZERO
	SPARC_INS_FZEROS      = C.SPARC_INS_FZEROS
	SPARC_INS_JMPL        = C.SPARC_INS_JMPL
	SPARC_INS_LDD         = C.SPARC_INS_LDD
	SPARC_INS_LD          = C.SPARC_INS_LD
	SPARC_INS_LDQ         = C.SPARC_INS_LDQ
	SPARC_INS_LDSB        = C.SPARC_INS_LDSB
	SPARC_INS_LDSH        = C.SPARC_INS_LDSH
	SPARC_INS_LDSW        = C.SPARC_INS_LDSW
	SPARC_INS_LDUB        = C.SPARC_INS_LDUB
	SPARC_INS_LDUH        = C.SPARC_INS_LDUH
	SPARC_INS_LDX         = C.SPARC_INS_LDX
	SPARC_INS_LZCNT       = C.SPARC_INS_LZCNT
	SPARC_INS_MEMBAR      = C.SPARC_INS_MEMBAR
	SPARC_INS_MOVDTOX     = C.SPARC_INS_MOVDTOX
	SPARC_INS_MOV         = C.SPARC_INS_MOV
	SPARC_INS_MOVRGEZ     = C.SPARC_INS_MOVRGEZ
	SPARC_INS_MOVRGZ      = C.SPARC_INS_MOVRGZ
	SPARC_INS_MOVRLEZ     = C.SPARC_INS_MOVRLEZ
	SPARC_INS_MOVRLZ      = C.SPARC_INS_MOVRLZ
	SPARC_INS_MOVRNZ      = C.SPARC_INS_MOVRNZ
	SPARC_INS_MOVRZ       = C.SPARC_INS_MOVRZ
	SPARC_INS_MOVSTOSW    = C.SPARC_INS_MOVSTOSW
	SPARC_INS_MOVSTOUW    = C.SPARC_INS_MOVSTOUW
	SPARC_INS_MULX        = C.SPARC_INS_MULX
	SPARC_INS_NOP         = C.SPARC_INS_NOP
	SPARC_INS_ORCC        = C.SPARC_INS_ORCC
	SPARC_INS_ORNCC       = C.SPARC_INS_ORNCC
	SPARC_INS_ORN         = C.SPARC_INS_ORN
	SPARC_INS_OR          = C.SPARC_INS_OR
	SPARC_INS_PDIST       = C.SPARC_INS_PDIST
	SPARC_INS_PDISTN      = C.SPARC_INS_PDISTN
	SPARC_INS_POPC        = C.SPARC_INS_POPC
	SPARC_INS_RD          = C.SPARC_INS_RD
	SPARC_INS_RESTORE     = C.SPARC_INS_RESTORE
	SPARC_INS_RETT        = C.SPARC_INS_RETT
	SPARC_INS_SAVE        = C.SPARC_INS_SAVE
	SPARC_INS_SDIVCC      = C.SPARC_INS_SDIVCC
	SPARC_INS_SDIVX       = C.SPARC_INS_SDIVX
	SPARC_INS_SDIV        = C.SPARC_INS_SDIV
	SPARC_INS_SETHI       = C.SPARC_INS_SETHI
	SPARC_INS_SHUTDOWN    = C.SPARC_INS_SHUTDOWN
	SPARC_INS_SIAM        = C.SPARC_INS_SIAM
	SPARC_INS_SLLX        = C.SPARC_INS_SLLX
	SPARC_INS_SLL         = C.SPARC_INS_SLL
	SPARC_INS_SMULCC      = C.SPARC_INS_SMULCC
	SPARC_INS_SMUL        = C.SPARC_INS_SMUL
	SPARC_INS_SRAX        = C.SPARC_INS_SRAX
	SPARC_INS_SRA         = C.SPARC_INS_SRA
	SPARC_INS_SRLX        = C.SPARC_INS_SRLX
	SPARC_INS_SRL         = C.SPARC_INS_SRL
	SPARC_INS_STBAR       = C.SPARC_INS_STBAR
	SPARC_INS_STB         = C.SPARC_INS_STB
	SPARC_INS_STD         = C.SPARC_INS_STD
	SPARC_INS_ST          = C.SPARC_INS_ST
	SPARC_INS_STH         = C.SPARC_INS_STH
	SPARC_INS_STQ         = C.SPARC_INS_STQ
	SPARC_INS_STX         = C.SPARC_INS_STX
	SPARC_INS_SUBCC       = C.SPARC_INS_SUBCC
	SPARC_INS_SUBX        = C.SPARC_INS_SUBX
	SPARC_INS_SUBXCC      = C.SPARC_INS_SUBXCC
	SPARC_INS_SUB         = C.SPARC_INS_SUB
	SPARC_INS_SWAP        = C.SPARC_INS_SWAP
	SPARC_INS_TADDCCTV    = C.SPARC_INS_TADDCCTV
	SPARC_INS_TADDCC      = C.SPARC_INS_TADDCC
	SPARC_INS_T           = C.SPARC_INS_T
	SPARC_INS_TSUBCCTV    = C.SPARC_INS_TSUBCCTV
	SPARC_INS_TSUBCC      = C.SPARC_INS_TSUBCC
	SPARC_INS_UDIVCC      = C.SPARC_INS_UDIVCC
	SPARC_INS_UDIVX       = C.SPARC_INS_UDIVX
	SPARC_INS_UDIV        = C.SPARC_INS_UDIV
	SPARC_INS_UMULCC      = C.SPARC_INS_UMULCC
	SPARC_INS_UMULXHI     = C.SPARC_INS_UMULXHI
	SPARC_INS_UMUL        = C.SPARC_INS_UMUL
	SPARC_INS_UNIMP       = C.SPARC_INS_UNIMP
	SPARC_INS_FCMPED      = C.SPARC_INS_FCMPED
	SPARC_INS_FCMPEQ      = C.SPARC_INS_FCMPEQ
	SPARC_INS_FCMPES      = C.SPARC_INS_FCMPES
	SPARC_INS_WR          = C.SPARC_INS_WR
	SPARC_INS_XMULX       = C.SPARC_INS_XMULX
	SPARC_INS_XMULXHI     = C.SPARC_INS_XMULXHI
	SPARC_INS_XNORCC      = C.SPARC_INS_XNORCC
	SPARC_INS_XNOR        = C.SPARC_INS_XNOR
	SPARC_INS_XORCC       = C.SPARC_INS_XORCC
	SPARC_INS_XOR         = C.SPARC_INS_XOR
	SPARC_INS_RET         = C.SPARC_INS_RET
	SPARC_INS_RETL        = C.SPARC_INS_RETL
	SPARC_INS_ENDING      = C.SPARC_INS_ENDING
)

const (
	SPARC_GRP_INVALID  = C.SPARC_GRP_INVALID
	SPARC_GRP_JUMP     = C.SPARC_GRP_JUMP
	SPARC_GRP_HARDQUAD = C.SPARC_GRP_HARDQUAD
	SPARC_GRP_V9       = C.SPARC_GRP_V9
	SPARC_GRP_VIS      = C.SPARC_GRP_VIS
	SPARC_GRP_VIS2     = C.SPARC_GRP_VIS2
	SPARC_GRP_VIS3     = C.SPARC_GRP_VIS3
	SPARC_GRP_32BIT    = C.SPARC_GRP_32BIT
	SPARC_GRP_64BIT    = C.SPARC_GRP_64BIT
	SPARC_GRP_ENDING   = C.SPARC_GRP_ENDING
)
