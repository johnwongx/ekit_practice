package slice

//差集， 已去重
func DiffSet[T comparable](src, dst []T) []T {
	srcMap := toMap[T](src)
	for _, val := range dst {
		delete(srcMap, val)
	}

	res := make([]T, 0, len(srcMap))
	for val := range srcMap {
		res = append(res, val)
	}
	return res
}

//差集，已去重
func DiffSetFunc[T any](src, dst []T, equal equalFunc[T]) []T {
	res := make([]T, 0, len(src))
	for _, val := range src {
		if !ContainsFunc[T](dst, func(t T) bool {
			return equal(val, t)
		}) {
			res = append(res, val)
		}
	}
	return deduplicateFunc[T](res, equal)
}
