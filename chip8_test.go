package Chip8

import (
	"testing"
)

func TestCls(t *testing.T) {
	// create a new Chip8 instance
	c := newChip8()

	// set some values in the display
	c.display[0][0] = 1
	c.display[1][1] = 1
	c.display[2][2] = 1

	// call the cls method
	c.cls()

	// check that all values in the display are now 0
	for i := 0; i < len(c.display); i++ {
		for j := 0; j < len(c.display[i]); j++ {
			if c.display[i][j] != 0 {
				t.Errorf("Error: display[%d][%d] should be 0, got %d", i, j, c.display[i][j])
			}
		}
	}
}
