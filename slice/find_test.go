package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFind(t *testing.T) {
	testCase := []struct {
		name      string
		inputSlic []int
		equalFunc matchFunc[int]
		wantVal   int
		wantRet   bool
	}{
		{
			name:      "find",
			inputSlic: []int{1, 2, 3, 4},
			equalFunc: func(src int) bool { return src == 2 },
			wantVal:   2,
			wantRet:   true,
		},
		{
			name:      "not find",
			inputSlic: []int{1, 2, 3, 4},
			equalFunc: func(src int) bool { return src == 233 },
			wantRet:   false,
		},
		{
			name:      "empty",
			inputSlic: []int{},
			equalFunc: func(src int) bool { return src == 2 },
			wantRet:   false,
		},
		{
			name:      "nil",
			inputSlic: nil,
			equalFunc: func(src int) bool { return src == 2 },
			wantRet:   false,
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			val, res := Find[int](tc.inputSlic, tc.equalFunc)
			assert.Equal(t, res, tc.wantRet)
			if !tc.wantRet {
				return
			}
			assert.Equal(t, val, tc.wantVal)
		})
	}
}

func TestFindAll(t *testing.T) {
	testCase := []struct {
		name      string
		inputSlic []int
		equalFunc matchFunc[int]
		wantSlic  []int
	}{
		{
			name:      "find one",
			inputSlic: []int{1, 2, 3, 4},
			equalFunc: func(src int) bool { return src == 2 },
			wantSlic:  []int{2},
		},
		{
			name:      "find multi",
			inputSlic: []int{1, 2, 3, 4},
			equalFunc: func(src int) bool { return src <= 2 },
			wantSlic:  []int{1, 2},
		},
		{
			name:      "not find",
			inputSlic: []int{1, 2, 3, 4},
			equalFunc: func(src int) bool { return src > 20 },
			wantSlic:  []int{},
		},
		{
			name:      "empty",
			inputSlic: []int{},
			equalFunc: func(src int) bool { return src <= 2 },
			wantSlic:  []int{},
		},
		{
			name:      "nil",
			inputSlic: nil,
			equalFunc: func(src int) bool { return src <= 2 },
			wantSlic:  []int{},
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			res := FindAll[int](tc.inputSlic, tc.equalFunc)
			assert.Equal(t, res, tc.wantSlic)
		})
	}
}
