// https://adventofcode.com/2018/day/22
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

//go:embed input.txt
var input string

const (
	ROCKY  RegionType = 0
	WET    RegionType = 1
	NARROW RegionType = 2
)
const (
	NEITHER EquipmentType = iota
	CRAMPONS
	TORCH
)

type (
	RegionType    int
	EquipmentType int
	Point         struct {
		y, x int
	}
	Region struct {
		geologicIndex, erosionLevel int
		risk                        RegionType // 'risk' doubles up as 'type'
	}
	Cave map[Point]Region
)

func (p Point) add(q Point) Point {
	return Point{y: p.y + q.y, x: p.x + q.x}
}

func (p Point) manhatten() int {
	// sum, not product, you eejit
	return abs(p.x) + abs(p.y)
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
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

func getParams(in string) (depth int, target Point) {
	scanner := bufio.NewScanner(strings.NewReader(in))
	scanner.Scan()
	if n, err := fmt.Sscanf(scanner.Text(), "depth: %d", &depth); n != 1 {
		fmt.Println(err)
	}
	scanner.Scan()
	if n, err := fmt.Sscanf(scanner.Text(), "target: %d,%d", &target.x, &target.y); n != 2 {
		fmt.Println(err)
	}
	return
}

func makeCave(depth int, target, extent Point) Cave {
	var cave Cave = make(Cave)
	for y := 0; y <= extent.y; y++ {
		for x := 0; x <= extent.x; x++ {
			var geo, ero, risk int
			if (x == 0 && y == 0) || (x == target.x && y == target.y) {
				geo = 0
			} else if x == 0 {
				geo = y * 48271
			} else if y == 0 {
				geo = x * 16807
			} else {
				geo = cave[Point{y: y, x: x - 1}].erosionLevel * cave[Point{y: y - 1, x: x}].erosionLevel
			}
			ero = (geo + depth) % 20183
			risk = ero % 3
			cave[Point{y: y, x: x}] = Region{geologicIndex: geo, erosionLevel: ero, risk: RegionType(risk)}
		}
	}
	return cave
}

func part1(in string, expected int) (result int) {
	defer duration(time.Now(), "part 1")

	var depth int
	var target Point
	depth, target = getParams(in)

	// make cave with no extra regions for part 1
	var cave Cave = makeCave(depth, target, target)

	for _, r := range cave {
		result += int(r.risk)
	}

	report(expected, result)
	return result
}

/*
			crampons	torch	neither
	rocky	Y			Y		N
	wet		Y			N		Y
	narrow	N			Y		Y
*/

func canTravelTo(next RegionType, equipped EquipmentType) bool {
	switch next {
	case ROCKY:
		return equipped != NEITHER // equipped == CRAMPONS || equipped == TORCH
	case WET:
		return equipped != TORCH // equipped == CRAMPONS || equipped == NEITHER
	case NARROW:
		return equipped != CRAMPONS // equipped == NEITHER || equipped == TORCH
	}
	panic("cannot come here")
}

func otherTool(risk RegionType, equipped EquipmentType) EquipmentType {
	switch risk {
	case ROCKY:
		if equipped == CRAMPONS {
			return TORCH
		} else if equipped == TORCH {
			return CRAMPONS
		}
	case WET:
		if equipped == CRAMPONS {
			return NEITHER
		} else if equipped == NEITHER {
			return CRAMPONS
		}
	case NARROW:
		if equipped == NEITHER {
			return TORCH
		} else if equipped == TORCH {
			return NEITHER
		}
	}
	panic("cannot come here")
}

func switchEquipment(equipped EquipmentType) []EquipmentType {
	var out []EquipmentType
	for _, eq := range []EquipmentType{CRAMPONS, TORCH, NEITHER} {
		if eq != equipped {
			out = append(out, eq)
		}
	}
	return out
}

func dijkstra(cave Cave, target Point) int {
	type SeenItem struct {
		Point
		EquipmentType
	}
	var seen map[SeenItem]int = map[SeenItem]int{}
	// "You start at 0,0 (the mouth of the cave) with the torch equipped"
	var q PriorityQueue = PriorityQueue{&PQItem{Point: Point{}, equipped: TORCH, minutes: 0}}
	heap.Init(&q)

	for len(q) > 0 {
		var item = heap.Pop(&q).(*PQItem)
		if item.Point == target {
			// "Finally, once you reach the target, you need the torch equipped
			// before you can find him in the dark. The target is always in a rocky region,
			// so if you arrive there with climbing gear equipped,
			// you will need to spend seven minutes switching to your torch."
			if item.equipped != TORCH {
				return item.minutes + 7
			} else {
				return item.minutes
			}
		}
		if _, ok := seen[SeenItem{item.Point, item.equipped}]; ok {
			continue
		}
		seen[SeenItem{item.Point, item.equipped}] = item.minutes

		// can change the result by changing the order of these directions
		for _, dir := range []Point{{-1, 0}, {1, 0}, {0, 1}, {0, -1}} {
			var np Point = item.Point.add(dir)
			// "The regions with negative X or Y are solid rock and cannot be traversed."
			if np.x < 0 || np.y < 0 {
				continue
			}
			if _, ok := cave[np]; !ok {
				fmt.Println("overflow", np)
				continue
			}
			if canTravelTo(cave[np].risk, item.equipped) {
				// "Moving to an adjacent region takes one minute."
				heap.Push(&q, &PQItem{Point: np, equipped: item.equipped, minutes: item.minutes + 1})
			} else {
				// "Switching to using the climbing gear, torch, or neither always takes seven minutes"
				for _, eq := range switchEquipment(item.equipped) {
					heap.Push(&q, &PQItem{
						Point:    item.Point,
						equipped: eq,
						minutes:  item.minutes + 7,
					})
				}
			}
		}
	}
	return -1
}

func part2(in string, expected int) (result int) {
	defer duration(time.Now(), "part 2")

	var depth int
	var target Point
	depth, target = getParams(in)

	// make cave with no extra regions for part 1
	var cave map[Point]Region = makeCave(depth, target, Point{y: target.y * 2, x: target.x * 80})

	result = dijkstra(cave, target)

	report(expected, result)
	return result
}

func main() {
	defer duration(time.Now(), "main")

	part1(test1, 114)
	part1(input, 8090)
	part2(test1, 45)
	part2(input, 992)

	{
		var memStats runtime.MemStats
		runtime.ReadMemStats(&memStats)
		fmt.Printf("Heap memory (in bytes): %d\n", memStats.HeapAlloc)
		fmt.Printf("Number of garbage collections: %d\n", memStats.NumGC)
	}
}

/*
$ go run .
*/
