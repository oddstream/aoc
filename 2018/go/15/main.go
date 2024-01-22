// https://adventofcode.com/2018/day/15 Beverage Bandits
package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"sort"
	"strings"
	"time"
)

//go:embed test1.txt
var test1 string

//go:embed test2.txt
var test2 string

//go:embed test3.txt
var test3 string

//go:embed test4.txt
var test4 string

//go:embed test5.txt
var test5 string

//go:embed test6.txt
var test6 string

//go:embed input.txt
var input string

type Square struct {
	Point
	wall bool
	unit *Unit
	// open space if wall == false && unit == nil
}

// Unit describes an Elf or a Goblin
type Unit struct {
	Point
	species                string // "E"lf or "G"oblin
	attackPower, hitPoints int    // starts with 3, 200
}

type Area [][]Square

func (a Area) wall(p Point) bool {
	return a[p.y][p.x].wall
}

func (a Area) unit(p Point) *Unit {
	return a[p.y][p.x].unit
}

func (a Area) setUnit(p Point, unit *Unit) {
	a[p.y][p.x].unit = unit
}

// square in area is unoccupied if it is "." and has no units on it
func (a Area) unoccupied(p Point) bool {
	return !a.wall(p) && a.unit(p) == nil
}

type Point struct {
	y, x int
}

func (p Point) add(q Point) Point {
	return Point{y: p.y + q.y, x: p.x + q.x}
}

// changing the order of these changes the input result (but not the test results)
// "adjacent (immediately up, down, left, or right)"
var toutesDirections []Point = []Point{
	{y: -1, x: 0}, // n/u
	{y: 0, x: 1},  // e/r
	{y: 0, x: -1}, // w/l
	{y: 1, x: 0},  // s/d
}

var toutesDirectionsBfs []Point = []Point{
	{y: -1, x: 0}, // n/u
	{y: 0, x: 1},  // e/r
	{y: 0, x: -1}, // w/l
	{y: 1, x: 0},  // s/d
}

// TODO try using different versions of toutesDirections
// for bfs(), findEnemiesInRange(), findTargets()
// to see which one is order-sensitive?
// ... just the bfs

func duration(invocation time.Time, name string) {
	fmt.Println(name, time.Since(invocation))
}

// load the area and units from embedded input
func loadInput(s string, elfAttackPower int) Area {
	var area Area
	scanner := bufio.NewScanner(strings.NewReader(s))
	var y int
	for scanner.Scan() {
		var row []Square
		for x, ch := range scanner.Text() {
			var pt Point = Point{y: y, x: x}
			if ch == '#' {
				row = append(row, Square{Point: pt, wall: true})
			} else if ch == '.' {
				row = append(row, Square{Point: pt})
			} else {
				var u Unit = Unit{Point: pt, species: string(ch), hitPoints: 200}
				if ch == 'E' {
					u.attackPower = elfAttackPower
				} else {
					u.attackPower = 3
				}
				row = append(row, Square{Point: pt, unit: &u})
			}
		}
		area = append(area, row)
		y += 1
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)

	}
	return area
}

// utility function for sort.Slice()
func readingOrder(a, b Point) bool {
	if a.y != b.y {
		return a.y < b.y
	} else {
		return a.x < b.x
	}
}

// return a list of enemy units adjacent to this unit
// sorted into attack order (hit points, reading order)
func findEnemiesInRange(area Area, unit *Unit) []*Unit {
	var inrange []*Unit
	for _, dir := range toutesDirections {
		var q Point = unit.Point.add(dir)
		if enemy := area.unit(q); enemy != nil {
			if enemy.hitPoints <= 0 {
				panic("findEnemiesInRange: dead enemy")
			}
			if enemy.species != unit.species {
				inrange = append(inrange, enemy)
			}
		}
		sort.Slice(inrange, func(a, b int) bool {
			if inrange[a].hitPoints != inrange[b].hitPoints {
				return inrange[a].hitPoints < inrange[b].hitPoints
			} else {
				return readingOrder(inrange[a].Point, inrange[b].Point)
			}
		})
	}
	return inrange
}

