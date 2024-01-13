// https://adventofcode.com/2017/day/25 The Halting Problem
package main

import (
	"fmt"
	"time"
)

type SubInstruction struct {
	write, move int
	next        string
}

type Instruction struct {
	instructions [2]SubInstruction
}

var (
	pos  int
	tape map[int]int = make(map[int]int)
)

// input has states A .. F
// only ever move one slot at a time
// at this point, I figured it would be quicker to parse the input by hand...

// var (
// 	state   string                 = "A"
// 	steps   int                    = 6
// 	program map[string]Instruction = map[string]Instruction{
// 		"A": {[2]SubInstruction{{1, 1, "B"}, {0, -1, "B"}}},
// 		"B": {[2]SubInstruction{{1, -1, "A"}, {1, 1, "A"}}},
// 	}
// )

var (
	state   string                 = "A"
	steps   int                    = 12425180
	program map[string]Instruction = map[string]Instruction{
		"A": {[2]SubInstruction{{1, 1, "B"}, {0, 1, "F"}}},
		"B": {[2]SubInstruction{{0, -1, "B"}, {1, -1, "C"}}},
		"C": {[2]SubInstruction{{1, -1, "D"}, {0, 1, "C"}}},
		"D": {[2]SubInstruction{{1, -1, "E"}, {1, 1, "A"}}},
		"E": {[2]SubInstruction{{1, -1, "F"}, {0, -1, "D"}}},
		"F": {[2]SubInstruction{{1, 1, "A"}, {0, -1, "E"}}},
	}
)

func duration(invocation time.Time, name string) {
	fmt.Println(name, time.Since(invocation))
}

func partOne() int {
	defer duration(time.Now(), "part 1")
	for i := 0; i < steps; i++ {
		var subinstr SubInstruction = program[state].instructions[tape[pos]]
		tape[pos] = subinstr.write
		pos += subinstr.move
		state = subinstr.next
	}
	// fmt.Println("tape length", len(tape)) // 6196
	var result int
	for _, v := range tape {
		result += v
	}
	return result
}

func main() {
	defer duration(time.Now(), "main")

	fmt.Println(partOne()) // 3099
}

/*
$ go run main.go
part 1 499.425801ms
3099
main 499.446435ms
*/
