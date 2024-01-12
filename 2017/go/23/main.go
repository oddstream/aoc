// https://adventofcode.com/2017/day/23 Coprocessor Conflagration
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
	var pc, result int
	for pc >= 0 && pc < len(prog) {
		instr := prog[pc]
		switch instr[0] {
		case "set":
			reg[instr[1]] = getval(reg, instr[2])
			pc += 1
		case "sub":
			reg[instr[1]] -= getval(reg, instr[2])
			pc += 1
		case "mul":
			result += 1
			reg[instr[1]] *= getval(reg, instr[2])
			pc += 1
		case "jnz":
			if getval(reg, instr[1]) != 0 {
				pc += getval(reg, instr[2])
			} else {
				pc += 1
			}
		}
	}
	return result
}

func run2(prog [][]string) int {
	var reg map[string]int = map[string]int{"a": 1}
	var pc int
	for pc >= 0 && pc < len(prog) {
		instr := prog[pc]
		switch instr[0] {
		case "set":
			reg[instr[1]] = getval(reg, instr[2])
			pc += 1
		case "sub":
			reg[instr[1]] -= getval(reg, instr[2])
			pc += 1
		case "mul":
			reg[instr[1]] *= getval(reg, instr[2])
			pc += 1
		case "jnz":
			if getval(reg, instr[1]) != 0 {
				pc += getval(reg, instr[2])
			} else {
				pc += 1
			}
		}
	}
	return reg["h"]
}

func run3() int {
	var a, b, c, d, e, f, g, h int

	a = 1       // just a part 1, 2 flag
	b = 57      // 1	b := 57
	c = b       // 2	c := 57
	if a != 0 { // 3	part 2 starts at L5
		goto L5 // +2
	}
	if 1 != 0 { // 4	always true
		goto L9 // +5	skip part 2 variables
	}
L5: // setup part 2 constants
	b *= 100     // 5	b := 5700
	b -= -100000 // 6	b := 105700
	c = b        // 7	c := 105700
	c -= -17000  // 8	c := 122700
L9: // start of outer loop
	f = 1 // 9
	d = 2 // 10
L11:
	e = 2 // 11
L12:
	g = d       // 12
	g *= e      // 13
	g -= b      // 14
	if g != 0 { // 15
		goto L17 // +2
	}
	f = 0 // 16
L17:
	e -= -1     // 17
	g = e       // 18
	g -= b      // 19
	if g != 0 { // 20
		goto L12 // -8
	}
	d -= -1     // 21
	g = d       // 22
	g -= b      // 23
	if g != 0 { // 24
		goto L11 // -13
	}
	if f != 0 { // 25
		goto L27 // +2			skip counting something
	}
	h -= -1 // 26				count something
L27:
	g = b       // 27
	g -= c      // 28
	if g != 0 { // 29
		goto L31 // +2
	}
	if 1 != 0 { // 30			always true ...
		goto L33 // +3			... exit
	}
L31:
	b -= -17    // 31			increment b
	if 1 != 0 { // 32			always true ...
		goto L9 // -23			top of outer loop
	}
L33:
	return h
}

// https://www.reddit.com/r/adventofcode/comments/7lms6p/comment/drnh5sx/?utm_source=share&utm_medium=web2x&context=3
func run4() int {
	b := 105700
	c := 122700

	var h int
	for {
		f := 1
		for d := 2; d*d <= b; d++ {
			if b%d == 0 {
				f = 0
				break
			}
		}
		if f == 0 {
			h++
		}
		if b == c {
			break
		}
		b += 17
	}

	return h
}

func partOne(prog [][]string) int {
	defer duration(time.Now(), "part 1")
	return run1(prog)
}

func partTwo(prog [][]string) int {
	defer duration(time.Now(), "part 2")
	// return run2(prog)
	return run3()
	// return run4()
}

func main() {
	defer duration(time.Now(), "main")

	prog := load()
	fmt.Println(partOne(prog)) // 3025
	fmt.Println(partTwo(prog)) // 915?
}

/*
$ go run main.go
part 1 1.969456ms
3025
part 2 2h6m19.902020012s
915
main 2h6m19.904137945s
*/
