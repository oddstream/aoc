// https://adventofcode.com/2022/day/18
// ProggyVector
package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"runtime"
	"strings"
	"time"
)

//go:embed test1.txt
var test1 string

//go:embed input.txt
var input string

type Point struct {
	x, y, z int
}

func (p Point) sides() []Point {
	var out []Point = []Point{
		{p.x + 1, p.y, p.z},
		{p.x - 1, p.y, p.z},
		{p.x, p.y + 1, p.z},
		{p.x, p.y - 1, p.z},
		{p.x, p.y, p.z + 1},
		{p.x, p.y, p.z - 1},
	}
	return out
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

func part1(in string, expected int) (result int) {
	defer duration(time.Now(), "part 1")

	var cubes map[Point]struct{} = map[Point]struct{}{}
	scanner := bufio.NewScanner(strings.NewReader(in))
	for scanner.Scan() {
		var pt Point
		if n, err := fmt.Sscanf(scanner.Text(), "%d,%d,%d", &pt.x, &pt.y, &pt.z); n != 3 {
			fmt.Println(err)
			break
		}
		cubes[pt] = struct{}{}
	}

	for p := range cubes {
		for _, q := range p.sides() {
			if _, ok := cubes[q]; !ok {
				result += 1
			}
		}
	}
	report(expected, result)
	return result
}

func part2(in string, expected int) (result int) {
	defer duration(time.Now(), "part 2")

	var cubes map[Point]struct{} = map[Point]struct{}{}
	scanner := bufio.NewScanner(strings.NewReader(in))
	for scanner.Scan() {
		var pt Point
		if n, err := fmt.Sscanf(scanner.Text(), "%d,%d,%d", &pt.x, &pt.y, &pt.z); n != 3 {
			fmt.Println(err)
			break
		}
		cubes[pt] = struct{}{}
	}
	/*
		var minp Point = Point{32767, 32767, 32767} // {1 1 0}
		var maxp Point                              // {18 18 19}
		for p := range cubes {
			minp.x = min(minp.x, p.x)
			minp.y = min(minp.y, p.y)
			minp.z = min(minp.z, p.z)
			maxp.x = max(maxp.x, p.x)
			maxp.y = max(maxp.y, p.y)
			maxp.z = max(maxp.z, p.z)
		}
	*/
	var seen map[Point]struct{} = map[Point]struct{}{}
	var todo []Point = []Point{
		{-1, -1, -1},
	}

	ingrid := func(p Point) bool {
		return p.x >= -1 && p.x <= 20 && p.y >= -1 && p.y <= 20 && p.z >= -1 && p.z <= 20
	}

	for len(todo) > 0 {
		here := todo[0]
		todo = todo[1:]

		if _, ok := seen[here]; ok {
			continue
		}
		if !ingrid(here) {
			continue
		}
		if _, ok := cubes[here]; ok {
			result += 1
			continue
		}
		seen[here] = struct{}{}

		for _, s := range here.sides() {
			if _, ok := seen[s]; ok {
				continue
			}
			todo = append(todo, s)
		}
	}

	report(expected, result)
	return result
}

func main() {
	defer duration(time.Now(), "main")

	// part1(test1, 64)
	part1(input, 3466)
	// part2(test1, 58)
	part2(input, 2012)

	{
		var memStats runtime.MemStats
		runtime.ReadMemStats(&memStats)
		fmt.Printf("Heap memory (in bytes): %d\n", memStats.HeapAlloc)
		fmt.Printf("Number of garbage collections: %d\n", memStats.NumGC)
	}
}

/*
$ go run .
*/
