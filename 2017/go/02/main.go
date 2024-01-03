// https://adventofcode.com/2017/day/2 Corruption Checksum
package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

//go:embed input.txt
var input string

func duration(invocation time.Time, name string) {
	fmt.Println(name, time.Since(invocation))
}

func pairs(arr []int) [][2]int {
	n := len(arr)
	pairs := make([][2]int, 0, n*(n-1)/2)
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			pairs = append(pairs, [2]int{arr[i], arr[j]})
		}
	}
	return pairs
}

func main() {
	defer duration(time.Now(), "main")

	var spreadsheet [][]int
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		tokens := strings.Split(scanner.Text(), "\t")
		var row []int
		for _, t := range tokens {
			if n, err := strconv.Atoi(t); err == nil {
				row = append(row, n)
			} else {
				fmt.Println(err)
			}
		}
		spreadsheet = append(spreadsheet, row)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	var result int
	for _, row := range spreadsheet {
		var min int = math.MaxInt64
		var max int = 0
		for _, n := range row {
			if n > max {
				max = n
			}
			if n < min {
				min = n
			}
		}
		result += max - min
	}
	fmt.Println(result) // part 1: 30994

	result = 0
	for _, row := range spreadsheet {
		for _, p := range pairs(row) {
			if p[0] == p[1] {
				continue
			}
			if p[0] > p[1] {
				z := p[0] % p[1]
				if z == 0 {
					// fmt.Println(p[0], p[1])
					result += p[0] / p[1]
				}
			} else {
				z := p[1] % p[0]
				if z == 0 {
					// fmt.Println(p[1], p[0])
					result += p[1] / p[0]
				}
			}
		}
	}
	// fmt.Println(pairs([]int{0, 1, 2}))
	fmt.Println(result) // part 2: 233
}

/*
$ go run main.go
30994
233
main 89.552µs
*/
