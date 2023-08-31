package slice

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	testCase := []struct {
		name     string
		intput   []int
		wantSlic []string
	}{
		{
			name:     "normal",
			intput:   []int{1, 2, 3},
			wantSlic: []string{"1", "2", "3"},
		},
		{
			name:     "empty",
			intput:   []int{},
			wantSlic: []string{},
		},
		{
			name:     "nil",
			intput:   nil,
			wantSlic: []string{},
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			res := Map[int, string](tc.intput, func(i1, i2 int) string {
				str := strconv.Itoa(i2)
				return str
			})
			assert.Equal(t, res, tc.wantSlic)
		})
	}
}

func TestFilterMap(t *testing.T) {
	testCase := []struct {
		name     string
		intput   []int
		wantSlic []string
	}{
		{
			name:     "normal",
			intput:   []int{1, 2, 3},
			wantSlic: []string{"1", "2"},
		},
		{
			name:     "empty",
			intput:   []int{},
			wantSlic: []string{},
		},
		{
			name:     "nil",
			intput:   nil,
			wantSlic: []string{},
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			res := FilterMap[int, string](tc.intput, func(i1, i2 int) (bool, string) {
				var res string
				if i2 > 2 {
					return false, res
				}
				str := strconv.Itoa(i2)
				return true, str
			})
			assert.Equal(t, res, tc.wantSlic)
		})
	}
}

func TestToMap(t *testing.T) {
	testCase := []struct {
		name    string
		intput  []int
		wantMap map[int]struct{}
	}{
		{
			name:   "normal",
			intput: []int{1, 2, 3},
			wantMap: map[int]struct{}{
				1: {},
				2: {},
				3: {},
			},
		},
		{
			name:   "exist duplicate",
			intput: []int{1, 2, 2, 3, 3},
			wantMap: map[int]struct{}{
				1: {},
				2: {},
				3: {},
			},
		},
		{
			name:   "all duplicate",
			intput: []int{1, 1, 1},
			wantMap: map[int]struct{}{
				1: {},
			},
		},
		{
			name:    "empty",
			intput:  []int{},
			wantMap: map[int]struct{}{},
		},
		{
			name:    "nil",
			intput:  nil,
			wantMap: map[int]struct{}{},
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			res := toMap[int](tc.intput)
			assert.Equal(t, res, tc.wantMap)
		})
	}
}

func TestDeduplicateFunc(t *testing.T) {
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
			assert.ElementsMatch(t, res, tc.wantSlic)
		})
	}
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
			name:     "all duplicate",
			input:    []int{1, 1, 1, 1},
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
			res := deduplicate[int](tc.input)
			assert.ElementsMatch(t, res, tc.wantSlic)
		})
	}
}
