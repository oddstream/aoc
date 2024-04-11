// https://adventofcode.com/2020/day/23

// ProggyVector

// https://gitlab.com/kurisuchan/advent-of-code-2020/-/blob/master/pkg/day23/day23.go

package main

import (
	"container/ring"
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"time"
)

var test1 string = "389125467"

var input string = "942387615"

func duration(invocation time.Time, name string) {
	fmt.Println(name, "duration", time.Since(invocation))
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

func game(cup *ring.Ring, moves int) *ring.Ring {
	var maxv int
	var cuplen int = cup.Len() // cup.Len() is VERY slow for large rings

	cupMap := make(map[int]*ring.Ring, cuplen)
	for i := 0; i < cuplen; i++ {
		var v int = cup.Value.(int)
		maxv = max(maxv, v)
		cupMap[v] = cup
		cup = cup.Next()
	}

	for i := 0; i < moves; i++ {
		// The crab picks up the three cups that are immediately clockwise of the current cup.
		// They are removed from the circle; cup spacing is adjusted as necessary to maintain the circle.
		removed := cup.Unlink(3)

		// find destination
		destination := cup.Value.(int)

		for {
			// The crab selects a destination cup: the cup with a label equal to the current cup's label minus one.
			destination -= 1
			if destination < 1 {
				destination = maxv // wrap around
			}

			// If this would select one of the cups that was just picked up,
			// the crab will keep subtracting one until it finds a cup that wasn't just picked up.
			var inRemoved bool
			removed.Do(func(v any) {
				if v.(int) == destination {
					inRemoved = true
				}
			})
			if !inRemoved {
				break
			}
		}

		// The crab places the cups it just picked up so that they are immediately clockwise of the destination cup.
		// They keep the same order as when they were picked up.
		cupMap[destination].Link(removed)

		// The crab selects a new current cup: the cup which is immediately clockwise of the current cup.
		cup = cup.Next()
	}

	return cupMap[1]
}

func part1(in string, expected string) {
	defer duration(time.Now(), "part 1")

	var cup *ring.Ring = ring.New(len(in))
	for _, r := range in {
		cup.Value = atoi(string(r))
		cup = cup.Next()
	}

	cup = game(cup, 100)

	var out strings.Builder
	cup.Do(func(v interface{}) {
		if v.(int) != 1 {
			out.WriteString(strconv.Itoa(v.(int)))
		}
	})

	var result string = out.String()
	if result != expected {
		fmt.Println("ERROR: got", result, "expected", expected)
	} else {
		fmt.Println("RIGHT ANSWER:", result)
	}
}

func part2(in string, expected int) {
	defer duration(time.Now(), "part 2")

	var cup *ring.Ring = ring.New(1_000_000)
	// Your labeling is still correct for the first few cups;
	for _, r := range in {
		cup.Value = atoi(string(r))
		cup = cup.Next()
	}
	// after that, the remaining cups are just numbered in an increasing fashion
	// starting from the number after the highest number in your list
	// and proceeding one by one until one million is reached.
	for i := len(input) + 1; i <= 1_000_000; i++ {
		cup.Value = i
		cup = cup.Next()
	}
	cup = game(cup, 10_000_000)

	// two cups that will end up immediately clockwise of cup 1
	var result int = cup.Next().Value.(int) * cup.Next().Next().Value.(int)
	if result != expected {
		fmt.Println("ERROR: got", result, "expected", expected)
	} else {
		fmt.Println("RIGHT ANSWER:", result)
	}
}

func main() {
	defer duration(time.Now(), "main")

	// part1(test1, "67384529")
	part1(input, "36542897")
	// part2(test1, 149245887792)
	part2(input, 562136730660)

	{
		var memStats runtime.MemStats
		runtime.ReadMemStats(&memStats)
		fmt.Printf("Heap memory (in bytes): %d\n", memStats.HeapAlloc)
		fmt.Printf("Number of garbage collections: %d\n", memStats.NumGC)
	}
}

/*
$ go run .
RIGHT ANSWER: 36542897
part 1 duration 33.022µs
RIGHT ANSWER: 562136730660
part 2 duration 1.656586473s
Heap memory (in bytes): 80256064
Number of garbage collections: 4
main duration 1.65665481s
*/
