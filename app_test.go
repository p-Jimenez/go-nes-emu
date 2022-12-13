package main

import (
	"cpu/nes"
	"testing"
)

func TestCPU(t *testing.T) {
	var program = []uint8{0x00, 0xA9, 0x00, 0xE8}

	cpu := nes.ProgramLoop(program)

	if cpu.RegisterA != 0 {
		t.Errorf("RegisterA is %d, want 0", cpu.RegisterA)
	}

	if cpu.RegisterX != 1 {
		t.Errorf("RegisterX is %d, want 1", cpu.RegisterX)
	}

	if cpu.StatusRegister != 0 {
		t.Errorf("StatusRegister is %b, want 0", cpu.StatusRegister)
	}

	println("RegisterA:", cpu.RegisterA)
	println("RegisterX:", cpu.RegisterX)
	println("StatusRegister:", cpu.StatusRegister)
}
