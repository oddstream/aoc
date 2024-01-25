// https://adventofcode.com/2019/day/5 Sunny with a Chance of Asteroids
package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
	"time"
)

//go:embed input.txt
var input string

const (
	ADD       int = 1
	MUL       int = 2
	INPUT     int = 3 // eg 3,50 takes an input value (from where?) and store it at address 50
	OUTPUT    int = 4 // eg 4,50 would output the value (to where?) at address 50
	JIT       int = 5 // jump if true
	JIF       int = 6 // jump if false
	LT        int = 7 // less than
	EQ        int = 8 // equals
	HALT      int = 99
	POSITION  int = 0
	IMMEDIATE int = 1
)

// given an opcode, gives number of ints this opcode will consume
var nparams map[int]int = map[int]int{
	ADD:    3,
	MUL:    3,
	INPUT:  1,
	OUTPUT: 1,
	JIT:    2,
	JIF:    2,
	LT:     3,
	EQ:     3,
	HALT:   0,
}

type Instruction struct {
	opcode int    // one of ADD, MUL, INPUT, OUTPUT, HALT
	params [3]int // 0..3, depending on opcode
}

func duration(invocation time.Time, name string) {
	fmt.Println(name, time.Since(invocation))
}

// or we could use K&R p61
func atoi(s string) int {
	if n, err := strconv.Atoi(s); err == nil {
		return n
	} else {
		fmt.Println(err)
	}
	return 0
}

func readInstruction(program []int, pc int) Instruction {
	var ins Instruction
	var n = program[pc]
	// opcode is lowest two decimal digits
	ins.opcode = n % 100
	n /= 100

	if _, ok := nparams[ins.opcode]; !ok {
		fmt.Println("ERROR: unknown opcode", ins.opcode, "at", pc)
	}

	for i := 0; i < nparams[ins.opcode]; i++ {
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
		default:
			fmt.Println("ERROR: unknown mode", mode)
			return Instruction{opcode: HALT}
		}
	}

	return ins
}

func partOne(program []int, input int) {
	defer duration(time.Now(), "part 1")
	var pc int
	for pc < len(program) {
		ins := readInstruction(program, pc)
		switch ins.opcode {
		case ADD:
			program[ins.params[2]] = program[ins.params[0]] + program[ins.params[1]]
			pc += 4
		case MUL:
			program[ins.params[2]] = program[ins.params[0]] * program[ins.params[1]]
			pc += 4
		case INPUT:
			program[ins.params[0]] = input
			pc += 2
		case OUTPUT:
			fmt.Println("OUTPUT", program[ins.params[0]]) // produce lots of 0 (tests passed) then result
			pc += 2
		case HALT:
			fmt.Println("EXIT")
			pc += 1
			return
		default:
			fmt.Println("unknown opcode", ins.opcode)
		}
	}
}

func partTwo(program []int, input int) {
	defer duration(time.Now(), "part 2")
	var pc int
	for pc < len(program) {
		ins := readInstruction(program, pc)
		switch ins.opcode {
		case ADD:
			program[ins.params[2]] = program[ins.params[0]] + program[ins.params[1]]
			pc += 4
		case MUL:
			program[ins.params[2]] = program[ins.params[0]] * program[ins.params[1]]
			pc += 4
		case INPUT:
			program[ins.params[0]] = input
			pc += 2
		case OUTPUT:
			fmt.Println("OUTPUT", program[ins.params[0]])
			pc += 2
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
		case HALT:
			fmt.Println("EXIT")
			pc += 1
			return
		default:
			fmt.Println("unknown opcode", ins.opcode)
		}
	}
}

func main() {
	defer duration(time.Now(), "main")

	var tokens []string = strings.Split(strings.Trim(input, "\n"), ",")
	var masterProgram []int
	for _, tok := range tokens {
		masterProgram = append(masterProgram, atoi(tok))
	}

	var program []int = make([]int, len(masterProgram))
	copy(program, masterProgram)

	partOne(program, 1) // 13294380

	program = make([]int, len(masterProgram))
	copy(program, masterProgram)

	partTwo(program, 5) // 11460760
}

/*
$ go run main.go
OUTPUT 0
OUTPUT 0
OUTPUT 0
OUTPUT 0
OUTPUT 0
OUTPUT 0
OUTPUT 0
OUTPUT 0
OUTPUT 0
OUTPUT 13294380
EXIT
part 1 27.955µs
OUTPUT 11460760
EXIT
part 2 9.048µs
main 153.793µs
*/
