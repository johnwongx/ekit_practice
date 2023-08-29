package slice

import "ekit_practice/internal/errs"

func Add[T any](src []T, index int, elem T) ([]T, error) {
	length := len(src)
	if index < 0 || index >= length {
		return src, errs.NewErrIndexOutOfRange(length, index)
	}

	var zero T
	src = append(src, zero)
	for i := length; i > index; i-- {
		if i-1 >= 0 {
			src[i] = src[i-1]
		}
	}
	src[index] = elem
	return src, nil
}
