// https://adventofcode.com/2016/day/8
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

const (
	// text.txt
	// WIDTH int = 7
	// HEIGHT int = 3

	// input.txt
	WIDTH  int = 50
	HEIGHT int = 6
)

func duration(invocation time.Time, name string) {
	fmt.Println(name, time.Since(invocation))
}

func partOne(width, height int) int {
	var screen [][]bool // screen[y rows height][x columns width]

	var display func(string) = func(line string) {
		fmt.Println(line)
		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				if screen[y][x] {
					fmt.Print("#")
				} else {
					fmt.Print(".")
				}
			}
			fmt.Println()
		}
		fmt.Println()
	}

	for y := 0; y < height; y++ {
		screen = append(screen, make([]bool, width))
	}
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "rect ") {
			var x, y int
			if _, err := fmt.Sscanf(line, "rect %dx%d", &x, &y); err != nil {
				fmt.Println(line, err)
				break
			}
			for i := 0; i < y; i++ {
				for j := 0; j < x; j++ {
					screen[i][j] = true
				}
			}
		} else if strings.HasPrefix(line, "rotate row ") {
			var y, n int
			if _, err := fmt.Sscanf(line, "rotate row y=%d by %d", &y, &n); err != nil {
				fmt.Println(line, err)
				break
			}
			for ; n > 0; n-- {
				tmp := screen[y][width-1]
				for i := width - 1; i > 0; i-- {
					screen[y][i] = screen[y][i-1]
				}
				screen[y][0] = tmp
			}
		} else if strings.HasPrefix(line, "rotate column ") {
			var x, n int
			if _, err := fmt.Sscanf(line, "rotate column x=%d by %d", &x, &n); err != nil {
				fmt.Println(line, err)
				break
			}
			for ; n > 0; n-- {
				tmp := screen[height-1][x]
				for i := height - 1; i > 0; i-- {
					screen[i][x] = screen[i-1][x]
				}
				screen[0][x] = tmp
			}
		}
		display(line)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	var count int
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if screen[y][x] {
				count += 1
			}
		}
	}
	return count
}

func main() {
	defer duration(time.Now(), "main")

	fmt.Println(partOne(WIDTH, HEIGHT)) // 123, AFBUPZBJPS
}
