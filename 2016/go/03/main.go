// https://adventofcode.com/2016/day/3
package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"slices"
	"strings"
	"time"
)

//go:embed input.txt
var input string

// input.txt has 1635 rows, 1635/3 == 545

func duration(invocation time.Time, name string) {
	fmt.Println(name, time.Since(invocation))
}

func isTriangle(arr []int) bool {
	slices.Sort(arr)
	return arr[0]+arr[1] > arr[2]
}

func partOne() int {
	defer duration(time.Now(), "partOne")
	var result = 0
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		var a, b, c int
		var err error
		if _, err = fmt.Sscanf(scanner.Text(), " %d %d %d", &a, &b, &c); err != nil {
			fmt.Println(err)
		}
		if isTriangle([]int{a, b, c}) {
			result += 1
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	return result
}

func partTwo() int {
	defer duration(time.Now(), "partTwo")
	var result = 0
	var lines [][]int

	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		var a, b, c int
		var err error
		if _, err = fmt.Sscanf(scanner.Text(), " %d %d %d", &a, &b, &c); err != nil {
			fmt.Println(err)
		}
		lines = append(lines, []int{a, b, c})
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	for i := 0; i < len(lines); i += 3 {
		for j := 0; j < 3; j++ {
			a := lines[i][j]
			b := lines[i+1][j]
			c := lines[i+2][j]
			if isTriangle([]int{a, b, c}) {
				result += 1
			}
		}
	}
	return result
}

func main() {
	defer duration(time.Now(), "main")

	fmt.Println("part one", partOne()) // 862
	fmt.Println("part two", partTwo()) // 1577
}
