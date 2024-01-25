// https://adventofcode.com/2019/day/4 Secure Container
package main

import (
	"fmt"
	"time"
)

var inputMin, inputMax int = 137683, 596253

func duration(invocation time.Time, name string) {
	fmt.Println(name, time.Since(invocation))
}

func int2array(n int) [6]int {
	var a [6]int
	for i := 5; n > 0; i-- {
		a[i] = n % 10
		n /= 10
	}
	return a
}

func test(i int) (part1, part2 bool) {
	a := int2array(i) // eg 137683 := [1,3,7,6,8,3]
	count := 0
	last := -1
	for i := 0; i < 6; i++ {
		val := a[i]
		if val < last {
			return false, false
		} else if val == last {
			part1 = true
			count++
		} else {
			if count == 1 {
				part2 = true
			}
			count = 0
		}
		last = val
	}
	return part1, part2 || count == 1
}

func partOne(expected int) {
	defer duration(time.Now(), "part 1")

	var result int

	for i := inputMin; i <= inputMax; i++ {
		p1, _ := test(i)
		if p1 {
			result += 1
		}
	}

	if result != expected {
		fmt.Println("ERROR: got", result, "expected", expected)
	} else {
		fmt.Println("CORRECT:", result)
	}
}

func partTwo(expected int) {
	defer duration(time.Now(), "part 2")

	var result int

	for i := inputMin; i <= inputMax; i++ {
		_, p2 := test(i)
		if p2 {
			result += 1
		}
	}

	if result != expected {
		fmt.Println("ERROR: got", result, "expected", expected)
	} else {
		fmt.Println("CORRECT:", result)
	}
}

func main() {
	defer duration(time.Now(), "main")

	partOne(1864)
	partTwo(1258)
}

/*
$ seq 137683 596253 | grep -P '^(?=1*2*3*4*5*6*7*8*9*$).*(\d)\1' | wc -l
1864

$ seq 137683 596253 | grep -P '^(?=1*2*3*4*5*6*7*8*9*$).*(\d)(?<!(?=\1)..)\1(?!\1)' | wc -l
1258

$ go run main.go
*/
