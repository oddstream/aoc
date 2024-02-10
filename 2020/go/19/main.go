// https://adventofcode.com/2020/day/19
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

// numbers rules are FOLLOWED-BY/AND'd and OR'd together
// never more than one pipe in a rule
// rules up to and including 128
// one rule for "a", one rule for "b"

// https://github.com/alexchao26/advent-of-code-go/blob/main/2020/day19/main.go
// recursively resolves the rules and then matches against messages

// Two ways of doing this:
// 1. taking the rules as-is and then applying them to the messages
// 2. resolving the rules before doing simpler tests against the messages
// There is some fiddly looping here, so try to reduce that by
// helping as much as possible while inputting/parsing the rules,
// which leads (kinda) to using the second approach.

type Rule struct {
	resolved []string
	options  [][]int
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

func report(expected, result int) {
	if expected != -1 {
		if result != expected {
			fmt.Println("ERROR: got", result, "expected", expected)
		} else {
			fmt.Println("RIGHT ANSWER:", result)
		}
	}
}

func resolveRules(rules map[int]*Rule, entry int) []string {
	if len(rules[entry].resolved) != 0 {
		// return a copy of resolved otherwise there's all kinds of side effect errors
		return append([]string{}, rules[entry].resolved...)
	}

	// iterate through options, resolve children and append resolved paths
	// for the current entry point
	for _, option := range rules[entry].options {
		// this will be all permutations generated from recursive calls to resolveRules
		// Note: there's probably a cleaner algorithm to do this kind of perm generation...
		resolved := []string{""}
		for _, entryPoint := range option {
			nestedResolveVals := resolveRules(rules, entryPoint)
			var newResolved []string
			for _, nextPiece := range nestedResolveVals {
				for _, prev := range resolved {
					newResolved = append(newResolved, prev+nextPiece)
				}
			}
			resolved = newResolved
		}
		rules[entry].resolved = append(rules[entry].resolved, resolved...)
	}

	return rules[entry].resolved
}

func part1(in string, expected int) (result int) {
	defer duration(time.Now(), "part 1")

	var rules map[int]*Rule = make(map[int]*Rule)

	scanner := bufio.NewScanner(strings.NewReader(in))
	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}
		tokens := strings.Split(scanner.Text(), ": ")
		rulenum := atoi(tokens[0])

		// demungle "a" and "b" for ease of parsing later
		if tokens[1] == "\"a\"" {
			rules[rulenum] = &Rule{resolved: []string{"a"}}
		} else if tokens[1] == "\"b\"" {
			rules[rulenum] = &Rule{resolved: []string{"b"}}
		} else {
			var newrule Rule = Rule{}
			// split on | to create one or two separate subrules
			subrules := strings.Split(tokens[1], " | ")
			for _, r := range subrules {
				nums := strings.Split(r, " ")
				var option []int
				for _, n := range nums {
					option = append(option, atoi(n))
				}
				newrule.options = append(newrule.options, option)
			}
			rules[rulenum] = &newrule
		}
	}

	fmt.Print("resolving rule 0 ... ")
	resolveRules(rules, 0)
	fmt.Println("done")

	for scanner.Scan() {
		message := scanner.Text() // don't use slow scanner.Text() in the following loop
		for _, res := range rules[0].resolved {
			if message == res {
				result += 1
				break
			}
		}
	}
	report(expected, result)
	return result
}

func main() {
	defer duration(time.Now(), "main")

	part1(test1, 2)
	part1(input, 226)
	//	part2(test1, 0)

	{
		var memStats runtime.MemStats
		runtime.ReadMemStats(&memStats)
		fmt.Printf("Heap memory (in bytes): %d\n", memStats.HeapAlloc)
		fmt.Printf("Number of garbage collections: %d\n", memStats.NumGC)
	}
}

/*
$ go run main.go
*/
