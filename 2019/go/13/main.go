// https://adventofcode.com/2019/day/13
package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
	"time"
)

//go:embed input.txt
var input string

const (
	EMPTY  int = 0
	WALL   int = 1
	BLOCK  int = 2
	PADDLE int = 3
	BALL   int = 4
)

type Point struct {
	x, y int
}

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

func play1(program []int) int {
	var tiles map[Point]int = make(map[Point]int)

	in := func() int {
		fmt.Println("INPUT") // part 1 has no input
		return 0
	}

	var xytile []int
	out := func(val int) {
		xytile = append(xytile, val)
		if len(xytile) == 3 {
			tiles[Point{x: xytile[0], y: xytile[1]}] = xytile[2]
			xytile = nil
		}
	}

	intcode(program, in, out)
	/*
		in := make(chan int)
		go func() {
			in <- 0
		}()
		var xytile []int
		out := make(chan int)
		go func() {
			xytile = append(xytile, <-out)
			if len(xytile) == 3 {
				tiles[Point{x: xytile[0], y: xytile[1]}] = xytile[2]
				xytile = nil
			}
		}()
		intcode2(program, in, out)
	*/
	var result int
	for _, v := range tiles {
		if v == BLOCK {
			result += 1
		}
	}
	display(tiles)
	return result
}

func play2(program []int) int {
	var tiles map[Point]int = make(map[Point]int)

	in := func() int {
		findObjectX := func(obj int) int {
			for p, o := range tiles {
				if o == obj {
					return p.x
				}
			}
			return 0
		}
		ballX := findObjectX(BALL)
		paddleX := findObjectX(PADDLE)
		if paddleX > ballX {
			return -1
		} else if paddleX < ballX {
			return 1
		} else {
			return 0
		}
	}

	var xytile []int
	var score int
	out := func(val int) {
		xytile = append(xytile, val)
		if len(xytile) == 3 {
			if xytile[0] == -1 && xytile[1] == 0 {
				score = xytile[2]
			} else {
				tiles[Point{x: xytile[0], y: xytile[1]}] = xytile[2]
			}
			xytile = nil
		}
	}

	intcode(program, in, out)

	return score
}

func display(tiles map[Point]int) {
	for y := 0; y < 24; y++ {
		for x := 0; x < 45; x++ {
			switch tiles[Point{y: y, x: x}] {
			case EMPTY:
				fmt.Print(" ")
			case WALL:
				fmt.Print("#")
			case BLOCK:
				fmt.Print(".")
			case PADDLE:
				fmt.Print("-")
			case BALL:
				fmt.Print("o")
			}
		}
		fmt.Println()
	}
}

func part1() {
	defer duration(time.Now(), "part 1")

	var tokens []string = strings.Split(strings.Trim(input, "\n"), ",")
	var masterProgram []int
	for _, tok := range tokens {
		masterProgram = append(masterProgram, atoi(tok))
	}
	// The computer's available memory should be much larger than the initial program.
	// Memory beyond the initial program starts with the value 0
	// and can be read or written like any other memory.
	var program []int = make([]int, len(masterProgram)+RAMSIZE)
	copy(program, masterProgram)
	var result int = play1(program)
	fmt.Println("part 1", result) // 329
}

func part2() {
	defer duration(time.Now(), "part 2")

	var tokens []string = strings.Split(strings.Trim(input, "\n"), ",")
	var masterProgram []int
	for _, tok := range tokens {
		masterProgram = append(masterProgram, atoi(tok))
	}
	// The computer's available memory should be much larger than the initial program.
	// Memory beyond the initial program starts with the value 0
	// and can be read or written like any other memory.
	var program []int = make([]int, len(masterProgram)+RAMSIZE)
	copy(program, masterProgram)
	program[0] = 2
	var result int = play2(program)
	fmt.Println("part 2", result) // 15973
}

func main() {
	defer duration(time.Now(), "main")

	part1() // 329
	part2() // 15973
}

/*
$ go run main.go
#############################################
#                                           #
# ..    .... ..... .  .. .   .... ...  ...  #
#   . ..   .. ..  .   . ..  .   .   . .   . #
# ... ....  . ..  .  . .... . .  ...  .   . #
# . . ..    ....   ...........      .  .  . #
#  . ...   ..   ..   .  .. .   .... . .  .  #
# .. .. ..   ... .  .. . .. .. . ....   ..  #
# .. .  ..  .  ..  .......   .....   .  ... #
# ...    .   ...  ..     .. .. . .. . ..    #
#  . . .    .  . .     . ..... .. .. . ...  #
# ..   .. ..   ....   ...    .....  .. .... #
# .  ...  . . . .. . . ....... .  . . ....  #
#    ..       ... .. . .  .. ... ..  . . .  #
# . .. ..  .   .  .. . .  .. . . ...  . ... #
# . .. ...  .  .....  .. ...   ..... . ...  #
#     .. .. . ...  . .  . .. ....   ...  .  #
#                                           #
#                   o                       #
#                                           #
#                                           #
#                     -                     #
#                                           #

part 1 329
part 1 2.306692ms
part 2 15973
part 2 83.1133ms
main 85.431986ms
*/
