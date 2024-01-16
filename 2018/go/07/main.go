// https://adventofcode.com/2018/day/7 The Sum of Its Parts
package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"sort"
	"strings"
	"time"
)

//go:embed input.txt
var input string

type Rule struct{ a, b string }

func duration(invocation time.Time, name string) {
	fmt.Println(name, time.Since(invocation))
}

func Map[T, U any](ts []T, f func(T) U) []U {
	us := make([]U, len(ts))
	for i := range ts {
		us[i] = f(ts[i])
	}
	return us
}

func partOne() string {
	defer duration(time.Now(), "part 1")

	// steps is a set of all known steps (A..Z)
	var steps = map[string]struct{}{}
	// rules is a set of [step step]; a condensed version of the input
	var rules = map[Rule]struct{}{}
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		var a, b string
		if n, err := fmt.Sscanf(scanner.Text(), "Step %s must be finished before step %s can begin.", &a, &b); n != 2 {
			fmt.Println(err)
			break
		}
		steps[a] = struct{}{}
		steps[b] = struct{}{}
		rules[Rule{a, b}] = struct{}{}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	// fmt.Println(lines)
	// fmt.Println(steps)

	// return a (sorted) list of steps that can be executed next
	// ie ones where the remaining steps in the set of steps
	// does not have second step in the lines of steps
	nextStep := func() []string {
		var res []string
		for s := range steps {
			found := false
			for r := range rules {
				if r.b == s {
					found = true
				}
			}
			if !found {
				res = append(res, s)
			}
		}
		sort.Strings(res)
		return res
	}

	var order string
	for len(steps) > 0 {
		cand := nextStep() // ["B", "Q", "R"] first time around
		thisStep := cand[0]
		order += thisStep
		delete(steps, thisStep) // remove finished step

		// remove all rules that contain this step
		// first time around, this is 4 lines that start with "Step B"
		for r := range rules {
			if r.a == thisStep {
				delete(rules, r)
			}
		}
	}
	return order
}

func partTwo() int {
	defer duration(time.Now(), "part 2")

	// steps is a set of all known steps (A..Z)
	var steps = map[string]struct{}{}
	// rules is a set of [step step]; a condensed version of the input
	var rules = map[Rule]struct{}{}
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		var a, b string
		if n, err := fmt.Sscanf(scanner.Text(), "Step %s must be finished before step %s can begin.", &a, &b); n != 2 {
			fmt.Println(err)
			break
		}
		steps[a] = struct{}{}
		steps[b] = struct{}{}
		rules[Rule{a, b}] = struct{}{}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	time := func(ch string) int {
		r := []rune(ch)[0]
		return 60 + int(r) - int('A')
	}

	// return a (sorted) list of steps that can be executed next
	// ie ones where the remaining steps in the set of steps
	// does not have second step in the lines of steps
	nextStep := func() []string {
		var res []string
		for s := range steps {
			found := false
			for r := range rules {
				if r.b == s {
					found = true
				}
			}
			if !found {
				res = append(res, s)
			}
		}
		sort.Strings(res)
		return res
	}

	var workers = [5]int{} // time to complete task
	var work = [5]string{} // task worker is working on

	activeWorkers := func() bool {
		for i := 0; i < 5; i++ {
			if workers[i] > 0 {
				return true
			}
		}
		return false
	}

	var t int

	for len(steps) > 0 || activeWorkers() {
		cand := nextStep() // ["B", "Q", "R"] first time around
		for i := 0; i < 5; i++ {
			workers[i] = max(workers[i]-1, 0) // count down time
			if workers[i] == 0 {
				// this worker is idle
				// delete step worker was working on
				if work[i] != "" {
					for r := range rules {
						if r.a == work[i] {
							delete(rules, r)
						}
					}
					work[i] = ""
				}
				// give worker a new step
				if len(cand) > 0 {
					var thisStep string
					thisStep, cand = cand[0], cand[1:]
					workers[i] = time(thisStep)
					work[i] = thisStep
					delete(steps, thisStep)
				}
			}
		}
		t += 1
	}

	return t
}

func main() {
	defer duration(time.Now(), "main")

	fmt.Println(partOne()) // BGKDMJCNEQRSTUZWHYLPAFIVXO
	fmt.Println(partTwo()) // 941
}

/*
$ go run main.go
part 1 725.891µs
BGKDMJCNEQRSTUZWHYLPAFIVXO
part 2 9.386386ms
941
main 10.151155ms
*/
