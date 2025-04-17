****************
Platform: X86 16bit (Intel syntax)
Code: 0x8d 0x4c 0x32 0x08 0x01 0xd8 0x81 0xc6 0x34 0x12 0x00 0x00 
Disasm:
0x1000:	lea		cx, [si + 0x32] // insn-ID: 332, insn-mnem: lea
0x1003:	or		byte ptr [bx + di], al // insn-ID: 512, insn-mnem: or
	Implicit registers modified: flags 
0x1005:	fadd		dword ptr [bx + di + 0x34c6] // insn-ID: 15, insn-mnem: fadd
	Implicit registers modified: fpsw 
	This instruction belongs to groups: fpu 
0x1009:	adc		al, byte ptr [bx + si] // insn-ID: 6, insn-mnem: adc
	Implicit registers read: flags 
	Implicit registers modified: flags 

****************
Platform: X86 32bit (ATT syntax)
Code: 0x8d 0x4c 0x32 0x08 0x01 0xd8 0x81 0xc6 0x34 0x12 0x00 0x00 
Disasm:
0x1000:	leal		8(%edx, %esi), %ecx // insn-ID: 332, insn-mnem: lea
	This instruction belongs to groups: not64bitmode 
0x1004:	addl		%ebx, %eax // insn-ID: 8, insn-mnem: add
	Implicit registers modified: eflags 
0x1006:	addl		$0x1234, %esi // insn-ID: 8, insn-mnem: add
	Implicit registers modified: eflags 

****************
Platform: X86 32 (Intel syntax)
Code: 0x8d 0x4c 0x32 0x08 0x01 0xd8 0x81 0xc6 0x34 0x12 0x00 0x00 
Disasm:
0x1000:	lea		ecx, [edx + esi + 8] // insn-ID: 332, insn-mnem: lea
	This instruction belongs to groups: not64bitmode 
0x1004:	add		eax, ebx // insn-ID: 8, insn-mnem: add
	Implicit registers modified: eflags 
0x1006:	add		esi, 0x1234 // insn-ID: 8, insn-mnem: add
	Implicit registers modified: eflags 

****************
Platform: X86 64 (Intel syntax)
Code: 0x55 0x48 0x8b 0x05 0xb8 0x13 0x00 0x00 
Disasm:
0x1000:	push		rbp // insn-ID: 609, insn-mnem: push
	Implicit registers read: rsp 
	Implicit registers modified: rsp 
	This instruction belongs to groups: mode64 
0x1001:	mov		rax, qword ptr [rip + 0x13b8] // insn-ID: 460, insn-mnem: mov

****************
Platform: ARM
Code: 0xed 0xff 0xff 0xeb 0x04 0xe0 0x2d 0xe5 0x00 0x00 0x00 0x00 0xe0 0x83 0x22 0xe5 0xf1 0x02 0x03 0x0e 0x00 0x00 0xa0 0xe3 0x02 0x30 0xc1 0xe7 0x00 0x00 0x53 0xe3 
Disasm:
0x1000:	bl		#0xfbc // insn-ID: 16, insn-mnem: bl
	Implicit registers read: pc 
	Implicit registers modified: lr pc 
	This instruction belongs to groups: call branch_relative arm jump 
