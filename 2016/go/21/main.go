// https://adventofcode.com/2016/day/21 Scrambled Letters and Hash

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

// The Go Programming Language 4.2
func reverse(s []byte) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func reverse2pos(s []byte, a, b int) {
	for i, j := a, b; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func swap2pos(s []byte, a, b int) {
	s[a], s[b] = s[b], s[a]
}

func swap2letters(s []byte, a, b byte) {
	var posa, posb int
	for i := 0; i < len(s); i++ {
		if s[i] == a {
			posa = i
		}
		if s[i] == b {
			posb = i
		}
	}
	if posa == posb {
		fmt.Println("error in swap2letters", a, b)
	}
	s[posa], s[posb] = s[posb], s[posa]
}

// The Go Programming Language 4.2
func rotleftn(s []byte, n int) {
	for n > 0 {
		reverse(s[:1])
		reverse(s[1:])
		reverse(s)
		n -= 1
	}
}

// The Go Programming Language 4.2
func rotrightn(s []byte, n int) {
	for n > 0 {
		reverse(s)
		reverse(s[:1])
		reverse(s[1:])
		n -= 1
	}
}

func rotpos(s []byte, b byte) {
	var i int
	for i = 0; i < len(s); i++ {
		if s[i] == b {
			break
		}
	}
	if i >= 4 {
		i += 2
	} else {
		i += 1
	}
	rotrightn(s, i)

}

func move2pos(s []byte, a, b int) {
	// remove byte at pos a
	tmp := s[a]
	s = append(s[:a], s[a+1:]...)
	// insert byte at pos b
	s = append(s[:b], append([]byte{tmp}, s[b:]...)...)
}

func loadInput() []string {
	var lines []string
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	return lines
}

func partOne(lines []string, password string) string {
	var pw []byte = []byte(password)
	var num1, num2 int
	var letter1, letter2 string

	for _, line := range lines {
		if n, _ := fmt.Sscanf(line, "move position %d to position %d", &num1, &num2); n == 2 {
			move2pos(pw, num1, num2)
		} else if n, _ := fmt.Sscanf(line, "reverse positions %d through %d", &num1, &num2); n == 2 {
			reverse2pos(pw, num1, num2)
		} else if n, _ := fmt.Sscanf(line, "swap position %d with position %d", &num1, &num2); n == 2 {
			swap2pos(pw, num1, num2)
		} else if n, _ := fmt.Sscanf(line, "swap letter %s with letter %s", &letter1, &letter2); n == 2 {
			swap2letters(pw, []byte(letter1)[0], []byte(letter2)[0])
		} else if n, _ := fmt.Sscanf(line, "rotate based on position of letter %s", &letter1); n == 1 {
			rotpos(pw, []byte(letter1)[0])
		} else if n, _ := fmt.Sscanf(line, "rotate left %d step", &num1); n == 1 {
			rotleftn(pw, num1)
		} else if n, _ := fmt.Sscanf(line, "rotate right %d step", &num1); n == 1 {
			rotrightn(pw, num1)
		} else {
			fmt.Println("cannot parse", line)
		}
		// println(string(pw))
	}
	return string(pw)
}

var perms []string

func generatePermutations(s []rune, left, right int) {
	if left == right {
		perms = append(perms, string(s))
	} else {
		for i := left; i <= right; i++ {
			s[left], s[i] = s[i], s[left]
			generatePermutations(s, left+1, right)
			s[left], s[i] = s[i], s[left]
		}
	}
}

func partTwo(lines []string, scrambled string) string {
	// 8! is only 40320, so...
	generatePermutations([]rune(scrambled), 0, len(scrambled)-1)
	// println("made", len(perms), "permutations")
	for _, s := range perms {
		if partOne(lines, s) == scrambled {
			return s
		}
	}
	return "not found"
}

func main() {
	defer duration(time.Now(), "main")
	var lines []string = loadInput()
	fmt.Println(partOne(lines, "abcdefgh")) // dgfaehcb
	fmt.Println(partTwo(lines, "fbgdceah")) // fdhgacbe

}

/*
$ go run main.go
dgfaehcb
made 40320 permutations
fdhgacbe
main 442.896512ms
*/
