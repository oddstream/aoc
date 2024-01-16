// https://adventofcode.com/2018/day/8 Memory Maneuver
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

// input contains 16783 integers

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

func partOne() int {
	defer duration(time.Now(), "part 1")

	var numbers []int

	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		numbers = append(numbers, atoi(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	var result, pos int
	var metadata func()
	metadata = func() {
		next := func() int {
			n := numbers[pos]
			pos += 1
			return n
		}
		// if !(pos < len(numbers)-2) {
		// 	return
		// }
		children := next()
		metas := next()
		for ; children > 0; children-- {
			metadata()
		}
		for ; metas > 0; metas-- {
			result += next()
		}
	}

	metadata()

	return result
}

func partTwo() int {
	defer duration(time.Now(), "part 2")

	var numbers []int

	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		numbers = append(numbers, atoi(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	var index int

	var values func() int
	values = func() int {
		next := func() int {
			n := numbers[index]
			index += 1
			return n
		}
		var c = next()
		var m = next()
		var value int
		var children []int
		for i := 0; i < c; i++ {
			children = append(children, values())
		}
		// If a node has no child nodes,
		// its value is the sum of its metadata entries.
		// if a node does have child nodes,
		// the metadata entries become indexes which refer to those child nodes
		// A metadata entry of 1 refers to the first child node,
		// 2 to the second, 3 to the third, and so on.
		for i := 0; i < m; i++ {
			v := next()
			value += func() int {
				if c == 0 {
					return v
				} else if v-1 < len(children) {
					return children[v-1] // indexes are 1-based
				} else {
					return 0
				}
			}()
		}
		return value
	}
	return values()
}

func main() {
	defer duration(time.Now(), "main")

	fmt.Println(partOne()) // 138, 42768
	fmt.Println(partTwo()) // 66, 34348
}

/*
$ go run main.go
part 1 648.207µs
42768
part 2 635.545µs
34348
main 1.30337ms
*/
