package main

import (
	"fmt"
	"math"
	"time"
)

var grid [300][300]int

func duration(invocation time.Time, name string) {
	fmt.Println(name, time.Since(invocation))
}

// func powerLevel(x, y, serial int) int {
// 	var rackId int = x + 10
// 	var power = rackId * y
// 	power += serial
// 	power *= rackId
// 	power = (power / 100) % 10
// 	power -= 5
// 	return power
// }

func fillGrid(serial int) {
	for i := 0; i < 300; i++ {
		for j := 0; j < 300; j++ {
			var x int = i + 1
			var y int = j + 1
			var rackId int = x + 10
			var power = rackId * y
			power += serial
			power *= rackId
			power = (power / 100) % 10
			power -= 5
			grid[i][j] = power
		}
	}
}

func partOne(serial int) string {
	defer duration(time.Now(), "part 1")
	fillGrid(serial)

	var largestTotalPower int = -math.MaxInt64
	var largestX, largestY int
	for i := 0; i < 300-2; i++ {
		for j := 0; j < 300-2; j++ {
			var tp int
			for k := 0; k < 3; k++ {
				for l := 0; l < 3; l++ {
					tp += grid[i+k][j+l]
				}
			}
			if tp > largestTotalPower {
				largestTotalPower = tp
				largestX = i + 1
				largestY = j + 1
			}
		}
	}
	return fmt.Sprintf("%d -> %d @ %d,%d", serial, largestTotalPower, largestX, largestY)
}

func partTwo(serial int) string {
	defer duration(time.Now(), "part 2")
	fillGrid(serial)

	var largestTotalPower int = -math.MaxInt64
	var largestX, largestY, largestSz int
	for sz := 1; sz <= 300; sz++ {
		for i := 0; i < 300-sz+1; i++ {
			for j := 0; j < 300-sz+1; j++ {
				var tp int
				for k := 0; k < sz; k++ {
					for l := 0; l < sz; l++ {
						tp += grid[i+k][j+l]
					}
				}
				if tp > largestTotalPower {
					largestTotalPower = tp
					largestX = i + 1
					largestY = j + 1
					largestSz = sz
				}
			}
		}
	}
	return fmt.Sprintf("%d -> %d @ %d,%d,%d", serial, largestTotalPower, largestX, largestY, largestSz)
}

func main() {
	defer duration(time.Now(), "main")

	// fmt.Println(powerLevel(3, 5, 8))

	// 18 := 33,45 (power 29)
	// 42 := 21,61 (power 30)
	// 57 := 122,79 (power -5)
	// 39 := 217,196 (power 0)
	// 71 := 101,153 (power 4)
	// 2866 := 20,50 (power 30)
	fmt.Println(partOne(2866)) // 20,50 (power 30)
	fmt.Println(partTwo(2866)) // 238,278,9 (power 88)
}

/*
$ go run main.go
part 1 1.403942ms
2866 -> 30 @ 20,50
part 2 1m6.50428875s
2866 -> 88 @ 238,278,9
main 1m6.505723483s
*/
