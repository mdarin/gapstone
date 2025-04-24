package gapstone

import (
	"reflect"
	"unsafe"
)

// Accessed via insn.Riscv.XXX
type RISCVInstruction struct {
	CC          uint         ///< Conditional code for this instruction
	UpdateFlags bool         ///< Does this instruction update flags?
	Writeback   bool         ///< Does this instruction write back?
	MemBarrier  int          ///< Memory barrier type (if any)
	Operands    []RISCVOperand ///< Operands for this instruction
}

type RISCVShifter struct {
	Type  uint
	Value uint
}

type RISCVOperand struct {
	VectorIndex int   ///< Vector Index for some vector operands (or -1 if irrelevant)
	Shift       RISCVShifter
	Type        uint              // RISCV_OP_* - determines which field is set below
	Reg         uint              ///< Register value for REG operand
	Imm         int64             ///< Immediate value for IMM operand 
	FP          float64           ///< Floating point value for FP operand
	Mem         RISCVMemoryOperand ///< Base/index/scale/disp value for MEM operand
	
	/// How is this operand accessed? (READ, WRITE or READ|WRITE)
	Access uint8
	
	/// Neon lane index for NEON instructions (or -1 if irrelevant)
	NeonLane int8
}

type RISCVMemoryOperand struct {
	Base  uint ///< Base register
	Index uint ///< Index register
	Scale int  ///< Scale factor for index register 
	Disp  int  ///< Displacement/offset value
	LShift int ///< Left shift amount for index register (or 0 if irrelevant)
}

func fillRiscvHeader(raw C.cs_insn, insn *Instruction) {
	if raw.detail == nil {
		return
	}

	// Cast the cs_detail union
	cs_riscv := (*C.cs_riscv)(unsafe.Pointer(&raw.detail.anon0[0]))

	riscvInsn := RISCVInstruction{
		CC:          uint(cs_riscv.cc),
		UpdateFlags: bool(cs_riscv.update_flags),
		Writeback:   bool(cs_riscv.writeback),
		MemBarrier:  int(cs_riscv.mem_barrier),
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
			Shift: RISCVShifter{
				Type:  uint(cop.shift._type),
				Value: uint(cop.shift.value),
			},
			Type:        uint(cop._type),
			VectorIndex: int(cop.vector_index),
			Access:      uint8(cop.access),
			NeonLane:    int8(cop.neon_lane),
		}

		switch cop._type {
		case RISCV_OP_IMM, RISCV_OP_CIMM:
			gop.Imm = int64(*(*C.int64_t)(unsafe.Pointer(&cop.anon0[0])))
		case RISCV_OP_FP:
			gop.FP = float64(*(*C.double)(unsafe.Pointer(&cop.anon0[0])))
		case RISCV_OP_REG:
			gop.Reg = uint(*(*C.uint)(unsafe.Pointer(&cop.anon0[0])))
		case RISCV_OP_MEM:
			cmop := (*C.riscv_op_mem)(unsafe.Pointer(&cop.anon0[0]))
			gop.Mem = RISCVMemoryOperand{
				Base:   uint(cmop.base),
				Index:  uint(cmop.index), 
				Scale:  int(cmop.scale),
				Disp:   int(cmop.disp),
				LShift: int(cmop.lshift),
			}
		}
		
		riscvInsn.Operands = append(riscvInsn.Operands, gop)
	}

	insn.Riscv = &riscvInsn
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
