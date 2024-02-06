// https://adventofcode.com/2020/day/16
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

func intersect[T comparable](a, b []T) []T {
	// create an empty result array, with enough room to hold intersection
	var result = make([]T, 0, len(a))
	// create a set from the first array
	var set map[T]struct{} = make(map[T]struct{})
	for _, ele := range a {
		set[ele] = struct{}{}
	}
	// check if each element exists in the set by traversing through the second array
	for _, ele := range b {
		if _, ok := set[ele]; ok {
			// if an element exists in both arrays, add it to the intersection array
			result = append(result, ele)
		}
	}
	return result
}

func remove1[T comparable](a []T, b T) []T {
	for i, v := range a {
		if v == b {
			return append(a[:i], a[i+1:]...)
		}
	}
	return a
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

func part1(in string, expected int) (result int) {
	defer duration(time.Now(), "part 1")

	var fields map[string]map[int]struct{} = make(map[string]map[int]struct{})
	fieldrx := regexp.MustCompile("([a-z ]+): ([[:digit:]]+)-([[:digit:]]+) or ([[:digit:]]+)-([[:digit:]]+)")
	scanner := bufio.NewScanner(strings.NewReader(in))
	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}
		matches := fieldrx.FindAllStringSubmatch(scanner.Text(), -1)
		var field string = matches[0][1]
		var f1 int = atoi(matches[0][2])
		var f2 int = atoi(matches[0][3])
		var f3 int = atoi(matches[0][4])
		var f4 int = atoi(matches[0][5])
		var ranges map[int]struct{} = make(map[int]struct{})
		for i := f1; i <= f2; i++ {
			ranges[i] = struct{}{}
		}
		for i := f3; i <= f4; i++ {
			ranges[i] = struct{}{}
		}
		fields[field] = ranges
	}
	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}
		if scanner.Text() == "your ticket:" {
			continue
		}
	}
	for scanner.Scan() {
		if scanner.Text() == "nearby tickets:" {
			continue
		}
		for _, token := range strings.Split(scanner.Text(), ",") {
			num := atoi(token)
			// if num doesn't appear in any of the fields, then this
			// ticket is invalid
			var found bool = false
			for _, ranges := range fields {
				if _, ok := ranges[num]; ok {
					found = true
				}
			}
			if !found {
				result += num
			}
		}
	}

	// fmt.Println(fields)

	report(expected, result)
	return result
}

func part2(in string, expected int) (result int) {
	defer duration(time.Now(), "part 2")

	var validfields map[string]map[int]struct{} = make(map[string]map[int]struct{})
	fieldrx := regexp.MustCompile("([a-z ]+): ([[:digit:]]+)-([[:digit:]]+) or ([[:digit:]]+)-([[:digit:]]+)")
	scanner := bufio.NewScanner(strings.NewReader(in))
	// parse the field value clues
	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}
		matches := fieldrx.FindAllStringSubmatch(scanner.Text(), -1)
		var field string = matches[0][1]
		var f1 int = atoi(matches[0][2])
		var f2 int = atoi(matches[0][3])
		var f3 int = atoi(matches[0][4])
		var f4 int = atoi(matches[0][5])
		var ranges map[int]struct{} = make(map[int]struct{})
		for i := f1; i <= f2; i++ {
			ranges[i] = struct{}{}
		}
		for i := f3; i <= f4; i++ {
			ranges[i] = struct{}{}
		}
		validfields[field] = ranges
	}
	// parse your ticket
	var yourticket []int
	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}
		if scanner.Text() == "your ticket:" {
			continue
		}
		for _, token := range strings.Split(scanner.Text(), ",") {
			yourticket = append(yourticket, atoi(token))
		}
		// fmt.Println("yourticket", yourticket)
	}
	// parse the nearby tickets
	var nearbytickets [][]int
	for scanner.Scan() {
		if scanner.Text() == "nearby tickets:" {
			continue
		}
		var valid bool = true
		var ticket []int
		for _, token := range strings.Split(scanner.Text(), ",") {
			num := atoi(token)
			// if num doesn't appear in any of the fields, then this
			// ticket is invalid
			var found bool = false
			for _, ranges := range validfields {
				if _, ok := ranges[num]; ok {
					found = true
				}
			}
			if !found {
				valid = false
			}
			ticket = append(ticket, num)
		}
		if valid {
			nearbytickets = append(nearbytickets, ticket)
			// } else {
			// 	fmt.Println("invalid ticket", ticket)
		}
	}
	// "Using the valid ranges for each field, determine what order the fields appear on the tickets.
	// The order is consistent between all tickets: if seat is the third field,
	// it is the third field on every ticket, including your ticket."

	// nearbytickets always contains 20 integers (or 3 in test2 data)
	// for i, ticket := range nearbytickets {
	// 	fmt.Println(i, len(ticket))
	// }

	couldbe := func(num int) []string {
		var result []string
		for field, values := range validfields {
			if _, ok := values[num]; ok {
				result = append(result, field)
			}
		}
		return result
	}

	var nvalues = len(nearbytickets[0])
	var possibles map[int][]string = make(map[int][]string)
	for i := 0; i < nvalues; i++ {
		var colresult []string
		for j := 0; j < len(nearbytickets); j++ {
			// for each column, i = column, j = row
			num := nearbytickets[j][i]
			if len(colresult) == 0 {
				colresult = couldbe(num)
			} else {
				colresult = intersect(colresult, couldbe(num))
			}
		}
		possibles[i] = colresult
	}

	// fmt.Println(possibles)

	var results map[string]int = make(map[string]int)
	for len(possibles) > 0 {
		for num, lst := range possibles {
			if len(lst) == 1 {
				found := lst[0]
				results[found] = num
				for i := range possibles {
					possibles[i] = remove1(possibles[i], found)
				}
			}
		}
		for num, lst := range possibles {
			if len(lst) == 0 {
				delete(possibles, num)
			}
		}
	}

	// fmt.Println(results)

	/*
		 0 class
		 1 type
		 2 route
		 3 departure station
		 4 arrival track
		 5 train
		 6 departure date
		 7 price
		 8 arrival location
		 9 row
		10 duration
		11 arrival platform
		12 zone
		13 departure platform
		14 departure location
		15 departure track
		16 arrival station
		17 seat
		18 departure time
		19 wagon
	*/
	result = 1
	// for _, n := range [6]int{3, 6, 13, 14, 15, 18} {
	// 	result = result * yourticket[n]
	// }
	for field, num := range results {
		if strings.HasPrefix(field, "departure ") {
			result *= yourticket[num]
		}
	}
	report(expected, result)
	return result
}

func main() {
	defer duration(time.Now(), "main")

	// part1(test1, 71)
	part1(input, 18227)
	// part2(test2, 0)
	part2(input, 2355350878831)

	{
		var memStats runtime.MemStats
		runtime.ReadMemStats(&memStats)
		fmt.Printf("Heap memory (in bytes): %d\n", memStats.HeapAlloc)
		fmt.Printf("Number of garbage collections: %d\n", memStats.NumGC)
	}
}

/*
$ go run main.go
RIGHT ANSWER: 18227
part 1 4.615411ms
RIGHT ANSWER: 2355350878831
part 2 17.06362ms
Heap memory (in bytes): 2735344
Number of garbage collections: 3
main 21.698408ms
*/
