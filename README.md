# Chip8 Interpreter in Go

This is a project for a Chip8 interpreter written in Go. The Chip8 is a simple virtual machine designed in the 1970s to run games and other applications on early microcomputers.

The Chip8 system has a very basic hardware configuration, with only 4KB of memory, 16 general-purpose registers, a 16-bit address register, a program counter, a stack, and a simple display consisting of 64x32 pixels. It also has a 16-key hexadecimal keypad.

Chip8 programs are composed of instructions called opcodes, which are 16 bits long. There are 35 different opcodes in total, each of which represents a specific operation that the Chip8 system can perform. These operations can include basic arithmetic, branching, input/output, and more. The interpreter must be able to correctly decode and execute each of these opcodes in order to properly run Chip8 programs.

**Note: This project is currently a work in progress and is not yet ready for use.**



