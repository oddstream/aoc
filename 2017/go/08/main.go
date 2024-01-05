// https://adventofcode.com/2017/day/8 I Heard You Like Registers
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
	dstreg  string
	instr   string
	amt     int
	testreg string
	cond    string
	testamt int
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

func partOne() int {
	var result int
	var program []Instruction
	var registers map[string]int = make(map[string]int)

	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		tokens := strings.Split(scanner.Text(), " ")
		// fmt.Println(tokens)
		program = append(program, Instruction{
			dstreg:  tokens[0],
			instr:   tokens[1],
			amt:     atoi(tokens[2]),
			testreg: tokens[4],
			cond:    tokens[5],
			testamt: atoi(tokens[6]),
		})
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	for _, i := range program {
		if _, ok := registers[i.dstreg]; !ok {
			registers[i.dstreg] = 0
		}
		if _, ok := registers[i.testreg]; !ok {
			registers[i.testreg] = 0 // not strictly needed
		}
		tstval := registers[i.testreg]
		switch i.cond {
		case ">":
			if tstval > i.testamt {
				if i.instr == "inc" {
					registers[i.dstreg] += i.amt
				} else if i.instr == "dec" {
					registers[i.dstreg] -= i.amt
				}
			}
		case "<":
			if tstval < i.testamt {
				if i.instr == "inc" {
					registers[i.dstreg] += i.amt
				} else if i.instr == "dec" {
					registers[i.dstreg] -= i.amt
				}
			}
		case ">=":
			if tstval >= i.testamt {
				if i.instr == "inc" {
					registers[i.dstreg] += i.amt
				} else if i.instr == "dec" {
					registers[i.dstreg] -= i.amt
				}
			}
		case "<=":
			if tstval <= i.testamt {
				if i.instr == "inc" {
					registers[i.dstreg] += i.amt
				} else if i.instr == "dec" {
					registers[i.dstreg] -= i.amt
				}
			}
		case "==":
			if tstval == i.testamt {
				if i.instr == "inc" {
					registers[i.dstreg] += i.amt
				} else if i.instr == "dec" {
					registers[i.dstreg] -= i.amt
				}
			}
		case "!=":
			if tstval != i.testamt {
				if i.instr == "inc" {
					registers[i.dstreg] += i.amt
				} else if i.instr == "dec" {
					registers[i.dstreg] -= i.amt
				}
			}
		default:
			fmt.Println("unknown cond", i.cond)
		}

		// move this block outside of loop for part 1
		for _, v := range registers {
			if v > result {
				result = v
			}
		}
	}
	return result
}

func main() {
	defer duration(time.Now(), "main")

	fmt.Println(partOne()) // 7296, 8186
}

/*
$ go run main.go
*/
