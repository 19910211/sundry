package collection

// 获取map keys
func GetKeys[K comparable, V any](m map[K]V) []K {
	return MapKeys(m)
}

func MapKeys[K comparable, V any](m map[K]V) []K {
	var (
		ret = make([]K, 0, len(m))
	)
	for k := range m {
		ret = append(ret, k)
	}
	return ret
}
func MapValues[K comparable, V any](m map[K]V) []V {
	var (
		ret = make([]V, 0, len(m))
	)
	for _, v := range m {
		ret = append(ret, v)
	}
	return ret
}

// map K  V 反转
func MapReverse[K, V comparable](m map[K]V) map[V]K {
	t := make(map[V]K, len(m))
	for k, v := range m {
		t[v] = k
	}
	return t
}

//
// SliceMap
// @Description: 为数组的每个元素应用回调函数，SliceMap与SliceWalk的使用方法基本相同
func SliceMap[T, V any](arr []T, handler func(T) V) []V {
	return SliceWalk(arr, func(i int, t T) (V, bool) {
		return handler(t), true
	})
}

func MapWalk[K comparable, V, T any](m map[K]V, handler func(K, V) (T, bool)) []T {
	var ret = make([]T, 0, len(m))
	for k, v := range m {
		if t, ok := handler(k, v); ok {
			ret = append(ret, t)
		}
	}
	return ret
}
