// https://adventofcode.com/2017/day/14 Disk Defragmentation
package main

import (
	"fmt"
	"time"
)

func duration(invocation time.Time, name string) {
	fmt.Println(name, time.Since(invocation))
}

// func hammingWeight(n uint32) int {
// 	count := 0
// 	for n != 0 {
// 		count += int(n & 1) // Check the least significant bit
// 		n >>= 1             // Shift right to examine the next bit
// 	}
// 	return count
// }

// map from lower case hex char to Hamming weight
var ch2ham map[byte]int = map[byte]int{
	'0': 0, //hammingWeight(0),
	'1': 1, //hammingWeight(1),
	'2': 1, //hammingWeight(2),
	'3': 2, //hammingWeight(3),
	'4': 1, //hammingWeight(4),
	'5': 2, //hammingWeight(5),
	'6': 2, //hammingWeight(6),
	'7': 3, //hammingWeight(7),
	'8': 1, //hammingWeight(8),
	'9': 2, //hammingWeight(9),
	'a': 2, //hammingWeight(10),
	'b': 3, //hammingWeight(11),
	'c': 2, //hammingWeight(12),
	'd': 3, //hammingWeight(13),
	'e': 3, //hammingWeight(14),
	'f': 4, //hammingWeight(15),
}

// var hashmap map[byte]string = map[byte]string{
// 	'0': "....",
// 	'1': "...#",
// 	'2': "..#.",
// 	'3': "..##",
// 	'4': ".#..",
// 	'5': ".#.#",
// 	'6': ".##.",
// 	'7': ".###",
// 	'8': "#...",
// 	'9': "#..#",
// 	'a': "#.#.",
// 	'b': "#.##",
// 	'c': "##..",
// 	'd': "##.#",
// 	'e': "###.",
// 	'f': "####",
// }

// map from lower case hex char to [4]int
var ch24int map[byte][4]int = map[byte][4]int{
	'0': {0, 0, 0, 0},
	'1': {0, 0, 0, -1},
	'2': {0, 0, -1, 0},
	'3': {0, 0, -1, -1},
	'4': {0, -1, 0, 0},
	'5': {0, -1, 0, -1},
	'6': {0, -1, -1, 0},
	'7': {0, -1, -1, -1},
	'8': {-1, 0, 0, 0},
	'9': {-1, 0, 0, -1},
	'a': {-1, 0, -1, 0},
	'b': {-1, 0, -1, -1},
	'c': {-1, -1, 0, 0},
	'd': {-1, -1, 0, -1},
	'e': {-1, -1, -1, 0},
	'f': {-1, -1, -1, -1},
}

// func borrowed from https://adventofcode.com/2017/day/10 solution
// and hard-wired t0 256 list length
func knothash(input string) string {
	var list [256]int
	for i := 0; i < 256; i++ {
		list[i] = i
	}
	var pos, skip int // preserved between rounds
	var round func() = func() {
		for off := 0; off < len(input); off++ {
			len := int(input[off])
			for i, j := pos, pos+len-1; i < j; i, j = i+1, j-1 {
				list[i%256], list[j%256] = list[j%256], list[i%256]
			}
			pos += len + skip
			skip += 1
		}
		var suffix [5]int = [5]int{17, 31, 73, 47, 23}
		for off := 0; off < len(suffix); off++ {
			len := int(suffix[off])
			for i, j := pos, pos+len-1; i < j; i, j = i+1, j-1 {
				list[i%256], list[j%256] = list[j%256], list[i%256]
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

func partOne(key string) int {
	defer duration(time.Now(), "part 1")
	var result int
	for row := 0; row < 128; row++ {
		khash := knothash(fmt.Sprintf("%s-%d", key, row))
		for i := 0; i < len(khash); i++ {
			ch := khash[i]
			result += ch2ham[ch]
		}
	}
	return result
}

func partTwo(key string) int {
	defer duration(time.Now(), "part 2")

	var numgrid [128][128]int
	for y := 0; y < 128; y++ {
		khash := knothash(fmt.Sprintf("%s-%d", key, y))
		var z int
		for x := 0; x < len(khash); x++ {
			ch := khash[x]
			i4 := ch24int[ch]
			numgrid[y][z] = i4[0]
			numgrid[y][z+1] = i4[1]
			numgrid[y][z+2] = i4[2]
			numgrid[y][z+3] = i4[3]
			z += 4
		}
	}

	var coloring func(int, int, int)
	coloring = func(y, x, n int) {
		numgrid[y][x] = n
		for _, d := range [][2]int{{0, -1}, {1, 0}, {0, 1}, {-1, 0}} {
			ny := y + d[1]
			nx := x + d[0]
			if nx < 0 || ny < 0 || nx > 127 || ny > 127 {
				continue
			}
			if numgrid[ny][nx] == -1 {
				coloring(ny, nx, n)
			}
		}
	}

	var color int
	for y := 0; y < 128; y++ {
		for x := 0; x < 128; x++ {
			if numgrid[y][x] == -1 {
				coloring(y, x, color)
				color += 1
			}
		}
	}
	return color
}

func main() {
	defer duration(time.Now(), "main")

	fmt.Println(partOne("xlqgujun")) // 8204
	// fmt.Println(partTwo("flqrgnkx")) // 1242
	fmt.Println(partTwo("xlqgujun")) // 1089
}

/*
go run main.go
part 1 7.273282ms
8204
part 2 7.326349ms
1089
main 14.621783ms
*/
