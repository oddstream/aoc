// https://adventofcode.com/2016/day/22 Computing
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

type Node struct {
	x                      int // 0 .. 37
	y                      int // 0 .. 27
	size, used, avail, use int
}

func duration(invocation time.Time, name string) {
	fmt.Println(name, time.Since(invocation))
}

func loadInput() []Node {
	var nodes []Node
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		var line string = scanner.Text()
		if strings.HasPrefix(line, "root") {
			continue
		}
		if strings.HasPrefix(line, "Filesystem") {
			continue
		}
		var n Node = Node{}
		// Sscanf magically swallows multiple spaces
		if n, err := fmt.Sscanf(scanner.Text(), "/dev/grid/node-x%d-y%d %dT %dT %dT %d%%",
			&n.x, &n.y, &n.size, &n.used, &n.avail, &n.use); n != 6 || err != nil {
			fmt.Println(err, line)
			break
		}
		nodes = append(nodes, n)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	return nodes
}

func partOne() int {
	var result int
	var nodes = loadInput()
	for _, a := range nodes {
		if a.used == 0 {
			continue
		}
		for _, b := range nodes {
			if a == b {
				continue
			}
			if a.used <= b.avail {
				result += 1
			}
		}
	}
	return result
}

func partTwo() int {
	var result int = -1
	/*
		could have used an awk one-liner, but instead borrowed this:
		tail -n+3 input.txt # chop off first 2 line
		tail -n+3 input.txt | sed 's/ \+/ /g' # remove duplicate spaces
		tail -n+3 input.txt | sed 's/ \+/ /g' | cut -d' ' -f3 # extract Used field

		$ tail -n+3 input.txt | sed 's/ \+/ /g' | cut -d' ' -f3 | sort -n | uniq -c
			         1 0T
			       112 64T
			        93 65T
			        87 66T
			       102 67T
			       120 68T
			       100 69T
			       102 70T
			       116 71T
			        97 72T
			       105 73T
			         2 490T
			         1 491T
			         1 492T
			         3 493T
			         7 494T
			         1 495T
			         4 496T
			         3 497T
			         5 498T
			         2 499T

		which shows one empty and an some very full nodes
		which when displayed look liek this:

		.....................................G
		......................................
		......................................
		......................................
		......................................
		......................................
		......................................
		......................................
		......................................
		......................................
		......................................
		......................................
		......................................
		......................................
		......................................
		......................................
		......................................
		......................................
		......................................
		......................................
		......................................
		......................................
		......................................
		......................................
		.........#############################
		......................................
		.................................._...
		......................................

		The operation we have available is moving the _ by one step,
		swapping it with its target location, and avoiding any #

		80 to move _ next to (left of) G
		36*5 to move G to 0,0
		+1
		= 261
	*/
	var nodes = loadInput()
	for y := 0; y < 28; y++ {
		for x := 0; x < 38; x++ {
			if y == 0 && x == 37 {
				fmt.Print("G")
			} else {
				for _, n := range nodes {
					if x == n.x && y == n.y {
						if n.used == 0 {
							fmt.Print("_")
						} else if n.used > 400 {
							fmt.Print("#")
						} else {
							fmt.Print(".")
						}
						break
					}
				}
			}
		}
		fmt.Println()
	}
	return result
}

func main() {
	defer duration(time.Now(), "main")

	fmt.Println("part 1", partOne()) // 1034
	fmt.Println("part 2", partTwo()) // 261
}
