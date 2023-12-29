// https://adventofcode.com/2016/day/14 One-Time Pad

package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math"
	"time"
)

var input = "ahsbgdzn"

// var input = "abc"

func duration(invocation time.Time, name string) {
	fmt.Println(name, time.Since(invocation))
}

func consecutiveN(s string, n int) byte {
	count := 1
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			count++
			if count == n {
				return s[i]
			}
		} else {
			count = 1
		}
	}
	return 0
}

func consecutiveNCh(s string, n int, ch byte) bool {
	count := 0
	for i := 0; i < len(s); i++ {
		if s[i] == ch {
			count++
			if count == n {
				return true
			}
		} else {
			count = 0
		}
	}
	return false
}

var hashCache map[string]string = make(map[string]string)

func hash1(n int) string {
	in := fmt.Sprintf("%s%d", input, n)
	if str, ok := hashCache[in]; ok {
		return str
	}
	hash := md5.Sum([]byte(in))
	out := hex.EncodeToString(hash[:])
	hashCache[in] = out
	return out
}

func hash2(n int) string {
	in := fmt.Sprintf("%s%d", input, n)
	if str, ok := hashCache[in]; ok {
		return str
	}
	hash := md5.Sum([]byte(in))
	for i := 0; i < 2016; i++ {
		out := hex.EncodeToString(hash[:])
		hash = md5.Sum([]byte(out))
	}
	out := hex.EncodeToString(hash[:])
	hashCache[in] = out
	return out
}

func partOne() int {
	var keys int
	for i := 0; i < math.MaxInt; i++ {
		h := hash1(i)
		if ch := consecutiveN(h, 3); ch != 0 {
			for j := i + 1; j < i+1+1000; j++ {
				h2 := hash1(j)
				if consecutiveNCh(h2, 5, ch) {
					keys += 1
					if keys == 64 {
						return i
					}
				}
			}
		}
	}
	return -1
}

func partTwo() int {
	var keys int
	for i := 0; i < math.MaxInt; i++ {
		h := hash2(i)
		if ch := consecutiveN(h, 3); ch != 0 {
			for j := i + 1; j < i+1+1000; j++ {
				h2 := hash2(j)
				if consecutiveNCh(h2, 5, ch) {
					// println(keys, i, j)
					keys += 1
					if keys == 64 {
						return i
					}
				}
			}
		}
	}
	return -1
}

func main() {
	defer duration(time.Now(), "main")
	fmt.Println(partOne()) // 22728, 23890
	fmt.Println(partTwo()) // 22551, 22696
}

/*
$ go run main.go
23890
23890
main 1.123288537s

inside debugger (2x slower):

23890
23890
main 2.216892302s
*/
