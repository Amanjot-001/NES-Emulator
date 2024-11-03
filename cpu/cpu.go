package cpu

// CPU represents the 6502 structure
type CPU struct {
	PC uint16 // Program Counter
	S  uint8  // Stack Pointer

	A uint8 // Accumulator
	X uint8 // X Register
	Y uint8 // Y Register

	// Status flags (each bit can represent a flag)
	P uint8
}

func NewCPU() *CPU {
	return &CPU{
		PC: 0xFFFC,     // PC points to reset vector location
		S:  0xFD,       // S points to $FD
		P:  0b00100100, // set I flag to 1 with others as 0 and 3rd bit is always 1
	}
}

func (cpu *CPU) Reset() {
	cpu.PC = 0xFFFC     // re-init to $FFFC
	cpu.S -= 3          // S decrements by 3 as per behaviour
	cpu.P |= 0b00000100 // reset sets I flag to 1 (OR to achieve this)
}
