package slice

func Reverse[T any](src []T) []T {
	res := make([]T, 0, len(src))
	for i := len(src) - 1; i >= 0; i-- {
		res = append(res, src[i])
	}
	return res
}

func ReverseSelf[T any](src []T) {
	for i, j := 0, len(src)-1; i < j; i, j = i+1, j-1 {
		src[i], src[j] = src[j], src[i]
	}
}
