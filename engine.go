/*
Gapstone is a Go binding for the Capstone disassembly library. For examples,
try reading the *_test.go files.

	Library Author: Nguyen Anh Quynh
	Binding Author: Ben Nagy
	License: BSD style - see LICENSE file for details
    (c) 2013 COSEINC. All Rights Reserved.
*/

/*
Package gapstone provides a Go binding for the Capstone disassembly framework.

Capstone is a high-performance disassembly engine supporting multiple architectures
and providing detailed instruction information including operands, registers,
and semantic groups.

Features:
- Architecture-independent API
- Detailed instruction decomposition
- Support for x86, ARM, ARM64, MIPS, and other architectures
- Thread-safe design
- Customizable output formatting

Example usage:

	engine, err := capstone.New(capstone.CS_ARCH_X86, capstone.CS_MODE_64)
	if err != nil {
	    panic(err)
	}
	defer engine.Close()

	codeBytes := []byte{0x55, 0x48, 0x8b, 0x05, 0xab, 0xcd, 0xef, 0x12}
	instructions, err := engine.Disasm(codeBytes, 0x4000, 0)
	if err != nil {
	    panic(err)
	}

	for _, insn := range instructions {
	    fmt.Printf("0x%x:\t%s\t%s\n", insn.Address, insn.Mnemonic, insn.OpStr)
	}
*/
package gapstone

// #cgo LDFLAGS: -lcapstone
// #cgo freebsd CFLAGS: -I/usr/local/include
// #cgo freebsd LDFLAGS: -L/usr/local/lib
// #include <stdlib.h>
// #include <capstone/capstone.h>
// extern size_t trampoline(uint8_t *buffer, size_t buflen, size_t offset, void *user_data);
import "C"

import (
	"reflect"
	"unsafe"
	// "github.com/samber/lo"
)

// Since this is a build-time option for the C lib, it seems logical to have
// this as a static flag.
// Diet Mode Changes:
// - No regs_read, regs_written or groups
// - No response to reg_name or insn_name
// - No mnemonic or op_str
// If you want to see any operands in diet mode, then you need CS_DETAIL.
var dietMode = bool(C.cs_support(CS_SUPPORT_DIET))

// The arch and mode given at create time will determine how code is
// disassembled. After use you must close an Engine with engine.Close() to allow
// the C lib to free resources.
/* Engine represents a Capstone disassembler instance configured for a specific architecture and mode */
type Engine struct {
	handle   C.csh              // Handle to the underlying Capstone engine
	arch     int                // Architecture identifier (CS_ARCH_*)
	mode     int                // Mode identifier (CS_MODE_*)
	skipdata *C.cs_opt_skipdata // Skip data configuration pointer
}

// Information that exists for every Instruction, regardless of arch.
// Structure members here will be promoted, so every Instruction will have
// them available. Check the constants for each architecture for available
// Instruction groups etc.
/* InstructionHeader contains basic information common to all instructions */
type InstructionHeader struct {
	Id       uint   `json:"id"`       // Internal id for this instruction. Subject to change.
	Address  uint   `json:"address"`  // Nominal address ($ip) of this instruction
	Size     uint   `json:"size"`     // Size of the instruction in bytes
	Bytes    []byte `json:"bytes"`    // Raw instruction bytes as they appear in memory
	Mnemonic string `json:"mnemonic"` // Instruction mnemonic (e.g., "add", "mov")
	OpStr    string `json:"op_str"`   // Formatted operand string for the instruction
	// Not available without the decomposer. BE CAREFUL! By default,
	// CS_OPT_DETAIL is set to CS_OPT_OFF so the result of accessing these
	// members is undefined.
	AllRegistersRead    []uint // List of implicit and explicit registers read by this instruction
	AllRegistersWritten []uint // List of implicit and explicit registers written by this instruction
	RegistersRead       []uint // List of implicit registers read by this instruction
	RegistersWritten    []uint // List of implicit registers written by this instruction
	Groups              []uint // List of *_GRP_* groups this instruction belongs to.
}

// * INFO
// arch specific information will be filled in for exactly one of the
// substructures. Eg, an Engine created with New(CS_ARCH_ARM, CS_MODE_ARM) will
// fill in only the Arm structure member.
type Instruction struct {
	InstructionHeader
	X86   *X86Instruction
	Arm64 *Arm64Instruction
	Arm   *ArmInstruction
	Mips  *MipsInstruction
	PPC   *PPCInstruction
	Sparc *SparcInstruction
	SysZ  *SysZInstruction
	Xcore *XcoreInstruction
	BPF   *BpfInstruction
	RISCV *RISCVInstruction
	// TODO: Add new arch here
}

