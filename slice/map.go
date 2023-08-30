package slice

func toMap[T comparable](src []T) map[T]struct{} {
	dataMap := make(map[T]struct{}, len(src))
	for _, val := range src {
		dataMap[val] = struct{}{}
	}
	return dataMap
}

func deduplicateFunc[T any](src []T, equal equalFunc[T]) []T {
	res := make([]T, 0, len(src))
	for _, v := range src {
		if !ContainsFunc[T](res, func(t T) bool {
			return equal(v, t)
		}) {
			res = append(res, v)
		}
	}
	return res
}
