package collection

// 可比较类型集合
type Set[T comparable] map[T]struct{}

func NewSet[T comparable](keys ...T) Set[T] {
	if length := len(keys); length != 0 {
		set := make(Set[T], length)
		set.Add(keys...)
		return set
	} else {
		// 初始默认容量
		set := make(Set[T], defaultCapacity)
		return set
	}
}

func (s Set[T]) Add(keys ...T) Set[T] {
	for _, key := range keys {
		if _, ok := s[key]; !ok {
			s[key] = struct{}{}
		}
	}
	return s
}

func (s Set[T]) Remove(keys ...T) Set[T] {
	for _, key := range keys {
		if _, ok := s[key]; ok {
			delete(s, key)
		}
	}

	return s
}

func (s Set[T]) Contains(key T) bool {
	_, ok := s[key]

	return ok
}

func (s Set[T]) IsEmpty() bool {
	return len(s) == 0
}

func (s Set[T]) Count() int {
	return len(s)
}

func (s Set[T]) ToSlice() []T {
	var result = make([]T, 0, len(s))
	for v := range s {
		result = append(result, v)
	}

	return result
}