0x1004:	str		lr, [sp, #-4]! // insn-ID: 240, insn-mnem: str
	This instruction belongs to groups: arm 
0x1008:	andeq		r0, r0, r0 // insn-ID: 9, insn-mnem: and
	This instruction belongs to groups: arm 
0x100c:	str		r8, [r2, #-0x3e0]! // insn-ID: 240, insn-mnem: str
	This instruction belongs to groups: arm 
0x1010:	mcreq		p2, #0, r0, c3, c1, #7 // insn-ID: 100, insn-mnem: mcr
	This instruction belongs to groups: privilege arm 
0x1014:	mov		r0, #0 // insn-ID: 106, insn-mnem: mov
	This instruction belongs to groups: arm 
0x1018:	strb		r3, [r1, r2] // insn-ID: 241, insn-mnem: strb
	This instruction belongs to groups: arm 
0x101c:	cmp		r3, #0 // insn-ID: 29, insn-mnem: cmp
	Implicit registers modified: cpsr 
	This instruction belongs to groups: arm 

****************
Platform: THUMB-2
Code: 0x4f 0xf0 0x00 0x01 0xbd 0xe8 0x00 0x88 0xd1 0xe8 0x00 0xf0 
Disasm:
0x1000:	mov.w		r1, #0 // insn-ID: 106, insn-mnem: mov
	This instruction belongs to groups: thumb2 
0x1004:	pop.w		{fp, pc} // insn-ID: 127, insn-mnem: pop
	Implicit registers read: sp 
	Implicit registers modified: sp 
	This instruction belongs to groups: thumb2 
0x1008:	tbb		[r1, r0] // insn-ID: 263, insn-mnem: tbb
	This instruction belongs to groups: thumb2 jump 

****************
Platform: ARM: Cortex-A15 + NEON
Code: 0x10 0xf1 0x10 0xe7 0x11 0xf2 0x31 0xe7 0xdc 0xa1 0x2e 0xf3 0xe8 0x4e 0x62 0xf3 
Disasm:
0x1000:	sdiv		r0, r0, r1 // insn-ID: 156, insn-mnem: sdiv
	This instruction belongs to groups: arm 
0x1004:	udiv		r1, r1, r2 // insn-ID: 278, insn-mnem: udiv
	This instruction belongs to groups: arm 
0x1008:	vbit		q5, q15, q6 // insn-ID: 323, insn-mnem: vbit
	This instruction belongs to groups: neon 
0x100c:	vcgt.f32		q10, q9, q12 // insn-ID: 328, insn-mnem: vcgt
	This instruction belongs to groups: neon 

****************
Platform: THUMB
Code: 0x70 0x47 0xeb 0x46 0x83 0xb0 0xc9 0x68 
Disasm:
0x1000:	bx		lr // insn-ID: 19, insn-mnem: bx
	Implicit registers modified: pc 
	This instruction belongs to groups: thumb jump 
0x1002:	mov		fp, sp // insn-ID: 106, insn-mnem: mov
	This instruction belongs to groups: thumb thumb1only 
0x1004:	sub		sp, #0xc // insn-ID: 251, insn-mnem: sub
	This instruction belongs to groups: thumb thumb1only 
0x1006:	ldr		r1, [r1, #0xc] // insn-ID: 83, insn-mnem: ldr
	This instruction belongs to groups: thumb thumb1only 

****************
Platform: MIPS-32 (Big-endian)
Code: 0x0c 0x10 0x00 0x97 0x00 0x00 0x00 0x00 0x24 0x02 0x00 0x0c 0x8f 0xa2 0x00 0x00 0x34 0x21 0x34 0x56 0x00 0x80 0x04 0x08 
Disasm:
0x1000:	jal		0x40025c // insn-ID: 337, insn-mnem: jal
	Implicit registers modified: ra 
	This instruction belongs to groups: stdenc 
0x1004:	nop		 // insn-ID: 622, insn-mnem: nop
	This instruction belongs to groups: stdenc notinmicromips 
0x1008:	addiu		$v0, $zero, 0xc // insn-ID: 26, insn-mnem: addiu
	This instruction belongs to groups: stdenc notinmicromips 
0x100c:	lw		$v0, ($sp) // insn-ID: 373, insn-mnem: lw
	This instruction belongs to groups: stdenc notinmicromips 
0x1010:	ori		$at, $at, 0x3456 // insn-ID: 473, insn-mnem: ori
	This instruction belongs to groups: stdenc 
0x1014:	jr.hb		$a0 // insn-ID: 345, insn-mnem: jr
	This instruction belongs to groups: stdenc mips32 notmips32r6 notmips64r6 jump 

****************
Platform: MIPS-64-EL (Little-endian)
Code: 0x56 0x34 0x21 0x34 0xc2 0x17 0x01 0x00 
Disasm:
0x1000:	ori		$at, $at, 0x3456 // insn-ID: 473, insn-mnem: ori
	This instruction belongs to groups: stdenc 
0x1004:	srl		$v0, $at, 0x1f // insn-ID: 557, insn-mnem: srl
	This instruction belongs to groups: stdenc notinmicromips 

****************
Platform: ARM-64
Code: 0x09 0x00 0x38 0xd5 0xbf 0x40 0x00 0xd5 0x0c 0x05 0x13 0xd5 0x20 0x50 0x02 0x0e 0x20 0xe4 0x3d 0x0f 0x00 0x18 0xa0 0x5f 0xa2 0x00 0xae 0x9e 0x9f 0x37 0x03 0xd5 0xbf 0x33 0x03 0xd5 0xdf 0x3f 0x03 0xd5 0x21 0x7c 0x02 0x9b 0x21 0x7c 0x00 0x53 0x00 0x40 0x21 0x4b 0xe1 0x0b 0x40 0xb9 0x20 0x04 0x81 0xda 0x20 0x08 0x02 0x8b 0x10 0x5b 0xe8 0x3c 
Disasm:
0x1000:	mrs		x9, midr_el1 // insn-ID: 495, insn-mnem: mrs
	This instruction belongs to groups: privilege 
0x1004:	msr		spsel, #0 // insn-ID: 497, insn-mnem: msr
	Implicit registers modified: nzcv 
	This instruction belongs to groups: privilege 
0x1008:	msr		dbgdtrtx_el0, x12 // insn-ID: 497, insn-mnem: msr
	This instruction belongs to groups: privilege 
0x100c:	tbx		v0.8b, {v1.16b, v2.16b, v3.16b}, v2.8b // insn-ID: 835, insn-mnem: tbx
	This instruction belongs to groups: neon 
0x1010:	scvtf		v0.2s, v1.2s, #3 // insn-ID: 592, insn-mnem: scvtf
	This instruction belongs to groups: neon 
0x1014:	fmla		s0, s0, v0.s[3] // insn-ID: 231, insn-mnem: fmla
	This instruction belongs to groups: neon 
0x1018:	fmov		x2, v5.d[1] // insn-ID: 233, insn-mnem: fmov
	This instruction belongs to groups: fparmv8 
0x101c:	dsb		nsh // insn-ID: 156, insn-mnem: dsb
0x1020:	dmb		osh // insn-ID: 154, insn-mnem: dmb
0x1024:	isb		 // insn-ID: 276, insn-mnem: isb
0x1028:	mul		x1, x1, x2 // insn-ID: 499, insn-mnem: mul
0x102c:	lsr		w1, w1, #0 // insn-ID: 480, insn-mnem: lsr
0x1030:	sub		w0, w0, w1, uxtw // insn-ID: 805, insn-mnem: sub
0x1034:	ldr		w1, [sp, #8] // insn-ID: 393, insn-mnem: ldr
0x1038:	cneg		x0, x1, ne // insn-ID: 119, insn-mnem: cneg
	Implicit registers read: nzcv 
0x103c:	add		x0, x1, x2, lsl #2 // insn-ID: 4, insn-mnem: add
0x1040:	ldr		q16, [x24, w8, uxtw #4] // insn-ID: 393, insn-mnem: ldr

****************
Platform: PPC-64
Code: 0x80 0x20 0x00 0x00 0x80 0x3f 0x00 0x00 0x10 0x43 0x23 0x0e 0xd0 0x44 0x00 0x80 0x4c 0x43 0x22 0x02 0x2d 0x03 0x00 0x80 0x7c 0x43 0x20 0x14 0x7c 0x43 0x20 0x93 0x4f 0x20 0x00 0x21 0x4c 0xc8 0x00 0x21 0x40 0x82 0x00 0x14 
Disasm:
0x1000:	lwz		r1, 0(0) // insn-ID: 651, insn-mnem: lwz
0x1004:	lwz		r1, 0(r31) // insn-ID: 651, insn-mnem: lwz
0x1008:	vpkpx		v2, v3, v4 // insn-ID: 1350, insn-mnem: vpkpx
	This instruction belongs to groups: altivec 
0x100c:	stfs		f2, 0x80(r4) // insn-ID: 1046, insn-mnem: stfs
0x1010:	crand		cr0eq, cr0un, cr1lt // insn-ID: 223, insn-mnem: crand
0x1014:	cmpwi		cr2, r3, 0x80 // insn-ID: 215, insn-mnem: cmpwi
0x1018:	addc		r2, r3, r4 // insn-ID: 2, insn-mnem: addc
	Implicit registers modified: ca 
0x101c:	mulhd.		r2, r3, r4 // insn-ID: 830, insn-mnem: mulhd
	Implicit registers modified: cr0 
0x1020:	bdnzlrl+		 // insn-ID: 51, insn-mnem: bdnzlrl
	Implicit registers read: ctr lr rm 
	Implicit registers modified: ctr 
	This instruction belongs to groups: jump 
0x1024:	bgelrl-		cr2 // insn-ID: 99, insn-mnem: bgelrl
	Implicit registers read: ctr lr rm 
	Implicit registers modified: lr ctr 
	This instruction belongs to groups: jump 
0x1028:	bne		0x103c // insn-ID: 128, insn-mnem: bne
	Implicit registers read: ctr rm 
	Implicit registers modified: ctr 
	This instruction belongs to groups: jump 

****************
Platform: Sparc
Code: 0x80 0xa0 0x40 0x02 0x85 0xc2 0x60 0x08 0x85 0xe8 0x20 0x01 0x81 0xe8 0x00 0x00 0x90 0x10 0x20 0x01 0xd5 0xf6 0x10 0x16 0x21 0x00 0x00 0x0a 0x86 0x00 0x40 0x02 0x01 0x00 0x00 0x00 0x12 0xbf 0xff 0xff 0x10 0xbf 0xff 0xff 0xa0 0x02 0x00 0x09 0x0d 0xbf 0xff 0xff 0xd4 0x20 0x60 0x00 0xd4 0x4e 0x00 0x16 0x2a 0xc2 0x80 0x03 
Disasm:
0x1000:	cmp		%g1, %g2 // insn-ID: 33, insn-mnem: cmp
	Implicit registers modified: icc 
0x1004:	jmpl		%o1+8, %g2 // insn-ID: 194, insn-mnem: jmpl
0x1008:	restore		%g0, 1, %g2 // insn-ID: 226, insn-mnem: restore
0x100c:	restore		 // insn-ID: 226, insn-mnem: restore
0x1010:	mov		1, %o0 // insn-ID: 207, insn-mnem: mov
0x1014:	casx		[%i0], %l6, %o2 // insn-ID: 28, insn-mnem: casx
	This instruction belongs to groups: 64bit 
0x1018:	sethi		0xa, %l0 // insn-ID: 232, insn-mnem: sethi
0x101c:	add		%g1, %g2, %g3 // insn-ID: 6, insn-mnem: add
0x1020:	nop		 // insn-ID: 217, insn-mnem: nop
0x1024:	bne		0x1020 // insn-ID: 16, insn-mnem: b
	Implicit registers read: icc 
	This instruction belongs to groups: jump 
0x1028:	ba		0x1024 // insn-ID: 16, insn-mnem: b
	This instruction belongs to groups: jump 
0x102c:	add		%o0, %o1, %l0 // insn-ID: 6, insn-mnem: add
0x1030:	fbg		0x102c // insn-ID: 19, insn-mnem: fb
	Implicit registers read: fcc0 
	This instruction belongs to groups: jump 
0x1034:	st		%o2, [%g1] // insn-ID: 246, insn-mnem: st
0x1038:	ldsb		[%i0+%l6], %o2 // insn-ID: 198, insn-mnem: ldsb
0x103c:	brnz,a,pn		%o2, 0x1048 // insn-ID: 24, insn-mnem: brnz
	This instruction belongs to groups: 64bit jump 

****************
Platform: SparcV9
Code: 0x81 0xa8 0x0a 0x24 0x89 0xa0 0x10 0x20 0x89 0xa0 0x1a 0x60 0x89 0xa0 0x00 0xe0 
Disasm:
0x1000:	fcmps		%f0, %f4 // insn-ID: 70, insn-mnem: fcmps
0x1004:	fstox		%f0, %f4 // insn-ID: 181, insn-mnem: fstox
	This instruction belongs to groups: 64bit 
0x1008:	fqtoi		%f0, %f4 // insn-ID: 159, insn-mnem: fqtoi
	This instruction belongs to groups: hardquad 
0x100c:	fnegq		%f0, %f4 // insn-ID: 127, insn-mnem: fnegq
	This instruction belongs to groups: v9 

****************
Platform: SystemZ
Code: 0xed 0x00 0x00 0x00 0x00 0x1a 0x5a 0x0f 0x1f 0xff 0xc2 0x09 0x80 0x00 0x00 0x00 0x07 0xf7 0xeb 0x2a 0xff 0xff 0x7f 0x57 0xe3 0x01 0xff 0xff 0x7f 0x57 0xeb 0x00 0xf0 0x00 0x00 0x24 0xb2 0x4f 0x00 0x78 
Disasm:
0x1000:	adb		%f0, 0 // insn-ID: 2, insn-mnem: adb
	Implicit registers modified: cc 
0x1006:	a		%r0, 0xfff(%r15, %r1) // insn-ID: 1, insn-mnem: a
	Implicit registers modified: cc 
0x100a:	afi		%r0, -0x80000000 // insn-ID: 6, insn-mnem: afi
	Implicit registers modified: cc 
0x1010:	br		%r7 // insn-ID: 283, insn-mnem: br
	This instruction belongs to groups: jump 
0x1012:	xiy		0x7ffff(%r15), 0x2a // insn-ID: 678, insn-mnem: xiy
	Implicit registers modified: cc 
0x1018:	xy		%r0, 0x7ffff(%r1, %r15) // insn-ID: 681, insn-mnem: xy
	Implicit registers modified: cc 
0x101e:	stmg		%r0, %r0, 0(%r15) // insn-ID: 657, insn-mnem: stmg
0x1024:	ear		%r7, %a8 // insn-ID: 383, insn-mnem: ear

****************
Platform: XCore
Code: 0xfe 0x0f 0xfe 0x17 0x13 0x17 0xc6 0xfe 0xec 0x17 0x97 0xf8 0xec 0x4f 0x1f 0xfd 0xec 0x37 0x07 0xf2 0x45 0x5b 0xf9 0xfa 0x02 0x06 0x1b 0x10 
Disasm:
0x1000:	get		r11, ed // insn-ID: 43, insn-mnem: get
	Implicit registers modified: r11 
0x1002:	ldw		et, sp[4] // insn-ID: 66, insn-mnem: ldw
	Implicit registers read: sp 
0x1004:	setd		res[r3], r4 // insn-ID: 93, insn-mnem: setd
0x1006:	init		t[r2]:lr, r1 // insn-ID: 50, insn-mnem: init
0x100a:	divu		r9, r1, r3 // insn-ID: 26, insn-mnem: divu
0x100e:	lda16		r9, r3[-r11] // insn-ID: 62, insn-mnem: lda16
0x1012:	ldw		dp, dp[0x81c5] // insn-ID: 66, insn-mnem: ldw
0x1016:	lmul		r11, r0, r2, r5, r8, r10 // insn-ID: 68, insn-mnem: lmul
0x101a:	add		r1, r2, r3 // insn-ID: 1, insn-mnem: add

****************
Platform: M680X_6809
Code: 0x06 0x10 0x19 0x1a 0x55 0x1e 0x01 0x23 0xe9 0x31 0x06 0x34 0x55 0xa6 0x81 0xa7 0x89 0x7f 0xff 0xa6 0x9d 0x10 0x00 0xa7 0x91 0xa6 0x9f 0x10 0x00 0x11 0xac 0x99 0x10 0x00 0x39 
Disasm:
0x1000:	ror		$10 // insn-ID: 276, insn-mnem: ror
	Implicit registers read: cc 
	Implicit registers modified: cc 
0x1002:	daa		 // insn-ID: 117, insn-mnem: daa
	Implicit registers read: cc a 
	Implicit registers modified: cc a 
0x1003:	orcc		#85 // insn-ID: 243, insn-mnem: orcc
	Implicit registers read: cc 
	Implicit registers modified: cc 
0x1005:	exg		d, x // insn-ID: 153, insn-mnem: exg
	Implicit registers read: d x 
	Implicit registers modified: d x 
0x1007:	bls		$0ff2 // insn-ID: 58, insn-mnem: bls
	Implicit registers read: cc 
	This instruction belongs to groups: branch_relative jump 
0x1009:	leay		6, x // insn-ID: 209, insn-mnem: leay
	Implicit registers read: cc x 
	Implicit registers modified: cc y 
0x100b:	pshs		cc, b, x, u // insn-ID: 251, insn-mnem: pshs
	Implicit registers read: s cc b x u 
	Implicit registers modified: s 
0x100d:	lda		, x++ // insn-ID: 190, insn-mnem: lda
	Implicit registers read: cc x 
	Implicit registers modified: cc a x 
0x100f:	sta		32767, x // insn-ID: 298, insn-mnem: sta
	Implicit registers read: cc a x 
	Implicit registers modified: cc 
0x1013:	lda		[$2017, pcr] // insn-ID: 190, insn-mnem: lda
	Implicit registers read: cc pc 
	Implicit registers modified: cc a 
0x1017:	sta		[, x++] // insn-ID: 298, insn-mnem: sta
	Implicit registers read: cc a x 
	Implicit registers modified: cc x 
0x1019:	lda		[$1000] // insn-ID: 190, insn-mnem: lda
	Implicit registers read: cc 
	Implicit registers modified: cc a 
0x101d:	cmps		[4096, x] // insn-ID: 98, insn-mnem: cmps
	Implicit registers read: cc s x 
	Implicit registers modified: cc 
0x1022:	rts		 // insn-ID: 285, insn-mnem: rts
	Implicit registers read: s 
	Implicit registers modified: s pc 
	This instruction belongs to groups: return 

****************
Platform: MOS65XX
Code: 0x0d 0x34 0x12 0x08 0x09 0xff 0x10 0x80 0x20 0x00 0x00 0x98 
Disasm:
0x1000:	ora		0x1234 // insn-ID: 44, insn-mnem: ora
	Implicit registers read: A 
	Implicit registers modified: A P 
0x1003:	php		 // insn-ID: 52, insn-mnem: php
	Implicit registers read: P SP 
	Implicit registers modified: SP 
0x1004:	ora		#0xff // insn-ID: 44, insn-mnem: ora
	Implicit registers read: A 
	Implicit registers modified: A P 
0x1006:	bpl		0x0f88 // insn-ID: 12, insn-mnem: bpl
	Implicit registers read: P 
	This instruction belongs to groups: jump branch_relative 
0x1008:	jsr		0x0000 // insn-ID: 36, insn-mnem: jsr
	Implicit registers read: SP 
	Implicit registers modified: SP 
	This instruction belongs to groups: call 
0x100b:	tya		 // insn-ID: 91, insn-mnem: tya
	Implicit registers read: Y 
	Implicit registers modified: A P 

****************
Platform: eBPF
Code: 0x97 0x09 0x00 0x00 0x37 0x13 0x03 0x00 0xdc 0x02 0x00 0x00 0x20 0x00 0x00 0x00 0x30 0x00 0x00 0x00 0x00 0x00 0x00 0x00 0xdb 0x3a 0x00 0x01 0x00 0x00 0x00 0x00 0x84 0x02 0x00 0x00 0x00 0x00 0x00 0x00 0x6d 0x33 0x17 0x02 0x00 0x00 0x00 0x00 
Disasm:
0x1000:	mod64		r9, 0x31337 // insn-ID: 23, insn-mnem: mod64
	This instruction belongs to groups: alu 
0x1008:	be32		r2 // insn-ID: 31, insn-mnem: be32
	This instruction belongs to groups: alu 
0x1010:	ldb		[0x0] // insn-ID: 35, insn-mnem: ldb
	Implicit registers modified: r0 
	This instruction belongs to groups: load 
0x1018:	xadddw		[r10+0x100], r3 // insn-ID: 50, insn-mnem: xadddw
	This instruction belongs to groups: store 
0x1020:	neg		r2 // insn-ID: 9, insn-mnem: neg
	This instruction belongs to groups: alu 
0x1028:	jsgt		r3, r3, +0x217 // insn-ID: 57, insn-mnem: jsgt
	This instruction belongs to groups: jump 

****************
Platform: RISCV32
Code: 0x37 0x34 0x00 0x00 0x97 0x82 0x00 0x00 0xef 0x00 0x80 0x00 0xef 0xf0 0x1f 0xff 0xe7 0x00 0x45 0x00 0xe7 0x00 0xc0 0xff 0x63 0x05 0x41 0x00 0xe3 0x9d 0x61 0xfe 0x63 0xca 0x93 0x00 0x63 0x53 0xb5 0x00 0x63 0x65 0xd6 0x00 0x63 0x76 0xf7 0x00 0x03 0x88 0x18 0x00 0x03 0x99 0x49 0x00 0x03 0xaa 0x6a 0x00 0x03 0xcb 0x2b 0x01 0x03 0xdc 0x8c 0x01 0x23 0x86 0xad 0x03 0x23 0x9a 0xce 0x03 0x23 0x8f 0xef 0x01 0x93 0x00 0xe0 0x00 0x13 0xa1 0x01 0x01 0x13 0xb2 0x02 0x7d 0x13 0xc3 0x03 0xdd 0x13 0xe4 0xc4 0x12 0x13 0xf5 0x85 0x0c 0x13 0x96 0xe6 0x01 0x13 0xd7 0x97 0x01 0x13 0xd8 0xf8 0x40 0x33 0x89 0x49 0x01 0xb3 0x0a 0x7b 0x41 0x33 0xac 0xac 0x01 0xb3 0x3d 0xde 0x01 0x33 0xd2 0x62 0x40 0xb3 0x43 0x94 0x00 0x33 0xe5 0xc5 0x00 0xb3 0x76 0xf7 0x00 0xb3 0x54 0x39 0x01 0xb3 0x50 0x31 0x00 0x33 0x9f 0x0f 0x00 
Disasm:
0x1000:	lui		s0, 3 // insn-ID: 221, insn-mnem: lui
0x1004:	auipc		t0, 8 // insn-ID: 79, insn-mnem: auipc
0x1008:	jal		8 // insn-ID: 206, insn-mnem: jal
0x100c:	jal		-0x10 // insn-ID: 206, insn-mnem: jal
0x1010:	jalr		ra, a0, 4 // insn-ID: 207, insn-mnem: jalr
0x1014:	jalr		ra, zero, -4 // insn-ID: 207, insn-mnem: jalr
0x1018:	beq		sp, tp, 0xa // insn-ID: 80, insn-mnem: beq
	This instruction belongs to groups: jump 
0x101c:	bne		gp, t1, -6 // insn-ID: 85, insn-mnem: bne
	This instruction belongs to groups: jump 
0x1020:	blt		t2, s1, 0x14 // insn-ID: 83, insn-mnem: blt
	This instruction belongs to groups: jump 
0x1024:	bge		a0, a1, 6 // insn-ID: 81, insn-mnem: bge
	This instruction belongs to groups: jump 
0x1028:	bltu		a2, a3, 0xa // insn-ID: 84, insn-mnem: bltu
	This instruction belongs to groups: jump 
0x102c:	bgeu		a4, a5, 0xc // insn-ID: 82, insn-mnem: bgeu
	This instruction belongs to groups: jump 
0x1030:	lb		a6, 1(a7) // insn-ID: 208, insn-mnem: lb
0x1034:	lh		s2, 4(s3) // insn-ID: 211, insn-mnem: lh
0x1038:	lw		s4, 6(s5) // insn-ID: 222, insn-mnem: lw
0x103c:	lbu		s6, 0x12(s7) // insn-ID: 209, insn-mnem: lbu
0x1040:	lhu		s8, 0x18(s9) // insn-ID: 212, insn-mnem: lhu
0x1044:	sb		s10, 0x2c(s11) // insn-ID: 236, insn-mnem: sb
0x1048:	sh		t3, 0x34(t4) // insn-ID: 247, insn-mnem: sh
0x104c:	sb		t5, 0x1e(t6) // insn-ID: 236, insn-mnem: sb
0x1050:	addi		ra, zero, 0xe // insn-ID: 2, insn-mnem: addi
0x1054:	slti		sp, gp, 0x10 // insn-ID: 253, insn-mnem: slti
0x1058:	sltiu		tp, t0, 0x7d0 // insn-ID: 254, insn-mnem: sltiu
0x105c:	xori		t1, t2, -0x230 // insn-ID: 272, insn-mnem: xori
0x1060:	ori		s0, s1, 0x12c // insn-ID: 231, insn-mnem: ori
0x1064:	andi		a0, a1, 0xc8 // insn-ID: 78, insn-mnem: andi
0x1068:	slli		a2, a3, 0x1e // insn-ID: 249, insn-mnem: slli
0x106c:	srli		a4, a5, 0x19 // insn-ID: 262, insn-mnem: srli
0x1070:	srai		a6, a7, 0xf // insn-ID: 257, insn-mnem: srai
0x1074:	add		s2, s3, s4 // insn-ID: 1, insn-mnem: add
0x1078:	sub		s5, s6, s7 // insn-ID: 265, insn-mnem: sub
0x107c:	slt		s8, s9, s10 // insn-ID: 252, insn-mnem: slt
0x1080:	sltu		s11, t3, t4 // insn-ID: 255, insn-mnem: sltu
0x1084:	sra		tp, t0, t1 // insn-ID: 256, insn-mnem: sra
0x1088:	xor		t2, s0, s1 // insn-ID: 271, insn-mnem: xor
0x108c:	or		a0, a1, a2 // insn-ID: 230, insn-mnem: or
0x1090:	and		a3, a4, a5 // insn-ID: 77, insn-mnem: and
0x1094:	srl		s1, s2, s3 // insn-ID: 261, insn-mnem: srl
0x1098:	srl		ra, sp, gp // insn-ID: 261, insn-mnem: srl
0x109c:	sll		t5, t6, zero // insn-ID: 248, insn-mnem: sll

****************
Platform: RISCV64
Code: 0x13 0x04 0xa8 0x7a 
Disasm:
0x1000:	addi		s0, a6, 0x7aa // insn-ID: 2, insn-mnem: addi

