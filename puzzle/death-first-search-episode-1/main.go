package main

import (
	"fmt"
)

type Node struct {
	ID int

	Edges []*Edge
}

type Edge struct {
	ID int

	Node1 *Node
	Node2 *Node

	IsGateway bool
	IsSealed  bool
}

type Graph struct {
	Nodes []*Node
	Edges []*Edge
}

func NewGraph(size int) *Graph {
	graph := &Graph{
		Nodes: make([]*Node, 0, size),
	}
	for i := 0; i < size; i++ {
		graph.Nodes = append(graph.Nodes, &Node{ID: i})
	}
	return graph
}

func (g *Graph) AddEdge(n1, n2 int) {
	edge := &Edge{
		ID:    len(g.Edges),
		Node1: g.Nodes[n1],
		Node2: g.Nodes[n2],
	}
	g.Edges = append(g.Edges, edge)

	g.Nodes[n1].Edges = append(g.Nodes[n1].Edges, edge)
	g.Nodes[n2].Edges = append(g.Nodes[n2].Edges, edge)
}

func (g *Graph) AddGateway(n int) {
	for _, edge := range g.Nodes[n].Edges {
		edge.IsGateway = true
	}
}

func BFS(graph *Graph, start int) string {
	visited := make(map[int]bool)

	queue := make([]*Edge, 0, len(graph.Nodes))
	for _, edge := range graph.Nodes[start].Edges {
		queue = append(queue, edge)
	}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if visited[current.ID] {
			continue
		}
		visited[current.ID] = true

		if current.IsSealed {
			continue
		}

		if current.IsGateway {
			current.IsSealed = true
			return fmt.Sprintf("%d %d", current.Node1.ID, current.Node2.ID)
		}

		for _, edge := range current.Node1.Edges {
			if visited[edge.ID] {
				continue
			}
			queue = append(queue, edge)
		}
		for _, edge := range current.Node2.Edges {
			if visited[edge.ID] {
				continue
			}
			queue = append(queue, edge)
		}
	}

	return ""
}

func main() {
	// N: the total number of nodes in the level, including the gateways
	// L: the number of links
	// E: the number of exit gateways
	var N, L, E int
	fmt.Scan(&N, &L, &E)
	graph := NewGraph(N)

	for i := 0; i < L; i++ {
		// N1: N1 and N2 defines a link between these nodes
		var N1, N2 int
		fmt.Scan(&N1, &N2)
		graph.AddEdge(N1, N2)
	}

	for i := 0; i < E; i++ {
		// EI: the index of a gateway node
		var EI int
		fmt.Scan(&EI)
		graph.AddGateway(EI)
	}

	for {
		// SI: The index of the node on which the Bobnet agent is positioned this turn
		var SI int
		fmt.Scan(&SI)

		link := BFS(graph, SI)
		fmt.Println(link)
	}
}
