package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var circularGraph = func() *Graph {
	graph := NewGraph(8)

	graph.AddEdge(0, 1)
	graph.AddEdge(0, 2)
	graph.AddEdge(0, 3)
	graph.AddEdge(0, 4)
	graph.AddEdge(0, 5)
	graph.AddEdge(0, 6)

	graph.AddEdge(1, 2)
	graph.AddEdge(2, 3)
	graph.AddEdge(3, 4)
	graph.AddEdge(4, 5)
	graph.AddEdge(5, 6)
	graph.AddEdge(6, 1)

	graph.AddEdge(1, 7)
	graph.AddEdge(6, 7)

	graph.SetGateway(0)

	return graph
}

func TestIsFree(t *testing.T) {
	tests := []struct {
		name     string
		graph    func() *Graph
		start    int
		expected bool
	}{
		{
			name: `not free`,
			graph: func() *Graph {
				graph := circularGraph()
				return graph
			},
			start:    1,
			expected: false,
		},
		{
			name: `free`,
			graph: func() *Graph {
				graph := circularGraph()
				return graph
			},
			start:    7,
			expected: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, tc.graph().IsFree(tc.start))
		})
	}
}

func TestCircleAmbush(t *testing.T) {
	tests := []struct {
		name     string
		graph    func() *Graph
		expected string
		isAmbush bool
	}{
		{
			name: `simple`,
			graph: func() *Graph {
				graph := circularGraph()
				return graph
			},
			expected: `4 5`,
			isAmbush: true,
		},
		{
			name: `not ambush`,
			graph: func() *Graph {
				graph := testGraph()
				return graph
			},
			expected: ``,
			isAmbush: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			isAmbush, edge := tc.graph().CircleAmbush()
			assert.Equal(t, tc.expected, edge)
			assert.Equal(t, tc.isAmbush, isAmbush)
		})
	}
}
