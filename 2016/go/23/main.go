// https://adventofcode.com/2016/day/23 Safe Cracking
package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"strconv"
	"strings"
	"time"
)

//go:embed input.txt
var input string

func duration(invocation time.Time, name string) {
	fmt.Println(name, time.Since(invocation))
}

type Instruction struct {
	op, first, second string
}

func run(program []Instruction) int {
	var a, b, c, d, pc int
	var regMap map[string]*int = make(map[string]*int)
	regMap["a"] = &a
	regMap["b"] = &b
	regMap["c"] = &c
	regMap["d"] = &d
	// a = 7	// part 1	11640
	a = 12 // part 2	479008200
	for pc < len(program) {
		var instr = program[pc]
		// fmt.Println(pc, a, b, c, d, instr)
		switch instr.op {
		case "cpy":
			// cpy x y copies x (either an integer or the value of a register) into register y
			dst := regMap[instr.second]
			if src, ok := regMap[instr.first]; ok {
				*dst = *src
			} else {
				if n, err := strconv.Atoi(instr.first); err == nil {
					*dst = n
				} else {
					fmt.Println("skipping invalid cpy", instr.first, instr.second)
				}
			}
			pc = pc + 1
		case "inc":
			// inc x increases the value of register x by one.
			*regMap[instr.first]++
			pc = pc + 1
		case "dec":
			// dec x decreases the value of register x by one
			*regMap[instr.first]--
			pc = pc + 1
		case "jnz":
			// jnz x y jumps to an instruction y away (positive means forward;
			// negative means backward), but only if x is not zero.
			var x, y int
			if reg, ok := regMap[instr.first]; ok {
				x = *reg
			} else {
				if n, err := strconv.Atoi(instr.first); err == nil {
					x = n
				}
			}
			if reg, ok := regMap[instr.second]; ok {
				y = *reg
			} else {
				if n, err := strconv.Atoi(instr.second); err == nil {
					y = n
				}
			}
			if x != 0 {
				pc = pc + y
			} else {
				pc = pc + 1
			}
		case "tgl":
			// tgl x toggles the instruction x away (pointing at instructions like jnz
			// does: positive means forward; negative means backward):
			x := pc + *regMap[instr.first]
			// If an attempt is made to toggle an instruction outside the program,
			// nothing happens.
			if x < len(program) {
				switch program[x].op {
				// For one-argument instructions, inc becomes dec, and all other one-
				// argument instructions become inc.
				case "inc":
					program[x].op = "dec"
				case "dec", "tgl":
					program[x].op = "inc"
				// For two-argument instructions, jnz becomes cpy, and all other two-
				// instructions become jnz.
				case "jnz":
					program[x].op = "cpy"
				case "cpy":
					program[x].op = "jnz"
				}
				// The arguments of a toggled instruction are not affected.
			}
			// If toggling produces an invalid instruction (like cpy 1 2) and an
			// attempt is later made to execute that instruction, skip it instead.
			// If tgl toggles itself (for example, if a is 0, tgl a would target
			// itself and become inc a), the resulting instruction is not executed
			// until the next time it is reached.
			pc = pc + 1
		}
	}
	return a
}

func main() {
	defer duration(time.Now(), "main")

	var program []Instruction

	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		tokens := strings.Split(scanner.Text(), " ")
		var instr Instruction = Instruction{op: tokens[0], first: tokens[1]}
		if len(tokens) == 3 {
			instr.second = tokens[2]
		}
		program = append(program, instr)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	fmt.Println(run(program)) // part 1: 11640, part2: 479008200
}

/*
$ go run main.go
479008200
main 1m2.424041616s
*/
