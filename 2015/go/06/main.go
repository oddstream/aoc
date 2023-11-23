package main

import (
	"bufio"
	_ "embed"
	"flag"
	"fmt"
	"strings"
)

//go:embed "input.txt"
var input string

func part1() int {
	var lights [1000][1000]bool

	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		var x0, y0, x1, y1 int
		var err error
		switch {
		case strings.HasPrefix(scanner.Text(), "turn off "):
			_, err = fmt.Sscanf(scanner.Text(), "turn off %d,%d through %d,%d", &x0, &y0, &x1, &y1)
			for x := x0; x <= x1; x++ {
				for y := y0; y <= y1; y++ {
					lights[x][y] = false
				}
			}
		case strings.HasPrefix(scanner.Text(), "turn on "):
			_, err = fmt.Sscanf(scanner.Text(), "turn on %d,%d through %d,%d", &x0, &y0, &x1, &y1)
			for x := x0; x <= x1; x++ {
				for y := y0; y <= y1; y++ {
					lights[x][y] = true
				}
			}
		case strings.HasPrefix(scanner.Text(), "toggle "):
			_, err = fmt.Sscanf(scanner.Text(), "toggle %d,%d through %d,%d", &x0, &y0, &x1, &y1)
			for x := x0; x <= x1; x++ {
				for y := y0; y <= y1; y++ {
					lights[x][y] = !lights[x][y]
				}
			}
		}
		if err != nil {
			panic(err)
		}
		// fmt.Println(x0, y0, x1, y1)
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("scanner error: %v\n", err)
	}

	var count int
	for x := 0; x < 1000; x++ {
		for y := 0; y < 1000; y++ {
			if lights[x][y] {
				count++
			}
		}
	}
	return count // 400410
}

func part2() int {
	var lights [1000][1000]int

	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		var x0, y0, x1, y1 int
		var err error
		switch {
		case strings.HasPrefix(scanner.Text(), "turn off "):
			_, err = fmt.Sscanf(scanner.Text(), "turn off %d,%d through %d,%d", &x0, &y0, &x1, &y1)
			for x := x0; x <= x1; x++ {
				for y := y0; y <= y1; y++ {
					if lights[x][y] > 0 {
						lights[x][y] -= 1
					}
				}
			}
		case strings.HasPrefix(scanner.Text(), "turn on "):
			_, err = fmt.Sscanf(scanner.Text(), "turn on %d,%d through %d,%d", &x0, &y0, &x1, &y1)
			for x := x0; x <= x1; x++ {
				for y := y0; y <= y1; y++ {
					lights[x][y] += 1
				}
			}
		case strings.HasPrefix(scanner.Text(), "toggle "):
			_, err = fmt.Sscanf(scanner.Text(), "toggle %d,%d through %d,%d", &x0, &y0, &x1, &y1)
			for x := x0; x <= x1; x++ {
				for y := y0; y <= y1; y++ {
					lights[x][y] += 2
				}
			}
		}
		if err != nil {
			panic(err)
		}
		// fmt.Println(x0, y0, x1, y1)
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("scanner error: %v\n", err)
	}

	var count int
	for x := 0; x < 1000; x++ {
		for y := 0; y < 1000; y++ {
			count += lights[x][y]
		}
	}
	return count // 15343601
}

func main() {
	var part int
	flag.IntVar(&part, "part", 2, "1 or 2")
	flag.Parse()

	var count int
	if part == 1 {
		count = part1()
	} else {
		count = part2()
	}
	fmt.Println(count) // 400410
}
