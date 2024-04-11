// https://adventofcode.com/2020/day/24
// ProggyVector
// cube coordinates with pointy tops, see https://www.redblobgames.com/grids/hexagons/

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

type Coord struct {
	q, r, s int
}

type TileMap map[Coord]struct{}

var dir2coord map[string]Coord = map[string]Coord{
	"e":  Coord{q: 1, r: 0, s: -1},
	"se": Coord{q: 0, r: 1, s: -1},
	"sw": Coord{q: -1, r: 1, s: 0},
	"w":  Coord{q: -1, r: 0, s: 1},
	"nw": Coord{q: 0, r: -1, s: 1},
	"ne": Coord{q: 1, r: -1, s: 0},
	// constraint is that q + r + s == 0
}

func (c Coord) add(d Coord) Coord {
	return Coord{c.q + d.q, c.r + d.r, c.s + d.s}
}

func (tm TileMap) neighbours(c Coord) int {
	var count int
	for _, v := range dir2coord {
		var d Coord = c.add(v)
		if _, ok := tm[d]; ok {
			count += 1
		}
	}
	return count
}

func duration(invocation time.Time, name string) {
	fmt.Println(name, "duration", time.Since(invocation))
}

func report(expected, result int) {
	if result != expected {
		fmt.Println("ERROR: got", result, "expected", expected)
	} else {
		fmt.Println("RIGHT ANSWER:", result)
	}
}

func load(in string) TileMap {
	var tiles TileMap = TileMap{}

	scanner := bufio.NewScanner(strings.NewReader(in))
	for scanner.Scan() {
		var line = scanner.Text()
		var pos Coord = Coord{0, 0, 0}
		for i := 0; i < len(line); i++ {
			var dir string
			if line[i] == 'e' || line[i] == 'w' {
				dir = string(line[i])
			} else {
				dir = string(line[i])
				i = i + 1
				dir = dir + string(line[i])
			}
			pos = pos.add(dir2coord[dir])
		}
		// all tiles are initially white
		// only record the black ones
		if _, ok := tiles[pos]; ok {
			// already a black tile here, so delete it
			delete(tiles, pos)
		} else {
			// add a black tile
			tiles[pos] = struct{}{}
		}
	}

	return tiles
}

func part1(in string, expected int) (result int) {
	defer duration(time.Now(), "part 1")

	var blacks TileMap = load(in)
	result = len(blacks)
	report(expected, result)
	return result
}

func part2(in string, expected int) (result int) {
	defer duration(time.Now(), "part 2")

	var blacks TileMap = load(in)
	for day := 0; day < 100; day++ {
		var dst TileMap = TileMap{}
		for pos := range blacks {
			var n = blacks.neighbours(pos)
			// "Any black tile with zero or more than 2 black tiles immediately adjacent to it is flipped to white."
			if n == 0 || n > 2 {
				// tile becomes white in new map
			} else {
				// tile remains black in new map
				dst[pos] = struct{}{}
			}
		}
		// "Any white tile with exactly 2 black tiles immediately adjacent to it is flipped to black."
		var whites TileMap = TileMap{}
		for pos := range blacks {
			for _, p2 := range dir2coord {
				var p3 = pos.add(p2)
				if _, ok := blacks[p3]; !ok {
					whites[p3] = struct{}{}
				}
			}
		}
		for pos := range whites {
			var n = blacks.neighbours(pos)
			if n == 2 {
				dst[pos] = struct{}{}
			}
		}
		// println(day+1, len(dst))
		blacks = dst
	}
	result = len(blacks)
	report(expected, result)
	return result
}

func main() {
	defer duration(time.Now(), "main")

	// part1(test1, 10)
	part1(input, 434)
	// part2(test1, 2208)
	part2(input, 3955)

	{
		var memStats runtime.MemStats
		runtime.ReadMemStats(&memStats)
		fmt.Printf("Heap memory (in bytes): %d\n", memStats.HeapAlloc)
		fmt.Printf("Number of garbage collections: %d\n", memStats.NumGC)
	}
}

/*
$ go run .
RIGHT ANSWER: 434
part 1 duration 485.323µs
RIGHT ANSWER: 3955
part 2 duration 197.701407ms
Heap memory (in bytes): 3227456
Number of garbage collections: 14
main duration 198.216195ms
*/
