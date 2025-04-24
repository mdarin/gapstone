/*
Gapstone is a Go binding for the Capstone disassembly library.
For examples, try reading the *_test.go files.

    Library Author: Nguyen Anh Quynh
    Binding Author: Ben Nagy
    License: BSD style - see LICENSE file for details
    (c) 2013 COSEINC. All Rights Reserved.
*/

package gapstone

// #cgo LDFLAGS: -lcapstone
// #cgo freebsd CFLAGS: -I/usr/local/include
// #cgo freebsd LDFLAGS: -L/usr/local/lib
// #include <stdlib.h>
// #include <capstone/capstone.h>
import "C"

import (
	"reflect"
	"unsafe"
)

// Accessed via insn.BPF.XXX
type BpfInstruction struct {
	Operands []BpfOperand
}

type BpfOperand struct {
	Type uint             // BPF_OP_* - determines which field is set below
	Reg  uint8            ///< register value for REG operand
	Mem  BpfMemoryOperand ///< base/disp value for MEM operand
	Imm  uint64           ///< immediate value IMM operand
	Off  uint32           ///< offset value, used in jump & call
	Mmem uint32           ///< M[k] in cBPF [cBPF only]
	Msh  uint32           ///< corresponds to cBPF's BPF_MSH mode [cBPF only]
	Ext  uint32           ///< cBPF's extension (not eBPF) [cBPF only]
}

type BpfMemoryOperand struct {
	Base uint
	Disp int32
}

// Number of Operands of a given BPF_OP_* type (uint8 originaly)
func (insn BpfInstruction) OpCount(optype uint) int {
	count := 0
	for _, op := range insn.Operands {
		if op.Type == optype {
			count++
		}
	}
	return count
}

func fillBpfHeader(raw C.cs_insn, insn *Instruction) {
	// detail can be NULL on "data" instruction if SKIPDATA option is turned ON
	if raw.detail == nil {
		return
	}

	// Cast the cs_detail union
	cs_bpf := (*C.cs_bpf)(unsafe.Pointer(&raw.detail.anon0[0]))

	bpf := BpfInstruction{}

	// Cast the op_info to a []C.cs_bpf_op
	var ops []C.cs_bpf_op

	h := (*reflect.SliceHeader)(unsafe.Pointer(&ops)) // FIXME: reflect.SliceHeader is deprecated: Use unsafe.Slice or unsafe.SliceData
	h.Data = uintptr(unsafe.Pointer(&cs_bpf.operands[0]))
	h.Len = int(cs_bpf.op_count)
	h.Cap = int(cs_bpf.op_count)

	bpf.Operands = make([]BpfOperand, 0, len(ops))

	// Create the Go object for each operand
	for _, cop := range ops {

		if cop._type == BPF_OP_INVALID {
			break
		}

		gop := BpfOperand{
			Type: uint(cop._type),
		}

		switch cop._type {
		case BPF_OP_IMM:
			gop.Imm = uint64(*(*C.int64_t)(unsafe.Pointer(&cop.anon0[0])))

		case BPF_OP_REG:
			gop.Reg = uint8(*(*C.uint8_t)(unsafe.Pointer(&cop.anon0[0])))

		case BPF_OP_MEM:
			cmop := (*C.bpf_op_mem)(unsafe.Pointer(&cop.anon0[0]))
			gop.Mem = BpfMemoryOperand{
				Base: uint(cmop.base),
				Disp: int32(cmop.disp),
			}

		case BPF_OP_OFF:
			gop.Off = uint32(*(*C.uint32_t)(unsafe.Pointer(&cop.anon0[0])))

		case BPF_OP_MMEM:
			gop.Mmem = uint32(*(*C.uint32_t)(unsafe.Pointer(&cop.anon0[0])))

		case BPF_OP_MSH:
			gop.Msh = uint32(*(*C.uint32_t)(unsafe.Pointer(&cop.anon0[0])))

		case BPF_OP_EXT:
			gop.Ext = uint32(*(*C.uint32_t)(unsafe.Pointer(&cop.anon0[0])))

		}

		bpf.Operands = append(bpf.Operands, gop)
	}

	insn.BPF = &bpf
}

func decomposeBpf(e *Engine, raws []C.cs_insn) []Instruction {
	decomposed := []Instruction{}
	for _, raw := range raws {
		decomp := new(Instruction)
		fillGenericHeader(e, raw, decomp)
		fillBpfHeader(raw, decomp)
		decomposed = append(decomposed, *decomp)
	}
	return decomposed
}
