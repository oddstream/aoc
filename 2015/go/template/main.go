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

func parseInput() {
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("scanner error: %v\n", err)
	}
}

func main() {
	defer duration(time.Now(), "main")
	parseInput()
}

