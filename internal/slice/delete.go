package slice

import "ekit_practice/internal/errs"

func Delete[T any](src []T, index int) ([]T, T, error) {
	length := len(src)
	if index < 0 || index >= length {
		var zero T
		return src, zero, errs.NewErrIndexOutOfRange(length, index)
	}
	res := src[index]
	for i := index; i < length-1; i++ {
		src[i] = src[i+1]
	}
	src = src[:length-1]
	return src, res, nil
}
