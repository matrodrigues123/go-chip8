package Chip8

// 00E0: clear screen
func (c *Chip8) OP_00E0() {
	for i := 0; i < len(c.display); i++ {
		for j := 0; j < len(c.display[i]); j++ {
			c.display[i][j] = 0
		}
	}
}

// 00EE: return from a subroutine
func (c *Chip8) OP_00EE() {
	c.sp--
	c.pc = c.stack[c.sp]
}

// 1nnn: jump to location nnn
func (c *Chip8) OP_1nnn() {
	address := getNNN(c.opcode)
	c.pc = address
}

// 2nnn: call subroutine at nnn
func (c *Chip8) OP_2nnn() {
	// get the address, ex: CALL $208 (get the 208)
	address := getNNN(c.opcode)

	// put the previous PC at the top of the stack
	// and increase the stack pointer
	c.stack[c.sp] = c.pc
	c.sp++

	// update PC to 208
	c.pc = address

}

// 3xkk: The interpreter compares register Vx to kk, and if they are equal, increments the program counter by 2.
func (c *Chip8) OP_3xkk() {
	x := getX(c.opcode)
	kk := getKK(c.opcode)

	if c.registers[x] == kk {
		c.pc += 2
	}
}

// 4xkk: The interpreter compares register Vx to kk, and if they are not equal, increments the program counter by 2.
func (c *Chip8) OP_4xkk() {
	x := getX(c.opcode)
	kk := getKK(c.opcode)

	if c.registers[x] != kk {
		c.pc += 2
	}
}

//5xy0: compare registers Vx and Vy. If they are equal, increment program counter by 2
func (c *Chip8) OP_5xy0() {
	x := getX(c.opcode)
	y := getY(c.opcode)

	if c.registers[x] == c.registers[y] {
		c.pc += 2
	}
}

// 6xkk: set Vx to kk
func (c *Chip8) OP_6xkk() {
	x := getX(c.opcode)
	kk := getKK(c.opcode)

	c.registers[x] = kk
}

// 7xkk: set Vx to vx + kk
func (c *Chip8) OP_7xkk() {
	x := getX(c.opcode)
	kk := getKK(c.opcode)

	c.registers[x] += kk
}

//8xy0: set Vx to Vy
func (c *Chip8) OP_8xy0() {
	x := getX(c.opcode)
	y := getY(c.opcode)

	c.registers[x] = c.registers[y]
}

//8xy1: Vx = Vx or Vy
func (c *Chip8) OP_8xy1() {
	x := getX(c.opcode)
	y := getY(c.opcode)

	c.registers[x] |= c.registers[y]
}

//8xy2: Vx = Vx and Vy
func (c *Chip8) OP_8xy2() {
	x := getX(c.opcode)
	y := getY(c.opcode)

	c.registers[x] &= c.registers[y]
}

//8xy3: Vx = Vx xor Vy
func (c *Chip8) OP_8xy3() {
	x := getX(c.opcode)
	y := getY(c.opcode)

	c.registers[x] ^= c.registers[y]
}

//8xy4: The values of Vx and Vy are added together. If the result is greater than 8 bits (i.e., > 255,) VF is set to 1, otherwise 0. Only the lowest 8 bits of the result are kept, and stored in Vx.
func (c *Chip8) OP_8xy4() {
	x := getX(c.opcode)
	y := getY(c.opcode)

	sumRes := c.registers[x] + c.registers[y]

	if sumRes > 255 {
		c.registers[0xF] = 1
	} else {
		c.registers[0xF] = 0
	}

	// only the lowest 8 bit of the result are kept
	c.registers[x] = sumRes & 0xFF
}

//8xy5: If Vx > Vy, then VF is set to 1, otherwise 0. Then Vy is subtracted from Vx, and the results stored in Vx.
func (c *Chip8) OP_8xy5() {
	x := getX(c.opcode)
	y := getY(c.opcode)

	if c.registers[x] > c.registers[y] {
		c.registers[0xF] = 1
	} else {
		c.registers[0xF] = 0
	}

	c.registers[x] -= c.registers[y]
}

// 8xy6: shr Vx
func (c *Chip8) OP_8xy6() {
	x := getX(c.opcode)

	// save the lsb of x to VF
	c.registers[0xF] = c.registers[x] & 1

	c.registers[x] = c.registers[x] >> 1
}

// 8xy7: subn Vx, Vy
func (c *Chip8) OP_8xy7() {
	x := getX(c.opcode)
	y := getY(c.opcode)

	if c.registers[y] > c.registers[x] {
		c.registers[0xF] = 1
	} else {
		c.registers[0xF] = 0
	}

	c.registers[x] = c.registers[y] - c.registers[x]
}

