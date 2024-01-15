// https://adventofcode.com/2018/day/6 Chronal Coordinates
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
	x, y int
}

// func (p Point) add(q Point) Point {
// 	return Point{p.x + q.x, p.y + q.y}
// }

func (p Point) manhatten(q Point) int {
	return abs(q.x-p.x) + abs(q.y-p.y)
}

type Bounds struct {
	min, max Point
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func duration(invocation time.Time, name string) {
	fmt.Println(name, time.Since(invocation))
}

func loadInput() ([]Point, Bounds) {
	var points []Point
	var bounds Bounds = Bounds{min: Point{x: math.MaxInt64, y: math.MaxInt64}, max: Point{x: -math.MaxInt64, y: -math.MaxInt64}}
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		var y, x int
		if n, err := fmt.Sscanf(scanner.Text(), "%d, %d", &x, &y); n != 2 {
			fmt.Println(err)
			break
		}
		points = append(points, Point{x, y})
		if x < bounds.min.x {
			bounds.min.x = x
		}
		if y < bounds.min.y {
			bounds.min.y = y
		}
		if x > bounds.max.x {
			bounds.max.x = x
		}
		if y > bounds.max.y {
			bounds.max.y = y
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	return points, bounds
}

// var toutesDirections []Point = []Point{
// 	{0, -1},  // n
// 	{1, 0},   // e
// 	{0, 1},   // s
// 	{-1, 0},  // w
// 	{1, -1},  // ne
// 	{1, 1},   // se
// 	{-1, 1},  // sw
// 	{-1, -1}, // nw
// }

func partOne() int {
	defer duration(time.Now(), "part 1")

	var points, bounds = loadInput()
	fmt.Println("points", len(points))
	fmt.Println("bounds", bounds)

	var inf map[Point]bool = make(map[Point]bool)
	var area map[Point]int = make(map[Point]int)

	// for all points in our imaginary infinite 'grid'
	for y := bounds.min.y; y <= bounds.max.y; y++ {
		for x := bounds.min.x; x <= bounds.max.x; x++ {
			var tie bool = false
			var min_dist int = math.MaxInt64
			var min_pt Point
			for _, pt := range points {
				cur_dist := pt.manhatten(Point{x, y})
				if cur_dist == min_dist {
					tie = true
				}
				if cur_dist < min_dist {
					tie = false
					min_dist = cur_dist
					min_pt = pt
				}
			}
			if !tie {
				// filter infinite
				if inf[min_pt] || x == bounds.min.x || x == bounds.max.y || y == bounds.min.y || y == bounds.max.y {
					inf[min_pt] = true
					continue
				}
				area[min_pt]++
			}
		}
	}

	// var max_pt Point
	var max_area int
	for pt, a := range area {
		if a > max_area && !inf[pt] {
			max_area = a
			// max_pt = pt
		}
	}
	// fmt.Println(max_pt)
	// fmt.Println(max_area)
	return max_area
}

func partTwo() int {
	defer duration(time.Now(), "part 2")

	var points, bounds = loadInput()
	fmt.Println("points", len(points))
	fmt.Println("bounds", bounds)

	// What is the size of the region containing all locations
	// which have a total distance to all given coordinates of less than 10000?
	var big_t int
	for y := bounds.min.y; y <= bounds.max.y; y++ {
		for x := bounds.min.x; x <= bounds.max.x; x++ {
			var t int
			for _, pt := range points {
				t += pt.manhatten(Point{x, y})
			}
			if t < 10000 {
				big_t += 1
			}
		}
	}

	return big_t
}

func main() {
	defer duration(time.Now(), "main")

	// test1 should be 5, 17
	fmt.Println(partOne()) // 3420 (3998 WRONG)
	fmt.Println(partTwo()) // 46667
}

/*
	Two previous failed attempts:
	1. Tried modeling an actual grid (using points in a map)
	and growing the area of each point outwards, stepwise
	Produced random answers (some of which were right)
	because the order of items in a map is not only undefined
	but also (apparently) varies between runs.
	2. Tried implementing a algorithm (from the solutions subreddit)
	that I didn't properly understand, it produced a commonly-seen
	'wrong' answer (3998 for part 1).
*/

/*
$ go run main.go
points 50
bounds {{48 42} {347 338}}
part 1 10.631217ms
3420
points 50
bounds {{48 42} {347 338}}
part 2 5.231317ms
46667
main 15.886124ms
*/
