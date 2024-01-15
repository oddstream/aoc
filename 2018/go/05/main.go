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

func removeAdjacentPairs(in string) string {
	if len(in) < 2 {
		return in
	}
	// var out []byte = make([]byte, 0, len(in))	slower!
	var out []byte
	for i := 0; i < len(in); i++ {
		l1 := len(out) - 1
		if len(out) > 0 && (out[l1] == in[i]+32 || out[l1] == in[i]-32) {
			out = out[:l1]
		} else {
			out = append(out, in[i])
		}
	}
	return string(out)
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
