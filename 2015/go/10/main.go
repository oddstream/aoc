package main

import (
	"fmt"
	"strings"
)

func look_and_say(in string) string {
	var sb strings.Builder
	var i int = 0
	for i < len(in) {
		var j int
		for j = i + 1; j < len(in); j++ {
			if in[i] != in[j] {
				break
			}
		}
		sb.WriteString(fmt.Sprintf("%c%c", j-i+int(byte('0')), in[i]))
		i = j
	}
	return sb.String()
}

func main() {
	s := "3113322113"
	for i := 0; i <= 50; i++ {
		fmt.Println(i, len(s))
		s = look_and_say(s)
	}
	// 40 times =  329356
	// 50 times = 4666278
}
