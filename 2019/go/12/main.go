// https://adventofcode.com/2019/day/12
package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"strings"
	"time"
)

//go:embed test1.txt
var test1 string

//go:embed test2.txt
var test2 string

//go:embed input.txt
var input string

type Vector struct {
	x, y, z int
}

func (v Vector) add(w Vector) Vector {
	return Vector{x: v.x + w.x, y: v.y + w.y, z: v.z + w.z}
}

func (v Vector) sumabs() int {
	return abs(v.x) + abs(v.y) + abs(v.z)
}

type Moon struct {
	pos, vel Vector
}

type Number interface {
	int | float32 | float64
}

func abs[T Number](n T) T {
	if n < 0 {
		return -n
	}
	return n
}

func duration(invocation time.Time, name string) {
	fmt.Println(name, time.Since(invocation))
}

func step(moons []Moon) {
	// gravity
	for i := 0; i < len(moons)-1; i++ {
		for j := i + 1; j < len(moons); j++ {
			if moons[i].pos.x > moons[j].pos.x {
				moons[i].vel.x -= 1
				moons[j].vel.x += 1
			} else if moons[i].pos.x < moons[j].pos.x {
				moons[i].vel.x += 1
				moons[j].vel.x -= 1
			}
			if moons[i].pos.y > moons[j].pos.y {
				moons[i].vel.y -= 1
				moons[j].vel.y += 1
			} else if moons[i].pos.y < moons[j].pos.y {
				moons[i].vel.y += 1
				moons[j].vel.y -= 1
			}
			if moons[i].pos.z > moons[j].pos.z {
				moons[i].vel.z -= 1
				moons[j].vel.z += 1
			} else if moons[i].pos.z < moons[j].pos.z {
				moons[i].vel.z += 1
				moons[j].vel.z -= 1
			}
		}
	}
	// velocity
	// nb don't use range 'cos it makes copy of args
	for i := 0; i < len(moons); i++ {
		moons[i].pos = moons[i].pos.add(moons[i].vel)
	}
}

func energy(moons []Moon) int {
	var e int
	for _, moon := range moons {
		e += moon.pos.sumabs() * moon.vel.sumabs()
	}
	return e
}

func part1(in string, steps int, expected int) {
	defer duration(time.Now(), "part 1")

	var moons []Moon
	scanner := bufio.NewScanner(strings.NewReader(in))
	for scanner.Scan() {
		var x, y, z int
		if n, err := fmt.Sscanf(scanner.Text(), "<x=%d, y=%d, z=%d>", &x, &y, &z); n != 3 {
			fmt.Println(err)
			break
		}
		moons = append(moons, Moon{pos: Vector{x: x, y: y, z: z}})
		// "the x, y, and z velocity of each moon starts at 0"
	}
	// fmt.Printf("0 %+v\n", moons[0])
	for i := 0; i < steps; i++ {
		step(moons)
		// fmt.Printf("%d %+v\n", i+1, moons[0])
	}

	var result int = energy(moons)
	if expected != -1 {
		if result != expected {
			fmt.Println("ERROR: got", result, "expected", expected)
		} else {
			fmt.Println("RIGHT ANSWER:", result)
		}
	}
}

// GCD returns the greatest common divisor of a and b
func GCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// LCM returns the least common multiple of a and b
func LCM(a, b int) int {
	return a / GCD(a, b) * b
}

// LCMList returns the least common multiple of a list of integers
func LCMList(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	lcm := nums[0]
	for _, n := range nums[1:] {
		lcm = LCM(lcm, n)
	}
	return lcm
}

