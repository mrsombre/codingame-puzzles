package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntToStr(t *testing.T) {
	tests := []struct {
		name     string
		x        int
		expected string
	}{
		{
			name:     `>0`,
			x:        1,
			expected: "1",
		},
		{
			name:     `<0`,
			x:        -1,
			expected: "-1",
		},
		{
			name:     `0`,
			x:        0,
			expected: "0",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, IntToStr(tc.x))
		})
	}
}

func TestStrToInt(t *testing.T) {
	tests := []struct {
		name     string
		str      string
		expected int
	}{
		{
			name:     "Positive",
			str:      "42",
			expected: 42,
		},
		{
			name:     "Negative",
			str:      "-10",
			expected: -10,
		},
		{
			name:     "Zero",
			str:      "0",
			expected: 0,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			num := StrToInt(tc.str)
			assert.Equal(t, tc.expected, num)
		})
	}
}

func TestBoolToInt(t *testing.T) {
	tests := []struct {
		name     string
		x        bool
		expected int
	}{
		{
			name:     `true`,
			x:        true,
			expected: 1,
		},
		{
			name:     `false`,
			x:        false,
			expected: 0,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, BoolToInt(tc.x))
		})
	}
}

func TestIntToBool(t *testing.T) {
	tests := []struct {
		name     string
		x        int
		expected bool
	}{
		{
			name:     `true`,
			x:        1,
			expected: true,
		},
		{
			name:     `false`,
			x:        0,
			expected: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, IntToBool(tc.x))
		})
	}
}
