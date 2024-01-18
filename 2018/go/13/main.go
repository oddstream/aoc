// https://adventofcode.com/2018/day/13 Mine Cart Madness
package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"sort"
	"strings"
	"time"
)

//go:embed input.txt
var input string

type Point struct {
	y, x int
}

func (p Point) add(q Point) Point {
	return Point{p.y + q.y, p.x + q.x}
}

type Cart struct {
	Point
	dir     string
	next    int // 0: left, 1: straight, 2: right (repeat)
	removed bool
}

var dirMap map[string]Point = map[string]Point{
	"^": {y: -1, x: 0},
	"v": {y: 1, x: 0},
	"<": {y: 0, x: -1},
	">": {y: 0, x: 1},
}

var leftTurnMap map[string]string = map[string]string{
	"^": "<",
	"v": ">",
	"<": "v",
	">": "^",
}

var straightTurnMap map[string]string = map[string]string{
	"^": "^",
	"v": "v",
	"<": "<",
	">": ">",
}

var rightTurnMap map[string]string = map[string]string{
	"^": ">",
	"v": "<",
	"<": "^",
	">": "v",
}

var turnMaps []map[string]string = []map[string]string{leftTurnMap, straightTurnMap, rightTurnMap}

var slashCornerMap map[string]string = map[string]string{
	"^": ">",
	"v": "<",
	"<": "v",
	">": "^",
}

var backslashCornerMap map[string]string = map[string]string{
	"^": "<",
	"v": ">",
	"<": "^",
	">": "v",
}

func duration(invocation time.Time, name string) {
	fmt.Println(name, time.Since(invocation))
}

func display(tracks [][]string) {
	for y := 0; y < len(tracks); y++ {
		for x := 0; x < len(tracks[0]); x++ {
			fmt.Print(tracks[y][x])
		}
		fmt.Println()
	}
}

func loadInput() [][]string {
	var tracks [][]string
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		tracks = append(tracks, strings.Split(scanner.Text(), ""))
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	return tracks
}

func findCarts(tracks [][]string) []*Cart {
	var carts []*Cart
	for y := 0; y < len(tracks); y++ {
		for x := 0; x < len(tracks[0]); x++ {
			ch := tracks[y][x]
			if ch == "^" || ch == "v" || ch == ">" || ch == "<" {
				carts = append(carts, &Cart{Point: Point{y: y, x: x}, dir: ch})
				if ch == "<" || ch == ">" {
					tracks[y][x] = "-"
				} else if ch == "^" || ch == "v" {
					tracks[y][x] = "|"
				}
			}
		}
	}
	return carts
}

func tick(tracks [][]string, carts []*Cart, remove bool) (Point, bool) {
	sort.Slice(carts, func(a, b int) bool {
		return carts[a].y*len(carts)+carts[a].x < carts[b].y*len(carts)+carts[b].x
	})

	for i := 0; i < len(carts); i++ {
		c := carts[i]
		if c.removed {
			continue
		}
		c.Point = c.Point.add(dirMap[c.dir])
		switch tracks[c.y][c.x] {
		case "+":
			turnMap := turnMaps[c.next]
			c.dir = turnMap[c.dir]
			c.next += 1
			if c.next == 3 { // 0: left, 1: straight, 2: right (repeat)
				c.next = 0
			}
		case "|", "-":
		case " ":
			panic("ran into space!")
		case "/":
			c.dir = slashCornerMap[c.dir]
		case "\\":
			c.dir = backslashCornerMap[c.dir]
		default:
			panic("unknown thing in map")
		}
		if remove {
			if pt, ok := removeCrashedCarts(carts); ok {
				return pt, true
			}
		} else {
			if pt, ok := crash(carts); ok {
				return pt, true
			}
		}
	}
	return Point{-1, -1}, false
}

func crash(carts []*Cart) (Point, bool) {
	var m map[Point]struct{} = make(map[Point]struct{})
	for _, c := range carts {
		pt := Point{y: c.y, x: c.x}
		if _, ok := m[pt]; ok {
			return pt, true
		}
		m[pt] = struct{}{}
	}
	return Point{-1, -1}, false
}

func removeCrashedCarts(carts []*Cart) (Point, bool) {
	var m map[Point]*Cart = make(map[Point]*Cart)
	for _, c := range carts {
		if c.removed {
			continue
		}
		pt := Point{y: c.y, x: c.x}
		if oldc, ok := m[pt]; ok {
			oldc.removed = true
			c.removed = true
		}
		m[pt] = c
	}
	var ncarts int
	for _, c := range carts {
		if !c.removed {
			ncarts++
		}
	}
	if ncarts == 1 {
		for _, c := range carts {
			if !c.removed {
				return Point{y: c.y, x: c.x}, true
			}
		}
	}
	return Point{-1, -1}, false
}

func simulate(remove bool) string {
	tracks := loadInput()
	carts := findCarts(tracks)

	for i := 0; i < 18000; i++ {
		if pt, ok := tick(tracks, carts, remove); ok {
			fmt.Println("crash after tick", i)
			return fmt.Sprintf("%d,%d", pt.x, pt.y)
		}
	}

	return "-"
}

func main() {
	defer duration(time.Now(), "main")

	fmt.Println(simulate(false)) // 116,91 after tick 530
	fmt.Println(simulate(true))  // 8,23 after tick 17316
}

/*
$ go run main.go
*/
