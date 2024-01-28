// https://adventofcode.com/2019/day/9
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
	RELBASE   int = 9 // relative base
	HALT      int = 99
	POSITION  int = 0
	IMMEDIATE int = 1
	RELATIVE  int = 2
)

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

// return output, pc
func run(program []int, pc int, input []int) (int, int) {

	var ninput, relbase int // Intcode interpreter state
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
			fmt.Println("INPUT", input[ninput])
			program[ins.params[0]] = input[ninput]
			ninput += 1
			pc += 2
		case OUTPUT:
			fmt.Println("OUTPUT", program[ins.params[0]])
			pc += 2
			return program[ins.params[0]], pc
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
			fmt.Println("EXIT")
			return -1, -1
		default:
			fmt.Println("run: unknown opcode", ins.opcode)
			return -1, pc
		}
	}
	return -1, -1
}

func part1() {
	defer duration(time.Now(), "part 1")

	// input is 3403 bytes, 973 []int

	// var test = "109,1,204,-1,1001,100,1,100,1008,100,16,101,1006,101,0,99"
	// var test = "1102,34915192,34915192,7,4,7,99,0"
	// var test = "104,1125899906842624,99"
	var tokens []string = strings.Split(strings.Trim(input, "\n"), ",")
	var masterProgram []int
	for _, tok := range tokens {
		masterProgram = append(masterProgram, atoi(tok))
	}

	// The computer's available memory should be much larger than the initial program.
	// Memory beyond the initial program starts with the value 0
	// and can be read or written like any other memory.
	var program []int = make([]int, len(masterProgram)+100)
	copy(program, masterProgram)
	output, _ := run(program, 0, []int{1})
	fmt.Println("part 1", output)
}

func part2() {
	defer duration(time.Now(), "part 2")

	// input is 3403 bytes, 973 []int

	var tokens []string = strings.Split(strings.Trim(input, "\n"), ",")
	var masterProgram []int
	for _, tok := range tokens {
		masterProgram = append(masterProgram, atoi(tok))
	}

	// The computer's available memory should be much larger than the initial program.
	// Memory beyond the initial program starts with the value 0
	// and can be read or written like any other memory.
	var program []int = make([]int, len(masterProgram)+200)
	copy(program, masterProgram)
	output, _ := run(program, 0, []int{2})
	fmt.Println("part 2", output)
}

func main() {
	defer duration(time.Now(), "main")

	part1()
	part2()
}

/*
$ go run main.go
INPUT 1
OUTPUT 3742852857
part 1 3742852857
part 1 68.165µs
INPUT 2
OUTPUT 73439
part 2 73439
part 2 11.737381ms
main 11.814823ms
*/
