// https://adventofcode.com/2018/day/16 Chronal Classification
package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"strings"
	"time"
)

//go:embed input1.txt
var input1 string

//go:embed input2.txt
var input2 string

type Registers [4]int // four registers, 0 .. 3

type Instruction struct {
	N, A, B, C int // N is opcode (0..15), A, B are inputs, C is output
}

type OpFunc func(int, int, Registers) int // output always goes to C

type Sample struct {
	before Registers
	i      Instruction
	after  Registers
}

var ops map[string]OpFunc = map[string]OpFunc{
	"addr": func(A, B int, r Registers) int { return r[A] + r[B] },
	"addi": func(A, B int, r Registers) int { return r[A] + B },
	"mulr": func(A, B int, r Registers) int { return r[A] * r[B] },
	"muli": func(A, B int, r Registers) int { return r[A] * B },
	"banr": func(A, B int, r Registers) int { return r[A] & r[B] },
	"bani": func(A, B int, r Registers) int { return r[A] & B },
	"borr": func(A, B int, r Registers) int { return r[A] | r[B] },
	"bori": func(A, B int, r Registers) int { return r[A] | B },
	"setr": func(A, B int, r Registers) int { return r[A] },
	"seti": func(A, B int, r Registers) int { return A },
	"gtir": func(A, B int, r Registers) int {
		if A > r[B] {
			return 1
		} else {
			return 0
		}
	},
	"gtri": func(A, B int, r Registers) int {
		if r[A] > B {
			return 1
		} else {
			return 0
		}
	},
	"gtrr": func(A, B int, r Registers) int {
		if r[A] > r[B] {
			return 1
		} else {
			return 0
		}
	},
	"eqir": func(A, B int, r Registers) int {
		if A == r[B] {
			return 1
		} else {
			return 0
		}
	},
	"eqri": func(A, B int, r Registers) int {
		if r[A] == B {
			return 1
		} else {
			return 0
		}
	},
	"eqrr": func(A, B int, r Registers) int {
		if r[A] == r[B] {
			return 1
		} else {
			return 0
		}
	},
}

func duration(invocation time.Time, name string) {
	fmt.Println(name, time.Since(invocation))
}

/*
               111111
op   0123456789012345
addr       1     1  1
mulr       1
bani 1    11 1 11
gtri 1  11 1111 1
addi     1 1
borr       1     1
eqrr 11 1 1  11 1 11
banr 1   111   11
bori     1 1   1 1
gtir 1 11 11 1  1
eqri 11 1 1 11  1 1
muli      11
setr   1 1 1   1 1
seti 1     1        1
gtrr 1  1 1111  11
eqir 11 1 1 111 1

00 bani gtri banr gtir eqri seti gtrr eqir
01 eqri eqir
02 gtir setr
03 gtri gtir eqri gtrr eqir
04 gtri addi banr bori setr
05 bani banr gtir eqri muli gtrr eqir
06 addr mulr bani gtri addi borr banr bori gtir muli setr seti gtrr
07 gtri eqri gtrr eqir
08 bani gtri gtir eqri, gtrr eqir
09 gtri eqir
10 bani banr bori setr
11 bani banr gtir eqri gtrr
12 addr borr bori setr gtrr
13 eqri
14 eqrr
15 addr seti

14 eqrr
13 eqri
01 eqir
09 gtri
07 gtrr
03 gtir
02 setr
08 bani
11 banr
10 bori
00 seti
04 addi
15 addr
12 borr
05 muli
06 mulr
*/

var num2func map[int]OpFunc = map[int]OpFunc{
	// 14: ops["eqrr"],
	// 13: ops["eqri"],
	// 1:  ops["eqir"],
	// 9:  ops["gtri"],
	// 7:  ops["gtrr"],
	// 3:  ops["gtir"],
	// 2:  ops["setr"],
	// 8:  ops["bani"],
	// 11: ops["banr"],
	// 10: ops["bori"],
	// 0:  ops["seti"],
	// 4:  ops["addi"],
	// 15: ops["addr"],
	// 12: ops["borr"],
	// 5:  ops["muli"],
	// 6:  ops["mulr"],
}