func attack(area Area, unit *Unit, enemy *Unit) {
	// attack target with lowest hit points (in reading order)
	if enemy.hitPoints <= 0 {
		panic("attack: dead enemy")
	}
	enemy.hitPoints -= unit.attackPower
	if enemy.hitPoints <= 0 {
		fmt.Println(enemy.species, "killed")
		area.setUnit(enemy.Point, nil)
	}
}

// return (up to four) points surrounding this unit's point
// sorted into reading order
func findAdjacentSquares(area Area, unit *Unit) []Point {
	var targets = []Point{}
	for _, dir := range toutesDirections {
		var q Point = unit.Point.add(dir)
		if area.unoccupied(q) {
			targets = append(targets, q)
		}
	}
	sort.Slice(targets, func(a, b int) bool {
		return readingOrder(targets[a], targets[b])
	})
	return targets
}

// return a list of all the enemies (targets) of this unit
// in reading order
func findAllEnemies(area Area, unit *Unit) []*Unit {
	var enemies []*Unit
	for y := 0; y < len(area); y++ {
		for x := 0; x < len(area[0]); x++ {
			if enemy := area.unit(Point{y: y, x: x}); enemy != nil {
				if enemy.hitPoints <= 0 {
					panic("findAllEnemies: dead enemy")
				}
				if enemy.species != unit.species {
					enemies = append(enemies, enemy)
				}
			}
		}
	}
	return enemies
}

func bfs(area Area, start, end Point) (int, Point) {
	type BfsPoint struct {
		Point
		steps  int
		parent *BfsPoint
	}

	var seen map[Point]struct{} = make(map[Point]struct{})
	seen[start] = struct{}{}
	var q []BfsPoint = []BfsPoint{{Point: start}}
	for len(q) > 0 {
		var p BfsPoint
		p, q = q[0], q[1:]
		for _, dir := range toutesDirectionsBfs {
			var np Point = p.Point.add(dir)
			if _, ok := seen[np]; !ok {
				seen[np] = struct{}{}
				if area.unoccupied(np) {
					var bfsp BfsPoint = BfsPoint{Point: np, parent: &p, steps: p.steps + 1}
					if np.y == end.y && np.x == end.x {
						var firstStep Point
						for par := &bfsp; par.parent != nil; par = par.parent {
							firstStep = par.Point
						}
						return p.steps, firstStep
					}
					q = append(q, bfsp)
					// sort.Slice(q, func(a, b int) bool {
					// 	return readingOrder(q[a].Point, q[b].Point)
					// })
				}
			}
		}
	}
	return -1, Point{0, 0}
}

// select the nearest reachable target square
func selectTarget(area Area, unit *Unit, targets []Point) (bool, Point) {
	type Target struct {
		steps     int
		target    Point
		firstStep Point
	}
	var t []Target
	for _, end := range targets {
		if steps, firstStep := bfs(area, unit.Point, end); steps != -1 {
			t = append(t, Target{steps, end, firstStep})
		}
	}
	if len(t) == 0 {
		return false, Point{0, 0}
	}
	/*
	   https://www.reddit.com/r/adventofcode/comments/a6urok/2018_day_15_something_thats_not_very_clear_and/
	   https://www.reddit.com/r/adventofcode/comments/a6hldy/comment/ebvifry/
	   1. "To move, the unit first considers the squares that are in range and determines which of those squares it could reach in the fewest steps"
	   2. "If multiple squares are in range and tied for being reachable in the fewest steps, the square which is first in reading order is chosen. "
	   3. "If multiple steps would put the unit equally closer to its destination, the unit chooses the step which is first in reading order."
	*/
	// sort first by distance (number of steps),
	// then secondly by target square,
	// then thirdly by first step toward the target square.
	sort.Slice(t, func(a, b int) bool {
		if t[a].steps != t[b].steps {
			return t[a].steps < t[b].steps
		} else if t[a].target != t[b].target {
			return readingOrder(t[a].target, t[b].target)
		} else {
			return readingOrder(t[a].firstStep, t[b].firstStep)
		}
	})
	return true, t[0].firstStep
}

