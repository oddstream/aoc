// https://adventofcode.com/2017/day/17 Spinlock
package main

import (
	"fmt"
	"time"
)

func duration(invocation time.Time, name string) {
	fmt.Println(name, time.Since(invocation))
}

func onestar(steps int) int {
	var lock []int = []int{0}
	var pos int = 0
	for cycle := 0; cycle < 2017; cycle++ {
		new := (steps + pos) % len(lock)
		new += 1
		// thank you, https://ueokande.github.io/go-slice-tricks/
		lock = append(lock[:new], append([]int{cycle + 1}, lock[new:]...)...)
		pos = new
	}
	return lock[pos+1]
}

func twostarbrute(steps int) int {
	var lock []int = []int{0}
	var pos int = 0
	for cycle := 0; cycle < 50e6; cycle++ {
		new := (steps + pos) % len(lock)
		new += 1
		lock = append(lock[:new], append([]int{cycle + 1}, lock[new:]...)...)
		pos = new
	}
	for i := 0; i < len(lock); i++ {
		if lock[i] == 0 {
			return lock[i+1]
		}
	}
	return -1
}

// we only need the position of 0 in the lock
// so we don't need to maintain the []int lock
func twostarpy(steps int) int {
	var len int = 1 // initial state is []int{0}
	var pos int = 0
	var out int = 0
	for i := 0; i < 50e6; i++ {
		to_ins := i + 1 // the value to insert
		new := (pos + steps) % len
		if new == 0 {
			out = to_ins
		}
		new += 1 // where we 'insert' new value
		pos = new
		len += 1 // the lock has grown in length by one
	}
	return out
}

func main() {
	defer duration(time.Now(), "main")
	fmt.Println("test one", onestar(3))
	fmt.Println("     one", onestar(345))   // 866
	fmt.Println("     two", twostarpy(345)) // 11995607
}

/*
$ go run main.go
test one 638
     one 866
     two 11995607
main 267.63708ms
*/
