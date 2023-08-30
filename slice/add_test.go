package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	testCase := []struct {
		name      string
		src       []int
		index     int
		addVal    int
		wantErr   error
		wantSlice []int
	}{
		{
			name:      "index 0",
			src:       []int{1, 2, 3, 4, 5},
			index:     0,
			addVal:    233,
			wantErr:   nil,
			wantSlice: []int{233, 1, 2, 3, 4, 5},
		},
		{
			name:      "index middle",
			src:       []int{1, 2, 3, 4, 5},
			index:     2,
			addVal:    233,
			wantErr:   nil,
			wantSlice: []int{1, 2, 233, 3, 4, 5},
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			ret, err := Add[int](tc.src, tc.index, tc.addVal)
			assert.Equal(t, err, tc.wantErr)
			if err != nil {
				return
			}
			assert.Equal(t, ret, tc.wantSlice)
		})
	}
}
