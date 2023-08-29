package slice

import (
	"ekit_practice/internal/errs"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDelete(t *testing.T) {
	testCase := []struct {
		name      string
		src       []int
		index     int
		wantErr   error
		wantSlice []int
		wantVal   int
	}{
		{
			name:      "index 0",
			src:       []int{1, 2, 3, 4, 5},
			index:     0,
			wantErr:   nil,
			wantSlice: []int{2, 3, 4, 5},
			wantVal:   1,
		},
		{
			name:      "index middle",
			src:       []int{1, 2, 3, 4, 5},
			index:     2,
			wantErr:   nil,
			wantSlice: []int{1, 2, 4, 5},
			wantVal:   3,
		},
		{
			name:      "index last",
			src:       []int{1, 2, 3, 4, 5},
			index:     4,
			wantErr:   nil,
			wantSlice: []int{1, 2, 3, 4},
			wantVal:   5,
		},
		{
			name:    "index less than 0",
			src:     []int{1, 2, 3, 4, 5},
			index:   -1,
			wantErr: errs.NewErrIndexOutOfRange(5, -1),
		},
		{
			name:    "index out of range",
			src:     []int{1, 2, 3, 4, 5},
			index:   5,
			wantErr: errs.NewErrIndexOutOfRange(5, 5),
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			ret, res, err := Delete[int](tc.src, tc.index)
			assert.Equal(t, err, tc.wantErr)
			if err != nil {
				return
			}
			assert.Equal(t, ret, tc.wantSlice)
			assert.Equal(t, res, tc.wantVal)
		})
	}
}
