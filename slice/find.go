package slice

func Find[T any](src []T, equal func(T) bool) (T, bool) {
	for _, val := range src {
		if equal(val) {
			return val, true
		}
	}

	var zero T
	return zero, false
}

func FindAll[T any](src []T, equal func(T) bool) []T {
	//默认符合条件的元素是少数的
	//初始化切片长度为原始长度的1/8, 加一保证至少有一个元素
	res := make([]T, 0, len(src)>>3+1)
	for _, val := range src {
		if equal(val) {
			res = append(res, val)
		}
	}
	return res
}
