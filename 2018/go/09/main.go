// https://adventofcode.com/2018/day/9 Marble Mania
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

// elegant algorithm from u/marcusandrews et al
// but using slices like this is too slow; part 2 runs (seemingly) forever
func playSlices(nplayers, lastMarble int) int {
	var scores map[int]int = map[int]int{}
	var marbles []int = []int{0} // contains starting marble

	push := func(x int) {
		marbles = append(marbles, x)
	}

	pushleft := func(x int) {
		marbles = append([]int{x}, marbles...)
	}

	pop := func() int {
		var x int
		x, marbles = marbles[len(marbles)-1], marbles[:len(marbles)-1]
		return x
	}

	popleft := func() int {
		var x int
		x, marbles = marbles[0], marbles[1:]
		return x
	}

	rotate := func(n int) {
		if n == 0 {
			return
		} else if n < 0 {
			for ; n < 0; n++ {
				push(popleft())
			}
		} else {
			for ; n > 0; n-- {
				pushleft(pop())
			}
		}
	}

	for marble := 1; marble < lastMarble+1; marble++ {
		if marble%23 == 0 {
			rotate(7)
			scores[marble%nplayers] += marble + pop()
			rotate(-1)
		} else {
			rotate(-1)
			push(marble)
		}
	}

	var maxScore int
	for _, score := range scores {
		maxScore = max(maxScore, score)
	}
	return maxScore
}

// using a linked list of (a lot of) nodes is much faster
// than reslicing
func playLinkedList(nplayers, nmarbles int) int {
	var scores map[int]int = map[int]int{}

	type Node struct {
		value       int
		left, right *Node
	}

	// create a circular doubly-linked list containing one 0 marble
	var current = &Node{value: 0}
	current.left, current.right = current, current

	for i := 1; i <= nmarbles; i++ {
		player := ((i - 1) % nplayers) + 1
		if i%23 == 0 {
			scores[player] += i
			for j := 0; j < 7; j++ {
				current = current.left
			}
			current.left.right, current.right.left = current.right, current.left
			scores[player] += current.value
			current = current.right
		} else {
			current = current.right
			new := Node{right: current.right, left: current, value: i}
			current.right.left = &new
			current.right = &new
			current = &new
		}
	}

	var maxScore int
	for _, score := range scores {
		maxScore = max(maxScore, score)
	}
	return maxScore
}

func play(multiplier int) int {
	defer duration(time.Now(), "part 1")

	var result int
	scanner := bufio.NewScanner(strings.NewReader(input))
	var players, last int
	for scanner.Scan() {
		if n, err := fmt.Sscanf(scanner.Text(), "%d players; last marble is worth %d points", &players, &last); n != 2 {
			fmt.Println(err)
			break
		}
		result = playLinkedList(players, last*multiplier)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	return result
}

func main() {
	defer duration(time.Now(), "main")

	fmt.Println(play(1))   // 422980
	fmt.Println(play(100)) // 552041936
}

/*
$ go run main.go
part 1 2.177718ms
422980
part 1 325.53282ms
3552041936
main 327.743835ms
*/