func move(area Area, unit *Unit, nextPoint Point) {
	if !area.unoccupied(nextPoint) {
		panic("unit overlap")
	}
	// if !(unit.Point.y == nextPoint.y+1 || unit.Point.y == nextPoint.y-1 || unit.Point.x == nextPoint.x+1 || unit.Point.x == nextPoint.x-1) {
	// 	panic("move: a point too far")
	// }
	area.setUnit(unit.Point, nil)
	unit.Point = nextPoint
	area.setUnit(unit.Point, unit)
}

func round(area Area) (gameOver bool) {
	// in each round, each unit that is still alive takes a turn,
	// resolving all of its actions before the next unit's turn begins.
	// On each unit's turn,
	// it tries to move into range of an enemy (if it isn't already)
	// and then attack (if it is in range).

	// make a list of all units in reading order of their starting positions
	var unitsToProcess []*Unit
	for y := 0; y < len(area); y++ {
		for x := 0; x < len(area[0]); x++ {
			if unit := area.unit(Point{y: y, x: x}); unit != nil {
				unitsToProcess = append(unitsToProcess, unit)
			}
		}
	}

	for _, unit := range unitsToProcess {
		if area.unit(unit.Point) == nil {
			continue // this unit was killed earlier in this round
		}
		// "Each unit begins its turn by identifying all possible targets (enemy units).
		// If no targets remain, combat ends."
		if enemies := findAllEnemies(area, unit); len(enemies) == 0 {
			return true // gameOver
		}

		inrange := findEnemiesInRange(area, unit)
		if len(inrange) > 0 {
			// "If the unit is already in range of a target,
			// it does not move, but continues its turn with an attack."
			attack(area, unit, inrange[0])
		} else {
			// "the unit identifies all of the open squares (.)
			// that are in range of each target;
			// these are the squares which are adjacent (immediately up, down, left, or right)
			// to any target and which aren't already occupied by a wall or another unit."
			enemies := findAllEnemies(area, unit)
			if len(enemies) == 0 {
				return true // gameOver
			}
			var targets []Point
			for _, enemy := range enemies {
				if enemy.hitPoints <= 0 {
					panic("round: dead enemy")
				}
				targets = append(targets, findAdjacentSquares(area, enemy)...)
			}
			if len(targets) > 0 {
				if false {
					// dedupe targets
					var tmp map[Point]struct{} = make(map[Point]struct{})
					for _, target := range targets {
						tmp[target] = struct{}{}
					}
					targets = nil
					for p := range tmp {
						targets = append(targets, p)
					}
					if len(targets) != len(tmp) {
						fmt.Println("dedupe", len(targets), len(tmp))
					}
					// sort into reading order
					sort.Slice(targets, func(a, b int) bool {
						return readingOrder(targets[a], targets[b])
					})
				}
				if ok, nextPoint := selectTarget(area, unit, targets); ok {
					// take a single step
					move(area, unit, nextPoint)
					// now try to attack again
					inrange := findEnemiesInRange(area, unit)
					if len(inrange) > 0 {
						attack(area, unit, inrange[0])
					}
				}
			}
			// if the unit cannot move to a target (route is blocked),
			// the turn ends
		}
	}
	return false // not gameOver
}

func numSpecies(area Area, species string) int {
	var count int
	for y := 0; y < len(area); y++ {
		for x := 0; x < len(area[0]); x++ {
			if unit := area.unit(Point{y: y, x: x}); unit != nil {
				if unit.species == species {
					count += 1
				}
			}
		}
	}
	return count
}

