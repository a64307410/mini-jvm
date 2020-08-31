package bcode

const (
	Nop byte = 0x00

	Iconst0 = 0x03
	Iconst1 = 0x04
	Iconst2 = 0x05
	Iconst3 = 0x06
	Iconst4 = 0x07
	Iconst5 = 0x08

	Ldc = 0x12

	Iaload = 0x2e

	Aaload = 0x32
	Caload = 0x34

	Istore0 = 0x3b
	Istore1 = 0x3c
	Istore2 = 0x3d
	Istore3 = 0x3e

	Bipush = 0x10
	Sipush = 0x11

	Iload = 0x15
	Iload0 = 0x1a
	Iload1 = 0x1b
	Iload2 = 0x1c
	Iload3 = 0x1d

	Aload = 0x19
	Aload0 = 0x2a
	Aload1 = 0x2b
	Aload2 = 0x2c
	Aload3 = 0x2d

	Getstatic = 0xb2
	Putstatic = 0xb3

	Athrow = 0xbf

	Monitorenter = 0xc2
	Monitorexit = 0xc3

	Istore = 0x36
	Lstore1 = 0x40

	Astore = 0x3a
	Astore0 = 0x4b
	Astore1 = 0x4c
	Astore2 = 0x4d
	Astore3 = 0x4e
	Iastore = 0x4f

	Aastore = 0x53
	Castore = 0x55

	Dup = 0x59

	Iadd = 0x60
	Isub = 0x64

	Iinc = 0x84

	Ifeq = 0x99
	Ifne = 0x9a
	Iflt = 0x9b
	Ifge = 0x9c
	Ifgt = 0x9d
	Ifle = 0x9e

	Ificmpeq = 0x9f
	Ificmpne = 0xa0
	Ificmplt = 0xa1
	Ificmpge = 0xa2
	Ificmpgt = 0xa3
	Ificmple = 0xa4
	Goto = 0xa7

	Return = 0xb1

	GetField = 0xb4
	Putfield = 0xb5

	Newarray = 0xbc
	Anewarray = 0xbd

	Invokevirtual = 0xb6
	Invokespecial = 0xb7
	Invokestatic = 0xb8
	Invokeinterface = 0xb9

	New = 0xbb

	Arraylength = 0xbe

	Ireturn = 0xac

	Wide = 0xc4
)
