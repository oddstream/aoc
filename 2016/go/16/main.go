package main

import (
	"fmt"
	"time"
)

var input []rune = []rune("10001110011110000") // 17 bits wide

func duration(invocation time.Time, name string) {
	fmt.Println(name, time.Since(invocation))
}

func copyr(r []rune) []rune {
	r2 := make([]rune, len(r))
	copy(r2, r)
	return r2
}

func reverser(r []rune) {
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
}

func invertr(r []rune) {
	for i := 0; i < len(r); i++ {
		switch r[i] {
		case '1':
			r[i] = '0'
		case '0':
			r[i] = '1'
		}
	}
}

func joinr(a, b []rune) []rune {
	r := make([]rune, len(a)+1+len(b))
	i := 0
	for j := 0; j < len(a); j++ {
		r[i] = a[j]
		i += 1
	}
	r[i] = '0'
	i += 1
	for j := 0; j < len(b); j++ {
		r[i] = b[j]
		i += 1
	}
	return r
}

func checksumr(in []rune, limit int) []rune {
	out := make([]rune, limit/2)
	j := 0
	for i := 0; i < limit; {
		if in[i] == in[i+1] {
			out[j] = '1'
		} else {
			out[j] = '0'
		}
		i += 2
		j += 1
	}
	return out
}

func checksumr2(in []rune, limit int) []rune {
	r := checksumr(in, limit)
	for len(r)%2 == 0 {
		r = checksumr(r, len(r))
	}
	return r
}

func processr(in []rune) []rune {
	a := copyr(in)
	b := copyr(in)
	reverser(b)
	invertr(b)
	return joinr(a, b)
}

func partOne() {
	r := copyr(input)
	for {
		r = processr(r)
		if len(r) > 272 {
			break
		}
	}
	// fmt.Println(len(r), r, string(r))
	chk := checksumr2(r, 272)
	fmt.Println(len(chk), chk, string(chk)) // 10010101010011101
}

func partTwo() {
	r := copyr(input)
	for {
		r = processr(r)
		if len(r) > 35651584 {
			break
		}
	}
	// fmt.Println(len(r), r, string(r))
	chk := checksumr2(r, 35651584)
	fmt.Println(len(chk), chk, string(chk)) // 01100111101101111
}

func main() {
	defer duration(time.Now(), "main")

	partOne()
	partTwo()
}

/*
$ go run main.go
17 [49 48 48 49 48 49 48 49 48 49 48 48 49 49 49 48 49] 10010101010011101
17 [48 49 49 48 48 49 49 49 49 48 49 49 48 49 49 49 49] 01100111101101111
main 194.18036ms
*/
