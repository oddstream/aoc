package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"time"
)

//go:embed input.txt
var input string

func duration(invocation time.Time, name string) {
	fmt.Println(name, time.Since(invocation))
}

type Point struct{ x, y int }

func (p Point) Add(q Point) Point {
	return Point{x: p.x + q.x, y: p.y + q.y}
}

/*
NW|N |..
--+--+--
SW|..|NE
--+--+--
..|S |SE
*/
var directions map[string]Point = map[string]Point{
	"n":  {0, -1},
	"ne": {1, 0},
	"se": {1, 1},
	"s":  {0, 1},
	"sw": {-1, 0},
	"nw": {-1, -1},
}

func partOne() int {
	defer duration(time.Now(), "part 1")
	re := regexp.MustCompile("[[:alpha:]]+")
	matches := re.FindAllStringSubmatch(input, -1)
	goal := Point{0, 0}
	for i := range matches {
		dir := matches[i][0]
		// if _, ok := directions[dir]; !ok {
		// 	fmt.Println("bad direction", dir)
		// 	break
		// }
		goal = goal.Add(directions[dir])
	}
	fmt.Println(goal)
	return (goal.x + goal.y + (goal.x - goal.y)) / 2
}

func partTwo() int {
	defer duration(time.Now(), "part 2")
	re := regexp.MustCompile("[[:alpha:]]+")
	matches := re.FindAllStringSubmatch(input, -1)
	goal := Point{0, 0}
	var max int = 0
	for i := range matches {
		dir := matches[i][0]
		goal = goal.Add(directions[dir])
		dist := (goal.x + goal.y + (goal.x - goal.y)) / 2
		if dist > max {
			max = dist
		}
	}
	return max
}

func main() {
	defer duration(time.Now(), "main")

	fmt.Println(partOne()) // 696
	fmt.Println(partTwo()) // 1461
}

/*
$ go run main.go
{696 483}
part 1 2.239637ms
696
part 2 2.047607ms
1461
main 4.301084ms
*/
