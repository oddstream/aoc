package main

import (
	"bufio"
	_ "embed"
	"flag"
	"fmt"
	"math"
	"strconv"
	"strings"
)

//go:embed "input.txt"
var input string

type Instruction struct {
	op1, op2 string // operands
	CMD      string // operator
	dst      string
}

var wires map[string]Instruction = make(map[string]Instruction)
var signals map[string]uint16 = make(map[string]uint16)

// wires are always lowercase alphas
// CMD is always UPPERCASE alphas
func parseInstruction(line string) Instruction {
	var i Instruction = Instruction{}
	// split into left and right
	lr := strings.Split(line, " -> ")
	i.dst = lr[1]
	lhs := strings.Split(lr[0], " ")
	switch len(lhs) {
	case 1:
		// <wire> | <number>
		i.op1 = lhs[0]
		// CMD will be ""
	case 2:
		// NOT <wire> | <number>
		if lhs[0] == "NOT" {
			i.CMD = "NOT"
		} else {
			fmt.Println("arity 2 CMD is not NOT", lhs[0])
		}
		i.op1 = lhs[1]
	case 3:
		// <wire> | <number> AND | OR | LSHIFT | RSHIFT <wire> | <number>
		i.op1 = lhs[0]
		if lhs[1] == "AND" || lhs[1] == "OR" || lhs[1] == "LSHIFT" || lhs[1] == "RSHIFT" {
			i.CMD = lhs[1]
		} else {
			fmt.Println("arity 3 unknown CMD", lhs[1])
		}
		i.op2 = lhs[2]
	}
	return i
}

func execute(wire string) uint16 {

	if result, ok := signals[wire]; ok {
		return result
	}

	if val, err := strconv.Atoi(wire); err == nil {
		return uint16(val)
	}

	var result uint16
	if i, ok := wires[wire]; ok {
		switch i.CMD {
		case "":
			result = execute(i.op1)
		case "NOT":
			result = math.MaxUint16 ^ execute(i.op1)
		case "AND":
			result = execute(i.op1) & execute(i.op2)
		case "OR":
			result = execute(i.op1) | execute(i.op2)
		case "LSHIFT":
			result = execute(i.op1) << execute(i.op2)
		case "RSHIFT":
			result = execute(i.op1) >> execute(i.op2)
		}
	}
	signals[wire] = result
	return result
}

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "1 or 2")
	flag.Parse()

	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		var i = parseInstruction(scanner.Text())
		wires[i.dst] = i
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("scanner error: %v\n", err)
	}
	// for k := range wires {
	// 	fmt.Println(k, wires[k])
	// }
	part1result := execute("a")
	// store result from "a" into "b"
	wires["b"] = Instruction{op1: fmt.Sprintf("%v", part1result), dst: "b"}
	// clear the signals map
	signals = make(map[string]uint16)
	part2result := execute("a")
	fmt.Println(part1result, part2result) // 956, 40149
}
