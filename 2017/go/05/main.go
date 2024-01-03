// https://adventofcode.com/2017/day/5 A Maze of Twisty Trampolines, All Alike
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

func partOne() int {
	var program []int

	var run func() int = func() int {
		var pc, steps int
		for pc < len(program) {
			program[pc] += 1
			pc += program[pc] - 1
			steps += 1
		}
		return steps
	}

	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		if n, err := strconv.Atoi(scanner.Text()); err == nil {
			program = append(program, n)
		} else {
			fmt.Println(err)
			break
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	return run()
}

func partTwo() int {
	var program []int

	var run func() int = func() int {
		var pc, steps int
		for pc < len(program) {
			if program[pc] >= 3 {
				program[pc] -= 1
				pc += program[pc] + 1
			} else {
				program[pc] += 1
				pc += program[pc] - 1
			}
			steps += 1
		}
		return steps
	}

	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		if n, err := strconv.Atoi(scanner.Text()); err == nil {
			program = append(program, n)
		} else {
			fmt.Println(err)
			break
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	return run()
}

func main() {
	defer duration(time.Now(), "main")

	fmt.Println(partOne()) // 391540
	fmt.Println(partTwo()) // 30513679
}

/*
$ go run main.go
391540
30513679
main 136.572984ms
*/
