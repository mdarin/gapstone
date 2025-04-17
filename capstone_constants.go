/*
Gapstone is a Go binding for the Capstone disassembly library.
For examples, try reading the *_test.go files.

	Library Author: Nguyen Anh Quynh
	Binding Author: Ben Nagy
	License: BSD style - see LICENSE file for details
    (c) 2013 COSEINC. All Rights Reserved.

    THIS FILE WAS AUTO-GENERATED -- DO NOT EDIT!
	Command: ./genconst.rb capstone/bindings/python/capstone/
	Created at: 2025-04-14T19:17:55+00:00

*/

package gapstone

// #cgo LDFLAGS: -lcapstone
// #cgo freebsd CFLAGS: -I/usr/local/include
// #cgo freebsd LDFLAGS: -L/usr/local/lib
// #include <stdlib.h>
// #include <capstone/capstone.h>
import "C"

// Capstone Python bindings, by Nguyen Anh Quynnh <aquynh@gmail.com>
// Capstone C interface
// API version
const (
	CS_API_MAJOR = C.CS_API_MAJOR
	CS_API_MINOR = C.CS_API_MINOR
)

// Package version
const (
	CS_VERSION_MAJOR = C.CS_VERSION_MAJOR
	CS_VERSION_MINOR = C.CS_VERSION_MINOR
	CS_VERSION_EXTRA = C.CS_VERSION_EXTRA
)

// architectures
const (
	CS_ARCH_ARM        = C.CS_ARCH_ARM
	CS_ARCH_ARM64      = C.CS_ARCH_ARM64
	CS_ARCH_MIPS       = C.CS_ARCH_MIPS
	CS_ARCH_X86        = C.CS_ARCH_X86
	CS_ARCH_PPC        = C.CS_ARCH_PPC
	CS_ARCH_SPARC      = C.CS_ARCH_SPARC
	CS_ARCH_SYSZ       = C.CS_ARCH_SYSZ
	CS_ARCH_XCORE      = C.CS_ARCH_XCORE
	CS_ARCH_M68K       = C.CS_ARCH_M68K
	CS_ARCH_TMS320C64X = C.CS_ARCH_TMS320C64X
	CS_ARCH_M680X      = C.CS_ARCH_M680X
	CS_ARCH_EVM        = C.CS_ARCH_EVM
	CS_ARCH_MOS65XX    = C.CS_ARCH_MOS65XX
	CS_ARCH_WASM       = C.CS_ARCH_WASM
	CS_ARCH_BPF        = C.CS_ARCH_BPF
	CS_ARCH_RISCV      = C.CS_ARCH_RISCV
	CS_ARCH_MAX        = C.CS_ARCH_MAX
	CS_ARCH_ALL        = C.CS_ARCH_ALL
)

