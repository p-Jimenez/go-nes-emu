package nes

// opcodes

type AddressingMode struct {
	bytes  int
	cycles int
}

var (
	Immediate AddressingMode = AddressingMode{bytes: 2, cycles: 2}
	ZeroPage  AddressingMode = AddressingMode{bytes: 2, cycles: 3}
	ZeroPageX AddressingMode = AddressingMode{bytes: 2, cycles: 4}
	ZeroPageY AddressingMode = AddressingMode{bytes: 2, cycles: 4}
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

func getAddress(mode AddressingMode, address uint16) uint16 {
	switch mode {
	case Immediate:
		cpu.ProgramCounter++ // skip the address
		return cpu.ProgramCounter
	case ZeroPage:
		return uint16(cpu.memory[address])
	case ZeroPageX:
		return uint16(cpu.memory[address]) + uint16(cpu.RegisterX)
	case ZeroPageY:
		return uint16(cpu.memory[address]) + uint16(cpu.RegisterY)
	case Absolute:
		return read16(cpu.ProgramCounter)
	case AbsoluteX:
		return read16(cpu.ProgramCounter) + uint16(cpu.RegisterX)
	case AbsoluteY:
		return read16(cpu.ProgramCounter) + uint16(cpu.RegisterY)
	case IndirectX:
		return read16(uint16(cpu.memory[address]) + uint16(cpu.RegisterX))
	case IndirectY:
		return read16(uint16(cpu.memory[address])) + uint16(cpu.RegisterY)
	default:
		return 0
	}
}
