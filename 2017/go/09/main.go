// https://adventofcode.com/2017/day/9 Stream Processing
package main

import (
	_ "embed"
	"fmt"
	"time"
)

//go:embed input.txt
var input string

func duration(invocation time.Time, name string) {
	fmt.Println(name, time.Since(invocation))
}

func scan(s string) (int, int) {
	var garbage bool = false
	var score int
	var depth = 1
	var gcount int
	for i := 0; i < len(s); i++ {
		ch := s[i]
		if ch == '!' {
			i++
		} else if garbage && ch != '>' {
			gcount += 1
		} else if ch == '<' {
			garbage = true
		} else if ch == '>' {
			garbage = false
		} else if ch == '{' {
			score += depth
			depth += 1
		} else if ch == '}' {
			depth -= 1
		}
	}
	return score, gcount
}

func main() {
	defer duration(time.Now(), "main")

	fmt.Println(scan(input)) // 12505, 6671
}

/*
$ go run main.go
12505 6671
main 99.09µs
*/
