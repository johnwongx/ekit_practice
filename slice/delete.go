package slice

import "ekit_practice/internal/slice"

func Delete[T any](src []T, index int) ([]T, T, error) {
	return slice.Delete[T](src, index)
}
