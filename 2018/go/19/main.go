// https://adventofcode.com/2018/day/19
package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"runtime"
	"strings"
	"time"
)

//go:embed test1.txt
var test1 string

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

func divisorSum(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			// fmt.Println("factor", i)
			sum += i
		}
	}
	return sum
}

func execute(ipreg int, program []Instruction, reg0 int) int {
	var reg Registers
	reg[0] = reg0
	var cycles int
	for reg[ipreg] >= 0 && reg[ipreg] < len(program) {
		// fmt.Print(reg)
		var i Instruction = program[reg[ipreg]]
		var fn OpFunc = ops[i.opcode]
		reg[i.C] = fn(i.A, i.B, reg)
		// fmt.Println(" ", i, " ", reg)
		reg[ipreg] += 1
		cycles += 1
		// if cycles%1_000_000 == 0 {
		// 	fmt.Println(reg)
		// }
		if cycles > 10_000_000 {
			// ... "cheating" alert ...
			// the code seems to compute a large number (10551370)
			// and store it in reg[4] (which seems to stay there,
			// the other registers fluctuate, notably reg[5] counts
			// up) it then spends DAYS doing something with that number;
			// from previous experience with AoC you can guess that
			// it's the sum of all factors (integer divisors) of
			// that large number, so compute that directly...
			return divisorSum(reg[4])
		}
	}
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
	result = execute(ip, program, 0)
	report(expected, result)
	return result
}

func part2(in string, expected int) (result int) {
	defer duration(time.Now(), "part 2")

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
	result = execute(ip, program, 1)
	report(expected, result)
	return result
}

func main() {
	defer duration(time.Now(), "main")

	// part1(test1, 7)
	part1(input, 1764)
	part2(input, 18992484)
	{
		var memStats runtime.MemStats
		runtime.ReadMemStats(&memStats)
		fmt.Printf("Heap memory (in bytes): %d\n", memStats.HeapAlloc)
		fmt.Printf("Number of garbage collections: %d\n", memStats.NumGC)
	}
}

/*
$ go run main.go
*/
