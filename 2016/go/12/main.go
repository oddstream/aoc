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
	c = 1 // 0 for part 1, 1 for part 2
	for pc < len(program) {
		var i = program[pc]
		switch i.op {
		case "cpy":
			dst := regMap[i.second]
			if src, ok := regMap[i.first]; ok {
				*dst = *src
			} else {
				if n, err := strconv.Atoi(i.first); err == nil {
					*dst = n
				}
			}
			pc = pc + 1
		case "inc":
			reg := regMap[i.first]
			*reg = *reg + 1
			pc = pc + 1
		case "dec":
			reg := regMap[i.first]
			*reg = *reg - 1
			pc = pc + 1
		case "jnz":
			var val int
			if reg, ok := regMap[i.first]; ok {
				val = *reg
			} else {
				if n, err := strconv.Atoi(i.first); err == nil {
					val = n
				}
			}
			if n, err := strconv.Atoi(i.second); err == nil {
				if val != 0 {
					pc = pc + n
				} else {
					pc = pc + 1
				}
			}
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
		var i Instruction
		switch len(tokens) {
		case 2:
			i.op = tokens[0]
			i.first = tokens[1]
		case 3:
			i.op = tokens[0]
			i.first = tokens[1]
			i.second = tokens[2]
		}
		program = append(program, i)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	fmt.Println(run(program)) // part 1: 318007, part 2 (c=1): 9227661
}
