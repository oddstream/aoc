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

type RuleMap map[string][][]string

func duration(invocation time.Time, name string) {
	fmt.Println(name, time.Since(invocation))
}

// flipy: invert y
func flipy(in [][]string) [][]string {
	var out [][]string = make([][]string, len(in))
	for o, i := 0, len(in)-1; o < len(in); o, i = o+1, i-1 {
		out[o] = make([]string, len(in[i]))
		copy(out[o], in[i])
	}
	return out
}

// flipx: invert x
func flipx(in [][]string) [][]string {

	// adapted from canonical Go string reverse
	reverse := func(in []string) []string {
		var out []string = make([]string, len(in))
		copy(out, in)
		for i, j := 0, len(in)-1; i < j; i, j = i+1, j-1 {
			out[i], out[j] = out[j], out[i]
		}
		return out
	}

	n := len(in)
	var out [][]string = make([][]string, n)
	for i := 0; i < n; i++ {
		out[i] = reverse(in[i])
	}
	return out
}

// transpose: write columns of in as rows of out
// https://en.wikipedia.org/wiki/Transpose
func transpose(input [][]string) [][]string {
	if len(input) == 0 {
		return input
	}
	output := make([][]string, len(input[0]))
	for i := range output {
		output[i] = make([]string, len(input))
	}
	for i, row := range input {
		for j, val := range row {
			output[j][i] = val
		}
	}
	return output
}

func rotate(in [][]string) [][]string {
	out := transpose(in)
	return flipy(out)
}

func makeGridFromString(str string) [][]string {
	var grid [][]string
	for _, line := range strings.Split(str, "/") {
		grid = append(grid, strings.Split(line, ""))
	}
	return grid
}

func stringifyGrid(grid [][]string) string {
	var str string
	for _, row := range grid {
		for _, v := range row {
			str += v
		}
	}
	return str
}

func loadRules() RuleMap {

	var rm RuleMap = make(RuleMap)

	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		var left, right string
		if n, err := fmt.Sscanf(scanner.Text(), "%s => %s", &left, &right); n != 2 {
			fmt.Println(err)
			break
		}
		key := makeGridFromString(left)
		rule := makeGridFromString(right)
		// Oddly pleased with this ...
		for _, k := range [][][]string{key, flipx(key), flipy(key)} {
			rm[stringifyGrid(k)] = rule
			for i := 0; i < 3; i++ {
				k = rotate(k) // rotate 90, 180, 270
				rm[stringifyGrid(k)] = rule
			}
		}
		// ...although it would be neater with a directly-comparable map key.
	}

	return rm
}

// Sanity-saving logic from https://github.com/alexchao26/advent-of-code-go/tree/main/2017/day21
// State is [][]string.
// No math.sqrt()!
// All "/" have been removed.
// No separate split() and join().
// All grid elements saved as (pointers to immutable) strings.
// Was trying to do this by saving state as / separated strings, and got in a terrible mess.
// Injected logic for flip/transpose/rotate, and loading/building of rules map, from failed attempt.
func tick(grid [][]string, rules map[string][][]string) [][]string {
	var nextState [][]string

	// determine the size of break up the grid by. prioritize 2x2 grids
	var edgeSize int
	if len(grid)%2 == 0 {
		edgeSize = 2
	} else if len(grid)%3 == 0 {
		edgeSize = 3
	} else {
		panic("len(grid) is not evenly divisible by 2 or 3")
	}
	if len(grid) != len(grid[0]) {
		panic("grid is not square")
	}

	// iterate over like a sudoku grid, r and c iterate over the top left corner
	// of each sub-square
	for r := 0; r < len(grid); r += edgeSize {
		// a new row of sub-squares is being iterated over, add edgeSize+1 number
		// of empty slices onto the nextState grid
		for i := 0; i < edgeSize+1; i++ {
			nextState = append(nextState, []string{})
		}
		for c := 0; c < len(grid[0]); c += edgeSize {
			// generate the string to match a key in the rules map
			var strToMatch string
			for i := 0; i < edgeSize; i++ {
				for j := 0; j < edgeSize; j++ {
					// r+i and c+j point at coords within the original grid
					strToMatch += grid[r+i][c+j]
				}
			}

			// finding the result of the enhancement rule for the string to match
			resulting, ok := rules[strToMatch]
			if !ok {
				panic("No matching pattern found for " + strToMatch)
			}

			// append the values from the result onto the appropriate nextState row
			for i, vals := range resulting {
				nextStateIndex := len(nextState) - edgeSize - 1 + i
				nextState[nextStateIndex] = append(nextState[nextStateIndex], vals...)
			}
		}
	}

	return nextState
}

func run(rounds int) int {

	var rules RuleMap = loadRules()
	var state [][]string = [][]string{
		{".", "#", "."},
		{".", ".", "#"},
		{"#", "#", "#"},
	}

	for i := 0; i < rounds; i++ {
		state = tick(state, rules)
	}

	var count int
	for _, row := range state {
		for _, v := range row {
			if v == "#" {
				count++
			}
		}
	}
	return count
}

func main() {
	defer duration(time.Now(), "main")

	// fmt.Println(len(loadRules()))	// 528 (duplicates removed)
	// fmt.Println(flipx(makeGridFromString("123/456/789")))
	// fmt.Println(flipy(makeGridFromString("123/456/789")))
	// fmt.Println(transpose(makeGridFromString("123/456/789")))
	// fmt.Println(rotate(makeGridFromString("123/456/789")))

	fmt.Println(run(5))  // 162
	fmt.Println(run(18)) // 2264586
}

/*
$ go run main.go
162
2264586
main 272.182382ms
*/