func sumHitPoints(area Area) int {
	var sum int
	for y := 0; y < len(area); y++ {
		for x := 0; x < len(area[0]); x++ {
			if unit := area.unit(Point{y: y, x: x}); unit != nil {
				sum += unit.hitPoints
			}
		}
	}
	return sum
}

func displayArea(title string, area Area) {
	fmt.Println(title)
	for y := 0; y < len(area); y++ {
		var displayUnits []*Unit
		for x := 0; x < len(area[0]); x++ {
			pt := Point{y: y, x: x}
			if unit := area.unit(pt); unit != nil {
				fmt.Print(unit.species)
				displayUnits = append(displayUnits, unit)
			} else if area.wall(pt) {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		if len(displayUnits) > 0 {
			fmt.Print("\t")
			for _, u := range displayUnits {
				fmt.Printf("%s(%d) ", u.species, u.hitPoints)
			}
		}
		fmt.Println()
	}
}

func partOne(in string, expected int) int {
	defer duration(time.Now(), "part 1")
	area := loadInput(in, 3)

	// displayArea("Initially:", area)
	// round(area)
	// displayArea("After 1 round:", area)
	// round(area)
	// displayArea("After 2 rounds:", area)

	var result int

	for i := 0; i < 110; i++ {
		if round(area) {
			fmt.Println("break after", i, "rounds with", sumHitPoints(area), "hit points")
			result = i * sumHitPoints(area)
			break
		}
	}
	// displayArea("Finally:", area)

	fmt.Println(numSpecies(area, "E"), "elves", numSpecies(area, "G"), "goblins,", sumHitPoints(area), "hit points")

	// fmt.Println(bfs(area,
	// 	Point{y: 1, x: 1},
	// 	Point{y: 5, x: 1}))
	// fmt.Println(bfs(area,
	// 	Point{y: 1, x: 1},
	// 	Point{y: 5, x: 5}))
	// fmt.Println(bfs(area,
	// 	Point{y: 1, x: 1},
	// 	Point{y: 3, x: 3}))
	// fmt.Println(bfs(area,
	// 	Point{y: 1, x: 1},
	// 	Point{y: 3, x: 5}))

	if expected != -1 {
		if result != expected {
			fmt.Println("ERROR: expected", expected, "got", result)
		} else {
			fmt.Println("CORRECT:", result)
		}
	}
	return result
}

func partTwo(in string, expected int) int {
	defer duration(time.Now(), "part 2")

	var result int

	for elfPower := 29; elfPower < 30; elfPower++ {
		area := loadInput(in, elfPower)
		var elves = numSpecies(area, "E") // 10
		for i := 0; i < 110; i++ {
			if round(area) || numSpecies(area, "E") < elves {
				fmt.Println("power", elfPower, "break after", i, "rounds with", sumHitPoints(area), "hit points", numSpecies(area, "E"), "elves", numSpecies(area, "G"), "goblins")
				result = i * sumHitPoints(area)
				break
			}
		}
	}

	if expected != -1 {
		if result != expected {
			fmt.Println("ERROR: expected", expected, "got", result)
		} else {
			fmt.Println("CORRECT:", result)
		}
	}
	return result
}

func main() {
	defer duration(time.Now(), "main")

	fmt.Println("part 1", partOne(test1, 27730))
	fmt.Println("part 1", partOne(test2, 36334))
	fmt.Println("part 1", partOne(test3, 39514))
	fmt.Println("part 1", partOne(test4, 27755))
	fmt.Println("part 1", partOne(test5, 28944))
	fmt.Println("part 1", partOne(test6, 18740))

	fmt.Println("part 1", partOne(input, 237996))

	// 245619 too high (2481 * 99)
	// 243138 wrong (2481 * 98)
	// 244020 wrong (2490 * 98)
	// 244213 wrong (2371 * 103)
	// 237996

	// fmt.Println("part 2", partTwo(input, 69700)) // rounds 50, elves left 10, attack power 4 .. 19
}

/*
$ go run main.go
*/
