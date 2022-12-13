package main

import "cpu/nes"

func main() {
	println("hello world")

	var program = []uint8{0x0000}

	nes.ProgramLoop(program)
}
