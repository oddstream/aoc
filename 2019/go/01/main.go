// https://adventofcode.com/2019/day/1 The Tyranny of the Rocket Equation
package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"strings"
	"time"
)

//go:embed input.txt
var input string

func duration(invocation time.Time, name string) {
	fmt.Println(name, time.Since(invocation))
}

func partOne(in string, expected int) {
	defer duration(time.Now(), "part 1")

	var totalMass int
	scanner := bufio.NewScanner(strings.NewReader(in))
	for scanner.Scan() {
		var mass int
		if n, err := fmt.Sscanf(scanner.Text(), "%d", &mass); n != 1 {
			fmt.Println(err)
		}
		totalMass += (mass / 3) - 2
	}

	var result int = totalMass
	if expected != -1 {
		if result != expected {
			fmt.Println("ERROR: got", result, "expected", expected)
		} else {
			fmt.Println("RIGHT:", result)
		}
	}
}

func fuel(mass int) int {
	var f int = (mass / 3) - 2
	if f <= 0 {
		return 0
	}
	return f + fuel(f)
}

func partTwo(in string, expected int) {
	defer duration(time.Now(), "part 2")

	var totalFuel int
	scanner := bufio.NewScanner(strings.NewReader(in))
	for scanner.Scan() {
		var mass int
		if n, err := fmt.Sscanf(scanner.Text(), "%d", &mass); n != 1 {
			fmt.Println(err)
		}
		totalFuel += fuel(mass)
	}

	var result = totalFuel
	if expected != -1 {
		if result != expected {
			fmt.Println("ERROR: got", result, "expected", expected)
		} else {
			fmt.Println("RIGHT:", result)
		}
	}
}

func main() {
	defer duration(time.Now(), "main")

	partOne(input, 3223398)
	// fmt.Println((14 / 3) - 2)
	// fmt.Println(fuel(14))
	// fmt.Println(fuel(1969))
	// fmt.Println(fuel(100756))
	partTwo(input, 4832253)
}

/*
$ go run main.go
*/
