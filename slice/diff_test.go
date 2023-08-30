package slice

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestDiff(t *testing.T) {
	testCase := []struct {
		name     string
		inputSrc []int
		inputDst []int
		wantSlic []int
	}{
		{
			name:     "exist diff val",
			inputSrc: []int{1, 2, 2, 3, 4},
			inputDst: []int{3, 4, 5},
			wantSlic: []int{1, 2},
		},
		{
			name:     "not inter val",
			inputSrc: []int{1, 2, 2, 3, 4},
			inputDst: []int{5},
			wantSlic: []int{1, 2, 3, 4},
		},
		{
			name:     "dst less than src",
			inputSrc: []int{1, 2, 3, 3, 4},
			inputDst: []int{3, 4},
			wantSlic: []int{1, 2},
		},
		{
			name:     "src less than dst",
			inputSrc: []int{3, 3, 4},
			inputDst: []int{1, 2, 3, 4},
			wantSlic: []int{},
		},
		{
			name:     "empty dst",
			inputSrc: []int{1, 2, 3, 4},
			inputDst: []int{},
			wantSlic: []int{1, 2, 3, 4},
		},
		{
			name:     "empty both",
			inputSrc: []int{},
			inputDst: []int{},
			wantSlic: []int{},
		},
		{
			name:     "nil dst",
			inputSrc: []int{1, 2, 3, 4},
			inputDst: nil,
			wantSlic: []int{1, 2, 3, 4},
		},
		{
			name:     "both nil",
			inputSrc: nil,
			inputDst: nil,
			wantSlic: []int{},
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			res := DiffSet[int](tc.inputSrc, tc.inputDst)
			assert.Equal(t, res, tc.wantSlic)
		})
	}
}

func TestDiffFunc(t *testing.T) {
	testCase := []struct {
		name     string
		inputSrc []int
		inputDst []int
		wantSlic []int
	}{
		{
			name:     "exist diff val",
			inputSrc: []int{1, 2, 2, 3, 4},
			inputDst: []int{3, 4, 5},
			wantSlic: []int{1, 2},
		},
		{
			name:     "not inter val",
			inputSrc: []int{1, 2, 2, 3, 4},
			inputDst: []int{5},
			wantSlic: []int{1, 2, 3, 4},
		},
		{
			name:     "dst less than src",
			inputSrc: []int{1, 2, 3, 3, 4},
			inputDst: []int{3, 4},
			wantSlic: []int{1, 2},
		},
		{
			name:     "src less than dst",
			inputSrc: []int{3, 3, 4},
			inputDst: []int{1, 2, 3, 4},
			wantSlic: []int{},
		},
		{
			name:     "empty dst",
			inputSrc: []int{1, 2, 3, 4},
			inputDst: []int{},
			wantSlic: []int{1, 2, 3, 4},
		},
		{
			name:     "empty both",
			inputSrc: []int{},
			inputDst: []int{},
			wantSlic: []int{},
		},
		{
			name:     "nil dst",
			inputSrc: []int{1, 2, 3, 4},
			inputDst: nil,
			wantSlic: []int{1, 2, 3, 4},
		},
		{
			name:     "both nil",
			inputSrc: nil,
			inputDst: nil,
			wantSlic: []int{},
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			res := DiffSetFunc[int](tc.inputSrc, tc.inputDst, func(src, dst int) bool {
				return src == dst
			})
			assert.Equal(t, res, tc.wantSlic)
		})
	}
}
