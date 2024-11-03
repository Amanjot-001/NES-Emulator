package cpu

// CPU represents the 6502 CPU structure with necessary registers.
type CPU struct {
	PC uint16 // Program Counter
	S  uint8  // Stack Pointer

	A uint8 // Accumulator
	X uint8 // X Register
	Y uint8 // Y Register

	// Status flags (using single byte for simplicity, each bit can represent a flag)
	P uint8
}

func NewCPU() *CPU {
	return &CPU{
		PC: 0xFFFC,
		S:  0xFD,
		P:  0b00100100,
	}
}

func (cpu *CPU) Reset() {
	cpu.PC = 0xFFFC
	cpu.S -= 3
	cpu.P |= 0b00000100
}