// disasm mode
const (
	CS_MODE_LITTLE_ENDIAN         = C.CS_MODE_LITTLE_ENDIAN
	CS_MODE_ARM                   = C.CS_MODE_ARM
	CS_MODE_16                    = C.CS_MODE_16
	CS_MODE_32                    = C.CS_MODE_32
	CS_MODE_64                    = C.CS_MODE_64
	CS_MODE_THUMB                 = C.CS_MODE_THUMB
	CS_MODE_MCLASS                = C.CS_MODE_MCLASS
	CS_MODE_V8                    = C.CS_MODE_V8
	CS_MODE_MICRO                 = C.CS_MODE_MICRO
	CS_MODE_MIPS3                 = C.CS_MODE_MIPS3
	CS_MODE_MIPS32R6              = C.CS_MODE_MIPS32R6
	CS_MODE_MIPS2                 = C.CS_MODE_MIPS2
	CS_MODE_V9                    = C.CS_MODE_V9
	CS_MODE_QPX                   = C.CS_MODE_QPX
	CS_MODE_SPE                   = C.CS_MODE_SPE
	CS_MODE_BOOKE                 = C.CS_MODE_BOOKE
	CS_MODE_M68K_000              = C.CS_MODE_M68K_000
	CS_MODE_M68K_010              = C.CS_MODE_M68K_010
	CS_MODE_M68K_020              = C.CS_MODE_M68K_020
	CS_MODE_M68K_030              = C.CS_MODE_M68K_030
	CS_MODE_M68K_040              = C.CS_MODE_M68K_040
	CS_MODE_M68K_060              = C.CS_MODE_M68K_060
	CS_MODE_BIG_ENDIAN            = C.CS_MODE_BIG_ENDIAN
	CS_MODE_MIPS32                = C.CS_MODE_MIPS32
	CS_MODE_MIPS64                = C.CS_MODE_MIPS64
	CS_MODE_M680X_6301            = C.CS_MODE_M680X_6301
	CS_MODE_M680X_6309            = C.CS_MODE_M680X_6309
	CS_MODE_M680X_6800            = C.CS_MODE_M680X_6800
	CS_MODE_M680X_6801            = C.CS_MODE_M680X_6801
	CS_MODE_M680X_6805            = C.CS_MODE_M680X_6805
	CS_MODE_M680X_6808            = C.CS_MODE_M680X_6808
	CS_MODE_M680X_6809            = C.CS_MODE_M680X_6809
	CS_MODE_M680X_6811            = C.CS_MODE_M680X_6811
	CS_MODE_M680X_CPU12           = C.CS_MODE_M680X_CPU12
	CS_MODE_M680X_HCS08           = C.CS_MODE_M680X_HCS08
	CS_MODE_BPF_CLASSIC           = C.CS_MODE_BPF_CLASSIC
	CS_MODE_BPF_EXTENDED          = C.CS_MODE_BPF_EXTENDED
	CS_MODE_RISCV32               = C.CS_MODE_RISCV32
	CS_MODE_RISCV64               = C.CS_MODE_RISCV64
	CS_MODE_RISCVC                = C.CS_MODE_RISCVC
	CS_MODE_MOS65XX_6502          = C.CS_MODE_MOS65XX_6502
	CS_MODE_MOS65XX_65C02         = C.CS_MODE_MOS65XX_65C02
	CS_MODE_MOS65XX_W65C02        = C.CS_MODE_MOS65XX_W65C02
	CS_MODE_MOS65XX_65816         = C.CS_MODE_MOS65XX_65816
	CS_MODE_MOS65XX_65816_LONG_M  = C.CS_MODE_MOS65XX_65816_LONG_M
	CS_MODE_MOS65XX_65816_LONG_X  = C.CS_MODE_MOS65XX_65816_LONG_X
	CS_MODE_MOS65XX_65816_LONG_MX = C.CS_MODE_MOS65XX_65816_LONG_MX
)

// Capstone option type
const (
	CS_OPT_SYNTAX         = C.CS_OPT_SYNTAX
	CS_OPT_DETAIL         = C.CS_OPT_DETAIL
	CS_OPT_MODE           = C.CS_OPT_MODE
	CS_OPT_MEM            = C.CS_OPT_MEM
	CS_OPT_SKIPDATA       = C.CS_OPT_SKIPDATA
	CS_OPT_SKIPDATA_SETUP = C.CS_OPT_SKIPDATA_SETUP
	CS_OPT_MNEMONIC       = C.CS_OPT_MNEMONIC
	CS_OPT_UNSIGNED       = C.CS_OPT_UNSIGNED
)

// Capstone option value
const (
	CS_OPT_OFF = C.CS_OPT_OFF
	CS_OPT_ON  = C.CS_OPT_ON
)

// Common instruction operand types - to be consistent across all architectures.
const (
	CS_OP_INVALID = C.CS_OP_INVALID
	CS_OP_REG     = C.CS_OP_REG
	CS_OP_IMM     = C.CS_OP_IMM
	CS_OP_MEM     = C.CS_OP_MEM
	CS_OP_FP      = C.CS_OP_FP
)

// Common instruction groups - to be consistent across all architectures.
const (
	CS_GRP_INVALID   = C.CS_GRP_INVALID
	CS_GRP_JUMP      = C.CS_GRP_JUMP
	CS_GRP_CALL      = C.CS_GRP_CALL
	CS_GRP_RET       = C.CS_GRP_RET
	CS_GRP_INT       = C.CS_GRP_INT
	CS_GRP_IRET      = C.CS_GRP_IRET
	CS_GRP_PRIVILEGE = C.CS_GRP_PRIVILEGE
)

