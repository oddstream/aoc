// https://adventofcode.com/2019/day/10
package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"math"
	"sort"
	"strings"
	"time"
)

//go:embed test1.txt
var test1 string

//go:embed test2.txt
var test2 string

//go:embed test3.txt
var test3 string

//go:embed test4.txt
var test4 string

//go:embed test5.txt
var test5 string

//go:embed input.txt
var input string

type Point struct {
	x, y int
}

func duration(invocation time.Time, name string) {
	fmt.Println(name, time.Since(invocation))
}

// slope between two points
//
// Deprecated: because slope(Point{0, 0}, Point{0, 4}) == +Inf
// func slope(p1, p2 Point) float64 {
// 	return (float64(p2.y) - float64(p1.y)) / (float64(p2.x) - float64(p1.x))
// }

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

// func manhatten(p, q Point) int {
// 	return abs(q.x-p.x) + abs(q.y-p.y)
// }

// returns true if three points are in line
//
// Deprecated: because uses slope
// func isCollinear(p1, p2, p3 Point) bool {
// 	// If the slope of the line between p1 and p2 is equal to the slope of the line between p2 and p3,
// 	// then the points are collinear
// 	return slope(p1, p2) == slope(p1, p3)
// }

// return true if p3 lies between p1 and p2
func between(p1, p2, p3 Point) bool {
	// https://stackoverflow.com/questions/11907947/how-to-check-if-a-point-lies-on-a-line-between-2-other-points

	// first check whether the point lies on the line p1-p2
	// for that you simply need a "cross-product" of vectors p1 -> p3 and p1 -> p2
	var dxc = p3.x - p1.x
	var dyc = p3.y - p1.y

	var dxl = p2.x - p1.x
	var dyl = p2.y - p1.y

	var cross = dxc*dyl - dyc*dxl
	// p3 lies on the p1-p2 line if and only if cross == 0
	if cross != 0 {
		return false
	}
	// does p3 lie BETWEEN p1-p2?
	if abs(dxl) >= abs(dyl) {
		if dxl > 0 {
			return p1.x <= p3.x && p3.x <= p2.x
		} else {
			return p2.x <= p3.x && p3.x <= p1.x
		}
	} else {
		if dyl > 0 {
			return p1.y <= p3.y && p3.y <= p2.y
		} else {
			return p2.y <= p3.y && p3.y <= p1.y
		}
	}
}

func part1(in string, expected int) {
	defer duration(time.Now(), "part 1")

	var asteroids map[Point]struct{} = make(map[Point]struct{})
	scanner := bufio.NewScanner(strings.NewReader(in))
	var width, height, y int
	for scanner.Scan() {
		width = len(scanner.Text())
		for x, ch := range strings.Split(scanner.Text(), "") {
			if ch == "#" {
				asteroids[Point{y: y, x: x}] = struct{}{}
			}
		}
		y += 1
	}
	height = y
	fmt.Println(len(asteroids), "asteroids, width height is", width, height)

	obstructed := func(pt1, pt2 Point) bool {
		for pt3 := range asteroids {
			if pt3 == pt1 || pt3 == pt2 {
				continue
			}
			if between(pt1, pt2, pt3) {
				return true
			}
		}
		return false
	}

	var max_p1_sees int
	var max_p1_sees_pt Point
	for pt1 := range asteroids {
		var p1_sees int
		for pt2 := range asteroids {
			if pt2 == pt1 {
				continue
			}
			if !obstructed(pt1, pt2) {
				p1_sees += 1
			}
		}
		if p1_sees > max_p1_sees {
			max_p1_sees = p1_sees
			max_p1_sees_pt = pt1
		}
	}

	fmt.Println(max_p1_sees_pt, max_p1_sees)

	var result = max_p1_sees
	if expected != -1 {
		if result != expected {
			fmt.Println("WRONG ANSWER: got", result, "expected", expected)
		} else {
			fmt.Println("RIGHT ANSWER:", result)
		}
	}
}

