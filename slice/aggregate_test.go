package slice

import (
	ekitpractice "ekit_practice"
	"testing"

	"github.com/stretchr/testify/assert"
)

type integer int

func TestMax(t *testing.T) {
	testCase := []struct {
		name    string
		src     []integer
		wantVal integer
	}{
		{
			name:    "value",
			src:     []integer{1},
			wantVal: 1,
		},
		{
			name:    "values",
			src:     []integer{1, 2, 3, 4},
			wantVal: 4,
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			res := Max[integer](tc.src)
			assert.Equal(t, res, tc.wantVal)
		})
	}

	assert.Panics(t, func() {
		Max[int](nil)
	})
	assert.Panics(t, func() {
		Max[int]([]int{})
	})

	testMaxTypes[int](t)
	testMaxTypes[int8](t)
	testMaxTypes[int16](t)
	testMaxTypes[int32](t)
	testMaxTypes[int64](t)
	testMaxTypes[uint](t)
	testMaxTypes[uint8](t)
	testMaxTypes[uint16](t)
	testMaxTypes[uint32](t)
	testMaxTypes[uint64](t)
	testMaxTypes[float32](t)
	testMaxTypes[float64](t)
}

func TestMin(t *testing.T) {
	testCase := []struct {
		name    string
		src     []integer
		wantVal integer
	}{
		{
			name:    "value",
			src:     []integer{1},
			wantVal: 1,
		},
		{
			name:    "values",
			src:     []integer{1, 2, 3, 4},
			wantVal: 1,
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			res := Min[integer](tc.src)
			assert.Equal(t, res, tc.wantVal)
		})
	}

	assert.Panics(t, func() {
		Min[int](nil)
	})
	assert.Panics(t, func() {
		Min[int]([]int{})
	})

	testMinTypes[int](t)
	testMinTypes[int8](t)
	testMinTypes[int16](t)
	testMinTypes[int32](t)
	testMinTypes[int64](t)
	testMinTypes[uint](t)
	testMinTypes[uint8](t)
	testMinTypes[uint16](t)
	testMinTypes[uint32](t)
	testMinTypes[uint64](t)
	testMinTypes[float32](t)
	testMinTypes[float64](t)
}

func TestSum(t *testing.T) {
	testCase := []struct {
		name    string
		src     []integer
		wantVal integer
	}{
		{
			name: "nil",
		},
		{
			name:    "empty",
			src:     []integer{},
			wantVal: 0,
		},
		{
			name:    "value",
			src:     []integer{1},
			wantVal: 1,
		},
		{
			name:    "values",
			src:     []integer{1, 2, 3, 4},
			wantVal: 10,
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			res := Sum[integer](tc.src)
			assert.Equal(t, res, tc.wantVal)
		})
	}

	testSumTypes[int](t)
	testSumTypes[int8](t)
	testSumTypes[int16](t)
	testSumTypes[int32](t)
	testSumTypes[int64](t)
	testSumTypes[uint](t)
	testSumTypes[uint8](t)
	testSumTypes[uint16](t)
	testSumTypes[uint32](t)
	testSumTypes[uint64](t)
	testSumTypes[float32](t)
	testSumTypes[float64](t)
}

func testMaxTypes[T ekitpractice.RealNumber](t *testing.T) {
	res := Max[T]([]T{1, 2, 3})
	assert.Equal(t, res, T(3))
}

func testMinTypes[T ekitpractice.RealNumber](t *testing.T) {
	res := Min[T]([]T{1, 2, 3})
	assert.Equal(t, res, T(1))
}

func testSumTypes[T ekitpractice.RealNumber](t *testing.T) {
	res := Sum[T]([]T{1, 2, 3})
	assert.Equal(t, res, T(6))
}
