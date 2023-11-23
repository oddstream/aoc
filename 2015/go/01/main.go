package main

import (
	_ "embed"
	"flag"
	"fmt"
)

//go:embed "input.txt"
var input string

func main() {
	var part int
	flag.IntVar(&part, "part", 2, "1 or 2")
	flag.Parse()

	var level int
	for i, r := range input {
		if r == '(' {
			level++
		} else if r == ')' {
			level--
		} else {
			fmt.Println("unexpected char", r)
		}
		if part == 2 && level < 0 {
			fmt.Println("part 2 ", i)
			break
		}
	}
	if part == 1 {
		fmt.Println("part 1", level)
	}
}
