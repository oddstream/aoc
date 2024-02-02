package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var data string

type program struct {
	mem []int
	pc  int
	rel int
	in  chan int
	out chan int
}

func newProgram(mem []int, in, out chan int) *program {
	return &program{
		mem: append([]int(nil), mem...),
		in:  in,
		out: out,
	}
}

func (p *program) run() {
	for {
		op := p.mem[p.pc] % 100
		mode1 := p.mem[p.pc] / 100 % 10
		mode2 := p.mem[p.pc] / 1000 % 10
		mode3 := p.mem[p.pc] / 10000 % 10

		switch op {
		case 1:
			p.write(3, p.read(1, mode1)+p.read(2, mode2), mode3)
			p.pc += 4
		case 2:
			p.write(3, p.read(1, mode1)*p.read(2, mode2), mode3)
			p.pc += 4
		case 3:
			p.write(1, <-p.in, mode1)
			p.pc += 2
		case 4:
			p.out <- p.read(1, mode1)
			p.pc += 2
		case 5:
			if p.read(1, mode1) != 0 {
				p.pc = p.read(2, mode2)
			} else {
				p.pc += 3
			}
		case 6:
			if p.read(1, mode1) == 0 {
				p.pc = p.read(2, mode2)
			} else {
				p.pc += 3
			}
		case 7:
			if p.read(1, mode1) < p.read(2, mode2) {
				p.write(3, 1, mode3)
			} else {
				p.write(3, 0, mode3)
			}
			p.pc += 4
		case 8:
			if p.read(1, mode1) == p.read(2, mode2) {
				p.write(3, 1, mode3)
			} else {
				p.write(3, 0, mode3)
			}
			p.pc += 4
		case 9:
			p.rel += p.read(1, mode1)
			p.pc += 2
		case 99:
			close(p.out)
			return
		default:
			panic(fmt.Sprintf("unknown opcode %d", op))
		}
	}
}

func (p *program) read(offset, mode int) int {
	switch mode {
	case 0:
		return p.mem[p.mem[p.pc+offset]]
	case 1:
		return p.mem[p.pc+offset]
	case 2:
		return p.mem[p.mem[p.pc+offset]+p.rel]
	default:
		panic(fmt.Sprintf("unknown mode %d", mode))
	}
}

func (p *program) write(offset, value, mode int) {
	switch mode {
	case 0:
		p.mem[p.mem[p.pc+offset]] = value
	case 1:
		panic("writing to immediate mode")
	case 2:
		p.mem[p.mem[p.pc+offset]+p.rel] = value
	default:
		panic(fmt.Sprintf("unknown mode %d", mode))
	}
}

func main() {
	var mem []int
	for _, s := range strings.Split(strings.TrimSpace(string(data)), ",") {
		n, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		mem = append(mem, n)
	}

	in := make(chan int)
	out := make(chan int)
	go newProgram(mem, in, out)

	go func() {
		for i := 0; i < 10; i++ {
			in <- i
		}
		close(in)
	}()

	for n := range out {
		fmt.Println(n)
	}
}
