package main

import (
	"bufio"
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

func partOne() int {
	defer duration(time.Now(), "part 1")
	var result int
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		result += atoi(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	return result
}

func partTwo() int {
	defer duration(time.Now(), "part 2")
	var freq int
	var seen map[int]int = map[int]int{freq: 1}
	for {
		scanner := bufio.NewScanner(strings.NewReader(input))
		for scanner.Scan() {
			freq += atoi(scanner.Text())
			seen[freq]++
			if seen[freq] == 2 {
				return freq
			}
		}
		if err := scanner.Err(); err != nil {
			fmt.Println(err)
		}
	}
}

func main() {
	defer duration(time.Now(), "main")

	fmt.Println(partOne()) // 435
	fmt.Println(partTwo()) // 245
}

/*
$ go run main.go
part 1 43.002µs
435
part 2 19.691773ms
245
main 19.782349ms
*/
