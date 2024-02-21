// https://adventofcode.com/2021/day/16
package main

import (
	_ "embed"
	"fmt"
	"math"
	"runtime"
	"strconv"
	"strings"
	"time"
)

//go:embed input.txt
var input string

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

var rune2bin map[rune]string = map[rune]string{
	'0': "0000",
	'1': "0001",
	'2': "0010",
	'3': "0011",
	'4': "0100",
	'5': "0101",
	'6': "0110",
	'7': "0111",
	'8': "1000",
	'9': "1001",
	'A': "1010",
	'B': "1011",
	'C': "1100",
	'D': "1101",
	'E': "1110",
	'F': "1111",
}

// https://github.com/pscosta/aoc_2021/blob/main/go/day_16.go

type Packet struct {
	version    int
	typeID     int
	data       int
	subPackets []Packet
}

func hex2bin(s string) string {
	var sb strings.Builder
	for _, ch := range s {
		sb.WriteString(rune2bin[ch])
	}
	return sb.String()
}

func readInt(bits string, size int, pc *int) int {
	i, _ := strconv.ParseInt(bits[*pc:*pc+size], 2, 64)
	*pc += size
	return int(i)
}

func readLiteral(bits string, pc *int) Packet {
	var data string
	for prefix := -1; prefix != 0; *pc += 4 {
		prefix = readInt(bits, 1, pc)
		data += bits[*pc : *pc+4]
	}
	return Packet{data: readInt(data, len(data), new(int))}
}

func readLengthOperator(bits string, pc *int) Packet {
	var subPackets []Packet
	var subPacketLen int = readInt(bits, 15, pc)
	var initPc int = *pc
	for *pc-initPc < subPacketLen {
		subPackets = append(subPackets, readPacket(bits, pc))
	}
	return Packet{subPackets: subPackets}
}

func readCountOperator(bits string, pc *int) Packet {
	var subPackets []Packet
	var subPacketCount int = readInt(bits, 11, pc)
	for i := 0; i < subPacketCount; i++ {
		subPackets = append(subPackets, readPacket(bits, pc))
	}
	return Packet{subPackets: subPackets}
}

func readPacket(bits string, pc *int) (pkt Packet) {
	var version int = readInt(bits, 3, pc)
	var typeID int = readInt(bits, 3, pc)
	switch typeID {
	case 4:
		pkt = readLiteral(bits, pc)
	default:
		switch readInt(bits, 1, pc) {
		case 0:
			pkt = readLengthOperator(bits, pc)
		case 1:
			pkt = readCountOperator(bits, pc)
		}
	}
	pkt.version = version
	pkt.typeID = typeID
	return
}

func sumVersions(pkt Packet) (sum int) {
	for _, sp := range pkt.subPackets {
		sum += sumVersions(sp)
	}
	return pkt.version + sum
}

func part1(hex string, expected int) (result int) {
	defer duration(time.Now(), "part 1")

	var bits string = hex2bin(hex)
	// fmt.Println(len(hex), hex)
	// fmt.Println(len(bits), bits)

	var pc int
	var outerPacket Packet = readPacket(bits, &pc)

	result = sumVersions(outerPacket)

	report(expected, result)
	return result
}

/*
var thisOps map[int]func(Packet) int

func init() {
	thisOps = ops
}

var ops map[int]func(Packet) int = map[int]func(Packet) int{
	0: func(pkt Packet) int {
		var sum int
		for _, sp := range pkt.subPackets {
			sum += eval(sp)
		}
		return sum
	},
	1: func(pkt Packet) int {
		var prod int = 1
		for _, sp := range pkt.subPackets {
			prod *= eval(sp)
		}
		return prod
	},
	2: func(pkt Packet) int {
		var m int = math.MaxInt64
		for _, sp := range pkt.subPackets {
			m = min(m, eval(sp))
		}
		return m
	},
	3: func(pkt Packet) int {
		var m int = math.MaxInt64
		for _, sp := range pkt.subPackets {
			m = max(m, eval(sp))
		}
		return m
	},
	4: func(pkt Packet) int {
		return pkt.data
	},
	5: func(pkt Packet) int {
		if eval(pkt.subPackets[0]) > eval(pkt.subPackets[1]) {
			return 1
		} else {
			return 0
		}
	},
	6: func(pkt Packet) int {
		if eval(pkt.subPackets[0]) < eval(pkt.subPackets[1]) {
			return 1
		} else {
			return 0
		}
	},
	7: func(pkt Packet) int {
		if eval(pkt.subPackets[0]) == eval(pkt.subPackets[1]) {
			return 1
		} else {
			return 0
		}
	},
}
*/

func eval(pkt Packet) int {
	switch pkt.typeID {
	case 0:
		var sum int
		for _, sp := range pkt.subPackets {
			sum += eval(sp)
		}
		return sum

	case 1:
		var prod int = 1
		for _, sp := range pkt.subPackets {
			prod *= eval(sp)
		}
		return prod

	case 2:
		var m int = math.MaxInt64
		for _, sp := range pkt.subPackets {
			m = min(m, eval(sp))
		}
		return m

	case 3:
		var m int
		for _, sp := range pkt.subPackets {
			m = max(m, eval(sp))
		}
		return m

	case 4:
		return pkt.data

	case 5:
		if eval(pkt.subPackets[0]) > eval(pkt.subPackets[1]) {
			return 1
		} else {
			return 0
		}

	case 6:
		if eval(pkt.subPackets[0]) < eval(pkt.subPackets[1]) {
			return 1
		} else {
			return 0
		}
	case 7:
		if eval(pkt.subPackets[0]) == eval(pkt.subPackets[1]) {
			return 1
		} else {
			return 0
		}
	}

	return -1
}

func part2(hex string, expected int) (result int) {
	defer duration(time.Now(), "part 2")

	var bits string = hex2bin(hex)
	var pc int
	var outerPacket Packet = readPacket(bits, &pc)

	result = eval(outerPacket)

	report(expected, result)
	return result
}

func main() {
	defer duration(time.Now(), "main")

	// part1("8A004A801A8002F478", 16)
	// part1("620080001611562C8802118E34", 12)
	// part1("C0015000016115A2E0802F182340", 23)
	// part1("A0016C880162017C3686B18A3D4780", 31)

	part1(input, 974)

	// part2("C200B40A82", 3)
	// part2("04005AC33890", 54)
	// part2("880086C3E88112", 7)
	// part2("CE00C43D881120", 9)
	// part2("D8005AC2A8F0", 1)
	// part2("F600BC2D8F", 0)
	// part2("9C005AC2F8F0", 0)
	// part2("9C0141080250320F1802104A08", 1)

	part2(input, 180616437720)

	{
		var memStats runtime.MemStats
		runtime.ReadMemStats(&memStats)
		fmt.Printf("Heap memory (in bytes): %d\n", memStats.HeapAlloc)
		fmt.Printf("Number of garbage collections: %d\n", memStats.NumGC)
	}
}

/*
$ go run main.go
RIGHT ANSWER: 974
part 1 duration 153.965µs
RIGHT ANSWER: 180616437720
part 2 duration 107.714µs
Heap memory (in bytes): 245784
Number of garbage collections: 0
main duration 364.478µs
*/
