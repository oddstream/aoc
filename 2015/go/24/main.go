package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"log"
	"math"
	"slices"
	"strconv"
	"strings"
	"time"
)

//go:embed input.txt
var input string

func duration(invocation time.Time, name string) {
	fmt.Println(name, time.Since(invocation))
}

type Package struct {
	amount, quantum int
}

var presents []int

func sum(arr []int) int {
	var s int
	for _, n := range arr {
		s += n
	}
	return s
}

func parseInput() {
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		if num, err := strconv.Atoi(scanner.Text()); err != nil {
			log.Println(err)
		} else {
			presents = append(presents, num)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("scanner error: %v\n", err)
	}
	// sort into descending order
	slices.SortFunc(presents, func(a, b int) int { return b - a })
	// log.Println(packages)
}

func calc(remaining int, presents []int, combinations *[]Package, used int, qe int) {
	if remaining == 0 {
		*combinations = append(*combinations, Package{used, qe})
	} else if remaining > 0 && used < 6 && len(presents) > 0 {
		first := presents[0]
		rest := presents[1:]
		calc(remaining-first, rest, combinations, used+1, qe*first)
		calc(remaining, rest, combinations, used, qe)
	}
}

func findSolution(goal int) []Package {
	var combinations []Package
	calc(goal, presents, &combinations, 0, 1)
	return combinations
}

func main() {
	defer duration(time.Now(), "main")
	parseInput()
	/*
		var firstGoal = sum(presents) / 3
		var soln1 = findSolution(firstGoal)
		// log.Println(soln1)
		var min_amt = len(presents)
		var min_qe int = math.MaxInt64
		for _, p := range soln1 {
			if p.amount < min_amt {
				min_amt = p.amount
				if p.quantum < min_qe {
					min_qe = p.quantum
				}
			}
		}
		log.Println(min_amt, min_qe) // 11266889531
	*/
	var secondGoal = sum(presents) / 4
	var soln2 = findSolution(secondGoal)
	log.Println(soln2)
	var min_amt = len(presents)
	var min_qe = math.MaxInt64
	for _, p := range soln2 {
		if p.amount < min_amt {
			min_amt = p.amount
			if p.quantum < min_qe {
				min_qe = p.quantum
			}
		}
	}
	log.Println(min_amt, min_qe) // 77387711
}

/*
Python. Just need to find the smallest combination of numbers that adds up to the sum
 of the weights divided by 3 (or 4 for part 2) since you need 3 (or 4) equal groups.
 Of the combinations that satisfy that condition, find the minimum quantum entanglement.

day = 24

from functools import reduce
from itertools import combinations
from operator import mul

wts = [int(x) for x in get_input(day).split('\n')]

def day24(num_groups):
    group_size = sum(wts) // num_groups
    for i in range(len(wts)):
        qes = [reduce(mul, c) for c in combinations(wts, i)
              if sum(c) == group_size]
        if qes:
            return min(qes)

print(day24(3))
print(day24(4))
*/
