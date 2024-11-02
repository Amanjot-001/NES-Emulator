package cpu

// CPU represents the 6502 CPU structure with necessary registers.
type CPU struct {
	PC uint16 // Program Counter
	S  uint8  // Stack Pointer

	A uint8 // Accumulator
	X uint8 // X Register
	Y uint8 // Y Register

	// Status flags (using single byte for simplicity, each bit can represent a flag)
	Status uint8
}

func main() {

}
