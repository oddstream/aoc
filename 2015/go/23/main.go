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
	op  string
	reg *int
	num int
}

var pc, a, b int
var instructions []Instruction

func parseInput() {
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		var i Instruction
		line := scanner.Text()
		line = strings.ReplaceAll(line, ",", "")
		tokens := strings.Split(line, " ")
		var err error
		i.op = tokens[0]
		switch i.op {
		case "inc", "tpl", "hlf":
			switch tokens[1] {
			case "a":
				i.reg = &a
			case "b":
				i.reg = &b
			default:
				fmt.Println("unknown reg", tokens)
			}
		case "jmp":
			i.num, err = strconv.Atoi(tokens[1])
			if err != nil {
				fmt.Println(tokens, err)
			}
		case "jio", "jie":
			switch tokens[1] {
			case "a":
				i.reg = &a
			case "b":
				i.reg = &b
			default:
				fmt.Println("unknown reg", tokens)
			}
			i.num, err = strconv.Atoi(tokens[2])
			if err != nil {
				fmt.Println(tokens, err)
			}
		}
		instructions = append(instructions, i)
		// fmt.Println(i)
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("scanner error: %v\n", err)
	}
}

func run(starta, startb int) {
	pc, a, b = 0, starta, startb
	for pc >= 0 && pc < len(instructions) {
		var i Instruction = instructions[pc]
		switch i.op {
		case "hlf":
			*i.reg = *i.reg / 2
			pc += 1
		case "tpl":
			*i.reg = *i.reg * 3
			pc += 1
		case "inc":
			*i.reg += 1
			pc += 1
		case "jmp":
			pc += i.num
		case "jie":
			if *i.reg%2 == 0 {
				pc += i.num
			} else {
				pc += 1
			}
		case "jio":
			if *i.reg == 1 {
				pc += i.num
			} else {
				pc += 1
			}
		default:
			fmt.Println("unhandled op", i)
		}
	}
}

func main() {
	defer duration(time.Now(), "main")

	parseInput()
	run(0, 0)
	fmt.Println("part 1", b) // 255
	run(1, 0)
	fmt.Println("part 2", b) // 334
}
