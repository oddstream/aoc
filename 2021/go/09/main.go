// https://adventofcode.com/2021/day/9
package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"runtime"
	"slices"
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

func (g Grid) lowPoint(p Point) bool {
	var height int = g.get(p)
	for _, dir := range directions {
		var q Point = p.add(dir)
		if g.ingrid(q) {
			if g.get(q) <= height {
				return false
			}
		}
	}
	return true
}

var directions []Point = []Point{
	{y: -1, x: 0}, // N
	{y: 0, x: 1},  // E
	{y: 0, x: -1}, // W
	{y: 1, x: 0},  // S
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

	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			var p Point = Point{y: y, x: x}
			if grid.lowPoint(p) {
				result += 1 + grid.get(p)
			}
		}
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

	// luckily (?), the three largest basins correspond
	// to the lowest points already found

	var lowPoints []Point
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			var p Point = Point{y: y, x: x}
			if grid.lowPoint(p) {
				lowPoints = append(lowPoints, p)
			}
		}
	}
	// 237 low points in input data set
	// fmt.Println(len(lowPoints), "low points found")

	var basinSizes []int

	for _, start := range lowPoints {
		var size int = 1 // includes start point
		var seen map[Point]struct{} = map[Point]struct{}{
			start: /*struct{}*/ {},
		}
		var q []Point = []Point{start}
		for len(q) > 0 {
			var p Point
			p, q = q[0], q[1:]
			for _, dir := range directions {
				var np Point = p.add(dir)
				if !grid.ingrid(np) {
					continue
				}
				if _, ok := seen[np]; !ok {
					seen[np] = struct{}{}
					if grid.get(np) != 9 {
						q = append(q, np)
						size += 1
					}
				}
			}
		}
		basinSizes = append(basinSizes, size)
	}

	if len(basinSizes) > 2 { // will be same size as lowPoints
		slices.SortFunc[[]int](basinSizes, func(a, b int) int {
			return b - a // reverse (decreasing) sort
		})
		// sort.Slice(basinSizes, func(a, b int) bool {
		// 	return basinSizes[a] > basinSizes[b]
		// })
		result = basinSizes[0] * basinSizes[1] * basinSizes[2]
	} else {
		fmt.Println("not enough low points/basin sizes")
	}
	report(expected, result)
	return result
}

func main() {
	defer duration(time.Now(), "main")

	// part1(test1, 15)
	part1(input, 545)
	// part2(test1, 1134)
	part2(input, 950600)

	{
		var memStats runtime.MemStats
		runtime.ReadMemStats(&memStats)
		fmt.Printf("Heap memory (in bytes): %d\n", memStats.HeapAlloc)
		fmt.Printf("Number of garbage collections: %d\n", memStats.NumGC)
	}
}

/*
$ go run main.go
RIGHT ANSWER: 545
part 1 400.121µs
RIGHT ANSWER: 950600
part 2 2.635952ms
Heap memory (in bytes): 2059256
Number of garbage collections: 0
main 3.155417ms
*/
