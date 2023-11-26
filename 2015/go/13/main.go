package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"oddstream/aoc/utils"
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
	// namesSet is  used to assemble namesList,
	// a list of all people's names
	// that can be used to permute all name combinations
	var namesSet utils.Set[string] = utils.NewSet[string]()

	re := regexp.MustCompile(`(?P<subject>\w+) would (?P<gainlose>\w+) (?P<amount>\d+) happiness units by sitting next to (?P<partner>\w+).`)
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		match := re.FindStringSubmatch(scanner.Text())
		result := make(map[string]string)
		for i, name := range re.SubexpNames() {
			if i != 0 && name != "" {
				result[name] = match[i]
			}
		}

		namesSet.Add(result["subject"])

		delta, _ := strconv.Atoi(result["amount"])
		if result["gainlose"] == "lose" {
			delta = -delta
		}
		if smap[result["subject"]] == nil {
			smap[result["subject"]] = make(map[string]int)
		}
		smap[result["subject"]][result["partner"]] = delta
	}
	namesList = namesSet.Members()
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
	// fmt.Println(namesList)

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
