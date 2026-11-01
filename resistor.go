package main

// 2026 Kenneth Galle

import (
	"fmt"
	"sort"
)

// list of resistor values to combine together
var resistor_values = [...]float32{0.0, 1.0, 2.2, 3.3, 4.7, 5.1, 10.0,
	15.0, 22.0, 33.0, 47.0, 51.0, 62.0, 82.0, 100.0, 120.0, 150.0, 180.0,
	200.0, 220.0, 240.0, 270.0, 300.0, 330.0, 470.0, 510.0, 560.0, 680.0,
	820.0, 1000.0, 1200.0, 1500.0, 1800.0, 2000.0, 2200.0, 3000.0, 3300.0,
	4700.0, 5100.0, 5600.0, 7500.0, 10000.0, 12000.0, 15000.0, 18000.0,
	20000.0, 22000.0, 33000.0, 43000.0, 47000.0, 51000.0, 75000.0,
	100000.0, 200000.0, 300000.0, 390000.0, 470000.0, 680000.0, 1000000.0,
	10000000.0}

// var resistor_values = [...]float32{1.0, 10.0, 100.0}

// number_of_values
var resistor_count = len(resistor_values)

/*
tree like with rv nodes in the root
where each childnode can likewise have rv nodes as children
depth of the tree determines number of resistors combined together
*/

func main() {
	const DEPTH = 3
	nodes := new(Nodes)
	nodes.initial_fill(DEPTH)
	fmt.Printf("%d nodes arranged %d deep\n\n", len(*nodes), DEPTH)
	node_result := nodes.getResults()

	sort.Slice(node_result, func(i, j int) bool {
		return node_result[i].resistor_value < node_result[j].resistor_value
	})

	for i := 0; i < len(node_result); i++ {
		fmt.Println(node_result[i].asString())
	}

}
