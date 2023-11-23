package main

import (
	_ "embed"
	"fmt"
	"log"
	"math"
	"time"
)

/*
	The boss's actual stats are in your puzzle input:
	Hit Points: 103
	Damage: 9
	Armor: 2

	Weapons:    Cost  Damage  Armor
	Dagger        8     4       0
	Shortsword   10     5       0
	Warhammer    25     6       0
	Longsword    40     7       0
	Greataxe     74     8       0

	Armor:      Cost  Damage  Armor
	Leather      13     0       1
	Chainmail    31     0       2
	Splintmail   53     0       3
	Bandedmail   75     0       4
	Platemail   102     0       5

	Rings:      Cost  Damage  Armor
	Damage +1    25     1       0
	Damage +2    50     2       0
	Damage +3   100     3       0
	Defense +1   20     0       1
	Defense +2   40     0       2
	Defense +3   80     0       3
*/

type ShopItem struct {
	name                string
	cost, damage, armor int
}

// nb including "None" items made part 2 answer too high (282 instead of 201)

var weapons []ShopItem = []ShopItem{
	// {"None", 0, 0, 0},
	{"Dagger", 8, 4, 0},
	{"Shortsword", 10, 5, 0},
	{"Warhammer", 25, 6, 0},
	{"Longsword", 40, 7, 0},
	{"Greataxe", 74, 8, 0},
}

var armor []ShopItem = []ShopItem{
	// {"None", 0, 0, 0},
	{"Leather", 13, 0, 1},
	{"Chainmail", 31, 0, 2},
	{"Splintmail", 53, 0, 3},
	{"Bandedmail", 75, 0, 4},
	{"Platemail", 102, 0, 5},
}

var rings []ShopItem = []ShopItem{
	// {"None 1", 0, 0, 0},
	// {"None 2", 0, 0, 0},
	{"Damage +1", 25, 1, 0},
	{"Damage +2", 50, 2, 0},
	{"Damage +3", 100, 3, 0},
	{"Defense +1", 20, 0, 1},
	{"Defense +2", 40, 0, 2},
	{"Defense +3", 80, 0, 3},
}

type Player struct {
	name              string
	HP, damage, armor int
}

func duration(invocation time.Time, name string) {
	fmt.Println(name, time.Since(invocation))
}

func report(player, boss Player) {
	log.Println(player.name, player.HP, ":", boss.HP, boss.name)
}

func round(player, boss *Player) {
	// An attacker always does at least 1 damage
	dmg := player.damage - boss.armor
	if dmg <= 0 {
		dmg = 1
	}
	boss.HP -= dmg
	if boss.HP < 0 {
		return
	}
	dmg = boss.damage - player.armor
	if dmg <= 0 {
		dmg = 1
	}
	player.HP -= dmg
}

func player_wins(player, boss *Player) bool {
	for player.HP > 0 && boss.HP > 0 {
		round(player, boss)
	}
	return player.HP > boss.HP
}

func part1and2() {
	var mingold int = math.MaxInt32
	var maxgold int = 0
	var minweap, minarm, minring1, minring2 string
	var maxweap, maxarm, maxring1, maxring2 string
	for _, weap := range weapons {
		for _, arm := range armor {
			for r1, ring1 := range rings {
				for r2, ring2 := range rings {
					if r1 == r2 {
						continue
					}
					player := &Player{name: "Player", HP: 100, damage: weap.damage + ring1.damage + ring2.damage, armor: arm.armor + ring1.armor + ring2.armor}
					boss := &Player{name: "Boss", HP: 103, damage: 9, armor: 2}
					gold := weap.cost + arm.cost + ring1.cost + ring2.cost
					if player_wins(player, boss) {
						if gold < mingold {
							mingold = gold
							minweap = weap.name
							minarm = arm.name
							minring1 = ring1.name
							minring2 = ring2.name
						}
					} else {
						if gold > maxgold {
							maxgold = gold
							maxweap = weap.name
							maxarm = arm.name
							maxring1 = ring1.name
							maxring2 = ring2.name
						}
					}
				}
			}
		}
	}
	log.Println("mingold", mingold, minweap, minarm, minring1, minring2) // 121
	log.Println("maxgold", maxgold, maxweap, maxarm, maxring1, maxring2) // 201
}

func main() {
	defer duration(time.Now(), "main")

	part1and2()
}
