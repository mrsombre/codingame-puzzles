package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMedian(t *testing.T) {
	values := []float64{
		0, 2, 4, 6, 8, 10, 1000,
	}

	assert.EqualValues(t, 6, Median(values))
}

func TestExpectedExpDiff(t *testing.T) {
	tests := []struct {
		name    string
		current float64
		exp     float64
		decay   float64
		want    float64
	}{
		{
			name:    "positive ok",
			current: 10,
			exp:     10,
			want:    1,
		},
		{
			name:    "positive far",
			current: 21,
			exp:     10,
			want:    0,
		},
		{
			name:    "positive middle",
			current: 15,
			exp:     10,
			decay:   2,
			want:    0.36,
		},
		{
			name:    "positive far",
			current: 18,
			exp:     10,
			decay:   2,
			want:    0.2,
		},
		{
			name:    "zero ok",
			current: 0,
			exp:     0,
			want:    1,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.InDelta(t, tc.want, ExpectedExpDiff(tc.exp, tc.current, tc.decay), 0.1)
		})
	}
}
