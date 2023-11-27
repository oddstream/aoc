package main

import (
	"log"
	"math"
)

type Difficulty int

const (
	normal Difficulty = iota
	hard
)

type State struct {
	mana, spentMana                         int
	hp, bossHP                              int
	shieldTurns, poisonTurns, rechargeTurns int
}

type MagicItem struct {
	cost   int
	effect func(*State)
	active func(State) bool
}

var spells map[string]MagicItem = map[string]MagicItem{
	"magicMissile": {
		cost:   53,
		effect: func(s *State) { s.bossHP -= 4 },
		active: func(s State) bool { return false },
	},
	"drain": {
		cost:   73,
		effect: func(s *State) { s.bossHP -= 2; s.hp += 2 },
		active: func(s State) bool { return false },
	},
	"shield": {
		cost:   113,
		effect: func(s *State) { s.shieldTurns = 6 },
		active: func(s State) bool { return s.shieldTurns > 0 },
	},
	"poison": {
		cost:   173,
		effect: func(s *State) { s.poisonTurns = 6 },
		active: func(s State) bool { return s.poisonTurns > 0 },
	},
	"recharge": {
		cost:   229,
		effect: func(s *State) { s.rechargeTurns = 5 },
		active: func(s State) bool { return s.rechargeTurns > 0 },
	},
}

const bossDmg int = 9

func (s *State) applyEffects() {
	if s.poisonTurns > 0 {
		s.bossHP -= 3
		s.poisonTurns -= 1
	}
	if s.rechargeTurns > 0 {
		s.mana += 101
		s.rechargeTurns -= 1
	}
}

func battle(s State, difficulty Difficulty) int {
	var states []State = []State{s}
	var state State
	var result = math.MaxInt64

	for len(states) > 0 {
		// Go's ugly pop
		state = states[len(states)-1]
		states = states[:len(states)-1]

		if difficulty == hard {
			state.hp -= 1
			if state.hp <= 0 {
				continue
			}
		}

		state.applyEffects()
		if state.bossHP <= 0 {
			// min is now a built-in generic
			result = min(result, state.spentMana)
			continue
		}

		if state.shieldTurns > 0 {
			state.shieldTurns -= 1
		}

		for _, spell := range spells {
			if spell.cost <= state.mana &&
				state.spentMana+spell.cost < result &&
				!spell.active(state) {
				// Go primitive types are copied by value
				var newState State = state
				newState.mana -= spell.cost
				newState.spentMana += spell.cost
				spell.effect(&newState)

				newState.applyEffects()
				if newState.bossHP <= 0 {
					result = newState.spentMana
					continue
				}

				if newState.shieldTurns > 0 {
					newState.hp -= bossDmg - 7
					newState.shieldTurns -= 1
				} else {
					newState.hp -= bossDmg
				}

				if newState.hp > 0 {
					states = append(states, newState)
				}
			}
		}
	}
	return result
}

func main() {
	var startingPosition State = State{mana: 500, hp: 50, bossHP: 500}

	log.Println(battle(startingPosition, normal))
	log.Println(battle(startingPosition, hard))
}
