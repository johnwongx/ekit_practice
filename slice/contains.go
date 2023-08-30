package slice

//判断src是否包含dst
func Contains[T comparable](src []T, dst T) bool {
	return ContainsFunc[T](src, func(src T) bool {
		return dst == src
	})
}

//判断src是否包含dst
func ContainsFunc[T any](src []T, equal func(T) bool) bool {
	for _, val := range src {
		if equal(val) {
			return true
		}
	}
	return false
}

//判断src是否包含dst中任何一个元素
func ContainsAny[T comparable](src, dst []T) bool {
	srcMap := toMap[T](src)
	for _, val := range dst {
		if _, ok := srcMap[val]; ok {
			return true
		}
	}
	return false
}

//判断src是否包含dst中任何一个元素
func ContainsAnyFunc[T any](src, dst []T, equal equalFunc[T]) bool {
	for _, l := range src {
		for _, r := range dst {
			if equal(l, r) {
				return true
			}
		}
	}
	return false
}

//判断src是否包含dst中任何所有元素
func ContainsAll[T comparable](src, dst []T) bool {
	srcMap := toMap[T](src)
	for _, val := range dst {
		if _, ok := srcMap[val]; !ok {
			return false
		}
	}
	return true
}

//判断src是否包含dst中任何一个元素
func ContainsAllFunc[T any](src, dst []T, equal equalFunc[T]) bool {
	for _, dstVal := range dst {
		if !ContainsFunc[T](src, func(t T) bool {
			return equal(t, dstVal)
		}) {
			return false
		}
	}
	return true
}
