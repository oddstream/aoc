// https://adventofcode.com/2021/day/3
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

func btoi(s string) int {
	var i64 int64
	var err error
	if i64, err = strconv.ParseInt(s, 2, 64); err != nil {
		fmt.Println(s, err)
	}
	return int(i64)
}

func readInput(in string) []string {
	var out []string
	scanner := bufio.NewScanner(strings.NewReader(in))
	for scanner.Scan() {
		out = append(out, scanner.Text())
	}
	return out
}

func mostCommonValues(report []string, x int) (ones int, zeros int) {
	for y := 0; y < len(report); y++ {
		if report[y][x] == '1' {
			ones += 1
		}
	}
	return ones, len(report) - ones
}

func part1(in string, expected int) (result int) {
	defer duration(time.Now(), "part 1")

	var diagnosticReport []string = readInput(in)

	var gamma, epsilon string
	for x := 0; x < len(diagnosticReport[0]); x++ {
		var ones, zeros int = mostCommonValues(diagnosticReport, x)
		if ones > zeros {
			gamma += "1"
			epsilon += "0"
		} else {
			gamma += "0"
			epsilon += "1"
		}
	}

	var g int = btoi(gamma)
	var e int = btoi(epsilon)
	result = g * e
	report(expected, result)
	return result
}

func collectColumn(report []string, i int, digit byte) []string {
	var out []string
	for y := 0; y < len(report); y++ {
		if report[y][i] == digit {
			out = append(out, report[y])
		}
	}
	return out
}

func part2(in string, expected int) (result int) {
	defer duration(time.Now(), "part 2")

	var oxygenGeneratorRating, CO2ScrubberRating int

	var src []string = readInput(in)
	for x := 0; x < len(src[0]); x++ {
		var ones, zeros int = mostCommonValues(src, x)
		var o []string = collectColumn(src, x, '1') // numbers with "1" in column x
		var z []string = collectColumn(src, x, '0') // numbers with "0" in column x
		var dst []string
		if ones > zeros || ones == zeros {
			dst = append(dst, o...)
		} else {
			dst = append(dst, z...)
		}
		if len(dst) == 0 {
			fmt.Println(x, "dst is empty")
			break
		}
		if len(dst) == 1 {
			var ogr int = btoi(dst[0])
			// fmt.Println("ogr", dst[0], ogr)
			oxygenGeneratorRating = int(ogr)
			break
		}
		src = dst
	}

	src = readInput(in)
	for x := 0; x < len(src[0]); x++ {
		var ones, zeros int = mostCommonValues(src, x)
		var o []string = collectColumn(src, x, '1') // numbers with "1" in column x
		var z []string = collectColumn(src, x, '0') // numbers with "0" in column x
		var dst []string
		if ones > zeros || ones == zeros {
			dst = append(dst, z...)
		} else {
			dst = append(dst, o...)
		}
		if len(dst) == 0 {
			fmt.Println(x, "dst is empty")
			break
		}
		if len(dst) == 1 {
			var csr int = btoi(dst[0])
			// fmt.Println("csr", dst[0], csr)
			CO2ScrubberRating = int(csr)
			break
		}
		src = dst
	}

	result = oxygenGeneratorRating * CO2ScrubberRating
	report(expected, result)
	return result
}

func main() {
	defer duration(time.Now(), "main")

	// part1(test1, 198)
	part1(input, 741950)
	// part2(test1, 230)
	part2(input, 903810)

	{
		var memStats runtime.MemStats
		runtime.ReadMemStats(&memStats)
		fmt.Printf("Heap memory (in bytes): %d\n", memStats.HeapAlloc)
		fmt.Printf("Number of garbage collections: %d\n", memStats.NumGC)
	}
}

/*
$ go run main.go
RIGHT ANSWER: 741950
part 1 90.16µs
RIGHT ANSWER: 903810
part 2 236.858µs
Heap memory (in bytes): 531128
Number of garbage collections: 0
main 455.442µs
*/
