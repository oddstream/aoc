// https://adventofcode.com/2018/day/2 Inventory Management System
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

func partOne() int {
	defer duration(time.Now(), "part 1")
	var twos, threes int
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		var freq map[rune]int = make(map[rune]int)
		for _, ch := range scanner.Text() {
			freq[ch]++
		}
		var has2, has3 bool
		for _, v := range freq {
			if v == 2 {
				has2 = true
			} else if v == 3 {
				has3 = true
			}
		}
		if has2 {
			twos++
		}
		if has3 {
			threes++
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	return twos * threes
}

func differ1(a, b string) (bool, []byte) {
	var cnt int // diff count
	var diffs []byte
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			cnt += 1
			diffs = append(diffs, a[i], b[i])
		}
	}
	return cnt == 1, diffs
}

func partTwo() string {
	defer duration(time.Now(), "part 2")
	scanner := bufio.NewScanner(strings.NewReader(input))
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	for i := 0; i < len(lines); i++ {
		for j := i + 1; j < len(lines); j++ {
			if ok, diffs := differ1(lines[i], lines[j]); ok {
				fmt.Println(string(diffs))
				fmt.Println(lines[i])
				fmt.Println(lines[j])
				var result []byte
				for k := 0; k < len(lines[i]); k++ {
					if lines[i][k] == lines[j][k] {
						result = append(result, lines[i][k])
					}
				}
				return string(result)
			}
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	return ""

	// di
	// mxhwoglxgeauywfkztndcvjqr
	// mxhwoglxgeauywfkztndcvjqr
}

func main() {
	defer duration(time.Now(), "main")

	fmt.Println(partOne()) // 6474
	fmt.Println(partTwo()) // mxhwoglxgeauywfkztndcvjqr
}

/*
$ go run main.go
part 1 479.218µs
6474
di
mxhwoglxgeauywfdkztndcvjqr
mxhwoglxgeauywfikztndcvjqr
part 2 2.354746ms
mxhwoglxgeauywfkztndcvjqr
main 2.854713ms
*/
