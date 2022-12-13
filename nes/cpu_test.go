package nes

import "testing"

func TestCPU(t *testing.T) {
	var program = []int{0x00, 0xA9, 0x00, 0xE8}

	ProgramLoop(program)

	if cpu.registerA != 0 {
		t.Errorf("registerA is %d, want 0", cpu.registerA)
	}

	if cpu.registerX != 1 {
		t.Errorf("registerX is %d, want 1", cpu.registerX)
	}

	if cpu.statusRegister != 0 {
		t.Errorf("statusRegister is %d, want 0", cpu.statusRegister)
	}

	println("registerA:", cpu.registerA)
	println("registerX:", cpu.registerX)
	println("statusRegister:", cpu.statusRegister)
}
