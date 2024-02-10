// https://adventofcode.com/
package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"math"
	"strings"
	"time"
)

// all tiles (test1 and input) are 10x10
// 12 arrangements of tile (original, flipx, flipy) each rotated (0, 90, 180, 270)
// TODO are any of them the same?
// 9 tiles in test1 input (3x3)
// 144 tiles in input (12x12)

// not sure where to start,
// so try brute force
// every tile in every position in every rotation
// test will abort quickly if edges don't line up

// only need to store edges

// https://sethgeoghegan.com/advent-of-code-2020

var opp map[string]string = map[string]string{
	"top":    "bottom",
	"bottom": "top",
	"left":   "right",
	"right":  "left",
}

type Tile struct {
	id      int
	pattern [][]string // only need this for part 2
	edges   map[string]string
}

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

func newTile(id int, pattern [][]string) Tile {
	var tile Tile = Tile{id: id, pattern: pattern, edges: make(map[string]string)}
	tile.edges["left"] = leftEdge(tile.pattern)
	tile.edges["right"] = rightEdge(tile.pattern)
	tile.edges["top"] = topEdge(tile.pattern)
	tile.edges["bottom"] = bottomEdge(tile.pattern)
	return tile
}

func same(a, b [][]string) bool {
	for y := 0; y < 10; y++ {
		for x := 0; x < 10; x++ {
			if a[y][x] != b[y][x] {
				return false
			}
		}
	}
	return true
}

func contains(lst []Tile, p [][]string) bool {
	for _, item := range lst {
		if same(item.pattern, p) {
			return true
		}
	}
	return false
}

func (t Tile) display() {
	for y := 0; y < len(t.pattern); y++ {
		for x := 0; x < len(t.pattern[y]); x++ {
			fmt.Print(t.pattern[y][x])
		}
		fmt.Println()
	}
	fmt.Println("top    ", t.edges["top"])
	fmt.Println("left   ", t.edges["left"])
	fmt.Println("right  ", t.edges["right"])
	fmt.Println("bottom ", t.edges["bottom"])
}

func part1(in string, expected int) (result int) {
	defer duration(time.Now(), "part 1")

	var tileset []Tile
	var ids []int

	scanner := bufio.NewScanner(strings.NewReader(in))
	for scanner.Scan() {
		// get tile id number
		var id int
		if n, err := fmt.Sscanf(scanner.Text(), "Tile %d:", &id); n != 1 {
			fmt.Println(err)
			break
		}
		ids = append(ids, id)

		// get tile pattern
		var pattern [][]string
		for scanner.Scan() {
			if scanner.Text() == "" {
				break
			}
			pattern = append(pattern, strings.Split(scanner.Text(), ""))
		}
		// create all permutations (rotations and flips) of pattern and save in a slice
		var tiles []Tile
		for _, k := range [][][]string{pattern, flipx(pattern), flipy(pattern)} {
			if !contains(tiles, k) {
				tiles = append(tiles, newTile(id, k))
			}
			for i := 0; i < 3; i++ { // 90, 180, 270
				k = rotate(k)
				if !contains(tiles, k) {
					tiles = append(tiles, newTile(id, k))
				}
			}
		}
		// duplicates happen with flipx, 0, 90, 180 degrees
		// len(tiles) will be 8 (would be 12 (3*4) without duplicate removal)

		// add list of assembled tiles to map
		tileset = append(tileset, tiles...)
	}

	// test1 will have 9*8 = 72 tiles
	// input will have 144*8 = 1152 tiles

	// we only need the four corners to satisfy part 1

	possibleNeighbours := func(tile Tile, dir string) []int {
		var out []int
		for _, t := range tileset {
			if tile.id == t.id {
				continue
			}
			if tile.edges[dir] == t.edges[opp[dir]] {
				out = append(out, t.id)
			}
		}
		return out
	}

	// test1 totals: 16, 12, 8
	// input totals: 16, 12, 8
	// corners are tiles with 8 possibles
	// edges are tiles with 12 possibles
	// internal are tiles with 16 possibles

	var corners, edges, internals []int

	result = 1 // we are calculating PRODUCT of four ints

	for _, id := range ids {
		var total int
		for _, dir := range []string{"top", "bottom", "left", "right"} {
			var possibles map[int]struct{} = make(map[int]struct{})
			for _, tile := range tileset {
				if tile.id == id {
					for _, t := range possibleNeighbours(tile, dir) {
						possibles[t] = struct{}{}
					}
				}
			}
			total += len(possibles)
		}

		switch total {
		case 8:
			corners = append(corners, id)
			result *= id
		case 12:
			edges = append(edges, id)
		case 16:
			internals = append(internals, id)
		default:
			fmt.Println("unexpected number of possible neighbours", total)
		}
	}

	// test1 has 4 corners, 4 edges, 1 internals
	// input has 4 corners, 40 edges, 100 internals

	// let's make a picture

	var SZ int = int(math.Sqrt(float64(len(corners) + len(edges) + len(internals))))
	var picture [][]*Tile = make([][]*Tile, SZ)
	for y := 0; y < SZ; y++ {
		picture[y] = make([]*Tile, SZ)
	}

	// start by placing the corners
	// there will be 4! (24) ways of placing the corners
	// also, we know which tile id goes in the corners,
	// but we don't know which orientation that tile should be in
	// also, because of rotations, we might assemble the picture
	// in any orientation (ie there is >1 correct solution)
	/*
		for p := make([]int, len(corners)); p[0] < len(p); nextPerm(p) {
			perms := getPerm(corners, p)
			picture[0][0] = perms[0]
			picture[0][SZ-1] = perms[1]
			picture[SZ-1][0] = perms[2]
			picture[SZ-1][SZ-1] = perms[3]

			// how many edge tiles can go at picture[0][1]?
			// now try placing edges and backtracking
		}
	*/
	report(expected, result)
	return result
}

func main() {
	part1(test1, 20899048083289)  // 1951 * 3079 * 2971 * 1171
	part1(input, 174206308298779) // 3821 * 3343 * 3677 * 3709
}
