// https://adventofcode.com/2020/day/18
package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"time"
)

//go:embed input.txt
var input string

// Eric, this is too easy.
// Single character numeric constants,
// only two operators,
// and everything except () separated by single spaces.

func duration(invocation time.Time, name string) {
	fmt.Println(name, time.Since(invocation))
}

func report(expected, result int) {
	if expected != -1 {
		if result != expected {
			fmt.Println("ERROR: got", result, "expected", expected)
		} else {
			fmt.Println("RIGHT ANSWER:", result)
		}
	}
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

func infixToPostfix(infix string, precedence map[rune]int) string {
	var stack []rune
	var postfix strings.Builder
	for _, char := range infix {
		switch {
		case char == ' ':
		case char >= '0' && char <= '9':
			postfix.WriteRune(char)
		case char == '(':
			stack = append(stack, char)
		case char == ')':
			for stack[len(stack)-1] != '(' {
				postfix.WriteRune(stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			stack = stack[:len(stack)-1]
		case precedence[char] != 0:
			for len(stack) > 0 && precedence[char] <= precedence[stack[len(stack)-1]] {
				postfix.WriteRune(stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, char)
		default:
			fmt.Println("unknown char in infix", char)
		}
	}
	for len(stack) > 0 {
		postfix.WriteRune(stack[len(stack)-1])
		stack = stack[:len(stack)-1]
	}
	return postfix.String()
}

func evalPostfix(postfix string) int {
	var stack []int
	for _, char := range postfix {
		switch {
		case char >= '0' && char <= '9':
			stack = append(stack, atoi(string(char)))
		case char == '+':
			var a, b int
			a, stack = stack[len(stack)-1], stack[:len(stack)-1]
			b, stack = stack[len(stack)-1], stack[:len(stack)-1]
			stack = append(stack, a+b)
		case char == '*':
			var a, b int
			a, stack = stack[len(stack)-1], stack[:len(stack)-1]
			b, stack = stack[len(stack)-1], stack[:len(stack)-1]
			stack = append(stack, a*b)
		default:
			fmt.Println("unknown char in postfix", char)
		}
	}
	return stack[0]
}

func part1(in string, expected int) (result int) {
	defer duration(time.Now(), "part 1")

	var precedence map[rune]int = map[rune]int{'*': 2, '+': 2, '(': 1}
	scanner := bufio.NewScanner(strings.NewReader(in))
	for scanner.Scan() {
		var s string = infixToPostfix(scanner.Text(), precedence)
		result += evalPostfix(s)
	}
	report(expected, result)
	return result
}

func part2(in string, expected int) (result int) {
	defer duration(time.Now(), "part 1")

	var precedence map[rune]int = map[rune]int{'+': 3, '*': 2, '(': 1}
	scanner := bufio.NewScanner(strings.NewReader(in))
	for scanner.Scan() {
		var s string = infixToPostfix(scanner.Text(), precedence)
		result += evalPostfix(s)
	}
	report(expected, result)
	return result
}

func main() {
	defer duration(time.Now(), "main")

	part1(input, 654686398176)
	part2(input, 8952864356993)

	{
		var memStats runtime.MemStats
		runtime.ReadMemStats(&memStats)
		fmt.Printf("Heap memory (in bytes): %d\n", memStats.HeapAlloc)
		fmt.Printf("Number of garbage collections: %d\n", memStats.NumGC)
	}
}

/*
$ go run main.go
RIGHT ANSWER: 654686398176
part 1 393.677µs
RIGHT ANSWER: 8952864356993
part 1 367.173µs
Heap memory (in bytes): 305120
Number of garbage collections: 0
main 845.726µs
*/
