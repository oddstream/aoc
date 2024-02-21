// https://adventofcode.com/2018/day/17
package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

//go:embed test1.txt
var test1 string

//go:embed input.txt
var input string

var clay map[Point]struct{} = make(map[Point]struct{})
var settled map[Point]struct{} = make(map[Point]struct{})
var flowing map[Point]struct{} = make(map[Point]struct{})

var spring Point = Point{x: 500, y: 0}
var ymin, ymax = math.MaxInt64, 0

type Point struct {
	x, y int
}

func (p Point) add(q Point) Point {
	return Point{y: p.y + q.y, x: p.x + q.x}
}

func duration(invocation time.Time, name string) {
	fmt.Println(name, "duration", time.Since(invocation))
}

// or we could use K&R p61
func atoi(s string) int {
	if n, err := strconv.Atoi(s); err == nil {
		return n
	} else {
		fmt.Println(err)
	}
	return 0
}

func token2range(s string) []int {
	lr := strings.Split(s, "..")
	if len(lr) == 1 {
		return []int{atoi(s)}
	} else {
		var a, b int = atoi(lr[0]), atoi(lr[1])
		var out []int = make([]int, 0, b-a+1)
		for i := a; i <= b; i++ {
			out = append(out, i)
		}
		return out
	}
}

func isClay(p Point) bool {
	var ok bool
	_, ok = clay[p]
	return ok
}

func isFlowing(p Point) bool {
	var ok bool
	_, ok = flowing[p]
	return ok
}

func isSettled(p Point) bool {
	var ok bool
	_, ok = settled[p]
	return ok
}

func between(a, n, b int) bool {
	return n >= a && n <= b
}

func fill(p Point, direction string) bool {
	flowing[p] = struct{}{}
	var below Point = p.add(Point{0, 1})
	var left Point = p.add(Point{-1, 0})
	var right Point = p.add(Point{1, 0})
	// real = x, imaginary = y
	// 1i, point.imag : down
	// -1 : left
	// 1 : right
	if !isClay(below) {
		if !isFlowing(below) && between(1, below.y, ymax) {
			fill(below, "down")
		}
		if !isSettled(below) {
			return false
		}
	}

	lFilled := isClay(left) || !isFlowing(left) && fill(left, "left")
	rFilled := isClay(right) || !isFlowing(right) && fill(right, "right")

	if direction == "down" && lFilled && rFilled {
		settled[p] = struct{}{}

		for isFlowing(left) {
			settled[left] = struct{}{}
			left.x -= 1
		}
		for isFlowing(right) {
			settled[right] = struct{}{}
			right.x += 1
		}
	}

	return (direction == "left" && (lFilled || isClay(left))) ||
		(direction == "right" && (rFilled || isClay(right)))
}

func loadClay(in string) {
	scanner := bufio.NewScanner(strings.NewReader(in))
	for scanner.Scan() {
		lr := strings.Split(scanner.Text(), ", ")
		var xstr, ystr string

		if strings.HasPrefix(lr[0], "x=") {
			xstr = lr[0][2:]
		} else if strings.HasPrefix(lr[1], "x=") {
			xstr = lr[1][2:]
		}
		if strings.HasPrefix(lr[0], "y=") {
			ystr = lr[0][2:]
		} else if strings.HasPrefix(lr[1], "y=") {
			ystr = lr[1][2:]
		}
		for _, x := range token2range(xstr) {
			for _, y := range token2range(ystr) {
				clay[Point{x: x, y: y}] = struct{}{}
			}
		}
	}
	for p := range clay {
		ymin = min(ymin, p.y)
		ymax = max(ymax, p.y)
	}
	fmt.Println("clay from", ymin, "to", ymax, "total", len(clay)) // 6 -> 1679, 20971
}

func display(xmin, xmax, ymax int) {
	for y := 0; y < ymax; y++ {
		for x := xmin; x <= xmax; x++ {
			if y == 0 && x == 500 {
				fmt.Print("+")
			} else if _, ok := settled[Point{x: x, y: y}]; ok {
				fmt.Print("-")
			} else if _, ok := clay[Point{x: x, y: y}]; ok {
				fmt.Print("#")
				// } else if _, ok := flowing[Point{x: x, y: y}]; ok {
				// 	fmt.Print("~")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func main() {
	defer duration(time.Now(), "main")

	loadClay(input)
	fill(spring, "down")
	// display(490, 510, 15)
	var part1, part2 int
	for p := range flowing {
		if between(ymin, p.y, ymax) {
			part1 += 1
		}
	}
	for p := range settled {
		if between(ymin, p.y, ymax) {
			part2 += 1
		}
	}
	// test1	input
	// 57		38364
	// 29		30551
	fmt.Println("part 1:", part1)
	fmt.Println("part 2:", part2)
}

/*
$ go run main.go
clay from 6 to 1679 total 20971
part 1: 38364
part 2: 30551
main duration 20.798933ms
*/
