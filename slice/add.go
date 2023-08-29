package slice

import "ekit_practice/internal/slice"

func Add[T any](src []T, index int, addVal T) ([]T, error) {
	return slice.Add[T](src, index, addVal)
}
