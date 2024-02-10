// https://adventofcode.com/2020/day/21
package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"regexp"
	"runtime"
	"strings"
	"time"
)

//go:embed test1.txt
var test1 string

//go:embed input.txt
var input string

func duration(invocation time.Time, name string) {
	fmt.Println(name, time.Since(invocation))
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

func Contains[T comparable](elems []T, v T) bool {
	for _, s := range elems {
		if v == s {
			return true
		}
	}
	return false
}

func Union[T comparable](a []T, b []T) []T {
	var result []T = a
	for _, elem := range b {
		if !Contains(a, elem) {
			result = append(result, elem)
		}
	}
	return result
}

// brute force (inefficient) method O(m∗n)
func Intersection[T comparable](a, b []T) []T {
	var result []T
	for _, elem := range a {
		if Contains(b, elem) {
			result = append(result, elem)
		}
	}
	return result
}

func part1(in string, expected int) (result int) {
	defer duration(time.Now(), "part 1")

	var allAllergens map[string]struct{} = make(map[string]struct{})
	var allIngredients map[string]struct{} = make(map[string]struct{})
	// foods is indexed by allergen and has a list of list of ingredients
	// dairy
	// 	 0 mxmxvkd kfcds sqjhc nhms
	// 	 1 trh fvjkl sbzzf mxmxvkd
	// fish
	// 	 0 mxmxvkd kfcds sqjhc nhms
	// 	 1 sqjhc mxmxvkd sbzzf
	// soy
	// 	 0 sqjhc fvjkl
	var foods map[string][][]string = make(map[string][][]string)

	// potentials is indexed by ingredient and has a set of potential allergens
	// sqjhc
	// 	 soy
	// 	 fish
	var potentials map[string]map[string]struct{} = make(map[string]map[string]struct{})

	rxAllergens := regexp.MustCompile(`([a-z ]+) \(contains ([a-z, ]+)\)`)
	scanner := bufio.NewScanner(strings.NewReader(in))
	for scanner.Scan() {
		matches := rxAllergens.FindAllStringSubmatch(scanner.Text(), -1)
		for _, allergen := range strings.Split(strings.ReplaceAll(matches[0][2], ",", ""), " ") {
			allAllergens[allergen] = struct{}{}

			foods[allergen] = append(foods[allergen], strings.Split(matches[0][1], " "))
		}
		for _, ingredient := range strings.Split(matches[0][1], " ") {
			allIngredients[ingredient] = struct{}{}

			if _, ok := potentials[ingredient]; !ok {
				potentials[ingredient] = make(map[string]struct{})
			}
			for _, allergen := range strings.Split(strings.ReplaceAll(matches[0][2], ",", ""), " ") {
				potentials[ingredient][allergen] = struct{}{}
			}
		}
	}
	fmt.Println("all ingredients")
	fmt.Println(len(allIngredients), ":", allIngredients)

	fmt.Println("all allergens")
	fmt.Println(len(allAllergens), ":", allAllergens)

	fmt.Println("foods")
	for k, v := range foods {
		fmt.Println(k)
		for _, l := range v {
			fmt.Println("\t", l)
		}
	}

	// fmt.Println("potentials")
	// for k, v := range potentials {
	// 	fmt.Println(k)
	// 	for i := range v {
	// 		fmt.Println("\t", i)
	// 	}
	// }

	var allergens map[string]struct{} = make(map[string]struct{})
	for _, lstlst := range foods {
		var inter []string
		for _, lst := range lstlst {
			if len(inter) == 0 {
				inter = append(inter, lst...)
			} else {
				inter = Intersection(inter, lst)
			}
		}
		for _, item := range inter {
			allergens[item] = struct{}{}
		}
	}
	fmt.Println("allergens", len(allergens), allergens)

	var safer map[string]struct{} = make(map[string]struct{})
	for ingredient := range allIngredients {
		if _, ok := allergens[ingredient]; !ok {
			safer[ingredient] = struct{}{}
		}
	}

	fmt.Println("safer", len(safer), safer)

	scanner = bufio.NewScanner(strings.NewReader(in))
	for scanner.Scan() {
		matches := rxAllergens.FindAllStringSubmatch(scanner.Text(), -1)
		for _, ingredient := range strings.Split(matches[0][1], " ") {
			if _, ok := safer[ingredient]; ok {
				result += 1
			}
		}
	}
	report(expected, result)

	fmt.Println("foods without safer")
	for allergen, lstlst := range foods {
		for _, lst := range lstlst {
			fmt.Print(allergen, ":")
			for _, ingredient := range lst {
				if _, ok := safer[ingredient]; !ok {
					fmt.Print(" ", ingredient)
				}
			}
			fmt.Println()
		}
	}
	return result
}

func main() {
	defer duration(time.Now(), "main")

	part1(test1, 5)
	// part1(input, 2659)
	//	part2(test1, "mxmxvkd,sqjhc,fvjkl")
	//	part2(input, "rcqb,cltx,nrl,qjvvcvz,tsqpn,xhnk,tfqsb,zqzmzl")

	{
		var memStats runtime.MemStats
		runtime.ReadMemStats(&memStats)
		fmt.Printf("Heap memory (in bytes): %d\n", memStats.HeapAlloc)
		fmt.Printf("Number of garbage collections: %d\n", memStats.NumGC)
	}
}

/*
$ go run main.go
*/
