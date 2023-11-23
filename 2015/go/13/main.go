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

//go:embed "input.txt"
var input string

type Seatingmap map[string]map[string]int

// smap is used to look up happiness when two people sit next to each other
var smap Seatingmap = make(Seatingmap)
var namesList []string

// Person is a record of each person's current happiness
type Person struct {
	name      string
	happiness int
}

// Duration of a func call
// Arguments to a defer statement are immediately evaluated and stored.
// The deferred function receives the pre-evaluated values when its invoked.
// usage: defer uDuration(time.Now(), "IntFactorial")
func Duration(invocation time.Time, name string) {
	fmt.Println(name, time.Since(invocation))
}

// parseInput and set up seating map and list of names
func parseInput() {
	// namesMap is just a set of names, used to assemble a list of all people's names
	var namesMap map[string]struct{} = make(map[string]struct{})

	scanner := bufio.NewScanner(strings.NewReader(input))
	// re := regexp.MustCompile(`(?P<name>\w+) would (?P<change>\w+) (?P<number>\d+) happiness units by sitting next to (?P<partner>\w+).`)
	re := regexp.MustCompile(`(\w+) would (\w+) (\d+) happiness units by sitting next to (\w+).`)
	// fmt.Printf("Pattern: %v\n", re.String())
	for scanner.Scan() {
		matches := re.FindAllStringSubmatchIndex(scanner.Text(), -1)
		// len(matches) == 1
		// len(matches[0]) == 10
		// 0:1 is whole match
		subject := scanner.Text()[matches[0][2]:matches[0][3]]
		ch := scanner.Text()[matches[0][4]:matches[0][5]]
		numstr := scanner.Text()[matches[0][6]:matches[0][7]]
		delta, _ := strconv.Atoi(numstr)
		if ch == "lose" {
			delta = -delta
		}
		partner := scanner.Text()[matches[0][8]:matches[0][9]]
		// fmt.Println(s)
		namesMap[subject] = struct{}{}

		if smap[subject] == nil {
			smap[subject] = make(map[string]int)
		}
		smap[subject][partner] = delta
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("scanner error: %v\n", err)
	}
	for name := range namesMap {
		namesList = append(namesList, name)
	}
}

func permute(data []Person, c chan []Person) {
	var calc func([]Person) = func(data []Person) {
		for i := 0; i < len(data)-1; i++ {
			data[i].happiness = smap[data[i].name][data[i+1].name]
			data[i].happiness += smap[data[i+1].name][data[i].name]
		}
		data[len(data)-1].happiness = smap[data[len(data)-1].name][data[0].name]
		data[len(data)-1].happiness += smap[data[0].name][data[len(data)-1].name]
	}
	var helper func([]Person, int)
	helper = func(data []Person, i int) {
		if i == len(data) {
			calc(data)
			c <- append([]Person{}, data...)
		} else {
			for j := i; j < len(data); j++ {
				data[i], data[j] = data[j], data[i]
				helper(data, i+1)
				data[i], data[j] = data[j], data[i]
			}
		}
	}
	helper(data, 0)
	close(c)
}

func main() {
	defer Duration(time.Now(), "main")

	parseInput()
	fmt.Println(namesList)

	{
		var lst []Person
		for _, name := range namesList {
			lst = append(lst, Person{name: name, happiness: 0})
		}
		c := make(chan []Person)
		go permute(lst, c)
		var biggestTotal int
		for r := range c {
			var total int
			for _, p := range r {
				total += p.happiness
			}
			// fmt.Println(total, r)
			if total > biggestTotal {
				biggestTotal = total
			}
		}
		fmt.Println("part 1", biggestTotal) // 664
	}

	// for part 2, inject person Zero and re-run
	smap["Zero"] = make(map[string]int)
	for _, name := range namesList {
		smap["Zero"][name] = 0
	}
	namesList = append(namesList, "Zero")

	{
		var lst []Person
		for _, name := range namesList {
			lst = append(lst, Person{name: name, happiness: 0})
		}
		c := make(chan []Person)
		go permute(lst, c)
		var biggestTotal int
		for r := range c {
			var total int
			for _, p := range r {
				total += p.happiness
			}
			// fmt.Println(total, r)
			if total > biggestTotal {
				biggestTotal = total
			}
		}
		fmt.Println("part 2", biggestTotal) // 640
	}

}
