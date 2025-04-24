package gapstone

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/avito-tech/normalize"
)

func riscvInsnDetail(insn Instruction, engine *Engine, buf *bytes.Buffer) {
	fmt.Fprintf(buf, "\tGroups:")
	for _, group := range insn.Groups {
		fmt.Fprintf(buf, " %s", engine.GroupName(group))
	}
	fmt.Fprintf(buf, "\n")

	if oplen := len(insn.Riscv.Operands); oplen >= 0 {
		fmt.Fprintf(buf, "\tOperand count: %v\n", oplen)
	}

	for i, op := range insn.Riscv.Operands {
		switch op.Type {
		case RISCV_OP_IMM:
			fmt.Fprintf(buf, "\t\toperands[%v].type: IMM = 0x%x\n", i, op.Imm)
		case RISCV_OP_REG:
			fmt.Fprintf(buf, "\t\toperands[%v].type: REG = %s\n", i, engine.RegName(op.Reg))
		case RISCV_OP_MEM:
			fmt.Fprintf(buf, "\t\toperands[%v].type: MEM\n", i)
			if op.Mem.Base != RISCV_REG_INVALID {
				fmt.Fprintf(buf, "\t\t\toperands[%v].mem.base: REG = %s\n", i, engine.RegName(op.Mem.Base))
			}
			fmt.Fprintf(buf, "\t\t\toperands[%v].mem.disp: 0x%x\n", i, uint64(op.Mem.Disp))
		}

		if op.Access != 0 {
			fmt.Fprintf(buf, "\t\t\tAccess type: %v\n", op.Access)
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

func TestRiscv(t *testing.T) {
	t.Parallel()

	final := new(bytes.Buffer)
	spec_file := "riscv_spec.txt"

	for i, platform := range riscvPlatforms {
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
			t.Logf("Arch: RISC-V. Capstone Version: %v.%v", maj, min)
			check := checks[CS_ARCH_RISCV]
			if check.grpMax != RISCV_GRP_ENDING ||
				check.insMax != RISCV_INS_ENDING ||
				check.regMax != RISCV_REG_ENDING {
				t.Errorf("Failed in sanity check. Constants out of sync with core.")
			} else {
				t.Logf("Sanity Check: PASS")
			}
		}
		defer engine.Close()

		input := []byte(platform.code)
		address := uint64(0x00)
		count := uint64(0)

		insns, err := engine.Disasm(input, address, count)
		if err == nil {
			fmt.Fprintf(final, "****************\n")
			fmt.Fprintf(final, "Platform: %s\n", platform.comment)
			fmt.Fprintf(final, "Code: ")
			dumpHex([]byte(platform.code), final)
			fmt.Fprintf(final, "Disasm:\n")
			for _, insn := range insns {
				fmt.Fprintf(final, "0x%x:\t%s\t%s\n", insn.Address, insn.Mnemonic, insn.OpStr)
				riscvInsnDetail(insn, &engine, final)
			}
		} else {
			t.Errorf("Disassembly error: %v\n", err)
		}

		spec, err := os.ReadFile(spec_file)
		if err != nil {
			t.Errorf("Cannot read spec file %v: %v", spec_file, err)
		}

		similarityThreshold := 0.01
		isSimilar := normalize.AreStringsSimilar(string(spec), final.String(), similarityThreshold)
		t.Logf("String similariti Levenstein distance: %f isSimilar: %t", similarityThreshold, isSimilar)

		if !CompareNormalized(final.String(), spec) {
			fmt.Println(final.String())
			t.Errorf("Output failed to match spec!")
		} else {
			t.Logf("Clean diff with %v.\n", spec_file)
		}
	}
}

func dumpHex(data []byte, buf *bytes.Buffer) {
	for i, b := range data {
		if i > 0 && i%8 == 0 {
			fmt.Fprintf(buf, " ")
		}
		fmt.Fprintf(buf, "%02x ", b)
	}
	fmt.Fprintf(buf, "\n")
}
