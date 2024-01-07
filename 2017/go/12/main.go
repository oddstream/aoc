// https://adventofcode.com/2017/day/12 Digital Plumber
package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"slices"
	"strconv"
	"strings"
	"time"
)

//go:embed input.txt
var input string

func duration(invocation time.Time, name string) {
	fmt.Println(name, time.Since(invocation))
}

type IntSet map[int]struct{}

func (is IntSet) contains(n int) bool {
	var ok bool
	_, ok = is[n]
	return ok
}

func (is IntSet) add(n int) {
	is[n] = struct{}{}
}

// slow way of generating a key for a map
func (is IntSet) key() string {
	var arr []int = make([]int, len(is))
	for num := range is {
		arr = append(arr, num)
	}
	slices.Sort(arr)
	var sb strings.Builder
	for _, v := range arr {
		if v != 0 {
			sb.WriteString(strconv.Itoa(v))
			sb.WriteByte(',')
		}
	}
	return sb.String()
	// b, _ := json.Marshal(is) // Marshal sorts keys, apprently ...
	// return string(b)         // ... but this is much slower
}

// or we could use K&R p61
// and add leading/trailing non-digit ignoring
func atoi(s string) int {
	if n, err := strconv.Atoi(s); err == nil {
		return n
	} else {
		fmt.Println(err)
	}
	return 0
}

func loadConexMap() map[int][]int {
	var conex map[int][]int = make(map[int][]int)
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		tokens := strings.Split(scanner.Text(), " ")
		p0 := atoi(strings.TrimSuffix(tokens[0], ","))
		// tokens[1] == "<->"
		var plst []int
		for i := 2; i < len(tokens); i++ {
			plst = append(plst, atoi(strings.TrimSuffix(tokens[i], ",")))
		}
		conex[p0] = plst
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	return conex
}

func findGroup(conex map[int][]int, seed int) IntSet {
	var cset IntSet = make(IntSet)

	var add func(int)
	add = func(n int) {
		cset.add(n)
		for _, ch := range conex[n] {
			if !cset.contains(ch) {
				add(ch)
			}
		}
	}

	add(seed)

	return cset
}

func partOne() int {
	defer duration(time.Now(), "part 1")
	return len(findGroup(loadConexMap(), 0))
}

func partTwo() int {
	defer duration(time.Now(), "part 2")
	conex := loadConexMap()
	gmap := make(map[string]struct{})
	for n := range conex {
		g := findGroup(conex, n)
		gmap[g.key()] = struct{}{}
	}
	return len(gmap)
}

func main() {
	defer duration(time.Now(), "main")

	fmt.Println(partOne()) // 115
	fmt.Println(partTwo()) // 221
}

/*
part 1 811.87µs
115
part 2 27.696523ms
221
main 28.539215ms
*/
