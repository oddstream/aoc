// https://adventofcode.com/2016/day/13 A Maze of Twisty Little Cubicles

package main

import (
	"fmt"
	"time"
)

// const favorite int = 10
const favorite int = 1364

type Point struct {
	x, y int
}

type Position struct {
	x, y, moves int
}

func (p Position) point() Point {
	return Point{x: p.x, y: p.y}
}

func duration(invocation time.Time, name string) {
	fmt.Println(name, time.Since(invocation))
}

func wall(x, y int) bool {
	n := x*x + 3*x + 2*x*y + y + y*y + favorite
	var bits int = 0
	for n > 0 {
		if n&1 == 1 {
			bits = bits + 1
		}
		n = n >> 1
	}
	return bits%2 == 1
}

func display(w, h int) {
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			if wall(x, y) {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Print("\n")
	}
}

func bfs1(start Position, stop Point) int {
	var seen map[Point]struct{} = make(map[Point]struct{})
	seen[start.point()] = struct{}{}
	var q []Position = []Position{start}
	for len(q) > 0 {
		var p Position
		p, q = q[0], q[1:]
		for _, pos := range []Point{{x: 0, y: -1}, {x: 1, y: 0}, {x: 0, y: 1}, {x: -1, y: 0}} {
			var np Position = Position{x: p.x + pos.x, y: p.y + pos.y}
			if np.x < 0 || np.y < 0 {
				continue
			}
			if _, ok := seen[np.point()]; !ok {
				if !wall(np.x, np.y) {
					seen[np.point()] = struct{}{}
					np.moves = p.moves + 1
					if np.x == stop.x && np.y == stop.y {
						return np.moves
					}
					q = append(q, np)
				}
			}
		}
	}
	return -1
}

func bfs2(start Position, limit int) int {
	var seen map[Point]struct{} = make(map[Point]struct{})
	seen[Point{start.x, start.y}] = struct{}{}
	var q []Position = []Position{start}
	for len(q) > 0 {
		var p Position
		p, q = q[0], q[1:]
		for _, pos := range []Point{{x: 0, y: -1}, {x: 1, y: 0}, {x: 0, y: 1}, {x: -1, y: 0}} {
			var np Position = Position{x: p.x + pos.x, y: p.y + pos.y}
			if np.x < 0 || np.y < 0 {
				continue
			}
			if _, ok := seen[np.point()]; !ok {
				if !wall(np.x, np.y) {
					seen[np.point()] = struct{}{}
					np.moves = p.moves + 1
					if np.moves < limit {
						q = append(q, np)
					}
				}
			}
		}
	}
	return len(seen)
}

func main() {
	defer duration(time.Now(), "main")

	// display(10, 7)
	// fmt.Println(bfs(Position{x: 1, y: 1, moves: 0}, Point{x: 7, y: 4})) // with favorite = 10
	if wall(1, 1) {
		fmt.Println("start position is a wall!")
	}
	fmt.Println(bfs1(Position{x: 1, y: 1, moves: 0}, Point{x: 31, y: 39})) // part1: 86
	fmt.Println(bfs2(Position{x: 1, y: 1, moves: 0}, 50))                  // part2: 127
}
