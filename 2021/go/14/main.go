// https://adventofcode.com/2021/day/14
package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"math"
	"runtime"
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

func apply(rules map[string]string, in string) (out string) {
	/*
		NNCB (len 4)
		NN NC CB (len 4, pairs 3)
		NCN NBC CHB
		NC + NB + CH + B (first two + last from template)
		NCNBCHB (len 7)
		first char is always first from template
		last char is always last from template
	*/

	// this could be made faster
	// by ranging over runes
	// but we can't use it
	// for part 2
	//
	// so there's no point
	var sb strings.Builder
	sb.Grow(len(in) * 2)
	for i := 0; i < len(in)-1; i++ {
		var pair string = in[i : i+2]
		var ins string = rules[pair]
		sb.WriteString(pair[:1])
		sb.WriteString(ins)
	}
	sb.WriteString(in[len(in)-1:])
	out = sb.String()
	return
}

func score(in string) int {
	var scores map[rune]int = make(map[rune]int)
	for _, ch := range in {
		scores[ch] += 1
	}
	var lce int = math.MaxInt64 // least common element
	var mce int                 // moist common element
	for _, n := range scores {
		lce = min(lce, n)
		mce = max(mce, n)
	}
	return mce - lce
}

func part1(in string, steps int, expected int) (result int) {
	defer duration(time.Now(), fmt.Sprintf("part 1 %d steps", steps))

	var template string
	var rules map[string]string = make(map[string]string)
	scanner := bufio.NewScanner(strings.NewReader(in))
	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}
		template = scanner.Text()
	}
	for scanner.Scan() {
		var rule = strings.Split(scanner.Text(), " -> ")
		rules[rule[0]] = rule[1]
	}
	// fmt.Println(template)
	// fmt.Println(rules)

	var polymer string = template
	for i := 0; i < steps; i++ {
		fmt.Print(".")
		polymer = apply(rules, polymer)
	}
	fmt.Println()
	fmt.Println("polymer length", len(polymer))

	result = score(polymer)

	report(expected, result)
	return result
}

func part2(in string, steps int, expected int) (result int) {
	defer duration(time.Now(), fmt.Sprintf("part 2 %d steps", steps))

	var template string
	var rules map[string]string = make(map[string]string)
	scanner := bufio.NewScanner(strings.NewReader(in))
	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}
		template = scanner.Text()
	}
	for scanner.Scan() {
		var rule = strings.Split(scanner.Text(), " -> ")
		rules[rule[0]] = rule[1]
	}
	// fmt.Println(template)
	// fmt.Println(rules)

	type counter map[string]int

	// make a map for tracking ALL possible pairs
	// start with all the pairs in the template
	var trackPairs counter = make(counter)
	for i := 0; i < len(template)-1; i++ {
		var pair string = template[i : i+2]
		trackPairs[pair] += 1
	}
	for i := 0; i < steps; i++ {
		var update counter = make(counter)
		for k, v := range trackPairs {
			// eg if rule is FV -> C and k is FV
			// add FC and CV to pair tracker
			// eg or rule is FV -> C and k is EE
			// add EC and CE to rules tracker
			var pair string = string(k[0]) + rules[k]
			update[pair] += v
			pair = string(rules[k]) + string(k[1])
			update[pair] += v
		}
		trackPairs = update
	}
	var counts counter = make(counter)
	for k, v := range trackPairs {
		counts[string(k[0])] += v
	}
	var lce int = math.MaxInt64 // least common element
	var mce int                 // moist common element
	for _, n := range counts {
		lce = min(lce, n)
		mce = max(mce, n)
	}
	result = mce - lce + 1

	// https://skarlso.github.io/2021/12/14/aoc-day14/

	report(expected, result)
	return result
}

func main() {
	defer duration(time.Now(), "main")

	part1(test1, 10, 1588)
	// part1(input, 10, 3587)
	part2(test1, 40, 2188189693529)
	// part2(input, 40, 3906445077999)

	{
		var memStats runtime.MemStats
		runtime.ReadMemStats(&memStats)
		fmt.Printf("Heap memory (in bytes): %d\n", memStats.HeapAlloc)
		fmt.Printf("Number of garbage collections: %d\n", memStats.NumGC)
	}
}

/*
$ go run main.go
..........
polymer length 3073
RIGHT ANSWER: 1588
part 1 10 steps 145.179µs
RIGHT ANSWER: 2188189693529
part 2 40 steps 174.127µs
Heap memory (in bytes): 217424
Number of garbage collections: 0
main 422.262µs
*/
