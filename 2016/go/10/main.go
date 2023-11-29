// https://adventofcode.com/2016/day/10
package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"strconv"
	"strings"
	"time"
)

//go:embed input.txt
var input string

// lowest value is 1
// highest value is 73?

// lowest bot is 0
// highest bot is 209

// lowest output is 0
// highest output is 20

const (
	// MAXBOTS  int = 4
	// ANSWERLO int = 2
	// ANSWERHI int = 5
	MAXBOTS    int = 210
	ANSWERLO       = 17
	ANSWERHI       = 61
	MAXOUTPUTS int = 21
)

var bots [MAXBOTS][2]int

var outputs [MAXOUTPUTS]int

func duration(invocation time.Time, name string) {
	fmt.Println(name, time.Since(invocation))
}

func addValueToBot(val, bot int) {
	if bots[bot][0] == 0 {
		bots[bot][0] = val
	} else if bots[bot][1] == 0 {
		bots[bot][1] = val
	} else {
		fmt.Printf("BOT %d IS FULL\n", bot)
	}
}

func addValueToOutput(val, out int) {
	if outputs[out] == 0 {
		outputs[out] = val
	} else {
		fmt.Printf("OUTPUT %d IS FULL\n", out)
	}
}

func extractValues(bot int) (int, int) {
	if bots[bot][0] == 0 || bots[bot][1] == 0 {
		return 0, 0
		// fmt.Println("BOT HAS ZERO VALUE", bot)
	}
	var lo, hi int
	if bots[bot][0] > bots[bot][1] {
		hi = bots[bot][0]
		lo = bots[bot][1]
	} else {
		hi = bots[bot][1]
		lo = bots[bot][0]
	}
	bots[bot][0] = 0
	bots[bot][1] = 0
	return lo, hi
}

func loadBots() {
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "value ") {
			var val, bot int
			var err error
			_, err = fmt.Sscanf(line, "value %d goes to bot %d", &val, &bot)
			if err != nil {
				fmt.Println("ERROR ON LINE", line)
				return
			}
			addValueToBot(val, bot)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	// fmt.Println(bots)
}

func runBots(part int) int {
	// part 1 16 loops
	// part 2 23 loops
	for {
		scanner := bufio.NewScanner(strings.NewReader(input))
		for scanner.Scan() {
			if strings.HasPrefix(scanner.Text(), "bot ") {
				fields := strings.Fields(scanner.Text())
				var srcbot, dstlo, dsthi int
				var dstloname, dsthiname string
				var err error
				// 0   1 2     3   4  5      6 7   8    9  10  11
				// bot 1 gives low to output 1 and high to bot 0

				if !(fields[3] == "low" && fields[4] == "to") {
					fmt.Println("ERROR", scanner.Text())
					break
				}
				if !(fields[8] == "high" && fields[9] == "to") {
					fmt.Println("ERROR", scanner.Text())
					break
				}
				if srcbot, err = strconv.Atoi(fields[1]); err != nil {
					fmt.Println(err)
					break
				}
				dstloname = fields[5]
				if dstlo, err = strconv.Atoi(fields[6]); err != nil {
					fmt.Println(err)
					break
				}
				dsthiname = fields[10]
				if dsthi, err = strconv.Atoi(fields[11]); err != nil {
					fmt.Println(err)
					break
				}

				// fmt.Println(srcbot, dstloname, dstlo, dsthiname, dsthi)

				lo, hi := extractValues(srcbot)
				if lo == 0 || hi == 0 {
					// bot does not have two microchips
					continue
				}
				if part == 1 && (lo == ANSWERLO && hi == ANSWERHI) {
					return srcbot
				}
				switch dstloname {
				case "output":
					addValueToOutput(lo, dstlo)
				case "bot":
					addValueToBot(lo, dstlo)
				default:
					fmt.Println(scanner.Text(), dstloname)
				}
				switch dsthiname {
				case "output":
					addValueToOutput(hi, dsthi)
				case "bot":
					addValueToBot(hi, dsthi)
				default:
					fmt.Println(scanner.Text(), dsthiname)
				}
			}
		}
		if err := scanner.Err(); err != nil {
			fmt.Println(err)
		}
		if part == 2 && (outputs[0] != 0 && outputs[1] != 0 && outputs[2] != 0) {
			return outputs[0] * outputs[1] * outputs[2]
		}
	}
}

func partOne() int {
	loadBots()
	return runBots(1)
}

func partTwo() int {
	loadBots()
	return runBots(2)
}

func main() {
	defer duration(time.Now(), "main")

	// TODO optimize by loading the rules ONCE

	fmt.Println(partOne()) // 116
	for i := 0; i < len(bots); i++ {
		bots[i][0] = 0
		bots[i][1] = 0
	}
	for i := 0; i < len(outputs); i++ {
		outputs[i] = 0
	}
	fmt.Println(partTwo()) // 23903
}
