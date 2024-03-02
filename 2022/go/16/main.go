// https://adventofcode.com/2022/16
package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"regexp"
	"runtime"
	"sort"
	"strings"
	"time"
)

//go:embed test1.txt
var test1 string

//go:embed input.txt
var input string

type (
	Valve struct {
		name     string
		flowRate int
		links    []string
	}
	Path struct {
		totalFlow int
		visited   []string
	}
)

var (
	valves    map[string]Valve
	distances map[string]map[string]int
)

func (p Path) clone() Path {
	var visited []string = make([]string, len(p.visited))
	copy(visited, p.visited)
	return Path{p.totalFlow, visited}
}

func (p *Path) add(flow int, valve string) {
	p.totalFlow += flow
	p.visited = append(p.visited, valve)
}

func contains(arr []string, val string) bool {
	for _, v := range arr {
		if v == val {
			return true
		}
	}
	return false
}

func cloneSet(in map[string]struct{}) map[string]struct{} {
	out := make(map[string]struct{})
	for k := range in {
		out[k] = struct{}{}
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

func loadValves(in string) map[string]Valve {
	var valves map[string]Valve = make(map[string]Valve)
	var rx *regexp.Regexp = regexp.MustCompile("[[:upper:]][[:upper:]]")
	scanner := bufio.NewScanner(strings.NewReader(in))
	for scanner.Scan() {
		var v Valve
		parts := strings.Split(scanner.Text(), ";")
		lhs := parts[0]
		rhs := parts[1]
		if n, err := fmt.Sscanf(lhs, "Valve %s has flow rate=%d", &v.name, &v.flowRate); n != 2 {
			fmt.Println(err)
		}
		found := rx.FindAllStringSubmatch(rhs, -1)
		for _, n := range found {
			v.links = append(v.links, n[0])
		}
		valves[v.name] = v
	}
	return valves
}

// https://en.wikipedia.org/wiki/Floyd%E2%80%93Warshall_algorithm
// provides the lengths of the paths between all pairs of valves
func floydWarshall() map[string]map[string]int {
	var dist map[string]map[string]int = make(map[string]map[string]int)

	// set up the sub maps
	for i := range valves {
		for j := range valves {
			if _, ok := dist[i]; !ok {
				dist[i] = make(map[string]int)
			}
			if i == j {
				dist[i][j] = 0 // self-connects, nothing will be less than this
			} else if contains(valves[i].links, j) {
				dist[i][j] = 1 // direct link from A to B
			} else {
				dist[i][j] = 999 // greater than the longest path
			}
		}
	}

	// now for the black magic
	for k := range valves {
		for i := range valves {
			for j := range valves {
				dist[i][j] = min(dist[i][j], dist[i][k]+dist[k][j])
			}
		}
	}

	return dist
}

// https://en.wikipedia.org/wiki/Depth-first_search
func search(current string, time int, path Path, visited map[string]struct{}) []Path {
	var paths []Path = []Path{path}
	for _, next := range valves {
		if _, ok := visited[next.name]; ok {
			continue
		}
		if next.flowRate == 0 {
			continue
		}
		nextTime := time - distances[current][next.name] - 1 // count down from 30
		if nextTime <= 0 {
			continue
		}
		nextMap := cloneSet(visited)
		nextMap[next.name] = struct{}{}
		nextPath := path.clone()
		nextPath.add(nextTime*valves[next.name].flowRate, next.name)
		paths = append(paths, search(next.name, nextTime, nextPath, nextMap)...)
	}
	return paths
}

func part1(in string, expected int) (result int) {
	defer duration(time.Now(), "part 1")

	valves = loadValves(in)
	distances = floydWarshall()

	var possiblePaths []Path = search("AA", 30, Path{}, make(map[string]struct{}))
	// fmt.Println(len(possiblePaths), "possible paths found")	// 199812
	sort.Slice(possiblePaths, func(i, j int) bool {
		return possiblePaths[i].totalFlow > possiblePaths[j].totalFlow
	})
	result = possiblePaths[0].totalFlow

	report(expected, result)
	return result
}

func part2(in string, expected int) (result int) {
	defer duration(time.Now(), "part 2")

	valves = loadValves(in)
	distances = floydWarshall()

	// find all possible paths with the new 26 minute constraint
	var possiblePaths []Path = search("AA", 26, Path{}, make(map[string]struct{}))
	// fmt.Println(len(possiblePaths), "possible paths found")	// 44702

	// slow test for an empty intersection
	containsAny := func(slice1, slice2 []string) bool {
		for _, s1 := range slice1 {
			for _, s2 := range slice2 {
				if s1 == s2 {
					return true
				}
			}
		}
		return false
	}

	// find a second path that does not cross with any of the first path
	for i := 0; i < len(possiblePaths)-1; i++ {
		a := possiblePaths[i]
		for j := i + 1; j < len(possiblePaths); j++ {
			b := possiblePaths[j]
			var f int = a.totalFlow + b.totalFlow
			if f > result && !containsAny(a.visited, b.visited) {
				result = f
			}
		}
	}

	report(expected, result)
	return result
}

func main() {
	defer duration(time.Now(), "main")

	// part1(test1, 1651)
	part1(input, 1850)
	// part2(test1, 1707)
	part2(input, 2306)

	{
		var memStats runtime.MemStats
		runtime.ReadMemStats(&memStats)
		fmt.Printf("Heap memory (in bytes): %d\n", memStats.HeapAlloc)
		fmt.Printf("Number of garbage collections: %d\n", memStats.NumGC)
	}
}

/*
$ go run .
RIGHT ANSWER: 1850
part 1 duration 645.920452ms
RIGHT ANSWER: 2306
part 2 duration 909.622617ms
Heap memory (in bytes): 12620896
Number of garbage collections: 28
main duration 1.555577389s
*/
