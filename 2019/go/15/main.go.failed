// https://adventofcode.com/2019/day/15
package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
	"time"
)

//go:embed input.txt
var input string

const (
	NORTH int = 1
	SOUTH int = 2
	WEST  int = 3
	EAST  int = 4

	WALL   int = 0
	MOVED  int = 1
	OXYGEN int = 2
)

type Point struct {
	x, y int
}

func (p Point) add(q Point) Point {
	return Point{x: p.x + q.x, y: p.y + q.y}
}

var directions map[int]Point = map[int]Point{
	NORTH: {x: 0, y: -1},
	SOUTH: {x: 0, y: 1},
	WEST:  {x: -1, y: 0},
	EAST:  {x: 1, y: 0},
}

func duration(invocation time.Time, name string) {
	fmt.Println(name, time.Since(invocation))
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

func play1(program []int) int {

	type BfsPoint struct {
		Point
		steps int
		parent *BfsPoint
	}

	var p BfsPoint = BfsPoint{{Point{x: 0, y: 0}} // one assumes start position is 0,0
	var np BfsPoint
	var section map[Point]rune = make(map[Point]rune)
	section[p.Point] = '.'

	in := func() int {
		return 0
	}
	out := func(val int) {
		switch val {
		case WALL:
			section[p.Point] = '#'
		case MOVED:
			p = np
			section[p.Point] = '.'
		case OXYGEN:
			section[p.Point] = 'O'
			fmt.Println("oxygen at", p)
		}
	}

	var q []BfsPoint = []BfsPoint{p}
	for len(q) > 0 {
		p, q = q[0], q[1:]
		for d, dir := range directions {
			np = BfsPoint{Point:p.Point.add(dir)}
			if _, ok := section[np.Point]; !ok {
				in(d)
			}
		}
	}

	return 0
}

func partOne(in string, expected int) (result int) {
	defer duration(time.Now(), "part 1")

	var tokens []string = strings.Split(strings.Trim(input, "\n"), ",")
	var masterProgram []int
	for _, tok := range tokens {
		masterProgram = append(masterProgram, atoi(tok))
	}
	// The computer's available memory should be much larger than the initial program.
	// Memory beyond the initial program starts with the value 0
	// and can be read or written like any other memory.
	var program []int = make([]int, len(masterProgram)+1000)
	copy(program, masterProgram)

	result = play1(program)

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

	partOne(input, 0)
	//	partTwo()

	// {
	// 	var memStats runtime.MemStats
	// 	runtime.ReadMemStats(&memStats)
	// 	fmt.Printf("Heap memory (in bytes): %d\n", memStats.HeapAlloc)
	// 	fmt.Printf("Number of garbage collections: %d\n", memStats.NumGC)
	// }
}

// https://github.com/xorkevin/advent2019/blob/master/day15/main.go

/*
$ go run main.go
*/
