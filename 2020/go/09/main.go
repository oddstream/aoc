// https://adventofcode.com/2020/day/9
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

// or we could use K&R p61
func atoi(s string) int {
	if n, err := strconv.Atoi(s); err == nil {
		return n
	} else {
		fmt.Println(err)
	}
	return 0
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

func valid(nums []int, preamble, pos int) bool {
	for i := pos - preamble; i < pos; i++ {
		for j := i + 1; j < pos; j++ {
			if nums[i]+nums[j] == nums[pos] {
				return true
			}
		}
	}
	return false
}

func part1(in string, preamble, expected int) (result int) {
	defer duration(time.Now(), "part 1")

	var nums []int
	scanner := bufio.NewScanner(strings.NewReader(in))
	for scanner.Scan() {
		nums = append(nums, atoi(scanner.Text()))
	}
	for pos := preamble; pos < len(nums); pos++ {
		if !valid(nums, preamble, pos) {
			result = nums[pos]
			break
		}
	}

	report(expected, result)
	return result
}

func part2(in string, invalid, expected int) (result int) {
	defer duration(time.Now(), "part 2")

	var nums []int
	sum := func(a, b int) int {
		var n int
		for i := a; i <= b; i++ {
			n += nums[i]
		}
		return n
	}
	scanner := bufio.NewScanner(strings.NewReader(in))
	for scanner.Scan() {
		nums = append(nums, atoi(scanner.Text()))
	}
	// contiguous set of at least two numbers that add up to invalid
	for anchor := 0; anchor < len(nums)-2; anchor++ {
		for end := anchor + 1; end < len(nums)-1; end++ {
			if sum(anchor, end) == invalid {
				// add together the smallest and largest number in this contiguous range
				var r []int = nums[anchor : end+1]
				// increasingly weird that the second slice number must be +1
				// s[:5] up to (but excluding) s[5]
				// sort.Slice(r, func(a, b int) bool { return r[a] < r[b] })
				slices.Sort(r)
				result = r[0] + r[len(r)-1]
				goto report
			}
		}
	}
report:
	report(expected, result)
	return result
}

func main() {
	defer duration(time.Now(), "main")

	// part1(test1, 5, 127)
	part1(input, 25, 69316178)
	// part2(test1, 127, 62)
	part2(input, 69316178, 9351526)

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
