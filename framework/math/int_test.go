package math

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAbs(t *testing.T) {
	tests := []struct {
		name     string
		x        int
		expected int
	}{
		{
			name:     `>0`,
			x:        1,
			expected: 1,
		},
		{
			name:     `<0`,
			x:        -1,
			expected: 1,
		},
		{
			name:     `0`,
			x:        0,
			expected: 0,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, Abs(tc.x))
		})
	}
}

func TestMin(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{
			name:     `1 2`,
			a:        1,
			b:        2,
			expected: 1,
		},
		{
			name:     `2 1`,
			a:        2,
			b:        1,
			expected: 1,
		},
		{
			name:     `-1 1`,
			a:        -1,
			b:        1,
			expected: -1,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, Min(tc.a, tc.b))
		})
	}
}

func TestMax(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{
			name:     `1 2`,
			a:        1,
			b:        2,
			expected: 2,
		},
		{
			name:     `2 1`,
			a:        2,
			b:        1,
			expected: 2,
		},
		{
			name:     `-1 1`,
			a:        -1,
			b:        1,
			expected: 1,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, Max(tc.a, tc.b))
		})
	}
}