func part2(in string, expected int) {
	defer duration(time.Now(), "part 2")

	// tried running part 2 by brute force;
	// machine ran out of memory and locked up
	// so, use the three-separate-runs-and-find-the-LCM trick

	var moons []Moon

	keyx := func() [8]int {
		var k [8]int
		for i := 0; i < len(moons); i++ {
			k[i] = moons[i].pos.x
			k[i+4] = moons[i].vel.x
		}
		return k
	}
	keyy := func() [8]int {
		var k [8]int
		for i := 0; i < len(moons); i++ {
			k[i] = moons[i].pos.y
			k[i+4] = moons[i].vel.y
		}
		return k
	}
	keyz := func() [8]int {
		var k [8]int
		for i := 0; i < len(moons); i++ {
			k[i] = moons[i].pos.z
			k[i+4] = moons[i].vel.z
		}
		return k
	}

	var keyx0, keyy0, keyz0 [8]int
	var steps, stepsx, stepsy, stepsz int

	// The instructions suggest that we should cycle until
	// reaching a previous state, but that previous state always seems
	// to be the initial state

	// because of this, we don't need to reset the moons slice
	// before the second and third test, but we do

	moons = nil
	scanner := bufio.NewScanner(strings.NewReader(in))
	for scanner.Scan() {
		var x, y, z int
		if n, err := fmt.Sscanf(scanner.Text(), "<x=%d, y=%d, z=%d>", &x, &y, &z); n != 3 {
			fmt.Println(err)
			break
		}
		moons = append(moons, Moon{pos: Vector{x: x, y: y, z: z}})
		// "the x, y, and z velocity of each moon starts at 0"
	}
	// var state map[[8]int]int = make(map[[8]int]int)
	// state[keyx()] = 0
	keyx0 = keyx()
	steps = 0
	for {
		step(moons)
		steps += 1
		mk := keyx()
		if mk == keyx0 {
			stepsx = steps
			break
		}
		// if s, ok := state[mk]; ok {
		// 	stepsx = steps - s
		// 	break
		// }
		// state[mk] = steps
	}

	moons = nil
	scanner = bufio.NewScanner(strings.NewReader(in))
	for scanner.Scan() {
		var x, y, z int
		if n, err := fmt.Sscanf(scanner.Text(), "<x=%d, y=%d, z=%d>", &x, &y, &z); n != 3 {
			fmt.Println(err)
			break
		}
		moons = append(moons, Moon{pos: Vector{x: x, y: y, z: z}})
		// "the x, y, and z velocity of each moon starts at 0"
	}
	// state = make(map[[8]int]int)
	// state[keyy()] = 0
	keyy0 = keyy()
	steps = 0
	for {
		step(moons)
		steps += 1
		mk := keyy()
		if mk == keyy0 {
			stepsy = steps
			break
		}
		// if s, ok := state[mk]; ok {
		// 	stepsy = steps - s
		// 	break
		// }
		// state[mk] = steps
	}

	moons = nil
	scanner = bufio.NewScanner(strings.NewReader(in))
	for scanner.Scan() {
		var x, y, z int
		if n, err := fmt.Sscanf(scanner.Text(), "<x=%d, y=%d, z=%d>", &x, &y, &z); n != 3 {
			fmt.Println(err)
			break
		}
		moons = append(moons, Moon{pos: Vector{x: x, y: y, z: z}})
		// "the x, y, and z velocity of each moon starts at 0"
	}
	// state = make(map[[8]int]int)
	// state[keyz()] = 0
	keyz0 = keyz()
	steps = 0
	for {
		step(moons)
		steps += 1
		mk := keyz()
		if mk == keyz0 {
			stepsz = steps
			break
		}
		// if s, ok := state[mk]; ok {
		// 	stepsz = steps - s
		// 	break
		// }
		// state[mk] = steps
	}

	fmt.Println(stepsx, stepsy, stepsz)

	var result int = LCMList([]int{stepsx, stepsy, stepsz})
	if expected != -1 {
		if result != expected {
			fmt.Println("ERROR: got", result, "expected", expected)
		} else {
			fmt.Println("RIGHT ANSWER:", result)
		}
	}
}

func main() {
	defer duration(time.Now(), "main")

	// part1(test1, 10, 179)
	// part1(test2, 100, 1940)
	part1(input, 1000, 7636)

	// part2(test1, 2772)
	// part2(test2, 4686774924)
	part2(input, 281691380235984)
}

/*
$ go run main.go
RIGHT ANSWER: 7636
part 1 133.944µs
161428 193052 144624
RIGHT ANSWER: 281691380235984
part 2 27.168485ms
main 27.308296ms
*/
