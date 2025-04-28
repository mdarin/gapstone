****************
Platform: TriCore
Code:0x09 0xcf 0xbc 0xf5 0x09 0xf4 0x01 0x00 0x89 0xfb 0x8f 0x74 0x89 0xfe 0x48 0x01 0x29 0x00 0x19 0x25 0x29 0x03 0x09 0xf4 0x85 0xf9 0x68 0x0f 0x16 0x01 
Disasm:
0x1000:	ld.a	a15, [+a12]#-4
	op_count: 2
		operands[0].type: REG = a15
		operands[1].type: MEM
			operands[1].mem.base: REG = a12
			operands[1].mem.disp: 0xfffffffc

0x1004:	ld.b	d4, [a15+]#1
	op_count: 2
		operands[0].type: REG = d4
		operands[1].type: MEM
			operands[1].mem.base: REG = a15
			operands[1].mem.disp: 0x1

0x1008:	st.h	[+a15]#0x1cf, d11
	op_count: 2
		operands[0].type: MEM
			operands[0].mem.base: REG = a15
			operands[0].mem.disp: 0x1cf
		operands[1].type: REG = d11

0x100c:	st.d	[a15+]#8, e14
	op_count: 2
		operands[0].type: MEM
			operands[0].mem.base: REG = a15
			operands[0].mem.disp: 0x8
		operands[1].type: REG = e14

0x1010:	ld.w	d0, [p0+c]#0x99
	op_count: 2
		operands[0].type: REG = d0
		operands[1].type: MEM
			operands[1].mem.base: REG = p0
			operands[1].mem.disp: 0x99

0x1014:	ld.b	d3, [p0+c]#-0x37
	op_count: 2
		operands[0].type: REG = d3
		operands[1].type: MEM
			operands[1].mem.base: REG = p0
			operands[1].mem.disp: 0xffffffc9

0x1018:	ld.da	p8, #0xf0003428
	op_count: 2
		operands[0].type: REG = p8
		operands[1].type: IMM = 0xf0003428

0x101c:	and	d15, #1
	op_count: 1
		operands[0].type: IMM = 0x1

0x101e:

