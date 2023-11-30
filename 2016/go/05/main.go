// https://adventofcode.com/2016/day/5
package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math"
	"strings"
	"time"
)

var input string = "uqwqemis"

// var input string = "abc"

func duration(invocation time.Time, name string) {
	fmt.Println(name, time.Since(invocation))
}

func partOne() string {
	defer duration(time.Now(), "partOne")
	var password []byte
	for i := 0; i < math.MaxInt; i++ {
		in := fmt.Sprintf("%s%d", input, i)
		hash := md5.Sum([]byte(in))
		out := hex.EncodeToString(hash[:])
		if strings.HasPrefix(out, "00000") {
			password = append(password, out[5])
			if len(password) == 8 {
				// fmt.Println(string(c), i, in, hash, out)
				// fmt.Println("max i", i)
				break
			}
		}
	}
	return string(password)
}

func partTwo() string {
	defer duration(time.Now(), "partTwo")
	var password [8]byte
	for i := 0; i < math.MaxInt; i++ {
		in := fmt.Sprintf("%s%d", input, i)
		hash := md5.Sum([]byte(in))
		out := hex.EncodeToString(hash[:])
		if strings.HasPrefix(out, "00000") {
			pos := out[5] - '0'
			ch := out[6]
			// fmt.Printf("pos %d char %s", pos, string(ch))
			if pos < 8 && password[pos] == 0 {
				password[pos] = ch
				var blanks int
				for n := 0; n < 8; n++ {
					if password[n] == 0 {
						blanks += 1
					}
				}
				if blanks == 0 {
					// fmt.Println("max i", i)
					break
				}
			}
			// fmt.Printf(" password '%s'\n", string(password[:]))
		}
	}
	return string(password[:])
}

func main() {
	defer duration(time.Now(), "main")
	fmt.Println(partOne()) // 1a3099aa, 30s, max i 16734551
	fmt.Println(partTwo()) // 694190cd, 43s, max i 26326685
}
