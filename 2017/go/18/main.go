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

// or we could use K&R p61
func atoi(s string) int {
	if n, err := strconv.Atoi(s); err == nil {
		return n
	} else {
		fmt.Println(err)
	}
	return 0
}

func getval(reg map[string]int, tok string) int {
	if tok[0] >= 'a' && tok[0] <= 'z' {
		return reg[tok]
	} else {
		return atoi(tok)
	}
}

func load() [][]string {
	var prog [][]string
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		prog = append(prog, strings.Split(scanner.Text(), " "))
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	return prog
}

func run1(prog [][]string) int {
	var reg map[string]int = make(map[string]int)
	var pc, playing int
	for pc >= 0 && pc < len(prog) {
		instr := prog[pc]
		switch instr[0] {
		case "snd":
			playing = getval(reg, instr[1])
			pc += 1
		case "set":
			reg[instr[1]] = getval(reg, instr[2])
			pc += 1
		case "add":
			reg[instr[1]] += getval(reg, instr[2])
			pc += 1
		case "mul":
			reg[instr[1]] *= getval(reg, instr[2])
			pc += 1
		case "mod":
			reg[instr[1]] %= getval(reg, instr[2])
			pc += 1
		case "rcv":
			if getval(reg, instr[1]) != 0 {
				return playing
			}
			pc += 1
		case "jgz":
			if getval(reg, instr[1]) > 0 {
				pc += getval(reg, instr[2])
			} else {
				pc += 1
			}
		}
	}
	return -1
}

func run2(prog [][]string, id int, in, out chan int) int {
	var reg map[string]int = make(map[string]int)
	reg["p"] = id

	var pc, result int
	for pc >= 0 && pc < len(prog) {
		instr := prog[pc]
		switch instr[0] {
		case "snd":
			select {
			case out <- getval(reg, instr[1]):
				result += 1
			case <-time.After(time.Second / 10):
				return -1
			}
			pc += 1
		case "set":
			reg[instr[1]] = getval(reg, instr[2])
			pc += 1
		case "add":
			reg[instr[1]] += getval(reg, instr[2])
			pc += 1
		case "mul":
			reg[instr[1]] *= getval(reg, instr[2])
			pc += 1
		case "mod":
			reg[instr[1]] %= getval(reg, instr[2])
			pc += 1
		case "rcv":
			select {
			case reg[instr[1]] = <-in:
			case <-time.After(time.Second / 10):
				return result
			}
			pc += 1
		case "jgz":
			if getval(reg, instr[1]) > 0 {
				pc += getval(reg, instr[2])
			} else {
				pc += 1
			}
		}
	}
	return -1
}

func partOne(prog [][]string) int {
	defer duration(time.Now(), "part 1")
	return run1(prog)
}

func partTwo(prog [][]string) int {
	defer duration(time.Now(), "part 2")

	ch0 := make(chan int, 1000)
	ch1 := make(chan int, 1000)

	go run2(prog, 0, ch0, ch1)

	result := run2(prog, 1, ch1, ch0)

	close(ch0)
	close(ch1)

	return result
}

func main() {
	defer duration(time.Now(), "main")

	prog := load()
	fmt.Println(partOne(prog)) // 8600
	fmt.Println(partTwo(prog)) // 7239
}

/*
$ go run main.go
*/
