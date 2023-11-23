package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"sort"
	"strings"
)

//go:embed "input.txt"
var input string

func main() {
	var dims []int = make([]int, 3)
	var totalPaper, totalRibbon int
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		_, err := fmt.Sscanf(scanner.Text(), "%dx%dx%d", &dims[0], &dims[1], &dims[2])
		if err != nil {
			panic(err)
		}
		sort.Ints(dims) // or sort.Ints(dims[:]) if using an array like [3]int
		// fmt.Println(dimensions[0], dimensions[1], dimensions[2])
		// area = 2*l*w + 2*w*h + 2*h*l
		area := 2*dims[0]*dims[1] + 2*dims[1]*dims[2] + 2*dims[2]*dims[0]
		extra := dims[0] * dims[1]
		totalPaper += area + extra
		ribbon := dims[0] + dims[0] + dims[1] + dims[1]
		bow := dims[0] * dims[1] * dims[2]
		totalRibbon += ribbon + bow
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("scanner error: %v\n", err)
	}

	fmt.Println("paper", totalPaper)   // 1606483
	fmt.Println("ribbon", totalRibbon) // 3842356
}
