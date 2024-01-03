package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"sort"
	"strings"
	"time"
)

//go:embed input.txt
var input string

func duration(invocation time.Time, name string) {
	fmt.Println(name, time.Since(invocation))
}

func partOne() int {
	var result int
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		tokens := strings.Split(scanner.Text(), " ")
		sort.Slice(tokens, func(a, b int) bool { return tokens[a] < tokens[b] })
		var valid bool = true
		for i := 1; i < len(tokens); i++ {
			if tokens[i-1] == tokens[i] {
				valid = false
				break
			}
		}
		if valid {
			result += 1
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	return result
}

func partTwo() int {
	var sorted = func(s string) string {
		rs := []rune(s)
		sort.Slice(rs, func(i, j int) bool { return rs[i] < rs[j] })
		return string(rs)
	}
	var result int
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		words := strings.Split(scanner.Text(), " ")
		var tokens []string
		for _, word := range words {
			tokens = append(tokens, sorted(word))
		}
		sort.Slice(tokens, func(a, b int) bool { return tokens[a] < tokens[b] })
		var valid bool = true
		for i := 1; i < len(tokens); i++ {
			if tokens[i-1] == tokens[i] {
				valid = false
				break
			}
		}
		if valid {
			result += 1
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	return result
}

func main() {
	defer duration(time.Now(), "main")

	fmt.Println(partOne()) // 451
	fmt.Println(partTwo()) // 223
}

/*
$ go run main.go
451
223
main 1.775622ms
*/
