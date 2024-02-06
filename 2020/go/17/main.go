// https://adventofcode.com/2020/day/17
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
var test1 string // 3x3

//go:embed input.txt
var input string // 8x8

// don't store inactive cells at all
// use a min/max, extended by 1 in every dimension every loop

type Point struct {
	x, y, z int
}

type Cube map[Point]struct{}

type HyperPoint struct {
	x, y, z, w int
}

type HyperCube map[HyperPoint]struct{}

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

	var cube Cube = make(Cube)

	// read layer z=0 from input
	scanner := bufio.NewScanner(strings.NewReader(in))
	var y int
	for scanner.Scan() {
		tokens := strings.Split(scanner.Text(), "")
		for x, ch := range tokens {
			if ch == "#" {
				cube[Point{x - len(tokens)/2, y - len(tokens)/2, 0}] = struct{}{}
			}
		}
		y += 1
	}
	var minmax int = (y / 2) + 1
	// fmt.Println(cube)

	var dirs []int = []int{-1, 0, 1}

	// neighbours := func(p Point) []Point {
	// 	var n []Point
	// 	for _, x := range dirs {
	// 		for _, y := range dirs {
	// 			for _, z := range dirs {
	// 				var q Point = Point{p.x + x, p.y + y, p.z + z}
	// 				if p != q {
	// 					n = append(n, q)
	// 				}
	// 			}
	// 		}
	// 	}
	// 	return n
	// }

	activeNeighbours := func(p Point) int {
		var n int
		for _, x := range dirs {
			for _, y := range dirs {
				for _, z := range dirs {
					var q Point = Point{p.x + x, p.y + y, p.z + z}
					if p != q {
						if _, ok := cube[q]; ok {
							n += 1
						}
						// not in map or "." will count as inactive
					}
				}
			}
		}
		return n
	}

	// "During a cycle, all cubes simultaneously change their state"
	// so we need to take all input from one cube while building a new cube
	var newcube Cube

	for cycle := 0; cycle < 6; cycle++ {
		// fmt.Println(cycle, minmax, len(cube))
		newcube = make(Cube)
		for x := -minmax; x <= minmax; x++ {
			for y := -minmax; y <= minmax; y++ {
				for z := -minmax; z <= minmax; z++ {
					var p Point = Point{x: x, y: y, z: z}
					var n int = activeNeighbours(p)
					if _, ok := cube[p]; ok {
						// If a cube is active and exactly 2 or 3 of its neighbors are also active,
						// the cube remains active
						if n == 2 || n == 3 {
							newcube[p] = struct{}{}
						}
					} else {
						// If a cube is inactive but exactly 3 of its neighbors are active
						// the cube becomes active. Otherwise, the cube remains inactive.
						if n == 3 {
							newcube[p] = struct{}{}
						}
					}
				}
			}
		}
		minmax += 1
		cube = newcube
	}

	result = len(newcube)
	report(expected, result)
	return result
}

func part2(in string, expected int) (result int) {
	defer duration(time.Now(), "part 2")

	var cube HyperCube = make(HyperCube)

	// read layer z=0 from input
	scanner := bufio.NewScanner(strings.NewReader(in))
	var y int
	for scanner.Scan() {
		tokens := strings.Split(scanner.Text(), "")
		for x, ch := range tokens {
			if ch == "#" {
				cube[HyperPoint{x - len(tokens)/2, y - len(tokens)/2, 0, 0}] = struct{}{}
			}
		}
		y += 1
	}
	var minmax int = (y / 2) + 1
	// fmt.Println(cube)

	var dirs []int = []int{-1, 0, 1}

	activeNeighbours := func(p HyperPoint) int {
		var n int
		for _, x := range dirs {
			for _, y := range dirs {
				for _, z := range dirs {
					for _, w := range dirs {
						var q HyperPoint = HyperPoint{p.x + x, p.y + y, p.z + z, p.w + w}
						if p != q {
							if _, ok := cube[q]; ok {
								n += 1
							}
							// not in map or "." will count as inactive
						}
					}
				}
			}
		}
		return n
	}

	// "During a cycle, all cubes simultaneously change their state"
	// so we need to take all input from one cube while building a new cube
	var newcube HyperCube

	for cycle := 0; cycle < 6; cycle++ {
		newcube = make(HyperCube)
		for x := -minmax; x <= minmax; x++ {
			for y := -minmax; y <= minmax; y++ {
				for z := -minmax; z <= minmax; z++ {
					for w := -minmax; w <= minmax; w++ {
						var p HyperPoint = HyperPoint{x: x, y: y, z: z, w: w}
						var n int = activeNeighbours(p)
						if _, ok := cube[p]; ok {
							// If a cube is active and exactly 2 or 3 of its neighbors are also active,
							// the cube remains active
							if n == 2 || n == 3 {
								newcube[p] = struct{}{}
							}
						} else {
							// If a cube is inactive but exactly 3 of its neighbors are active
							// the cube becomes active. Otherwise, the cube remains inactive.
							if n == 3 {
								newcube[p] = struct{}{}
							}
						}
					}
				}
			}
		}
		minmax += 1
		cube = newcube
	}

	result = len(newcube)
	report(expected, result)
	return result
}

func main() {
	defer duration(time.Now(), "main")

	// part1(test1, 112)
	part1(input, 359)
	// part2(test1, 848)
	part2(input, 2228)

	{
		var memStats runtime.MemStats
		runtime.ReadMemStats(&memStats)
		fmt.Printf("Heap memory (in bytes): %d\n", memStats.HeapAlloc)
		fmt.Printf("Number of garbage collections: %d\n", memStats.NumGC)
	}
}

/*
$ go run main.go
RIGHT ANSWER: 359
part 1 21.066185ms
RIGHT ANSWER: 2228
part 2 1.12160988s
Heap memory (in bytes): 1221656
Number of garbage collections: 0
main 1.142793047s
*/
