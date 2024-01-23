// https://adventofcode.com/2019/day/2 1202 Program Alarm
package main

import (
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

func run(prog []int, noun, verb int) {
	prog[1] = noun
	prog[2] = verb

	var pc int = 0
	for prog[pc] != 99 {
		if prog[pc] == 1 {
			a := prog[pc+1]
			b := prog[pc+2]
			c := prog[pc+3]
			prog[c] = prog[a] + prog[b]
			pc += 4
		} else if prog[pc] == 2 {
			a := prog[pc+1]
			b := prog[pc+2]
			c := prog[pc+3]
			prog[c] = prog[a] * prog[b]
			pc += 4
		} else {
			fmt.Println("unexpected value", prog[pc], "at position", pc)
			break
		}
	}
}

func partOne(program []int, expected int) {
	defer duration(time.Now(), "part 1")

	var prog []int = make([]int, len(program))
	copy(prog, program)

	run(prog, 12, 2)

	var result int = prog[0]
	if result != expected {
		fmt.Println("ERROR: got", result, "expected", expected)
	} else {
		fmt.Println("CORRECT:", result)
	}
}

func partTwo(program []int, expected int) {
	defer duration(time.Now(), "part 2")

	var result int

	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {

			var prog []int = make([]int, len(program))
			copy(prog, program)

			run(prog, noun, verb)

			if prog[0] == 19690720 {
				result = 100*noun + verb
				goto exit
			}
		}
	}

exit:
	if result != expected {
		fmt.Println("ERROR: got", result, "expected", expected)
	} else {
		fmt.Println("CORRECT:", result)
	}
}

func main() {
	defer duration(time.Now(), "main")

	var tokens []string = strings.Split(strings.Trim(input, "\n"), ",")
	var program []int
	for _, tok := range tokens {
		program = append(program, atoi(tok))
	}

	partOne(program, 3101878)
	partTwo(program, 8444)
}

/*
$ go run main.go
CORRECT: 3101878
part 1 33.308µs
CORRECT: 8444
part 2 3.265397ms
main 3.335786ms
*/
