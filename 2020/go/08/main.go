// https://adventofcode.com/2020/day/8
package main

import (
	"bufio"
	_ "embed"
	"errors"
	"fmt"
	"runtime"
	"strings"
	"time"
)

//go:embed test1.txt
var test1 string

//go:embed input.txt
var input string

type Instruction struct {
	operation string
	argument  int
	deployed  int
}

func duration(invocation time.Time, name string) {
	fmt.Println(name, time.Since(invocation))
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

func loadBootCode(in string) []Instruction {
	var bootCode []Instruction
	scanner := bufio.NewScanner(strings.NewReader(in))
	for scanner.Scan() {
		var i Instruction
		if n, err := fmt.Sscanf(scanner.Text(), "%s %d", &i.operation, &i.argument); n != 2 {
			fmt.Println(err)
			break
		}
		bootCode = append(bootCode, i)
	}
	return bootCode
}

func run(bootCode []Instruction) (int, error) {
	var accumulator, pc int
	for pc < len(bootCode) {
		if bootCode[pc].deployed > 0 {
			return accumulator, errors.New("infinite loop")
		}
		switch bootCode[pc].operation {
		case "nop":
			bootCode[pc].deployed += 1
			pc += 1
		case "acc":
			accumulator += bootCode[pc].argument
			bootCode[pc].deployed += 1
			pc += 1
		case "jmp":
			bootCode[pc].deployed += 1
			pc += bootCode[pc].argument
		}
	}
	return accumulator, nil
}

func part1(in string, expected int) (result int) {
	defer duration(time.Now(), "part 1")

	var bootCode []Instruction = loadBootCode(in)
	result, _ = run(bootCode)
	report(expected, result)
	return result
}

func part2(in string, expected int) (result int) {
	defer duration(time.Now(), "part 2")

	var change int
	for {
		var err error
		var bootCode []Instruction = loadBootCode(in)
		if change >= len(bootCode) {
			panic("change has overflowed")
		}
		if bootCode[change].operation == "nop" {
			bootCode[change].operation = "jmp"
			result, err = run(bootCode)
			if err == nil {
				break
			}
		} else if bootCode[change].operation == "jmp" {
			bootCode[change].operation = "nop"
			result, err = run(bootCode)
			if err == nil {
				break
			}
		}
		change += 1
	}

	report(expected, result)
	return result
}

func main() {
	defer duration(time.Now(), "main")

	// part1(test1, 5)
	part1(input, 1200)
	// part2(test1, 8)
	part2(input, 1023)

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
