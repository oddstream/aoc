// https://adventofcode.com/2016/day/11
package main

import (
	"fmt"
	"time"
)

type Element int
type Device int

const (
	hydrogen Element = iota
	lithium
	thulium
	plutonium
	promethium
	strontium
	ruthenium
	elerium
	dilithium
)

const (
	generator Device = iota
	microchip
)

type Item struct {
	f int
	e Element
	d Device
}

type State struct {
	items []Item
	key   int
	floor int
	moves int
}

var testInput = State{
	items: []Item{
		{f: 0, e: hydrogen, d: microchip},
		{f: 0, e: lithium, d: microchip},
		{f: 1, e: hydrogen, d: generator},
		{f: 2, e: lithium, d: generator},
	},
	key:   0,
	floor: 0,
	moves: 0,
}

var part1Input = State{
	items: []Item{
		{f: 0, e: thulium, d: generator},
		{f: 0, e: thulium, d: microchip},
		{f: 0, e: plutonium, d: generator},
		{f: 0, e: strontium, d: generator},

		{f: 1, e: plutonium, d: microchip},
		{f: 1, e: strontium, d: microchip},

		{f: 2, e: promethium, d: generator},
		{f: 2, e: promethium, d: microchip},
		{f: 2, e: ruthenium, d: generator},
		{f: 2, e: ruthenium, d: microchip},
	},
	key:   0,
	floor: 0,
	moves: 0,
}

var part2Input = State{
	items: []Item{
		{f: 0, e: thulium, d: generator},
		{f: 0, e: thulium, d: microchip},
		{f: 0, e: plutonium, d: generator},
		{f: 0, e: strontium, d: generator},
		{f: 0, e: elerium, d: generator},
		{f: 0, e: elerium, d: microchip},
		{f: 0, e: dilithium, d: generator},
		{f: 0, e: dilithium, d: microchip},

		{f: 1, e: plutonium, d: microchip},
		{f: 1, e: strontium, d: microchip},

		{f: 2, e: promethium, d: generator},
		{f: 2, e: promethium, d: microchip},
		{f: 2, e: ruthenium, d: generator},
		{f: 2, e: ruthenium, d: microchip},
	},
	key:   0,
	floor: 0,
	moves: 0,
}

func duration(invocation time.Time, name string) {
	fmt.Println(name, time.Since(invocation))
}

func combinations(start, stop int) <-chan [2]int {
	c := make(chan [2]int)
	go func() {
		defer close(c)
		for i := start; i < stop; i++ {
			for j := i + 1; j <= stop; j++ {
				c <- [2]int{i, j}
			}
		}
	}()
	return c
}

func (s State) genkey() int {
	var k int = s.floor + 1
	for _, item := range s.items {
		k = k*10 + (item.f + 1)
	}
	return k
}

func (s State) complete() bool {
	for _, i := range s.items {
		if i.f != 3 {
			return false
		}
	}
	return true
}

func (s State) clone() State {
	var ns = State{key: s.key, floor: s.floor, moves: s.moves}
	ns.items = append(ns.items, s.items...)
	return ns
}

func permutations(state State) <-chan State {
	c := make(chan State)
	go func() {
		defer close(c)
		var atfloor []int
		for i, item := range state.items {
			if item.f == state.floor {
				atfloor = append(atfloor, i)
			}
		}
		if len(atfloor) == 0 {
			panic("atfloor empty")
		}
		for _, newf := range []int{state.floor + 1, state.floor - 1} {
			if newf < 0 || newf > 3 {
				continue
			}
			for combi := range combinations(-1, len(atfloor)-1) {
				var n0, n1 int
				if combi[0] != -1 {
					n0 = atfloor[combi[0]]
				}
				n1 = atfloor[combi[1]]
				var newstate State = state.clone()
				newstate.floor = newf
				if combi[0] != -1 {
					newstate.items[n0].f = newf
				}
				newstate.items[n1].f = newf
				newstate.key = newstate.genkey()
				c <- newstate
			}
		}
	}()
	return c
}

func (s State) valid() bool {
	for _, a := range s.items {
		if a.f != s.floor && a.d == microchip {
			var found_another bool = false
			for _, b := range s.items {
				if b.f == a.f && b.e != a.e && b.d == generator {
					found_another = true
					break
				}
			}
			if found_another {
				var found_own bool = false
				for _, b := range s.items {
					if b.f == a.f && b.e == a.e && b.d == generator {
						found_own = true
						break
					}
				}
				if !found_own {
					return false
				}
			}
		}
	}
	return true
}

func bfs(start State) int {
	var seen map[int]struct{} = make(map[int]struct{})
	seen[start.key] = struct{}{}
	var q []State = []State{start}
	for len(q) > 0 {
		var state State
		state, q = q[0], q[1:]
		for newstate := range permutations(state) {
			if _, ok := seen[newstate.key]; !ok {
				seen[newstate.key] = struct{}{}
				if newstate.valid() {
					newstate.moves = state.moves + 1
					if newstate.complete() {
						return newstate.moves
					}
					q = append(q, newstate)
				}
			}
		}
	}
	return -1
}

func test(s State) int {
	for pair := range combinations(-1, 1) {
		fmt.Println(pair)
	}
	for newstate := range permutations(s) {
		fmt.Println(newstate.key, newstate.valid())
		for newstate2 := range permutations(newstate) {
			fmt.Println("\t", newstate2.key, newstate2.valid())
			for newstate3 := range permutations(newstate2) {
				fmt.Println("\t\t", newstate3.key, newstate3.valid())
			}
		}
	}
	return -1
}

func main() {
	defer duration(time.Now(), "main")
	part1Input.key = part1Input.genkey()
	// fmt.Println("complete", part1Input.complete())
	// fmt.Println("valid", part1Input.valid())
	fmt.Println(bfs(part1Input))
}

/*
31
main 7.086546212s

55
main 13m35.069347204s
*/
