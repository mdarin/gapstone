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

func sysZInsnDetail(insn Instruction, engine *Engine, buf *bytes.Buffer) {

	if len(insn.SysZ.Operands) > 0 {
		fmt.Fprintf(buf, "\top_count: %v\n", len(insn.SysZ.Operands))
	}

	for i, op := range insn.SysZ.Operands {
		switch op.Type {
		case SYSZ_OP_REG:
			fmt.Fprintf(buf, "\t\toperands[%v].type: REG = %v\n", i, engine.RegName(op.Reg))
		case SYSZ_OP_ACREG:
			fmt.Fprintf(buf, "\t\toperands[%v].type: ACREG = %v\n", i, op.Reg)
		case SYSZ_OP_IMM:
			fmt.Fprintf(buf, "\t\toperands[%v].type: IMM = 0x%x\n", i, (uint64(op.Imm)))
		case SYSZ_OP_MEM:
			fmt.Fprintf(buf, "\t\toperands[%v].type: MEM\n", i)
			if op.Mem.Base != SYSZ_REG_INVALID {
				fmt.Fprintf(buf, "\t\t\toperands[%v].mem.base: REG = %s\n",
					i, engine.RegName(uint(op.Mem.Base)))
			}
			if op.Mem.Index != SYSZ_REG_INVALID {
				fmt.Fprintf(buf, "\t\t\toperands[%v].mem.index: REG = %s\n",
					i, engine.RegName(uint(op.Mem.Index)))
			}
			if op.Mem.Length != 0 {
				fmt.Fprintf(buf, "\t\t\toperands[%v].mem.length: 0x%x\n", i, uint64(op.Mem.Length))
			}
			if op.Mem.Disp != 0 {
				fmt.Fprintf(buf, "\t\t\toperands[%v].mem.disp: 0x%x\n", i, uint64(op.Mem.Disp))
			}
		}

	}

	if insn.SysZ.CC != 0 {
		fmt.Fprintf(buf, "\tCode condition: %v\n", insn.SysZ.CC)
	}

	fmt.Fprintf(buf, "\n")
}

func TestSysZ(t *testing.T) {

	t.Parallel()

	final := new(bytes.Buffer)
	spec_file := SystemzSpec
	for i, platform := range sysZTests {

		engine, err := New(platform.arch, platform.mode)
		if err != nil {
			t.Errorf("Failed to initialize engine %v", err)
			return
		}
		defer engine.Close()

		for _, opt := range platform.options {
			engine.SetOption(opt.ty, opt.value)
		}
		if i == 0 {
			maj, min := engine.Version()
			t.Logf("Arch: SystemZ. Capstone Version: %v.%v", maj, min)
			check := checks[CS_ARCH_SYSZ]
			if check.grpMax != SYSZ_GRP_ENDING ||
				check.insMax != SYSZ_INS_ENDING ||
				check.regMax != SYSZ_REG_ENDING {
				t.Errorf("Failed in sanity check. Constants out of sync with core.")
			} else {
				t.Logf("Sanity Check: PASS")
			}
		}

		insns, err := engine.Disasm([]byte(platform.code), address, 0)
		if err == nil {
			fmt.Fprintf(final, "****************\n")
			fmt.Fprintf(final, "Platform: %s\n", platform.comment)
			fmt.Fprintf(final, "Code:")
			dumpHex([]byte(platform.code), final)
			fmt.Fprintf(final, "Disasm:\n")
			for _, insn := range insns {
				fmt.Fprintf(final, "0x%x:\t%s\t%s\n", insn.Address, insn.Mnemonic, insn.OpStr)
				sysZInsnDetail(insn, &engine, final)
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
	final.Reset()

}
