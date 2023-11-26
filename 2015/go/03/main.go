// https://adventofcode.com/2015/day/3

package main

import (
	_ "embed"
	"log"
)

//go:embed input.txt
var input string

type mapKey [2]int

func visit(presents map[mapKey]int, r rune, x, y int) (int, int) {
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
		log.Println("unexpected rune", r)
	}
	presents[mapKey{x, y}]++ // auto creates new entry
	return x, y
}

func part1() int {
	var sx, sy int
	var presents = map[mapKey]int{{0, 0}: 1} // explicit 0,0 for legibility
	for _, r := range input {
		sx, sy = visit(presents, r, sx, sy)
	}
	return len(presents)
}

func part2() int {
	var sx, sy int
	var rx, ry int
	var presents = map[mapKey]int{{0, 0}: 1} // explicit 0,0 for legibility
	for i, r := range input {
		if i%2 == 0 {
			sx, sy = visit(presents, r, sx, sy)
		} else {
			rx, ry = visit(presents, r, rx, ry)
		}
	}
	return len(presents)
}

func main() {
	log.Println("part 1", part1()) // 2081
	log.Println("part 2", part2()) // 2341
}
