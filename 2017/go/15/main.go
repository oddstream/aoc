// https://adventofcode.com/2017/day/15  Dueling Generators
package main

import (
	"fmt"
	"time"
)

func duration(invocation time.Time, name string) {
	fmt.Println(name, time.Since(invocation))
}

func partOne(a, b int) int {
	defer duration(time.Now(), "part 1")

	var result int
	for i := 0; i < 40e6; i++ {
		a *= 16807
		a %= 2147483647
		b *= 48271
		b %= 2147483647
		if a&0xffff == b&0xffff {
			result += 1
		}
	}
	return result
}

func partTwo(a, b int) int {
	defer duration(time.Now(), "part 2")

	var result int
	cha := make(chan int)
	chb := make(chan int)

	go func() {
		var count int
		for {
			a *= 16807
			a %= 2147483647
			if a%4 == 0 {
				cha <- a
				count += 1
				if count == 5e6 {
					break
				}
			}
		}
		close(cha)
	}()

	go func() {
		var count int
		for {
			b *= 48271
			b %= 2147483647
			if b%8 == 0 {
				chb <- b
				count += 1
				if count == 5e6 {
					break
				}
			}
		}
		close(chb)
	}()

	for {
		vala, oka := <-cha
		valb, okb := <-chb
		if !oka || !okb {
			break
		}
		if vala&0xffff == valb&0xffff {
			result += 1
		}
	}
	return result
}

func main() {
	defer duration(time.Now(), "main")

	// fmt.Println("part 1", partOne(65, 8921)) // 588
	fmt.Println("part 1", partOne(116, 299)) // 569
	// fmt.Println("part 2", partTwo(65, 8921)) // 309
	fmt.Println("part 2", partTwo(116, 299)) // 298

}

/*
$ go run main.go
part 1 569
part 2 298
main 2.134356086s
*/

/* some neat ES6 for part 2
let A = 116, B = 299;
let score = 0;

for (let i = 0; i < 5E6; i++) {
    do { A = (A * 16807) % 2147483647; } while (A & 3);
    do { B = (B * 48271) % 2147483647; } while (B & 7);
    if ((A & 0xFFFF) == (B & 0xFFFF)) score++;
}

console.log(score);
*/
