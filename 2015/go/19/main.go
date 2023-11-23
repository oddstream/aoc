package main

import (
	"bufio"
	_ "embed"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"regexp"
	"strings"
	"time"
)

//go:embed input.txt
var input string

type Subst struct {
	src, dst string
}

var subs []Subst = []Subst{}
var rev_subs map[string]string = map[string]string{}

var molecule string

func duration(invocation time.Time, name string) {
	fmt.Println(name, time.Since(invocation))
}

// replace the first src with dst in s, starting from s[pos]
// return new string and pos (which can be used to further call replace)
func replaceFirstInstance(s, src, dst string, idx int) (string, int) {
	var left string = s[:idx]
	var right string = s[idx:]
	var i int = strings.Index(right, src)
	if i == -1 {
		return s, -1
	}
	right = strings.Replace(right, src, dst, 1)
	return left + right, idx + i
}

func parseInput() {
	re := regexp.MustCompile(`(\w+) => (\w+)`)
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		matches := re.FindAllStringSubmatchIndex(line, -1)
		src := line[matches[0][2]:matches[0][3]]
		dst := line[matches[0][4]:matches[0][5]]
		s := Subst{
			src: src,
			dst: dst,
		}
		subs = append(subs, s)
		if rev_subs[dst] != "" {
			log.Println("rev_sub duplicate", dst)
		}
		rev_subs[dst] = src
	}
	if scanner.Scan() {
		molecule = scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		log.Println("scanner error", err)
	}
}

// do one pass through the machine with one input
// generate a set of outputs
func onepass(molecule string) map[string]struct{} {
	var results map[string]struct{} = make(map[string]struct{})
	for _, s := range subs {
		mol := molecule
		for i := 0; i < len(molecule); i++ {
			mol, j := replaceFirstInstance(mol, s.src, s.dst, i)
			if j == -1 {
				break
			}
			results[mol] = struct{}{}
			i = j
		}
	}
	return results
}

// first in-memory attempt used up all memory and exited with code -9
func multipass(input map[string]struct{}) map[string]struct{} {
	var results map[string]struct{} = make(map[string]struct{})
	for molecule := range input {
		for _, s := range subs {
			mol := molecule
			for i := 0; i < len(molecule); i++ {
				mol, j := replaceFirstInstance(mol, s.src, s.dst, i)
				if j == -1 {
					break
				}
				results[mol] = struct{}{}
				i = j
			}
		}
	}
	return results
}

// read in-file one line at a time, write results to out-file
// by step 9 was creating multi-GByte files and slowing
func multipass2(step int) {
	infile, err := os.Open(fmt.Sprintf("results%d.txt", step))
	if err != nil {
		log.Fatal(err)
	}
	defer infile.Close()

	outfile, err := os.Create(fmt.Sprintf("results%d.txt", step+1))
	if err != nil {
		log.Fatal(err)
	}
	defer outfile.Close()

	scanner := bufio.NewScanner(infile)
	for scanner.Scan() {
		results := onepass(scanner.Text())
		for line := range results {
			fmt.Fprintln(outfile, line)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func part1() {
	results := onepass(molecule)
	log.Println(len(results)) // part 1 576
}

func part2() {
	// make a list of all things we search for
	// so order can be shuffled
	// if we get stuck
	var reps []string = []string{}
	for thing := range rev_subs {
		reps = append(reps, thing)
	}
	// rand.New(rand.NewSource(1))
	rand.Shuffle(len(reps), func(i, j int) { reps[i], reps[j] = reps[j], reps[i] })
	// log.Println("reps", reps)

	var target string = molecule
	var part2 int = 0
	var shuffles int = 0
	for target != "e" {
		var tmp string = target
		for _, a := range reps {
			b := rev_subs[a]
			if strings.Count(target, a) == 0 {
				continue
			}
			target = strings.Replace(target, a, b, 1)
			part2 += 1
		}
		if tmp == target {
			// nothing happened, so shuffle replace order and try again
			target = molecule
			part2 = 0
			rand.Shuffle(len(reps), func(i, j int) { reps[i], reps[j] = reps[j], reps[i] })
			shuffles += 1
		}
	}
	log.Println("part 2", part2, "(", shuffles, "shuffles)") // 207, 117-4000 shuffles
}

func main() {
	defer duration(time.Now(), "main")

	var part int
	flag.IntVar(&part, "part", 2, "1 or 2")
	flag.Parse()

	parseInput()
	// for i, s := range substs {
	// 	log.Println(i, s.src, s.dst)
	// }
	// log.Println(molecule)
	// log.Println("rev_subs", rev_subs)

	if part == 1 {
		part1()
	} else if part == 2 {
		// turn "e" into molecule in the fewest number of "steps"
		// assume one pass through the replacement list is a "step"
		// each step src has a number of dst, which creates branches
		// https://markheath.net/post/advent-of-code-day19
		// https://www.reddit.com/r/adventofcode/comments/3xflz8/day_19_solutions/
		// start with
		part2()
	}
	// log.Println(replaceFirstInstance("abba", "b", "x", 0))
}