// fillGenericHeader Called by the arch specific decomposers
//
// Parameters:
//   - e *Engine  The Engine object, must been already inited
//   - raw C.cs_insn  Raw code instruction
//   - insn *Instruction  Decomposed instruction to go structure
func fillGenericHeader(e *Engine, raw C.cs_insn, insn *Instruction) {

	insn.Id = uint(raw.id)
	insn.Address = uint(raw.address)
	insn.Size = uint(raw.size)

	if !dietMode {
		insn.Mnemonic = C.GoString(&raw.mnemonic[0])
		insn.OpStr = C.GoString(&raw.op_str[0])
	}

	bslice := make([]byte, raw.size)
	for i := 0; i < int(raw.size); i++ {
		bslice[i] = byte(raw.bytes[i])
	}
	insn.Bytes = bslice

	if raw.detail != nil && !dietMode {
		for i := 0; i < int(raw.detail.regs_read_count); i++ {
			insn.RegistersRead = append(insn.RegistersRead, uint(raw.detail.regs_read[i]))
		}

		for i := 0; i < int(raw.detail.regs_write_count); i++ {
			insn.RegistersWritten = append(insn.RegistersWritten, uint(raw.detail.regs_write[i]))
		}

		for i := 0; i < int(raw.detail.groups_count); i++ {
			insn.Groups = append(insn.Groups, uint(raw.detail.groups[i]))
		}

		var regsRead C.cs_regs
		var regsReadCount C.uint8_t
		var regsWrite C.cs_regs
		var regsWriteCount C.uint8_t
		res := C.cs_regs_access(
			e.handle,
			&raw,
			&regsRead[0],
			&regsReadCount,
			&regsWrite[0],
			&regsWriteCount)

		if Errno(res) == ErrOK {
			for i := 0; i < int(regsReadCount); i++ {
				insn.AllRegistersRead = append(insn.AllRegistersRead, uint(regsRead[i]))
			}

			for i := 0; i < int(regsWriteCount); i++ {
				insn.AllRegistersWritten = append(insn.AllRegistersWritten, uint(regsWrite[i]))
			}
		}
	}

}

// Close  Closes the underlying C handle and resources used by this Engine
//
// Returns:
//   - error
func (e *Engine) Close() error {
	res := C.cs_close(&e.handle)
	if e.skipdata != nil {
		C.free(unsafe.Pointer(e.skipdata.mnemonic))
	}
	return Errno(res)
}

// Arch  Accessor for the Engine architecture CS_ARCH_*
//
// Returns:
//   - int  The Engine architecture set
func (e *Engine) Arch() int { return e.arch }

// Mode  Accessor for the Engine mode CS_MODE_*
//
// Returns:
//   - int  The Engine mode set
func (e *Engine) Mode() int { return e.mode }

// Support  Check if a particular arch is supported by this engine.
// To verify if this engine supports everything, use CS_ARCH_ALL
//
// Parameters:
//   - arch int  An architecture id
//
// Returns:
//   - bool  Yes/No
func (e *Engine) Support(arch int) bool { return bool(C.cs_support(C.int(arch))) }

// Version Returns libcapstone version information.
//
// Returns:
//   - maj int  Major part of version
//   - min int  Minor part of version
func (e *Engine) Version() (maj, min int) {
	C.cs_version((*C.int)(unsafe.Pointer(&maj)), (*C.int)(unsafe.Pointer(&min)))
	return
}

// Errno Getter for the last Errno from the engine. Normal code shouldn't need to
// access this directly, but it's exported just in case.
//
// Returns:
//   - error
func (e *Engine) Errno() error { return Errno(C.cs_errno(e.handle)) }

// RegName converts a register number to its mnemonic name
//
// Example: For x86, RegName(0) returns "al", RegName(1) returns "cl"
//
// WARNING: Always returns "" if capstone built with CAPSTONE_DIET
// The arch is implicit in the Engine. Accepts either a constant like ARM_REG_R0
// or insn.Arm.Operands[0].Reg, or anything that refers to a Register like
// insn.X86.SibBase etc
//
// Parameters:
//   - reg uint  A register number
//
// Returns:
//   - string  Its mnemonic name
func (e *Engine) RegName(reg uint) string {

	if dietMode {
		return ""
	}
	return C.GoString(C.cs_reg_name(e.handle, C.uint(reg)))
}

// AI: релаизуй подобные RegNameGeneric фунцкци обёртки для
// 1 для InsnName функции
// 2 для  GroupName функции AI!

