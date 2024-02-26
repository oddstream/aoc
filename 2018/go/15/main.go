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

const debugChecks bool = true

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

type (
	Area [][]string
	// Unit describes an Elf or a Goblin
	Unit struct {
		Point
		species                string // "E"lf or "G"oblin
		attackPower, hitPoints int    // starts with 3, 200
	}
	// Units is a container for all alive units
	Units map[Point]*Unit
	Point struct {
		y, x int
	}
)

// TODO at what point does the output from this, and
// python3 15.y 3<input.txt
// diverge, and why?
// diverges between rounds 38 .. 39

// TODO store pointers to units in a map, not as pointers in Area squares
// eg units = map[Point]struct{species, attackPower, hitPoints}
// grid will become [][]string with "#" wall else not-a-wall
// this is making a change just to keep the fingers busy,
// to see if anything shakes out

func (a Area) ingrid(p Point) bool {
	return p.x >= 0 && p.y >= 0 && p.x < len(a[0]) && p.y < len(a)
}

func (a Area) wall(p Point) bool {
	return a[p.y][p.x] == "#"
}

func (a Area) unoccupied(p Point) bool {
	return a[p.y][p.x] == "."
}

func (p Point) add(q Point) Point {
	return Point{y: p.y + q.y, x: p.x + q.x}
}

// used by findEnemiesInRange and findAdjacentSquares
// changing the order of these points has NO effect on result
var toutesDirections []Point = []Point{
	{y: -1, x: 0}, // n/u
	{y: 0, x: -1}, // w/l
	{y: 0, x: 1},  // e/r
	{y: 1, x: 0},  // s/d
}

// used by bfs
// changing the order of these points has DOES effect the result
// which is weird and a clue to why result is wrong
var toutesDirectionsBfs []Point = []Point{
	{y: -1, x: 0}, // n/u
	{y: 0, x: -1}, // w/l
	{y: 0, x: 1},  // e/r
	{y: 1, x: 0},  // s/d
}

func duration(invocation time.Time, name string) {
	fmt.Println(name, "duration", time.Since(invocation))
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

// load the area and units from embedded input
func loadInput(s string, elfAttackPower int) (Area, map[Point]*Unit) {
	var area Area
	var units Units = make(Units)

	scanner := bufio.NewScanner(strings.NewReader(s))
	var y int
	for scanner.Scan() {
		var row []string
		for x, ch := range scanner.Text() {
			var p Point = Point{y: y, x: x}
			if ch == '#' {
				row = append(row, "#")
			} else if ch == '.' {
				row = append(row, ".")
			} else if ch == 'E' {
				row = append(row, ".")
				var u Unit = Unit{Point: p, species: "E", hitPoints: 200, attackPower: elfAttackPower}
				units[p] = &u
			} else if ch == 'G' {
				row = append(row, ".")
				var u Unit = Unit{Point: p, species: "G", hitPoints: 200, attackPower: 3}
				units[p] = &u
			}
		}
		area = append(area, row)
		y += 1
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)

	}
	return area, units
}

func cull(in Units) Units {
	var out Units = make(Units)
	for p, u := range in {
		if u.hitPoints > 0 {
			out[p] = u
		}
	}
	// if len(in) != len(out) {
	// 	fmt.Println(len(in)-len(out), "units culled")
	// }
	return out
}

func alive(units Units, p Point) bool {
	var u *Unit = units[p]
	if u == nil {
		return false
	}
	if u.hitPoints > 0 {
		return true
	}
	return false
}

// utility function for sort.Slice()
func readingOrder(a, b Point) bool {
	if a.y != b.y {
		return a.y < b.y
	} else {
		return a.x < b.x
	}
}

