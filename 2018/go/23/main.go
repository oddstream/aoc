// https://adventofcode.com/2018/day/23
package main

import (
	"bufio"
	"container/heap"
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

//go:embed input.txt
var input string

type (
	Point struct {
		x, y, z int
	}
	Nanobot struct {
		Point
		r int
	}
	PQItem struct {
		dist, dir int
	}
	PriorityQueue []*PQItem
)

// manhatten distance from origin {0,0,0}
func (p Point) manhatten() int {
	return abs(p.x) + abs(p.y) + abs(p.z)
}

// manhatten distance between two nanobots
func (n Nanobot) manhatten(m Nanobot) int {
	return abs(n.x-m.x) + abs(n.y-m.y) + abs(n.z-m.z)
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].dist < pq[j].dist }
func (pq PriorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PriorityQueue) Push(x any) {
	item := x.(*PQItem)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
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

func loadNanobots(in string) []Nanobot {
	var nanobots []Nanobot
	scanner := bufio.NewScanner(strings.NewReader(in))
	for scanner.Scan() {
		var nanobot Nanobot
		if n, err := fmt.Sscanf(scanner.Text(), "pos=<%d,%d,%d>, r=%d",
			&nanobot.x, &nanobot.y, &nanobot.z, &nanobot.r); n != 4 {
			fmt.Println(err)
			break
		}
		nanobots = append(nanobots, nanobot)
	}
	return nanobots
}

func part1(in string, expected int) (result int) {
	defer duration(time.Now(), "part 1")

	var nanobots []Nanobot = loadNanobots(in)
	var largest Nanobot = nanobots[0]
	for _, bot := range nanobots {
		if bot.r > largest.r {
			largest = bot
		}
	}
	for _, bot := range nanobots {
		if bot.manhatten(largest) <= largest.r {
			result += 1
		}
	}
	report(expected, result)
	return result
}

func part2(in string, expected int) (result int) {
	defer duration(time.Now(), "part 2")

	var q PriorityQueue = PriorityQueue{}
	heap.Init(&q)

	// for each bot, caluclate the manhatten distance to the origin {0,0,0},
	// add furthest and nearest radius to make a radius,
	// and add each to a priority queue
	for _, bot := range loadNanobots(in) {
		var d int = bot.Point.manhatten()
		heap.Push(&q, &PQItem{dist: max(0, d-bot.r), dir: 1})
		heap.Push(&q, &PQItem{dist: d + bot.r + 1, dir: -1})
	}
	// the queue is holding entries for the start and end of each "line segment"
	// as measured by manhattan distance from the origin.
	// At the start of the segment the e=1 adds to the total of overlapping segments.
	// The e=-1 marks the segment's end, and is used to decrease the counter.
	var count, maxCount int
	for len(q) > 0 {
		var pqi *PQItem = heap.Pop(&q).(*PQItem)
		count += pqi.dir // either +1 or -1
		// calculate the maximum number of overlapping segments
		// and the point where the maximum is hit (a manhatten distance)
		if count > maxCount {
			result = pqi.dist
			maxCount = count
		}
	}
	// idea from u/EriiKKo solution in subreddit
	report(expected, result)
	return result
}

func main() {
	defer duration(time.Now(), "main")

	// part1(test1, 7)
	part1(input, 510)
	// part2(test2, 36)
	part2(input, 108889300)

	{
		var memStats runtime.MemStats
		runtime.ReadMemStats(&memStats)
		fmt.Printf("Heap memory (in bytes): %d\n", memStats.HeapAlloc)
		fmt.Printf("Number of garbage collections: %d\n", memStats.NumGC)
	}
}

/*
$ go run .
RIGHT ANSWER: 510
part 1 duration 2.06518ms
RIGHT ANSWER: 108889300
part 2 duration 2.376158ms
Heap memory (in bytes): 795352
Number of garbage collections: 0
main duration 4.566828ms
*/
