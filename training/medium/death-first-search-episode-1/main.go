package main

import (
	"fmt"
)

type ID = int

type Node struct {
	ID     ID
	Edges  []*Edge
	IsGate bool
}

type Edge struct {
	NodeA    *Node
	NodeB    *Node
	IsSealed bool
}

// Next returns the next node in the edge based on the given node.
func (e *Edge) Next(node *Node) *Node {
	if e.NodeA == node {
		return e.NodeB
	}
	if e.NodeB == node {
		return e.NodeA
	}
	panic("node not found in the edge")
}

func (e *Edge) String() string {
	return fmt.Sprintf("%d %d", e.NodeA.ID, e.NodeB.ID)
}

type Graph struct {
	Nodes []*Node

	NodesCnt int
	EdgesCnt int
	GatesCnt int
}

func NewGraph(size int) *Graph {
	graph := &Graph{
		Nodes:    make([]*Node, 0, size),
		NodesCnt: size,
	}

	for i := 0; i < size; i++ {
		graph.Nodes = append(graph.Nodes, &Node{ID: i})
	}
	return graph
}

func (g *Graph) AddEdge(a, b ID) {
	a, b = minmax(a, b)

	edge := &Edge{
		NodeA: g.Nodes[a],
		NodeB: g.Nodes[b],
	}

	g.Nodes[a].Edges = append(g.Nodes[a].Edges, edge)
	g.Nodes[b].Edges = append(g.Nodes[b].Edges, edge)
	g.EdgesCnt++
}

func (g *Graph) SetGateway(n ID) {
	g.Nodes[n].IsGate = true
	g.GatesCnt++
}

// IsFree checks whether the agent is adjacent to a gate and if there is no then free move available.
func (g *Graph) IsFree(start ID) bool {
	agent := g.Nodes[start]

	for _, edge := range agent.Edges {
		if edge.IsSealed {
			continue
		}
		if edge.Next(agent).IsGate {
			return false
		}
	}
	return true
}

// CircleAmbush try to block circular paths for the Ambush Achievement.
func (g *Graph) CircleAmbush() (bool, string) {
	gates := make([]*Node, 0, g.GatesCnt)
	for _, node := range g.Nodes {
		if !node.IsGate {
			continue
		}
		edgesCnt := 0
		for _, edge := range node.Edges {
			if edge.IsSealed {
				continue
			}
			edgesCnt++
		}
		if edgesCnt < 3 {
			continue
		}
		gates = append(gates, node)
	}

	for _, gate := range gates {
		adjacent := make(map[ID]*Node, len(gate.Edges))

		var start *Node
		for _, edge := range gate.Edges {
			if edge.IsSealed {
				continue
			}

			next := edge.Next(gate)
			adjacent[next.ID] = next
			start = next
		}

		// DFS to find a cycle among adjacent nodes.
		visited := make(map[ID]bool)
		var path []*Node
		var dfs func(*Node, *Node) bool
		dfs = func(node *Node, from *Node) bool {
			if visited[node.ID] {
				// If we return to the start node, we found a cycle.
				return node == start
			}
			visited[node.ID] = true
			path = append(path, node)

			for _, edge := range node.Edges {
				// Ensure the next node is unsealed, adjacent, and not the node we came from.
				if edge.IsSealed || edge.Next(node) == from {
					continue
				}
				next := edge.Next(node)
				if _, exists := adjacent[next.ID]; exists {
					// Continue DFS with the next node.
					if dfs(next, node) {
						return true
					}
				}
			}

			// Backtrack if no cycle found.
			visited[node.ID] = false
			path = path[:len(path)-1]
			return false
		}

		if dfs(start, nil) {
			// We found a cycle, seal all edges in the path.
			for _, node := range path {
				if len(node.Edges) != 3 {
					continue
				}
				for _, edge := range node.Edges {
					if edge.IsSealed {
						continue
					}
					if edge.Next(node).IsGate {
						continue
					}
					edge.IsSealed = true
					return true, edge.String()
				}
			}
		}
	}

	return false, ""
}

// BFS finds the shortest path to a gateway.
func (g *Graph) BFS(agent ID) string {
	visited := make(map[ID]bool)
	queue := make([]*Node, 0, g.EdgesCnt)
	queue = append(queue, g.Nodes[agent])
	path := make(map[ID]*Edge)

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if visited[current.ID] {
			continue
		}
		visited[current.ID] = true

		for _, edge := range current.Edges {
			if edge.IsSealed {
				continue
			}

			next := edge.Next(current)
			if next.IsGate {
				// We found a gateway, backtrack to find the path.
				for current.ID != agent {
					edge = path[current.ID]
					current = edge.Next(current)
				}
				edge.IsSealed = true
				return edge.String()
			}

			if visited[next.ID] {
				continue
			}

			path[next.ID] = edge
			queue = append(queue, edge.Next(current))
		}
	}

	return ""
}

func minmax(a, b ID) (ID, ID) {
	if a < b {
		return a, b
	}
	return b, a
}

func main() {
	var err error

	// N: the total number of nodes in the level, including the gateways
	// L: the number of links
	// E: the number of exit gateways
	var N, L, E int
	_, err = fmt.Scan(&N, &L, &E)
	if err != nil {
		panic(err)
	}

	graph := NewGraph(N)

	for i := 0; i < L; i++ {
		// N1 and N2 defines a link between these nodes
		var N1, N2 int
		_, err = fmt.Scan(&N1, &N2)
		if err != nil {
			panic(err)
		}
		graph.AddEdge(N1, N2)
	}

	for i := 0; i < E; i++ {
		// EI: the index of a gateway node
		var EI int
		_, err = fmt.Scan(&EI)
		if err != nil {
			panic(err)
		}
		graph.SetGateway(EI)
	}

	for {
		// SI: The index of the node on which the Bobnet agent is positioned this turn
		var SI int
		_, err = fmt.Scan(&SI)
		if err != nil {
			panic(err)
		}

		if graph.IsFree(SI) {
			if ambush, edge := graph.CircleAmbush(); ambush {
				fmt.Println(edge)
				continue
			}
		}

		link := graph.BFS(SI)
		fmt.Println(link)
	}
}
