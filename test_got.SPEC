****************
Platform: X86 16bit (Intel syntax)
Code: 0x8d 0x4c 0x32 0x08 0x01 0xd8 0x81 0xc6 0x34 0x12 0x00 0x00 
Disasm:
0x1000: lea             cx, [si + 0x32]
0x1003: or              byte ptr [bx + di], al
0x1005: fadd            dword ptr [bx + di + 0x34c6]
0x1009: adc             al, byte ptr [bx + si]
0x100b:

****************
Platform: X86 32bit (ATT syntax)
Code: 0xba 0xcd 0xab 0x00 0x00 0x8d 0x4c 0x32 0x08 0x01 0xd8 0x81 0xc6 0x34 0x12 0x00 0x00 
Disasm:
0x1000: movl            $0xabcd, %edx
0x1005: leal            8(%edx, %esi), %ecx
0x1009: addl            %ebx, %eax
0x100b: addl            $0x1234, %esi
0x1011:

****************
Platform: X86 32 (Intel syntax)
Code: 0xba 0xcd 0xab 0x00 0x00 0x8d 0x4c 0x32 0x08 0x01 0xd8 0x81 0xc6 0x34 0x12 0x00 0x00 
Disasm:
0x1000: mov             edx, 0xabcd
0x1005: lea             ecx, [edx + esi + 8]
0x1009: add             eax, ebx
0x100b: add             esi, 0x1234
0x1011:

****************
Platform: X86 32 (MASM syntax)
Code: 0xba 0xcd 0xab 0x00 0x00 0x8d 0x4c 0x32 0x08 0x01 0xd8 0x81 0xc6 0x34 0x12 0x00 0x00 
Disasm:
0x1000: mov             edx, 0abcdh
0x1005: lea             ecx, [edx + esi + 8]
0x1009: add             eax, ebx
0x100b: add             esi, 1234h
0x1011:

****************
Platform: X86 64 (Intel syntax)
Code: 0x55 0x48 0x8b 0x05 0xb8 0x13 0x00 0x00 
Disasm:
0x1000: push            rbp
0x1001: mov             rax, qword ptr [rip + 0x13b8]
0x1008:

