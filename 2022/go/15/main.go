// https://adventofcode.com/2022/15
package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"math"
	"runtime"
	"strings"
	"time"
)

//go:embed test1.txt
var test1 string

//go:embed input.txt
var input string

type (
	Point struct {
		y, x int
	}
	Sensor struct {
		Point
		radius int
	}
)

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func manhattan(a, b Point) int {
	return abs(a.x-b.x) + abs(a.y-b.y)
}

func duration(invocation time.Time, name string) {
	fmt.Println(name, "duration", time.Since(invocation))
}

func report(expected, result int) {
	if expected != -1 {
		if result != expected {
			fmt.Println("ERROR: got", result, "expected", expected)
		} else {
			fmt.Println("RIGHT ANSWER:", result)
		}
	}
}

func part1(in string, testRow int, expected int) (result int) {
	defer duration(time.Now(), "part 1")

	var sensors []Sensor
	var beacons map[Point]struct{} = make(map[Point]struct{})
	scanner := bufio.NewScanner(strings.NewReader(in))
	for scanner.Scan() {
		var s Sensor
		var b Point
		if n, err := fmt.Sscanf(scanner.Text(), "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d",
			&s.x, &s.y, &b.x, &b.y); n != 4 {
			fmt.Println(err)
			break
		}
		s.radius = manhattan(s.Point, b)
		sensors = append(sensors, s)
		beacons[b] = struct{}{}
	}
	// fmt.Println(sensors)
	// fmt.Println(beacons)

	var xmin = math.MaxInt32
	var xmax = -math.MaxInt32
	for _, s := range sensors {
		xmin = min(xmin, s.x-s.radius)
		xmax = max(xmax, s.x+s.radius)
	}
	// fmt.Println(xmin, xmax)

	var tmp map[Point]struct{} = make(map[Point]struct{})
	for x := xmin; x <= xmax; x++ {
		var p Point = Point{x: x, y: testRow}
		if _, ok := beacons[p]; ok {
			continue
		}
		for _, s := range sensors {
			// if this point's x is in radius of any sensor
			// then there can't be a beacon here
			if manhattan(s.Point, p) > s.radius {
				continue
			}
			tmp[p] = struct{}{}
		}
	}
	result = len(tmp)
	report(expected, result)
	return result
}

func part2(in string, coordLimit int, expected int) (result int) {
	defer duration(time.Now(), "part 2")

	var sensors []Sensor

	isReachable := func(p Point) bool {
		for _, s := range sensors {
			if s.radius >= manhattan(p, s.Point) {
				return true
			}

		}
		return false
	}

	scanner := bufio.NewScanner(strings.NewReader(in))
	for scanner.Scan() {
		var s Sensor
		var b Point
		if n, err := fmt.Sscanf(scanner.Text(), "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d",
			&s.x, &s.y, &b.x, &b.y); n != 4 {
			fmt.Println(err)
			break
		}
		s.radius = manhattan(s.Point, b)
		sensors = append(sensors, s)
		// don't need to use beacon positions
	}
	// fmt.Println(sensors)

	for _, s := range sensors {
		var dist1 int = s.radius + 1
		for dist := -dist1; dist <= dist1; dist++ {
			var y int = s.y + dist
			if y < 0 {
				continue
			}
			if y > coordLimit {
				break
			}
			var xOffset int = dist1 - abs(dist)
			var xLeft int = s.x - xOffset
			var xRight int = s.x + xOffset
			if xLeft >= 0 && xLeft <= coordLimit && !isReachable(Point{x: xLeft, y: y}) {
				result = xLeft*4_000_000 + y
				goto exit
			}
			if xRight >= 0 && xRight <= coordLimit && !isReachable(Point{x: xRight, y: y}) {
				result = xRight*4_000_000 + y
				goto exit
			}
		}
	}
exit:
	report(expected, result)
	return result
}

func main() {
	defer duration(time.Now(), "main")

	// part1(test1, 10, 26)
	part1(input, 2_000_000, 5_838_453)

	// part2(test1, 20, 56_000_011)
	part2(input, 4_000_000, 12_413_999_391_794)

	{
		var memStats runtime.MemStats
		runtime.ReadMemStats(&memStats)
		fmt.Printf("Heap memory (in bytes): %d\n", memStats.HeapAlloc)
		fmt.Printf("Number of garbage collections: %d\n", memStats.NumGC)
	}
}

/*
$ go run .
RIGHT ANSWER: 5838453
part 1 duration 1.467973369s
RIGHT ANSWER: 12413999391794
part 2 duration 288.659831ms
Heap memory (in bytes): 260861600
Number of garbage collections: 7
main duration 1.756708758s
*/
