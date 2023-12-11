package main

import (
	"fmt"
	"image"
	"math"
	"os"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")
	split := strings.Fields(string(input))

	run := func(expand int) (d int) {
		galax := map[image.Point]struct{}{}
		dy := 0
		for y, s := range split {
			if !strings.Contains(s, "#") {
				dy += expand - 1
			}

			dx := 0
			for x, r := range s {
				col := ""
				for _, s := range split {
					col += string(s[x])
				}
				if !strings.Contains(col, "#") {
					dx += expand - 1
				}

				if r == '#' {
					for g := range galax {
						d += int(math.Abs(float64(g.X-(x+dx))) + math.Abs(float64(g.Y-(y+dy))))
					}
					galax[image.Point{x + dx, y + dy}] = struct{}{}
				}
			}
		}
		return
	}

	fmt.Println(run(2))
	fmt.Println(run(1000000))
}
