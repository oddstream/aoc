// https://adventofcode.com/2021/day/15
package main

import (
	"bufio"
	"container/heap"
	_ "embed"
	"fmt"
	"runtime"
	"sort"
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

// https://tutorialhorizon.com/algorithms/dynamic-programming-minimum-cost-path-problem/
// O(n*n)
// works for test1, but not for input,
// (can't immediately see that it ought to work at all)
// I think the important bit from the min cost path algorithm is:
// "You can move only right or down
// which is coincidentally true for the smaller example puzzle,
// but not necessarily for the full one"
func minimumSumPath(grid Grid) (result int) {
	for y, rows := range grid {
		for x := range rows {
			if y == 0 && x == 0 {
				// starting square - do nothing
				continue
			} else if y == 0 {
				// first row
				grid[y][x] += grid[y][x-1]
			} else if x == 0 {
				// first column
				grid[y][x] += grid[y-1][x]
			} else {
				// path will be either down or right, choose the cheapest
				if grid[y-1][x] < grid[y][x-1] {
					grid[y][x] += grid[y-1][x]
				} else {
					grid[y][x] += grid[y][x-1]
				}
			}
		}
	}
	result = grid[len(grid)-1][len(grid[0])-1] - grid[0][0]

	for y, rows := range grid {
		for x := range rows {
			fmt.Printf("%4d", grid[y][x])
		}
		fmt.Println()
	}
	return
}

// correct but slow
func bfsWithSort(grid Grid) int {
	type BfsNode struct {
		Point
		risk int
	}
	var start Point = Point{0, 0}
	var end Point = Point{y: len(grid) - 1, x: len(grid[0]) - 1}
	var seen map[Point]struct{} = make(map[Point]struct{})
	var q []BfsNode = []BfsNode{BfsNode{Point: start, risk: 0}}
	for len(q) > 0 {
		var bn BfsNode
		bn, q = q[0], q[1:]
		if bn.Point == end {
			return bn.risk
		}
		if _, ok := seen[bn.Point]; ok {
			continue
		}
		seen[bn.Point] = struct{}{}

		for _, dir := range directions {
			var np Point = bn.Point.add(dir)
			if !grid.ingrid(np) {
				continue
			}
			q = append(q, BfsNode{Point: np, risk: bn.risk + grid.get(np)})
			sort.Slice(q, func(a, b int) bool {
				return q[a].risk < q[b].risk
			})
		}
	}
	return -1
}

// correct and 50x times faster than version with sort
// i'm thinking dijkstra == bfs with priority queue
func bfsWithPQ(grid Grid) int {
	var start Point = Point{0, 0}
	var end Point = Point{y: len(grid) - 1, x: len(grid[0]) - 1}
	var seen map[Point]struct{} = make(map[Point]struct{})
	var q PriorityQueue = PriorityQueue{&Item{Point: start, risk: 0, index: 0}}
	heap.Init(&q)

	for len(q) > 0 {
		var item = heap.Pop(&q).(*Item)
		if item.Point == end {
			return item.risk
		}
		if _, ok := seen[item.Point]; ok {
			continue
		}
		seen[item.Point] = struct{}{}

		for _, dir := range directions {
			var np Point = item.Point.add(dir)
			if !grid.ingrid(np) {
				continue
			}
			heap.Push(&q, &Item{Point: np, risk: item.risk + grid.get(np)})
		}
	}
	return -1
}

// a*

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

	result = bfsWithPQ(grid)

	report(expected, result)
	return result
}

// thank you alexchao26
func expandGrid(grid Grid) Grid {
	bigGrid := make([][]int, len(grid)*5)
	for i := range bigGrid {
		bigGrid[i] = make([]int, len(grid[0])*5)
	}
	for r, row := range grid {
		copy(bigGrid[r], row)
	}

	assignGrid := func(baseGrid [][]int, newGrid [][]int, r, c int) {
		for i := 0; i < len(newGrid); i++ {
			for j := 0; j < len(newGrid[0]); j++ {
				baseGrid[r+i][c+j] = newGrid[i][j]
			}
		}
	}

	incrementGrid := func(baseGrid [][]int, by int) [][]int {
		newGrid := make([][]int, len(baseGrid))
		for i := range newGrid {
			newGrid[i] = make([]int, len(baseGrid[0]))
		}
		for i := range baseGrid {
			for j := range baseGrid[0] {
				newGrid[i][j] = baseGrid[i][j] + by
				for newGrid[i][j] > 9 {
					newGrid[i][j] -= 9
				}
			}
		}
		return newGrid
	}

	for r := 0; r < 5; r++ {
		for c := 0; c < 5; c++ {
			if r == 0 && c == 0 {
				continue
			}
			nextGrid := incrementGrid(grid, r+c)
			assignGrid(bigGrid, nextGrid, r*len(grid), c*len(grid[0]))
		}
	}

	return bigGrid
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

	fmt.Printf("grid %dx%d", len(grid), len(grid[0]))
	grid = expandGrid(grid)
	fmt.Printf(" := %dx%d\n", len(grid), len(grid[0]))

	// result = bfsWithSort(grid)
	result = bfsWithPQ(grid)

	report(expected, result)
	return result
}

func main() {
	defer duration(time.Now(), "main")

	// part1(test1, 40)
	part1(input, 456)
	// part2(test1, 315)
	part2(input, 2831)

	{
		var memStats runtime.MemStats
		runtime.ReadMemStats(&memStats)
		fmt.Printf("Heap memory (in bytes): %d\n", memStats.HeapAlloc)
		fmt.Printf("Number of garbage collections: %d\n", memStats.NumGC)
	}
}

/*
$ go run main.go pq.go
RIGHT ANSWER: 456
part 1 8.391646ms
grid 100x100 := 500x500
RIGHT ANSWER: 2831
part 2 229.751358ms
Heap memory (in bytes): 24662160
Number of garbage collections: 10
main 238.175862ms
*/
