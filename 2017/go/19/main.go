// https://adventofcode.com/2017/day/19 A Series of Tubes
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

// apart from start/end points, grid is surrounded by '.'

type Point struct {
	y, x int
}

func (p Point) add(q Point) Point {
	return Point{p.y + q.y, p.x + q.x}
}

type Grid [][]byte

func (g Grid) get(p Point) byte {
	return g[p.y][p.x]
}

func (g Grid) set(p Point, ch byte) {
	g[p.y][p.x] = ch
}

func (g Grid) width() int {
	return len(g[0])
}

func (g Grid) height() int {
	return len(g)
}

func (g Grid) display() {
	for y := 0; y < g.height(); y++ {
		for x := 0; x < g.width(); x++ {
			fmt.Print(string(g.get(Point{y, x})))
		}
		fmt.Println()
	}
}

var dirmap map[string]Point = map[string]Point{
	"n": {-1, 0},
	"e": {0, 1},
	"s": {1, 0},
	"w": {0, -1},
}

var altdirmap map[string][]string = map[string][]string{
	"n": {"e", "w"},
	"e": {"n", "s"},
	"s": {"e", "w"},
	"w": {"n", "s"},
}

func isletter(ch byte) bool {
	return ch >= 'A' && ch <= 'Z'
}

func duration(invocation time.Time, name string) {
	fmt.Println(name, time.Since(invocation))
}

// load grid and return grid and starting point
func load() (Grid, Point) {
	var grid Grid
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		var row []byte = []byte(scanner.Text())
		grid = append(grid, row)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	// "Starting at the only line touching the top of the diagram ..."
	for x := 0; x < grid.width(); x++ {
		if grid.get(Point{0, x}) == '|' {
			return grid, Point{0, x}
		}
	}
	return Grid{}, Point{-1, -1}
}

func walk() (string, int) {
	grid, pos := load()
	var letters []byte = make([]byte, 0, 26)
	var steps int
	var direction string = "s"
	for {
		ch := grid.get(pos)

		if ch == ' ' {
			break
		} else if isletter(ch) {
			// the letters don't change the direction
			letters = append(letters, ch)
		} else if ch == '|' || ch == '-' {
		} else if ch == '+' {
			// look left and right
			for _, dir := range altdirmap[direction] {
				np := pos.add(dirmap[dir])
				nch := grid.get(np)
				if nch == '|' || nch == '-' || isletter(nch) {
					direction = dir
					break
				}
			}
		}
		pos = pos.add(dirmap[direction])
		steps += 1
	}
	return string(letters), steps
}

func main() {
	defer duration(time.Now(), "main")

	fmt.Println(walk()) // DTOUFARJQ, 16642
}

/*
$ go run main.go
DTOUFARJQ 16642
main 325.517µs
*/
