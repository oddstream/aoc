// https://adventofcode.com/2016/day/20 Firewall Rules
package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"math"
	"sort"
	"strings"
	"time"
)

//go:embed input.txt
var input string

func duration(invocation time.Time, name string) {
	fmt.Println(name, time.Since(invocation))
}

type pair struct {
	a, b int
}

func maxint(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	defer duration(time.Now(), "main")

	scanner := bufio.NewScanner(strings.NewReader(input))
	var a, b int
	var records []pair
	for scanner.Scan() {
		// amazed this works, thought '-' would be interpreted as -ve prefix
		if n, err := fmt.Sscanf(scanner.Text(), "%d-%d", &a, &b); n != 2 {
			fmt.Println(err)
			break
		}
		records = append(records, pair{a, b})
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	sort.Slice(records, func(i, j int) bool { return records[i].a < records[j].a })

	var allowedCount int
	var lastMax int = -1
	var firstAllowed int = -1

	for _, record := range records {
		// count ips between 0 and last record.a
		var num int = maxint(0, record.a-lastMax-1)
		allowedCount += num
		if firstAllowed == -1 && num != 0 {
			firstAllowed = lastMax + 1
		}
		lastMax = maxint(lastMax, record.b)
	}
	allowedCount += maxint(0, math.MaxInt32-lastMax)
	fmt.Println("part 1", firstAllowed)
	fmt.Println("part 2", allowedCount)
}
