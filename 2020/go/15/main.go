// https://adventofcode.com/2020/day/15
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
		tokens := strings.Split(scanner.Text(), ",")
		for _, token := range tokens {
			nums = append(nums, atoi(token))
		}
	}
	for turn := len(nums); turn < 2020; turn++ {
		i := len(nums) - 1
		lastnum := nums[i]
		for i > 0 {
			i -= 1
			if nums[i] == lastnum {
				new := len(nums) - 1 - i
				nums = append(nums, new)
				goto next
			}
		}
		nums = append(nums, 0)
	next:
		// fmt.Println(turn+1, nums)
	}
	result = nums[len(nums)-1]
	report(expected, result)
	return result
}

func part2(in string, expected int) (result int) {
	defer duration(time.Now(), "part 2")

	var memory map[int]int = make(map[int]int)

	scanner := bufio.NewScanner(strings.NewReader(in))
	for scanner.Scan() {
		tokens := strings.Split(scanner.Text(), ",")
		for turn, token := range tokens {
			// memory contains (actual turn number - 1)
			memory[atoi(token)] = turn
		}
	}

	var spoken int

	for turn := len(memory) + 1; turn < 30_000_000; turn++ {
		lastSpoken, spokenBefore := memory[spoken]
		memory[spoken] = turn - 1
		if spokenBefore {
			spoken = turn - 1 - lastSpoken
		} else {
			spoken = 0
		}
	}
	result = spoken
	report(expected, result)
	return result
}

func main() {
	defer duration(time.Now(), "main")

	part1(test1, 436)
	part1(input, 206)
	part2(test1, 175594)
	part2(input, 955)

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
