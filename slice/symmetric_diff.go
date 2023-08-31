package slice

//对称差集,已去重
//删除共同元素，剩余元素取并集
func SymmetricDiff[T comparable](src, dst []T) []T {
	srcMap, dstMap := toMap[T](src), toMap[T](dst)
	for k := range dstMap {
		if _, ok := srcMap[k]; ok {
			delete(srcMap, k)
		} else {
			srcMap[k] = struct{}{}
		}
	}

	res := make([]T, 0, len(srcMap))
	for k := range srcMap {
		res = append(res, k)
	}

	return res
}

func SymmetricDiffFunc[T any](src, dst []T, equal equalFunc[T]) []T {
	res := []T{}

	//找出在src不在dst的元素
	for _, v := range src {
		if !ContainsFunc[T](dst, func(t T) bool {
			return equal(t, v)
		}) {
			res = append(res, v)
		}
	}

	//找出在dst不在src的元素
	for _, v := range dst {
		if !ContainsFunc[T](src, func(t T) bool {
			return equal(t, v)
		}) {
			res = append(res, v)
		}
	}

	return deduplicateFunc[T](res, equal)
}
