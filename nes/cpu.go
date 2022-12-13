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
			// todo: implement other addressing modes

			// INX implementation
		case 0xE8:
			cpu.RegisterX++

			checkFlag(cpu.RegisterX, ZeroFlag)
			checkFlag(cpu.RegisterX, NegativeFlag)
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
