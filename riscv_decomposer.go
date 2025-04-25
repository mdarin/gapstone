/*
Gapstone is a Go binding for the Capstone disassembly library.
For examples, try reading the *_test.go files.

	Library Author: Nguyen Anh Quynh
	Binding Author: Ben Nagy
	License: BSD style - see LICENSE file for details
	(c) 2013 COSEINC. All Rights Reserved.
*/
package gapstone

import (
	"reflect"
	"unsafe"
)

// #cgo LDFLAGS: -lcapstone
// #cgo freebsd CFLAGS: -I/usr/local/include
// #cgo freebsd LDFLAGS: -L/usr/local/lib
// #include <stdlib.h>
// #include <capstone/capstone.h>
import "C"

// Accessed via insn.RISCV.XXX

// Instruction structure
type RISCVInstruction struct {
	NeedEffectiveAddr bool           ///< // Does this instruction need effective address or not.
	Operands          []RISCVOperand ///< Operands for this instruction
}

// Instruction operand
type RISCVOperand struct {
	Type uint               // RISCV_OP_* - determines which field is set below
	Reg  uint               ///< Register value for REG operand
	Imm  uint64             ///< Immediate value for IMM operand
	Mem  RISCVMemoryOperand ///< Base/disp value for MEM operand
}

// Instruction's operand referring to memory
type RISCVMemoryOperand struct {
	Base uint  ///< Base register
	Disp int64 ///< Displacement/offset value
}

func fillRiscvHeader(raw C.cs_insn, insn *Instruction) {
	if raw.detail == nil {
		return
	}

	// Cast the cs_detail union
	cs_riscv := (*C.cs_riscv)(unsafe.Pointer(&raw.detail.anon0[0]))

	riscvInsn := RISCVInstruction{
		NeedEffectiveAddr: bool(cs_riscv.need_effective_addr),
	}

	// Cast the op_info to a []C.cs_riscv_op
	var ops []C.cs_riscv_op
	h := (*reflect.SliceHeader)(unsafe.Pointer(&ops))
	h.Data = uintptr(unsafe.Pointer(&cs_riscv.operands[0]))
	h.Len = int(cs_riscv.op_count)
	h.Cap = int(cs_riscv.op_count)

	// Create the Go object for each operand
	for _, cop := range ops {
		if cop._type == RISCV_OP_INVALID {
			break
		}

		gop := RISCVOperand{
			Type: uint(cop._type),
		}

		switch cop._type {
		case RISCV_OP_IMM:
			gop.Imm = uint64(*(*C.int64_t)(unsafe.Pointer(&cop.anon0[0])))

		case RISCV_OP_REG:
			gop.Reg = uint(*(*C.uint)(unsafe.Pointer(&cop.anon0[0])))

		case RISCV_OP_MEM:
			cmop := (*C.riscv_op_mem)(unsafe.Pointer(&cop.anon0[0]))
			gop.Mem = RISCVMemoryOperand{
				Base: uint(cmop.base),
				Disp: int64(cmop.disp),
			}
		}

		riscvInsn.Operands = append(riscvInsn.Operands, gop)
	}

	insn.RISCV = &riscvInsn
}

func decomposeRiscv(e *Engine, raws []C.cs_insn) []Instruction {
	decomposed := []Instruction{}
	for _, raw := range raws {
		decomp := new(Instruction)
		fillGenericHeader(e, raw, decomp)
		fillRiscvHeader(raw, decomp)
		decomposed = append(decomposed, *decomp)
	}
	return decomposed
}
