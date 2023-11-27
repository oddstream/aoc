package main

import "log"

const (
	value   = 20151125
	base    = 252533
	divisor = 33554393
	row     = 2981
	col     = 3075
)

func findExponent(row, col int) int {
	var result int = col - 1
	startingRow := row + col - 1
	for i := 1; i < startingRow; i++ {
		result += i
	}
	return result
}

func findValue(b, e, d int) int {
	var result int = 1
	var base int = b
	var exp int = e
	var divisor int = d
	for exp > 0 {
		if exp%2 == 1 {
			result = result * base % divisor
		}
		exp /= 2
		base = base * base % divisor
	}
	result = result * value % d
	return result
}

func main() {
	var exp int = findExponent(row, col)
	log.Println(findValue(base, exp, divisor)) // 9132360
}
