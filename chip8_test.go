package Chip8

import (
	"strings"
	"testing"
)

func TestCls003(t *testing.T) {
	// create a new Chip8 instance
	c := newChip8()

	// set some values in the display
	c.display[0][0] = 1
	c.display[1][1] = 1
	c.display[2][2] = 1

	// call the cls method
	c.OP_00E0()

	// check that all values in the display are now 0
	for i := 0; i < len(c.display); i++ {
		for j := 0; j < len(c.display[i]); j++ {
			if c.display[i][j] != 0 {
				t.Errorf("Error: display[%d][%d] should be 0, got %d", i, j, c.display[i][j])
			}
		}
	}
}

func TestDRW(t *testing.T) {
	// Initialize chip8
	c := Chip8{}
	c.registers[0] = 0                     // Set V0 to 0
	c.registers[1] = 1                     // Set V1 to 1
	c.indexRegister = 0x200                // Set index register to start of program memory
	c.memory[c.indexRegister] = 0b11110000 // sprite data at memory location I
	c.memory[c.indexRegister+1] = 0b00001111

	// Set display pixels to 1 at (1,1) and (2,1)
	c.display[1][1] = 1
	c.display[1][2] = 1

	// Execute DRW instruction
	c.opcode = 0xD012
	c.OP_Dxyn()

	// Check that the display is updated correctly
	expectedDisplay := [64][32]byte{}
	expectedDisplay[1][1] = 0
	expectedDisplay[1][2] = 0
	expectedDisplay[2][1] = 1
	expectedDisplay[2][2] = 1

	for y := 0; y < 32; y++ {
		for x := 0; x < 64; x++ {
			if c.display[x][y] != expectedDisplay[x][y] {
				t.Errorf("Expected display:\n%v\n\nActual display:\n%v\n\nat x=%d, y=%d",
					formatDisplay(expectedDisplay), formatDisplay(c.display), x, y)
				return
			}
		}
	}

}

func formatDisplay(display [64][32]byte) string {
	var b strings.Builder
	for y := 0; y < 32; y++ {
		for x := 0; x < 64; x++ {
			if display[x][y] == 1 {
				b.WriteString("*")
			} else {
				b.WriteString(" ")
			}
		}
		b.WriteString("\n")
	}
	return b.String()
}
