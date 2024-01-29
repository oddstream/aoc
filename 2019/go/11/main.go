// https://adventofcode.com/2019/day/11
package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
	"time"
)

//go:embed input.txt
var input string

const (
	BLACK   int = 0
	WHITE   int = 1
	LEFT90  int = 0
	RIGHT90 int = 1
)

const (
	ADD       int = 1
	MUL       int = 2
	INPUT     int = 3 // eg 3,50 takes an input value (from where?) and store it at address 50
	OUTPUT    int = 4 // eg 4,50 would output the value (to where?) at address 50
	JIT       int = 5 // jump if true
	JIF       int = 6 // jump if false
	LT        int = 7 // less than
	EQ        int = 8 // equals
	RELBASE   int = 9 // relative base
	HALT      int = 99
	POSITION  int = 0
	IMMEDIATE int = 1
	RELATIVE  int = 2
)

type Point struct {
	x, y int
}

func (p Point) add(q Point) Point {
	return Point{y: p.y + q.y, x: p.x + q.x}
}

type Robot struct {
	Point
	dir           string
	awaitingPaint bool // paint, then turn
}

var direction map[string]Point = map[string]Point{
	"U": {x: 0, y: -1},
	"D": {x: 0, y: 1},
	"L": {x: -1, y: 0},
	"R": {x: 1, y: 0},
}

var turnLeft map[string]string = map[string]string{
	"U": "L",
	"D": "R",
	"L": "D",
	"R": "U",
}

var turnRight map[string]string = map[string]string{
	"U": "R",
	"D": "L",
	"L": "U",
	"R": "D",
}

// given an opcode,
// gives number of ints this opcode will consume,
// not including the opcode itself
var arity map[int]int = map[int]int{
	ADD:     3,
	MUL:     3,
	INPUT:   1,
	OUTPUT:  1,
	JIT:     2,
	JIF:     2,
	LT:      3,
	EQ:      3,
	RELBASE: 1,
	HALT:    0,
}

