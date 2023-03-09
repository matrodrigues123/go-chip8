package Chip8

func getNNN(opcode uint16) uint16 {
	return opcode & 0x0FFF
}

func getN(opcode uint16) uint16 {
	return opcode & 0x000F
}

func getX(opcode uint16) uint16 {
	return (opcode & 0x0F00) >> 8
}

func getY(opcode uint16) uint16 {
	return (opcode & 0x00F0) >> 4
}

func getKK(opcode uint16) uint16 {
	return opcode & 0x00FF
}
