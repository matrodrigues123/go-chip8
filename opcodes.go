package Chip8

// 00E0: clear screen
func (c *Chip8) cls() {
	for i := 0; i < len(c.display); i++ {
		for j := 0; j < len(c.display[i]); j++ {
			c.display[i][j] = 0
		}
	}
}

// 00EE: return from a subroutine
func (c *Chip8) ret() {
	c.sp--
	c.pc = c.stack[c.sp]
}

// 1nnn: jump to location nnn
func (c *Chip8) jpAddr() {
	address := getNNN(c.opcode)
	c.pc = address
}

// 2nnn: call subroutine at nnn
func (c *Chip8) callAddr() {
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
func (c *Chip8) seVx() {
	x := getX(c.opcode)
	kk := getKK(c.opcode)

	if c.registers[x] == kk {
		c.pc += 2
	}
}

// 4xkk: The interpreter compares register Vx to kk, and if they are not equal, increments the program counter by 2.
func (c *Chip8) notSeVx() {
	x := getX(c.opcode)
	kk := getKK(c.opcode)

	if c.registers[x] != kk {
		c.pc += 2
	}
}

//5xy0: compare registers Vx and Vy. If they are equal, increment program counter by 2
func (c *Chip8) seVxVy() {
	x := getX(c.opcode)
	y := getY(c.opcode)

	if c.registers[x] == c.registers[y] {
		c.pc += 2
	}
}

// 6xkk: set Vx to kk
func (c *Chip8) ldVx() {
	x := getX(c.opcode)
	kk := getKK(c.opcode)

	c.registers[x] = kk
}

// 7xkk: set Vx to vx + kk
func (c *Chip8) addVx() {
	x := getX(c.opcode)
	kk := getKK(c.opcode)

	c.registers[x] += kk
}

//8xy0: set Vx to Vy
func (c *Chip8) ldVxVy() {
	x := getX(c.opcode)
	y := getY(c.opcode)

	c.registers[x] = c.registers[y]
}

//8xy1: Vx = Vx or Vy
func (c *Chip8) orVxVy() {
	x := getX(c.opcode)
	y := getY(c.opcode)

	c.registers[x] |= c.registers[y]
}

//8xy2: Vx = Vx and Vy
func (c *Chip8) andVxVy() {
	x := getX(c.opcode)
	y := getY(c.opcode)

	c.registers[x] &= c.registers[y]
}

//8xy3: Vx = Vx xor Vy
func (c *Chip8) xorVxVy() {
	x := getX(c.opcode)
	y := getY(c.opcode)

	c.registers[x] ^= c.registers[y]
}

//8xy4: The values of Vx and Vy are added together. If the result is greater than 8 bits (i.e., > 255,) VF is set to 1, otherwise 0. Only the lowest 8 bits of the result are kept, and stored in Vx.
func (c *Chip8) addVxVy() {
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
func (c *Chip8) subVxVy() {
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
func (c *Chip8) shrVx() {
	x := getX(c.opcode)

	// save the lsb of x to VF
	c.registers[0xF] = c.registers[x] & 1

	c.registers[x] = c.registers[x] >> 1
}

// 8xy7: subn Vx, Vy
func (c *Chip8) subnVxVy() {
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
func (c *Chip8) shlVx() {
	x := getX(c.opcode)

	// save the msb of x to VF
	c.registers[0xF] = (c.registers[x] & 0x80) >> 7

	c.registers[x] = c.registers[x] << 1
}

// 9xy0: skip next instruction if Vx != Vy
func (c *Chip8) sneVxVy() {
	x := getX(c.opcode)
	y := getY(c.opcode)

	if c.registers[y] != c.registers[x] {
		c.pc += 2
	}
}

// Annn - set I to nnn
func (c *Chip8) ldIaddr() {
	address := getNNN(c.opcode)
	c.indexRegister = int(address)
}

// Bnnn: jump to location nnn + V0
func (c *Chip8) jpV0Addr() {
	address := getNNN(c.opcode) + c.registers[0]
	c.pc = address
}

// Cxkk: set Vx to random byte and kk
func (c *Chip8) rndVx() {
	x := getX(c.opcode)
	kk := getKK(c.opcode)

	randomByte := make([]byte, 1)
	c.rng.Read(randomByte)
	c.registers[x] = uint16(randomByte[0]) & kk
}

// Dxyn: Display n-byte sprite starting at memory location I at (Vx, Vy), set VF = collision.
func (c *Chip8) DRW() {
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
				if c.display[yPos+row][xPos+col] == 1 {
					c.registers[0xF] = 1
				}
				c.display[yPos+row][xPos+col] ^= 1
			}

		}
	}

}
