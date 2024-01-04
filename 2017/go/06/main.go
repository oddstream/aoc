// https://adventofcode.com/2017/day/6 Memory Reallocation
package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
	"time"
)

//go:embed input.txt
var input string

func duration(invocation time.Time, name string) {
	fmt.Println(name, time.Since(invocation))
}

// or we could use K&R p61
func atoi(s string) int {
	if n, err := strconv.Atoi(s); err == nil {
		return n
	} else {
		fmt.Println(err)
	}
	return 0
}

func solve(part int) int {
	var banks []int

	var key = func() string {
		var b []byte = make([]byte, len(banks))
		for i := 0; i < len(banks); i++ {
			b[i] = byte(banks[i])
		}
		return string(b)
	}

	var most = func() int {
		var max, imax int
		for i := 0; i < len(banks); i++ {
			if banks[i] > max {
				max = banks[i]
				imax = i
			}
		}
		return imax
	}

	for _, str := range strings.Split(input, "\t") {
		banks = append(banks, atoi(str))
	}

	var redistributions int
	var seen map[string]int = make(map[string]int)
	seen[key()] = 0
	for {
		var i int = most()
		var blocks int = banks[i]
		banks[i] = 0
		for blocks > 0 {
			i += 1
			if i >= len(banks) {
				i = 0
			}
			banks[i] += 1
			blocks -= 1
		}

		redistributions += 1

		var k = key()
		if cycles, ok := seen[k]; ok {
			if part == 1 {
				return redistributions
			} else if part == 2 {
				return redistributions - cycles
			}
		} else {
			seen[k] = redistributions
		}
	}
}

func main() {
	defer duration(time.Now(), "main")

	fmt.Println(solve(1)) // 12841
	fmt.Println(solve(2)) // 8038
}

/*
$ go run main.go
12841
8038
main 6.229516ms
*/
