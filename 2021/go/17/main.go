// https://adventofcode.com/2021/17
package main

import (
	_ "embed"
	"fmt"
	"runtime"
	"time"
)

//go:embed test1.txt
var test1 string

//go:embed input.txt
var input string

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
	var xmin, ymin, xmax, ymax int
	if n, err := fmt.Sscanf(in, "target area: x=%d..%d, y=%d..%d", &xmin, &xmax, &ymin, &ymax); n != 4 {
		fmt.Println(err)
	}
	result = ((ymin + 1) * ymin) / 2
	report(expected, result)
	return result
}

func part2(in string, expected int) (result int) {
	defer duration(time.Now(), "part 2")
	var xmin, ymin, xmax, ymax int
	if n, err := fmt.Sscanf(in, "target area: x=%d..%d, y=%d..%d", &xmin, &xmax, &ymin, &ymax); n != 4 {
		fmt.Println(err)
	}
	between := func(a, n, b int) bool {
		return n >= a && n <= b
	}

	var RANGE int = 300 // found by trial and error; as small as possible
	var results map[[2]int]struct{} = make(map[[2]int]struct{})
	for vx := 1; vx < RANGE; vx++ {
		for vy := -RANGE; vy < RANGE; vy++ {
			var xvelo int = vx
			var yvelo int = vy
			var x, y int
			for n := 0; n < RANGE; n++ {
				x += xvelo
				// drag
				if xvelo > 0 {
					xvelo -= 1
				} else if xvelo < 0 {
					xvelo += 1
				}
				y += yvelo
				// gravity
				yvelo = yvelo - 1
				if between(xmin, x, xmax) && between(ymin, y, ymax) {
					results[[2]int{vx, vy}] = struct{}{}
				}
			}
		}
	}
	result = len(results)
	report(expected, result)
	return result
}

func main() {
	defer duration(time.Now(), "main")

	// part1(test1, 45)
	part1(input, 7875)
	// part2(test1, 112)
	part2(input, 2321)

	{
		var memStats runtime.MemStats
		runtime.ReadMemStats(&memStats)
		fmt.Printf("Heap memory (in bytes): %d\n", memStats.HeapAlloc)
		fmt.Printf("Number of garbage collections: %d\n", memStats.NumGC)
	}
}

/*
$ go run main.go
RIGHT ANSWER: 7875
part 1 duration 32.445µs
RIGHT ANSWER: 2321
part 2 duration 41.541903ms
Heap memory (in bytes): 304920
Number of garbage collections: 0
main duration 41.689213ms
*/
