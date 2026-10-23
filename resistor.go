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

type Node struct {
	total_value float32
	level       int
	item_list   []int
}

/*
node
  array of
    total value (uint32)
	items that make up that total (index into the resistor_values array, []int)
	{0, []} is null value/node
*/

func add_resistor(depth int, nodes []Node) []Node {
	// if depth is down to 0, stop
	if depth > 0 {

		// create new []node to store results
		//		new_nodes := make([]Node, 0, resistor_count)

		// if nodes is empty, do initial fill
		if len(nodes) == 0 {
			for i := 0; i < resistor_count; i++ {
				node := Node{resistor_values[i], depth, []int{}}
				node.item_list = append(node.item_list, i)
				nodes = append(nodes, node)
			}
		} else {
			// add new set of resistor values onto each existing leaf node

		}

		// add resistor values to each item in nodes

		// call another_round on each node in nodes with i - 1 when i > 1

	}
	return nodes
}

func main() {
	nodes := add_resistor(1, []Node{})
	fmt.Print(len(nodes))
}