// "To attack, the unit first determines all of the targets that are in range of it
// by being immediately adjacent to it. If there are no such targets,
// the unit ends its turn. Otherwise, the adjacent target with the fewest hit points is selected;
// in a tie, the adjacent target with the fewest hit points which is first in reading order is selected."
func findEnemyInRange(units Units, p Point) *Unit {
	var inrange []*Unit
	var unit = units[p]
	for _, dir := range toutesDirections {
		var q Point = p.add(dir)
		if enemy, ok := units[q]; ok {
			if enemy.species == unit.species {
				continue // not an enemy
			}
			if enemy.hitPoints > 0 {
				inrange = append(inrange, enemy)
			}
		}
	}
	if len(inrange) == 0 {
		return nil
	}
	sort.Slice(inrange, func(a, b int) bool {
		if inrange[a].hitPoints != inrange[b].hitPoints {
			return inrange[a].hitPoints < inrange[b].hitPoints
		} else {
			return readingOrder(inrange[a].Point, inrange[b].Point)
		}
	})
	return inrange[0]
}

func attack(units Units, unit *Unit, enemy *Unit) {
	// attack target with lowest hit points (in reading order)
	if debugChecks && enemy.hitPoints <= 0 {
		panic("attack: dead enemy")
	}
	enemy.hitPoints -= unit.attackPower
	if enemy.hitPoints <= 0 {
		fmt.Println(enemy.species, "killed")
		// delete(units, unit.Point) // TODO maybe postpone this, cull at end of round
	}
}

// return (up to four) points surrounding a unit's point
func findAdjacentSquares(area Area, units Units, p Point) []Point {
	var targets = []Point{}
	for _, dir := range toutesDirections {
		var q Point = p.add(dir)
		if area.unoccupied(q) {
			u := units[q]
			if u == nil || u.hitPoints <= 0 {
				targets = append(targets, q)
			}
		}
	}
	return targets
}

// return a list of all the enemies (targets) of this unit
// in reading order
func findAllEnemies(units Units, unit *Unit) []*Unit {
	var enemies []*Unit
	for _, enemy := range units {
		if enemy.species == unit.species {
			continue // not an enemy
		}
		if enemy.hitPoints > 0 {
			enemies = append(enemies, enemy)
		}
	}
	sort.Slice(enemies, func(a, b int) bool {
		return readingOrder(enemies[a].Point, enemies[b].Point)
	})
	return enemies
}

func bfs(area Area, units Units, start, end Point) (int, Point) {
	type BfsPoint struct {
		Point
		steps  int
		parent *BfsPoint
	}
	unoccupied := func(p Point) bool {
		return !area.wall(p) && !alive(units, p)
	}

	var seen map[Point]struct{} = make(map[Point]struct{})
	seen[start] = struct{}{}
	var q []BfsPoint = []BfsPoint{{Point: start}}
	for len(q) > 0 {
		var p BfsPoint
		p, q = q[0], q[1:]
		for _, dir := range toutesDirectionsBfs {
			var np Point = p.Point.add(dir)
			if !area.ingrid(np) {
				continue
			}
			if !unoccupied(np) {
				continue
			}
			if _, ok := seen[np]; ok {
				continue
			}
			seen[np] = struct{}{}
			var bfsp BfsPoint = BfsPoint{Point: np, parent: &p, steps: p.steps + 1}
			if np == end {
				var firstStep Point
				for par := &bfsp; par.parent != nil; par = par.parent {
					firstStep = par.Point
					if debugChecks {
						if !unoccupied(firstStep) {
							panic("bfs: firstStep is occupied")
						}
					}
				}
				return p.steps + 1, firstStep
			}
			q = append(q, bfsp)
		}
	}
	return -1, Point{}
}

