// https://adventofcode.com/2018/day/20
package main

import (
	_ "embed"
	"fmt"
	"runtime"
	"time"
)

//go:embed test1.txt
var test1 string

//go:embed test2.txt
var test2 string

//go:embed test3.txt
var test3 string

//go:embed test4.txt
var test4 string

//go:embed test5.txt
var test5 string

//go:embed input.txt
var input string

func duration(invocation time.Time, name string) {
	fmt.Println(name, "duration", time.Since(invocation))
}

func solve(in string) {

	type (
		Step struct {
			y, x, distance int
		}
		Key [2]int
	)
	var (
		rooms            map[Key]int = make(map[Key]int) // x,y = distance
		stack            []Step
		greatestDistance int
		distantRooms     map[Key]struct{} = make(map[Key]struct{})
		y, x, distance   int
	)

	for i := 1; i < len(in)-1; i++ { // skip leading ^
		var ch byte = in[i]
		if ch == '(' { // push
			stack = append(stack, Step{y: y, x: x, distance: distance})
			continue
		} else if ch == '|' { // peek
			var step = stack[len(stack)-1]
			y, x, distance = step.y, step.x, step.distance
			continue
		} else if ch == ')' { // pop
			var step = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			y, x, distance = step.y, step.x, step.distance
			continue
		} else if ch == 'N' {
			y -= 1
		} else if ch == 'S' {
			y += 1
		} else if ch == 'W' {
			x -= 1
		} else if ch == 'E' {
			x += 1
		} else if ch == '$' {
			break
		} else {
			fmt.Println("unexpected char in input", string(ch))
		}
		distance += 1
		var key Key = Key{y, x}
		if _, ok := rooms[key]; !ok {
			rooms[key] = distance
		} else if rooms[key] > distance {
			rooms[key] = distance
		}
		var dist = rooms[key]
		greatestDistance = max(greatestDistance, dist)
		if dist >= 1000 {
			distantRooms[key] = struct{}{}
		}
	}
	fmt.Println("part 1", greatestDistance)
	fmt.Println("part 2", len(distantRooms))
}

func main() {
	defer duration(time.Now(), "main")

	// solve(test1)
	// solve(test2)
	// solve(test3)
	// solve(test4)
	// solve(test5)
	solve(input)

	{
		var memStats runtime.MemStats
		runtime.ReadMemStats(&memStats)
		fmt.Printf("Heap memory (in bytes): %d\n", memStats.HeapAlloc)
		fmt.Printf("Number of garbage collections: %d\n", memStats.NumGC)
	}
}

/*
$ go run .
part 1 4247
part 2 8356
Heap memory (in bytes): 1781624
Number of garbage collections: 0
main duration 2.684752ms
*/
