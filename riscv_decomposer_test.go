package gapstone

import (
	"testing"
)

func TestRiscvDecomposition(t *testing.T) {
	e := NewEngine(CS_ARCH_RISCV, CS_MODE_RISCV64)
	defer e.Close()

	if err := e.SetSkipData(nil); err != nil {
		t.Fatalf("SetSkipData failed: %v", err)
	}

	// Test RISC-V code from riscv_codes
	for _, code := range []string{riscvCodes["RISCV_CODE64"]} {
		instructions, err := e.Disasm(code, 0x1000, 4)
		if err != nil {
			t.Fatalf("Disasm failed: %v", err)
		}

		for _, insn := range instructions {
			if len(insn.Riscv.Operands) == 0 {
				t.Error("Expected operands but got none")
			}
		}
	}
}