// select the nearest reachable target square from a list of target squares
// (which is are unoccupied squares adjacent to an enemy)
func selectTarget(area Area, units Units, unit *Unit, targets map[Point]struct{}) (bool, Point) {
	type Target struct {
		steps     int
		target    Point
		firstStep Point
	}
	var t []Target
	for end := range targets {
		if debugChecks {
			if !area.unoccupied(end) {
				panic("selectTarget: end is a wall")
			}
			if u := units[end]; u != nil && u.hitPoints > 0 {
				panic("selectTarget: end has an alive unit")
			}
		}
		if steps, firstStep := bfs(area, units, unit.Point, end); steps != -1 {
			if debugChecks {
				if steps == 0 {
					panic("selectTarget: steps is zero") // plenty of these if don't have steps + 1 from bfs()...
				}
				if firstStep == unit.Point {
					panic("selectTarget: first step is same as start") // ...but none of these
				}
				if area.wall(firstStep) {
					panic("selectTarget: first step is a wall") // ...but none of these
				}
			}
			t = append(t, Target{steps: steps, target: end, firstStep: firstStep})
		}
	}
	if len(t) == 0 {
		return false, Point{}
	}
	/*
		   https://www.reddit.com/r/adventofcode/comments/a6urok/2018_day_15_something_thats_not_very_clear_and/
		   https://www.reddit.com/r/adventofcode/comments/a6hldy/comment/ebvifry/
		   1. "To move, the unit first considers the squares that are in range
		       and determines which of those squares it could reach in the fewest steps"
		   2. "If multiple squares are in range and tied for being reachable in the fewest steps,
		       the square which is first in reading order is chosen. "
		   3. "If multiple steps would put the unit equally closer to its destination,
		       the unit chooses the step which is first in reading order."

			Moving: you don't just take the path that takes you the fastest to an enemy,
			broken by reading order.
			You first choose the square adjacent to an enemy that you want to go to
			(closest, break ties by reading order),
			and then choose the move that takes you the fastest to it
			(break ties by reading order, again)."
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
			// nb this is never reached for any test or input
			return readingOrder(t[a].firstStep, t[b].firstStep)
		}
	})
	if debugChecks && !area.unoccupied(t[0].firstStep) {
		panic("chosen step is a wall")
	}
	return true, t[0].firstStep
}

func move(units Units, unit *Unit, newPoint Point) {
	if debugChecks {
		if unit != units[unit.Point] {
			panic("move: unit has moved")
		}
		if u, ok := units[newPoint]; ok {
			if u.hitPoints > 0 {
				panic("move: new point already has an alive unit")
			}
		}
	}
	delete(units, unit.Point)
	unit.Point = newPoint
	units[unit.Point] = unit
}

func round(area Area, units Units) (gameOver bool) {
	// in each round, each unit that is still alive takes a turn,
	// resolving all of its actions before the next unit's turn begins.
	// On each unit's turn,
	// it tries to move into range of an enemy (if it isn't already)
	// and then attack (if it is in range).

	// "the order in which units take their turns within a round
	// is the reading order of their starting positions in that round"
	// make a list of all units in reading order of their starting positions
	var unitsToProcess []*Unit
	for _, unit := range units {
		if unit.hitPoints <= 0 {
			continue
			// panic("round: found a dead unit")
		}
		unitsToProcess = append(unitsToProcess, unit)
	}
	sort.Slice(unitsToProcess, func(a, b int) bool {
		return readingOrder(unitsToProcess[a].Point, unitsToProcess[b].Point)
	})

	/*
		The ending condition is a bit subtle.
		The combat ends the first time a unit starts its turn with no enemies alive
		(which may be in the middle of a round)
		Make sure you break when the first goblin/elf notices a lack of enemies,
		not when an entire round goes without anyone doing anything.
		If the very last goblin/elf to go in a round kills its last enemy,
		that round should be counted but otherwise
		(if, say, the second-to-last elf/goblin to move kills its last target),
		the round is considered incomplete and not counted.
	*/
	for _, unit := range unitsToProcess {
		if unit.hitPoints <= 0 {
			continue // this unit was killed earlier in this round
		}
		// "Each unit begins its turn by identifying all possible targets (enemy units).
		// If no targets remain, combat ends."
		var enemies []*Unit
		if enemies = findAllEnemies(units, unit); len(enemies) == 0 {
			fmt.Println("no more enemies")
			return true // gameOver
		}

		if enemy := findEnemyInRange(units, unit.Point); enemy != nil {
			// "If the unit is already in range of a target,
			// it does not move, but continues its turn with an attack."
			attack(units, unit, enemy)
		} else {
			// "the unit identifies all of the open squares (.)
			// that are in range of each target;
			// these are the squares which are adjacent (immediately up, down, left, or right)
			// to any target and which aren't already occupied by a wall or another unit."
			// enemies := findAllEnemies(area, unit)
			// if len(enemies) == 0 {
			// 	return true // gameOver
			// }
			var adjacentSquares map[Point]struct{} = make(map[Point]struct{})
			for _, enemy := range enemies {
				if enemy.hitPoints <= 0 {
					panic("round: dead enemy")
				}
				for _, pt := range findAdjacentSquares(area, units, enemy.Point) {
					adjacentSquares[pt] = struct{}{}
				}
			}
			if len(adjacentSquares) > 0 {
				if debugChecks {
					for p := range adjacentSquares {
						if !area.unoccupied(p) {
							panic("round: adjacent point is a wall")
						}
					}
				}
				if ok, nextPoint := selectTarget(area, units, unit, adjacentSquares); ok {
					if debugChecks {
						if !area.unoccupied(nextPoint) {
							panic("round: next point is a wall")
						}
						if u := units[nextPoint]; u != nil && u.hitPoints > 0 {
							panic("selectTarget: nextPoint has an alive unit")
						}
					}

					// take a single step
					move(units, unit, nextPoint)
					// now try to attack again
					if enemy := findEnemyInRange(units, unit.Point); enemy != nil {
						attack(units, unit, enemy)
					}
				}
			}
			// "If the unit cannot reach (find an open path to) any of the squares that are in range, it ends its turn."
		}
	}
	return false // not gameOver
}

