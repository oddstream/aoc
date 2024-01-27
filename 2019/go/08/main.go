// https://adventofcode.com/2019/day/8
package main

import (
	_ "embed"
	"fmt"
	"math"
	"time"
)

//go:embed input.txt
var input []byte

const (
	BLACK       = '0'
	WHITE       = '1'
	TRANSPARENT = '2'
)

func duration(invocation time.Time, name string) {
	fmt.Println(name, time.Since(invocation))
}

func partOne() {
	defer duration(time.Now(), "part 1")

	fmt.Println("input is", len(input), "bytes ==", len(input)/(25*6), "layers of", 25*6, "bytes each")
	var layers map[int][3]int = make(map[int][3]int)
	var layer int
	for i := 0; i < len(input); i += 150 {
		var count0, count1, count2 int
		for j := 0; j < 150; j++ {
			switch input[i+j] {
			case '0':
				count0 += 1
			case '1':
				count1 += 1
			case '2':
				count2 += 1
			}
		}
		layers[layer] = [3]int{count0, count1, count2}
		layer += 1
	}

	var min0layer int
	var min0 int = math.MaxInt64
	for i, v := range layers {
		if v[0] < min0 {
			min0layer = i
			min0 = v[0]
		}
	}

	var result int = layers[min0layer][1] * layers[min0layer][2]
	fmt.Println("part 1", result) // 1905
}

func getLayer(y, x int) [100]byte {
	var layer [100]byte
	var loffset = y*25 + x     // offset within layer
	for i := 0; i < 100; i++ { // for each layer
		layer[i] = input[loffset+(i*150)]
	}
	return layer
}

func decodeLayer(layer [100]byte) byte {
	var result byte = TRANSPARENT
	for i := 99; i >= 0; i-- {
		if layer[i] != TRANSPARENT {
			result = layer[i]
		}
	}
	return result
}

func partTwo() {
	defer duration(time.Now(), "part 2")

	for y := 0; y < 6; y++ {
		for x := 0; x < 25; x++ {
			layer := getLayer(y, x)
			if decodeLayer(layer) == WHITE {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func main() {
	defer duration(time.Now(), "main")

	partOne() // 1905
	partTwo() // ACKPZ
}

/*
$ go run main.go
input is 15000 bytes == 100 layers of 150 bytes each
part 1 1905
part 1 140.917µs
 **   **  *  * ***  ****
*  * *  * * *  *  *    *
*  * *    **   *  *   *
**** *    * *  ***   *
*  * *  * * *  *    *
*  *  **  *  * *    ****
part 2 261.565µs
main 413.996µs
*/
