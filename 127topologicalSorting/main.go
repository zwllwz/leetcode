package main

import (
	"fmt"
)

var i = 0

type DirectedGraphNode struct {
	Id        int
	Neighbors []DirectedGraphNode
}

func topsort(graph []DirectedGraphNode) []DirectedGraphNode {
	m := make(map[int]int)

	for _, node := range graph {
		for _, neighbor := range node.Neighbors {
			if v, ok := m[neighbor.Id]; ok {
				m[neighbor.Id] = v + 1
			} else {
				m[neighbor.Id] = 1
			}
		}
	}
	fmt.Println(m)

	result := []DirectedGraphNode{}
	queue := []DirectedGraphNode{}

	for _, node := range graph {
		if _, ok := m[node.Id]; !ok {
			queue = append(queue, node)
			result = append(result, node)
		}
	}
	//fmt.Println(queue)
	//fmt.Println(result)

	for len(queue) > 0 {
		fmt.Println("queue", queue)

		node := queue[0]
		queue = queue[1:]

		for _, n := range node.Neighbors {
			m[n.Id] = m[n.Id] - 1
			fmt.Println("map", m)
			if m[n.Id] == 0 {
				result = append(result, n)
				queue = append(queue, n)
			}
		}
	}

	return result
}

func main() {
	node0 := DirectedGraphNode{0, []DirectedGraphNode{}}
	node1 := DirectedGraphNode{1, []DirectedGraphNode{}}
	node2 := DirectedGraphNode{2, []DirectedGraphNode{}}
	node3 := DirectedGraphNode{3, []DirectedGraphNode{}}
	node4 := DirectedGraphNode{4, []DirectedGraphNode{}}
	node5 := DirectedGraphNode{5, []DirectedGraphNode{}}

	node0.Neighbors = []DirectedGraphNode{node1, node2, node3}
	node1.Neighbors = []DirectedGraphNode{node4}
	node2.Neighbors = []DirectedGraphNode{node4, node5}
	node3.Neighbors = []DirectedGraphNode{node4, node5}

	result := []DirectedGraphNode{node0, node1, node2, node3, node4, node5}

	// wrong answer = =
	fmt.Println(topsort(result))
}
