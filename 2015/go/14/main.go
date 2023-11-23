package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

//go:embed "input.txt"
var input string

// Duration of a func call
// Arguments to a defer statement are immediately evaluated and stored.
// The deferred function receives the pre-evaluated values when its invoked.
// usage: defer uDuration(time.Now(), "IntFactorial")
func Duration(invocation time.Time, name string) {
	fmt.Println(name, time.Since(invocation))
}

type Reindeer struct {
	// static
	name        string
	speedKms    int
	durationSec int
	restSec     int
	// dynamic
	resting      bool
	countdownSec int
	distanceKm   int
	points       int
}

var herd []*Reindeer

func parseInput() {
	re := regexp.MustCompile(`(\w+) can fly (\d+) km/s for (\d+) seconds, but then must rest for (\d+) seconds.`)
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		matches := re.FindAllStringSubmatchIndex(scanner.Text(), -1)
		// len(matches) == 1
		// len(matches[0]) == 10
		// 0:1 is whole match
		var r Reindeer = Reindeer{}
		r.name = scanner.Text()[matches[0][2]:matches[0][3]]
		speedStr := scanner.Text()[matches[0][4]:matches[0][5]]
		r.speedKms, _ = strconv.Atoi(speedStr)
		durationStr := scanner.Text()[matches[0][6]:matches[0][7]]
		r.durationSec, _ = strconv.Atoi(durationStr)
		restStr := scanner.Text()[matches[0][8]:matches[0][9]]
		r.restSec, _ = strconv.Atoi(restStr)
		herd = append(herd, &r)
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("scanner error: %v\n", err)
	}
}

func run(secs int) {
	for _, r := range herd {
		r.resting = false
		r.points = 0
		r.countdownSec = r.durationSec
	}
	for sec := 0; sec <= secs; sec++ {
		for _, r := range herd {
			if r.resting {
				r.countdownSec -= 1
				if r.countdownSec == 0 {
					r.resting = false
					r.countdownSec = r.durationSec
				}
			} else {
				r.distanceKm += r.speedKms
				r.countdownSec -= 1
				if r.countdownSec == 0 {
					r.resting = true
					r.countdownSec = r.restSec
				}
			}
		}
		var furthestKm = 0
		for _, r := range herd {
			if r.distanceKm > furthestKm {
				furthestKm = r.distanceKm
			}
		}
		for _, r := range herd {
			if r.distanceKm == furthestKm {
				r.points += 1
			}
		}
	}
}

func reportDistance() {
	fmt.Println("- Distance --------------------")
	sort.Slice(herd, func(a, b int) bool { return herd[a].distanceKm < herd[b].distanceKm })
	for _, r := range herd {
		fmt.Println(r.name, r.distanceKm)
	}
}

func reportPoints() {
	fmt.Println("- Points ----------------------")
	sort.Slice(herd, func(a, b int) bool { return herd[a].points < herd[b].points })
	for _, r := range herd {
		fmt.Println(r.name, r.points)
	}
}

func main() {
	defer Duration(time.Now(), "main")
	parseInput()
	run(2503)
	reportDistance() // part 1 Vixen   2660
	reportPoints()   // part 2 Blitzen 1256
}
