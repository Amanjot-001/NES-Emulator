package main

import (
	"fmt"
	"goNESis/cpu"
)

func main() {
	cpuInstance := cpu.NewCPU()

	// Power-up state for CPU
	fmt.Printf("Power-Up State:\n PC: %04X\n S: %02X\n A: %02X\n X: %02X\n Y: %02X\n Status: %08b\n",
		cpuInstance.PC, cpuInstance.S, cpuInstance.A, cpuInstance.X, cpuInstance.Y, cpuInstance.P)

	// Simulate a reset and print CPU state again
	cpuInstance.Reset()
	fmt.Printf("After Reset:\n PC: %04X\n S: %02X\n A: %02X\n X: %02X\n Y: %02X\n Status: %08b\n",
		cpuInstance.PC, cpuInstance.S, cpuInstance.A, cpuInstance.X, cpuInstance.Y, cpuInstance.P)
}
