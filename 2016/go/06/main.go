package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"math"
	"strings"
	"time"
)

//go:embed input.txt
var input string

const LINELEN int = 8 // 6 for test.txt, 8 for input.txt

func duration(invocation time.Time, name string) {
	fmt.Println(name, time.Since(invocation))
}

func partOne() string {
	var charfreq [LINELEN][26]int
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()
		for column := 0; column < LINELEN; column++ {
			by := line[column]
			charfreq[column][by-byte('a')] += 1
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	// fmt.Println(charfreq)
	var result [LINELEN]byte
	for column := 0; column < LINELEN; column++ {
		var maxi, maxval int
		for i := 0; i < 26; i++ {
			if charfreq[column][i] > maxval {
				maxval = charfreq[column][i]
				maxi = i
			}
		}
		result[column] = byte(byte(maxi) + byte('a'))
	}
	return string(result[:])
}

func partTwo() string {
	var charfreq [LINELEN][26]int
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()
		for column := 0; column < LINELEN; column++ {
			by := line[column]
			charfreq[column][by-byte('a')] += 1
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	// fmt.Println(charfreq)
	var result [LINELEN]byte
	for column := 0; column < LINELEN; column++ {
		var mini int
		var minval int = math.MaxInt64
		for i := 0; i < 26; i++ {
			if charfreq[column][i] != 0 && charfreq[column][i] < minval {
				minval = charfreq[column][i]
				mini = i
			}
		}
		result[column] = byte(byte(mini) + byte('a'))
	}
	return string(result[:])
}

func main() {
	defer duration(time.Now(), "main")

	fmt.Println(partOne()) // easter, qqqluigu
	fmt.Println(partTwo()) // advent, lsoypmia
}
