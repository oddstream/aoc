// https://adventofcode.com/2020/day/1
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

func sum(nums ...int) int {
	var result int = 0
	for _, num := range nums {
		result += num
	}
	return result
}

func product(nums ...int) int {
	var result int = 1
	for _, num := range nums {
		result *= num
	}
	return result
}

func pairs(s []int) <-chan []int {
	c := make(chan []int)
	go func() {
		defer close(c)
		for i, x := range s {
			for _, y := range s[i+1:] {
				c <- []int{x, y}
			}
		}
	}()
	return c
}

func triples(s []int) <-chan []int {
	c := make(chan []int)
	go func() {
		defer close(c)
		for i, x := range s {
			for _, y := range s[i+1:] {
				for _, z := range s[i+2:] {
					c <- []int{x, y, z}
				}
			}
		}
	}()
	return c
}

func part1(in string, expected int) (result int) {
	defer duration(time.Now(), "part 1")

	var nums []int
	scanner := bufio.NewScanner(strings.NewReader(in))
	for scanner.Scan() {
		nums = append(nums, atoi(scanner.Text()))
	}
	for pair := range pairs(nums) {
		if sum(pair...) == 2020 {
			result = product(pair...)
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

func part2(in string, expected int) (result int) {
	defer duration(time.Now(), "part 2")

	var nums []int
	scanner := bufio.NewScanner(strings.NewReader(in))
	for scanner.Scan() {
		nums = append(nums, atoi(scanner.Text()))
	}
	for triple := range triples(nums) {
		if sum(triple...) == 2020 {
			result = product(triple...)
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

	part1(test1, 514579)
	part1(input, 1019371)
	part2(test1, 241861950)
	part2(input, 278064990)

	// for triple := range triples([]int{0, 1, 2, 3}) {
	// 	fmt.Println(triple, sum(triple...), product(triple...))
	// }
	{
		var memStats runtime.MemStats
		runtime.ReadMemStats(&memStats)
		fmt.Printf("Heap memory (in bytes): %d\n", memStats.HeapAlloc)
		fmt.Printf("Number of garbage collections: %d\n", memStats.NumGC)
	}
}

/*
$ go run main.go
RIGHT ANSWER: 514579
part 1 27.374µs
RIGHT ANSWER: 1019371
part 1 249.08µs
RIGHT ANSWER: 241861950
part 2 19.444µs
RIGHT ANSWER: 278064990
part 2 44.770967ms
Heap memory (in bytes): 150776
Number of garbage collections: 15
main 45.131486ms
*/
