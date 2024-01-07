package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

//go:embed input.txt
var input string

func duration(invocation time.Time, name string) {
	fmt.Println(name, time.Since(invocation))
}

// or we could use K&R p61
func atoi(s string) int {
	if n, err := strconv.Atoi(s); err == nil {
		return n
	} else {
		fmt.Println(err)
	}
	return 0
}

type Layer struct {
	depth, scanner, dir int
}

func loadInput() []Layer {
	re := regexp.MustCompile("([[:digit:]]+): ([[:digit:]]+)")
	entries := make(map[int]int)
	maxk := 0
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		matches := re.FindAllStringSubmatch(scanner.Text(), -1)
		// matches[0][0] is whole string
		k := atoi(matches[0][1])
		v := atoi(matches[0][2])
		entries[k] = v
		if k > maxk {
			maxk = k
		}
	}
	var layers []Layer = make([]Layer, maxk+1)
	for i := 0; i < maxk+1; i++ {
		if v, ok := entries[i]; ok {
			layers[i].depth = v
		}
		// layers not in entries will have depth == 0
		// all scanner will be == 0
		layers[i].dir = 1 // start by going down
	}
	return layers
}

func move(layers []Layer) {
	for layer := 0; layer < len(layers); layer++ {
		if layers[layer].depth > 0 {
			switch layers[layer].dir {
			case -1:
				if layers[layer].scanner == 0 {
					layers[layer].scanner = 1
					layers[layer].dir = 1
				} else {
					layers[layer].scanner -= 1
				}
			case 1:
				if layers[layer].scanner == layers[layer].depth-1 {
					layers[layer].scanner -= 1
					layers[layer].dir = -1
				} else {
					layers[layer].scanner += 1
				}
			}
		}
	}
}

func run1(layers []Layer) (int, int) {
	catches := 0
	severity := 0
	// fmt.Println("-", layers)

	for pp := 0; pp < len(layers); pp++ { // packet position
		if layers[pp].depth > 0 {
			if layers[pp].scanner == 0 {
				// fmt.Println("caught on layer", pp)
				catches += 1
				severity += pp * layers[pp].depth
			}
		}
		move(layers)
		// fmt.Println(pp, layers)
	}
	return catches, severity
}

func run2(layers []Layer, delay int) bool {
	for delay != 0 {
		move(layers)
		delay -= 1
	}
	for pp := 0; pp < len(layers); pp++ { // packet position
		if layers[pp].depth > 0 {
			if layers[pp].scanner == 0 {
				// fmt.Println("caught on layer", pp)
				return false
			}
		}
		move(layers)
		// fmt.Println(pp, layers)
	}
	return true
}

func partOne() int {
	defer duration(time.Now(), "part 1")

	layers := loadInput()
	_, severity := run1(layers)
	return severity
}

func partTwo() int {
	defer duration(time.Now(), "part 2")

	layers := loadInput()
	var delay int
	for delay = 3878060; delay < 3878064; delay++ {
		if run2(layers, delay) {
			break
		}
		for i := 0; i < len(layers); i++ {
			layers[i].scanner = 0
			layers[i].dir = 1
		}
	}
	return delay
}

func partTwoPy() {
	defer duration(time.Now(), "part 2 py")

	re := regexp.MustCompile("([[:digit:]]+): ([[:digit:]]+)")
	valDict := make(map[int]int)
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		matches := re.FindAllStringSubmatch(scanner.Text(), -1)
		// matches[0][0] is whole string
		k := atoi(matches[0][1])
		v := atoi(matches[0][2])
		valDict[k] = v
	}
	caught := false
	for delay := 0; delay < 4000000; delay++ {
		caught = false
		for i := range valDict {
			if (i+delay)%(2*valDict[i]-2) == 0 {
				caught = true
				break
			}
		}
		if !caught {
			fmt.Println(delay)
			break
		}
	}
}

func main() {
	defer duration(time.Now(), "main")

	fmt.Println(partOne()) // 3184
	partTwoPy()            // 3878062
}

/*
$ go run main.go
part 1 124.524µs
3184
3878062
part 2 py 818.034376ms
main 818.179514ms
*/
