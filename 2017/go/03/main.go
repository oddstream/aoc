// https://adventofcode.com/2017/day/3 Spiral Memory
package main

import (
	"fmt"
	"time"
)

const input int = 277678

type Pair [2]int
type Grid map[Pair]int

func duration(invocation time.Time, name string) {
	fmt.Println(name, time.Since(invocation))
}

func (grid Grid) display(n int) {
	for y := -n; y <= n; y++ {
		for x := -n; x <= n; x++ {
			fmt.Print(grid[[2]int{x, y}], "\t")
		}
		fmt.Println()
	}
}

/*
	part 1 - no code required
	each SE step in the spiral is an increasing series of
	odd-number * odd-number, viz 1, 3, 5, 7, 9 = 1, 9, 25, 49, 81 ...
	the manhatten distance from that number (eg 25) is sqrt(num)-1
	eg sqrt(25) = 5, -1 = 4
	input is 277678, sqrt(277678) = 526.95
	ceil(sqrt(277678)) = 527
	the next biggest 'corner' is 527x527 = 277729, which is 51 over input
	manhatten distance from that corner is 527-1 = 526
	so steps from input to '1' is 526-51, = 475
*/

func partOne() int {
	var grid Grid = make(Grid)
	grid[Pair{0, 0}] = 1

	var d int = 1
	var val = 1
	var x, y = 0, 0
	for val <= 25 {
		// right
		for i := 0; i < d; i++ {
			x = x + 1
			val += 1
			grid[Pair{x, y}] = val
		}
		// up
		for i := 0; i < d; i++ {
			y = y - 1
			val += 1
			grid[Pair{x, y}] = val
		}

		d += 1

		// left
		for i := 0; i < d; i++ {
			x = x - 1
			val += 1
			grid[Pair{x, y}] = val
		}

		// down
		for i := 0; i < d; i++ {
			y = y + 1
			val += 1
			grid[Pair{x, y}] = val
		}

		d += 1
	}
	grid.display(3)
	return 475
}

var dirs []Pair = []Pair{
	{0, -1},  // N
	{1, -1},  // NE
	{1, 0},   // E
	{1, 1},   // SE
	{0, 1},   // S
	{-1, 1},  // SW
	{-1, 0},  // W
	{-1, -1}, // NW
}

func partTwo() int {
	var grid Grid = make(Grid)
	grid[Pair{0, 0}] = 1

	var adjacent = func(x, y int) int {
		var n int
		for _, pair := range dirs {
			var q = Pair{x + pair[0], y + pair[1]}
			n += grid[q]
		}
		return n
	}

	var d int = 1
	var x, y = 0, 0
	for {
		// right
		for i := 0; i < d; i++ {
			x = x + 1
			n := adjacent(x, y)
			if n > input {
				return n
			}
			grid[Pair{x, y}] = n
		}
		// up
		for i := 0; i < d; i++ {
			y = y - 1
			n := adjacent(x, y)
			if n > input {
				return n
			}
			grid[Pair{x, y}] = n
		}

		d += 1

		// left
		for i := 0; i < d; i++ {
			x = x - 1
			n := adjacent(x, y)
			if n > input {
				return n
			}
			grid[Pair{x, y}] = n
		}

		// down
		for i := 0; i < d; i++ {
			y = y + 1
			n := adjacent(x, y)
			if n > input {
				return n
			}
			grid[Pair{x, y}] = n
		}

		d += 1
	}
}

func main() {
	defer duration(time.Now(), "main")

	// fmt.Println(partOne()) // 475
	fmt.Println(partTwo()) // 279138
}

/*
$ go run main.go
475
279138
main 61.289µs
*/
