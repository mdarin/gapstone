package arm64

import (
	"unsafe"
)

// Arm64OpMem represents memory operand information
type Arm64OpMem struct {
	Base   uint32
	Index  uint32
	Disp   int32
}

// Arm64OpShift represents shift information
type Arm64OpShift struct {
	Type uint32
	Value uint32
}

// Arm64OpValue represents different possible operand values
type Arm64OpValue struct {
	Reg      uint32
	Imm      int64
	Fp       float64
	Mem      Arm64OpMem
	Pstate   int32
	Sys      uint32
	Prefetch int32
	Barrier  int32
}

// Arm64Op represents an ARM64 instruction operand
type Arm64Op struct {
	VectorIndex int32
	Vas         int32
	Shift       Arm64OpShift
	Ext         uint32
	Type        uint32
	Value       Arm64OpValue
	Access      uint8
}

// CsArm64 represents an ARM64 instruction
type CsArm64 struct {
Cc           uint32
UpdateFlags  bool
Writeback    bool
OpCount      byte
Operands     [8]Arm64Op
}

// GetArchInfo returns architecture specific information
func GetArchInfo(cs *CsArm64) (uint32, bool, bool, []Arm64Op) {
	return cs.Cc, cs.UpdateFlags, cs.Writeback, cs.Operands[:cs.OpCount]
}
