// https://adventofcode.com/2018/day/10 The Stars Align
package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"math"
	"strings"
	"time"
)

//go:embed input.txt
var input string

type Point struct {
	x, y, dx, dy int
}

func duration(invocation time.Time, name string) {
	fmt.Println(name, time.Since(invocation))
}

func step(points []Point) {
	for i := 0; i < len(points); i++ {
		points[i].x += points[i].dx
		points[i].y += points[i].dy
	}
}

func bounds(points []Point) (int, int, int, int) {
	var minx int = math.MaxInt64
	var miny int = math.MaxInt64
	var maxx int = -math.MaxInt64
	var maxy int = -math.MaxInt64
	for i := 0; i < len(points); i++ {
		if points[i].x < minx {
			minx = points[i].x
		}
		if points[i].y < miny {
			miny = points[i].y
		}
		if points[i].x > maxx {
			maxx = points[i].x
		}
		if points[i].y > maxy {
			maxy = points[i].y
		}
	}
	return minx, miny, maxx, maxy
}

func display(points []Point) {

	contains := func(x, y int) bool {
		for i := 0; i < len(points); i++ {
			if points[i].x == x && points[i].y == y {
				return true
			}
		}
		return false
	}

	minx, miny, maxx, maxy := bounds(points)

	for y := miny; y <= maxy; y++ {
		for x := minx; x <= maxx; x++ {
			if contains(x, y) {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func partOne() {
	defer duration(time.Now(), "part 1")

	var points []Point
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		var pt Point
		if n, err := fmt.Sscanf(scanner.Text(), "position=<%d, %d> velocity=<%d, %d>", &pt.x, &pt.y, &pt.dx, &pt.dy); n != 4 {
			fmt.Println(err)
			break
		}
		points = append(points, pt)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	// fmt.Println(points)

	var prevArea int = math.MaxInt64
	for i := 0; i < 10003; i++ {
		step(points)
		minx, miny, maxx, maxy := bounds(points)
		area := (maxx - minx) * (maxy - miny)
		if area > prevArea {
			fmt.Println(i, area)
			break
		}
		prevArea = area
	}
	display(points)
}

func main() {
	defer duration(time.Now(), "main")

	partOne()
}

/*
$ go run main.go
*/
