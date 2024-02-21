// https://adventofcode.com/
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

func (g Grid) getExtended(p Point) string {
	var x int = p.x
	var y int = p.y
	var W int = len(g[0])
	var H = len(g)
	for ; x >= W; x -= W {
	}
	for ; x < 0; x += W {
	}
	for ; y >= H; y -= H {
	}
	for ; y < 0; y += H {
	}
	if !g.ingrid(Point{y: y, x: x}) {
		fmt.Println("not ingrid", y, x)
		return "!"
	}
	return g[y][x]
}

func (g Grid) set(p Point, value string) {
	g[p.y][p.x] = value
}

func (g Grid) foreach(fn func(Point)) {
	for y := 0; y < len(g); y++ {
		for x := 0; x < len(g[0]); x++ {
			fn(Point{y: y, x: x})
		}
	}
}

var directions []Point = []Point{
	{y: -1, x: 0}, // N
	// {y: -1, x: 1},  // NE
	{y: 0, x: 1}, // E
	// {y: 1, x: 1},   // SE
	{y: 1, x: 0}, // S
	// {y: 1, x: -1},  // SW
	{y: 0, x: -1}, // W
	// {y: -1, x: -1}, // NW
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

func (g Grid) display(elves map[Point]struct{}) {
	for y := 0; y < len(g); y++ {
		for x := 0; x < len(g[0]); x++ {
			if _, ok := elves[Point{y: y, x: x}]; ok {
				fmt.Print("O")
			} else if g.get(Point{y: y, x: x}) == "#" {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
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

func part1(in string, steps int, expected int) (result int) {
	defer duration(time.Now(), "part 1")

	var grid Grid
	scanner := bufio.NewScanner(strings.NewReader(in))
	for scanner.Scan() {
		var row []string = strings.Split(scanner.Text(), "")
		grid = append(grid, row)
	}
	var visited map[Point]struct{} = make(map[Point]struct{})
	var elves map[Point]struct{} = make(map[Point]struct{})
	grid.foreach(func(p Point) {
		if grid.get(p) == "S" {
			elves[p] = struct{}{}
			visited[p] = struct{}{}
			grid.set(p, ".")
		}
	})
	// fmt.Println(grid)
	// fmt.Println(elves)
	for step := 0; step < steps; step++ {
		var newElves map[Point]struct{} = make(map[Point]struct{})
		for p := range elves {
			for _, dir := range directions {
				var q Point = p.add(dir)
				if grid.ingrid(q) && grid.get(q) == "." {
					visited[q] = struct{}{}
					newElves[q] = struct{}{}
				}
			}
		}
		elves = newElves
	}
	// grid.display(elves)
	fmt.Println(len(visited))

	result = len(elves)
	report(expected, result)
	return result
}

func part2(in string, steps int, expected int) (result int) {
	defer duration(time.Now(), "part 2")

	var grid Grid
	scanner := bufio.NewScanner(strings.NewReader(in))
	for scanner.Scan() {
		var row []string = strings.Split(scanner.Text(), "")
		grid = append(grid, row)
	}
	var elves map[Point]struct{} = make(map[Point]struct{})
	grid.foreach(func(p Point) {
		if grid.get(p) == "S" {
			elves[p] = struct{}{}
			grid.set(p, ".")
		}
	})
	// fmt.Println(grid)
	// fmt.Println(elves)
	for step := 0; step < steps; step++ {
		var newElves map[Point]struct{} = make(map[Point]struct{})
		for p := range elves {
			for _, dir := range directions {
				var q Point = p.add(dir)
				if grid.getExtended(q) == "." {
					newElves[q] = struct{}{}
				}
			}
		}
		elves = newElves
	}

	result = len(elves)
	report(expected, result)
	return result
}

// https://github.com/villuna/aoc23/wiki/A-Geometric-solution-to-advent-of-code-2023,-day-21
// grid is 131x131 (17161)
// 65 steps from starting point to edge
// starting row and column are clear of rocks
// 2282 rocks
// 14878 plots (not including start)
// at each step, the center row grows 1 left and 1 right, the center column grows 1 up and 1 down
// (part 1) only every other square is occupied by an elf after each turn
// because steps are odd - even - odd - even numbered squares
// the shape is a diamond, 129x129
// so the number of steps = area of diamond/2 - number of rocks on places where step would have been
// area of diamond is diagonal * diagonal / 2 (64*64/2) = 4096/2 = 2048
// 4096 - 3642 = 454 (seems like a reasonable number of rocks in diamond)
// https://github.com/JoanaBLate/advent-of-code-js/blob/main/2023/day21/solve2.js

func main() {
	defer duration(time.Now(), "main")

	// part1(test1, 6, 16)

	part1(input, 64, 3642)

	// part2(test1, 6, 16)
	// part2(test1, 10, 50)
	// part2(test1, 50, 1594)
	// part2(test1, 100, 6536)
	// part2(test1, 500, 167004) // 29s
	// part2(test1, 1000, 668697)	// 2m28s
	// part2(test1, 5000, 16733044)	// 6h7m

	// part2(input, 26501365, 0)

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
