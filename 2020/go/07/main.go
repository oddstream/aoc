// https://adventofcode.com/2020/day/7
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

//go:embed input.txt
var input string

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

func loadRules(in string) map[string]map[string]int {
	rhsrx := regexp.MustCompile("([[:digit:]]+) ([[:alpha:]]+ [[:alpha:]]+)")
	var rules map[string]map[string]int = make(map[string]map[string]int)
	scanner := bufio.NewScanner(strings.NewReader(in))
	for scanner.Scan() {
		var contents map[string]int = make(map[string]int)
		line := scanner.Text()
		line = strings.ReplaceAll(line, " bags", "") // don't you know it? just annoys
		line = strings.ReplaceAll(line, " bag", "")  // don't you know it? just a noise
		bags := strings.Split(line, " contain ")
		lhs := bags[0] // eg "dull aqua"
		rhs := bags[1] // eg "5 wavy cyan." or "4 dark fushia, 1 shiny purple."
		if rhs != "no other." {
			matches := rhsrx.FindAllStringSubmatch(rhs, -1)
			for _, match := range matches {
				contents[match[2]] = atoi(match[1])
			}
		}
		rules[lhs] = contents
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	return rules
}

func part1(rules map[string]map[string]int, expected int) (result int) {
	defer duration(time.Now(), "part 1")

	var found map[string]struct{} = make(map[string]struct{})
	var findBags func(string)
	findBags = func(col string) {
		for color, contents := range rules {
			if _, ok := contents[col]; ok {
				if _, ok := found[color]; !ok {
					found[color] = struct{}{}
					findBags(color)
				}
			}
		}
	}
	findBags("shiny gold")
	result = len(found)

	if expected != -1 {
		if result != expected {
			fmt.Println("ERROR: got", result, "expected", expected)
		} else {
			fmt.Println("RIGHT ANSWER:", result)
		}
	}
	return result
}

func part2(rules map[string]map[string]int, expected int) (result int) {
	defer duration(time.Now(), "part 2")

	var countBags func(string) int
	countBags = func(col string) int {
		var total int
		var contents map[string]int = rules[col]
		for color, count := range contents {
			var subcount int = countBags(color)
			if subcount > 0 {
				total += count * subcount
			}
			total += count
		}
		return total
	}
	result = countBags("shiny gold")

	if expected != -1 {
		if result != expected {
			fmt.Println("ERROR: got", result, "expected", expected)
		} else {
			fmt.Println("RIGHT ANSWER:", result)
		}
	}
	return result
}

func main() {
	defer duration(time.Now(), "main")

	var rules map[string]map[string]int = loadRules(input)

	// part1(rules, 4)
	// part2(rules, 32)
	part1(rules, 192)
	part2(rules, 12128)

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
