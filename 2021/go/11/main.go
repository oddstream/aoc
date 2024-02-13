// https://adventofcode.com/2021/day/11
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

type Point struct {
	y, x int
}

func (p Point) add(q Point) Point {
	return Point{y: p.y + q.y, x: p.x + q.x}
}

type Grid [][]int

func (g Grid) ingrid(p Point) bool {
	return p.x >= 0 && p.y >= 0 && p.x < len(g[0]) && p.y < len(g)
}

func (g Grid) get(p Point) int {
	return g[p.y][p.x]
}

func (g Grid) set(p Point, value int) {
	g[p.y][p.x] = value
}

func (g Grid) inc(p Point) {
	g[p.y][p.x] += 1
}

func (g Grid) foreach(fn func(Point)) {
	for y := 0; y < len(g); y++ {
		for x := 0; x < len(g[0]); x++ {
			fn(Point{y: y, x: x})
		}
	}
}

var directions []Point = []Point{
	{y: -1, x: 0},  // N
	{y: -1, x: 1},  // NE
	{y: 0, x: 1},   // E
	{y: 1, x: 1},   // SE
	{y: 1, x: 0},   // S
	{y: 1, x: -1},  // SW
	{y: 0, x: -1},  // W
	{y: -1, x: -1}, // NW
}

func (g Grid) neighbours(p Point) []Point {
	var out []Point
	for _, dir := range directions {
		var q Point = p.add(dir)
		if g.ingrid(q) {
			out = append(out, q)
		}
	}
	return out
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

	var grid Grid
	scanner := bufio.NewScanner(strings.NewReader(in))
	for scanner.Scan() {
		var row []int
		for _, num := range strings.Split(scanner.Text(), "") {
			row = append(row, atoi(num))
		}
		grid = append(grid, row)
	}

	// tried and failed with a recursive solution, so there's this ...
	for step := 0; step < 100; step++ {
		var q []Point // only for >9
		var flashed map[Point]struct{} = make(map[Point]struct{})

		// step 1 - the energy level of each octopus increases by 1
		grid.foreach(func(p Point) {
			grid.inc(p)
			if grid.get(p) > 9 {
				q = append(q, p)
			}
		})
		// step 2
		for len(q) > 0 {
			var p Point
			p, q = q[0], q[1:]
			if _, ok := flashed[p]; ok {
				continue
			}
			flashed[p] = struct{}{}
			for _, np := range grid.neighbours(p) {
				grid.inc(np)
				if grid.get(np) > 9 {
					q = append(q, np)
				}
			}
		}
		// step 3
		for p := range flashed {
			grid.set(p, 0)
		}
		result += len(flashed)
	}
	report(expected, result)
	return result
}

func part2(in string, expected int) (result int) {
	defer duration(time.Now(), "part 2")

	var grid Grid
	scanner := bufio.NewScanner(strings.NewReader(in))
	for scanner.Scan() {
		var row []int
		for _, num := range strings.Split(scanner.Text(), "") {
			row = append(row, atoi(num))
		}
		grid = append(grid, row)
	}

	for step := 1; step < 500; step++ { // nb step count is 1-based, not 0-based
		var q []Point // only for >9
		var flashed map[Point]struct{} = make(map[Point]struct{})

		// step 1 - the energy level of each octopus increases by 1
		grid.foreach(func(p Point) {
			grid.inc(p)
			if grid.get(p) > 9 {
				q = append(q, p)
			}
		})
		// step 2
		for len(q) > 0 {
			var p Point
			p, q = q[0], q[1:]
			if _, ok := flashed[p]; ok {
				continue
			}
			flashed[p] = struct{}{}
			for _, np := range grid.neighbours(p) {
				grid.inc(np)
				if grid.get(np) > 9 {
					q = append(q, np)
				}
			}
		}
		// step 3
		for p := range flashed {
			grid.set(p, 0)
		}

		if len(flashed) == 100 {
			result = step
			break
		}
	}
	report(expected, result)
	return result
}

func main() {
	defer duration(time.Now(), "main")

	// part1(test1, 1656)
	part1(input, 1694)
	// part2(test1, 195)
	part2(input, 346)

	{
		var memStats runtime.MemStats
		runtime.ReadMemStats(&memStats)
		fmt.Printf("Heap memory (in bytes): %d\n", memStats.HeapAlloc)
		fmt.Printf("Number of garbage collections: %d\n", memStats.NumGC)
	}
}

/*
$ go run main.go
RIGHT ANSWER: 1694
part 1 971.56µs
RIGHT ANSWER: 346
part 2 2.966655ms
Heap memory (in bytes): 3342552
Number of garbage collections: 0
main 4.081892ms
*/
