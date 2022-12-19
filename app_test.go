package main

import (
	"cpu/nes"
	"fmt"
	"testing"
)

func TestBaseCPU(t *testing.T) {
	// BRK, LDA #0, INX
	var program = []uint8{0x00, 0xA9, 0x00, 0xE8}

	println("----------------------------------")
	println("Program:", program)
	println("Test: TestBaseCPU")

	cpu := nes.ProgramLoop(program)

	fmt.Printf("RegisterA: %d \n", cpu.RegisterA)
	fmt.Printf("RegisterX: %d \n", cpu.RegisterX)
	fmt.Printf("StatusRegister: %b \n", cpu.StatusRegister)

	if cpu.RegisterA != 0 {
		t.Errorf("RegisterA is %d, want 0", cpu.RegisterA)
	}

	if cpu.RegisterX != 1 {
		t.Errorf("RegisterX is %d, want 1", cpu.RegisterX)
	}

	if cpu.StatusRegister&nes.ZeroFlag == nes.ZeroFlag {
		t.Errorf("Expected StatusRegister to not have ZeroFlag")
	}
}

func TestOverflowFlag(t *testing.T) {
	// BRK, LDX #FF, INX
	var program = []uint8{0x00, 0xA2, 0xFF, 0xE8}

	println("----------------------------------")
	println("Program:", program)
	println("Test: TestOverflowFlag")

	cpu := nes.ProgramLoop(program)

	fmt.Printf("RegisterA: %d \n", cpu.RegisterA)
	fmt.Printf("RegisterX: %d \n", cpu.RegisterX)
	fmt.Printf("StatusRegister: %b \n", cpu.StatusRegister)

	if cpu.RegisterA != 0 {
		t.Errorf("RegisterA is %d, want 0", cpu.RegisterA)
	}

	if cpu.RegisterX <= 162 {
		t.Errorf("RegisterX is %d, want > 162", cpu.RegisterX)
	}

	if cpu.StatusRegister&nes.OverflowFlag != nes.OverflowFlag {
		t.Errorf("Expected StatusRegister to have OverflowFlag")
	}

}