// returns angle (in degrees) between origin p1 and other point p2
// adjusted by 90 degrees so N = 0, E = 90, S = 180, W = 270
func angle(p1, p2 Point) float64 {
	a := math.Atan2(float64(p2.y-p1.y), float64(p2.x-p1.x))
	a = a * 180.0 / math.Pi
	a = a + 90.0
	if a < 0 {
		// some angles come out -ve
		// don't have enough math to know exactly why
		// this fudge (wrapping them around) seems to get me by
		a = a + 360.0
	}
	return a
}

func part2(in string, laser Point, expected int) {
	defer duration(time.Now(), "part 2")

	var asteroids map[Point]struct{} = make(map[Point]struct{})
	scanner := bufio.NewScanner(strings.NewReader(in))
	var width, height, y int
	for scanner.Scan() {
		width = len(scanner.Text())
		for x, ch := range strings.Split(scanner.Text(), "") {
			if ch == "#" {
				asteroids[Point{y: y, x: x}] = struct{}{}
			}
		}
		y += 1
	}
	height = y
	fmt.Println(len(asteroids), "asteroids, width height is", width, height)

	// make a list of all asteroids visible from laser
	// sort list by angle (straight up is zero) angle = arctan2(y2-y1/x2-x1)
	// delete asteroids in list (make a note of Point of 200th destroyed)

	obstructed := func(pt1, pt2 Point) bool {
		for pt3 := range asteroids {
			if pt3 == pt1 || pt3 == pt2 {
				continue
			}
			if between(pt1, pt2, pt3) {
				return true
			}
		}
		return false
	}
	var result, deleted int
	for len(asteroids) > 0 {
		var destroyable []Point
		for p2 := range asteroids {
			if !obstructed(laser, p2) {
				destroyable = append(destroyable, p2)
			}
		}
		sort.Slice(destroyable, func(a, b int) bool {
			return angle(laser, destroyable[a]) < angle(laser, destroyable[b])
		})
		for _, pt := range destroyable {
			// fmt.Println(i, pt, angle(laser, pt))
			delete(asteroids, pt)
			deleted += 1
			if deleted == 200 {
				result = pt.x*100 + pt.y
			}
		}
	}
	if expected != -1 {
		if result != expected {
			fmt.Println("WRONG ANSWER: got", result, "expected", expected)
		} else {
			fmt.Println("RIGHT ANSWER:", result)
		}
	}
}

func main() {
	defer duration(time.Now(), "main")

	// fmt.Println(between(Point{x: 0, y: 0}, Point{x: 4, y: 4}, Point{x: 1, y: 1})) // true
	// fmt.Println(between(Point{x: 0, y: 0}, Point{x: 4, y: 4}, Point{x: 2, y: 2})) // true
	// fmt.Println(between(Point{x: 0, y: 0}, Point{x: 4, y: 4}, Point{x: 3, y: 3})) // true
	// fmt.Println(between(Point{x: 2, y: 2}, Point{x: 4, y: 4}, Point{x: 0, y: 0})) // false
	// fmt.Println(between(Point{x: 0, y: 0}, Point{x: 4, y: 4}, Point{x: 4, y: 0})) // false
	// fmt.Println(between(Point{x: 2, y: 2}, Point{x: 4, y: 4}, Point{x: 2, y: 3})) // false
	// part1(test1, 8)   // 3,4 8
	// part1(test2, 33)  // 5,8 33
	// part1(test3, 35)  // 1,2 35
	// part1(test4, 41)  // 6,3 41
	// part1(test5, 210) // 11,13 210
	part1(input, 334) // 23,20 334

	// fmt.Println(angle(Point{0, 0}, Point{0, -4})) // n
	// fmt.Println(angle(Point{0, 0}, Point{4, 0}))  // e
	// fmt.Println(angle(Point{0, 0}, Point{0, 4}))  // s
	// fmt.Println(angle(Point{0, 0}, Point{-4, 0})) // w

	// part2(test1, Point{x: 3, y: 4}, -1)
	// part2(test5, Point{x: 11, y: 13}, 802)
	part2(input, Point{x: 23, y: 20}, 1119)

}

/*
$ go run main.go
417 asteroids, width height is 34 34
{23 20} 334
RIGHT: 334
part 1 930.434359ms
417 asteroids, width height is 34 34
RIGHT: 1119
part 2 2.79069ms
main 933.259658ms
*/
