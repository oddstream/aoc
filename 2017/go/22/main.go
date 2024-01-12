// https://adventofcode.com/2017/day/22 Sporifica Virus
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

const (
	Clean int = iota
	Weakened
	Infected
	Flagged
)

type Point struct {
	y, x int
}

func (p Point) add(q Point) Point {
	return Point{p.y + q.y, p.x + q.x}
}

var directions map[string]Point = map[string]Point{
	"u": {-1, 0},
	"d": {1, 0},
	"l": {0, -1},
	"r": {0, 1},
}

var left map[string]string = map[string]string{
	"u": "l",
	"d": "r",
	"l": "d",
	"r": "u",
}

var right map[string]string = map[string]string{
	"u": "r",
	"d": "l",
	"l": "u",
	"r": "d",
}

var reverse map[string]string = map[string]string{
	"u": "d",
	"d": "u",
	"l": "r",
	"r": "l",
}

type Grid map[Point]int

func duration(invocation time.Time, name string) {
	fmt.Println(name, time.Since(invocation))
}

func loadInput() (Grid, Point) {
	defer duration(time.Now(), "input")

	var g Grid = Grid{}
	var y int
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		for x, ch := range strings.Split(scanner.Text(), "") {
			if ch == "#" {
				g[Point{y, x}] = Infected
			}
		}
		y += 1
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	return g, Point{y / 2, y / 2} // input is 25x25
}

func partOne() int {
	defer duration(time.Now(), "part 1")

	infected, pos := loadInput()
	var dir string = "u"
	var result int

	burst := func() {
		if _, ok := infected[pos]; ok {
			dir = right[dir]
			delete(infected, pos)
		} else {
			dir = left[dir]
			infected[pos] = Infected
			result += 1
		}
		pos = pos.add(directions[dir])
	}

	for i := 0; i < 10000; i++ {
		burst()
	}

	return result
}

func partTwo() int {
	defer duration(time.Now(), "part 2")

	grid, pos := loadInput()
	var dir string = "u"
	var result int

	burst := func() {
		var state int
		var ok bool
		if state, ok = grid[pos]; !ok {
			state = Clean
		}
		switch state {
		case Clean:
			state = Weakened
			dir = left[dir]
		case Weakened:
			state = Infected
			// dir stays the same
			result += 1
		case Infected:
			state = Flagged
			dir = right[dir]
		case Flagged:
			state = Clean
			dir = reverse[dir]
		}
		if state == Clean {
			delete(grid, pos)
		} else {
			grid[pos] = state
		}
		pos = pos.add(directions[dir])
	}

	for i := 0; i < 10000000; i++ {
		burst()
	}

	return result
}

func main() {
	defer duration(time.Now(), "main")

	fmt.Println(partOne()) // 5411
	fmt.Println(partTwo()) // 2511416
}

/*
$ go run main.go
input 70.4µs
part 1 958.986µs
5411
input 51.391µs
part 2 757.127673ms
2511416
main 758.10927ms
*/
