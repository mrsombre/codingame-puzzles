package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testGraph = func() *Graph {
	graph := NewGraph(4)

	graph.AddEdge(0, 1)
	graph.AddEdge(0, 2)
	graph.AddEdge(1, 3)
	graph.AddEdge(2, 3)

	return graph
}

func TestNewGraph(t *testing.T) {
	graph := NewGraph(1)

	assert.Equal(t, 0, graph.Nodes[0].ID)
	assert.Len(t, graph.Nodes, 1)
}

func TestGraph_AddEdge(t *testing.T) {
	graph := NewGraph(3)

	graph.AddEdge(0, 1)
	graph.AddEdge(0, 2)

	assert.Len(t, graph.Nodes[0].Edges, 2)
	assert.Len(t, graph.Nodes[1].Edges, 1)
}

func TestGraph_AddGateway(t *testing.T) {
	graph := NewGraph(3)

	graph.AddEdge(0, 1)
	graph.AddEdge(0, 2)
	graph.AddGateway(2)

	assert.False(t, graph.Nodes[0].Edges[0].IsGateway)
	assert.True(t, graph.Nodes[0].Edges[1].IsGateway)
}

func TestBFS(t *testing.T) {
	createGraph := func() *Graph {
		graph := NewGraph(3)

		graph.AddEdge(0, 1)
		graph.AddEdge(1, 2)
		graph.AddGateway(2)

		return graph
	}

	tests := []struct {
		name     string
		graph    func() *Graph
		si       int
		expected string
	}{
		{
			name: `center`,
			graph: func() *Graph {
				graph := createGraph()
				return graph
			},
			si:       1,
			expected: "1 2",
		},
		{
			name: `sealed`,
			graph: func() *Graph {
				graph := createGraph()
				graph.AddGateway(1)
				graph.Nodes[1].Edges[0].IsSealed = true

				return graph
			},
			si:       1,
			expected: "1 2",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, BFS(tc.graph(), tc.si))
		})
	}
}
