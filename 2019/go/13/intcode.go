package main

import "fmt"

const (
	ADD       int = 1
	MUL       int = 2
	INPUT     int = 3 // eg 3,50 takes an input value and stores it at address 50
	OUTPUT    int = 4 // eg 4,50 would output the value at address 50
	JIT       int = 5 // jump if true
	JIF       int = 6 // jump if false
	LT        int = 7 // less than
	EQ        int = 8 // equals
	RELBASE   int = 9 // relative base
	HALT      int = 99
	POSITION  int = 0
	IMMEDIATE int = 1
	RELATIVE  int = 2
	RAMSIZE   int = 1024
)

type Interpreter struct {
	memory  []int
	pc      int
	relbase int
	input   chan int
	output  chan int
}

func NewInterpreter(memory []int) *Interpreter {
	return &Interpreter{
		memory: memory,
		input:  make(chan int, 2),
		output: make(chan int, 2),
	}
}

// given an opcode,
// gives number of ints this opcode will consume,
// not including the opcode itself
var arity map[int]int = map[int]int{
	ADD:     3,
	MUL:     3,
	INPUT:   1,
	OUTPUT:  1,
	JIT:     2,
	JIF:     2,
	LT:      3,
	EQ:      3,
	RELBASE: 1,
	HALT:    0,
}

type Instruction struct {
	opcode int    // one of ADD, MUL, INPUT, OUTPUT, HALT
	params [3]int // 0..3, depending on opcode
}

func readInstruction(program []int, pc int, relbase int) Instruction {
	var ins Instruction
	var n = program[pc]
	// opcode is lowest two decimal digits
	ins.opcode = n % 100
	n /= 100

	if _, ok := arity[ins.opcode]; !ok {
		fmt.Println("ERROR: unknown opcode", ins.opcode, "at", pc)
	}

	for i := 0; i < arity[ins.opcode]; i++ {
		mode := n % 10
		n /= 10
		switch mode {
		case IMMEDIATE:
			// "a parameter is interpreted as a value
			// if the parameter is 50, its value is simply 50"
			ins.params[i] = pc + 1 + i
		case POSITION:
			// "the parameter to be interpreted as a position
			// if the parameter is 50, its value is the value stored at address 50 in memory"
			ins.params[i] = program[pc+1+i]
		case RELATIVE:
			// ins.params[i] = program[relbase+pc+1+i]	// gave 203 for part 1 (wrong, too low)
			ins.params[i] = relbase + program[pc+1+i] // gave 3742852857, just doesn't look right
		default:
			fmt.Println("ERROR: unknown mode", mode)
			return Instruction{opcode: HALT}
		}
	}

	return ins
}

func intcode(program []int, in func() int, out func(int)) {
	var pc, relbase int // Intcode interpreter state
	for {
		ins := readInstruction(program, pc, relbase)
		switch ins.opcode {
		case ADD:
			program[ins.params[2]] = program[ins.params[0]] + program[ins.params[1]]
			pc += 4
		case MUL:
			program[ins.params[2]] = program[ins.params[0]] * program[ins.params[1]]
			pc += 4
		case INPUT:
			// fmt.Println("INPUT")
			program[ins.params[0]] = in()
			pc += 2
		case OUTPUT:
			// fmt.Println("OUTPUT")
			pc += 2
			out(program[ins.params[0]])
		case JIT:
			if program[ins.params[0]] != 0 {
				pc = program[ins.params[1]]
			} else {
				pc += 3
			}
		case JIF:
			if program[ins.params[0]] == 0 {
				pc = program[ins.params[1]]
			} else {
				pc += 3
			}
		case LT:
			if program[ins.params[0]] < program[ins.params[1]] {
				program[ins.params[2]] = 1
			} else {
				program[ins.params[2]] = 0
			}
			pc += 4
		case EQ:
			if program[ins.params[0]] == program[ins.params[1]] {
				program[ins.params[2]] = 1
			} else {
				program[ins.params[2]] = 0
			}
			pc += 4
		case RELBASE:
			relbase += program[ins.params[0]]
			pc += 2
		case HALT:
			// fmt.Println("EXIT")
			return
		default:
			fmt.Println("run: unknown opcode", ins.opcode)
			return
		}
	}
}

func intcode2(program []int, in <-chan int, out chan<- int) {
	var pc, relbase int // Intcode interpreter state
	for {
		ins := readInstruction(program, pc, relbase)
		switch ins.opcode {
		case ADD:
			program[ins.params[2]] = program[ins.params[0]] + program[ins.params[1]]
			pc += 4
		case MUL:
			program[ins.params[2]] = program[ins.params[0]] * program[ins.params[1]]
			pc += 4
		case INPUT:
			// fmt.Println("INPUT")
			program[ins.params[0]] = <-in
			pc += 2
		case OUTPUT:
			// fmt.Println("OUTPUT")
			pc += 2
			out <- program[ins.params[0]]
		case JIT:
			if program[ins.params[0]] != 0 {
				pc = program[ins.params[1]]
			} else {
				pc += 3
			}
		case JIF:
			if program[ins.params[0]] == 0 {
				pc = program[ins.params[1]]
			} else {
				pc += 3
			}
		case LT:
			if program[ins.params[0]] < program[ins.params[1]] {
				program[ins.params[2]] = 1
			} else {
				program[ins.params[2]] = 0
			}
			pc += 4
		case EQ:
			if program[ins.params[0]] == program[ins.params[1]] {
				program[ins.params[2]] = 1
			} else {
				program[ins.params[2]] = 0
			}
			pc += 4
		case RELBASE:
			relbase += program[ins.params[0]]
			pc += 2
		case HALT:
			// fmt.Println("EXIT")
			return
		default:
			fmt.Println("run: unknown opcode", ins.opcode)
			return
		}
	}
}
