// https://adventofcode.com/2020/day/10
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

//go:embed test2.txt
var test2 string

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

func part1(in string, expected int) (result int) {
	defer duration(time.Now(), "part 1")

	var nums []int
	scanner := bufio.NewScanner(strings.NewReader(in))
	for scanner.Scan() {
		nums = append(nums, atoi(scanner.Text()))
	}
	slices.Sort(nums)
	nums = append(nums, nums[len(nums)-1]+3) // append target adapter (always +3)
	var prev, jolt1, jolt3 int
	for i := 0; i < len(nums); i++ {
		if nums[i] == prev+1 {
			jolt1 += 1
		} else if nums[i] == prev+3 {
			jolt3 += 1
		}
		prev = nums[i]
	}
	result = jolt1 * jolt3
	report(expected, result)
	return result
}

func part2(in string, expected int) (result int) {
	defer duration(time.Now(), "part 2")

	var nums []int
	scanner := bufio.NewScanner(strings.NewReader(in))
	for scanner.Scan() {
		nums = append(nums, atoi(scanner.Text()))
	}
	slices.Sort(nums)
	nums = append(nums, nums[len(nums)-1]+3) // append target adapter (always +3)

	var counter map[int]int = map[int]int{
		0: 1, // 1 way to reach 0
	}
	// the counter map has the lovely property
	// that it will return 0
	// if the key is not in the map
	for _, adapter := range nums {
		counter[adapter] = counter[adapter-3] + counter[adapter-2] + counter[adapter-1]
	}
	result = counter[nums[len(nums)-1]]
	report(expected, result)
	return result
}

func main() {
	defer duration(time.Now(), "main")

	// part1(test1, 35)
	// part1(test2, 220)
	part1(input, 1856)
	// part2(test1, 8)
	// part2(test2, 19208)
	part2(input, 2314037239808)

	{
		var memStats runtime.MemStats
		runtime.ReadMemStats(&memStats)
		fmt.Printf("Heap memory (in bytes): %d\n", memStats.HeapAlloc)
		fmt.Printf("Number of garbage collections: %d\n", memStats.NumGC)
	}
}

/*
$ go run main.go
RIGHT ANSWER: 1856
part 1 37.946µs
RIGHT ANSWER: 2314037239808
part 2 45.477µs
Heap memory (in bytes): 151560
Number of garbage collections: 0
main 178.633µs
*/
