// https://adventofcode.com/2020/day/14
package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"time"
)

//go:embed test1.txt
var test1 string

//go:embed test2.txt
var test2 string

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

func report(expected, result int) {
	if expected != -1 {
		if result != expected {
			fmt.Println("ERROR: got", result, "expected", expected)
		} else {
			fmt.Println("RIGHT ANSWER:", result)
		}
	}
}

func part1(in string, expected int) (result int) {
	defer duration(time.Now(), "part 1")

	var mask string
	var mem map[int]int64 = make(map[int]int64)

	scanner := bufio.NewScanner(strings.NewReader(in))
	for scanner.Scan() {
		tokens := strings.Split(scanner.Text(), " = ")
		if tokens[0] == "mask" {
			mask = tokens[1]
			if len(mask) != 36 {
				panic("len mask is not 36")
			}
		} else {
			var addr int
			if n, err := fmt.Sscanf(tokens[0], "mem[%d]", &addr); n != 1 {
				fmt.Println(err, tokens[0])
				break
			}
			var val string = fmt.Sprintf("%036b", atoi(tokens[1]))
			// fmt.Println(tokens[0], tokens[1], len(mask), mask, addr, len(val), val)

			var res []byte = []byte(val)
			for i := 0; i < 36; i++ {
				if mask[i] != 'X' {
					res[i] = mask[i]
				}
			}
			mem[addr], _ = strconv.ParseInt(string(res), 2, 64)
		}
	}

	for _, v := range mem {
		result += int(v)
	}

	report(expected, result)
	return result
}

func generatePermutations(s []byte) [][]byte {
	var permutations [][]byte
	var helper func([]byte, int)
	helper = func(s []byte, i int) {
		if i == len(s) {
			permutations = append(permutations, s)
			return
		}
		if s[i] == 'X' {
			helper(bytes.Replace(s, []byte{'X'}, []byte{'0'}, 1), i+1)
			helper(bytes.Replace(s, []byte{'X'}, []byte{'1'}, 1), i+1)
		} else {
			helper(s, i+1)
		}
	}
	helper(s, 0)
	return permutations
}

func part2(in string, expected int) (result int) {
	defer duration(time.Now(), "part 2")

	var mask string
	var mem map[int]int64 = make(map[int]int64)

	scanner := bufio.NewScanner(strings.NewReader(in))
	for scanner.Scan() {
		tokens := strings.Split(scanner.Text(), " = ")
		if tokens[0] == "mask" {
			mask = tokens[1]
			if len(mask) != 36 {
				panic("len mask is not 36")
			}
		} else {
			var addr int
			if n, err := fmt.Sscanf(tokens[0], "mem[%d]", &addr); n != 1 {
				fmt.Println(err, tokens[0])
				break
			}
			var val int = atoi(tokens[1])

			// fmt.Println(tokens[0], tokens[1], len(mask), mask, addr, len(val), val)

			var res []byte = []byte(fmt.Sprintf("%036b", addr))
			for i := 0; i < 36; i++ {
				switch mask[i] {
				case '0':
					// addr bit is unchanged
				case '1':
					res[i] = '1'
				case 'X':
					res[i] = 'X'
				}
			}
			for _, s := range generatePermutations(res) {
				a, _ := strconv.ParseInt(string(s), 2, 64)
				mem[int(a)] = int64(val)
			}
		}
	}

	for _, v := range mem {
		result += int(v)
	}

	report(expected, result)
	return result
}

func main() {
	defer duration(time.Now(), "main")

	// part1(test1, 165)
	part1(input, 5875750429995)
	// part2(test2, 208)
	part2(input, 5272149590143)

	{
		var memStats runtime.MemStats
		runtime.ReadMemStats(&memStats)
		fmt.Printf("Heap memory (in bytes): %d\n", memStats.HeapAlloc)
		fmt.Printf("Number of garbage collections: %d\n", memStats.NumGC)
	}
}

/*
$ go run main.go
RIGHT ANSWER: 5875750429995
part 1 645.752µs
RIGHT ANSWER: 5272149590143
part 2 24.200053ms
Heap memory (in bytes): 2862440
Number of garbage collections: 6
main 24.876215ms
*/
