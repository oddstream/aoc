// https://adventofcode.com/2021/day/2
package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"runtime"
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

	var horz, depth int
	scanner := bufio.NewScanner(strings.NewReader(in))
	for scanner.Scan() {
		var dir string
		var amt int
		if n, err := fmt.Sscanf(scanner.Text(), "%s %d", &dir, &amt); n != 2 {
			fmt.Println(err)
			break
		}
		switch dir {
		case "forward":
			horz += amt
		case "down":
			depth += amt
		case "up":
			depth -= amt
		default:
			fmt.Println("unknown direction", dir)
		}
	}
	result = horz * depth
	report(expected, result)
	return result
}

func part2(in string, expected int) (result int) {
	defer duration(time.Now(), "part 2")

	var horz, depth, aim int
	scanner := bufio.NewScanner(strings.NewReader(in))
	for scanner.Scan() {
		var dir string
		var amt int
		if n, err := fmt.Sscanf(scanner.Text(), "%s %d", &dir, &amt); n != 2 {
			fmt.Println(err)
			break
		}
		switch dir {
		case "forward":
			horz += amt
			depth += aim * amt
		case "down":
			aim += amt
		case "up":
			aim -= amt
		default:
			fmt.Println("unknown direction", dir)
		}
	}
	result = horz * depth
	report(expected, result)
	return result
}

func main() {
	defer duration(time.Now(), "main")

	// part1(test1, 150)
	part1(input, 2150351)
	// part2(test1, 900)
	part2(input, 1842742223)

	{
		var memStats runtime.MemStats
		runtime.ReadMemStats(&memStats)
		fmt.Printf("Heap memory (in bytes): %d\n", memStats.HeapAlloc)
		fmt.Printf("Number of garbage collections: %d\n", memStats.NumGC)
	}
}

/*
$ go run main.go
RIGHT ANSWER: 2150351
part 1 667.756µs
RIGHT ANSWER: 1842742223
part 2 626.949µs
Heap memory (in bytes): 353056
Number of garbage collections: 0
main 1.383439ms
*/
