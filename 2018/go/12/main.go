// https://adventofcode.com/2018/day/12 Subterranean Sustainability
package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"strings"
	"time"
)

//go:embed input.txt
var input string

func duration(invocation time.Time, name string) {
	fmt.Println(name, time.Since(invocation))
}

func partOne() int {
	defer duration(time.Now(), "part 1")

	var state string
	var rules map[string]string = make(map[string]string)

	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Scan()
	if n, err := fmt.Sscanf(scanner.Text(), "initial state: %s", &state); n != 1 {
		fmt.Println(err)
		return -1
	}
	scanner.Scan() // skip blank line
	for scanner.Scan() {
		var lhs, rhs string
		if n, err := fmt.Sscanf(scanner.Text(), "%s => %s", &lhs, &rhs); n != 2 {
			fmt.Println(err)
			return -1
		}
		rules[lhs] = rhs
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	// fmt.Println(state)
	// fmt.Println(rules)

	// center pot is at state[3]

	const GENERATIONS int = 20         // use 100 to find part 2 pattern
	const EXTRAS int = GENERATIONS * 4 // found by trial and error

	sum := func() int {
		var result int
		for i := 0; i < len(state); i++ {
			if state[i] == '#' {
				result += i - (EXTRAS - GENERATIONS*2)
			}
		}
		return result
	}

	state = strings.Repeat(".", EXTRAS) + state + strings.Repeat(".", EXTRAS)
	var prevSum int
	for gen := 0; gen < GENERATIONS; gen++ {
		// fmt.Println(gen, state)

		var newState string
		for i := 0; i < len(state)-4; i++ {
			frag := state[i : i+5]
			if len(frag) != 5 {
				panic("frag is not 5")
			}
			if ch, ok := rules[frag]; ok {
				newState += ch
			} else {
				newState += "." // test1/example only
			}
		}
		state = newState
		thisSum := sum()
		fmt.Println(gen, thisSum, prevSum-thisSum)
		// after 95 generations, 32 is added to the sum each generation
		prevSum = thisSum
	}
	fmt.Println("--", state)
	var result int
	for i := 0; i < len(state); i++ {
		if state[i] == '#' {
			result += i - (EXTRAS - GENERATIONS*2)
		}
	}

	return result
}

func main() {
	defer duration(time.Now(), "main")

	fmt.Println(partOne()) // 2063
	// ((50000000000 - GENERATIONS) * each_gen_sum) + sum_after_GENERATIONS
	fmt.Println(((50000000000 - 100) * 32) + 3528) // 1600000000328
}

/*
$ go run main.go
*/
