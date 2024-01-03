// https://adventofcode.com/2017/day/1 Inverse Captcha
package main

import (
	_ "embed"
	"fmt"
	"time"
)

//go:embed input.txt
var input []byte // remember to strip trailing \n

func duration(invocation time.Time, name string) {
	fmt.Println(name, time.Since(invocation))
}

func partOne(in []byte) int {
	if len(in) == 0 {
		in = input
	}
	var result int
	for i := 1; i < len(in); i++ {
		if in[i-1] == in[i] {
			result += int(in[i] - '0')
		}
	}
	if in[0] == in[len(in)-1] {
		result += int(in[0] - '0')
	}
	return result
}

func partTwo(in []byte) int {
	if len(in) == 0 {
		in = input
	}
	var result int
	off := len(in) / 2
	for i, b := range in {
		n1 := int(b - '0')
		n2 := int(in[(i+off)%len(in)] - '0')
		if n1 == n2 {
			result += n1
		}
	}
	return result
}

func main() {
	defer duration(time.Now(), "main")

	fmt.Println(partOne([]byte{})) // 1177
	fmt.Println(partTwo([]byte{})) // 1060
}
