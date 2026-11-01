package main

// 2026 Kenneth Galle

import (
	"fmt"
)

type Node struct {
	parent         *Node
	resistor_index int
	children       Nodes
}

type NodeResult struct {
	resistor_value       float32
	resistor_arrangement string
}

type Nodes []Node

func (nodeResult NodeResult) asString() string {
	return fmt.Sprintf("%f\t%s", nodeResult.resistor_value, nodeResult.resistor_arrangement)
}

func (nodes Nodes) getResults() []NodeResult {
	// walk the tree and for each node call node.getResults()
	node_result := []NodeResult{}
	for i := 0; i < len(nodes); i++ {
		node_result = append(node_result, (nodes[i].getResults())...)
		if len(nodes[i].children) > 0 {
			node_result = append(node_result, nodes[i].children.getResults()...)
		}
	}
	return node_result
}

// return both series and parallel values for node;return if this is result for a root (top/last) node
func (node *Node) getResults() []NodeResult {
	node_result := []NodeResult{}
	if node.parent == nil {
		node_result = append(node_result, NodeResult{
			resistor_values[node.resistor_index],
			fmt.Sprintf("%f", resistor_values[node.resistor_index]),
		})
	} else {
		// use the parent value to create two new results, one for series, one for parallel
		parent_result := node.parent.getResults()
		// for each parent_result value, create two new results, one for series, one for parallel
		for i := 0; i < len(parent_result); i++ {
			// series
			node_result = append(node_result, NodeResult{
				parent_result[i].resistor_value + resistor_values[node.resistor_index],
				fmt.Sprintf("(%s + %f)", parent_result[i].resistor_arrangement, resistor_values[node.resistor_index]),
			})

			// parallel
			node_result = append(node_result, NodeResult{
				1.0 / ((1.0 / parent_result[i].resistor_value) + (1.0 / resistor_values[node.resistor_index])),
				fmt.Sprintf("(%s | %f)", parent_result[i].resistor_arrangement, resistor_values[node.resistor_index]),
			})
		}
	}
	return node_result
}

func (nodes *Nodes) initial_fill(depth int) {
	//nodes := *np
	// add a new node for each value in resistor_values
	for i := 0; i < resistor_count; i++ {
		node := Node{nil, i, []Node{}}
		node.node_fill(depth)
		*nodes = append(*nodes, node)
	}
	//*np = nodes
}

func (node *Node) node_fill(depth int) {
	// if depth is down to 0, stop
	if depth > 1 {
		// add children to node
		for i := 0; i < resistor_count; i++ {
			child := Node{node, i, []Node{}}
			child.node_fill(depth - 1)
			node.children = append(node.children, child)
		}
	}
}
