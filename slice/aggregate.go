package slice

import ekitpractice "ekit_practice"

//Max 返回最大值
//假设切片至少存在一个元素
func Max[T ekitpractice.RealNumber](src []T) T {
	res := src[0]
	for _, val := range src {
		if val > res {
			res = val
		}
	}
	return res
}

//Max 返回最小值
//假设切片至少存在一个元素
func Min[T ekitpractice.RealNumber](src []T) T {
	res := src[0]
	for _, val := range src {
		if val < res {
			res = val
		}
	}
	return res
}

//Sum 求和
func Sum[T ekitpractice.RealNumber](src []T) T {
	var res T
	for _, val := range src {
		res += val
	}
	return res
}
