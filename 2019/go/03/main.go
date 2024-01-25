// https://adventofcode.com/2019/day/3 Crossed Wires
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

//go:embed input.txt
var input string

type Point struct {
	y, x int
}

func (p Point) add(q Point) Point {
	return Point{y: p.y + q.y, x: p.x + q.x}
}

type Path map[Point]int

var directions map[string]Point = map[string]Point{
	"L": {y: 0, x: -1},
	"R": {y: 0, x: 1},
	"U": {y: -1, x: 0},
	"D": {y: 1, x: 0},
}

func duration(invocation time.Time, name string) {
	fmt.Println(name, time.Since(invocation))
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func manhatten(p, q Point) int {
	return abs(p.y-q.y) + abs(p.x-q.x)
}

func wires2paths(wires [][]string) []Path {
	var paths []Path
	for _, wire := range wires {
		var pos Point
		var path Path = make(Path)
		var steps int
		for _, s := range wire {
			dir := s[0:1]
			dist, err := strconv.Atoi(s[1:])
			if err != nil {
				fmt.Println(err)
				break
			}
			for ; dist > 0; dist-- {
				pos = pos.add(directions[dir])
				steps += 1
				path[pos] = steps
			}
		}
		paths = append(paths, path)
	}
	return paths
}

func partOne(wires [][]string, expected int) {
	defer duration(time.Now(), "part 1")

	var paths []Path = wires2paths(wires)

	var crosses []Point
	for p0 := range paths[0] {
		if _, ok := paths[1][p0]; ok {
			crosses = append(crosses, p0)
		}
	}
	fmt.Println(len(crosses), "crosses")

	var result int = math.MaxInt64
	for _, cross := range crosses {
		m := manhatten(cross, Point{})
		if m < result {
			result = m
		}
	}

	if result != expected {
		fmt.Println("ERROR: got", result, "expected", expected)
	} else {
		fmt.Println("CORRECT:", result)
	}
}

func partTwo(wires [][]string, expected int) {
	defer duration(time.Now(), "part 2")

	var paths []Path = wires2paths(wires)

	var result int = math.MaxInt64
	for p0 := range paths[0] {
		if _, ok := paths[1][p0]; ok {
			result = min(result, paths[0][p0]+paths[1][p0])
		}
	}

	if result != expected {
		fmt.Println("ERROR: got", result, "expected", expected)
	} else {
		fmt.Println("CORRECT:", result)
	}
}

func main() {
	defer duration(time.Now(), "main")

	// var test1 string = "R75,D30,R83,U83,L12,D49,R71,U7,L72\nU62,R66,U55,R34,D71,R55,D58,R83\n"
	// var test1 string = "R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51\nU98,R91,D20,R16,D67,R40,U7,R15,U6,R7\n"

	var wires [][]string
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		var wire []string = strings.Split(strings.Trim(scanner.Text(), "\n"), ",")
		wires = append(wires, wire)
	}

	partOne(wires, 5357)
	partTwo(wires, 101956)
}

/*
$ go run main.go
70 crosses
CORRECT: 5357
part 1 47.669388ms
CORRECT: 101956
part 2 43.408056ms
main 91.122141ms
*/
