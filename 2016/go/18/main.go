// https://adventofcode.com/2016/day/18 Like a Rogue
package main

import (
	_ "embed"
	"fmt"
	"time"
)

var test []byte = []byte(".^^.^.^^^^")
var input []byte = []byte("^^^^......^...^..^....^^^.^^^.^.^^^^^^..^...^^...^^^.^^....^..^^^.^.^^...^.^...^^.^^^.^^^^.^^.^..^.^")

func duration(invocation time.Time, name string) {
	fmt.Println(name, time.Since(invocation))
}

var trans map[[3]byte]struct{} = map[[3]byte]struct{}{
	{94, 94, 46}: {}, // left and center are traps
	{46, 94, 94}: {}, // center and right are traps
	{94, 46, 46}: {}, // left only
	{46, 46, 94}: {}, // right only
}

func makerow(in []byte) []byte {
	// add leading and trailing imaginary safe tiles to the input
	var in2 []byte = []byte(".")
	in2 = append(in2, in...)
	in2 = append(in2, '.')
	var out []byte
	for i := 0; i < len(in2)-2; i++ {
		var three [3]byte = [3]byte{in2[i], in2[i+1], in2[i+2]}
		// was 3x faster using a switch instead of a map
		if _, ok := trans[three]; ok {
			out = append(out, 94)
		} else {
			out = append(out, 46)
		}
	}
	return out
}

func safetiles(rows [][]byte) int {
	var count int
	for _, row := range rows {
		for _, by := range row {
			if by == 46 {
				count += 1
			}
		}
	}
	return count
}

func display(rows [][]byte) {
	for _, row := range rows {
		for _, by := range row {
			fmt.Print(string(by))
		}
		fmt.Println()
	}
}

func main() {
	defer duration(time.Now(), "main")

	var rows [][]byte
	rows = append(rows, test)
	for i := 1; i < 10; i++ {
		rows = append(rows, makerow(rows[i-1]))
	}
	// display(rows)
	fmt.Println("test", safetiles(rows)) // 38

	rows = nil
	rows = append(rows, input)
	for i := 1; i < 40; i++ {
		rows = append(rows, makerow(rows[i-1]))
	}
	// display(rows)
	fmt.Println("part 1", safetiles(rows)) // 1978

	rows = nil
	rows = append(rows, input)
	for i := 1; i < 400000; i++ {
		rows = append(rows, makerow(rows[i-1]))
	}
	// display(rows)
	fmt.Println("part 2", safetiles(rows)) // 20003246
}

/*
$ go run main.go
test 38
part 1 1978
part 2 20003246
main 360.178642ms
*/
