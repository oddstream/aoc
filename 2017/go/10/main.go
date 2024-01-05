package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"time"
)

//go:embed input.txt
var input string // beware trailing \n

const LISTSIZE int = 256

func duration(invocation time.Time, name string) {
	fmt.Println(name, time.Since(invocation))
}

// or we could use K&R p61
func atoi(s string) int {
	if n, err := strconv.Atoi(s); err == nil {
		return n
	} else {
		fmt.Println(err)
	}
	return 0
}

func partOne() int {
	re := regexp.MustCompile("[[:digit:]]+")
	matches := re.FindAllStringSubmatch(input, -1)
	var list [LISTSIZE]int
	for i := 0; i < LISTSIZE; i++ {
		list[i] = i
	}
	var pos, skip int
	for i := range matches {
		len := atoi(matches[i][0])
		for i, j := pos, pos+len-1; i < j; i, j = i+1, j-1 {
			list[i%LISTSIZE], list[j%LISTSIZE] = list[j%LISTSIZE], list[i%LISTSIZE]
		}
		pos += len + skip
		skip += 1
	}
	return list[0] * list[1]
}

func partTwo() string {
	var list [LISTSIZE]int
	for i := 0; i < LISTSIZE; i++ {
		list[i] = i
	}
	var bytes = input // to allow for test inputs eg "1,2,3" or "AoC 2017"
	var pos, skip int // preserved between rounds
	var round func() = func() {
		for off := 0; off < len(bytes); off++ {
			len := int(bytes[off])
			for i, j := pos, pos+len-1; i < j; i, j = i+1, j-1 {
				list[i%LISTSIZE], list[j%LISTSIZE] = list[j%LISTSIZE], list[i%LISTSIZE]
			}
			pos += len + skip
			skip += 1
		}
		var suffix [5]int = [5]int{17, 31, 73, 47, 23}
		for off := 0; off < len(suffix); off++ {
			len := int(suffix[off])
			for i, j := pos, pos+len-1; i < j; i, j = i+1, j-1 {
				list[i%LISTSIZE], list[j%LISTSIZE] = list[j%LISTSIZE], list[i%LISTSIZE]
			}
			pos += len + skip
			skip += 1
		}
	}
	for i := 0; i < 64; i++ {
		round()
	}
	var dense [16]int
	for i := 0; i < 16; i++ {
		for j := 0; j < 16; j++ {
			dense[i] ^= list[i*16+j]
		}
	}
	var result string
	for i := 0; i < 16; i++ {
		result = result + fmt.Sprintf("%02x", dense[i])
	}
	if len(result) != 32 {
		fmt.Println("ERROR len result is", len(result))
	}
	return result
}

func main() {
	defer duration(time.Now(), "main")

	fmt.Println(partOne()) // 54675
	fmt.Println(partTwo()) // a7af2706aa9a09cf5d848c1e6605dd2a
}

/*
$ go run main.go
54675
a7af2706aa9a09cf5d848c1e6605dd2a
main 202.964µs
*/
