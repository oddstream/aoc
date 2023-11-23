package main

import (
	"bufio"
	_ "embed"
	"flag"
	"fmt"
	"regexp"
	"strings"
)

//go:embed "input.txt"
var input string

// var naughtyStrings []string = []string{"ab", "cd", "pq", "xy"}
var naughtyStringsRegex = regexp.MustCompile("ab|cd|pq|xy")

// contains at least 3 vowels (input is all lower case)
func countVowels(s string) int {
	var vs int
	for _, c := range s {
		switch c {
		case 'a', 'e', 'i', 'o', 'u':
			vs++
		}
	}
	return vs
}

// contains at least one repeated letter
func containsRepeatedChar(s string) bool {
	var prev rune
	for _, c := range s {
		if c == prev {
			return true
		}
		prev = c
	}
	return false
}

// xyxy, aabcdefgaa but not aaa
func test21(s string) bool {
	for i := 0; i < len(s)-1; i++ {
		for j := i + 2; j < len(s)-1; j++ {
			if s[j] == s[i] && s[j+1] == s[i+1] {
				return true
			}
		}
	}
	return false
}

// xyx, abcdefeghi (efe), aaa
func test22(s string) bool {
	var prev1, prev2 rune
	for _, c := range s {
		if c == prev2 {
			return true
		}
		prev2 = prev1
		prev1 = c
	}
	return false
}

// does not contain any of "ab", "cd", "pq", "xy"
func containsNaughtyStrings(s string) bool {
	return naughtyStringsRegex.Match([]byte(s))
}

func main() {
	var part int
	flag.IntVar(&part, "part", 2, "1 or 2")
	flag.Parse()

	if !test21("xyxy") {
		fmt.Println("fail 1")
	}
	if !test21("aabcdefgaa") {
		fmt.Println("fail 2")
	}
	if test21("aaa") {
		fmt.Println("fail 3")
	}
	if !test22("xyx") {
		fmt.Println("fail 4")
	}
	if !test22("abcdefeghi") {
		fmt.Println("fail 5")
	}
	if !test22("aaa") {
		fmt.Println("fail 6")
	}

	var nice int
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		if part == 1 {
			if containsNaughtyStrings(scanner.Text()) {
				continue
			}
			if !containsRepeatedChar(scanner.Text()) {
				continue
			}
			if countVowels(scanner.Text()) < 3 {
				continue
			}
		} else if part == 2 {
			if !test22(scanner.Text()) {
				continue
			}
			if !test21(scanner.Text()) {
				continue
			}
		}
		nice++
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("scanner error: %v\n", err)
	}
	fmt.Println("part", part, nice) // 1=255, 2=55
}
