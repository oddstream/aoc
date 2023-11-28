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
	fields := strings.Fields(line)
	// the fields in brackets (hypernet) will be odd
	for i, field := range fields {
		if i%2 == 1 {
			if containsABBA(field) {
				return false
			}
		}
	}
	// supernet fields (not in brackets) will be even
	for i, field := range fields {
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

// returns list of corresponding BABs
func containsABAs(s string) ([]string, bool) {
	var babs []string
	for i := 0; i < len(s)-2; i++ {
		if s[i] != s[i+1] && s[i] == s[i+2] {
			bab := []byte{s[i+1], s[i], s[i+1]}
			babs = append(babs, string(bab))
		}
	}
	if len(babs) > 0 {
		return babs, true
	} else {
		return nil, false
	}
}

func supportsSSL(line string) bool {
	line = strings.ReplaceAll(line, "[", " ")
	line = strings.ReplaceAll(line, "]", " ")
	fields := strings.Fields(line)
	// search supernet strings for ABAs, collect their BABs
	var babs []string
	for i, field := range fields {
		if i%2 == 0 {
			if babs2, ok := containsABAs(field); ok {
				babs = append(babs, babs2...)
			}
		}
	}
	if len(babs) == 0 {
		return false
	}
	// search hypernet strings for any bab
	for i, field := range fields {
		if i%2 == 1 {
			for _, bab := range babs {
				if strings.Contains(field, bab) {
					return true
				}
			}
		}
	}
	return false
}

func partTwo() int {
	var result int
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		// fmt.Println(scanner.Text(), supportsSSL(scanner.Text()))
		if supportsSSL(scanner.Text()) {
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
	fmt.Println(partTwo()) // 260
}
