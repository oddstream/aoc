// https://adventofcode.com/2017/day/24 Electromagnetic Moat
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

type Component struct {
	a, b int
	used bool
}

var (
	strongest,
	max_length,
	strength_of_longest int
	components []Component
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func run(pins, length, strength int) {
	strongest = max(strength, strongest)
	max_length = max(length, max_length)
	if length == max_length {
		strength_of_longest = max(strength, strength_of_longest)
	}
	for i := 0; i < len(components); i++ {
		// we're mutating items in an array, so we need a pointer
		c := &components[i]
		if !c.used && (c.a == pins || c.b == pins) {
			c.used = true
			var needpins int
			if c.a == pins {
				needpins = c.b
			} else {
				needpins = c.a
			}
			run(needpins, length+1, strength+c.strength())
			c.used = false
		}
	}
}

// func (c Component) connect(n int) (int, bool) {
// 	if n == c.a {
// 		return c.b, true
// 	} else if n == c.b {
// 		return c.a, true
// 	}
// 	return 0, false
// }

func (c Component) strength() int {
	return c.a + c.b
}

func duration(invocation time.Time, name string) {
	fmt.Println(name, time.Since(invocation))
}

// func helper(arr []Component, start int, c chan []Component) {
// 	if start >= len(arr) {
// 		c <- append([]Component{}, arr...)
// 		return
// 	}
// 	helper(arr, start+1, c)
// 	for i := start + 1; i < len(arr); i++ {
// 		arr[start], arr[i] = arr[i], arr[start]
// 		helper(arr, start+1, c)
// 		arr[start], arr[i] = arr[i], arr[start]
// 	}
// }

// func combinations(arr []Component) <-chan []Component {
// 	c := make(chan []Component)
// 	go func() {
// 		defer close(c)
// 		helper(arr, 0, c)
// 	}()

// 	return c
// }

func loadInput() []Component {
	var lst []Component
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		var c Component
		if n, err := fmt.Sscanf(scanner.Text(), "%d/%d", &c.a, &c.b); n != 2 {
			fmt.Println(err)
			break
		}
		lst = append(lst, c)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	return lst
}

func main() {
	defer duration(time.Now(), "main")

	components = loadInput()
	run(0, 0, 0)
	fmt.Println(strongest)           // 1511
	fmt.Println(strength_of_longest) // 1471
}

/*
$ go run main.go
1511
1471
main 81.519002ms
*/
