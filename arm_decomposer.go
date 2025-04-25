/*
Gapstone is a Go binding for the Capstone disassembly library. For examples,
try reading the *_test.go files.

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

// FIXME: somthing goes wrond for thumb2 decomposing for ex..

// Accessed via insn.Arm.XXX
type ArmInstruction struct {
	UserMode    bool         ///< User-mode registers to be loaded (for LDM/STM instructions)
	VectorSize  int          ///< Scalar size for vector instructions
	VectorData  int          ///< Data type for elements of vector instructions
	CPSMode     int          ///< CPS mode for CPS instruction
	CPSFlag     int          ///< CPS mode for CPS instruction //? flag
	CC          uint         ///< conditional code for this insn
	UpdateFlags bool         ///< does this insn update flags?
	Writeback   bool         ///< does this insn write-back?
	MemBarrier  int          ///< Option for some memory barrier instructions
	Operands    []ArmOperand ///< operands for this instruction.
}

type ArmShifter struct {
	Type  uint
	Value uint
}

type ArmOperand struct {
	VectorIndex int ///< Vector Index for some vector operands (or -1 if irrelevant)///< Vector Index for some vector operands (or -1 if irrelevant)
	Shift       ArmShifter
	// operand type
	Type   uint             // ARM_OP_* - determines which field is set below
	Reg    uint             ///< register value for REG/SYSREG operand
	Imm    int32            ///< immediate value for C-IMM, P-IMM or IMM operand
	FP     float64          ///< floating point value for FP operand
	Mem    ArmMemoryOperand ///< base/index/scale/disp value for MEM operand
	Setend int              ///< SETEND instruction's operand type

	/// in some instructions, an operand can be subtracted or added to
	/// the base register,
	/// if TRUE, this operand is subtracted. otherwise, it is added.
	Subtracted bool

	/// How is this operand accessed? (READ, WRITE or READ|WRITE)
	/// This field is combined of cs_ac_type.
	/// NOTE: this field is irrelevant if engine is compiled in DIET mode.
	Access uint8
	/// Neon lane index for NEON instructions (or -1 if irrelevant)
	NeonLane int8
}

type ArmMemoryOperand struct {
	Base  uint ///< base register
	Index uint ///< index register
	Scale int  ///< scale for index register (can be 1, or -1)
	Disp  int  ///< displacement/offset value
	/// left-shift on index register, or 0 if irrelevant
	/// INFO: this value can also be fetched via operand.shift.value
	LShift int
}

// Number of Operands of a given ARM_OP_* type
func (insn ArmInstruction) OpCount(optype uint) int {
	count := 0
	for _, op := range insn.Operands {
		if op.Type == optype {
			count++
		}
	}
	return count
}

func fillArmHeader(raw C.cs_insn, insn *Instruction) {

	if raw.detail == nil {
		return
	}

	// Cast the cs_detail union
	cs_arm := (*C.cs_arm)(unsafe.Pointer(&raw.detail.anon0[0]))

	arm := ArmInstruction{
		UserMode:    bool(cs_arm.usermode),
		VectorSize:  int(cs_arm.vector_size),
		VectorData:  int(cs_arm.vector_data),
		CPSMode:     int(cs_arm.cps_mode),
		CPSFlag:     int(cs_arm.cps_flag),
		CC:          uint(cs_arm.cc),
		UpdateFlags: bool(cs_arm.update_flags),
		Writeback:   bool(cs_arm.writeback),
		MemBarrier:  int(cs_arm.mem_barrier),
	}

	// Cast the op_info to a []C.cs_arm_op
	var ops []C.cs_arm_op
	h := (*reflect.SliceHeader)(unsafe.Pointer(&ops))
	h.Data = uintptr(unsafe.Pointer(&cs_arm.operands[0]))
	h.Len = int(cs_arm.op_count)
	h.Cap = int(cs_arm.op_count)

	// Create the Go object for each operand
	for _, cop := range ops {
		if cop._type == ARM_OP_INVALID {
			break
		}
		gop := ArmOperand{
			Shift: ArmShifter{
				Type:  uint(cop.shift._type),
				Value: uint(cop.shift.value),
			},
			Type:        uint(cop._type),
			VectorIndex: int(cop.vector_index),
			Subtracted:  bool(cop.subtracted),
			Access:      uint8(cop.access),
			NeonLane:    int8(cop.neon_lane),
		}
		switch cop._type {
		// fake a union by setting only the correct struct member
		case ARM_OP_IMM, ARM_OP_CIMM, ARM_OP_PIMM:
			gop.Imm = int32(*(*C.int32_t)(unsafe.Pointer(&cop.anon0[0])))
		case ARM_OP_FP:
			gop.FP = float64(*(*C.double)(unsafe.Pointer(&cop.anon0[0])))
		case ARM_OP_REG, ARM_OP_SYSREG:
			gop.Reg = uint(*(*C.int)(unsafe.Pointer(&cop.anon0[0])))
		case ARM_OP_MEM:
			cmop := (*C.arm_op_mem)(unsafe.Pointer(&cop.anon0[0]))
			gop.Mem = ArmMemoryOperand{
				Base:   uint(cmop.base),
				Index:  uint(cmop.index),
				Scale:  int(cmop.scale),
				Disp:   int(cmop.disp),
				LShift: int(cmop.lshift),
			}
		case ARM_OP_SETEND:
			gop.Setend = int(*(*C.int)(unsafe.Pointer(&cop.anon0[0])))
		}
		arm.Operands = append(arm.Operands, gop)
	}
	insn.Arm = &arm
}

func decomposeArm(e *Engine, raws []C.cs_insn) []Instruction {
	decomposed := []Instruction{}
	for _, raw := range raws {
		decomp := new(Instruction)
		fillGenericHeader(e, raw, decomp)
		fillArmHeader(raw, decomp)
		decomposed = append(decomposed, *decomp)
	}
	return decomposed
}
