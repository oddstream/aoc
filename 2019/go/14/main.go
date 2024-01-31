// https://adventofcode.com/2019/day/14
package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"runtime"
	"strings"
	"time"
)

//go:embed test1.txt
var test1 string

//go:embed test2.txt
var test2 string

//go:embed test3.txt
var test3 string

//go:embed input.txt
var input string

type Ingredient struct {
	quantity int
	chemical string
}

func str2Ingredient(s string) Ingredient {
	var ing Ingredient
	if n, err := fmt.Sscanf(s, "%d %s", &ing.quantity, &ing.chemical); n != 2 {
		fmt.Println(err)
		return Ingredient{}
	}
	return ing
}

type Recipe struct {
	output Ingredient
	inputs []Ingredient
}

func duration(invocation time.Time, name string) {
	fmt.Println(name, time.Since(invocation))
}

func loadMenu(in string) map[string]Recipe {
	var menu map[string]Recipe = make(map[string]Recipe)

	scanner := bufio.NewScanner(strings.NewReader(in))
	for scanner.Scan() {
		sides := strings.Split(scanner.Text(), " => ")

		var rhs Ingredient = str2Ingredient(sides[1])

		var lhs []Ingredient
		tokens := strings.Split(sides[0], ", ")
		for _, token := range tokens {
			lhs = append(lhs, str2Ingredient(token))
		}

		var r Recipe = Recipe{
			output: rhs,
			inputs: lhs,
		}

		menu[r.output.chemical] = r

	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	return menu
}

// https://github.com/lynerist/Advent-of-code-2019-golang/blob/master/Day_14/fourteen_a.go

func ore(menu map[string]Recipe, chemical string, quantity int, leftovers map[string]int) int {
	// fmt.Println(chemical)
	if chemical == "ORE" {
		return quantity
	}
	if leftovers[chemical] > 0 {
		if leftovers[chemical] >= quantity {
			leftovers[chemical] -= quantity
			return 0 // we didn't need any ore
		}
		quantity -= leftovers[chemical]
		leftovers[chemical] = 0
		return ore(menu, chemical, quantity, leftovers)
	}

	quantityProduced := menu[chemical].output.quantity
	reactionsNeeded := (quantity-1)/quantityProduced + 1
	oreNeeded := 0
	for _, inp := range menu[chemical].inputs {
		oreNeeded += ore(menu, inp.chemical, inp.quantity*reactionsNeeded, leftovers)
	}

	if reactionsNeeded*quantityProduced-quantity > 0 {
		leftovers[chemical] += reactionsNeeded*quantityProduced - quantity
	}
	return oreNeeded
}

func part1(in string, expected int) {
	defer duration(time.Now(), "part 1")

	m := loadMenu(in)
	var leftovers map[string]int = make(map[string]int)
	var result int = ore(m, "FUEL", 1, leftovers)

	if expected != -1 {
		if result != expected {
			fmt.Println("ERROR: got", result, "expected", expected)
		} else {
			fmt.Println("RIGHT ANSWER:", result)
		}
	}
}

func part2(in string, expected int) {
	defer duration(time.Now(), "part 2")

	m := loadMenu(in)
	var leftovers map[string]int = make(map[string]int)
	const TRILLION int = 1000000000000
	const STOCK int = 1000 // small as possible for speed
	var fuel, lastStock int
	for oreConsumed := 0; oreConsumed < TRILLION; fuel++ {
		lastStock = TRILLION - oreConsumed
		oreConsumed += ore(m, "FUEL", STOCK, leftovers)
	}
	fuel = (fuel - 1) * STOCK
	for oreConsumed := 0; oreConsumed < lastStock; fuel++ {
		oreConsumed += ore(m, "FUEL", 1, leftovers)
	}
	var result int = fuel - 1
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

	// part1(test1, 31)
	// part1(test2, 165)
	part1(input, 899155)
	// part2(test3, 82892753)
	part2(input, 2390226)

	{
		var memStats runtime.MemStats
		runtime.ReadMemStats(&memStats)
		// fmt.Printf("Total allocated memory: %d\n", memStats.Alloc)
		fmt.Printf("Heap memory: %d\n", memStats.HeapAlloc)
		fmt.Printf("Number of GCs: %d\n", memStats.NumGC)
	}
}

/*
$ go run main.go
*/
