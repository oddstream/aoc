// https://adventofcode.com/2019/day/6
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

func main() {
	defer duration(time.Now(), "main")

	// input is "PARENT(CHILD";
	// each child is unique (children only have one parent);
	// each parent can have >1 child
	var orbits map[string]string = make(map[string]string)
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()
		orbits[line[4:]] = line[0:3] // map[child]parent
	}
	// add up all the steps from every child, up through it's parents
	// until there are no more parents
	var count int
	for child := range orbits {
		for orbits[child] != "" {
			child = orbits[child]
			count += 1
		}
	}
	fmt.Println("part 1", count) // 253104

	var YOU, SAN []string
	for o := "YOU"; orbits[o] != ""; o = orbits[o] { // TFB .. COM len 310
		YOU = append(YOU, orbits[o])
	}
	for o := "SAN"; orbits[o] != ""; o = orbits[o] { // 7GP .. COM len 303
		SAN = append(SAN, orbits[o])
	}
	// step backwards from the ends of each list
	// (ie forewards from COM)
	// until they diverge
	var y = len(YOU) - 1
	var s = len(SAN) - 1
	for YOU[y] == SAN[s] && y > 0 && s > 0 {
		y--
		s--
	}
	// we stepped back 58 times
	// step forward one on each list (we went back one parent too far)
	y++
	s++
	fmt.Println("part 2", y+s) // 499
}

/*
$ go run main.go
part 1 253104
part 2 499
main 6.678675ms
*/
