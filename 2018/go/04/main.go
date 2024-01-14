// https://adventofcode.com/2018/day/4 Repose Record
package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

//go:embed input.txt
var input string

type Row struct {
	date string
	id   int
	min  [60]int
}

func (r Row) sleepMins() int {
	var n int
	for i := 0; i < 60; i++ {
		n += r.min[i]
	}
	return n
}

func sleepiestGuard(log []Row) int {
	var guards map[int]int = make(map[int]int)
	for _, row := range log {
		guards[row.id] += row.sleepMins()
	}
	// fmt.Println(guards)
	var max, sleepiest int
	for id, mins := range guards {
		if mins > max {
			max = mins
			sleepiest = id
		}
	}
	return sleepiest
}

func mostFreqAsleep(log []Row) (int, int) {
	var guards map[int]*[60]int = make(map[int]*[60]int)
	for _, row := range log {
		var gr *[60]int
		var ok bool
		if gr, ok = guards[row.id]; !ok {
			guards[row.id] = &[60]int{}
			gr = guards[row.id]
		}
		for i := 0; i < 60; i++ {
			gr[i] += row.min[i]
		}
		guards[row.id] = gr
	}
	var report []Row
	for id, gr := range guards {
		var r Row = Row{id: id}
		for i := 0; i < 60; i++ {
			r.min[i] = gr[i]
		}
		report = append(report, r)
	}
	// find the biggest number, remember row (id) and col (min)
	var max, maxid, maxmin int
	for _, row := range report {
		for i := 0; i < 60; i++ {
			if row.min[i] > max {
				max = row.min[i]
				maxid = row.id
				maxmin = i
			}
		}
	}
	fmt.Println(max, maxid, maxmin)
	for k, v := range report {
		fmt.Println(k, 32070.0/float32(v.id), v)
	}

	return maxid, maxmin
}

func mostAsleep(log []Row, id int) int {
	var r2 Row
	for _, v := range log {
		if v.id == id {
			for i := 0; i < 60; i++ {
				r2.min[i] += v.min[i]
			}
		}
	}
	// fmt.Println(r2)
	var max, result int
	for i := 0; i < 60; i++ {
		if r2.min[i] > max {
			max = r2.min[i]
			result = i
		}
	}
	return result
}

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

func run() (int, int) {
	var lines []string
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	sort.Slice(lines, func(a, b int) bool { return lines[a] < lines[b] })

	var sleeplog []Row = []Row{}
	var row *Row
	var sleep int

	re := regexp.MustCompile(`\[1518-(.+) [0-9][0-9]:(.+)\] (.+)`)
	for _, line := range lines {
		matches := re.FindAllStringSubmatch(line, -1)
		monday := matches[0][1]
		min := matches[0][2]
		rhs := matches[0][3]
		if strings.HasPrefix(rhs, "Guard #") {
			var guard int
			if n, err := fmt.Sscanf(rhs, "Guard #%d begins shift", &guard); n != 1 {
				fmt.Println(err)
				break
			}
			if row != nil {
				sleeplog = append(sleeplog, *row)
			}
			row = &Row{id: guard, date: monday}
		} else if rhs == "falls asleep" {
			sleep = atoi(min)
		} else if rhs == "wakes up" {
			wake := atoi(min)
			for i := sleep; i <= wake; i++ {
				row.min[i] = 1
			}
		} else {
			fmt.Println("unknown rhs", rhs)
		}
	}
	if row != nil {
		sleeplog = append(sleeplog, *row)
	}

	// for k, v := range sleeplog {
	// 	if v.id == 1933 {
	// 		fmt.Println(k, v)
	// 	}
	// }

	sleepiest := sleepiestGuard(sleeplog)
	most := mostAsleep(sleeplog, sleepiest)
	// fmt.Println(sleepiest, most)

	maxid, maxmin := mostFreqAsleep(sleeplog)

	return sleepiest * most, maxid * maxmin
}

func main() {
	defer duration(time.Now(), "main")

	fmt.Println(run()) // 2953 * 39 = 115167
	// 118120 2953 * 40 too high
	// 115167 2953 * 39 too high
	// 32070 (1069*30, 17 times)
}

/*
$ go run main.go
*/
