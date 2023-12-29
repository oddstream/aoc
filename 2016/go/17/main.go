// https://adventofcode.com/2016/day/17 Two Steps Forward
package main

import (
	"crypto/md5"
	_ "embed"
	"encoding/hex"
	"fmt"
	"time"
)

const (
	WIDTH  int = 4
	HEIGHT int = 4
)

type Point struct {
	x, y int
}

type Step struct {
	x, y int
	path string // eg DULR
}

var directions map[byte]Point = map[byte]Point{
	'U': {x: 0, y: -1},
	'D': {x: 0, y: 1},
	'L': {x: -1, y: 0},
	'R': {x: 1, y: 0},
}

var input string = "veumntbg"
var test1 string = "ihgpwlah"
var test2 string = "kglvqrro"
var test3 string = "ulqzkmiv"

func duration(invocation time.Time, name string) {
	fmt.Println(name, time.Since(invocation))
}

func ingrid(p Point) bool {
	return p.x >= 0 && p.y >= 0 && p.x < WIDTH && p.y < HEIGHT
}

func makekeys(in string) string {
	hash := md5.Sum([]byte(in))
	out := hex.EncodeToString(hash[:])
	return out[:4]
}

var dir2idx map[byte]int = map[byte]int{
	'U': 0,
	'D': 1,
	'L': 2,
	'R': 3,
}

func open(keys string, dir byte) bool {
	if k, ok := dir2idx[dir]; ok {
		b := []byte(keys)[k]
		// "Any b, c, d, e, or f means that the corresponding door is open;
		// any other character (any number or a) means that the corresponding door is closed and locked."
		if b >= 'b' && b <= 'f' {
			return true
		}
	} else {
		fmt.Println("unknown dir in open")
	}
	return false
}

func bfs1(passcode string) string {
	var q []Step = []Step{{x: 0, y: 0, path: ""}}
	for len(q) > 0 {
		var st Step
		st, q = q[0], q[1:]
		var keys string = makekeys(passcode + st.path)
		for dir, p := range directions {
			var np Point = Point{x: st.x + p.x, y: st.y + p.y}
			if ingrid(np) {
				if open(keys, dir) {
					var ns Step = Step{x: np.x, y: np.y, path: st.path + string(dir)}
					if ns.x == WIDTH-1 && ns.y == HEIGHT-1 {
						return ns.path
					}
					q = append(q, ns)
				}
			}
		}
	}
	return "not found"
}

func bfs2(passcode string) int {
	var longest int
	var q []Step = []Step{{x: 0, y: 0, path: ""}}
	for len(q) > 0 {
		var step Step
		step, q = q[0], q[1:]
		var keys string = makekeys(passcode + step.path)
		for dir, p := range directions {
			var np Point = Point{x: step.x + p.x, y: step.y + p.y}
			if ingrid(np) {
				if open(keys, dir) {
					var newstep Step = Step{x: np.x, y: np.y, path: step.path + string(dir)}
					if newstep.x == WIDTH-1 && newstep.y == HEIGHT-1 {
						if len(newstep.path) > longest {
							longest = len(newstep.path)
						}
					} else {
						q = append(q, newstep)
					}
				}
			}
		}
	}
	return longest
}

func main() {
	defer duration(time.Now(), "main")
	// if makekeys("hijkl") != "ced9" {
	// 	fmt.Println("makekeys not working")
	// 	return
	// }
	// fmt.Println(open(makekeys("hijkl"), 'U'))
	// fmt.Println(open(makekeys("hijkl"), 'D'))
	// fmt.Println(open(makekeys("hijkl"), 'L'))
	// fmt.Println(open(makekeys("hijkl"), 'R'))
	// fmt.Println(open(makekeys("hijkl"), 'Z'))
	// fmt.Println(test1, bfs1(test1))
	// fmt.Println(test2, bfs1(test2))
	// fmt.Println(test3, bfs1(test3))
	fmt.Println(bfs1(input)) // DDRRULRDRD
	// fmt.Println(bfs2("ihgpwlah")) // 370
	// fmt.Println(bfs2("kglvqrro")) // 492
	// fmt.Println(bfs2("ulqzkmiv")) // 830
	fmt.Println(bfs2(input)) // 536
}

/*
$ go run main.go
DDRRULRDRD
536
main 29.53414ms
*/
