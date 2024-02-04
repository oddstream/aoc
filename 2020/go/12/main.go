// https://adventofcode.com/
package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"time"
)

//go:embed test1.txt
var test1 string // 5 lines

//go:embed input.txt
var input string // 764 lines

type Point struct {
	y, x int
}

func (p Point) add(q Point) Point {
	return Point{p.y + q.y, p.x + q.x}
}

func (p Point) manhatten(q Point) int {
	return abs(p.y-q.y) + abs(p.x-q.x)
}

var directions map[string]Point = map[string]Point{
	"N": {x: 0, y: -1},
	"E": {x: 1, y: 0},
	"W": {x: -1, y: 0},
	"S": {x: 0, y: 1},
}

var turnLeft map[string]string = map[string]string{
	"N": "W",
	"E": "N",
	"W": "S",
	"S": "E",
}

var turnRight map[string]string = map[string]string{
	"N": "E",
	"E": "S",
	"W": "N",
	"S": "W",
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

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func duration(invocation time.Time, name string) {
	fmt.Println(name, time.Since(invocation))
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

func part1(in string, expected int) (result int) {
	defer duration(time.Now(), "part 1")

	var ship Point
	var heading string = "E"
	var dir string
	var amt int
	scanner := bufio.NewScanner(strings.NewReader(in))
	for scanner.Scan() {
		line := scanner.Text()
		dir = line[0:1]
		amt = atoi(line[1:])
		switch dir {
		case "N":
			ship.y -= amt
		case "E":
			ship.x += amt
		case "W":
			ship.x -= amt
		case "S":
			ship.y += amt
		case "L":
			for ; amt > 0; amt -= 90 {
				heading = turnLeft[heading]
			}
		case "R":
			for ; amt > 0; amt -= 90 {
				heading = turnRight[heading]
			}
		case "F":
			for ; amt > 0; amt-- {
				ship = ship.add(directions[heading])
			}
		default:
			fmt.Println("unknown input", scanner.Text())
		}
	}
	result = Point{}.manhatten(ship) // ship started at 0,0
	report(expected, result)
	return result
}

func part2(in string, expected int) (result int) {
	defer duration(time.Now(), "part 2")

	var ship Point
	// "The waypoint starts 10 units east and 1 unit north relative to the ship."
	// "The waypoint is relative to the ship;"
	// if the ship moves, the waypoint moves with it
	var wp Point = Point{x: 10, y: -1}

	var dir string
	var amt int
	scanner := bufio.NewScanner(strings.NewReader(in))
	for scanner.Scan() {
		line := scanner.Text()
		dir = line[0:1]
		amt = atoi(line[1:])
		switch dir {
		case "N":
			wp.y -= amt
		case "E":
			wp.x += amt
		case "W":
			wp.x -= amt
		case "S":
			wp.y += amt
		case "L":
			// rotate the waypoint left/anti-clockwise around the ship
			// eg 4 E, 10 S := 10 E, 4 N
			for ; amt > 0; amt -= 90 {
				wp.x, wp.y = wp.y, -wp.x
			}
		case "R":
			// rotate the waypoint right/clockwise around the ship
			// 10 E, 4 N := 4 E, 10 S
			for ; amt > 0; amt -= 90 {
				wp.x, wp.y = -wp.y, wp.x
			}
		case "F":
			// move the ship to the waypoint amt times
			ship.x += amt * wp.x
			ship.y += amt * wp.y
		default:
			fmt.Println("unknown input", scanner.Text())
		}
	}
	result = Point{}.manhatten(ship) // ship started at 0,0
	report(expected, result)
	return result
}

func main() {
	defer duration(time.Now(), "main")

	// part1(test1, 25)
	part1(input, 521)
	// part2(test1, 286)
	part2(input, 22848)

	{
		var memStats runtime.MemStats
		runtime.ReadMemStats(&memStats)
		fmt.Printf("Heap memory (in bytes): %d\n", memStats.HeapAlloc)
		fmt.Printf("Number of garbage collections: %d\n", memStats.NumGC)
	}
}

/*
$ go run main.go
RIGHT ANSWER: 521
part 1 147.156µs
RIGHT ANSWER: 22848
part 2 37.415µs
Heap memory (in bytes): 147968
Number of garbage collections: 0
main 268.765µs
*/
