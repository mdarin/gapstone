********************
Platform: SH_SH4A
Code: 0x0c 0x31 0x10 0x20 0x22 0x21 0x36 0x64 0x46 0x25 0x12 0x12 0x1c 0x02 0x08 0xc1 0x05 0xc7 0x0c 0x71 0x1f 0x02 0x22 0xcf 0x06 0x89 0x23 0x00 0x2b 0x41 0x0b 0x00 0x0e 0x40 0x32 0x00 0x0a 0xf1 0x09 0x00 
Disasm:
0x80000000: 0c31       add   r0,r1
	op_count: 2
		operands[0].type: REGISTER = r0
		operands[1].type: REGISTER = r1
	Registers read: r0
	Registers modified: r1

0x80000002: 1020       mov.b r1,@r0
	op_count: 2
		operands[0].type: REGISTER = r1
			size: 8
		operands[1].type: MEM Register indirect REG r0
			size: 8
	Registers read: r0 r1

0x80000004: 2221       mov.l r2,@r1
	op_count: 2
		operands[0].type: REGISTER = r2
			size: 32
		operands[1].type: MEM Register indirect REG r1
			size: 32
	Registers read: r1 r2

0x80000006: 3664       mov.l @r3+,r4
	op_count: 2
		operands[0].type: MEM Register indirect with predecrement REG r3
			size: 32
		operands[1].type: REGISTER = r4
			size: 32
	Registers modified: r3 r4

0x80000008: 4625       mov.l r4,@-r5
	op_count: 2
		operands[0].type: REGISTER = r4
			size: 32
		operands[1].type: MEM Register indirect with postincrement REG r5
			size: 32
	Registers read: r4
	Registers modified: r5

0x8000000a: 1212       mov.l r1,@(8,r2)
	op_count: 2
		operands[0].type: REGISTER = r1
			size: 32
		operands[1].type: MEM Register indirect with displacement REG r2, DISP 8
			size: 32
	Registers read: r2 r1

0x8000000c: 1c02       mov.b @(r0,r1),r2
	op_count: 2
		operands[0].type: MEM R0 indexed
			size: 8
		operands[1].type: REGISTER = r2
			size: 8
	Registers read: r0 r1
	Registers modified: r2

0x8000000e: 08c1       mov.w r0,@(16,gbr)
	op_count: 2
		operands[0].type: REGISTER = r0
			size: 16
		operands[1].type: MEM GBR base with displacement DISP 16
			size: 16
	Registers read: gbr r0

0x80000010: 05c7       mova  0x80000028,r0
	op_count: 2
		operands[0].type: MEM PC relative Address=0x80000028
		operands[1].type: REGISTER = r0
	Registers modified: r0

0x80000012: 0c71       add   #12,r1
	op_count: 2
		operands[0].type: IMMEDIATE = #12
		operands[1].type: REGISTER = r1
	Registers modified: r1

0x80000014: 1f02       mac.l @r1+,@r2+
	op_count: 2
		operands[0].type: MEM Register indirect with predecrement REG r1
		operands[1].type: MEM Register indirect with predecrement REG r2
	Registers modified: r1 r2

0x80000016: 22cf       or.b  #34,@(r0,gbr)
	op_count: 2
		operands[0].type: IMMEDIATE = #34
			size: 8
		operands[1].type: MEM GBR base with R0 indexed
			size: 8
	Registers read: gbr r0

0x80000018: 0689       bt    0x80000028
	op_count: 1
		operands[0].type: MEM PC relative Address=0x80000028
	groups_count: 2

0x8000001a: 2300       braf  r0
	op_count: 1
		operands[0].type: REGISTER = r0
	Registers read: r0
	groups_count: 2

0x8000001c: 2b41       jmp   @r1
	op_count: 1
		operands[0].type: MEM Register indirect REG r1
	Registers read: r1
	groups_count: 1

0x8000001e: 0b00       rts   

0x80000020: 0e40       ldc   r0,sr
	op_count: 2
		operands[0].type: REGISTER = r0
		operands[1].type: REGISTER = sr
	Registers read: r0
	Registers modified: sr

0x80000022: 3200       stc   ssr,r0
	op_count: 2
		operands[0].type: REGISTER = ssr
		operands[1].type: REGISTER = r0
	Registers read: ssr
	Registers modified: r0

0x80000024: 0af1       fmov  fr0,@r1
	op_count: 2
		operands[0].type: REGISTER = fr0
		operands[1].type: MEM Register indirect REG r1
	Registers read: r1 fr0

0x80000026: 0900       nop   

********************
Platform: SH_SH2A
Code: 0x32 0x11 0x92 0x00 0x32 0x49 0x31 0x00 
Disasm:
0x80000000: 32119200   movu.w@(1024,r1),r2
	op_count: 2
		operands[0].type: MEM Register indirect with displacement REG r1, DISP 1024
			size: 16
		operands[1].type: REGISTER = r2
			size: 16
	Registers read: r1
	Registers modified: r2

0x80000004: 32493100   bld.b #4,@(256,r2)
	op_count: 2
		operands[0].type: IMMEDIATE = #4
			size: 8
		operands[1].type: MEM Register indirect with displacement REG r2, DISP 256
			size: 8
	Registers read: r2

