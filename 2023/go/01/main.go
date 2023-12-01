// https://adventofcode.com/2023/day/1
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

var numbers []string = []string{
	"zero", // dummy to make "one" == 1
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
}

func getFirstLast1(line string) int {
	var first, last byte
	for i := 0; i < len(line); i++ {
		ch := line[i]
		if ch >= '0' && ch <= '9' {
			first = ch
			break
		}
	}
	for i := len(line) - 1; i >= 0; i-- {
		ch := line[i]
		if ch >= '0' && ch <= '9' {
			last = ch
			break
		}
	}
	if first == 0 || last == 0 {
		fmt.Println("ERROR NO DIGIT(S) FOUND", line, first, last)
		return 0
	}
	return int(first-'0')*10 + int(last-'0')
}

func getFirstLast2(line string) int {
	var first, last byte
	var tmp string = line // use a tmp string in case first and last collide/overlap (they do)
	for len(tmp) > 0 {
		ch := tmp[0]
		if ch >= '0' && ch <= '9' {
			first = ch
			goto findSuffix
		}
		for i, number := range numbers {
			if strings.HasPrefix(tmp, number) {
				first = '0' + byte(i)
				goto findSuffix
			}
		}
		tmp = tmp[1:] // chop off the first rune and try again
	}
findSuffix:
	for len(line) > 0 {
		ch := line[len(line)-1]
		if ch >= '0' && ch <= '9' {
			last = ch
			goto exit
		}
		for i, number := range numbers {
			if strings.HasSuffix(line, number) {
				last = '0' + byte(i)
				goto exit
			}
		}
		line = line[:len(line)-1] // chop off the last rune and try again
	}

exit:
	if first == 0 || last == 0 {
		fmt.Println("ERROR NO DIGIT(S) FOUND", line, first, last)
		return 0
	}
	return int(first-'0')*10 + int(last-'0')
}

func partOne() int {
	var result int
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		num := getFirstLast1(scanner.Text())
		result = result + num
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
		num := getFirstLast2(scanner.Text())
		result = result + num
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	return result
}

func main() {
	defer duration(time.Now(), "main")

	fmt.Println("part one", partOne()) // 55208
	fmt.Println("part two", partTwo()) // 54578
}
