package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

//go:embed input.txt
var input string

type Attrib struct {
	name  string
	value int
}

func duration(invocation time.Time, name string) {
	fmt.Println(name, time.Since(invocation))
}

func Contains[T comparable](elems []T, v T) bool {
	for _, s := range elems {
		if v == s {
			return true
		}
	}
	return false
}

func Union[T comparable](a []T, b []T) []T {
	var result []T = a
	for _, elem := range b {
		if !Contains(a, elem) {
			result = append(result, elem)
		}
	}
	return result
}

// brute force (inefficient) method O(m∗n)
func Intersection[T comparable](a, b []T) []T {
	var result []T
	for _, elem := range a {
		if Contains(b, elem) {
			result = append(result, elem)
		}
	}
	return result
}

// more efficient method using a set O(m+n)
func Intersection2[T comparable](a, b []T) []T {
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

func Exclusion[T comparable](a []T, b []T) []T {
	var result []T
	for _, elem := range a {
		if !Contains(b, elem) {
			result = append(result, elem)
		}
	}
	return result
}

var aunts [][]Attrib

var readings []Attrib = []Attrib{
	{"children", 3},
	{"cats", 7},
	{"samoyeds", 2},
	{"pomeranians", 3},
	{"akitas", 0},
	{"vizslas", 0},
	{"goldfish", 5},
	{"trees", 5},
	{"cars", 2},
	{"perfumes", 1},
}

func parseInput() {
	re := regexp.MustCompile(`Sue (\d+): (\w+): (\d+), (\w+): (\d+), (\w+): (\d+)`)
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		matches := re.FindAllStringSubmatchIndex(scanner.Text(), -1)
		// auntid, _ := strconv.Atoi(scanner.Text()[matches[0][2]:matches[0][3]])
		att0 := Attrib{}
		att0.name = scanner.Text()[matches[0][4]:matches[0][5]]
		att0.value, _ = strconv.Atoi(scanner.Text()[matches[0][6]:matches[0][7]])
		att1 := Attrib{}
		att1.name = scanner.Text()[matches[0][8]:matches[0][9]]
		att1.value, _ = strconv.Atoi(scanner.Text()[matches[0][10]:matches[0][11]])
		att2 := Attrib{}
		att2.name = scanner.Text()[matches[0][12]:matches[0][13]]
		att2.value, _ = strconv.Atoi(scanner.Text()[matches[0][14]:matches[0][15]])
		// fmt.Println(aunt, att0, att1, att2)
		aunts = append(aunts, []Attrib{att0, att1, att2})
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("scanner error: %v\n", err)
	}
}

// return a list of aunts who either have the exact attribute
// or who don't have the attribute in their list
func auntsWithAttrib1(a Attrib) []int {
	var result []int
	for i, attlst := range aunts {
		if i == 0 {
			continue
		}
		var hasAttrib bool = false
		for _, att := range attlst {
			if att.name == a.name {
				hasAttrib = true
				if att.value == a.value {
					result = append(result, i)
				}
			}
		}
		if !hasAttrib {
			result = append(result, i)
		}
	}
	return result
}

func auntsWithAttrib2(a Attrib) []int {
	var result []int
	for i, attlst := range aunts {
		if i == 0 {
			continue
		}
		var hasAttrib bool = false
		for _, att := range attlst {
			if att.name == a.name {
				hasAttrib = true
				switch a.name {
				case "cats", "trees":
					if att.value > a.value {
						result = append(result, i)
					}
				case "pomeranians", "goldfish":
					if att.value < a.value {
						result = append(result, i)
					}
				default:
					if att.value == a.value {
						result = append(result, i)
					}
				}
			}
		}
		if !hasAttrib {
			result = append(result, i)
		}
	}
	return result
}

func main() {
	duration(time.Now(), "main")
	aunts = append(aunts, []Attrib{})
	parseInput()
	var result []int
	for _, a := range readings {
		auntlist := auntsWithAttrib2(a)
		// fmt.Println(a.name, a.value, auntlist)
		if result == nil {
			result = auntlist
		} else {
			result = Intersection2(result, auntlist)
		}
	}
	fmt.Println(result) // part 1 103, part 2 405
	// fmt.Println(Intersection([]int{1, 2, 3}, []int{2}))
	// fmt.Println(Intersection([]int{1, 2, 3}, []int{1, 4, 5, 6}))
}
