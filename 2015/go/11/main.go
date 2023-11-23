package main

import (
	"fmt"
)

func incrementString(r []rune) []rune {
	for i := len(r) - 1; i >= 0; i-- {
		if r[i] == 'z' {
			r[i] = 'a'
		} else {
			r[i]++
			break
		}
	}
	return r
}

func increasingStraight3(r []rune) bool {
	for i := 0; i < len(r)-2; i++ {
		r0 := r[i]
		r1 := r0 + 1
		r2 := r1 + 1
		if r[i+1] == r1 && r[i+2] == r2 {
			return true
		}
	}
	return false
}

func confusingLetters(r []rune) bool {
	for i := 0; i < len(r); i++ {
		r0 := r[i]
		if r0 == 'i' || r0 == 'o' || r0 == 'l' {
			return true
		}
	}
	return false
}

func twoDifferentPairs(r []rune) bool {
	for i := 0; i < len(r)-1; i++ {
		r0 := r[i]
		r1 := r[i+1]
		if r0 == r1 {
			// we found a pair, look in rest of string for another
			for j := i + 2; j < len(r)-1; j++ {
				s0 := r[j]
				s1 := r[j+1]
				if s0 == s1 && s0 != r0 {
					return true
				}
			}
		}
	}
	return false
}

func main() {
	/*
		var str []rune = []rune("xyz")
		fmt.Println(str, string(str))
		for i := 0; i < 3; i++ {
			str = incrementString(str)
			fmt.Println(str, string(str))
		}
		fmt.Println(increasingStraight3([]rune("abd")))
		fmt.Println(increasingStraight3([]rune("hijklmmn")))
		fmt.Println(confusingLetters([]rune("abd")))
		fmt.Println(confusingLetters([]rune("abi")))
		fmt.Println(twoDifferentPairs(([]rune("abi"))))
		fmt.Println(twoDifferentPairs(([]rune("aabii"))))
	*/
	// var str []rune = []rune("hxbxwxba")
	var str []rune = []rune("hxbxxyzz")
	fmt.Println(0, str, string(str))
	for i := 0; i < 1000000; i++ {
		str = incrementString(str)
		if increasingStraight3(str) {
			if !confusingLetters(str) {
				if twoDifferentPairs(str) {
					fmt.Println(i, str, string(str))
					break
				}
			}
		}
	}
	// part 1 hxbxxyzz
	// part 2 hxcaabcc
}
