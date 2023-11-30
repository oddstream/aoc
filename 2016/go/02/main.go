// https://adventofcode.com/2016/day/2
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

// 1 2 3
// 4 5 6
// 7 8 9

var dirmap1 map[uint][4]uint = map[uint][4]uint{
	//  U  D  L  R
	1: {1, 4, 1, 2},
	2: {2, 5, 1, 3},
	3: {3, 6, 2, 3},
	4: {1, 7, 4, 5},
	5: {2, 8, 4, 6},
	6: {3, 9, 5, 6},
	7: {4, 7, 7, 8},
	8: {5, 8, 7, 9},
	9: {6, 9, 8, 9},
}

func partOne() uint {
	var result uint = 0
	var key uint = 5
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		for _, r := range scanner.Text() {
			switch r {
			case 'U':
				key = dirmap1[key][0]
			case 'D':
				key = dirmap1[key][1]
			case 'L':
				key = dirmap1[key][2]
			case 'R':
				key = dirmap1[key][3]
			default:
				fmt.Println("unknown rune in input", r)
			}
		}
		result = result*10 + key
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("scanner error: %v\n", err)
	}
	return result
}

var dirmap2 map[rune][4]rune = map[rune][4]rune{
	//     U    D    L    R
	'1': {'1', '3', '1', '1'},
	'2': {'2', '6', '2', '3'},
	'3': {'1', '7', '2', '4'},
	'4': {'4', '8', '3', '4'},
	'5': {'5', '5', '5', '6'},
	'6': {'2', 'A', '5', '7'},
	'7': {'3', 'B', '6', '8'},
	'8': {'4', 'C', '7', '9'},
	'9': {'9', '9', '8', '9'},
	'A': {'6', 'A', 'A', 'B'},
	'B': {'7', 'D', 'A', 'C'},
	'C': {'8', 'C', 'B', 'C'},
	'D': {'B', 'D', 'D', 'D'},
}

func partTwo() string {
	var result string = ""
	var key rune = '5'
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		for _, r := range scanner.Text() {
			switch r {
			case 'U':
				key = dirmap2[key][0]
			case 'D':
				key = dirmap2[key][1]
			case 'L':
				key = dirmap2[key][2]
			case 'R':
				key = dirmap2[key][3]
			default:
				fmt.Println("unknown rune in input", r)
			}
		}
		result = result + string(key)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	return result
}

func main() {
	duration(time.Now(), "main")

	fmt.Println(partOne()) // 14894
	fmt.Println(partTwo()) // 26B96
}
