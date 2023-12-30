// https://adventofcode.com/2016/day/19 An Elephant Named Joseph
package main

import (
	"fmt"
	"math/bits"
	"time"
)

var input int = 3017957

func duration(invocation time.Time, name string) {
	fmt.Println(name, time.Since(invocation))
}

func partOne() int {
	// Simply convert the number to binary,
	// drop the MSB, (won't it always be 1?)
	// append it as a new LSB,
	// and you're done.
	sigbits := 64 - bits.LeadingZeros64(uint64(input))
	hibit := (input >> (sigbits - 1)) & 1
	n := input & ((1 << (sigbits - 1)) - 1)
	n = (n << 1) | hibit
	return n
}

func partTwo() int {
	var w int = 1
	for i := 1; i < input; i++ {
		w = w%i + 1
		if w > (i+1)/2 {
			w++
		}
	}
	return w
}

func main() {
	defer duration(time.Now(), "main")

	fmt.Println(partOne())
	fmt.Println(partTwo())
}
