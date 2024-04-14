// https://adventofcode.com/2022/day/19
// ProggyVector

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

type (
	Blueprint struct {
		oreRobotOreCost,
		clayRobotOreCost,
		obsidianRobotOreCost,
		obsidianRobotClayCost,
		geodeRobotOreCost,
		geodeRobotClayCost int
	}
	State struct {
		ore, clay, obsidian, geodes                        int
		oreRobots, clayRobots, obsidianRobots, geodeRobots int
		t                                                  int
	}
)

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

func solve(blueprint Blueprint, minutes int) int {
	var best int

	// key optimization from jonathan paulson, reduces runtimes from minutes to seconds
	maxOre := max(blueprint.oreRobotOreCost,
		blueprint.clayRobotOreCost,
		blueprint.obsidianRobotOreCost,
		blueprint.geodeRobotOreCost)

	// "you have exactly one ore-collecting robot in your pack
	// that you can use to kickstart the whole operation"
	var q []State = []State{{0, 0, 0, 0, 1, 0, 0, 0, minutes}}
	var seen map[State]struct{} = map[State]struct{}{}
	for len(q) > 0 {
		var state State
		state, q = q[0], q[1:]

		best = max(best, state.geodes)
		if state.t == 0 {
			// out of time, but keep trying incase there's a better solution
			continue
		}

		// optimizations ...
		if state.oreRobots >= maxOre {
			state.oreRobots = maxOre
		}
		if state.clayRobots >= blueprint.obsidianRobotClayCost {
			state.clayRobots = blueprint.obsidianRobotClayCost
		}
		if state.obsidianRobots >= blueprint.geodeRobotClayCost {
			state.obsidianRobots = blueprint.geodeRobotClayCost
		}
		if state.ore >= state.t*maxOre-state.oreRobots*(state.t-1) {
			state.ore = state.t*maxOre - state.oreRobots*(state.t-1)
		}
		if state.clay >= state.t*blueprint.obsidianRobotClayCost-state.clayRobots*(state.t-1) {
			state.clay = state.t*blueprint.obsidianRobotClayCost - state.clayRobots*(state.t-1)
		}
		if state.obsidian >= state.t*blueprint.geodeRobotClayCost-state.obsidianRobots*(state.t-1) {
			state.obsidian = state.t*blueprint.geodeRobotClayCost - state.obsidianRobots*(state.t-1)
		}

		if _, ok := seen[state]; ok {
			continue
		}
		seen[state] = struct{}{}

		// just accumulate ore/clay/obsidian/geodes
		q = append(q, State{
			state.ore + state.oreRobots,
			state.clay + state.clayRobots,
			state.obsidian + state.obsidianRobots,
			state.geodes + state.geodeRobots,
			state.oreRobots,
			state.clayRobots,
			state.obsidianRobots,
			state.geodeRobots,
			state.t - 1,
		})
		if state.ore >= blueprint.oreRobotOreCost {
			// buy a new ore robot
			q = append(q, State{
				state.ore - blueprint.oreRobotOreCost + state.oreRobots,
				state.clay + state.clayRobots,
				state.obsidian + state.obsidianRobots,
				state.geodes + state.geodeRobots,
				state.oreRobots + 1,
				state.clayRobots,
				state.obsidianRobots,
				state.geodeRobots,
				state.t - 1,
			})
		}
		if state.ore >= blueprint.clayRobotOreCost {
			// buy a new clay robot
			q = append(q, State{
				state.ore - blueprint.clayRobotOreCost + state.oreRobots,
				state.clay + state.clayRobots,
				state.obsidian + state.obsidianRobots,
				state.geodes + state.geodeRobots,
				state.oreRobots,
				state.clayRobots + 1,
				state.obsidianRobots,
				state.geodeRobots,
				state.t - 1,
			})
		}
		if state.ore >= blueprint.obsidianRobotOreCost && state.clay >= blueprint.obsidianRobotClayCost {
			// buy a new obsidian robot
			q = append(q, State{
				state.ore - blueprint.obsidianRobotOreCost + state.oreRobots,
				state.clay - blueprint.obsidianRobotClayCost + state.clayRobots,
				state.obsidian + state.obsidianRobots,
				state.geodes + state.geodeRobots,
				state.oreRobots,
				state.clayRobots,
				state.obsidianRobots + 1,
				state.geodeRobots,
				state.t - 1,
			})
		}
		if state.ore >= blueprint.geodeRobotOreCost && state.obsidian >= blueprint.geodeRobotClayCost {
			// buy a new geode robot
			q = append(q, State{
				state.ore - blueprint.geodeRobotOreCost + state.oreRobots,
				state.clay + state.clayRobots,
				state.obsidian - blueprint.geodeRobotClayCost + state.obsidianRobots,
				state.geodes + state.geodeRobots,
				state.oreRobots,
				state.clayRobots,
				state.obsidianRobots,
				state.geodeRobots + 1,
				state.t - 1,
			})
		}
	}
	return best
}

func part1(in string, expected int) (result int) {
	defer duration(time.Now(), "part 1")

	scanner := bufio.NewScanner(strings.NewReader(in))
	for scanner.Scan() {
		var blueprint Blueprint
		var id int
		if n, err := fmt.Sscanf(scanner.Text(), "Blueprint %d: Each ore robot costs %d ore. Each clay robot costs %d ore. Each obsidian robot costs %d ore and %d clay. Each geode robot costs %d ore and %d obsidian.",
			&id,
			&blueprint.oreRobotOreCost,
			&blueprint.clayRobotOreCost,
			&blueprint.obsidianRobotOreCost,
			&blueprint.obsidianRobotClayCost,
			&blueprint.geodeRobotOreCost,
			&blueprint.geodeRobotClayCost); n != 7 {
			fmt.Println(err)
			break
		}
		result += id * solve(blueprint, 24)
	}
	report(expected, result)
	return result
}

func part2(in string, expected int) (result int) {
	defer duration(time.Now(), "part 2")

	result = 1 // we seek product
	scanner := bufio.NewScanner(strings.NewReader(in))
	for scanner.Scan() {
		var blueprint Blueprint
		var id int
		if n, err := fmt.Sscanf(scanner.Text(), "Blueprint %d: Each ore robot costs %d ore. Each clay robot costs %d ore. Each obsidian robot costs %d ore and %d clay. Each geode robot costs %d ore and %d obsidian.",
			&id,
			&blueprint.oreRobotOreCost,
			&blueprint.clayRobotOreCost,
			&blueprint.obsidianRobotOreCost,
			&blueprint.obsidianRobotClayCost,
			&blueprint.geodeRobotOreCost,
			&blueprint.geodeRobotClayCost); n != 7 {
			fmt.Println(err)
			break
		}
		result *= solve(blueprint, 32)

		if id == 3 {
			break
		}
	}
	report(expected, result)
	return result
}

func main() {
	defer duration(time.Now(), "main")

	// part1(test1, 33)
	part1(input, 1675)

	part2(input, 6840)

	{
		var memStats runtime.MemStats
		runtime.ReadMemStats(&memStats)
		fmt.Printf("Heap memory (in bytes): %d\n", memStats.HeapAlloc)
		fmt.Printf("Number of garbage collections: %d\n", memStats.NumGC)
	}
}

/*
$ go run .
RIGHT ANSWER: 1675
part 1 duration 1.383619524s
RIGHT ANSWER: 6840
part 2 duration 5.610840504s
Heap memory (in bytes): 1175247456
Number of garbage collections: 393
main duration 6.994484695s
*/
