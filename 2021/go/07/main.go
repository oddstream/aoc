// https://adventofcode.com/2021/day/7
package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"runtime"
	"slices"
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

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
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

	var crabs []int
	scanner := bufio.NewScanner(strings.NewReader(in))
	for scanner.Scan() {
		for _, pos := range strings.Split(scanner.Text(), ",") {
			crabs = append(crabs, atoi(pos))
		}
	}
	// some bright sparks pointed out that the answer is the geometric median
	// of the crab positions, so ...
	slices.Sort(crabs)
	var idx int
	if len(crabs)%2 == 1 {
		idx = crabs[(len(crabs)+1)/2]
	} else {
		var left int = crabs[len(crabs)/2-1]
		var right int = crabs[len(crabs)/2]
		idx = (int)((left + right) / 2)
	}
	for _, pos := range crabs {
		result += abs(pos - idx)
	}
	report(expected, result)
	return result
}

func part2(in string, expected int) (result int) {
	defer duration(time.Now(), "part 2")

	var crabs []int
	scanner := bufio.NewScanner(strings.NewReader(in))
	for scanner.Scan() {
		for _, pos := range strings.Split(scanner.Text(), ",") {
			crabs = append(crabs, atoi(pos))
		}
	}
	// some bright sparks argued that the answer is the mean
	// of the crab positions, so ...
	var sum int
	for _, pos := range crabs {
		sum += pos
	}
	var mean int
	if sum&1 == 1 {
		mean = (sum + 1) / len(crabs)
	} else {
		mean = sum / len(crabs)
	}

	for _, pos := range crabs {
		var n int = abs(pos - mean)
		result += (n * (n + 1) / 2)
	}
	report(expected, result)
	return result
}

func main() {
	defer duration(time.Now(), "main")

	part1(test1, 37)
	part1(input, 329389)
	part2(test1, 168)
	part2(input, 86397080)

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
