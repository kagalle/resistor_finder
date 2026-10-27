package main

import "fmt"

// number of resisters to combine in series
// const resistor_depth = 2

// list of resistor values to combine together
var resistor_values = [...]float32{0.0, 1.0, 2.2, 3.3, 4.7, 5.1, 10.0,
	15.0, 22.0, 33.0, 47.0, 51.0, 62.0, 82.0, 100.0, 120.0, 150.0, 180.0,
	200.0, 220.0, 240.0, 270.0, 300.0, 330.0, 470.0, 510.0, 560.0, 680.0,
	820.0, 1000.0, 1200.0, 1500.0, 1800.0, 2000.0, 2200.0, 3000.0, 3300.0,
	4700.0, 5100.0, 5600.0, 7500.0, 10000.0, 12000.0, 15000.0, 18000.0,
	20000.0, 22000.0, 33000.0, 43000.0, 47000.0, 51000.0, 75000.0,
	100000.0, 200000.0, 300000.0, 390000.0, 470000.0, 680000.0, 1000000.0,
	10000000.0}

// number_of_values
var resistor_count = len(resistor_values)

// current_value[resistor_count]

/*
tree like with rv nodes in the root
where each childnode can likewise have rv nodes as children
depth of the tree determines number of values in series
*/

func main() {
	nodes := new(Nodes)
	nodes.initial_fill(2)
	fmt.Print(len(*nodes))
	//	nodes := add_resistor(1, nil)

}
