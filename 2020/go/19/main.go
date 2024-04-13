// https://adventofcode.com/2020/day/19

// ProggyVector

// https://github.com/alexchao26/advent-of-code-go/blob/main/2020/day19/main.go

package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"time"
)

//go:embed test1.txt
var test1 string

//go:embed test2.txt
var test2 string

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

func loadRules(in string) (map[int]*Rule, *bufio.Scanner) {
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
	return rules, scanner
}

func resolveRules(rules map[int]*Rule, entry int) []string {
	// foreach rule, turn the options (eg 2 3 | 3 2) into list of instantiated (eg aaaabb)

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
	// no longer need rules[].options
}

func part1(in string, expected int) (result int) {
	defer duration(time.Now(), "part 1")

	var rules map[int]*Rule
	var scanner *bufio.Scanner
	rules, scanner = loadRules(in)
	fmt.Print("resolving rule 0 ... ")
	resolveRules(rules, 0)
	fmt.Println("to", len(rules[0].resolved), "strings")

	// "Your goal is to determine the number of messages that completely match rule 0"
	// for _, res := range rules[0].resolved {
	// 	fmt.Println(res)
	// }

	for scanner.Scan() {
		message := scanner.Text() // scanner.Text() is slow in a loop
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

func part2(in string, expected int) (result int) {
	// "you only need to handle the rules you have;"

	// It looks like ALL input files have 0: 8 11, with no other rules referring to 8 or 11.
	// Rule 8 is 1 or more rule 42, and if we squint, rule 11 is 1 or more rule 42
	// followed by 1 or more rule 31 plus a context check;
	// combining those, for rule 0, it is trivial to check which lines out of a
	// context-free match to 1 or more rule 42 then one or more rule 31 also
	// satisfy the constraint of more matches to 42 in the first half than to 31 in the second half.

	defer duration(time.Now(), "part 2")

	var rules map[int]*Rule
	var scanner *bufio.Scanner
	rules, scanner = loadRules(in)

	// manually inject the two new rules
	rules[8] = &Rule{
		options: [][]int{{42}, {42, 8}},
	}
	rules[11] = &Rule{
		options: [][]int{{42, 31}, {42, 11, 31}},
	}

	// 0: 8 11
	// 8: 42
	// 11: 42 31
	// 31: 64 69 | 50 67
	// 42: 64 54 | 50 40

	fmt.Print("resolving rule 31 ... ")
	resolveRules(rules, 31)
	fmt.Println("to", len(rules[31].resolved), "strings") // 128, all 8 chars long
	// for _, res := range rules[31].resolved {
	// 	fmt.Println(res)
	// }
	fmt.Print("resolving rule 42 ... ")
	resolveRules(rules, 42)
	fmt.Println("to", len(rules[42].resolved), "strings") // 128, all 8 chars long
	// for _, res := range rules[42].resolved {
	// 	fmt.Println(res)
	// }

	rule42 := fmt.Sprintf("(%s)", strings.Join(rules[42].resolved, "|"))
	rule31 := fmt.Sprintf("(%s)", strings.Join(rules[31].resolved, "|"))

	// 468 messages to check
	for scanner.Scan() {
		message := scanner.Text()
		if len(message)%8 == 0 { // turns out, they are all mod 8 = 0
			for i := 1; i < 5; i++ {
				// rule 0 is matched by X repetitions of rule 42 followed by Y repetitions of rule 31, where X > Y.
				// x{n}           exactly n x
				pattern := regexp.MustCompile(fmt.Sprintf("^%s+%s{%d}%s{%d}$", rule42, rule42, i, rule31, i))
				if pattern.MatchString(message) {
					result += 1
					break
				}
			}
		}
	}

	report(expected, result)
	return result
}

func main() {
	defer duration(time.Now(), "main")

	// part1(test1, 2)
	part1(input, 226)

	// part2(test2, 12)
	part2(input, 355)

	{
		var memStats runtime.MemStats
		runtime.ReadMemStats(&memStats)
		fmt.Printf("Heap memory (in bytes): %d\n", memStats.HeapAlloc)
		fmt.Printf("Number of garbage collections: %d\n", memStats.NumGC)
	}
}

/*
$ go run main.go
resolving rule 0 ... to 2097152 strings
RIGHT ANSWER: 226
part 1 1.705232771s
resolving rule 31 ... to 128 strings
resolving rule 42 ... to 128 strings
RIGHT ANSWER: 355
part 2 732.807133ms
Heap memory (in bytes): 2907944
Number of garbage collections: 445
main 2.438066861s
*/
