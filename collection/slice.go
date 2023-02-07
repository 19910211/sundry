package collection

import (
	"sort"
	"sundry/common"

	"golang.org/x/exp/constraints"
)

// SliceFlipUnique 切片去重
func SliceFlipUnique[T comparable](arr []T) []T {
	var (
		set = map[T]struct{}{}
		ret []T
	)
	for _, t := range arr {
		if _, ok := set[t]; ok {
			continue
		}
		ret = append(ret, t)
		set[t] = struct{}{}
	}
	return ret
}

// SliceUnique 切片去重
func SliceUnique[T constraints.Ordered](a []T) []T {
	if len(a) == 0 {
		return a
	}
	SliceSort(a)
	i := 0
	for j := 1; j < len(a); j++ {
		if a[i] != a[j] {
			i++
			a[i] = a[j]
		}
	}
	var t T
	for j := i + 1; j < len(a); j++ {
		a[j] = t
	}
	return a[:i+1]
}

//
// SliceAnyUnique[T any, V constraints.Ordered]
// @Description: 针对 所有类型切片 去重
// @param a
// @param getCompareVal
// @param extension
// @return []T
//
func SliceAnyUnique[T any, V constraints.Ordered](a []T, getCompareVal func(T) V, extension ...bool) []T {
	if len(a) == 0 {
		return a
	}
	// 如果没有传入排序参数就默认 升序
	Sort(a, getCompareVal, common.IfElseGet(len(extension) == 0, SortAsc, func() bool { return extension[0] }))

	i := 0
	for j := 1; j < len(a); j++ {
		if getCompareVal(a[i]) != getCompareVal(a[j]) {
			i++
			a[i] = a[j]
		}
	}

	// 如果传入了 删除多余的数据就清除 (主要是用于常驻内存)
	if len(extension) > 1 && extension[1] {
		// 把多余的数据删除 防止内存泄露
		var t T
		for j := i + 1; j < len(a); j++ {
			a[j] = t
		}
	}

	return a[:i+1]
}

// SliceFilter 切片过滤
func SliceFilter[T any](a []T, handler func(T) bool) []T {
	var ret []T
	for _, t := range a {
		if !handler(t) {
			ret = append(ret, t)
		}
	}

	return ret
}

func SliceSort[T constraints.Ordered](s []T) {
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})
}

// SliceEq 比较
func SliceEq[T constraints.Ordered](a, b []T) bool {
	if (a == nil) != (b == nil) || len(a) != len(b) {
		return false
	}
	SliceSort(a)
	SliceSort(b)
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// Reduce reduces a []T1 to a single value using a reduction function.
func SliceReduce[T1, T2 any](s []T1, initializer T2, f func(T2, T1) T2) T2 {
	r := initializer
	for _, v := range s {
		r = f(r, v)
	}
	return r
}

//
// SliceWalk
// @Description: 函数使用用户自定义函数对数组中的每个元素做回调处理
func SliceWalk[T, V any](arr []T, handler func(int, T) (V, bool)) []V {
	var ret = make([]V, 0, len(arr))
	for i, o := range arr {
		if v, ok := handler(i, o); ok {
			ret = append(ret, v)
		}
	}
	return ret
}

//
// SliceIntersect
// @Description: 获取多个切片的交集
func SliceIntersect[T comparable](arr ...[]T) []T {
	if len(arr) == 0 {
		return nil
	} else if len(arr) == 1 {
		return arr[0]
	}

	var (
		set          = NewSet[T]()
		intersection = NewSet[T](arr[0]...)
	)
	for i := 1; i < len(arr); i++ {
		list := arr[i]
		set = intersection
		intersection = NewSet[T]()

		for _, t := range list {
			if set.Contains(t) {
				intersection.Add(t)
			}

		}
		// 没有交集 不用继续计算
		if intersection.Count() == 0 {
			break
		}
	}
	return intersection.ToSlice()
}

// SliceDifference 获取切片差集
// 获取 a在b 中的差集
func SliceDifference[T comparable](a, b []T) []T {
	if len(a) == 0 || len(b) == 0 {
		return a
	}

	var (
		r   []T
		set = NewSet[T](b...)
	)
	for _, t := range a {
		if !set.Contains(t) {
			r = append(r, t)
		}
	}

	return r
}