func numSpecies(units Units, species string) int {
	var count int
	for _, unit := range units {
		if unit.species == species {
			count += 1
		}
	}
	return count
}

func sumHitPoints(units Units) int {
	var sum int
	for _, unit := range units {
		if unit.hitPoints > 0 {
			sum += unit.hitPoints
		}
	}
	return sum
}

func displayArea(title string, area Area, units Units) {
	fmt.Println(title)
	for y := 0; y < len(area); y++ {
		var displayUnits []*Unit
		for x := 0; x < len(area[0]); x++ {
			pt := Point{y: y, x: x}
			if unit, ok := units[pt]; ok {
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
	fmt.Println()
}

func partOne(in string, expected int) {
	defer duration(time.Now(), "part 1")
	area, units := loadInput(in, 3)

	var result int

	for i := 0; i < 110; i++ {
		if round(area, units) {
			fmt.Println("break after", i, "rounds with", sumHitPoints(units), "hit points")
			result = i * sumHitPoints(units)
			break
		}
		units = cull(units)
		// displayArea(fmt.Sprintf("=== round %d ===", i+1), area, units)
	}

	fmt.Println(numSpecies(units, "E"), "elves", numSpecies(units, "G"), "goblins,", sumHitPoints(units), "hit points")

	report(expected, result)
}

func partTwo(in string, expected int) {
	defer duration(time.Now(), "part 2")

	var result int

	for elfPower := 18; elfPower < 20; elfPower++ {
		area, units := loadInput(in, elfPower)
		var elves = numSpecies(units, "E")
		for i := 0; i < 90; i++ {
			if round(area, units) {
				result = i * sumHitPoints(units)
				fmt.Println("power", elfPower, "break after", i, "rounds with", sumHitPoints(units), "hit points", numSpecies(units, "E"), "elves", numSpecies(units, "G"), "goblins", "result", result)
				if numSpecies(units, "E") == elves {
					goto exit
				} else {
					break
				}
			}
		}
	}

exit:
	report(expected, result)
}

func main() {
	defer duration(time.Now(), "main")

	// partOne(test1, 27730)
	// partOne(test2, 36334)
	// partOne(test3, 39514)
	// partOne(test4, 27755)
	// partOne(test5, 28944)
	// partOne(test6, 18740)

	partOne(input, 237996)

	partTwo(input, 69700) // power 19 break after 50 rounds with 1394 hit points 10 elves 5 goblins
}

/*
$ go run main.go
*/
