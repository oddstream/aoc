// https://adventofcode.com/2021/day/10
package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"runtime"
	"slices"
	"strings"
	"time"
)

//go:embed test1.txt
var test1 string

//go:embed input.txt
var input string

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

type Stack []string

func (a Stack) push(item string) Stack {
	return append(a, item)
}

func (a Stack) pop() (string, Stack) {
	return a[len(a)-1], a[:len(a)-1]
}

func (a Stack) peek() string {
	return a[len(a)-1]
}

// return remaining stack, and illegal character (or "")
func checker(in []string) ([]string, string) {
	var stack Stack
	for _, item := range in {
		switch item {
		case "(":
			stack = stack.push(")")
		case "[":
			stack = stack.push("]")
		case "{":
			stack = stack.push("}")
		case "<":
			stack = stack.push(">")
		default:
			if item == stack.peek() {
				_, stack = stack.pop()
			} else {
				return stack, item
			}
		}
	}
	return stack, ""
}

func part1(in string, expected int) (result int) {
	defer duration(time.Now(), "part 1")

	var illegalScore map[string]int = map[string]int{
		")": 3,
		"]": 57,
		"}": 1197,
		">": 25137,
	}
	scanner := bufio.NewScanner(strings.NewReader(in))
	for scanner.Scan() {
		_, ch := checker(strings.Split(scanner.Text(), ""))
		if ch != "" {
			result += illegalScore[ch]
		}
	}
	report(expected, result)
	return result
}

func part2(in string, expected int) (result int) {
	defer duration(time.Now(), "part 2")

	var completionScore map[string]int = map[string]int{
		")": 1,
		"]": 2,
		"}": 3,
		">": 4,
	}
	var scores []int
	scanner := bufio.NewScanner(strings.NewReader(in))
	for scanner.Scan() {
		rest, ch := checker(strings.Split(scanner.Text(), ""))
		if len(rest) > 0 && ch == "" {
			slices.Reverse[[]string](rest)
			var score int = 0
			for _, ch := range rest {
				score *= 5
				score += completionScore[ch]
			}
			scores = append(scores, score)
		}
	}
	slices.Sort(scores)
	result = scores[len(scores)/2]
	report(expected, result)
	return result
}

func main() {
	defer duration(time.Now(), "main")

	// part1(test1, 26397)
	part1(input, 469755)
	// part2(test1, 288957)
	part2(input, 2762335572)

	{
		var memStats runtime.MemStats
		runtime.ReadMemStats(&memStats)
		fmt.Printf("Heap memory (in bytes): %d\n", memStats.HeapAlloc)
		fmt.Printf("Number of garbage collections: %d\n", memStats.NumGC)
	}
}

/*
$ go run main.go
RIGHT ANSWER: 469755
part 1 261.004µs
RIGHT ANSWER: 2762335572
part 2 246.225µs
Heap memory (in bytes): 578816
Number of garbage collections: 0
main 603.082µs
*/
