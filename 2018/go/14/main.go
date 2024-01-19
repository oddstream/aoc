// https://adventofcode.com/2018/day/14 Chocolate Charts
package main

import (
	"fmt"
	"time"
)

func duration(invocation time.Time, name string) {
	fmt.Println(name, time.Since(invocation))
}

func display(scores []int, e1, e2 int) {
	for i := 0; i < len(scores); i++ {
		if i == e1 {
			fmt.Printf("(%d)", scores[i])
		} else if i == e2 {
			fmt.Printf("[%d]", scores[i])
		} else {
			fmt.Printf(" %d ", scores[i])
		}
	}
	fmt.Println()
}

func last10(scores []int, n int) int {
	var out int
	for i := 0; i < 10; i++ {
		out *= 10
		out += scores[n+i]
	}
	return out
}

func indexOf(scores []int, n []int) int {
	for i := range scores {
		if len(scores)-i < len(n) {
			return -1
		}
		for j := range n {
			if scores[i+j] != n[j] {
				break
			} else if j == len(n)-1 {
				return i
			}
		}
	}
	return -1
}

func intToSlice(n int) []int {
	var digits []int
	for n > 0 {
		digits = append([]int{n % 10}, digits...)
		n /= 10
	}
	return digits
}

func partOne(recipes int) int {
	defer duration(time.Now(), "part 1")

	var scores []int = []int{3, 7} // always 0 .. 9, could be []byte
	var e1, e2 int = 0, 1
	for len(scores) < recipes+10 {
		prev := scores[e1] + scores[e2]
		if prev < 10 {
			scores = append(scores, prev)
		} else {
			scores = append(scores, prev/10, prev%10)
		}
		e1 = (e1 + 1 + scores[e1]) % len(scores)
		e2 = (e2 + 1 + scores[e2]) % len(scores)
	}
	return last10(scores, recipes)
}

func partTwo(recipes int) int {
	defer duration(time.Now(), "part 2")

	var scores []int = []int{3, 7} // always 0 .. 9, could be []byte
	var e1, e2 int = 0, 1
	// loops := 0
	for len(scores) < 50e6 {
		prev := scores[e1] + scores[e2]
		if prev < 10 {
			scores = append(scores, prev)
		} else {
			scores = append(scores, prev/10, prev%10)
		}
		e1 = (e1 + 1 + scores[e1]) % len(scores)
		e2 = (e2 + 1 + scores[e2]) % len(scores)
		// loops++
	}
	// fmt.Println(loops)	// 38353114
	var sl []int = intToSlice(recipes)
	return indexOf(scores, sl)
}

func main() {
	defer duration(time.Now(), "main")

	fmt.Println("part 1", partOne(637061)) // 3138510102
	fmt.Println("part 2", partTwo(637061)) // 20179081
}

/*
$ go run main.go
part 1 14.152028ms
part 1 3138510102
part 2 1.055154786s
part 2 20179081
main 1.069372657s
*/
