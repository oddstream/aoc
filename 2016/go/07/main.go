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

// Autonomous Bridge Bypass Annotation
func containsABBA(s string) bool {
	for i := 0; i < len(s)-3; i++ {
		if s[i] != s[i+1] && s[i+1] == s[i+2] && s[i] == s[i+3] {
			return true
		}
	}
	return false
}

func supportsTLS(line string) bool {
	line = strings.ReplaceAll(line, "[", " ")
	line = strings.ReplaceAll(line, "]", " ")
	// the fields in brackets will be odd
	for i, field := range strings.Fields(line) {
		if i%2 == 1 {
			if containsABBA(field) {
				return false
			}
		}
	}
	for i, field := range strings.Fields(line) {
		if i%2 == 0 {
			if containsABBA(field) {
				return true
			}
		}
	}
	return false
}

func partOne() int {
	var result int
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		// fmt.Println(scanner.Text(), supportsTLS(scanner.Text()))
		if supportsTLS(scanner.Text()) {
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

	fmt.Println(partOne()) // 118
}
