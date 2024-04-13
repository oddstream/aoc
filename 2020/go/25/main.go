// https://adventofcode.com/
// ProggyVector

package main

import (
	"fmt"
	"time"
)

func duration(invocation time.Time, name string) {
	fmt.Println(name, "duration", time.Since(invocation))
}

func report(expected, result int) {
	if expected != -1 {
		if result != expected {
			fmt.Println("ERROR: got", result, "expected", expected)
		} else {
			fmt.Println("RIGHT ANSWER:", result)
		}
	}
}

func part1(cardPublicKey, doorPublicKey, expected int) (result int) {
	defer duration(time.Now(), "part 1")

	var publicKey [2]int = [2]int{1, 1}
	var encryptionKey [2]int = [2]int{1, 1}
	var foundIdx = 0
	for {
		publicKey[0] = (publicKey[0] * 7) % 20201227
		publicKey[1] = (publicKey[1] * 7) % 20201227
		encryptionKey[0] = (encryptionKey[0] * doorPublicKey) % 20201227
		encryptionKey[1] = (encryptionKey[1] * cardPublicKey) % 20201227
		if publicKey[0] == cardPublicKey {
			break
		}
		if publicKey[1] == doorPublicKey {
			foundIdx = 1
			break
		}
	}
	result = encryptionKey[foundIdx]
	report(expected, result)
	return result
}

func main() {
	defer duration(time.Now(), "main")

	part1(8252394, 6269621, 181800)
}

/*
$ go run .
RIGHT ANSWER: 181800
part 1 duration 19.698738ms
main duration 19.706097ms
*/
