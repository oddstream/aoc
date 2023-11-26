// https://adventofcode.com/2015/day/9
package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

//go:embed input.txt
var input string

type Jmap map[string]map[string]int

var journeys Jmap = make(Jmap)

type Jchan struct {
	places []string // list of place names to visit
	dist   int      // total distance between places
}

var all_places []string

// create actual slices otherwise copy() will copy nothing
// copy returns the number of elements copied, which will be the minimum of len(dst) and len(src)
// var min_journey_places []string = make([]string, len(journeys))
// var max_journey_places []string = make([]string, len(journeys))
var min_journey_places []string
var max_journey_places []string
var min_journey_dist = math.MaxInt32
var max_journey_dist = 0

// duration of a func call
// Arguments to a defer statement are immediately evaluated and stored.
// The deferred function receives the pre-evaluated values when its invoked.
// usage: defer uDuration(time.Now(), "IntFactorial")
func duration(invocation time.Time, name string) {
	fmt.Println(name, time.Since(invocation))
}

func permute2(data []string, c chan Jchan) {
	var calcdist2 func([]string) int = func(places []string) int {
		var dist int
		for i := 1; i < len(places); i++ {
			dist += journeys[places[i-1]][places[i]]
		}
		return dist
	}
	var helper func([]string, int)
	helper = func(data []string, i int) {
		if i == len(data) {
			c <- Jchan{places: data, dist: calcdist2(data)}
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

func readInput() {
	defer duration(time.Now(), "readInput")
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		tokens := strings.Split(scanner.Text(), " ")
		// fmt.Println(tokens[0], tokens[2], tokens[4])
		from := tokens[0]
		dest := tokens[2]
		dist, _ := strconv.Atoi(tokens[4])
		if journeys[from] == nil {
			journeys[from] = make(map[string]int)
		}
		journeys[from][dest] = dist
		if journeys[dest] == nil {
			journeys[dest] = make(map[string]int)
		}
		journeys[dest][from] = dist
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("scanner error: %v\n", err)
	}
	for place := range journeys {
		// fmt.Println(place, lst)
		all_places = append(all_places, place)
	}
	// fmt.Println(len(all_places), all_places)
}

func calcResult() {
	defer duration(time.Now(), "calcResult")
	c := make(chan Jchan)
	go permute2(all_places, c)
	for p := range c {
		if p.dist < min_journey_dist {
			min_journey_dist = p.dist
			// copy(min_journey_places, p.places)
			min_journey_places = append([]string{}, p.places...)
		}
		if p.dist > max_journey_dist {
			max_journey_dist = p.dist
			// copy(max_journey_places, p.places)
			max_journey_places = append([]string{}, p.places...)
		}
		// fmt.Println(dist, p)
	}
}

func main() {
	defer duration(time.Now(), "main")

	readInput()
	calcResult()

	fmt.Println("min", min_journey_dist, min_journey_places) // 207
	fmt.Println("max", max_journey_dist, max_journey_places) // 804
}
