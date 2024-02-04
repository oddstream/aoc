// https://adventofcode.com/2020/day/11
package main

import (
	"bufio"
	"bytes"
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
	return Point{y: p.y + q.y, x: p.x + q.x}
}

var directions []Point = []Point{
	{y: -1, x: 0},  // n
	{y: -1, x: 1},  // ne
	{y: 0, x: 1},   // e
	{y: 1, x: 1},   // se
	{y: 1, x: 0},   // s
	{y: 1, x: -1},  // sw
	{y: 0, x: -1},  // w
	{y: -1, x: -1}, // nw

}

func ingrid(grid [][]byte, p Point) bool {
	if p.y < 0 || p.x < 0 {
		return false
	}
	if p.y >= len(grid) || p.x >= len(grid[0]) {
		return false
	}
	return true
}

func neighbours(grid [][]byte, p Point) int {
	var n int
	for _, dir := range directions {
		var q Point = p.add(dir)
		if ingrid(grid, q) && grid[q.y][q.x] == '#' {
			n += 1
		}
	}
	return n
}

func visible(grid [][]byte, p Point) int {
	var n int
	for _, dir := range directions {
		var q Point = p.add(dir)
		for {
			if !ingrid(grid, q) {
				break
			}
			if grid[q.y][q.x] == 'L' {
				break
			}
			if grid[q.y][q.x] == '#' {
				n += 1
				break
			}
			q = q.add(dir)
		}
	}
	return n
}

func round1(src [][]byte) (dst [][]byte, changed bool) {
	dst = make([][]byte, len(src))
	for i := range src {
		dst[i] = bytes.Clone(src[i])
	}
	// changed = false
	for y := 0; y < len(src); y++ {
		for x := 0; x < len(src[0]); x++ {
			switch src[y][x] {
			case 'L':
				if neighbours(src, Point{y: y, x: x}) == 0 {
					dst[y][x] = '#'
					changed = true
				}
			case '#':
				if neighbours(src, Point{y: y, x: x}) >= 4 {
					dst[y][x] = 'L'
					changed = true
				}
			}
		}
	}
	return dst, changed
}

func round2(src [][]byte) (dst [][]byte, changed bool) {
	dst = make([][]byte, len(src))
	for i := range src {
		dst[i] = bytes.Clone(src[i])
	}
	// changed = false
	for y := 0; y < len(src); y++ {
		for x := 0; x < len(src[0]); x++ {
			switch src[y][x] {
			case 'L':
				if visible(src, Point{y: y, x: x}) == 0 {
					dst[y][x] = '#'
					changed = true
				}
			case '#':
				if visible(src, Point{y: y, x: x}) >= 5 {
					dst[y][x] = 'L'
					changed = true
				}
			}
		}
	}
	return dst, changed
}

func countChars(grid [][]byte, ch byte) int {
	var result int
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			if grid[y][x] == ch {
				result += 1
			}
		}
	}
	return result
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

	var grid [][]byte
	scanner := bufio.NewScanner(strings.NewReader(in))
	for scanner.Scan() {
		grid = append(grid, []byte(scanner.Text()))
	}
	for i := 0; i < 100; i++ {
		var changed bool
		grid, changed = round1(grid)
		if !changed {
			fmt.Println(i, "rounds")
			break
		}
	}

	result = countChars(grid, '#')
	report(expected, result)
	return result
}

func part2(in string, expected int) (result int) {
	defer duration(time.Now(), "part 2")

	var grid [][]byte
	scanner := bufio.NewScanner(strings.NewReader(in))
	for scanner.Scan() {
		grid = append(grid, []byte(scanner.Text()))
	}
	for i := 0; i < 100; i++ {
		var changed bool
		grid, changed = round2(grid)
		if !changed {
			fmt.Println(i, "rounds")
			break
		}
	}

	result = countChars(grid, '#')
	report(expected, result)
	return result
}

func main() {
	defer duration(time.Now(), "main")

	// part1(test1, 37)   // 5 rounds
	part1(input, 2108) // 70 rounds
	// part2(test1, 26)   // 6 rounds
	part2(input, 1897) // 83 rounds

	{
		var memStats runtime.MemStats
		runtime.ReadMemStats(&memStats)
		fmt.Printf("Heap memory (in bytes): %d\n", memStats.HeapAlloc)
		fmt.Printf("Number of garbage collections: %d\n", memStats.NumGC)
	}
}

/*
$ go run main.go
70 rounds
RIGHT ANSWER: 2108
part 1 10.506163ms
83 rounds
RIGHT ANSWER: 1897
part 2 25.761599ms
Heap memory (in bytes): 1885504
Number of garbage collections: 0
main 36.393296ms
*/
