// https://adventofcode.com/2015/day/2

package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"log"
	"sort"
	"strings"
	"time"
)

//go:embed input.txt
var input string

func duration(invocation time.Time, name string) {
	fmt.Println(name, time.Since(invocation))
}

func part1() int {
	duration(time.Now(), "part1")

	var totalPaper int
	var dims [3]int // use an array
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		_, err := fmt.Sscanf(scanner.Text(), "%dx%dx%d", &dims[0], &dims[1], &dims[2])
		if err != nil {
			panic(err)
		}
		sort.Ints(dims[:]) // turn array into slice
		l := dims[0]
		w := dims[1]
		h := dims[2]
		area := 2*l*w + 2*w*h + 2*h*l
		extra := dims[0] * dims[1]
		totalPaper += area + extra
	}
	return totalPaper
}

func part2() int {
	duration(time.Now(), "part1")

	var totalRibbon int
	var dims []int = []int{0, 0, 0} // use an instantiated slice
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		_, err := fmt.Sscanf(scanner.Text(), "%dx%dx%d", &dims[0], &dims[1], &dims[2])
		if err != nil {
			panic(err)
		}
		sort.Ints(dims)
		ribbon := dims[0] + dims[0] + dims[1] + dims[1]
		bow := dims[0] * dims[1] * dims[2]
		totalRibbon += ribbon + bow
	}
	return totalRibbon
}

func main() {
	duration(time.Now(), "main")
	log.Println("part 1", part1()) // 1606483
	log.Println("part 2", part2()) // 3842356
}
