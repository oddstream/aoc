// https://adventofcode.com/2016/day/25 Clock Signal
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

func run(program []Instruction, initiala int) bool {
	var a, b, c, d, pc int
	var regMap map[string]*int = make(map[string]*int)
	regMap["a"] = &a
	regMap["b"] = &b
	regMap["c"] = &c
	regMap["d"] = &d
	a = initiala
	var cycles, collected int
	var expectZero = true
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
		// case "tgl": removed as that instruction does not appear in the input
		case "out":
			// out x transmits x (either an integer or the value of a register) as
			// the next value for the clock signal.
			var x int
			if reg, ok := regMap[instr.first]; ok {
				x = *reg
			} else {
				if n, err := strconv.Atoi(instr.first); err == nil {
					x = n
				} else {
					fmt.Println(err)
				}
			}
			if expectZero && x == 0 {
				collected += 1
				expectZero = false
			} else if !expectZero && x == 1 {
				collected += 1
				expectZero = true
			} else {
				return false
			}
			pc = pc + 1
		}
		cycles += 1
		if cycles > 100000 { // arrived at by trial and error
			if collected > 1 {
				fmt.Println(initiala, "collected", collected, "alternating binary digits before cycles exceeded")
			}
			return false
		}
	}
	return true
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
	for i := 0; i < 200; i++ { // arrived at by trial and error
		run(program, i)
	}
}

/*
$ go run main.go
158 collected 36 alternating binary digits before cycles exceeded
main 112.417469ms
*/
