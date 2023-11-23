package main

import (
	"bufio"
	_ "embed"
	"log"
	"strings"
	"time"
)

//go:embed input.txt
var input string

type Grid struct {
	cells []byte
}

var (
	width, height int
	grid          *Grid
)

var dirs [8][2]int = [8][2]int{
	{0, -1},  // n
	{1, -1},  // ne
	{1, 0},   // e
	{1, 1},   //se
	{0, 1},   //s
	{-1, 1},  // sw
	{-1, 0},  // w
	{-1, -1}, // nw
}

func duration(invocation time.Time, name string) {
	log.Println(name, time.Since(invocation))
}

func newGrid() *Grid {
	g := &Grid{}
	g.cells = make([]byte, width*height)
	return g
}

func (g *Grid) get(x, y int) byte {
	if x < 0 || x >= width {
		return 0
	}
	if y < 0 || y >= width {
		return 0
	}
	i := x + (y * width)
	if i < 0 || i > len(g.cells) {
		log.Fatal("Grid.get() index out of bounds\n")
	}
	return g.cells[i]
}

func (g *Grid) set(x, y int, b byte) {
	if x < 0 || x >= width {
		return
	}
	if y < 0 || y >= width {
		return
	}
	i := x + (y * width)
	if i < 0 || i > len(g.cells) {
		log.Fatal("Grid.set() index out of bounds\n")
	}
	g.cells[i] = b
}

func (g *Grid) lightsOn() int {
	var n int
	for _, b := range g.cells {
		if b == '#' {
			n += 1
		}
	}
	return n
}

func (g *Grid) neighboursOn(x, y int) int {
	var n int
	for _, dir := range dirs {
		nx := x + dir[0]
		ny := y + dir[1]
		b := g.get(nx, ny)
		if b == '#' {
			n += 1
		}
	}
	return n
}

func (g *Grid) show() {
	var i int
	for y := 0; y < height; y++ {
		line := g.cells[i : i+width]
		log.Println(y, string(line))
		i += width
	}
}

func (g *Grid) cycle(part int) {
	out := newGrid()
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if part == 2 {
				if x == 0 && y == 0 {
					out.set(x, y, '#')
					continue // top left
				}
				if x == width-1 && y == 0 {
					out.set(x, y, '#')
					continue // top right
				}
				if x == width-1 && y == height-1 {
					out.set(x, y, '#')
					continue // bottom right
				}
				if x == 0 && y == height-1 {
					out.set(x, y, '#')
					continue // bottom left
				}
			}
			b := g.get(x, y)
			n := g.neighboursOn(x, y)
			if b == '#' {
				//  light which is on stays on when 2 or 3 neighbors are on, and turns off otherwise
				if n == 2 || n == 3 {
					out.set(x, y, '#')
				} else {
					out.set(x, y, '.')
				}
			} else {
				// A light which is off turns on if exactly 3 neighbors are on, and stays off otherwise.
				if n == 3 {
					out.set(x, y, '#')
				} else {
					out.set(x, y, '.')
				}
			}
		}
	}
	g.cells = out.cells
}

func parseInput() {
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		if width == 0 {
			width = len(scanner.Text())
		}
		height += 1
	}
	if err := scanner.Err(); err != nil {
		log.Printf("scanner error: %v\n", err)
	}

	log.Println("grid is", width, "by", height)
	grid = newGrid()

	var i int
	scanner = bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()
		for j := 0; j < width; j++ {
			b := line[j]
			grid.cells[i] = b
			i += 1
		}
	}
}

func main() {
	defer duration(time.Now(), "main")
	parseInput()
	log.Println("initial grid has", grid.lightsOn(), "lights on")
	// grid.show()
	/*
		for i := 0; i < 100; i++ {
			grid.cycle(1)
			// log.Println("after", i+1, "step")
			// grid.show()
		}

		log.Println("grid has", grid.lightsOn(), "lights on")
		// part 1 768
	*/
	for i := 0; i < 100; i++ {
		grid.cycle(2)
		// log.Println("after", i+1, "step")
		// grid.show()
	}

	// grid.show()
	log.Println("grid has", grid.lightsOn(), "lights on")

	// part 2 781
}
