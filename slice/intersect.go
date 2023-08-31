package slice

//交集，已去重
func IntersectSet[T comparable](src, dst []T) []T {
	srcMap := toMap[T](src)
	res := make([]T, 0)
	for _, v := range dst {
		if _, ok := srcMap[v]; ok {
			res = append(res, v)
		}
	}
	return deduplicate[T](res)
}

//交集，已去重
func IntersectSetFunc[T any](src, dst []T, equal equalFunc[T]) []T {
	var res []T
	for _, v := range dst {
		if ContainsFunc[T](src, func(t T) bool {
			return equal(t, v)
		}) {
			res = append(res, v)
		}
	}
	return deduplicateFunc[T](res, equal)
}
