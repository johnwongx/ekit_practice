package slice

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestToMap(t *testing.T) {
	//TODO
}

func TestDeduplicate(t *testing.T) {
	testCase := []struct {
		name     string
		input    []int
		wantSlic []int
	}{
		{
			name:     "exist duplicate",
			input:    []int{1, 2, 1, 2},
			wantSlic: []int{1, 2},
		},
		{
			name:     "no duplicate",
			input:    []int{1, 2},
			wantSlic: []int{1, 2},
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
			res := deduplicateFunc[int](tc.input, func(src, dst int) bool {
				return src == dst
			})
			assert.Equal(t, res, tc.wantSlic)
		})
	}
}
