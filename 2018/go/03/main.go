// https://adventofcode.com/2018/day/3 No Matter How You Slice It
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

type Point struct {
	x, y int
}

func duration(invocation time.Time, name string) {
	fmt.Println(name, time.Since(invocation))
}

func partOne() int {
	defer duration(time.Now(), "part 1")
	var fabric map[Point]int = make(map[Point]int)
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		var c, x, y, w, h int
		if n, err := fmt.Sscanf(scanner.Text(), "#%d @ %d,%d: %dx%d", &c, &x, &y, &w, &h); n != 5 {
			fmt.Println(err)
			break
		}
		for i := x; i < x+w; i++ {
			for j := y; j < y+h; j++ {
				fabric[Point{i, j}]++
			}
		}
	}
	var result int
	for _, v := range fabric {
		if v > 1 {
			result += 1
		}
	}
	return result
}

func partTwo() int {
	defer duration(time.Now(), "part 2")
	var fabric map[Point]int = make(map[Point]int)
	// pass 1: make the claims
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		var c, x, y, w, h int
		if n, err := fmt.Sscanf(scanner.Text(), "#%d @ %d,%d: %dx%d", &c, &x, &y, &w, &h); n != 5 {
			fmt.Println(err)
			break
		}
		for i := x; i < x+w; i++ {
			for j := y; j < y+h; j++ {
				fabric[Point{i, j}]++
			}
		}
	}
	// pass 2: see if any claim is all 1
	scanner = bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		var c, x, y, w, h int
		if n, err := fmt.Sscanf(scanner.Text(), "#%d @ %d,%d: %dx%d", &c, &x, &y, &w, &h); n != 5 {
			fmt.Println(err)
			break
		}
		for i := x; i < x+w; i++ {
			for j := y; j < y+h; j++ {
				if fabric[Point{i, j}] > 1 {
					goto nextline
				}
			}
		}
		return c
	nextline:
	}
	return -1
}

func main() {
	defer duration(time.Now(), "main")

	fmt.Println(partOne()) // 4, 109785
	fmt.Println(partTwo()) // 3, 504s
}

/*
$ go run main.go
part 1 59.581253ms
109785
part 2 55.922612ms
504
main 115.562719ms
*/
