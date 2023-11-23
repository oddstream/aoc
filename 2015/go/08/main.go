package main

import (
	"bufio"
	_ "embed"
	"flag"
	"fmt"
	"strconv"
	"strings"
)

//go:embed "input.txt"
var input string

func main() {
	var part int
	flag.IntVar(&part, "part", 2, "1 or 2")
	flag.Parse()

	var len1, len2, len3 int
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		str, err := strconv.Unquote(scanner.Text())
		if err != nil {
			fmt.Println(err)
		}
		len1 += len(scanner.Text())
		len2 += len(str)
		str2 := strconv.Quote(scanner.Text())
		len3 += len(str2)
	}
	fmt.Println(len1, "-", len2, "=", len1-len2) // 1350
	fmt.Println(len3, "-", len1, "=", len3-len1) // 2085
}
