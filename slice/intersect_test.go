package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntersectSet(t *testing.T) {
	tests := []struct {
		name string
		src  []int
		dst  []int
		want []int
	}{
		{
			want: []int{1, 3, 5},
			src:  []int{1, 3, 5, 7},
			dst:  []int{1, 3, 5},
			name: "normal test",
		},
		{
			src:  []int{},
			dst:  []int{1, 3, 5, 7},
			want: []int{},
			name: "length of src is 0",
		},
		{
			dst:  []int{1, 3, 5, 7},
			want: []int{},
			name: "src nil",
		},
		{
			src:  []int{1, 3, 5, 5},
			dst:  []int{1, 3, 5},
			want: []int{1, 3, 5},
			name: "exist the same ele in src",
		},
		{
			src:  []int{1, 3, 5, 5},
			dst:  []int{},
			want: []int{},
			name: "dst empty",
		},
		{
			src:  []int{1, 3, 5, 5},
			dst:  []int{},
			want: []int{},
			name: "dst nil",
		},
		{
			src:  []int{1, 1, 3, 5, 7},
			dst:  []int{1, 3, 5, 5},
			want: []int{1, 3, 5},
			name: "exist the same ele in src and dst",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := IntersectSet[int](tt.src, tt.dst)
			assert.ElementsMatch(t, tt.want, res)
		})
	}
}

func TestIntersectSetFunc(t *testing.T) {
	tests := []struct {
		name string
		src  []int
		dst  []int
		want []int
	}{
		{
			want: []int{1, 3, 5},
			src:  []int{1, 3, 5, 7},
			dst:  []int{1, 3, 5},
			name: "normal test",
		},
		{
			src:  []int{},
			dst:  []int{1, 3, 5, 7},
			want: []int{},
			name: "length of src is 0",
		},
		{
			dst:  []int{1, 3, 5, 7},
			want: []int{},
			name: "src nil",
		},
		{
			src:  []int{1, 3, 5, 5},
			dst:  []int{1, 3, 5},
			want: []int{1, 3, 5},
			name: "exist the same ele in src",
		},
		{
			src:  []int{1, 3, 5, 5},
			dst:  []int{},
			want: []int{},
			name: "dst empty",
		},
		{
			src:  []int{1, 3, 5, 5},
			dst:  []int{},
			want: []int{},
			name: "dst nil",
		},
		{
			src:  []int{1, 1, 3, 5, 7},
			dst:  []int{1, 3, 5, 5},
			want: []int{1, 3, 5},
			name: "exist the same ele in src and dst",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := IntersectSetFunc[int](tt.src, tt.dst, func(src, dst int) bool {
				return src == dst
			})
			assert.ElementsMatch(t, tt.want, res)
		})
	}
}