type Instruction struct {
	opcode int    // one of ADD, MUL, INPUT, OUTPUT, HALT
	params [3]int // 0..3, depending on opcode
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

func readInstruction(program []int, pc int, relbase int) Instruction {
	var ins Instruction
	var n = program[pc]
	// opcode is lowest two decimal digits
	ins.opcode = n % 100
	n /= 100

	if _, ok := arity[ins.opcode]; !ok {
		fmt.Println("ERROR: unknown opcode", ins.opcode, "at", pc)
	}

	for i := 0; i < arity[ins.opcode]; i++ {
		mode := n % 10
		n /= 10
		switch mode {
		case IMMEDIATE:
			// "a parameter is interpreted as a value
			// if the parameter is 50, its value is simply 50"
			ins.params[i] = pc + 1 + i
		case POSITION:
			// "the parameter to be interpreted as a position
			// if the parameter is 50, its value is the value stored at address 50 in memory"
			ins.params[i] = program[pc+1+i]
		case RELATIVE:
			// ins.params[i] = program[relbase+pc+1+i]	// gave 203 for part 1 (wrong, too low)
			ins.params[i] = relbase + program[pc+1+i] // gave 3742852857, just doesn't look right
		default:
			fmt.Println("ERROR: unknown mode", mode)
			return Instruction{opcode: HALT}
		}
	}

	return ins
}

// return output, pc
func run(program []int, hull map[Point]int) {

	// " The robot starts facing up"
	var robot Robot = Robot{dir: "U", awaitingPaint: true} // Point will be 0,0
	var pc, relbase int                                    // Intcode interpreter state
	for {
		ins := readInstruction(program, pc, relbase)
		switch ins.opcode {
		case ADD:
			program[ins.params[2]] = program[ins.params[0]] + program[ins.params[1]]
			pc += 4
		case MUL:
			program[ins.params[2]] = program[ins.params[0]] * program[ins.params[1]]
			pc += 4
		case INPUT:
			// "The program uses input instructions to access the robot's camera:
			// provide 0 if the robot is over a black panel or 1 if the robot is over a white panel."
			// fmt.Println("INPUT", hull[robot.Point])
			program[ins.params[0]] = hull[robot.Point] // default to 0 (BLACK)
			pc += 2
		case OUTPUT:
			// "Then, the program will output two values:"
			// fmt.Println("OUTPUT", program[ins.params[0]])
			pc += 2
			if robot.awaitingPaint {
				// "First, it will output a value indicating the color to paint the panel the robot is over:
				// 0 means to paint the panel black, and 1 means to paint the panel white."
				hull[robot.Point] = program[ins.params[0]]
			} else {
				// "Second, it will output a value indicating the direction the robot should turn:
				// 0 means it should turn left 90 degrees, and 1 means it should turn right 90 degrees."
				switch program[ins.params[0]] {
				case LEFT90:
					robot.dir = turnLeft[robot.dir]
				case RIGHT90:
					robot.dir = turnRight[robot.dir]
				default:
					fmt.Println("ERROR: unknown turn", program[ins.params[0]])
					return
				}
				// "After the robot turns, it should always move forward exactly one panel."
				robot.Point = robot.Point.add(direction[robot.dir])
			}
			robot.awaitingPaint = !robot.awaitingPaint
		case JIT:
			if program[ins.params[0]] != 0 {
				pc = program[ins.params[1]]
			} else {
				pc += 3
			}
		case JIF:
			if program[ins.params[0]] == 0 {
				pc = program[ins.params[1]]
			} else {
				pc += 3
			}
		case LT:
			if program[ins.params[0]] < program[ins.params[1]] {
				program[ins.params[2]] = 1
			} else {
				program[ins.params[2]] = 0
			}
			pc += 4
		case EQ:
			if program[ins.params[0]] == program[ins.params[1]] {
				program[ins.params[2]] = 1
			} else {
				program[ins.params[2]] = 0
			}
			pc += 4
		case RELBASE:
			relbase += program[ins.params[0]]
			pc += 2
		case HALT:
			fmt.Println("EXIT")
			return
		default:
			fmt.Println("run: unknown opcode", ins.opcode)
			return
		}
	}
}

func display(hull map[Point]int) {
	for y := -1; y < 7; y++ {
		for x := 0; x < 40; x++ {
			if hull[Point{y: y, x: x}] == BLACK {
				fmt.Print(" ")
			} else {
				fmt.Print("#")
			}
		}
		fmt.Println()
	}
}

func part1() {
	defer duration(time.Now(), "part 1")

	var tokens []string = strings.Split(strings.Trim(input, "\n"), ",")
	var masterProgram []int
	for _, tok := range tokens {
		masterProgram = append(masterProgram, atoi(tok))
	}

	var hull map[Point]int = make(map[Point]int)

	// The computer's available memory should be much larger than the initial program.
	// Memory beyond the initial program starts with the value 0
	// and can be read or written like any other memory.
	var program []int = make([]int, len(masterProgram)+1000)
	copy(program, masterProgram)
	run(program, hull)
	fmt.Println("part 1", len(hull)) // 1930
}

func part2() {
	defer duration(time.Now(), "part 2")

	var tokens []string = strings.Split(strings.Trim(input, "\n"), ",")
	var masterProgram []int
	for _, tok := range tokens {
		masterProgram = append(masterProgram, atoi(tok))
	}

	var hull map[Point]int = map[Point]int{
		{0, 0}: WHITE,
	}

	// The computer's available memory should be much larger than the initial program.
	// Memory beyond the initial program starts with the value 0
	// and can be read or written like any other memory.
	var program []int = make([]int, len(masterProgram)+1000)
	copy(program, masterProgram)
	run(program, hull)
	display(hull) // PFKHECZU
}

func main() {
	defer duration(time.Now(), "main")

	part1()
	part2()
}

/*
$ go run main.go
EXIT
part 1 1930
part 1 4.056905ms
EXIT

 ###  #### #  # #  # ####  ##  #### #  #
 #  # #    # #  #  # #    #  #    # #  #
 #  # ###  ##   #### ###  #      #  #  #
 ###  #    # #  #  # #    #     #   #  #
 #    #    # #  #  # #    #  # #    #  #
 #    #    #  # #  # ####  ##  ####  ##

part 2 653.049µs
main 4.718048ms
*/
