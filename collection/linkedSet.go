package collection

type (
	LinkedSetOption[T any, K comparable] func(*LinkedSet[T, K])

	HashHandler[T any, K comparable] func(T) K // 获取判断对象是否相等的 值

	// Set 有序集合， 按插入的顺序
	LinkedSet[T any, K comparable] struct {
		isCover  bool //
		initData []T

		hash func(T) K
		data *LinkedMap[K, T]
	}
)

func NewLinkedSet[T any, K comparable](hashHandler HashHandler[T, K], opts ...LinkedSetOption[T, K]) *LinkedSet[T, K] {

	o := LinkedSet[T, K]{
		hash: hashHandler,
	}

	for _, opt := range opts {
		opt(&o)
	}

	cap := len(o.initData)
	if cap == 0 {
		cap = defaultCapacity
	}
	o.data = NewLinkedMap[K, T](
		WithMapIsCover[K, T](o.isCover), // 设置是否覆盖
		WithMapCapacity[K, T](cap),      // 设置容量
	)
	o.Add(o.initData...)
	o.initData = nil

	return &o
}

func WithIsCover[T any, K comparable](isCover bool) LinkedSetOption[T, K] {
	return func(o *LinkedSet[T, K]) {
		o.data.isCover = isCover
	}
}

func WithInitData[T any, K comparable](arr ...T) LinkedSetOption[T, K] {
	return func(s *LinkedSet[T, K]) {
		s.initData = arr
	}
}

func (s *LinkedSet[T, K]) Add(keys ...T) *LinkedSet[T, K] {
	for _, key := range keys {
		code := s.hash(key)
		s.data.Put(code, key)
	}
	return s
}

func (s *LinkedSet[T, K]) Remove(keys ...T) *LinkedSet[T, K] {
	for _, key := range keys {
		code := s.hash(key)
		s.data.Remove(code)
	}
	return s
}

func (s *LinkedSet[T, K]) Contains(key T) bool {
	return s.data.Contains(s.hash(key))
}

func (s *LinkedSet[T, K]) IsEmpty() bool {
	return s.Count() == 0
}

func (s *LinkedSet[T, K]) Count() int {
	return s.data.Count()
}

func (s *LinkedSet[T, K]) ForEach(consumer func(v T)) {
	s.data.ForEach(func(_ K, t T) { consumer(t) })
}

//
// ToSlice
// @Description: 转换成切片
func (s *LinkedSet[T, K]) ToSlice() []T {
	return s.data.ValueSlice()
}
