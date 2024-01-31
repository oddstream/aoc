// https://adventofcode.com/2019/day/13
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
	EMPTY  int = 0
	WALL   int = 1
	BLOCK  int = 2
	PADDLE int = 3
	BALL   int = 4
)

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

type Point struct {
	x, y int
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
func run(program []int, tiles map[Point]int) {

	findObjectX := func(obj int) int {
		for p, o := range tiles {
			if o == obj {
				return p.x
			}
		}
		return 0
	}

	var xytile []int    // OUTPUT will fill this in
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
			// get ballX and paddleX dynamically in case program has moved them
			ballX := findObjectX(BALL)
			paddleX := findObjectX(PADDLE)
			if paddleX > ballX {
				program[ins.params[0]] = -1
			} else if paddleX < ballX {
				program[ins.params[0]] = 1
			} else {
				program[ins.params[0]] = 0
			}
			pc += 2
		case OUTPUT:
			pc += 2
			xytile = append(xytile, program[ins.params[0]])
			if len(xytile) == 3 {
				if xytile[0] == -1 && xytile[1] == 0 {
					fmt.Println("SCORE", xytile[2])
				} else {
					tiles[Point{x: xytile[0], y: xytile[1]}] = xytile[2]
				}
				xytile = nil
			}
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
			return
		default:
			fmt.Println("run: unknown opcode", ins.opcode)
			return
		}
	}
}

func display(tiles map[Point]int) {
	for y := -1; y < 25; y++ {
		for x := -1; x < 50; x++ {
			switch tiles[Point{y: y, x: x}] {
			case EMPTY:
				fmt.Print(" ")
			case WALL:
				fmt.Print("#")
			case BLOCK:
				fmt.Print(".")
			case PADDLE:
				fmt.Print("-")
			case BALL:
				fmt.Print("o")
			}
		}
		fmt.Println()
	}
}

func part1() {
	defer duration(time.Now(), "part 1")

	var tokens []string = strings.Split(strings.Trim(input, "\n"), ",")
	var masterProgram []int
	for _, tok := range tokens {
		masterProgram = append(masterProgram, atoi(tok))
	}

	var tiles map[Point]int = make(map[Point]int)

	// The computer's available memory should be much larger than the initial program.
	// Memory beyond the initial program starts with the value 0
	// and can be read or written like any other memory.
	var program []int = make([]int, len(masterProgram)+1000)
	copy(program, masterProgram)
	run(program, tiles)
	fmt.Println(len(tiles), "tiles")
	var result int
	for _, v := range tiles {
		if v == BLOCK {
			result += 1
		}
	}
	display(tiles)
	fmt.Println("part 1", result) // 329
}

func part2() {
	defer duration(time.Now(), "part 2")

	var tokens []string = strings.Split(strings.Trim(input, "\n"), ",")
	var masterProgram []int
	for _, tok := range tokens {
		masterProgram = append(masterProgram, atoi(tok))
	}

	var tiles map[Point]int = make(map[Point]int)

	// The computer's available memory should be much larger than the initial program.
	// Memory beyond the initial program starts with the value 0
	// and can be read or written like any other memory.
	var program []int = make([]int, len(masterProgram)+1000)
	copy(program, masterProgram)
	program[0] = 2
	run(program, tiles)
}

func main() {
	defer duration(time.Now(), "main")

	// TODO rewrite run() to accept INPUT from a chan and send OUTPUT to a chan?
	part1() // 329
	part2() // 15973
}

/*
$ go run main.go
*/
