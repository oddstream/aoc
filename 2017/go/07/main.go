// https://adventofcode.com/2017/day/7 Recursive Circus
package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"strconv"
	"strings"
	"time"
)

//go:embed input.txt
var input string

type Node struct {
	name     string
	parent   *Node
	children []*Node
	load     int // weight of children
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

// partOne finds the name of the root node,
// by making a set of all the nodes listed as having children,
// then passing through the input again,
// removing all child nodes from that set.
// what remains in the set is the root node,
// because it's a node that has children, but isn't another node's child
func partOne() string {
	var bases map[string]struct{} = make(map[string]struct{})
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		tokens := strings.Split(scanner.Text(), " ")
		if len(tokens) > 2 {
			bases[tokens[0]] = struct{}{}
		}
	}
	scanner = bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		tokens := strings.Split(scanner.Text(), " ")
		if len(tokens) > 3 {
			for i := 3; i < len(tokens); i++ {
				tok := strings.TrimRight(tokens[i], ",")
				delete(bases, tok)
			}
		}
	}
	if len(bases) > 1 {
		fmt.Println(bases) // oops
	}
	var result string
	for base := range bases {
		result = base
	}
	return result
}

func partTwo() int {
	// weights - map from node name to weight
	var weights map[string]int = make(map[string]int)
	// nodes - map from node name to Node pointer
	var nodes map[string]*Node = make(map[string]*Node)
	// towers - list of []string, like lines of input;
	// only one entry == leaf node,
	// multiple entry == immediate children of first entry
	var towers [][]string

	var addchildren func(*Node)
	var weigh func(*Node) int
	var report func(*Node, string)
	var check func(*Node)

	addchildren = func(p *Node) {
		// find n.name in towers (first entry in list)
		for _, tower := range towers {
			if tower[0] == p.name {
				// create nodes for all children tower[1..n]
				for i := 1; i < len(tower); i++ {
					child := Node{name: tower[i], parent: p}
					nodes[tower[i]] = &child
					p.children = append(p.children, &child)
					// recurse to add grandchildren
					addchildren(&child)
				}
				break
			}
		}
	}

	weigh = func(p *Node) int {
		p.load += weights[p.name]
		for _, child := range p.children {
			p.load += weigh(child)
		}
		return p.load
	}

	report = func(p *Node, tabs string) {
		fmt.Printf("%s %d %s (%d)\n", tabs, p.load, p.name, weights[p.name])
		for _, child := range p.children {
			report(child, tabs+"\t")
		}
	}

	var unweighted []*Node

	check = func(p *Node) {
		var sameweight bool = true
		for i := 1; i < len(p.children); i++ {
			if p.children[0].load != p.children[i].load {
				sameweight = false
				break
			}
		}
		if !sameweight {
			unweighted = append(unweighted, p)
		}
		for _, c := range p.children {
			check(c)
		}
	}

	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		tokens := strings.Split(scanner.Text(), " ")
		weights[tokens[0]] = atoi(strings.Trim(tokens[1], "()"))
		var tower []string = []string{tokens[0]}
		if len(tokens) > 3 {
			for i := 3; i < len(tokens); i++ {
				tower = append(tower, strings.TrimRight(tokens[i], ","))
			}
		}
		towers = append(towers, tower)
	}

	var root Node = Node{name: partOne()}
	nodes[root.name] = &root
	addchildren(&root)
	weigh(&root)
	check(&root)
	// unweighted will now be like [cyrupz qjvtm boropxd]
	var n *Node = unweighted[len(unweighted)-1]
	// a child of this node is the wrong weight
	// find unique weight using XOR trick
	var unique int
	for _, c := range n.children {
		unique ^= c.load
	}
	// fmt.Println(unique) // eg 1131
	var un, nn *Node // find unique node, normal node
	for _, c := range n.children {
		if c.load == unique {
			un = c
		} else {
			nn = c
		}
	}
	// fmt.Println(un.name, un.load, nn.name, nn.load)
	// report(&root, "")

	/*
	    836031 cyrupz (55)
	   	 119424 whbqia (23682)
	   		 15957 uiges (14517)
	   		 15957 jtwfdu (3727)
	   		 15957 spqvn (12168)
	   		 15957 gworrlc (14712)
	   		 15957 mgene (12150)
	   		 15957 aycnip (42)
	   	 119424 sopjux (81)
	   		 17049 tddqw (11164)
	   		 17049 iumspgx (1713)
	   		 17049 ojsntix (6025)
	   		 17049 uxuwgmz (8379)
	   		 17049 ehwofnh (18)
	   		 17049 zrtrhph (6535)
	   		 17049 rujdblh (15833)
	   	 119424 cfcdlh (96088)
	   		 5834 dwtgak (1742)
	   		 5834 whuak (62)
	   		 5834 rqjpza (1043)
	   		 5834 qpdufu (3326)
	   	 119424 wuluv (27912)
	   		 15252 atgrdn (90)
	   		 15252 hdpqtg (10989)
	   		 15252 qidhm (4140)
	   		 15252 socuv (6510)
	   		 15252 boslv (10868)
	   		 15252 fcakejv (4140)
	   	 119424 uwmocg (72774)
	   		 7775 npwjod (7625)
	   		 7775 fmqxggg (6539)
	   		 7775 tkaax (7058)
	   		 7775 tbpyoxy (5672)
	   		 7775 mefsxl (5015)
	   		 7775 hhxofiu (80)
	   	 119424 xtcpynj (31920)
	   		 14584 tixcp (14140)
	   		 14584 wlkzwch (87)
	   		 14584 sclfvp (7741)
	   		 14584 nafrtm (2458)
	   		 14584 vozeer (10879)
	   		 14584 ocowqvd (7522)
	   	 119432 qjvtm (82986)
	   		 12146 myfhxk (38)
	   		 12154 boropxd (4285)
	   			 1123 slzaeep (1023)
	   			 1123 hiotqxu (877)
	   			 1123 qppggd (171)
	   			 1123 iahug (397)
	   			 1131 cwwwj (201)
	   			 1123 upfhsu (27)
	   			 1123 jjlodie (46)
	   		 12146 jixdvf (7598)
	   			 1516 dohxzvo (752)
	   			 1516 tsjzvs (39)
	   			 1516 kvmjx (850)
	*/

	// 1131 cwwwj (201) -- too heavy, needs to be 201-8=193 to be 1123

	return weights[un.name] - (un.load - nn.load)
}

func main() {
	defer duration(time.Now(), "main")

	fmt.Println(partOne()) // cyrupz
	fmt.Println(partTwo()) // 193
}

/*
$ go run main.go
cyrupz
193
main 6.333537ms
*/
