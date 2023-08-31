package slice

//返回第一个和dst相等的元素
func Index[T comparable](src []T, dst T) int {
	return IndexFunc[T](src, func(src T) bool {
		return dst == src
	})
}

//返回第一个match返回true的元素
func IndexFunc[T any](src []T, match matchFunc[T]) int {
	for k, v := range src {
		if match(v) {
			return k
		}
	}
	return -1
}

//返回最后一个和dst相等的元素
func LastIndex[T comparable](src []T, dst T) int {
	return LastIndexFunc[T](src, func(src T) bool {
		return src == dst
	})
}

//返回最后一个match返回true的元素
func LastIndexFunc[T any](src []T, match matchFunc[T]) int {
	for i := len(src) - 1; i >= 0; i-- {
		if match(src[i]) {
			return i
		}
	}
	return -1
}

//返回所有和dst相等的元素
func IndexAll[T comparable](src []T, dst T) []int {
	return IndexAllFunc[T](src, func(src T) bool {
		return dst == src
	})
}

//返回所有match返回true的元素下标
func IndexAllFunc[T any](src []T, match matchFunc[T]) []int {
	res := make([]int, 0)
	for k, v := range src {
		if match(v) {
			res = append(res, k)
		}
	}
	return res
}
