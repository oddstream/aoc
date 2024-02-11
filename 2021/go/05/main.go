// https://adventofcode.com/2021/day/5
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
	x, y int
}

type Line struct {
	x1, y1, x2, y2 int
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

func sign(n int) int {
	if n < 0 {
		return -1
	} else if n > 0 {
		return 1
	}
	return 0
}

func display(grid map[Point]int) {
	for y := 0; y < 10; y++ {
		for x := 0; x < 10; x++ {
			var n int = grid[Point{x: x, y: y}]
			if n == 0 {
				fmt.Print(".")
			} else {
				fmt.Print(n)
			}
		}
		fmt.Println()
	}
}

func part1(in string, expected int) (result int) {
	defer duration(time.Now(), "part 1")

	var lines []Line
	scanner := bufio.NewScanner(strings.NewReader(in))
	for scanner.Scan() {
		var line Line
		if n, err := fmt.Sscanf(scanner.Text(), "%d,%d -> %d,%d", &line.x1, &line.y1, &line.x2, &line.y2); n != 4 {
			fmt.Println(err)
			break
		}
		if (line.x1 == line.x2) || (line.y1 == line.y2) {
			lines = append(lines, line)
		}
	}

	var grid map[Point]int = make(map[Point]int)
	for _, line := range lines {
		var x, y int = line.x1, line.y1
		var dx, dy int = sign(line.x2 - line.x1), sign(line.y2 - line.y1)
		for {
			grid[Point{x: x, y: y}] += 1
			if x == line.x2 && y == line.y2 {
				break
			}
			x += dx
			y += dy
		}
	}

	for _, count := range grid {
		if count > 1 {
			result += 1
		}
	}
	report(expected, result)
	return result
}

func part2(in string, expected int) (result int) {
	defer duration(time.Now(), "part 2")

	var lines []Line
	scanner := bufio.NewScanner(strings.NewReader(in))
	for scanner.Scan() {
		var line Line
		if n, err := fmt.Sscanf(scanner.Text(), "%d,%d -> %d,%d", &line.x1, &line.y1, &line.x2, &line.y2); n != 4 {
			fmt.Println(err)
			break
		}
		lines = append(lines, line)
	}

	var grid map[Point]int = make(map[Point]int)
	for _, line := range lines {
		var x, y int = line.x1, line.y1
		var dx, dy int = sign(line.x2 - line.x1), sign(line.y2 - line.y1)
		for {
			grid[Point{x: x, y: y}] += 1
			if x == line.x2 && y == line.y2 {
				break
			}
			x += dx
			y += dy
		}
	}
	for _, count := range grid {
		if count > 1 {
			result += 1
		}
	}
	report(expected, result)
	return result
}

func main() {
	defer duration(time.Now(), "main")

	// part1(test1, 5)
	part1(input, 6461)
	// part2(test1, 12)
	part2(input, 18065)

	{
		var memStats runtime.MemStats
		runtime.ReadMemStats(&memStats)
		fmt.Printf("Heap memory (in bytes): %d\n", memStats.HeapAlloc)
		fmt.Printf("Number of garbage collections: %d\n", memStats.NumGC)
	}
}

/*
$ go run main.go
RIGHT ANSWER: 6461
part 1 19.545491ms
RIGHT ANSWER: 18065
part 2 20.947885ms
Heap memory (in bytes): 11600640
Number of garbage collections: 4
main 40.544495ms
*/
