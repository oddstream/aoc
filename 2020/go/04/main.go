// https://adventofcode.com/2020/day/4
package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"time"
	"unicode"
)

//go:embed test1.txt
var test1 string

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

func part1(in string, expected int) (result int) {
	defer duration(time.Now(), "part 1")

	var passport map[string]string = make(map[string]string)

	valid := func() bool {
		var required []string = []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
		for _, req := range required {
			if passport[req] == "" {
				return false
			}
		}
		return true
	}

	scanner := bufio.NewScanner(strings.NewReader(in))
	for scanner.Scan() {
		if scanner.Text() == "" {
			if valid() {
				result += 1
			}
			passport = make(map[string]string)
		} else {
			for _, token := range strings.Split(scanner.Text(), " ") {
				kv := strings.Split(token, ":")
				passport[kv[0]] = kv[1]
			}
		}
	}
	if len(passport) > 0 {
		if valid() {
			result += 1
		}
	}

	if expected != -1 {
		if result != expected {
			fmt.Println("ERROR: got", result, "expected", expected)
		} else {
			fmt.Println("RIGHT ANSWER:", result)
		}
	}
	return result
}

func part2(in string, expected int) (result int) {
	defer duration(time.Now(), "part 2")

	var passport map[string]string = make(map[string]string)
	hclrx := regexp.MustCompile("^#[0-9a-fA-F]+$")

	valid := func() bool {
		var required []string = []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
		for _, req := range required {
			var value string = passport[req]
			if value == "" {
				return false
			}
			switch req {
			case "byr":
				n := atoi(value)
				if n < 1920 || n > 2002 {
					return false
				}
			case "iyr":
				n := atoi(value)
				if n < 2010 || n > 2020 {
					return false
				}
			case "eyr":
				n := atoi(value)
				if n < 2020 || n > 2030 {
					return false
				}
			case "hgt":
				var digits, letters []rune
				for _, r := range value {
					if unicode.IsDigit(r) {
						digits = append(digits, r)
					} else if unicode.IsLetter(r) {
						letters = append(letters, r)
					}
				}
				if len(digits) == 0 {
					return false
				}
				n := atoi(string(digits))
				switch string(letters) {
				case "cm":
					if n < 150 || n > 193 {
						return false
					}
				case "in":
					if n < 59 || n > 76 {
						return false
					}
				default:
					return false
				}
			case "hcl":
				if len(value) != 7 {
					return false
				}
				// inefficient to compile a regexp in a loop
				if !hclrx.MatchString(value) {
					fmt.Println("hcl", value)
					return false
				}
			case "ecl":
				if len(value) != 3 {
					return false
				}
				switch value {
				case "amb", "blu", "brn", "gry", "grn", "hzl", "oth":
				default:
					return false
				}
			case "pid":
				if len(value) != 9 {
					return false
				}
				for _, r := range value {
					if !unicode.IsDigit(r) {
						return false
					}
				}
			case "":
				return false
			}
		}
		return true
	}

	scanner := bufio.NewScanner(strings.NewReader(in))
	for scanner.Scan() {
		if scanner.Text() == "" {
			if valid() {
				result += 1
			}
			passport = make(map[string]string)
		} else {
			for _, token := range strings.Split(scanner.Text(), " ") {
				kv := strings.Split(token, ":")
				passport[kv[0]] = kv[1]
			}
		}
	}
	if len(passport) > 0 {
		if valid() {
			result += 1
		}
	}

	if expected != -1 {
		if result != expected {
			fmt.Println("ERROR: got", result, "expected", expected)
		} else {
			fmt.Println("RIGHT ANSWER:", result)
		}
	}
	return result
}

func main() {
	defer duration(time.Now(), "main")

	// part1(test1, 2)
	part1(input, 208)
	part2(input, 167)

	{
		var memStats runtime.MemStats
		runtime.ReadMemStats(&memStats)
		fmt.Printf("Heap memory (in bytes): %d\n", memStats.HeapAlloc)
		fmt.Printf("Number of garbage collections: %d\n", memStats.NumGC)
	}
}

/*
$ go run main.go
RIGHT ANSWER: 208
part 1 389.953µs
RIGHT ANSWER: 167
part 2 504.059µs
Heap memory (in bytes): 622984
Number of garbage collections: 0
main 998.889µs
*/
