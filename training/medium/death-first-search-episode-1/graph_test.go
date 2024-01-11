package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testGraph = func() *Graph {
	graph := NewGraph(8)

	graph.AddEdge(0, 1)
	graph.AddEdge(0, 2)
	graph.AddEdge(1, 3)
	graph.AddEdge(2, 3)

	graph.AddEdge(3, 4)

	graph.AddEdge(4, 5)
	graph.AddEdge(4, 6)
	graph.AddEdge(5, 7)
	graph.AddEdge(6, 7)

	return graph
}

func TestNewGraph(t *testing.T) {
	graph := NewGraph(2)

	assert.Len(t, graph.Nodes, 2)
	assert.Equal(t, 0, graph.Nodes[0].ID)
	assert.Equal(t, 1, graph.Nodes[1].ID)
	assert.Equal(t, graph.NodesCnt, 2)
}

func TestGraph_AddEdge(t *testing.T) {
	graph := NewGraph(2)
	graph.AddEdge(0, 1)

	assert.Len(t, graph.Nodes[0].Edges, 1)
	assert.Len(t, graph.Nodes[1].Edges, 1)
	assert.Same(t, graph.Nodes[0].Edges[0], graph.Nodes[1].Edges[0])
	assert.Equal(t, graph.EdgesCnt, 1)
}

func TestGraph_AddGateway(t *testing.T) {
	graph := NewGraph(2)
	graph.SetGateway(1)

	assert.False(t, graph.Nodes[0].IsGate)
	assert.True(t, graph.Nodes[1].IsGate)
	assert.Equal(t, graph.GatesCnt, 1)
}

func TestBFS(t *testing.T) {
	tests := []struct {
		name     string
		graph    func() *Graph
		start    int
		expected string
	}{
		{
			name: `simple`,
			graph: func() *Graph {
				graph := testGraph()
				graph.SetGateway(4)
				return graph
			},
			start:    3,
			expected: "3 4",
		},
		{
			name: `sealed`,
			graph: func() *Graph {
				graph := testGraph()
				graph.Nodes[4].Edges[1].IsSealed = true
				graph.SetGateway(7)
				return graph
			},
			start:    4,
			expected: "4 6",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, tc.graph().BFS(tc.start))
		})
	}
}
