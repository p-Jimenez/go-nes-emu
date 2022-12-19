package nes

// Status Register flags
const (
	CarryFlag     uint8 = 0b0000_0001
	ZeroFlag      uint8 = 0b0000_0010
	InterruptFlag uint8 = 0b0000_0100
	DecimalFlag   uint8 = 0b0000_1000
	BreakFlag     uint8 = 0b0001_0000
	OverflowFlag  uint8 = 0b0100_0000
	NegativeFlag  uint8 = 0b1000_0000
)

// CPU type
type CPU struct {
	RegisterA      uint8
	RegisterX      uint8
	RegisterY      uint8
	ProgramCounter uint16
	StackPointer   uint8
	StatusRegister uint8
	memory         []uint8
}

// CPU instance
var cpu = CPU{
	RegisterA:      0,
	RegisterX:      0,
	RegisterY:      0,
	ProgramCounter: 0,
	StackPointer:   0,
	StatusRegister: 0,
	memory:         []uint8{},
}

// ProgramLoop runs the program
func ProgramLoop(program []uint8) CPU {

	cpu = CPU{
		RegisterA:      0,
		RegisterX:      0,
		RegisterY:      0,
		ProgramCounter: 0,
		StackPointer:   0,
		StatusRegister: 0,
		memory:         []uint8{},
	}

	cpu.memory = program

	for cpu.ProgramCounter < uint16(len(program)) {
		// fetch
		opcode := cpu.memory[cpu.ProgramCounter]

		switch opcode {
		// BRK implementation
		case 0x00:
			cpu.StatusRegister |= BreakFlag
			break
		// LDA implementations
		case 0xA9:
			lda(getAddress(Immediate, 0))
		case 0xA5:
			lda(getAddress(ZeroPage, 0))
		case 0xB5:
			lda(getAddress(ZeroPageX, 0))
		case 0xAD:
			lda(getAddress(Absolute, 0))
		case 0xBD:
			lda(getAddress(AbsoluteX, 0))
		case 0xB9:
			lda(getAddress(AbsoluteY, 0))
		case 0xA1:
			lda(getAddress(IndirectX, 0))
		case 0xB1:
			lda(getAddress(IndirectY, 0))
		// INX implementation
		case 0xE8:
			cpu.RegisterX++

			checkFlag(cpu.RegisterX, ZeroFlag)
			checkFlag(cpu.RegisterX, NegativeFlag)
			checkFlag(cpu.RegisterX, OverflowFlag)
		// INY implementation
		case 0xC8:
			cpu.RegisterY++

			checkFlag(cpu.RegisterY, ZeroFlag)
			checkFlag(cpu.RegisterY, NegativeFlag)
		// LDX implementations
		case 0xA2:
			ldx(getAddress(Immediate, 0))
		case 0xA6:
			ldx(getAddress(ZeroPage, 0))
		case 0xB6:
			ldx(getAddress(ZeroPageY, 0))
		case 0xAE:
			ldx(getAddress(Absolute, 0))
		case 0xBE:
			ldx(getAddress(AbsoluteY, 0))
		// NOP implementation
		case 0xEA:
			break
		// LDY implementations
		case 0xA0:
			ldy(getAddress(Immediate, 0))
		case 0xA4:
			ldy(getAddress(ZeroPage, 0))
		case 0xB4:
			ldy(getAddress(ZeroPageX, 0))
		case 0xAC:
			ldy(getAddress(Absolute, 0))
		case 0xBC:
			ldy(getAddress(AbsoluteX, 0))
		// TAX implementation
		case 0xAA:
			cpu.RegisterX = cpu.RegisterA

			checkFlag(cpu.RegisterX, ZeroFlag)
			checkFlag(cpu.RegisterX, NegativeFlag)
		// TAY implementation
		case 0xA8:
			cpu.RegisterY = cpu.RegisterA

			checkFlag(cpu.RegisterY, ZeroFlag)
			checkFlag(cpu.RegisterY, NegativeFlag)
		// TXA implementation
		case 0x8A:
			cpu.RegisterA = cpu.RegisterX

			checkFlag(cpu.RegisterA, ZeroFlag)
			checkFlag(cpu.RegisterA, NegativeFlag)
		// TYA implementation
		case 0x98:
			cpu.RegisterA = cpu.RegisterY

			checkFlag(cpu.RegisterA, ZeroFlag)
			checkFlag(cpu.RegisterA, NegativeFlag)
		// ADC implementations
		case 0x69:
			adc(getAddress(Immediate, 0))
		case 0x65:
			adc(getAddress(ZeroPage, 0))
		case 0x75:
			adc(getAddress(ZeroPageX, 0))
		case 0x6D:
			adc(getAddress(Absolute, 0))
		case 0x7D:
			adc(getAddress(AbsoluteX, 0))
		case 0x79:
			adc(getAddress(AbsoluteY, 0))
		case 0x61:
			adc(getAddress(IndirectX, 0))
		case 0x71:
			adc(getAddress(IndirectY, 0))
		}

		cpu.ProgramCounter++
	}

	// decode
	// execute

	return cpu
}

func lda(address uint16) {
	cpu.RegisterA = cpu.memory[address]
	checkFlag(cpu.RegisterA, ZeroFlag)
	checkFlag(cpu.RegisterA, NegativeFlag)
}

func ldx(address uint16) {
	cpu.RegisterX = cpu.memory[address]
	checkFlag(cpu.RegisterX, ZeroFlag)
	checkFlag(cpu.RegisterX, NegativeFlag)
}

func ldy(address uint16) {
	cpu.RegisterY = cpu.memory[address]
	checkFlag(cpu.RegisterY, ZeroFlag)
	checkFlag(cpu.RegisterY, NegativeFlag)
}

func adc(address uint16) {
	cpu.RegisterA += cpu.memory[address]
	checkFlag(cpu.RegisterA, ZeroFlag)
	checkFlag(cpu.RegisterA, NegativeFlag)
	checkFlag(cpu.RegisterA, CarryFlag)
	checkFlag(cpu.RegisterA, OverflowFlag)
}

// checkFlag checks the status register
func checkFlag(register uint8, flag uint8) {
	switch flag {
	case ZeroFlag:
		if register == 0 {
			cpu.StatusRegister |= ZeroFlag
		} else {
			cpu.StatusRegister &= ^ZeroFlag
		}
	case NegativeFlag:
		if register < 0 {
			cpu.StatusRegister |= NegativeFlag
		} else {
			cpu.StatusRegister &= ^NegativeFlag
		}
	case CarryFlag:
		if register > 255 {
			cpu.StatusRegister |= CarryFlag
		} else {
			cpu.StatusRegister &= ^CarryFlag
		}
	case OverflowFlag:
		// to implement
	default:
		return
	}
}

// read16 reads a 16 bit value from memory
func read16(address uint16) uint16 {
	return uint16(cpu.memory[address]) | uint16(cpu.memory[address+1])<<8
}

// write16 writes a 16 bit value to memory
func write16(address uint16, value uint16) {
	cpu.memory[address] = uint8(value)
	cpu.memory[address+1] = uint8(value >> 8)
}
