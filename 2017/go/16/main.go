// https://adventofcode.com/2017/day/16 Permutation Promenade
package main

import (
	_ "embed"
	"fmt"
	"time"
)

//go:embed input.txt
var input []byte

func duration(invocation time.Time, name string) {
	fmt.Println(name, time.Since(invocation))
}

func display(programs []byte) {
	for i := 0; i < len(programs); i++ {
		fmt.Print(string(programs[i]))
	}
	fmt.Println()
}

func dance(cycles int, prog string) int {
	defer duration(time.Now(), fmt.Sprintf("dance %d:", cycles))

	var N int = len(prog)

	var programs []byte = make([]byte, N)
	for i := 0; i < N; i++ {
		programs[i] = prog[i]
	}

	var find func(byte) int = func(b byte) int {
		for i := 0; i < N; i++ {
			if programs[i] == b {
				return i
			}
		}
		return -1
	}

	// display()

	// use repeat cycle + 1
	for cycle := 0; cycle < cycles%(59+1); cycle++ {
		for i := 0; i < len(input); i++ {
			var getn func() int = func() int {
				var n int = -1
				if input[i] >= '0' && input[i] <= '9' {
					n = int(input[i] - '0')
					i += 1
					if input[i] >= '0' && input[i] <= '9' {
						n = n*10 + int(input[i]-'0')
						i += 1
					}
				}
				return n
			}
			switch input[i] {
			case ',':
			case 's':
				i += 1
				var n int = getn()
				programs = append(programs[N-n:], programs[:N-n]...)
			case 'x':
				i += 1
				var n1 int = getn()
				if input[i] == '/' {
					i += 1
				} else {
					fmt.Println("x expected / at pos", i)
					return -1
				}
				var n2 int = getn()
				programs[n2], programs[n1] = programs[n1], programs[n2]
			case 'p':
				i += 1
				var p1, p2 byte
				var p1pos, p2pos int
				if input[i] >= 'a' && input[i] <= 'p' {
					p1 = input[i]
					i += 1
					p1pos = find(p1)
				}
				if input[i] == '/' {
					i += 1
				} else {
					fmt.Println("p expected / at pos", i)
					return -1
				}
				if input[i] >= 'a' && input[i] <= 'p' {
					p2 = input[i]
					i += 1
					p2pos = find(p2)
				}
				programs[p2pos], programs[p1pos] = programs[p1pos], programs[p2pos]
			default:
				fmt.Println("unknown move", input[i], "at pos", i)
				return -1
			}
		}
		// uncomment this block to find the cycle length
		// if cycle != 0 && prog == string(programs) {
		// 	fmt.Println("repeat at cycle", cycle)	// 59
		// 	break
		// }
	}
	display(programs)
	return -1
}

func main() {
	defer duration(time.Now(), "main")

	dance(1, "abcdefghijklmnop")          // ociedpjbmfnkhlga
	dance(1000000000, "abcdefghijklmnop") // gnflbkojhicpmead
}

/*
0123
abcd		0,1,2,3
cadb		2,0,3,1
map start pos to end pos, represent as []int{2,0,3,1}
[]int of positions from start to put in result

          111111
0123456789012345
abcdefghijklmnop
ociedpjbmfnkhlga

14,2,8,4,3,15,9,1,12,5,13,10,7,11,6,0

nice idea, but won't work because of p move
*/

/*
$ go run main.go
ociedpjbmfnkhlga
dance 1: 318.576µs
gnflbkojhicpmead
dance 1000000000: 9.5002ms
main 9.83993ms
*/
