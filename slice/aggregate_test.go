package slice

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

type Integer int

func TestMax(t *testing.T) {
	testCase := []struct {
		name    string
		src     []Integer
		wantVal Integer
	}{
		{
			name:    "value",
			src:     []Integer{1},
			wantVal: 1,
		},
		{
			name:    "values",
			src:     []Integer{1, 2, 3, 4},
			wantVal: 4,
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			res := Max[Integer](tc.src)
			assert.Equal(t, res, tc.wantVal)
		})
	}
}

func TestMin(t *testing.T) {
	testCase := []struct {
		name    string
		src     []Integer
		wantVal Integer
	}{
		{
			name:    "value",
			src:     []Integer{1},
			wantVal: 1,
		},
		{
			name:    "values",
			src:     []Integer{1, 2, 3, 4},
			wantVal: 1,
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			res := Min[Integer](tc.src)
			assert.Equal(t, res, tc.wantVal)
		})
	}
}

func TestSum(t *testing.T) {
	testCase := []struct {
		name    string
		src     []Integer
		wantVal Integer
	}{
		{
			name:    "empty",
			src:     []Integer{},
			wantVal: 0,
		},
		{
			name:    "value",
			src:     []Integer{1},
			wantVal: 1,
		},
		{
			name:    "values",
			src:     []Integer{1, 2, 3, 4},
			wantVal: 10,
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			res := Sum[Integer](tc.src)
			assert.Equal(t, res, tc.wantVal)
		})
	}
}
