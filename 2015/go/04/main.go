package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math"
)

// var input = "abcdef" // md5 hash (in hex) of abcdef609043 starts with 5 zeros
// var input = "pqrstuv" // 1048970

var input = "iwrupvqb" // 346386 (00000), 9958218 (000000)

func main() {
	var partOneAnswer, partTwoAnswer int
	for i := 0; i < math.MaxInt; i++ {
		in := fmt.Sprintf("%s%d", input, i)
		hash := md5.Sum([]byte(in))
		out := hex.EncodeToString(hash[:])
		if out[0] == '0' && out[1] == '0' && out[2] == '0' && out[3] == '0' && out[4] == '0' {
			fmt.Println(i, in, hash, out)
			partOneAnswer = i
			break
		}
	}
	for i := partOneAnswer; i < math.MaxInt; i++ {
		in := fmt.Sprintf("%s%d", input, i)
		hash := md5.Sum([]byte(in))
		out := hex.EncodeToString(hash[:])
		if out[0] == '0' && out[1] == '0' && out[2] == '0' && out[3] == '0' && out[4] == '0' && out[5] == '0' {
			fmt.Println(i, in, hash, out)
			partTwoAnswer = i
			break
		}
	}
	fmt.Println(partOneAnswer, partTwoAnswer)
}
