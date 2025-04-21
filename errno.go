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
// extern size_t trampoline(uint8_t *buffer, size_t buflen, size_t offset, void *user_data);
import "C"

import (
	"fmt"
)

type Errno int

func (e Errno) Error() string {
	s := C.GoString(C.cs_strerror(C.cs_err(e)))
	if s == "" {
		return fmt.Sprintf("Internal Error: No Error string for Errno %d", int(e))
	}
	return s
}

var (
	ErrOK       = Errno(0)  // No error: everything was fine
	ErrMem      = Errno(1)  // Out-Of-Memory error: cs_open(), cs_disasm()
	ErrArch     = Errno(2)  // Unsupported architecture: cs_open()
	ErrHandle   = Errno(3)  // Invalid handle: cs_op_count(), cs_op_index()
	ErrCsh      = Errno(4)  // Invalid csh argument: cs_close(), cs_errno(), cs_option()
	ErrMode     = Errno(5)  // Invalid/unsupported mode: cs_open()
	ErrOption   = Errno(6)  // Invalid/unsupported option: cs_option()
	ErrDetail   = Errno(7)  // Information is unavailable because detail option is OFF
	ErrMemSetup = Errno(8)  // Dynamic memory management uninitialized (see CS_OPT_MEM)
	ErrVersion  = Errno(9)  // Unsupported version (bindings)
	ErrDiet     = Errno(10) // Access irrelevant data in "diet" engine
	ErrSkipdata = Errno(11) // Access irrelevant data for "data" instruction in SKIPDATA mode
	ErrX86ATT   = Errno(12) // X86 AT&T syntax is unsupported (opt-out at compile time)
	ErrX86Intel = Errno(13) // X86 Intel syntax is unsupported (opt-out at compile time)
)
