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
	"bytes"
	"fmt"
	"os"
	"strings"
	"testing"
)

func bpfInsnDetail(insn Instruction, engine *Engine, buf *bytes.Buffer) {
	if oplen := len(insn.BPF.Operands); oplen > 0 {
		fmt.Fprintf(buf, "\tOperand count: %v\n", oplen)
	}

	for i, op := range insn.BPF.Operands {
		switch op.Type {
		case BPF_OP_REG:
			fmt.Fprintf(buf, "\t\toperands[%v].type: REG = %v\n", i, engine.RegName(op.Reg))

		case BPF_OP_IMM:
			fmt.Fprintf(buf, "\t\toperands[%v].type: IMM = 0x%x\n", i, op.Imm)

		case BPF_OP_OFF:
			fmt.Fprintf(buf, "\t\toperands[%v].type: OFF = +0x%x\n", i, op.Off)

		case BPF_OP_MEM:
			fmt.Fprintf(buf, "\t\toperands[%v].type: MEM\n", i)
			if op.Mem.Base != BPF_REG_INVALID {
				fmt.Fprintf(buf, "\t\t\toperands[%v].mem.base: REG = %s\n",
					i, engine.RegName(op.Mem.Base))
			}
			if op.Mem.Disp != 0 {
				fmt.Fprintf(buf, "\t\t\toperands[%v].mem.disp: 0x%x\n", i, uint64(op.Mem.Disp))
			}

		case BPF_OP_MMEM:
			fmt.Fprintf(buf, "MMEM = M[0x%x]\n", op.Mmem)

		case BPF_OP_MSH:
			fmt.Fprintf(buf, "MSH = 4*([0x%x]&0xf)\n", op.Msh)

		case BPF_OP_EXT:
			fmt.Fprintf(buf, "EXT = %s\n", i, "todo: from BPF_EXT_LEN")
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
	spec_file := BpfSpec

	for i, platform := range bpfPlatforms {

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
				bpfInsnDetail(insn, &engine, final)
			}
			// ? unnes output. Uncomment if needed
			// fmt.Fprintf(final, "0x%x:\n", insns[len(insns)-1].Address+insns[len(insns)-1].Size)
			fmt.Fprintf(final, "\n")
		} else {
			t.Errorf("Disassembly error: %v\n", err)
		}

	}

	spec, err := os.ReadFile(spec_file)
	if err != nil {
		t.Errorf("Cannot read spec file %v: %v", spec_file, err)
	}
	if fs := final.String(); string(spec) != fs {
		// INFO: Ucomment for debugging, diff this output with SPEC
		fmt.Println(fs)
		t.Errorf("Output failed to match spec!")
	} else {
		t.Logf("Clean diff with %v.\n", spec_file)
	}
}

// Функция преобразования константы в строку
func bpfGroupToString(group int) string {
	// Преимущества варианта (с map):

	// Легче поддерживать при добавлении новых констант
	// Более компактный код
	// Проще автоматически генерировать
	// если потребуется спефические действия можено реализовать
	// в связанных функциях доступных по ключу

	// Создаем карту соответствия констант их строковым представлениям
	groupNames := map[int]string{
		BPF_GRP_INVALID: "BPF_GRP_INVALID", // + func (args...) { actions ...}
		BPF_GRP_LOAD:    "BPF_GRP_LOAD",
		BPF_GRP_STORE:   "BPF_GRP_STORE",
		BPF_GRP_ALU:     "BPF_GRP_ALU",
		BPF_GRP_JUMP:    "BPF_GRP_JUMP",
		BPF_GRP_CALL:    "BPF_GRP_CALL",
		BPF_GRP_RETURN:  "BPF_GRP_RETURN",
		BPF_GRP_MISC:    "BPF_GRP_MISC",
		BPF_GRP_ENDING:  "BPF_GRP_ENDING",
	}

	// Получаем строковое представление константы
	fullName, ok := groupNames[group]
	if !ok {
		return "unknown"
	}

	// Отбрасываем префикс "BPF_GRP_" и приводим к нижнему регистру
	return strings.ToLower(strings.TrimPrefix(fullName, "BPF_GRP_"))
}

// AI: напиши функцию подобную  bpfGroupToString
// Функция преобразования константы в строку
// func bpfExtToString(group int) string
// для значений
// BPF_EXT_INVALID
// 	BPF_EXT_LEN
// AI!
