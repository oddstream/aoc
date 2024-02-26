// https://adventofcode.com/2018/day/21
package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"runtime"
	"strings"
	"time"
)

//go:embed input.txt
var input string

type (
	Instruction struct {
		opcode  string
		A, B, C int
	}
	Registers [6]int                        // 6 registers, 0 .. 5
	OpFunc    func(int, int, Registers) int // output always goes to Registers[C]
)

func duration(invocation time.Time, name string) {
	fmt.Println(name, "duration", time.Since(invocation))
}

func report(expected, result int) {
	if expected != -1 {
		if result != expected {
			fmt.Println("ERROR: got", result, "expected", expected)
		} else {
			fmt.Println("RIGHT ANSWER:", result)
		}
	}
}

// directly copied from day 16
var ops map[string]OpFunc = map[string]OpFunc{
	"addr": func(A, B int, r Registers) int { return r[A] + r[B] },
	"addi": func(A, B int, r Registers) int { return r[A] + B },
	"mulr": func(A, B int, r Registers) int { return r[A] * r[B] },
	"muli": func(A, B int, r Registers) int { return r[A] * B },
	"banr": func(A, B int, r Registers) int { return r[A] & r[B] },
	"bani": func(A, B int, r Registers) int { return r[A] & B },
	"borr": func(A, B int, r Registers) int { return r[A] | r[B] },
	"bori": func(A, B int, r Registers) int { return r[A] | B },
	"setr": func(A, B int, r Registers) int { return r[A] },
	"seti": func(A, B int, r Registers) int { return A },
	"gtir": func(A, B int, r Registers) int {
		if A > r[B] {
			return 1
		} else {
			return 0
		}
	},
	"gtri": func(A, B int, r Registers) int {
		if r[A] > B {
			return 1
		} else {
			return 0
		}
	},
	"gtrr": func(A, B int, r Registers) int {
		if r[A] > r[B] {
			return 1
		} else {
			return 0
		}
	},
	"eqir": func(A, B int, r Registers) int {
		if A == r[B] {
			return 1
		} else {
			return 0
		}
	},
	"eqri": func(A, B int, r Registers) int {
		if r[A] == B {
			return 1
		} else {
			return 0
		}
	},
	"eqrr": func(A, B int, r Registers) int {
		if r[A] == r[B] {
			return 1
		} else {
			return 0
		}
	},
}

// reg[0] is set by calling func
// reg[0] is only referenced by the "eqrr 4 0 5" instruction (program[28])
// which compares reg[4] with reg[0] and sets register C
func execute1(ipreg int, program []Instruction, reg Registers) int {
	for reg[ipreg] >= 0 && reg[ipreg] < len(program) {
		// fmt.Print(reg)
		var i Instruction = program[reg[ipreg]]
		var fn OpFunc = ops[i.opcode]
		reg[i.C] = fn(i.A, i.B, reg)
		// fmt.Println(" ", i, " ", reg)
		reg[ipreg] += 1
		// if reg[ipreg] == 28 {
		// 	fmt.Println(reg[4])
		// }
	}
	fmt.Println(reg)
	return reg[0]
}

func part1(in string, expected int) (result int) {
	defer duration(time.Now(), "part 1")

	var ip int
	scanner := bufio.NewScanner(strings.NewReader(in))
	scanner.Scan()
	if n, err := fmt.Sscanf(scanner.Text(), "#ip %d", &ip); n != 1 {
		fmt.Println(err)
		return -1
	}
	var program []Instruction
	for scanner.Scan() {
		var ins Instruction
		if n, err := fmt.Sscanf(scanner.Text(), "%s %d %d %d", &ins.opcode, &ins.A, &ins.B, &ins.C); n != 4 {
			fmt.Println(err)
			return -1
		}
		program = append(program, ins)
	}
	result = execute1(ip, program, Registers{16128384, 0, 0, 0, 0, 0})
	report(expected, result)
	return result
}

func part2(in string, expected int) (result int) {
	defer duration(time.Now(), "part 2")

	var ipreg int
	scanner := bufio.NewScanner(strings.NewReader(in))
	scanner.Scan()
	if n, err := fmt.Sscanf(scanner.Text(), "#ip %d", &ipreg); n != 1 {
		fmt.Println(err)
		return -1
	}
	var program []Instruction
	for scanner.Scan() {
		var ins Instruction
		if n, err := fmt.Sscanf(scanner.Text(), "%s %d %d %d", &ins.opcode, &ins.A, &ins.B, &ins.C); n != 4 {
			fmt.Println(err)
			return -1
		}
		program = append(program, ins)
	}

	var seen map[int]struct{} = map[int]struct{}{}
	var last int

	var reg Registers = Registers{0, 0, 0, 0, 0, 0}
	for reg[ipreg] >= 0 && reg[ipreg] < len(program) {
		// fmt.Print(reg)
		var i Instruction = program[reg[ipreg]]
		var fn OpFunc = ops[i.opcode]
		reg[i.C] = fn(i.A, i.B, reg)
		// fmt.Println(" ", i, " ", reg)
		reg[ipreg] += 1
		if reg[ipreg] == 28 {
			if _, ok := seen[reg[4]]; ok {
				result = last
				break
			}
			seen[reg[4]] = struct{}{}
			last = reg[4]
		}
	}
	report(expected, result)
	return result
}

func main() {
	defer duration(time.Now(), "main")

	part1(input, 16128384)
	part2(input, 7705368)
	{
		var memStats runtime.MemStats
		runtime.ReadMemStats(&memStats)
		fmt.Printf("Heap memory (in bytes): %d\n", memStats.HeapAlloc)
		fmt.Printf("Number of garbage collections: %d\n", memStats.NumGC)
	}
}

/*
$ go run main.go
[16128384 31 1 1 16128384 1]
RIGHT ANSWER: 16128384
part 1 duration 130.216µs
RIGHT ANSWER: 7705368
part 2 duration 28.381505986s
Heap memory (in bytes): 542088
Number of garbage collections: 0
main duration 28.38179502s
*/
