// https://adventofcode.com/2020/day/3
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

type Point struct {
	y, x int
}

func (p Point) add(q Point) Point {
	return Point{p.y + q.y, p.x + q.x}
}

func duration(invocation time.Time, name string) {
	fmt.Println(name, time.Since(invocation))
}

func part1(in string, expected int) (result int) {
	defer duration(time.Now(), "part 1")

	var trees []string
	scanner := bufio.NewScanner(strings.NewReader(in))
	for scanner.Scan() {
		trees = append(trees, scanner.Text())
	}
	var p Point = Point{}
	for p.y < len(trees) {
		if trees[p.y][p.x] == '#' {
			result += 1
		}
		p.y += 1
		p.x += 3
		p.x %= len(trees[0])
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

	var trees []string
	scanner := bufio.NewScanner(strings.NewReader(in))
	for scanner.Scan() {
		trees = append(trees, scanner.Text())
	}
	result = 1
	for _, slope := range []Point{{y: 1, x: 1}, {y: 1, x: 3}, {y: 1, x: 5}, {y: 1, x: 7}, {y: 2, x: 1}} {
		var p Point = Point{}
		var collisions int
		for p.y < len(trees) {
			if trees[p.y][p.x] == '#' {
				collisions += 1
			}
			p = p.add(slope)
			p.x %= len(trees[0])
		}
		result *= collisions
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

	part1(test1, 7)
	part1(input, 189)
	part2(test1, 336)
	part2(input, 1718180100)

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
