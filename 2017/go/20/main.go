// https://adventofcode.com/2017/day/20 Particle Swarm
package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"sort"
	"strings"
	"time"
)

//go:embed input.txt
var input string

type Vector struct {
	x, y, z int
}

func (v Vector) manhatten() int {
	// sum, not product, you eejit
	return abs(v.x) + abs(v.y) + abs(v.z)
}

type Particle struct {
	pos, vel, acc Vector
	n             int
	deleted       bool
}

func duration(invocation time.Time, name string) {
	fmt.Println(name, time.Since(invocation))
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func partOne() int {
	defer duration(time.Now(), "part 1")
	var particles []Particle = make([]Particle, 0, 1002)
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		var p Particle = Particle{n: len(particles)}
		if n, err := fmt.Sscanf(scanner.Text(), "p=<%d,%d,%d>, v=<%d,%d,%d>, a=<%d,%d,%d>",
			&p.pos.x, &p.pos.y, &p.pos.z,
			&p.vel.x, &p.vel.y, &p.vel.z,
			&p.acc.x, &p.acc.y, &p.acc.z); n != 9 {
			fmt.Println(err)
			break
		}
		particles = append(particles, p)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	// rather than run the loop,
	// just sort by acc, vel, pos
	// and take the first particle
	sort.Slice(particles, func(a, b int) bool {
		var acca int = particles[a].acc.manhatten()
		var accb int = particles[b].acc.manhatten()
		if acca == accb {
			var vela int = particles[a].vel.manhatten()
			var velb int = particles[b].vel.manhatten()
			if vela == velb {
				var posa int = particles[a].pos.manhatten()
				var posb int = particles[b].pos.manhatten()
				return posa < posb
			} else {
				return vela < velb
			}
		} else {
			return acca < accb
		}
	})
	return particles[0].n
}

func partTwo() int {
	defer duration(time.Now(), "part 2")

	var particles []Particle = make([]Particle, 0, 1024)
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		var p Particle = Particle{n: len(particles), deleted: false}
		if n, err := fmt.Sscanf(scanner.Text(), "p=<%d,%d,%d>, v=<%d,%d,%d>, a=<%d,%d,%d>",
			&p.pos.x, &p.pos.y, &p.pos.z,
			&p.vel.x, &p.vel.y, &p.vel.z,
			&p.acc.x, &p.acc.y, &p.acc.z); n != 9 {
			fmt.Println(err)
			break
		}
		particles = append(particles, p)
	}

	{
		// check for collisions at start
		// Vector makes a nice map key
		// because structs (with all fields comparable) are comparable
		var posmap map[Vector]int = make(map[Vector]int)
		for i := 0; i < len(particles); i++ {
			posmap[particles[i].pos] += 1
		}
		for v, count := range posmap {
			if count != 1 {
				fmt.Println("start collide at", v)
			}
		}
	}

	for cycles := 0; cycles < 39; cycles++ { // 39 found by trial and error
		var posmap map[Vector]int = make(map[Vector]int)
		// of course, range creates a copy of the particle,
		// so we have to use indexes, which look ugly
		// but hopefuly gets optimized
		for i := 0; i < len(particles); i++ {
			if particles[i].deleted {
				continue
			}
			particles[i].vel.x += particles[i].acc.x
			particles[i].vel.y += particles[i].acc.y
			particles[i].vel.z += particles[i].acc.z
			particles[i].pos.x += particles[i].vel.x
			particles[i].pos.y += particles[i].vel.y
			particles[i].pos.z += particles[i].vel.z

			posmap[particles[i].pos] += 1
		}
		// flag as deleted any pos that occured more than once
		for v, count := range posmap {
			if count != 1 {
				for j := 0; j < len(particles); j++ {
					if particles[j].pos == v {
						particles[j].deleted = true
					}
				}
			}
		}
	}
	var remaining int
	for i := 0; i < len(particles); i++ {
		if !particles[i].deleted {
			remaining += 1
		}
	}
	return remaining
}

func main() {
	defer duration(time.Now(), "main")

	fmt.Println(partOne()) // 144
	fmt.Println(partTwo()) // 477
}

/*
$ go run main.go
part 1 2.977368ms
144
part 2 7.188625ms
477
main 10.196893ms
*/
