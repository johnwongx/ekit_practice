package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverse(t *testing.T) {
	testCase := []struct {
		name     string
		input    []int
		wantSlic []int
	}{
		{
			name:     "normal",
			input:    []int{1, 2, 3, 4},
			wantSlic: []int{4, 3, 2, 1},
		},
		{
			name:     "one element",
			input:    []int{1},
			wantSlic: []int{1},
		},
		{
			name:     "empty",
			input:    []int{},
			wantSlic: []int{},
		},
		{
			name:     "nil",
			input:    nil,
			wantSlic: []int{},
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			res := Reverse[int](tc.input)
			assert.Equal(t, res, tc.wantSlic)
		})
	}
}

func TestReverseSelf(t *testing.T) {
	testCase := []struct {
		name     string
		input    []int
		wantSlic []int
	}{
		{
			name:     "normal",
			input:    []int{1, 2, 3, 4},
			wantSlic: []int{4, 3, 2, 1},
		},
		{
			name:     "one element",
			input:    []int{1},
			wantSlic: []int{1},
		},
		{
			name:     "empty",
			input:    []int{},
			wantSlic: []int{},
		},
		{
			name:     "nil",
			input:    nil,
			wantSlic: nil,
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			ReverseSelf[int](tc.input)
			assert.Equal(t, tc.input, tc.wantSlic)
		})
	}
}
