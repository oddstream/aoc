// https://adventofcode.com/2021/day/13
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
var test1 string

//go:embed input.txt
var input string

type Point struct {
	y, x int
}

func duration(invocation time.Time, name string) {
	fmt.Println(name, time.Since(invocation))
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

func size(grid map[Point]struct{}) (width, height int) {
	for p := range grid {
		width = max(width, p.x)
		height = max(height, p.y)
	}
	return
}

func display(grid map[Point]struct{}, width, height int) {
	for y := 0; y <= height; y++ {
		for x := 0; x <= width; x++ {
			if _, ok := grid[Point{x: x, y: y}]; ok {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func foldup(grid map[Point]struct{}, width, height int, n int) {
	for d := 1; d <= height/2; d++ {
		for x := 0; x <= width; x++ {
			var src Point = Point{x: x, y: n + d}
			if _, ok := grid[src]; ok {
				var dst Point = Point{x: x, y: n - d}
				grid[dst] = struct{}{}
				delete(grid, src)
			}
		}
	}
}

func foldleft(grid map[Point]struct{}, width, height int, n int) {
	for d := 1; d <= width/2; d++ {
		for y := 0; y <= height; y++ {
			var src Point = Point{x: n + d, y: y}
			if _, ok := grid[src]; ok {
				var dst Point = Point{x: n - d, y: y}
				grid[dst] = struct{}{}
				delete(grid, src)
			}
		}
	}
}

func fold(in string, once bool) {
	defer duration(time.Now(), "fold")

	var grid map[Point]struct{} = make(map[Point]struct{})
	scanner := bufio.NewScanner(strings.NewReader(in))
	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}
		var p Point
		if n, err := fmt.Sscanf(scanner.Text(), "%d,%d", &p.x, &p.y); n != 2 {
			fmt.Println(err, scanner.Text())
			break
		}
		grid[p] = struct{}{}
	}
	var width, height int = size(grid)
	// display(grid, width, height)

	for scanner.Scan() {
		var axis string
		var num int
		if line, ok := strings.CutPrefix(scanner.Text(), "fold along "); !ok {
			fmt.Println("no prefix")
			break
		} else {
			axisnum := strings.Split(line, "=")
			axis = axisnum[0]
			num = atoi(axisnum[1])
		}
		switch axis {
		case "y":
			foldup(grid, width, height, num)
			height /= 2
		case "x":
			foldleft(grid, width, height, num)
			width /= 2
		}
		if once {
			break
		}
	}

	if once {
		fmt.Println(len(grid), "dots")
	} else {
		display(grid, width, height)
	}
}

func main() {
	defer duration(time.Now(), "main")

	// fold(test1, true)  // 17
	fold(input, true)  // 661
	fold(input, false) // PFKLKCFP

	{
		var memStats runtime.MemStats
		runtime.ReadMemStats(&memStats)
		fmt.Printf("Heap memory (in bytes): %d\n", memStats.HeapAlloc)
		fmt.Printf("Number of garbage collections: %d\n", memStats.NumGC)
	}
}

/*
$ go run main.go
661 dots
fold 14.284877ms
###..####.#..#.#....#..#..##..####.###...
#..#.#....#.#..#....#.#..#..#.#....#..#..
#..#.###..##...#....##...#....###..#..#..
###..#....#.#..#....#.#..#....#....###...
#....#....#.#..#....#.#..#..#.#....#.....
#....#....#..#.####.#..#..##..#....#.....
.........................................
fold 26.350969ms
Heap memory (in bytes): 463120
Number of garbage collections: 0
main 40.804012ms
*/
