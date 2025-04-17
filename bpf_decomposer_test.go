/*
Gapstone is a Go binding for the Capstone disassembly library. For examples,
try reading the *_test.go files.

	Library Author: Nguyen Anh Quynh
	Binding Author: Ben Nagy
	License: BSD style - see LICENSE file for details
    (c) 2013 COSEINC. All Rights Reserved.
*/

package gapstone

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"testing"
)

func bpfInsnDetail(insn Instruction, engine *Engine, buf *bytes.Buffer) {
	if oplen := len(insn.Bpf.Operands); oplen > 0 {
		fmt.Fprintf(buf, "\top_count: %v\n", oplen)
	}

	for i, op := range insn.Bpf.Operands {
		switch op.Type {
		case BPF_OP_REG:
			fmt.Fprintf(buf, "\t\toperands[%v].type: REG = %v\n", i, engine.RegName(op.Reg))
		case BPF_OP_IMM:
			fmt.Fprintf(buf, "\t\toperands[%v].type: IMM = 0x%x\n", i, (uint64(op.Immediate)))
		case BPF_OP_FP:
			fmt.Fprintf(buf, "\t\toperands[%v].type: FP = %f\n", i, op.FP)
		case BPF_OP_MEM:
			fmt.Fprintf(buf, "\t\toperands[%v].type: MEM\n", i)
			if op.Mem.Base != BPF_REG_INVALID {
				fmt.Fprintf(buf, "\t\t\toperands[%v].mem.base: REG = %s\n",
					i, engine.RegName(op.Mem.Base))
			}
			if op.Mem.Disp != 0 {
				fmt.Fprintf(buf, "\t\t\toperands[%v].mem.disp: 0x%x\n", i, uint64(op.Mem.Disp))
			}
		case BPF_OP_CIMM:
			fmt.Fprintf(buf, "\t\toperands[%v].type: C-IMM = %v\n", i, op.Immediate)
		case BPF_OP_REG_MRS:
			fmt.Fprintf(buf, "\t\toperands[%v].type: REG_MRS = 0x%x\n", i, op.Reg)
		case BPF_OP_REG_MSR:
			fmt.Fprintf(buf, "\t\toperands[%v].type: REG_MSR = 0x%x\n", i, op.Reg)
		case BPF_OP_PSTATE:
			fmt.Fprintf(buf, "\t\toperands[%v].type: PSTATE = 0x%x\n", i, op.PState)
		case BPF_OP_SYS:
			fmt.Fprintf(buf, "\t\toperands[%v].type: SYS = 0x%x\n", i, op.Sys)
		case BPF_OP_PREFETCH:
			fmt.Fprintf(buf, "\t\toperands[%v].type: PREFETCH = 0x%x\n", i, op.Prefetch)
		case BPF_OP_BARRIER:
			fmt.Fprintf(buf, "\t\toperands[%v].type: BARRIER = 0x%x\n", i, op.Barrier)
		}

		switch op.Access {
		case CS_AC_READ:
			fmt.Fprintf(buf, "\t\toperands[%v].access: READ\n", i)
		case CS_AC_WRITE:
			fmt.Fprintf(buf, "\t\toperands[%v].access: WRITE\n", i)
		case CS_AC_READ | CS_AC_WRITE:
			fmt.Fprintf(buf, "\t\toperands[%v].access: READ | WRITE\n", i)
		}

		if op.Ext != BPF_EXT_INVALID {
			fmt.Fprintf(buf, "\t\t\tExt: %v\n", op.Ext)
		}
	}

	if len(insn.AllRegistersRead) > 0 {
		fmt.Fprintf(buf, "\tRegisters read:")
		for _, reg := range insn.AllRegistersRead {
			fmt.Fprintf(buf, " %s", engine.RegName(reg))
		}
		fmt.Fprintf(buf, "\n")
	}

	if len(insn.AllRegistersWritten) > 0 {
		fmt.Fprintf(buf, "\tRegisters modified:")
		for _, reg := range insn.AllRegistersWritten {
			fmt.Fprintf(buf, " %s", engine.RegName(reg))
		}
		fmt.Fprintf(buf, "\n")
	}

	fmt.Fprintf(buf, "\n")
}

func TestBpf(t *testing.T) {

	t.Parallel()

	final := new(bytes.Buffer)
	spec_file := BPFSpec

	for i, platform := range bpfTests {

		engine, err := New(platform.arch, platform.mode)
		if err != nil {
			t.Errorf("Failed to initialize engine %v", err)
			return
		}
		for _, opt := range platform.options {
			engine.SetOption(opt.ty, opt.value)
		}
		if i == 0 {
			maj, min := engine.Version()
			t.Logf("Arch: BPF. Capstone Version: %v.%v", maj, min)
			check := checks[CS_ARCH_BPF]
			if check.grpMax != BPF_GRP_ENDING ||
				check.insMax != BPF_INS_ENDING ||
				check.regMax != BPF_REG_ENDING {
				t.Errorf("Failed in sanity check. Constants out of sync with core.")
			} else {
				t.Logf("Sanity Check: PASS")
			}
		}
		defer engine.Close()

		insns, err := engine.Disasm([]byte(platform.code), 0x2c, 0)
		if err == nil {
			fmt.Fprintf(final, "****************\n")
			fmt.Fprintf(final, "Platform: %s\n", platform.comment)
			fmt.Fprintf(final, "Code: ")
			dumpHex([]byte(platform.code), final)
			fmt.Fprintf(final, "Disasm:\n")
			for _, insn := range insns {
				fmt.Fprintf(final, "0x%x:\t%s\t%s\n", insn.Address, insn.Mnemonic, insn.OpStr)
				bpfInsnDetail(insn, &engine, final)
			}
			fmt.Fprintf(final, "0x%x:\n", insns[len(insns)-1].Address+insns[len(insns)-1].Size)
			fmt.Fprintf(final, "\n")
		} else {
			t.Errorf("Disassembly error: %v\n", err)
		}

	}

	spec, err := ioutil.ReadFile(spec_file)
	if err != nil {
		t.Errorf("Cannot read spec file %v: %v", spec_file, err)
	}
	if fs := final.String(); string(spec) != fs {
		// fmt.Println(fs)
		t.Errorf("Output failed to match spec!")
	} else {
		t.Logf("Clean diff with %v.\n", spec_file)
	}

}