****************
Platform: ARM
Code: 0xed 0xff 0xff 0xeb 0x04 0xe0 0x2d 0xe5 0x00 0x00 0x00 0x00 0xe0 0x83 0x22 0xe5 0xf1 0x02 0x03 0x0e 0x00 0x00 0xa0 0xe3 0x02 0x30 0xc1 0xe7 0x00 0x00 0x53 0xe3 
Disasm:
0x1000: bl              #0xfbc
0x1004: str             lr, [sp, #-4]!
0x1008: andeq           r0, r0, r0
0x100c: str             r8, [r2, #-0x3e0]!
0x1010: mcreq           p2, #0, r0, c3, c1, #7
0x1014: mov             r0, #0
0x1018: strb            r3, [r1, r2]
0x101c: cmp             r3, #0
0x1020:

****************
Platform: THUMB-2
Code: 0x4f 0xf0 0x00 0x01 0xbd 0xe8 0x00 0x88 0xd1 0xe8 0x00 0xf0 
Disasm:
0x1000: mov.w           r1, #0
0x1004: pop.w           {fp, pc}
0x1008: tbb             [r1, r0]
0x100c:

****************
Platform: ARM: Cortex-A15 + NEON
Code: 0x10 0xf1 0x10 0xe7 0x11 0xf2 0x31 0xe7 0xdc 0xa1 0x2e 0xf3 0xe8 0x4e 0x62 0xf3 
Disasm:
0x1000: sdiv            r0, r0, r1
0x1004: udiv            r1, r1, r2
0x1008: vbit            q5, q15, q6
0x100c: vcgt.f32                q10, q9, q12
0x1010:

****************
Platform: THUMB
Code: 0x70 0x47 0xeb 0x46 0x83 0xb0 0xc9 0x68 
Disasm:
0x1000: bx              lr
0x1002: mov             fp, sp
0x1004: sub             sp, #0xc
0x1006: ldr             r1, [r1, #0xc]
0x1008:

****************
Platform: Thumb-MClass
Code: 0xef 0xf3 0x02 0x80 
Disasm:
0x1000: mrs             r0, eapsr
0x1004:

****************
Platform: Arm-V8
Code: 0xe0 0x3b 0xb2 0xee 0x42 0x00 0x01 0xe1 0x51 0xf0 0x7f 0xf5 
Disasm:
0x1000: vcvtt.f64.f16           d3, s1
0x1004: crc32b          r0, r1, r2
0x1008: dmb             oshld
0x100c:

****************
Platform: MIPS-32 (Big-endian)
Code: 0x0c 0x10 0x00 0x97 0x00 0x00 0x00 0x00 0x24 0x02 0x00 0x0c 0x8f 0xa2 0x00 0x00 0x34 0x21 0x34 0x56 
Disasm:
0x1000: jal             0x40025c
0x1004: nop
0x1008: addiu           $v0, $zero, 0xc
0x100c: lw              $v0, ($sp)
0x1010: ori             $at, $at, 0x3456
0x1014:

****************
Platform: MIPS-64-EL (Little-endian)
Code: 0x56 0x34 0x21 0x34 0xc2 0x17 0x01 0x00 
Disasm:
0x1000: ori             $at, $at, 0x3456
0x1004: srl             $v0, $at, 0x1f
0x1008:

****************
Platform: MIPS-32R6 | Micro (Big-endian)
Code: 0x00 0x07 0x00 0x07 0x00 0x11 0x93 0x7c 0x01 0x8c 0x8b 0x7c 0x00 0xc7 0x48 0xd0 
Disasm:
0x1000: break           7, 0
0x1004: wait            0x11
0x1008: syscall         0x18c
0x100c: rotrv           $t1, $a2, $a3
0x1010:

****************
Platform: MIPS-32R6 (Big-endian)
Code: 0xec 0x80 0x00 0x19 0x7c 0x43 0x22 0xa0 
Disasm:
0x1000: addiupc         $a0, 0x64
0x1004: align           $a0, $v0, $v1, 2
0x1008:

****************
Platform: ARM-64
Code: 0x21 0x7c 0x02 0x9b 0x21 0x7c 0x00 0x53 0x00 0x40 0x21 0x4b 0xe1 0x0b 0x40 0xb9 
Disasm:
0x1000: mul             x1, x1, x2
0x1004: lsr             w1, w1, #0
0x1008: sub             w0, w0, w1, uxtw
0x100c: ldr             w1, [sp, #8]
0x1010:

****************
Platform: PPC-64
Code: 0x80 0x20 0x00 0x00 0x80 0x3f 0x00 0x00 0x10 0x43 0x23 0x0e 0xd0 0x44 0x00 0x80 0x4c 0x43 0x22 0x02 0x2d 0x03 0x00 0x80 0x7c 0x43 0x20 0x14 0x7c 0x43 0x20 0x93 0x4f 0x20 0x00 0x21 0x4c 0xc8 0x00 0x21 
Disasm:
0x1000: lwz             r1, 0(0)
0x1004: lwz             r1, 0(r31)
0x1008: vpkpx           v2, v3, v4
0x100c: stfs            f2, 0x80(r4)
0x1010: crand           cr0eq, cr0un, cr1lt
0x1014: cmpwi           cr2, r3, 0x80
0x1018: addc            r2, r3, r4
0x101c: mulhd.          r2, r3, r4
0x1020: bdnzlrl+
0x1024: bgelrl-         cr2
0x1028:

****************
Platform: PPC-64
Code: 0x80 0x20 0x00 0x00 0x80 0x3f 0x00 0x00 0x10 0x43 0x23 0x0e 0xd0 0x44 0x00 0x80 0x4c 0x43 0x22 0x02 0x2d 0x03 0x00 0x80 0x7c 0x43 0x20 0x14 0x7c 0x43 0x20 0x93 0x4f 0x20 0x00 0x21 0x4c 0xc8 0x00 0x21 
Disasm:
0x1000: lwz             1, 0(0)
0x1004: lwz             1, 0(31)
0x1008: vpkpx           2, 3, 4
0x100c: stfs            2, 0x80(4)
0x1010: crand           0, 0, 1
0x1014: cmpwi           2, 3, 0x80
0x1018: addc            2, 3, 4
0x101c: mulhd.          2, 3, 4
0x1020: bdnzlrl+
0x1024: bgelrl-         cr2
0x1028:

****************
Platform: PPC-64 + QPX
Code: 0x10 0x60 0x2a 0x10 0x10 0x64 0x28 0x88 0x7c 0x4a 0x5d 0x0f 
Disasm:
0x1000: qvfabs          q3, q5
0x1004: qvfand          q3, q4, q5
0x1008: qvstfsxa                q2, r10, r11
0x100c:

****************
Platform: Sparc
Code: 0x80 0xa0 0x40 0x02 0x85 0xc2 0x60 0x08 0x85 0xe8 0x20 0x01 0x81 0xe8 0x00 0x00 0x90 0x10 0x20 0x01 0xd5 0xf6 0x10 0x16 0x21 0x00 0x00 0x0a 0x86 0x00 0x40 0x02 0x01 0x00 0x00 0x00 0x12 0xbf 0xff 0xff 0x10 0xbf 0xff 0xff 0xa0 0x02 0x00 0x09 0x0d 0xbf 0xff 0xff 0xd4 0x20 0x60 0x00 0xd4 0x4e 0x00 0x16 0x2a 0xc2 0x80 0x03 
Disasm:
0x1000: cmp             %g1, %g2
0x1004: jmpl            %o1+8, %g2
0x1008: restore         %g0, 1, %g2
0x100c: restore
0x1010: mov             1, %o0
0x1014: casx            [%i0], %l6, %o2
0x1018: sethi           0xa, %l0
0x101c: add             %g1, %g2, %g3
0x1020: nop
0x1024: bne             0x1020
0x1028: ba              0x1024
0x102c: add             %o0, %o1, %l0
0x1030: fbg             0x102c
0x1034: st              %o2, [%g1]
0x1038: ldsb            [%i0+%l6], %o2
0x103c: brnz,a,pn               %o2, 0x1048
0x1040:

****************
Platform: SparcV9
Code: 0x81 0xa8 0x0a 0x24 0x89 0xa0 0x10 0x20 0x89 0xa0 0x1a 0x60 0x89 0xa0 0x00 0xe0 
Disasm:
0x1000: fcmps           %f0, %f4
0x1004: fstox           %f0, %f4
0x1008: fqtoi           %f0, %f4
0x100c: fnegq           %f0, %f4
0x1010:

****************
Platform: SystemZ
Code: 0xed 0x00 0x00 0x00 0x00 0x1a 0x5a 0x0f 0x1f 0xff 0xc2 0x09 0x80 0x00 0x00 0x00 0x07 0xf7 0xeb 0x2a 0xff 0xff 0x7f 0x57 0xe3 0x01 0xff 0xff 0x7f 0x57 0xeb 0x00 0xf0 0x00 0x00 0x24 0xb2 0x4f 0x00 0x78 
Disasm:
0x1000: adb             %f0, 0
0x1006: a               %r0, 0xfff(%r15, %r1)
0x100a: afi             %r0, -0x80000000
0x1010: br              %r7
0x1012: xiy             0x7ffff(%r15), 0x2a
0x1018: xy              %r0, 0x7ffff(%r1, %r15)
0x101e: stmg            %r0, %r0, 0(%r15)
0x1024: ear             %r7, %a8
0x1028:

****************
Platform: XCore
Code: 0xfe 0x0f 0xfe 0x17 0x13 0x17 0xc6 0xfe 0xec 0x17 0x97 0xf8 0xec 0x4f 0x1f 0xfd 0xec 0x37 0x07 0xf2 0x45 0x5b 0xf9 0xfa 0x02 0x06 0x1b 0x10 
Disasm:
0x1000: get             r11, ed
0x1002: ldw             et, sp[4]
0x1004: setd            res[r3], r4
0x1006: init            t[r2]:lr, r1
0x100a: divu            r9, r1, r3
0x100e: lda16           r9, r3[-r11]
0x1012: ldw             dp, dp[0x81c5]
0x1016: lmul            r11, r0, r2, r5, r8, r10
0x101a: add             r1, r2, r3
0x101c:

****************
Platform: M68K
Code: 0xd4 0x40 0x87 0x5a 0x4e 0x71 0x02 0xb4 0xc0 0xde 0xc0 0xde 0x5c 0x00 0x1d 0x80 0x71 0x12 0x01 0x23 0xf2 0x3c 0x44 0x22 0x40 0x49 0x0e 0x56 0x54 0xc5 0xf2 0x3c 0x44 0x00 0x44 0x7a 0x00 0x00 0xf2 0x00 0x0a 0x28 
Disasm:
0x1000: add.w           d0, d2
0x1002: or.w            d3, (a2)+
0x1004: nop
0x1006: andi.l          #$c0dec0de, (a4, d5.l * 4)
0x100e: move.b          d0, ([a6, d7.w], $123)
0x1014: fadd.s          #3.141500, fp0
0x101c: scc.b           d5
0x101e: fmove.s         #1000.000000, fp0
0x1026: fsub            fp2, fp4
0x102a:

****************
Platform: TMS320C64x
Code: 0x01 0xac 0x88 0x40 0x81 0xac 0x88 0x43 0x00 0x00 0x00 0x00 0x02 0x90 0x32 0x96 0x02 0x80 0x46 0x9e 0x05 0x3c 0x83 0xe6 0x0b 0x0c 0x8b 0x24 
Disasm:
0x1000: add.D1          a11, a4, a3
0x1004: [ a1] add.D2            b11, b4, b3     ||
0x1008: NOP
0x100c: ldbu.D1T2               *++a4[1], b5
0x1010: ldbu.D2T2               *+b15[0x46], b5
0x1014: lddw.D1T2               *+a15[4], b11:b10
0x1018: ldndw.D1T1              *+a3(a4), a23:a22
0x101c:

****************
Platform: M680X_M6809
Code: 0x06 0x10 0x19 0x1a 0x55 0x1e 0x01 0x23 0xe9 0x31 0x06 0x34 0x55 0xa6 0x81 0xa7 0x89 0x7f 0xff 0xa6 0x9d 0x10 0x00 0xa7 0x91 0xa6 0x9f 0x10 0x00 0x11 0xac 0x99 0x10 0x00 0x39 
Disasm:
0x1000: ror             $10
0x1002: daa
0x1003: orcc            #85
0x1005: exg             d, x
0x1007: bls             $0ff2
0x1009: leay            6, x
0x100b: pshs            cc, b, x, u
0x100d: lda             , x++
0x100f: sta             32767, x
0x1013: lda             [$2017, pcr]
0x1017: sta             [, x++]
0x1019: lda             [$1000]
0x101d: cmps            [4096, x]
0x1022: rts
0x1023:

****************
Platform: EVM
Code: 0x60 0x61 
Disasm:
0x1000: push1           61
0x1002:

****************
Platform: WASM
Code: 0x20 0x00 0x20 0x01 0x41 0x20 0x10 0xc9 0x01 0x45 0x0b 
Disasm:
0x1000: get_local               0x0
0x1002: get_local               0x1
0x1004: i32.const               0x20
0x1006: call            0xc9
0x1009: i32.eqz
0x100a: end
0x100b:

****************
Platform: MOS65XX
Code: 0x0d 0x34 0x12 0x00 0x81 0x65 0x6c 0x01 0x00 0x85 0xff 0x10 0x00 0x19 0x42 0x42 0x00 0x49 0x42 
Disasm:
0x1000: ora             0x1234
0x1003: brk             0x81
0x1005: adc             0x6c
0x1007: ora             (0x00, x)
0x1009: sta             0xff
0x100b: bpl             0x100d
0x100d: ora             0x4242, y
0x1010: brk             0x49
0x1012:

****************
Platform: eBPF
Code: 0x97 0x09 0x00 0x00 0x37 0x13 0x03 0x00 0xdc 0x02 0x00 0x00 0x20 0x00 0x00 0x00 0x30 0x00 0x00 0x00 0x00 0x00 0x00 0x00 0xdb 0x3a 0x00 0x01 0x00 0x00 0x00 0x00 0x84 0x02 0x00 0x00 0x00 0x00 0x00 0x00 0x6d 0x33 0x17 0x02 0x00 0x00 0x00 0x00 
Disasm:
0x1000: mod64           r9, 0x31337
0x1008: be32            r2
0x1010: ldb             [0x0]
0x1018: xadddw          [r10+0x100], r3
0x1020: neg             r2
0x1028: jsgt            r3, r3, +0x217
0x1030:

****************
Platform: RISCV32
Code: 0x37 0x34 0x00 0x00 0x97 0x82 0x00 0x00 0xef 0x00 0x80 0x00 0xef 0xf0 0x1f 0xff 0xe7 0x00 0x45 0x00 0xe7 0x00 0xc0 0xff 0x63 0x05 0x41 0x00 0xe3 0x9d 0x61 0xfe 0x63 0xca 0x93 0x00 0x63 0x53 0xb5 0x00 0x63 0x65 0xd6 0x00 0x63 0x76 0xf7 0x00 0x03 0x88 0x18 0x00 0x03 0x99 0x49 0x00 0x03 0xaa 0x6a 0x00 0x03 0xcb 0x2b 0x01 0x03 0xdc 0x8c 0x01 0x23 0x86 0xad 0x03 0x23 0x9a 0xce 0x03 0x23 0x8f 0xef 0x01 0x93 0x00 0xe0 0x00 0x13 0xa1 0x01 0x01 0x13 0xb2 0x02 0x7d 0x13 0xc3 0x03 0xdd 0x13 0xe4 0xc4 0x12 0x13 0xf5 0x85 0x0c 0x13 0x96 0xe6 0x01 0x13 0xd7 0x97 0x01 0x13 0xd8 0xf8 0x40 0x33 0x89 0x49 0x01 0xb3 0x0a 0x7b 0x41 0x33 0xac 0xac 0x01 0xb3 0x3d 0xde 0x01 0x33 0xd2 0x62 0x40 0xb3 0x43 0x94 0x00 0x33 0xe5 0xc5 0x00 0xb3 0x76 0xf7 0x00 0xb3 0x54 0x39 0x01 0xb3 0x50 0x31 0x00 0x33 0x9f 0x0f 0x00 
Disasm:
0x1000: lui             s0, 3
0x1004: auipc           t0, 8
0x1008: jal             8
0x100c: jal             -0x10
0x1010: jalr            ra, a0, 4
0x1014: jalr            ra, zero, -4
0x1018: beq             sp, tp, 0xa
0x101c: bne             gp, t1, -6
0x1020: blt             t2, s1, 0x14
0x1024: bge             a0, a1, 6
0x1028: bltu            a2, a3, 0xa
0x102c: bgeu            a4, a5, 0xc
0x1030: lb              a6, 1(a7)
0x1034: lh              s2, 4(s3)
0x1038: lw              s4, 6(s5)
0x103c: lbu             s6, 0x12(s7)
0x1040: lhu             s8, 0x18(s9)
0x1044: sb              s10, 0x2c(s11)
0x1048: sh              t3, 0x34(t4)
0x104c: sb              t5, 0x1e(t6)
0x1050: addi            ra, zero, 0xe
0x1054: slti            sp, gp, 0x10
0x1058: sltiu           tp, t0, 0x7d0
0x105c: xori            t1, t2, -0x230
0x1060: ori             s0, s1, 0x12c
0x1064: andi            a0, a1, 0xc8
0x1068: slli            a2, a3, 0x1e
0x106c: srli            a4, a5, 0x19
0x1070: srai            a6, a7, 0xf
0x1074: add             s2, s3, s4
0x1078: sub             s5, s6, s7
0x107c: slt             s8, s9, s10
0x1080: sltu            s11, t3, t4
0x1084: sra             tp, t0, t1
0x1088: xor             t2, s0, s1
0x108c: or              a0, a1, a2
0x1090: and             a3, a4, a5
0x1094: srl             s1, s2, s3
0x1098: srl             ra, sp, gp
0x109c: sll             t5, t6, zero
0x10a0:

****************
Platform: RISCV64
Code: 0x13 0x04 0xa8 0x7a 
Disasm:
0x1000: addi            s0, a6, 0x7aa