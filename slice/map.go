package slice

//将一种切片类型T映射为另一种切片类型R
func Map[T any, R any](src []T, m func(int, T) R) []R {
	res := make([]R, len(src))
	for k, v := range src {
		res[k] = m(k, v)
	}
	return res
}

//将一种切片类型T映射为另一种R
//并且按照制定的方法过滤元素
func FilterMap[T any, R any](src []T, m func(int, T) (bool, R)) []R {
	res := make([]R, 0, len(src))
	for k, v := range src {
		if ok, dst := m(k, v); ok {
			res = append(res, dst)
		}
	}
	return res
}

// 将切片转化为map
func toMap[T comparable](src []T) map[T]struct{} {
	dataMap := make(map[T]struct{}, len(src))
	for _, val := range src {
		dataMap[val] = struct{}{}
	}
	return dataMap
}

// 对切片去重
func deduplicate[T comparable](src []T) []T {
	srcMap := toMap[T](src)
	res := make([]T, 0, len(srcMap))
	for k := range srcMap {
		res = append(res, k)
	}
	return res
}

// 对切片去重
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
