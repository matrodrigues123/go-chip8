package Chip8

import (
	"math/rand"
	"time"
)

type Chip8 struct {
	memory        [4096]byte
	registers     [16]uint16
	indexRegister int
	pc            uint16
	stack         [16]uint16
	sp            byte
	opcode        uint16
	delayTimer    byte
	soundTimer    byte
	display       [64][32]byte
	keypad        [16]bool
	rng           *rand.Rand
}

func newChip8() *Chip8 {
	// alocate memory for the new chip8 instance and initialize all its fields to their default values
	chip8 := &Chip8{}

	// initialize pc to the start adress
	chip8.pc = START_ADRESS

	// load fonts into memory
	var i uint16
	for i = 0; i < FONTSET_SIZE; i++ {
		chip8.memory[FONTSET_START_ADRESS+i] = fontset[i]
	}

	// initialize the random number generator with a seed
	chip8.rng = rand.New(rand.NewSource(time.Now().UnixNano()))

	return chip8
}

func (c *Chip8) emulateCycle() {
	// Fetch Opcode
	// Decode Opcode
	// Execute Opcode

	// Update timers
}