func display(possible map[string]*[16]bool) {
	fmt.Println("               111111")
	fmt.Println("op   0123456789012345")
	for opname := range ops {
		poss := possible[opname]
		fmt.Print(opname, " ")
		for i := 0; i < len(poss); i++ {
			if poss[i] {
				fmt.Print("1")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func partOne(input string, expected int) int {
	defer duration(time.Now(), "part 1")

	var samples []Sample
	var before, after Registers
	var i Instruction

	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			samples = append(samples, Sample{before, i, after})
			before = Registers{}
			i = Instruction{}
			after = Registers{}
		} else {
			if n, _ := fmt.Sscanf(line, "Before: [%d, %d, %d, %d]", &before[0], &before[1], &before[2], &before[3]); n == 4 {
				//fmt.Println(before)
			} else if n, _ := fmt.Sscanf(line, "After: [%d, %d, %d, %d]", &after[0], &after[1], &after[2], &after[3]); n == 4 {
				//fmt.Println(after)
			} else if n, _ := fmt.Sscanf(line, "%d %d %d %d", &i.N, &i.A, &i.B, &i.C); n == 4 {
				//fmt.Println(i)
			} else {
				fmt.Println("ERROR: cannot parse", line)
			}
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	fmt.Println(len(ops), "ops")         // 16 ops
	fmt.Println(len(samples), "samples") // 806 samples

	// map the op name to pointer to list of len(ops) bools
	var possible map[string]*[16]bool = make(map[string]*[16]bool)
	for opname := range ops {
		possible[opname] = &[16]bool{true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true}
	}

	var result int
	for _, sample := range samples {
		var valid int
		for opname, f := range ops {
			res := f(sample.i.A, sample.i.B, sample.before)
			// res is in the range 0 .. 9
			// i.C says which register to put result in
			// i.C will always be 0..3
			if sample.after[sample.i.C] == res {
				valid += 1
			} else {
				possible[opname][sample.i.N] = false
			}
		}
		if valid >= 3 {
			result += 1
		}
	}

	// use output from display(possible)
	// to create (by hand, Bletchley Park-style)
	// map of opcode number (via opcode string) to op func
	// display(possible)

	var occlst [][]string
	for i := 0; i < 16; i++ {
		occlst = append(occlst, []string{})
		for k := range ops {
			if possible[k][i] {
				occlst[i] = append(occlst[i], k)
			}
		}
	}
	// display occlst
	// for i := 0; i < 16; i++ {
	// 	fmt.Print(i)
	// 	for _, v := range occlst[i] {
	// 		fmt.Print(" ", v)
	// 	}
	// 	fmt.Println()
	// }

	// var num2name map[int]string = map[int]string{}

	for {
		for i, lst1 := range occlst {
			if len(lst1) == 1 {
				opname := lst1[0]
				// num2name[i] = opname
				num2func[i] = ops[opname]
				for j, lst2 := range occlst {
					for k, name := range lst2 {
						if name == opname {
							occlst[j] = append(lst2[:k], lst2[k+1:]...) // replace list
						}
					}
				}
				goto next // restart outer loop to find another entry with len == 1
			}
		}
		break // we didn't find any entries with len == 1, so we're done
		// could also have outer loop run until len(ops) == 16
	next:
	}
	// fmt.Println(num2name)

	if expected != -1 {
		if result != expected {
			fmt.Println("ERROR: got", result, "expected", expected)
		} else {
			fmt.Println("RIGHT:", result)
		}
	}
	return result
}

func partTwo(input string, expected int) int {
	defer duration(time.Now(), "part 1")

	var r Registers
	var i Instruction
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		if n, err := fmt.Sscanf(scanner.Text(), "%d %d %d %d", &i.N, &i.A, &i.B, &i.C); n != 4 {
			fmt.Println(err)
			break
		}
		f := num2func[i.N]
		r[i.C] = f(i.A, i.B, r)
	}

	var result int = r[0]

	if expected != -1 {
		if result != expected {
			fmt.Println("ERROR: got", result, "expected", expected)
		} else {
			fmt.Println("RIGHT:", result)
		}
	}
	return result
}

func main() {
	defer duration(time.Now(), "main")

	partOne(input1, 646)
	partTwo(input2, 681)
}

/*
$ go run main.go
16 ops
806 samples
RIGHT: 646
part 1 3.982451ms
RIGHT: 681
part 1 806.324µs
main 4.809792ms
*/
