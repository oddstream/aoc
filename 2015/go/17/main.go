package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

//go:embed input.txt
var input string

var containers []int

func duration(invocation time.Time, name string) {
	fmt.Println(name, time.Since(invocation))
}

func parseInput() {
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		capacity, _ := strconv.Atoi(scanner.Text())
		containers = append(containers, capacity)
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("scanner error: %v\n", err)
	}
}

func combifunc(arr []int) [][]int {
	var result [][]int
	for i := 0; i < (1 << uint(len(arr))); i++ {
		var subset []int
		for j := 0; j < len(arr); j++ {
			if i&(1<<uint(j)) != 0 {
				subset = append(subset, arr[j])
			}
		}
		result = append(result, subset)
	}
	return result
}

func permute(arr []int, c chan []int) {
	for i := 0; i < (1 << uint(len(arr))); i++ {
		var subset []int
		for j := 0; j < len(arr); j++ {
			if i&(1<<uint(j)) != 0 {
				subset = append(subset, arr[j])
			}
		}
		c <- subset
	}
	close(c)
}

func checksum(array []int) uint {
	var c uint = 0
	for i := 0; i < len(array); i++ {
		c += uint(array[i])
		c = c<<3 | c>>(32-3) // rotate a little
		c ^= 0xFFFFFFFF      // invert just for fun
	}
	return c
}

func main() {
	defer duration(time.Now(), "main")
	parseInput()
	fmt.Println(containers)

	c := make(chan []int)
	go permute(containers, c)
	var count, combicount, min int
	var seen map[uint]struct{} = make(map[uint]struct{})
	min = math.MaxInt64
	for arr := range c {
		combicount += 1
		var total int
		for _, ele := range arr {
			total += ele
		}
		if total == 150 && len(arr) == 4 {
			count += 1
			seen[checksum(arr)] = struct{}{}
			if len(arr) < min {
				min = len(arr)
				fmt.Println(arr)
			}
		}
	}
	fmt.Println(combicount, min, count, len(seen)) // part 1 4372 (allows duplicates)
	// fmt.Println(len(combifunc([]int{47, 31, 36, 36})))
	// part 2 4
}
