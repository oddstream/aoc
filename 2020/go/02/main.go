// https://adventofcode.com/2020/day2
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

func chars(s string, r rune) int {
	var n int
	for _, c := range s {
		if c == r {
			n += 1
		}
	}
	return n
}

func part1(in string, expected int) (result int) {
	defer duration(time.Now(), "part 1")

	re := regexp.MustCompile("([[:digit:]]+)-([[:digit:]]+) ([[:alpha:]]): ([[:alpha:]]+)")
	scanner := bufio.NewScanner(strings.NewReader(in))
	for scanner.Scan() {
		var submatches []string = re.FindAllStringSubmatch(scanner.Text(), -1)[0]
		var n1, n2 int = atoi(submatches[1]), atoi(submatches[2])
		var ch, str string = submatches[3], submatches[4]
		var n int = chars(str, rune(ch[0]))
		if n >= n1 && n <= n2 {
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

	re := regexp.MustCompile("([[:digit:]]+)-([[:digit:]]+) ([[:alpha:]]): ([[:alpha:]]+)")
	scanner := bufio.NewScanner(strings.NewReader(in))
	for scanner.Scan() {
		var submatches []string = re.FindAllStringSubmatch(scanner.Text(), -1)[0]
		var n1, n2 int = atoi(submatches[1]), atoi(submatches[2])
		var ch rune = rune(submatches[3][0])
		var str []rune = []rune(submatches[4])
		if (str[n1-1] == ch && str[n2-1] != ch) || (str[n1-1] != ch && str[n2-1] == ch) {
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

	part1(test1, 2)
	part1(input, 500)
	part2(test1, 1)
	part2(input, 313)

	{
		var memStats runtime.MemStats
		runtime.ReadMemStats(&memStats)
		fmt.Printf("Heap memory (in bytes): %d\n", memStats.HeapAlloc)
		fmt.Printf("Number of garbage collections: %d\n", memStats.NumGC)
	}
}

/*
$ go run main.go
RIGHT ANSWER: 2
part 1 53.904µs
RIGHT ANSWER: 500
part 1 708.66µs
RIGHT ANSWER: 1
part 2 9.708µs
RIGHT ANSWER: 313
part 2 718.133µs
Heap memory (in bytes): 1071496
Number of garbage collections: 0
main 1.604045ms
*/
