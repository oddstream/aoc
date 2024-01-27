// https://adventofcode.com/2019/day/7
package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
	"time"
)

//go:embed test1.txt
var test1 string

//go:embed test2.txt
var test2 string

//go:embed test3.txt
var test3 string

//go:embed test4.txt
var test4 string

//go:embed test5.txt
var test5 string

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

// Heap's algorithm
func permutations(arr []int) [][]int {
	var helper func([]int, int)
	res := [][]int{}

	helper = func(arr []int, n int) {
		if n == 1 {
			tmp := make([]int, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
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

// return output, pc
func run(program []int, pc int, input []int) (int, int) {

	var ninput int
	for {
		ins := readInstruction(program, pc)
		switch ins.opcode {
		case ADD:
			program[ins.params[2]] = program[ins.params[0]] + program[ins.params[1]]
			pc += 4
		case MUL:
			program[ins.params[2]] = program[ins.params[0]] * program[ins.params[1]]
			pc += 4
		case INPUT:
			// fmt.Println("INPUT", input[ninput])
			program[ins.params[0]] = input[ninput]
			ninput += 1
			pc += 2
		case OUTPUT:
			// fmt.Println("OUTPUT", program[ins.params[0]])
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
		case HALT:
			// fmt.Println("EXIT")
			return -1, -1
		default:
			fmt.Println("unknown opcode", ins.opcode)
			return -1, pc
		}
	}
	return -1, -1
}

func part1() {
	defer duration(time.Now(), "part 1")

	var tokens []string = strings.Split(strings.Trim(input, "\n"), ",")
	var masterProgram []int
	for _, tok := range tokens {
		masterProgram = append(masterProgram, atoi(tok))
	}

	// var seq []int = []int{4, 3, 2, 1, 0}	// test1, 43210
	// var seq []int = []int{0, 1, 2, 3, 4} // test2, 54321
	// var seq []int = []int{1, 0, 4, 3, 2} // test3, 65210
	var result int
	var perms = permutations([]int{0, 1, 2, 3, 4})
	for _, perm := range perms {
		var output int
		for _, phase := range perm {
			var program []int = make([]int, len(masterProgram))
			copy(program, masterProgram)
			// "When a copy of the program starts running on an amplifier,
			// it will first use an input instruction to ask the amplifier
			// for its current phase setting (an integer from 0 to 4).
			// Each phase setting is used exactly once"
			//
			// "The program will then call another input instruction
			// to get the amplifier's input signal,"
			//
			// "The first amplifier's input value is 0"
			output, _ = run(program, 0, []int{phase, output})
			if output == -1 {
				fmt.Print("error break")
				break
			}
			result = max(result, output)
		}
	}
	fmt.Println("part 1", result) // 46248
}

func part2() {
	defer duration(time.Now(), "part 2")

	var tokens []string = strings.Split(strings.Trim(input, "\n"), ",")
	var masterProgram []int
	for _, tok := range tokens {
		masterProgram = append(masterProgram, atoi(tok))
	}

	// perms = permutations([]int{5, 6, 7, 8, 9})
	// var phases []int = []int{9, 8, 7, 6, 5} // test4, 139629729, loop 4
	// var phases []int = []int{9, 7, 8, 5, 6} // test5, 18216

	// To start the process, a 0 signal is sent to amplifier A's input exactly once.

	// The phase is only fed to the Amps exactly once (i.e. when the Amps are first "initialised" or "started").
	// When an individual Amp meets opcode 4 (output), they output a signal to be taken in by the next amp,
	// and then they PAUSE EXECUTION to be resumed again when the loop goes back to the same particular amp
	// (that is, you have to keep track of the instruction pointer for each individual amp).
	var result int
	for _, phases := range permutations([]int{5, 6, 7, 8, 9}) {
		var amplifiers [5][]int
		for i := 0; i < 5; i++ {
			amplifiers[i] = make([]int, len(masterProgram))
			copy(amplifiers[i], masterProgram)
		}
		var pcs [5]int
		var output, lastOutput int
		for loop := 0; loop < 100; loop++ {
			for i, phase := range phases {
				if loop == 0 {
					output, pcs[i] = run(amplifiers[i], pcs[i], []int{phase, output})
				} else {
					output, pcs[i] = run(amplifiers[i], pcs[i], []int{output})
				}
				// fmt.Println(loop, i, amplifiers[i]) // check that amplifiers[i] has changed
				if pcs[i] == -1 {
					output = lastOutput
					goto exitloop
				}
				if output == -1 {
					fmt.Println("ERROR in amp", i, "output -1")
					goto exitloop
				}
			}
			// fmt.Println(loop, output)
			if output == -1 || output == lastOutput {
				break
			}
			lastOutput = output
		}
	exitloop:
		result = max(result, output)
	}
	fmt.Println("part 2", result)
}

func main() {
	defer duration(time.Now(), "main")

	part1() // 46248
	part2() // 4163586
}

/*
$ go run main.go
part 1 46248
part 1 998.529µs
part 2 54163586
part 2 1.948722ms
main 2.95669ms
*/
