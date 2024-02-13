// https://adventofcode.com/2021/day/12
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

//go:embed test2.txt
var test2 string

//go:embed test3.txt
var test3 string

//go:embed input.txt
var input string

// all caves differ in first two characters
// caves are either ALL UPPER or all lower case
// there aren't any infinite loops

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

func little(cave string) bool {
	return cave == strings.ToLower(cave)
}

func copyMap(in map[string]struct{}) map[string]struct{} {
	var out = make(map[string]struct{})
	for key := range in {
		out[key] = struct{}{}
	}
	return out
}

func recursive1(caves map[string][]string, current string, visited map[string]struct{}) int {
	if current == "end" {
		return 1
	}
	var count int
	for _, next := range caves[current] {
		if _, ok := visited[next]; ok && little(next) {
			continue
		}
		visited[current] = struct{}{}
		count += recursive1(caves, next, visited)
		delete(visited, current) // backtrack
	}
	return count
}

func queue1(caves map[string][]string) int {
	type QueueItem struct {
		name    string
		visited map[string]struct{}
	}
	var count int
	var q []QueueItem = []QueueItem{
		{visited: make(map[string]struct{}), name: "start"},
	}
	for len(q) > 0 {
		var cave QueueItem
		cave, q = q[0], q[1:]

		var newVisited map[string]struct{} = copyMap(cave.visited)
		if little(cave.name) {
			newVisited[cave.name] = struct{}{}
		}
		for _, next := range caves[cave.name] {
			if next == "end" {
				count += 1
				continue
			}
			if _, ok := cave.visited[next]; ok {
				continue
			}
			q = append(q, QueueItem{
				name:    next,
				visited: newVisited,
			})
		}
	}
	return count
}

func part1(in string, expected int) (result int) {
	defer duration(time.Now(), "part 1")

	var caves map[string][]string = make(map[string][]string)
	scanner := bufio.NewScanner(strings.NewReader(in))
	for scanner.Scan() {
		var line []string = strings.Split(scanner.Text(), "-")
		var a string = line[0]
		var b string = line[1]
		caves[a] = append(caves[a], b)
		caves[b] = append(caves[b], a)
	}
	// result = recursive1(caves, "start", map[string]struct{}{"start": {}})
	result = queue1(caves)
	report(expected, result)
	return result
}

func queue2(caves map[string][]string) int {
	type QueueItem struct {
		name    string
		visited map[string]struct{}
		small2  bool
	}
	var count int
	var q []QueueItem = []QueueItem{
		{name: "start", visited: make(map[string]struct{})},
	}
	for len(q) > 0 {
		var cave QueueItem
		cave, q = q[0], q[1:]

		var newVisited map[string]struct{} = copyMap(cave.visited)
		if little(cave.name) {
			if _, ok := newVisited[cave.name]; ok {
				cave.small2 = true
			} else {
				newVisited[cave.name] = struct{}{}
			}
		}
		for _, next := range caves[cave.name] {
			if next == "end" {
				count += 1
				continue
			}
			if _, ok := newVisited[next]; ok && (next == "start" || cave.small2) {
				continue
			}
			q = append(q, QueueItem{
				name:    next,
				visited: newVisited,
				small2:  cave.small2,
			})
		}
	}
	return count
}

func part2(in string, expected int) (result int) {
	defer duration(time.Now(), "part 2")

	var caves map[string][]string = make(map[string][]string)
	scanner := bufio.NewScanner(strings.NewReader(in))
	for scanner.Scan() {
		var line []string = strings.Split(scanner.Text(), "-")
		var a string = line[0]
		var b string = line[1]
		caves[a] = append(caves[a], b)
		caves[b] = append(caves[b], a)
	}
	// recursive solution; tried queue-based bfs,
	// but got in a muddle with the backtracking
	result = queue2(caves)
	report(expected, result)
	return result
}

func main() {
	defer duration(time.Now(), "main")

	// part1(test1, 10)
	// part1(test2, 19)
	// part1(test3, 226)
	part1(input, 4792)
	// part2(test1, 36)
	// part2(test2, 103)
	// part2(test3, 3509)
	part2(input, 133360)

	{
		var memStats runtime.MemStats
		runtime.ReadMemStats(&memStats)
		fmt.Printf("Heap memory (in bytes): %d\n", memStats.HeapAlloc)
		fmt.Printf("Number of garbage collections: %d\n", memStats.NumGC)
	}
}

/*
$ go run main.go
RIGHT ANSWER: 4792
part 1 4.782155ms
RIGHT ANSWER: 133360
part 2 156.961282ms
Heap memory (in bytes): 5615488
Number of garbage collections: 17
main 161.780349ms
*/
