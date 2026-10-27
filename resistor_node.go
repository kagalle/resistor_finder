package main

type Node struct {
	total_value float32
	depth       int
	item_list   []int
	children    []Node
}

type Nodes []Node

/*
node
  array of
    total value (uint32)
	items that make up that total (index into the resistor_values array, []int)
	{0, []} is null value/node
*/

func (nodes *Nodes) initial_fill(depth int) {
	//nodes := *np
	// add a new node for each value in resistor_values
	for i := 0; i < resistor_count; i++ {
		node := Node{resistor_values[i], depth, []int{}, []Node{}}
		node.item_list = append(node.item_list, i)
		node.node_fill()
		*nodes = append(*nodes, node)
	}
	//*np = nodes
}

func (node *Node) node_fill() {
	// if depth is down to 0, stop
	if node.depth > 1 {
		// add children to node
		for i := 0; i < resistor_count; i++ {
			child := Node{node.total_value + resistor_values[i], node.depth - 1, []int{}, []Node{}}
			child.item_list = append(child.item_list, node.item_list...)
			child.item_list = append(child.item_list, i)
			child.node_fill()
			node.children = append(node.children, child)
		}
	}
}
