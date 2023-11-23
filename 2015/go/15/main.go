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

//go:embed "input.txt"
var input string

const (
	CAP = iota
	DUR
	FLA
	TEX
	CAL
)

const MAX_PROPS = 5

type Ingredient struct {
	name  string
	props [MAX_PROPS]int
}

var iList []Ingredient

// Duration of a func call
// Arguments to a defer statement are immediately evaluated and stored.
// The deferred function receives the pre-evaluated values when its invoked.
// usage: defer uDuration(time.Now(), "IntFactorial")
func duration(invocation time.Time, name string) {
	fmt.Println(name, time.Since(invocation))
}

func notneg(i int) int {
	if i < 0 {
		return 0
	}
	return i
}

func parseInput() {
	re := regexp.MustCompile(`(\w+): capacity ([-0-9]+), durability ([-0-9]+), flavor ([-0-9]+), texture ([-0-9]+), calories ([-0-9]+)`)
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		matches := re.FindAllStringSubmatchIndex(scanner.Text(), -1)
		if len(matches) != 1 {
			fmt.Println("no match")
			continue
		}
		if len(matches[0]) != 14 {
			fmt.Println("no sub matches")
			continue
		}
		// 0:1 is whole match
		var i Ingredient = Ingredient{}
		i.name = scanner.Text()[matches[0][2]:matches[0][3]]
		i.props[CAP], _ = strconv.Atoi(scanner.Text()[matches[0][4]:matches[0][5]])
		i.props[DUR], _ = strconv.Atoi(scanner.Text()[matches[0][6]:matches[0][7]])
		i.props[FLA], _ = strconv.Atoi(scanner.Text()[matches[0][8]:matches[0][9]])
		i.props[TEX], _ = strconv.Atoi(scanner.Text()[matches[0][10]:matches[0][11]])
		i.props[CAL], _ = strconv.Atoi(scanner.Text()[matches[0][12]:matches[0][13]])
		iList = append(iList, i)
		fmt.Println(i)
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("scanner error: %v\n", err)
	}
}

// permute produces all combinations of the ingredients
// with the amount of each ingredient ranging from 0 to 100
// the slice it returns is the same length as the number of ingredients
// hardwired to 2 ingredients for the moment
func permute2(c chan [2]int) {
	// c <- [2]int{44, 56}
	for i := 0; i <= 100; i++ {
		j := 100 - i
		c <- [2]int{i, j}
	}
	close(c)
}

func permute4(c chan [4]int) {
	for i := 0; i <= 100; i++ {
		for j := 0; j <= 100; j++ {
			for k := 0; k <= 100; k++ {
				l := 100 - i - j - k
				c <- [4]int{i, j, k, l}
			}
		}
	}
	close(c)
}

func solve() (int, int) {
	var topScore, topScore500 int
	c := make(chan [4]int)
	go permute4(c)
	for lst := range c {
		// for all the ingredient combinations
		// len(lst) == len(iList)
		// lst will be like [100,0]; 100 Butterscotch, 0 Cinnamon
		var subscores [MAX_PROPS]int = [MAX_PROPS]int{}
		for i, amt := range lst {
			ingredient := iList[i]
			// add up the amt * property-value for each property
			for prop := 0; prop < MAX_PROPS; prop++ {
				subscores[prop] += amt * ingredient.props[prop]
			}
		}
		// fmt.Println(subscores)
		// for part 1, ignore last prop (calories)
		var score int = notneg(subscores[CAP]) * notneg(subscores[DUR]) * notneg(subscores[FLA]) * notneg(subscores[TEX])
		if score > topScore {
			topScore = score
		}
		if subscores[CAL] == 500 {
			if score > topScore500 {
				topScore500 = score
			}
		}
	}
	return topScore, topScore500
}

func main() {
	defer duration(time.Now(), "main")
	parseInput()
	fmt.Println(solve()) // 13882464, 11171160
}
