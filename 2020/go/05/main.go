// https://adventofcode.com/2020/day/5
package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"runtime"
	"sort"
	"strings"
	"time"
)

//go:embed input.txt
var input string

func duration(invocation time.Time, name string) {
	fmt.Println(name, time.Since(invocation))
}

func decode1(s string) (row, col int) {
	for i := 0; i < 7; i++ {
		if s[i] == 'B' {
			row |= 1
		}
		if i != 6 {
			row <<= 1
		}
	}
	for i := 7; i < 10; i++ {
		if s[i] == 'R' {
			col |= 1
		}
		if i != 9 {
			col <<= 1
		}
	}
	return
}

func decode2(s string) (row, col int) {
	for i := 0; i < 7; i++ {
		if s[i] == 'B' {
			row |= 1 << (6 - i)
		}
	}
	for i := 7; i < 10; i++ {
		if s[i] == 'R' {
			col |= 1 << (9 - i)
		}
	}
	return
}

func decode3(s string) (row, col int) {
	for i := 0; i < 6; i++ {
		if s[i] == 'B' {
			row |= 1
		}
		row <<= 1
	}
	if s[6] == 'B' {
		row |= 1
	}
	for i := 7; i < 9; i++ {
		if s[i] == 'R' {
			col |= 1
		}
		col <<= 1
	}
	if s[9] == 'R' {
		col |= 1
	}
	return
}

func decodeid(s string) int {
	r, c := decode3(s)
	return r*8 + c
}

func part1(in string, expected int) (result int) {
	defer duration(time.Now(), "part 1")

	scanner := bufio.NewScanner(strings.NewReader(in))
	for scanner.Scan() {
		result = max(result, decodeid(scanner.Text()))
	}

	if expected != -1 {
		if result != expected {
			fmt.Println("ERROR: got", result, "expected", expected)
		} else {
			fmt.Println("RIGHT ANSWER:", result)
		}
	}
	return result
}

func part2(in string, expected int) (result int) {
	defer duration(time.Now(), "part 2")

	var passes []int
	scanner := bufio.NewScanner(strings.NewReader(in))
	for scanner.Scan() {
		passes = append(passes, decodeid(scanner.Text()))
	}
	sort.Slice(passes, func(a, b int) bool {
		return passes[a] < passes[b]
	})
	// len(passes) == 874
	// passes[0] == 48
	// passes[len(passes)-1] == 922
	// 922 - 48 = 874
	// would be able to find missing pass without an array or a sort
	// sum of all possible passes - sum of actual passes = result
	// but using an array and a sort is shorter
	for i := 1; i < len(passes)-2; i++ {
		if passes[i-1] != passes[i]-1 || passes[i+1] != passes[i]+1 {
			// fmt.Println(passes[i])
			result = passes[i] + 1
			break
		}
	}
	if expected != -1 {
		if result != expected {
			fmt.Println("ERROR: got", result, "expected", expected)
		} else {
			fmt.Println("RIGHT ANSWER:", result)
		}
	}
	return result
}

func main() {
	defer duration(time.Now(), "main")

	// fmt.Println(decode("FBFBBFFRLR"))
	// fmt.Println(decode("BFFFBBFRRR"))
	// fmt.Println(decode("FFFBBBFRRR"))
	// fmt.Println(decode("BBFFBBFRLL"))

	part1(input, 922)
	part2(input, 747)

	{
		var memStats runtime.MemStats
		runtime.ReadMemStats(&memStats)
		fmt.Printf("Heap memory (in bytes): %d\n", memStats.HeapAlloc)
		fmt.Printf("Number of garbage collections: %d\n", memStats.NumGC)
	}
}

/*
$ go run main.go
RIGHT ANSWER: 922
part 1 56.527µs
RIGHT ANSWER: 747
part 2 117.366µs
Heap memory (in bytes): 167640
Number of garbage collections: 0
main 262.21µs
*/
