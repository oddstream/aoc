// https://adventofcode.com/2021/day/4
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

//go:embed test1.txt
var test1 string

//go:embed input.txt
var input string

type Board [][]int

func (b Board) bingo(dn map[int]struct{}) bool {
	for y := 0; y < 5; y++ {
		var row bool = true
		for x := 0; x < 5; x++ {
			if _, ok := dn[b[y][x]]; !ok {
				row = false
			}
		}
		if row {
			return true
		}
	}
	for x := 0; x < 5; x++ {
		var col bool = true
		for y := 0; y < 5; y++ {
			if _, ok := dn[b[y][x]]; !ok {
				col = false
			}
		}
		if col {
			return true
		}
	}
	return false
}

func (b Board) unmarked(dn map[int]struct{}) int {
	var n int
	for y := 0; y < 5; y++ {
		for x := 0; x < 5; x++ {
			var val int = b[y][x]
			if _, ok := dn[val]; !ok {
				n += val
			}
		}
	}
	return n
}

func (b Board) display(dn map[int]struct{}) {
	for y := 0; y < 5; y++ {
		for x := 0; x < 5; x++ {
			if _, ok := dn[b[y][x]]; !ok {
				fmt.Printf(" %2d ", b[y][x])
			} else {
				fmt.Printf("[%2d]", b[y][x])
			}
		}
		fmt.Println()
	}
}

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

func readInput(in string) (numbers []int, boards []Board) {
	// https://golangdocs.com/split-string-in-golang

	scanner := bufio.NewScanner(strings.NewReader(in))
	scanner.Scan() // get first line, the numbers
	for _, s := range strings.Split(scanner.Text(), ",") {
		numbers = append(numbers, atoi(s))
	}

	var board Board

	addBoard := func() {
		if len(board) > 0 {
			boards = append(boards, board)
			board = nil
		}
	}

	for scanner.Scan() {
		var line string = scanner.Text()
		if line == "" {
			addBoard()
		} else {
			var arr []int
			for _, s := range strings.Fields(line) {
				arr = append(arr, atoi(s))
			}
			board = append(board, arr)
		}
	}
	addBoard()

	return numbers, boards
}

func part1(in string, expected int) (result int) {
	defer duration(time.Now(), "part 1")

	numbers, boards := readInput(in)

	var drawnNumbers map[int]struct{} = make(map[int]struct{})

	// for _, board := range boards {
	// 	board.display(drawnNumbers)
	// 	fmt.Println()
	// }

	for _, drawn := range numbers {
		drawnNumbers[drawn] = struct{}{}
		for _, board := range boards {
			if board.bingo(drawnNumbers) {
				// board.display(drawnNumbers)
				result = board.unmarked(drawnNumbers) * drawn
				// fmt.Println("bingo", n)
				goto exit
			}
		}
	}

exit:
	report(expected, result)
	return result
}

func part2(in string, expected int) (result int) {
	defer duration(time.Now(), "part 2")

	numbers, boards := readInput(in)

	var drawnNumbers map[int]struct{} = make(map[int]struct{})

	var completed map[int]struct{} = make(map[int]struct{})
	for _, drawn := range numbers {
		drawnNumbers[drawn] = struct{}{}
		for b, board := range boards {
			if board.bingo(drawnNumbers) {
				completed[b] = struct{}{}
			}
			if len(completed) == len(boards) {
				result = board.unmarked(drawnNumbers) * drawn
				goto exit
			}
		}
	}

exit:
	report(expected, result)
	return result
}

func main() {
	defer duration(time.Now(), "main")

	// part1(test1, 4512)
	part1(input, 38594)
	// part2(test1, 1924)
	part2(input, 21184)

	{
		var memStats runtime.MemStats
		runtime.ReadMemStats(&memStats)
		fmt.Printf("Heap memory (in bytes): %d\n", memStats.HeapAlloc)
		fmt.Printf("Number of garbage collections: %d\n", memStats.NumGC)
	}
}

/*
$ go run main.go
RIGHT ANSWER: 38594
part 1 1.366479ms
RIGHT ANSWER: 21184
part 2 5.346775ms
Heap memory (in bytes): 460256
Number of garbage collections: 0
main 6.815838ms
*/
