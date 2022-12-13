package nes

const (
	CarryFlag     int = 0000_0001
	ZeroFlag      int = 0000_0010
	InterruptFlag int = 0000_0100
	DecimalFlag   int = 0000_1000
	BreakFlag     int = 0001_0000
	OverflowFlag  int = 0100_0000
	NegativeFlag  int = 1000_0000
)

// CPU type
type CPU struct {
	registerA      int
	registerX      int
	registerY      int
	programCounter int
	stackPointer   uint
	statusRegister int
	memory         []int
}

// CPU instance
var cpu = CPU{
	registerA:      0,
	registerX:      0,
	registerY:      0,
	programCounter: 0,
	stackPointer:   0,
	statusRegister: 0,
	memory:         []int{},
}

func ProgramLoop(program []int) {

	cpu.memory = program

	for cpu.programCounter < len(program) {
		// fetch
		opcode := cpu.memory[cpu.programCounter]
		cpu.programCounter++

		switch opcode {
		case BRK:
			break
		case LDA:
			param := cpu.memory[cpu.programCounter]
			cpu.registerA = param
			cpu.programCounter++

			checkFlag(cpu.registerA, ZeroFlag)
			checkFlag(cpu.registerA, NegativeFlag)
		case INX:
			cpu.registerX++

			checkFlag(cpu.registerX, ZeroFlag)
			checkFlag(cpu.registerX, NegativeFlag)
		}
	}

	// decode
	// execute
}

func checkFlag(register int, flag int) {
	switch flag {
	case ZeroFlag:
		if register == 0 {
			cpu.statusRegister |= ZeroFlag
		} else {
			cpu.statusRegister &= ^ZeroFlag
		}
	case NegativeFlag:
		if register < 0 {
			cpu.statusRegister |= NegativeFlag
		} else {
			cpu.statusRegister &= ^NegativeFlag
		}
	}

	return
}
