package nes

// opcodes
const (
	BRK = 0x00
	LDA = 0xA9
	TAX = 0xAA
	INX = 0xE8
)

type AddressingMode struct {
	bytes  int
	cycles int
}

var (
	Immediate AddressingMode = AddressingMode{bytes: 2, cycles: 2}
	ZeroPage  AddressingMode = AddressingMode{bytes: 2, cycles: 3}
	ZeroPageX AddressingMode = AddressingMode{bytes: 2, cycles: 4}
	Absolute  AddressingMode = AddressingMode{bytes: 3, cycles: 4}
	// +1 cycle if page crossed
	AbsoluteX AddressingMode = AddressingMode{bytes: 3, cycles: 4}
	// +1 cycle if page crossed
	AbsoluteY AddressingMode = AddressingMode{bytes: 3, cycles: 4}
	// +1 cycle if page crossed
	IndirectX AddressingMode = AddressingMode{bytes: 2, cycles: 6}
	// +1 cycle if page crossed
	IndirectY AddressingMode = AddressingMode{bytes: 2, cycles: 5}
)
