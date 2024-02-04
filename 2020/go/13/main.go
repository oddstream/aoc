// https://adventofcode.com/2020/day/13
package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"math"
	"runtime"
	"strconv"
	"strings"
	"time"
)

//go:embed test1.txt
var test1 string

//go:embed input.txt
var input string

// all the bus numbers are prime; Chinese remainder theorem incoming ...

// or we could use K&R p61
func atoi(s string) int {
	if n, err := strconv.Atoi(s); err == nil {
		return n
	} else {
		fmt.Println(err)
	}
	return 0
}

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

	var earliest int
	var busses []int
	scanner := bufio.NewScanner(strings.NewReader(in))
	if scanner.Scan() {
		earliest = atoi(scanner.Text())
	}
	if scanner.Scan() {
		tokens := strings.Split(scanner.Text(), ",")
		for _, token := range tokens {
			if token == "x" {
				continue
			}
			busses = append(busses, atoi(token))
		}
	}
	// fmt.Println(earliest)
	// fmt.Println(busses)

	// var mult int = 1
	// for _, bus := range busses {
	// 	if !big.NewInt(int64(bus)).ProbablyPrime(0) {
	// 		fmt.Println(bus, "is not prime")
	// 	}
	// 	mult *= bus
	// }
	// fmt.Println("mult", mult)	// 1473355587699697

	for t := earliest; t < math.MaxInt64; t++ {
		for _, bus := range busses {
			if t%bus == 0 {
				result = (t - earliest) * bus
				goto report
			}
		}
	}
report:
	report(expected, result)
	return result
}

func part2(in string, expected int) (result int) {
	defer duration(time.Now(), "part 2")

	var busses []int
	scanner := bufio.NewScanner(strings.NewReader(in))
	scanner.Scan() // ignore earliest departure time
	if scanner.Scan() {
		tokens := strings.Split(scanner.Text(), ",")
		for _, token := range tokens {
			if token == "x" {
				busses = append(busses, 0)
			} else {
				busses = append(busses, atoi(token))
			}
		}
	}

	/*
		Iterate over the bus schedule and checks if the departure time is 0
		(meaning the bus is not running).
		If the bus is running, find the next departure time that is a multiple of the bus number.
		Do this by repeatedly adding the departure time to the result
		until the result is a multiple of the bus number.
		Once the code finds a departure time that satisfies all the buses in the schedule,
		it prints the departure time and exits.
		Does not use the Chinese Remainder Theorem.
	*/

	var depart int = 1
	for i, bus := range busses {
		if bus == 0 {
			continue
		}
		for {
			result += depart
			if (result+i)%bus == 0 {
				depart *= bus
				break
			}
		}
	}

	report(expected, result)
	return result
}

func main() {
	defer duration(time.Now(), "main")

	// part1(test1, 295)
	part1(input, 174)
	// part2(test1, 1068781)
	part2(input, 780601154795940)

	{
		var memStats runtime.MemStats
		runtime.ReadMemStats(&memStats)
		fmt.Printf("Heap memory (in bytes): %d\n", memStats.HeapAlloc)
		fmt.Printf("Number of garbage collections: %d\n", memStats.NumGC)
	}
}

/*
$ go run main.go
RIGHT ANSWER: 174
part 1 44.12µs
RIGHT ANSWER: 780601154795940
part 2 19.453µs
Heap memory (in bytes): 148656
Number of garbage collections: 0
main 176.319µs
*/