// 8xyE: shl Vx
func (c *Chip8) OP_8xyE() {
	x := getX(c.opcode)

	// save the msb of x to VF
	c.registers[0xF] = (c.registers[x] & 0x80) >> 7

	c.registers[x] = c.registers[x] << 1
}

// 9xy0: skip next instruction if Vx != Vy
func (c *Chip8) OP_9xy0() {
	x := getX(c.opcode)
	y := getY(c.opcode)

	if c.registers[y] != c.registers[x] {
		c.pc += 2
	}
}

// Annn - set I to nnn
func (c *Chip8) OP_Annn() {
	address := getNNN(c.opcode)
	c.indexRegister = int(address)
}

// Bnnn: jump to location nnn + V0
func (c *Chip8) OP_Bnnn() {
	address := getNNN(c.opcode) + c.registers[0]
	c.pc = address
}

// Cxkk: set Vx to random byte and kk
func (c *Chip8) OP_Cxkk() {
	x := getX(c.opcode)
	kk := getKK(c.opcode)

	randomByte := make([]byte, 1)
	c.rng.Read(randomByte)
	c.registers[x] = uint16(randomByte[0]) & kk
}

// Dxyn: Display n-byte sprite starting at memory location I at (Vx, Vy), set VF = collision.
func (c *Chip8) OP_Dxyn() {
	x := getX(c.opcode)
	y := getY(c.opcode)
	n := getN(c.opcode)

	xPos, yPos := int(c.registers[x]), int(c.registers[y])

	c.registers[0xF] = 0
	for row := 0; row < int(n); row++ {
		spriteByte := c.memory[c.indexRegister+row]
		for col := 0; col < 8; col++ {
			spritePixel := spriteByte & (0x80 >> col)
			// check if sprite pixel is on
			if spritePixel != 0 {
				// check if theres is already a 'on' pixel in the screen
				if c.display[xPos+col][yPos+row] == 1 {
					c.registers[0xF] = 1
				}
				c.display[xPos+col][yPos+row] ^= 1
			}

		}
	}

}

// Ex9E: Skip next instruction if key with the value of register x is pressed
func (c *Chip8) OP_Ex9E() {
	x := getX(c.opcode)
	key := int(c.registers[x])

	if c.keypad[key] {
		c.pc += 2
	}
}

// ExA1: Skip next instruction if key with the value of register x is not pressed
func (c *Chip8) OP_ExA1() {
	x := getX(c.opcode)
	key := int(c.registers[x])

	if !c.keypad[key] {
		c.pc += 2
	}
}

// Fx07: set register x = delay timer value
func (c *Chip8) OP_Fx07() {
	x := getX(c.opcode)
	c.registers[x] = uint16(c.delayTimer)
}

// Fx0A: Wait for a key press, store the value of the key in Vx
func (c *Chip8) OP_Fx0A() {
	x := getX(c.opcode)
	for i := 0; i < len(c.keypad); i++ {
		if c.keypad[i] {
			c.registers[x] = uint16(i)
			return
		}
	}
	c.sp -= 2
}

// Fx15: set delay timer = register x
func (c *Chip8) OP_Fx15() {
	x := getX(c.opcode)
	c.delayTimer = byte(c.registers[x])
}

// Fx18: set sound timer = register x
func (c *Chip8) OP_Fx18() {
	x := getX(c.opcode)
	c.soundTimer = byte(c.registers[x])
}

// Fx1E: set I = I + Vx
func (c *Chip8) OP_Fx1E() {
	x := getX(c.opcode)
	c.indexRegister += int(c.registers[x])
}

// Fx29: set I = location of sprite for digit Vx
func (c *Chip8) OP_Fx29() {
	x := getX(c.opcode)
	digit := int(c.registers[x])
	c.indexRegister = int(FONTSET_START_ADRESS) + (5 * digit)
}

// Fx33: store BCD representation of Vx in memory locations I, I+1, I+2
func (c *Chip8) OP_Fx33() {
	x := getX(c.opcode)
	value := int(c.registers[x])

	// ones place
	c.memory[c.indexRegister+2] = byte(value % 10)
	value = value / 10

	// tens place
	c.memory[c.indexRegister+1] = byte(value % 10)
	value = value / 10

	// hundreds place
	c.memory[c.indexRegister] = byte(value % 10)
}

// Fx55: store registers V0 through Vx in memory starting at location I
func (c *Chip8) OP_Fx55() {
	x := int(getX(c.opcode))

	for i := 0; i <= x; i++ {
		c.memory[c.indexRegister+i] = byte(c.registers[i])
	}
}

// Fx65: read registers V0 through Vx from memory starting at location I
func (c *Chip8) OP_Fx65() {
	x := int(getX(c.opcode))

	for i := 0; i <= x; i++ {
		c.registers[i] = uint16(c.memory[c.indexRegister+i])
	}
}
