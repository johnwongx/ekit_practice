package slice

//并集，已去重
func UnionSet[T comparable](src, dst []T) []T {
	srcMap := toMap[T](src)
	for _, v := range dst {
		srcMap[v] = struct{}{}
	}

	res := make([]T, 0, len(srcMap))
	for k := range srcMap {
		res = append(res, k)
	}

	return res
}

//并集，已去重
func UnionSetFunc[T any](src, dst []T, equal equalFunc[T]) []T {
	res := make([]T, 0, len(src)+len(dst))
	res = append(res, src...)
	res = append(res, dst...)
	return deduplicateFunc[T](res, equal)
}