// RegNameGeneric RegNameGeneric - wrapper for any integer value.
// With CommonInt type  and convert it to uint.
//
// Parameters:
//   - reg any  Register id.
//
// Returns:
//   - string Register mnemonic name.
func (e *Engine) RegNameGeneric(reg any) string {
	// Приводим тип к uint
	var regUint uint
	switch v := any(reg).(type) {
	case int:
		regUint = uint(v)
	case uint:
		regUint = v
	case int8:
		regUint = uint(v)
	case uint8:
		regUint = uint(v)
	case int16:
		regUint = uint(v)
	case uint16:
		regUint = uint(v)
	case int32:
		regUint = uint(v)
	case uint32:
		regUint = uint(v)
	case int64:
		regUint = uint(v)
	case uint64:
		regUint = uint(v)
	default:
		return ""
	}

	return e.RegName(regUint)
}

// InsnName converts an instruction id to its mnemonic name
// Example: For x86, InsnName(1) returns "add"
//
// The arch is implicit in the Engine. Accepts a constant like
// ARM_INSN_ADD, or insn.Id
//
// WARNING: Always returns "" if capstone built with CAPSTONE_DIET
//
// Parameters:
//   - insn any  An instruction id as any integer type
//
// Returns:
//   - string  Its mnemonic name
func (e *Engine) InsnName(insn any) string {
	if dietMode {
		return ""
	}

	var insnUint uint
	switch v := any(insn).(type) {
	case int:
		insnUint = uint(v)
	case uint:
		insnUint = v
	case int8:
		insnUint = uint(v)
	case uint8:
		insnUint = uint(v)
	case int16:
		insnUint = uint(v)
	case uint16:
		insnUint = uint(v)
	case int32:
		insnUint = uint(v)
	case uint32:
		insnUint = uint(v)
	case int64:
		insnUint = uint(v)
	case uint64:
		insnUint = uint(v)
	default:
		return ""
	}

	return C.GoString(C.cs_insn_name(e.handle, C.uint(insnUint)))
}

// GroupName The arch is implicit in the Engine. Accepts a constant like
// ARM_GRP_JUMP, or insn.Groups[0]
//
// WARNING: Always returns "" if capstone built with CAPSTONE_DIET
//
// Parameters:
//   - grp any  Group of instructions as any integer type
//
// Returns:
//   - string  Name of Group of instructions
func (e *Engine) GroupName(grp any) string {
	if dietMode {
		return ""
	}

	var grpUint uint
	switch v := any(grp).(type) {
	case int:
		grpUint = uint(v)
	case uint:
		grpUint = v
	case int8:
		grpUint = uint(v)
	case uint8:
		grpUint = uint(v)
	case int16:
		grpUint = uint(v)
	case uint16:
		grpUint = uint(v)
	case int32:
		grpUint = uint(v)
	case uint32:
		grpUint = uint(v)
	case int64:
		grpUint = uint(v)
	case uint64:
		grpUint = uint(v)
	default:
		return ""
	}

	return C.GoString(C.cs_group_name(e.handle, C.uint(grpUint)))
}

// SetOption Setter for Engine options CS_OPT_*
//
// Parameters:
//   - ty uint  Engine option type
//   - value uint  Engine options value
//
// Returns:
//   - error
func (e *Engine) SetOption(ty, value uint) error {
	res := C.cs_option(
		e.handle,
		C.cs_opt_type(ty),
		C.size_t(value),
	)

	if Errno(res) == ErrOK {
		return nil
	}
	return Errno(res)
}

// TODO:
// opts := {init..}
// engine.AppyOptions(opts)
// insted of cycle...
func (e *Engine) ApplyOptions(opts []any) error {
	// lo.ForEach(opts, func(x any, _ int) {
	// 	println(x)
	// })
	for _, opt := range opts {
		_ = opt
	}

	return nil
}