// Access types for instruction operands.
const (
	CS_AC_INVALID = C.CS_AC_INVALID
	CS_AC_READ    = C.CS_AC_READ
	CS_AC_WRITE   = C.CS_AC_WRITE
)

// Capstone syntax value
const (
	CS_OPT_SYNTAX_DEFAULT   = C.CS_OPT_SYNTAX_DEFAULT
	CS_OPT_SYNTAX_INTEL     = C.CS_OPT_SYNTAX_INTEL
	CS_OPT_SYNTAX_ATT       = C.CS_OPT_SYNTAX_ATT
	CS_OPT_SYNTAX_NOREGNAME = C.CS_OPT_SYNTAX_NOREGNAME
	CS_OPT_SYNTAX_MASM      = C.CS_OPT_SYNTAX_MASM
	CS_OPT_SYNTAX_MOTOROLA  = C.CS_OPT_SYNTAX_MOTOROLA
)

// Capstone error type
const (
	CS_ERR_OK        = C.CS_ERR_OK
	CS_ERR_MEM       = C.CS_ERR_MEM
	CS_ERR_ARCH      = C.CS_ERR_ARCH
	CS_ERR_HANDLE    = C.CS_ERR_HANDLE
	CS_ERR_CSH       = C.CS_ERR_CSH
	CS_ERR_MODE      = C.CS_ERR_MODE
	CS_ERR_OPTION    = C.CS_ERR_OPTION
	CS_ERR_DETAIL    = C.CS_ERR_DETAIL
	CS_ERR_MEMSETUP  = C.CS_ERR_MEMSETUP
	CS_ERR_VERSION   = C.CS_ERR_VERSION
	CS_ERR_DIET      = C.CS_ERR_DIET
	CS_ERR_SKIPDATA  = C.CS_ERR_SKIPDATA
	CS_ERR_X86_ATT   = C.CS_ERR_X86_ATT
	CS_ERR_X86_INTEL = C.CS_ERR_X86_INTEL
	CS_ERR_X86_MASM  = C.CS_ERR_X86_MASM
)

// query id for cs_support()
const (
	CS_SUPPORT_DIET       = C.CS_SUPPORT_DIET
	CS_SUPPORT_X86_REDUCE = C.CS_SUPPORT_X86_REDUCE
)

// Capstone reverse lookup
const (
// CS_AC = C.CS_AC
// CS_ARCH = C.CS_ARCH
// CS_ERR = C.CS_ERR
// CS_GRP = C.CS_GRP
// CS_MODE = C.CS_MODE
// CS_OP = C.CS_OP
// CS_OPT = C.CS_OPT
)

// Loading attempts, in order
// - user-provided environment variable
// - pkg_resources can get us the path to the local libraries
// - we can get the path to the local libraries by parsing our filename
// - global load
// - python's lib directory
// - last-gasp attempt at some hardcoded paths on darwin and linux
// low-level structure for C code
// Weird import placement because these modules are needed by the below code but need the above functions
// callback for SKIPDATA option
const (
// CS_SKIPDATA_CALLBACK = C.CS_SKIPDATA_CALLBACK
)

// setup all the function prototype
// access to error code via @errno of CsError
// return the core's version
// return the binding's version
// dummy class resembling Cs class, just for cs_disasm_quick()
// this class only need to be referenced to via 2 fields: @csh & @arch
// Quick & dirty Python function to disasm raw binary code
// This function return CsInsn objects
// NOTE: you might want to use more efficient Cs class & its methods.
// Another quick, but lighter function to disasm raw binary code.
// This function is faster than cs_disasm_quick() around 20% because
// cs_disasm_lite() only return tuples of (address, size, mnemonic, op_str),
// rather than CsInsn objects.
// NOTE: you might want to use more efficient Cs class & its methods.
// Python-style class to disasm code
// print out debugging info
