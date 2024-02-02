// https://adventofcode.com/2020/day/6
package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"runtime"
	"strings"
	"time"
)

//go:embed input.txt
var input string

func duration(invocation time.Time, name string) {
	fmt.Println(name, time.Since(invocation))
}

func score1(lines []string) int {
	var people map[rune]struct{} = map[rune]struct{}{}
	for _, line := range lines {
		for _, r := range line {
			people[r] = struct{}{}
		}
	}
	return len(people)
}

func part1(in string, expected int) (result int) {
	defer duration(time.Now(), "part 1")

	scanner := bufio.NewScanner(strings.NewReader(in))
	var lines []string
	for scanner.Scan() {
		if scanner.Text() == "" {
			result += score1(lines)
			lines = nil
		} else {
			lines = append(lines, scanner.Text())
		}
	}
	if len(lines) > 0 {
		result += score1(lines)
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

func score2(lines []string) int {
	var people map[rune]int = make(map[rune]int)
	for _, line := range lines {
		for _, r := range line {
			people[r] += 1
		}
	}
	var count int
	for r := range people {
		if people[r] == len(lines) {
			count += 1
		}
	}
	return count
}

func part2(in string, expected int) (result int) {
	defer duration(time.Now(), "part 2")

	scanner := bufio.NewScanner(strings.NewReader(in))
	var lines []string
	for scanner.Scan() {
		if scanner.Text() == "" {
			result += score2(lines)
			lines = nil
		} else {
			lines = append(lines, scanner.Text())
		}
	}
	if len(lines) > 0 {
		result += score2(lines)
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

	part1(input, 6387)
	part2(input, 3039)

	{
		var memStats runtime.MemStats
		runtime.ReadMemStats(&memStats)
		fmt.Printf("Heap memory (in bytes): %d\n", memStats.HeapAlloc)
		fmt.Printf("Number of garbage collections: %d\n", memStats.NumGC)
	}
}

/*
$ go run main.go
RIGHT ANSWER: 6387
part 1 704.9µs
RIGHT ANSWER: 3039
part 2 890.478µs
Heap memory (in bytes): 670304
Number of garbage collections: 0
main 1.68777ms
*/