// Disasm  Disassemble a []byte full of opcodes.
// Underlying C resources are automatically free'd by this function.
//
// Parameters:
//   - input []byte  Machine instructions in byte codes
//   - address uint64  Address of the first instruction in the given code buffer.
//   - count uint64  Number of instructions to disassemble, 0 to disassemble the whole []byte
//
// Returns:
//   - []Instruction  Disassemled instructions
//   - error
func (e *Engine) Disasm(input []byte, address, count uint64) ([]Instruction, error) {

	var insn *C.cs_insn
	bptr := (*C.uint8_t)(unsafe.Pointer(&input[0]))
	disassembled := C.cs_disasm(
		e.handle,
		bptr,
		C.size_t(len(input)),
		C.uint64_t(address),
		C.size_t(count),
		&insn,
	)

	if disassembled > 0 {
		defer C.cs_free((*C.cs_insn)(unsafe.Pointer(insn)), C.size_t(disassembled))
		// Create a slice, and reflect its header
		var insns []C.cs_insn
		h := (*reflect.SliceHeader)(unsafe.Pointer(&insns))
		// Manually fill in the ptr, len and cap from the raw C data
		h.Data = uintptr(unsafe.Pointer(insn))
		h.Len = int(disassembled)
		h.Cap = int(disassembled)

		switch e.arch {
		case CS_ARCH_ARM:
			return decomposeArm(e, insns), nil

		case CS_ARCH_ARM64:
			return decomposeArm64(e, insns), nil

		case CS_ARCH_MIPS:
			return decomposeMips(e, insns), nil

		case CS_ARCH_X86:
			return decomposeX86(e, insns), nil

		case CS_ARCH_PPC:
			return decomposePPC(e, insns), nil

		case CS_ARCH_SYSZ:
			return decomposeSysZ(e, insns), nil

		case CS_ARCH_SPARC:
			return decomposeSparc(e, insns), nil

		case CS_ARCH_XCORE:
			return decomposeXcore(e, insns), nil

		case CS_ARCH_BPF:
			return decomposeBpf(e, insns), nil

		case CS_ARCH_RISCV:
			return decomposeRiscv(e, insns), nil

			// TODO:
			// CS_ARCH_M68K
			// CS_ARCH_TMS320C64X
			// CS_ARCH_M680X
			// CS_ARCH_EVM
			// CS_ARCH_MOS65XX
			// CS_ARCH_WASM
			// CS_ARCH_MAX
			// CS_ARCH_ALL

		default:
			return decomposeGeneric(e, insns), nil
		}
	}
	return []Instruction{}, e.Errno()
}

// decomposeGeneric Generic decomposer function.
//
// Parameters:
//   - e *Engine  The Engine object
//   - raws []C.cs_insn  Machine instructions
//
// Returns:
//   - []Instruction  Decomposed instructions
func decomposeGeneric(e *Engine, raws []C.cs_insn) []Instruction {
	decomposed := []Instruction{}
	for _, raw := range raws {
		decomp := new(Instruction)
		fillGenericHeader(e, raw, decomp)
		decomposed = append(decomposed, *decomp)
	}
	return decomposed
}

// user callback function prototype
type SkipDataCB func(buffer []byte, offset int, userData interface{}) int

// configuration options for CS_OPT_SKIPDATA, passed via SkipDataStart()
type SkipDataConfig struct {
	Mnemonic string
	Callback SkipDataCB
	UserData interface{}
}

type cbWrapper struct {
	fn SkipDataCB
	ud interface{}
}

// SkipDataStart Enables capstone CS_OPT_SKIPDATA. If no SkipDataConfig is passed ( nil )
// the default behaviour will be enabled. It is valid to pass any combination
// of the SkipDataConfig options, although UserData without a Callback will be
// ignored.
//
// Parameters:
//   - config *SkipDataConfig  Configuration to data skipping.
func (e *Engine) SkipDataStart(config *SkipDataConfig) {

	if config != nil {

		e.skipdata = &C.cs_opt_skipdata{}

		if config.Callback != nil {
			e.skipdata.callback = (C.cs_skipdata_cb_t)(C.trampoline)
			// Happily, we can use the opaque user_data pointer in C to hold both
			// the Go callback function and the Go userData
			e.skipdata.user_data = unsafe.Pointer(
				&cbWrapper{
					fn: config.Callback,
					ud: config.UserData,
				},
			)
		}

		if config.Mnemonic != "" {
			e.skipdata.mnemonic = C.CString(config.Mnemonic)
		} else {
			e.skipdata.mnemonic = C.CString(".byte")
		}

		C.cs_option(e.handle, CS_OPT_SKIPDATA_SETUP, C.size_t(uintptr(unsafe.Pointer(e.skipdata))))
	}

	// If there's no config, just turn on skipdata with the default behaviour
	C.cs_option(e.handle, CS_OPT_SKIPDATA, CS_OPT_ON)
}

// SkipDataStop Disable CS_OPT_SKIPDATA. Removes any registered callbacks and frees
// resources.
func (e *Engine) SkipDataStop() {
	C.cs_option(e.handle, CS_OPT_SKIPDATA, CS_OPT_OFF)
	if e.skipdata == nil {
		return
	}
	C.free(unsafe.Pointer(e.skipdata.mnemonic))
	e.skipdata = nil
}

// New Creates a new Engine with the specified arch and mode CS_ARCH_ , CS_MODE_
//
// Example: CS_ARCH_ARM, CS_MODE_ARM
//
// Parameters:
//   - arch int  Architecture id
//   - mode int  Mode id
//
// Returns:
//   - Engine  Initialized Engine object
//   - error
func New(arch int, mode int) (Engine, error) {
	var handle C.csh
	res := C.cs_open(C.cs_arch(arch), C.cs_mode(mode), &handle)
	if Errno(res) == ErrOK {
		return Engine{handle, arch, mode, nil}, nil
	}
	return Engine{0, CS_ARCH_MAX, 0, nil}, Errno(res)
}
