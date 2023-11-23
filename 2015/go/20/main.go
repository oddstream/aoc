package main

import (
	_ "embed"
	"flag"
	"fmt"
	"log"
	"math"
	"time"
)

var input int = 33100000

func duration(invocation time.Time, name string) {
	fmt.Println(name, time.Since(invocation))
}

// slow way of creating divisors of n
// func getDivisors(n int) []int {
// 	divisors := []int{}
// 	for i := 1; i <= n; i++ {
// 		if n%i == 0 {
// 			divisors = append(divisors, i)
// 		}
// 	}
// 	return divisors
// }

func getDivisors(n int) []int {
	divisors := []int{}
	for i := 1; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			divisors = append(divisors, i)
			if i != n/i {
				divisors = append(divisors, n/i)
			}
		}
	}
	return divisors
}

func sum(arr []int) int {
	var s int
	for _, n := range arr {
		s += n
	}
	return s
}

func part1() int {
	for i := 0; ; i++ {
		divisors := getDivisors(i)
		if sum(divisors)*10 > input {
			return i
		}
	}
}

func part2() int {
	for i := 0; ; i++ {
		divisors := getDivisors(i)
		var d50 []int
		for _, d := range divisors {
			if i/d <= 50 {
				d50 = append(d50, d)
			}
		}
		if sum(d50)*11 > input {
			return i
		}
	}
}

func main() {
	defer duration(time.Now(), "main")

	var part int
	flag.IntVar(&part, "part", 1, "1 or 2")
	flag.Parse()

	if part == 1 {
		log.Println(part1()) // 776160
	} else if part == 2 {
		log.Println(part2()) // 786240
	}
}

/*
			elf
	haus	1	2	3	4	5	6	7	8	9	Tot
	1		10									10
	2		10	20								30
	3		10		30							40
	4		10	20		40						70
	5		10				50					60
	6		10	20	30			60				120
	7		10						70			90
	8		10	20		40				80		150
	9		10		30						90	130

	Since each elf visits every i'th house,
	the number of presents house N will get
	is the sum of the divisors of N (inclusive).
*/
