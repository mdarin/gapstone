gapstone
====

![logo](assets/logo.png)

Gapstone is a Go binding for the Capstone disassembly library.

Clone capstone or run install_capstone.sh script
https://github.com/capstone-engine/capstone.git
(from next branch by default)
and run scripts step-by-step
```sh
cd path/to/capstone
make 
sudo make install
make clean
cd back/to/gapstone

install ruby
# or  perpare builder image
docker build -t gapstone-builder .

## generate

export RUN="docker run -t --rm -v $(pwd):/build -w /build gapstone-builder"

./genconst path/to/clonned/capstone/bindings/python/capstone
# or in docker builder
docker run -t --rm -v $(pwd):/build -w /build gapstone-builder ./genconst.rb capstone/bindings/python/capstone/

./genspec path/to/clonned/capstone/tests
# or in docker builder
docker run -t --rm -v $(pwd):/build -w /build  gapstone-builder ./genspec.rb capstone/tests/

eport EXCLUDEDIR="gapstone"
EXCLUDEDIR="gapstone" go test -skip-dirs="$GAPSTONE" ./...

# number of times
# try to run go test ./...
# If you see someting like this on run the command go test ./...
# ./riscv_constants.go:471:22: could not determine kind of name for C.RISCV_GRP_ISRV32C
# ./riscv_constants.go:472:23: could not determine kind of name for C.RISCV_GRP_ISRV32CF
# ./riscv_constants.go:474:22: could not determine kind of name for C.RISCV_GRP_ISRV64A
# ./riscv_constants.go:475:22: could not determine kind of name for C.RISCV_GRP_ISRV64C
# ./riscv_constants.go:476:22: could not determine kind of name for C.RISCV_GRP_ISRV64D
# try to fix it by this script
 ./gencomment_lines.sh 

./genresources.sh 

# fix junk after generation
# and try to run tests
# repeat until done :)

 go test ./...
 # all tests ok 



# Для одного файла
ruby parser.rb test.c

# Для всей директории (рекурсивно)
ruby parser.rb /path/to/code/
DEBUG=1 ruby parser.rb path/to/file.c
```

You should see generated .go source files.
Use it for binding in your project.

## CURRENT UPSTREAM VERSION: 4.0.2
[![Build Status](https://travis-ci.org/knightsc/gapstone.svg?branch=master)](https://travis-ci.org/knightsc/gapstone)

(head over to the next branch for the newest stuff)

SUMMARY
===

( FROM THE CAPSTONE README )

Capstone is a disassembly framework with the target of becoming the ultimate
disasm engine for binary analysis and reversing in the security community.

Created by Nguyen Anh Quynh, then developed and maintained by a small community,
Capstone offers some unparalleled features:

- Support multiple hardware architectures: ARM, ARM64 (ARMv8), Mips, PPC, Sparc,
  SystemZ, XCore and X86.

- Having clean/simple/lightweight/intuitive architecture-neutral API.

- Provide details on disassembled instruction (called “decomposer” by others).

- Provide semantics of the disassembled instruction, such as list of implicit
  registers read & written.

- Implemented in pure C language, with lightweight wrappers for C++, C#, Go,
  Java, NodeJS, Ocaml, Python, Ruby & Vala ready (available in main code,
  or provided externally by the community).

- Native support for all popular platforms: Windows, Mac OSX, iOS, Android,
  Linux, *BSD, Solaris, etc.

- Thread-safe by design.

- Special support for embedding into firmware or OS kernel.

- Distributed under the open source BSD license.

Further information is available at http://www.capstone-engine.org

To install:
----

First install the capstone library from either https://github.com/aquynh/capstone
or http://www.capstone-engine.org

Then, assuming you have set up your Go environment according to the docs, just:
```bash
go get -u github.com/knightsc/gapstone
```

Tests are provided. You should probably run them.
```
cd $GOPATH/src/github.com/knightsc/gapstone
go test
```

To start writing code:
----

Take a look at the examples *_test.go

Here's "Hello World":
```go
package main

import (
    "github.com/knightsc/gapstone"
    "log"
)

func main() {

    engine, err := gapstone.New(
        gapstone.CS_ARCH_X86,
        gapstone.CS_MODE_32,
    )

    if err == nil {

        defer engine.Close()

        maj, min := engine.Version()
        log.Printf("Hello Capstone! Version: %v.%v\n", maj, min)

        var x86Code32 = "\x8d\x4c\x32\x08\x01\xd8\x81\xc6\x34" +
            "\x12\x00\x00\x05\x23\x01\x00\x00\x36\x8b\x84\x91" +
            "\x23\x01\x00\x00\x41\x8d\x84\x39\x89\x67\x00\x00" +
            "\x8d\x87\x89\x67\x00\x00\xb4\xc6"

        insns, err := engine.Disasm(
            []byte(x86Code32), // code buffer
            0x10000,           // starting address
            0,                 // insns to disassemble, 0 for all
        )

        if err == nil {
            log.Printf("Disasm:\n")
            for _, insn := range insns {
                log.Printf("0x%x:\t%s\t\t%s\n", insn.Address, insn.Mnemonic, insn.OpStr)
            }
            return
        }
        log.Fatalf("Disassembly error: %v", err)
    }
    log.Fatalf("Failed to initialize engine: %v", err)
}
```

Autodoc is available at http://godoc.org/github.com/knightsc/gapstone

Contributing
----

If you feel like chipping in, especially with better tests or examples, fork and send me a pull req.
