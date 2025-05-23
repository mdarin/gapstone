/*
Gapstone is a Go binding for the Capstone disassembly library. For examples,
try reading the *_test.go files.

	Library Author: Nguyen Anh Quynh
	Binding Author: Ben Nagy
	License: BSD style - see LICENSE file for details
    (c) 2013 COSEINC. All Rights Reserved.
*/

// ! DEPRECATED

package gapstone

import (
	"bytes"
	"fmt"
	"os"
	"testing"

	"github.com/avito-tech/normalize"
)

func TestTest(t *testing.T) {

	t.SkipNow()

	// ! DEPRECATED

	final := new(bytes.Buffer)
	spec_file := BasicSpec
	var maj, min int
	if ver, err := New(0, 0); err == nil {
		maj, min = ver.Version()
		ver.Close()
	}

	t.Logf("Basic Test. Capstone Version: %v.%v", maj, min)

	for i, platform := range basicPlatforms {

		t.Logf("%2d> %s", i, platform.comment)
		engine, err := New(platform.arch, platform.mode)
		if err != nil {
			t.Errorf("Failed to initialize engine %v", err)
			return
		}

		defer engine.Close()

		for _, opt := range platform.options {
			engine.SetOption(opt.ty, opt.value)
		}

		insns, err := engine.Disasm([]byte(platform.code), address, 0)
		if err == nil {
			fmt.Fprintf(final, "****************\n")
			fmt.Fprintf(final, "Platform: %s\n", platform.comment)
			fmt.Fprintf(final, "Code: ")
			dumpHex([]byte(platform.code), final)
			fmt.Fprintf(final, "Disasm:\n")
			for _, insn := range insns {
				fmt.Fprintf(final, "0x%x:\t%s\t\t%s\n", insn.Address, insn.Mnemonic, insn.OpStr)
			}
			fmt.Fprintf(final, "0x%x:\n", insns[len(insns)-1].Address+insns[len(insns)-1].Size)
			fmt.Fprintf(final, "\n")
		} else {
			t.Errorf("Disassembly error: %v\n", err)
		}

	}

	spec, err := os.ReadFile(spec_file)
	if err != nil {
		t.Errorf("Cannot read spec file %v: %v", spec_file, err)
	}

	// * INFO:
	// Compare two strings with all normalizations described above applied.
	// Provide threshold parameters to tweak how similar strings must be to make the function return true.
	// threshold is relative value, so 0.5 roughly means "strings are 50% different after all normalizations applied".
	similarityThreshold := 0.05
	isSimilar := normalize.AreStringsSimilar(string(spec), final.String(), similarityThreshold)
	t.Logf("String similariti Levenstein distance: %f isSimilar: %t", similarityThreshold, isSimilar)

	if !isSimilar {
		// * Debugging - uncomment below and run the test | diff - test.SPEC
		// fmt.Println(final.String())
		t.Errorf("Output failed to match spec!")
	} else {
		t.Logf("Clean diff with %v.\n", spec_file)
	}
}
