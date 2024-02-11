// https://adventofcode.com/2021/day/6
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

func simulate(fish []int, days int) (total int) {

	for ; days > 1; days-- {
		// decrease the fish-countdown-timers by shifting the array to the left
		// [0, 155, 31, 33, 38, 43, 0, 0, 0]
		fish = append(fish[1:], fish[:1]...)
		// [155, 31, 33, 38, 43, 0, 0, 0, 0]
		// new fish start with timer == 8 (index 7)
		fish[7] += fish[0]
		// [155, 31, 33, 38, 43, 0, 0, 155, 0]
	}

	for _, n := range fish {
		total += n
	}
	return total
}

func run(in string, days int, expected int) (result int) {
	defer duration(time.Now(), fmt.Sprintf("run %d days", days))

	var fish []int = []int{0, 0, 0, 0, 0, 0, 0, 0, 0} // 1+8
	scanner := bufio.NewScanner(strings.NewReader(in))
	for scanner.Scan() {
		for _, num := range strings.Split(scanner.Text(), ",") {
			fish[atoi(num)] += 1
		}
	}
	fmt.Println(fish)
	result = simulate(fish, days)
	report(expected, result)
	return result
}

func main() {
	defer duration(time.Now(), "main")

	// run(test1, 18, 26)
	// run(test1, 80, 5934)
	run(input, 80, 372300)
	run(input, 256, 1675781200288)

	{
		var memStats runtime.MemStats
		runtime.ReadMemStats(&memStats)
		fmt.Printf("Heap memory (in bytes): %d\n", memStats.HeapAlloc)
		fmt.Printf("Number of garbage collections: %d\n", memStats.NumGC)
	}
}

/*
$ go run main.go
[0 155 31 33 38 43 0 0 0]
RIGHT ANSWER: 372300
run 80 days 49.974µs
[0 155 31 33 38 43 0 0 0]
RIGHT ANSWER: 1675781200288
run 256 days 12.213µs
Heap memory (in bytes): 159152
Number of garbage collections: 0
main 149.754µs
*/
