package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"strings"
	"time"
)

//go:embed test.txt
var input string

func duration(invocation time.Time, name string) {
	fmt.Println(name, time.Since(invocation))
}

func partOne() string {
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	return ""
}

func main() {
	defer duration(time.Now(), "main")

	fmt.Println(partOne())
}

