package geometry

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMovingDistance(t *testing.T) {
	tests := []struct {
		name                      string
		speed, acceleration, time float64
		expected                  float64
	}{
		{
			name:         `speed`,
			speed:        10,
			acceleration: 0,
			time:         5,
			expected:     50,
		},
		{
			name:         `acceleration`,
			speed:        0,
			acceleration: 10,
			time:         5,
			expected:     125,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, MovingDistance(tc.speed, tc.acceleration, tc.time))
		})
	}
}

func TestMovingTime(t *testing.T) {
	tests := []struct {
		name                          string
		distance, speed, acceleration float64
		expected                      float64
	}{
		{
			name:         `speed`,
			distance:     50,
			speed:        10,
			acceleration: 0,
			expected:     5,
		},
		{
			name:         `acceleration`,
			distance:     125,
			speed:        0,
			acceleration: 10,
			expected:     5,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, MovingTime(tc.distance, tc.speed, tc.acceleration))
		})
	}
}
