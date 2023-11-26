// https://adventofcode.com/2015/day/1

package main

import (
	_ "embed"
	"log"
)

//go:embed input.txt
var input string

func part1() int {
	var level int
	for _, r := range input {
		switch r {
		case '(':
			level++
		case ')':
			level--
		default:
			log.Println("unexpected rune", r)
		}
	}
	return level
}

func part2() int {
	var level int
	for i, r := range input {
		switch r {
		case '(':
			level++
		case ')':
			level--
		default:
			log.Println("unexpected rune", r)
		}
		if level == -1 {
			// add 1 because problem description mentions first
			// character is at position 1, and Go is 0-based
			return i + 1
		}
	}
	return -1
}

func main() {
	log.Println("part 1", part1())
	log.Println("part 2", part2())
}
