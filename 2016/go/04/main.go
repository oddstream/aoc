package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"sort"
	"strings"
	"time"
	"unicode"
)

//go:embed input.txt
var input string

type Letter struct {
	r rune
	n int
}

type Room struct {
	name     string
	letters  [26]Letter
	sectorID int
	checksum string
}

func shiftCipher(s string, n int) string {
	var out []rune
	runes := []rune(s)
	for _, r := range runes {
		if r == '-' {
			out = append(out, ' ')
		} else {
			for i := 0; i < n; i++ {
				r += 1
				if r > 'z' {
					r = 'a'
				}
			}
			out = append(out, r)
		}
	}
	return string(out)
}

func parseRoom(s string) Room {
	var rm Room
	for r := 'a'; r <= 'z'; r++ {
		rm.letters[r-'a'].r = r
	}
	var i int
	var r rune
	for i, r = range s {
		if unicode.IsDigit(r) {
			break
		}
		rm.name = rm.name + string(r)
		if r == '-' {
			continue
		}
		rm.letters[r-'a'].r = r
		rm.letters[r-'a'].n += 1
	}
	fmt.Sscanf(s[i:], `%d%s`, &rm.sectorID, &rm.checksum)
	rm.checksum = strings.TrimPrefix(rm.checksum, "[")
	rm.checksum = strings.TrimSuffix(rm.checksum, "]")
	return rm
}

func validRoom(rm Room) bool {
	sl := rm.letters[:]
	sort.Slice(sl, func(a, b int) bool {
		if sl[a].n == sl[b].n {
			return sl[a].r < sl[b].r
		}
		return sl[a].n > sl[b].n
	})
	var runes []rune
	for i := 0; i < 5; i++ {
		runes = append(runes, sl[i].r)
	}
	return string(runes) == rm.checksum
}

func decodeRoom(rm Room, target string) {
	s := shiftCipher(rm.name, rm.sectorID)
	if strings.Contains(s, target) {
		fmt.Println(rm.sectorID, s)
	}
}

func duration(invocation time.Time, name string) {
	fmt.Println(name, time.Since(invocation))
}

func partOne() int {
	var result int
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		rm := parseRoom(scanner.Text())
		if validRoom(rm) {
			result += rm.sectorID
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	return result
}

func partTwo() {
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		rm := parseRoom(scanner.Text())
		if validRoom(rm) {
			decodeRoom(rm, "north")
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}

func main() {
	defer duration(time.Now(), "main")

	fmt.Println(partOne()) // 278221
	partTwo()              // 267 northpole object storage
}
