package main

import (
	_ "embed"
	"flag"
	"fmt"
)

//go:embed "input.txt"
var input string

type mapKey [2]int

var sx, sy int
var rx, ry int
var presents = map[mapKey]int{mapKey{0, 0}: 1} // explicit 0,0 for legibility

func visit(r rune, x, y int) (int, int) {
	switch r {
	case '<':
		x--
	case '>':
		x++
	case '^':
		y--
	case 'v':
		y++
	default:
		fmt.Println("unexpected char", r)
	}
	presents[mapKey{x, y}]++ // auto creates new entry
	return x, y
}

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "1 or 2")
	flag.Parse()

	if part == 1 {
		for _, r := range input {
			sx, sy = visit(r, sx, sy)
		}
	} else if part == 2 {
		for i, r := range input {
			if i%2 == 0 {
				sx, sy = visit(r, sx, sy)
			} else {
				rx, ry = visit(r, rx, ry)
			}
		}
	}
	fmt.Println("part", part, len(presents)) // 2081, 2341
}
