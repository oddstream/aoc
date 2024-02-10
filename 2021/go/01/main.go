// https://adventofcode.com/2021/day/1
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

// or we could use K&R p61
func atoi(s string) int {
	if n, err := strconv.Atoi(s); err == nil {
		return n
	} else {
		fmt.Println(err)
	}
	return 0
}

func part1(in string, expected int) (result int) {
	defer duration(time.Now(), "part 1")

	var prev int
	scanner := bufio.NewScanner(strings.NewReader(in))
	for scanner.Scan() {
		var depth int = atoi(scanner.Text())
		if prev != 0 {
			if depth > prev {
				result += 1
			}
		}
		prev = depth
	}
	report(expected, result)
	return result
}

func part2(in string, expected int) (result int) {
	defer duration(time.Now(), "part 2")

	var depths []int

	scanner := bufio.NewScanner(strings.NewReader(in))
	for scanner.Scan() {
		depths = append(depths, atoi(scanner.Text()))
	}

	var prev int
	for i := 0; i < len(depths)-2; i++ {
		var sum3 int = depths[i] + depths[i+1] + depths[i+2]
		if prev != 0 {
			if sum3 > prev {
				result += 1
			}
		}
		prev = sum3
	}
	report(expected, result)
	return result
}

func main() {
	defer duration(time.Now(), "main")

	part1(test1, 7)
	part1(input, 1195)
	part2(test1, 5)
	part2(input, 1235)

	{
		var memStats runtime.MemStats
		runtime.ReadMemStats(&memStats)
		fmt.Printf("Heap memory (in bytes): %d\n", memStats.HeapAlloc)
		fmt.Printf("Number of garbage collections: %d\n", memStats.NumGC)
	}
}

/*
$ go run main.go
*/
