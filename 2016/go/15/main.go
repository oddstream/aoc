// https://adventofcode.com/2016/day/15 Timing is Everything
package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"strings"
	"time"
)

//go:embed input.txt
var input string

func duration(invocation time.Time, name string) {
	fmt.Println(name, time.Since(invocation))
}

type Disc struct {
	positions, start int
}

func loadInput() []Disc {
	var discs []Disc
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		var num, positions, start int
		if n, err := fmt.Sscanf(scanner.Text(), "Disc #%d has %d positions; at time=0, it is at position %d.", &num, &positions, &start); n == 3 {
			discs = append(discs, Disc{positions: positions, start: start})
		} else {
			fmt.Println(n, err)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	return discs
}

func calc(discs []Disc) int {
	var time int
	for {
		var bad bool = false
		for i, disc := range discs {
			if (disc.start+time+i+1)%disc.positions != 0 {
				bad = true
				break
			}
		}
		if !bad {
			break
		}
		time += 1
	}
	return time
}

func partOne(expecting int) int {
	var result int
	var discs []Disc = loadInput()
	result = calc(discs)
	if expecting != 0 && result != expecting {
		fmt.Println("expecting", expecting)
	}
	return result
}

func partTwo(expecting int) int {
	var result int
	var discs []Disc = loadInput()
	discs = append(discs, Disc{positions: 11, start: 0})
	result = calc(discs)
	if expecting != 0 && result != expecting {
		fmt.Println("expecting", expecting)
	}
	return result
}

func main() {
	defer duration(time.Now(), "main")

	fmt.Println(partOne(121834))
	fmt.Println(partTwo(3208099))
}

/*
	go run main.go
	121834
	3208099
	main 10.786848ms
*/
