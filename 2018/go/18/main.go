// https://adventofcode.com/2018/day/18
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
var test1 string // 10x10

//go:embed input.txt
var input string // 50x50

type (
	Point struct {
		y, x int
	}
	Grid [][]string
)

func (p Point) add(q Point) Point {
	return Point{y: p.y + q.y, x: p.x + q.x}
}

func (g Grid) ingrid(p Point) bool {
	return p.x >= 0 && p.y >= 0 && p.x < len(g[0]) && p.y < len(g)
}

func (g Grid) get(p Point) string {
	return g[p.y][p.x]
}

func (g Grid) set(p Point, value string) {
	g[p.y][p.x] = value
}

func (g Grid) adjacent(p Point, thing string) int {
	var count int
	for _, dir := range directions {
		var q Point = p.add(dir)
		if g.ingrid(q) {
			if g.get(q) == thing {
				count += 1
			}
		}
	}
	return count
}

func (g Grid) count(thing string) int {
	var result int
	for y := 0; y < len(g); y++ {
		for x := 0; x < len(g[0]); x++ {
			if g[y][x] == thing {
				result += 1
			}
		}
	}
	return result
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

func (g Grid) display() {
	for y := 0; y < len(g); y++ {
		for x := 0; x < len(g[0]); x++ {
			fmt.Print(g.get(Point{y: y, x: x}))
		}
		fmt.Println()
	}
}

func (g Grid) stringify() string {
	var sb strings.Builder = strings.Builder{}
	for y := 0; y < len(g); y++ {
		for x := 0; x < len(g[0]); x++ {
			sb.WriteString(g[y][x])
		}
	}
	return sb.String()
}

func duration(invocation time.Time, name string) {
	fmt.Println(name, "duration", time.Since(invocation))
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

func step(grid Grid) Grid {
	var newGrid Grid
	for i := 0; i < len(grid); i++ {
		newGrid = append(newGrid, make([]string, len(grid[i])))
	}
	grid.foreach(func(p Point) {
		if grid.get(p) == "." && grid.adjacent(p, "|") >= 3 {
			newGrid.set(p, "|")
		} else if grid.get(p) == "|" && grid.adjacent(p, "#") >= 3 {
			newGrid.set(p, "#")
		} else if grid.get(p) == "#" {
			if grid.adjacent(p, "#") >= 1 && grid.adjacent(p, "|") >= 1 {
				newGrid.set(p, "#")
			} else {
				newGrid.set(p, ".")
			}
		} else {
			var thing string = grid.get(p)
			newGrid.set(p, thing)
		}
	})
	return newGrid
}

func part1(in string, expected int) (result int) {
	defer duration(time.Now(), "part 1")

	var grid Grid
	scanner := bufio.NewScanner(strings.NewReader(in))
	for scanner.Scan() {
		grid = append(grid, strings.Split(scanner.Text(), ""))
	}

	for i := 0; i < 10; i++ {
		grid = step(grid)
	}
	// grid.display()
	result = grid.count("|") * grid.count("#")
	report(expected, result)
	return result
}

func part2(in string, expected int) (result int) {
	defer duration(time.Now(), "part 2")

	var grid Grid
	scanner := bufio.NewScanner(strings.NewReader(in))
	for scanner.Scan() {
		grid = append(grid, strings.Split(scanner.Text(), ""))
	}

	var cycles map[string]int = make(map[string]int)
	for i := 0; i < 1_000_000_000; i++ {
		grid = step(grid)
		// using a strings key for this map,
		// instead of just using trees*lumberyards as an int key
		// because the latter gave a wrong answer (162610)
		var key string = grid.stringify()
		if previ, ok := cycles[key]; ok {
			// got confused by all the % 1_000_000_000 stuff
			// and was possibly off-by-one somewhere
			// so let's just bump the counter
			freq := i - previ
			for i+freq < 1_000_000_000 {
				i += freq
			}
		}
		cycles[key] = i
	}
	result = grid.count("|") * grid.count("#")

	report(expected, result)
	return result
}

func main() {
	defer duration(time.Now(), "main")

	// part1(test1, 1147)
	part1(input, 355918)
	part2(input, 202806)

	{
		var memStats runtime.MemStats
		runtime.ReadMemStats(&memStats)
		fmt.Printf("Heap memory (in bytes): %d\n", memStats.HeapAlloc)
		fmt.Printf("Number of garbage collections: %d\n", memStats.NumGC)
	}
}

/*
$ go run main.go
RIGHT ANSWER: 355918
part 1 duration 2.015829ms
RIGHT ANSWER: 202806
part 2 duration 78.971728ms
Heap memory (in bytes): 2422312
Number of garbage collections: 11
main duration 81.013447ms
*/
