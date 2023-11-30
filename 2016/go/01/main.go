// https://adventofcode.com/2016/day/1
package main

import (
	_ "embed"
	"fmt"
	util "oddstream/aoc/util"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string // caveat, remove trailing \n

var steps map[int][2]int = map[int][2]int{
	0: {0, -1}, // N
	1: {1, 0},  // E
	2: {0, 1},  // S
	3: {-1, 0}, // W
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func partOne() int {
	var dir int = 0 // N
	var dx, dy int
	var x, y int = 0, 0
	substrs := strings.Split(input, ", ")
	for _, substr := range substrs {
		turn := rune(substr[0])
		n, err := strconv.Atoi(substr[1:])
		if err != nil {
			fmt.Println(err)
		}
		// fmt.Println(string(turn), n)
		switch turn {
		case 'L':
			dir -= 1
			if dir < 0 {
				dir = 3
			}
			dx = steps[dir][0]
			dy = steps[dir][1]
		case 'R':
			dir += 1
			if dir > 3 {
				dir = 0
			}
			dx = steps[dir][0]
			dy = steps[dir][1]
		}
		for ; n > 0; n-- {
			x += dx
			y += dy
		}
	}
	// fmt.Println(x, y)
	return Abs(x) + Abs(y)
}

func partTwo() int {
	var dir int = 0 // N
	var dx, dy int
	var x, y int = 0, 0
	var visited util.Set[[2]int] = util.NewSet[[2]int]([2]int{x, y})
	substrs := strings.Split(input, ", ")
	for _, substr := range substrs {
		turn := rune(substr[0])
		n, err := strconv.Atoi(substr[1:])
		if err != nil {
			fmt.Println(err)
		}
		// fmt.Println(string(turn), n)
		switch turn {
		case 'L':
			dir -= 1
			if dir < 0 {
				dir = 3
			}
			dx = steps[dir][0]
			dy = steps[dir][1]
		case 'R':
			dir += 1
			if dir > 3 {
				dir = 0
			}
			dx = steps[dir][0]
			dy = steps[dir][1]
		}
		for ; n > 0; n-- {
			x += dx
			y += dy
			key := [2]int{x, y}
			if visited.Contains(key) {
				goto exit
			} else {
				visited.Add(key)
			}
		}
	}
exit:
	// fmt.Println(x, y)
	return Abs(x) + Abs(y)
}

func main() {
	if strings.HasSuffix(input, "\n") {
		fmt.Println("WARNING! input has trailing \\n")
	}
	fmt.Println(partOne()) // 291
	fmt.Println(partTwo()) // 159
}
