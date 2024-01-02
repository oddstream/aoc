// https://adventofcode.com/2016/day/24 Air Duct Spelunking
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

// input is 37 lines 184 columns, surrounded by #
// start point (0) line 12, column 36
// highest waypoint is 7

func duration(invocation time.Time, name string) {
	fmt.Println(name, time.Since(invocation))
}

type Point struct {
	y, x int
}

type Position struct {
	y, x, moves int
}

func (p Position) point() Point {
	return Point{x: p.x, y: p.y}
}

func loadInput() [][]byte {
	var grid [][]byte
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		grid = append(grid, []byte(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	return grid
}

func findWaypoint(grid [][]byte, wp byte) Point {
	for y, line := range grid {
		for x, b := range line {
			if b == wp {
				return Point{y: y, x: x}
			}
		}
	}
	return Point{}
}

func waypoints(grid [][]byte) []byte {
	var wps []byte
	for _, line := range grid {
		for _, b := range line {
			if b >= '0' && b <= '9' {
				wps = append(wps, b)
			}
		}
	}
	return wps
}

// non-backtracking https://en.wikipedia.org/wiki/Breadth-first_search
// between two points
func bfs(grid [][]byte, start byte, goal byte) int {
	var seen map[Point]struct{} = make(map[Point]struct{})
	var startpos Point = findWaypoint(grid, start)
	var q []Position = []Position{{y: startpos.y, x: startpos.x, moves: 0}}
	for len(q) > 0 {
		var p Position
		p, q = q[0], q[1:]
		for _, pos := range []Point{{x: 0, y: -1}, {x: 1, y: 0}, {x: 0, y: 1}, {x: -1, y: 0}} {
			var np Position = Position{x: p.x + pos.x, y: p.y + pos.y, moves: p.moves + 1}
			var b byte = grid[np.y][np.x]
			if b == '#' {
				continue
			}
			if _, ok := seen[np.point()]; !ok {
				if b == goal {
					return np.moves
				}
				seen[np.point()] = struct{}{}
				q = append(q, np)
			}
		}
	}
	return -1
}

// https://en.wikipedia.org/wiki/Heap%27s_algorithm
// to generate all possible permutations of []byte
func permutations(arr []byte) [][]byte {
	var helper func([]byte, int)
	res := [][]byte{}

	helper = func(arr []byte, n int) {
		if n == 1 {
			tmp := make([]byte, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
}

// brute-force https://en.wikipedia.org/wiki/Travelling_salesman_problem
func partOne(grid [][]byte, distmap map[[2]byte]int, perms [][]byte) int {
	var result = math.MaxInt64
	for _, perm := range perms {
		// always start from '0' ...
		if perm[0] != '0' {
			continue
		}
		var dist = 0
		for i := 1; i < len(perm); i++ {
			if d, ok := distmap[[2]byte{perm[i-1], perm[i]}]; ok {
				dist += d
			} else {
				fmt.Println("no distance for", perm[i-1], perm[i])
			}
		}
		if dist < result {
			result = dist
		}
	}
	return result
}

// brute-force https://en.wikipedia.org/wiki/Travelling_salesman_problem
func partTwo(grid [][]byte, distmap map[[2]byte]int, perms [][]byte) int {
	var result = math.MaxInt64
	for _, perm := range perms {
		// always start from '0' ...
		if perm[0] != '0' {
			continue
		}
		var perm2 []byte = make([]byte, len(perm)+1)
		copy(perm2, perm)
		perm2[len(perm)] = '0'
		var dist = 0
		for i := 1; i < len(perm2); i++ {
			if d, ok := distmap[[2]byte{perm2[i-1], perm2[i]}]; ok {
				dist += d
			} else {
				fmt.Println("no distance for", perm2[i-1], perm2[i])
			}
		}
		if dist < result {
			result = dist
		}
	}
	return result
}

func main() {
	defer duration(time.Now(), "main")
	grid := loadInput()
	wps := waypoints(grid)
	var distmap map[[2]byte]int = make(map[[2]byte]int)
	for i := 0; i < len(wps); i++ {
		for j := i + 1; j < len(wps); j++ {
			a, b := wps[i], wps[j]
			dist := bfs(grid, a, b)
			distmap[[2]byte{a, b}] = dist
			distmap[[2]byte{b, a}] = dist
		}
	}
	var perms [][]byte = permutations(wps)
	fmt.Println(partOne(grid, distmap, perms)) // 502
	fmt.Println(partTwo(grid, distmap, perms)) // 724
}

/*
$ go run main.go
502
724
main 17.464215ms
*/
