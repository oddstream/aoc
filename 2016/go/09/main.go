// https://adventofcode.com/2016/day/9
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

func decompress(str string, part int) int {
	// fmt.Println(str)
	var answer, idx int
	for idx < len(str) {
		switch str[idx] {
		case '(':
			var marker []byte
			for j := idx + 1; str[j] != ')'; j++ {
				marker = append(marker, str[j])
			}
			var len_marker = 1 + len(marker) + 1
			// chars - number of characters to take
			// times - number of times to repeat chars into output
			var chars, times int
			if _, err := fmt.Sscanf(string(marker), "%dx%d", &chars, &times); err != nil {
				fmt.Println(err)
				return -1
			}
			var substr string = str[idx+len_marker : idx+len_marker+chars]
			if part == 1 {
				answer += len(substr) * times
			} else {
				// fmt.Println(">>>", substr)
				inc := times * decompress(substr, part)
				// fmt.Println("<<<", substr, inc)
				answer += inc
			}
			idx += len_marker + chars
		default:
			answer += 1
			idx += 1
		}
	}
	// fmt.Println(str, ":=", answer)
	return answer
}

func partOne() int {
	var result int
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		result += decompress(scanner.Text(), 1)
		// fmt.Println()
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	return result
}

func partTwo() int {
	var result int
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		result += decompress(scanner.Text(), 2)
		// fmt.Println()
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	return result
}

func main() {
	defer duration(time.Now(), "main")

	fmt.Println(partOne()) // 120765
	fmt.Println(partTwo()) // 1658395076
}
