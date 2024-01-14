// https://adventofcode.com/2018/day/5 Alchemical Reduction
package main

import (
	_ "embed"
	"fmt"
	"math"
	"strings"
	"time"
)

//go:embed input.txt
var input string // remove trailing \n

func duration(invocation time.Time, name string) {
	fmt.Println(name, time.Since(invocation))
}

func removeCharacters(input string, characters string) string {
	filter := func(r rune) rune {
		if !strings.ContainsRune(characters, r) {
			return r
		}
		return -1
	}
	return strings.Map(filter, input)
}

func removeAdjacentPairs(s string) string {
	if len(s) < 2 {
		return s
	}
	var stack []byte
	for i := 0; i < len(s); i++ {
		if len(stack) > 0 && (stack[len(stack)-1] == s[i]+32 || stack[len(stack)-1] == s[i]-32) {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}
	return string(stack)
}

func partOne() int {
	defer duration(time.Now(), "part 1")
	fmt.Println("input len", len(input))

	return len(removeAdjacentPairs(input))
}

func partTwo() int {
	defer duration(time.Now(), "part 2")
	var shortest int = math.MaxInt64

	for a := 'A'; a <= 'Z'; a++ {
		s := removeCharacters(input, string([]rune{a, a + 32}))
		n := len(removeAdjacentPairs(s))
		if n < shortest {
			shortest = n
		}
	}
	return shortest
}

func main() {
	defer duration(time.Now(), "main")

	fmt.Println(partOne()) // 10978
	fmt.Println(partTwo()) // 4840
}

/*
$ go run main.go
input len 50000
part 1 291.802µs
10978
part 2 15.050609ms
4840
main 15.363411ms
*/
