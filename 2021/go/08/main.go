// https://adventofcode.com/2021/day/8
package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"time"
)

//go:embed test1.txt
var test1 string

//go:embed input.txt
var input string

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

func part1(in string, expected int) (result int) {
	defer duration(time.Now(), "part 1")

	scanner := bufio.NewScanner(strings.NewReader(in))
	for scanner.Scan() {
		var sides []string = strings.Split(scanner.Text(), " | ")
		var outputs []string = strings.Fields(sides[1])
		for _, output := range outputs {
			// 1, 4, 7 or 8
			if len(output) == 2 || len(output) == 4 || len(output) == 3 || len(output) == 7 {
				result += 1
			}
		}
	}
	report(expected, result)
	return result
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

type CharSet map[string]struct{}

func set(str string) CharSet {
	var s CharSet = make(CharSet)
	for _, ch := range strings.Split(str, "") {
		s[ch] = struct{}{}
	}
	return s
}

func intersection(a, b CharSet) CharSet {
	var result CharSet = make(CharSet)
	for key := range a {
		if _, ok := b[key]; ok {
			result[key] = struct{}{}
		}
	}
	return result
}

func part2(in string, expected int) (result int) {
	defer duration(time.Now(), "part 2")

	scanner := bufio.NewScanner(strings.NewReader(in))
	for scanner.Scan() {
		var sides []string = strings.Split(scanner.Text(), " | ")

		var signals []string = strings.Fields(sides[0])

		// relying on signals to always contain a 2- and 4-length item
		var d map[int]CharSet = make(map[int]CharSet)
		for _, s := range signals {
			l := len(s)
			if l == 2 || l == 4 {
				d[l] = set(s)
			}
		}
		if len(d) != 2 {
			panic("set d is too short")
		}

		var outputs []string = strings.Fields(sides[1])
		var n string
		for _, o := range outputs {
			// determine the four obvious numbers from their unique lengths
			// (2, 4, 3, 7 can only be "1", "4", "7", "8")
			// (strong hint from part 1)
			// determine the rest (lengths 5 and 6)
			// by comparing the overlap of segments
			// between known outputs and the ambiguous ones
			l := len(o)
			if l == 2 {
				n += "1"
			} else if l == 4 {
				n += "4"
			} else if l == 3 {
				n += "7"
			} else if l == 7 {
				n += "8"
			} else if l == 5 {
				s := set(o)
				if len(intersection(s, d[2])) == 2 {
					// two segments in common with "1"
					n += "3"
				} else if len(intersection(s, d[4])) == 2 {
					// two segments in common with "4"
					n += "2"
				} else {
					// one segment in common with "1"
					// three segments in common with "4"
					n += "5"
				}
			} else if l == 6 {
				s := set(o)
				if len(intersection(s, d[2])) == 1 {
					// one segment in common with "1"
					n += "6"
				} else if len(intersection(s, d[4])) == 4 {
					// four segments in common with "4"
					n += "9"
				} else {
					// two segments in common with "1"
					// three segments in common with "4"
					n += "0"
				}
			} else {
				panic("stange output length")
			}
		}
		result += atoi(n)
	}
	report(expected, result)
	return result
}

func main() {
	defer duration(time.Now(), "main")

	// part1(test1, 26)
	part1(input, 532)
	// part2(test1, 61229)
	part2(input, 1011284)

	{
		var memStats runtime.MemStats
		runtime.ReadMemStats(&memStats)
		fmt.Printf("Heap memory (in bytes): %d\n", memStats.HeapAlloc)
		fmt.Printf("Number of garbage collections: %d\n", memStats.NumGC)
	}
}

/*
$ go run main.go
RIGHT ANSWER: 532
part 1 125.057µs
RIGHT ANSWER: 1011284
part 2 526.005µs
Heap memory (in bytes): 422256
Number of garbage collections: 0
main 740.571µs
*/
