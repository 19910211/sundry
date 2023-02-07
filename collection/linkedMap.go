package collection

type (
	LinkedMapOption[K comparable, V any] func(*LinkedMap[K, V])

	//有序map 插入顺序
	LinkedMap[K comparable, V any] struct {
		isCover bool                 // 如果相同 是否覆盖  如果要设置覆盖，LinkedSetOption 参数 WithIsCover 需要在 WithInitData 的前面
		data    map[K]*mapNode[K, V] // 数据存储
		hash    func(V) K            // 获取比较值的函数

		initCap int //

		head *mapNode[K, V] // 头节点
		tail *mapNode[K, V] // 尾节点
	}

	mapNode[K comparable, V any] struct {
		key  K              // key
		val  V              // value
		pre  *mapNode[K, V] // 上一个节点
		next *mapNode[K, V] // 下一个节点
	}
)

//
// NewLinkedMap
// @Description: 创建
func NewLinkedMap[K comparable, V any](opts ...LinkedMapOption[K, V]) *LinkedMap[K, V] {
	tail := &mapNode[K, V]{} // 尾节点
	head := tail             // 头节点

	o := LinkedMap[K, V]{
		head: head,
		tail: tail,
	}

	for _, opt := range opts {
		opt(&o)
	}

	cap := o.initCap
	if cap == 0 {
		cap = defaultCapacity
	}
	o.data = make(map[K]*mapNode[K, V], cap)

	return &o
}

//
// WithMapIsCover
// @Description: 设置key是否覆盖 默认抛弃
func WithMapIsCover[K comparable, V any](isCover bool) LinkedMapOption[K, V] {
	return func(o *LinkedMap[K, V]) {
		o.isCover = isCover
	}
}

//
// WithMapCapacity
// @Description: 设置初始容量
func WithMapCapacity[K comparable, V any](cap int) LinkedMapOption[K, V] {
	return func(o *LinkedMap[K, V]) {
		o.initCap = cap
	}
}

func (m *LinkedMap[K, V]) Put(k K, v V) *LinkedMap[K, V] {

	if n, ok := m.data[k]; ok {
		//isCover(覆盖)true 则直接覆盖源值
		if m.isCover {
			n.val = v
		}
	} else {
		v := &mapNode[K, V]{key: k, val: v}
		m.data[k] = v

		v.pre, m.tail.next = m.tail, v
		// 尾节点指向新的节点
		m.tail = v
	}

	return m
}

func (m *LinkedMap[K, V]) PutAll(arr map[K]V) *LinkedMap[K, V] {
	for k, v := range arr {
		m.Put(k, v)
	}
	return m
}

func (m *LinkedMap[K, V]) Remove(keys ...K) *LinkedMap[K, V] {
	for _, key := range keys {
		if o, ok := m.data[key]; ok {
			if o == m.tail {
				o.pre.next = nil
				m.tail = o.pre
			} else {
				o.pre.next = o.next
				o.next.pre = o.pre
			}

			delete(m.data, key)
		}
	}
	return m
}

func (m *LinkedMap[K, V]) Contains(key K) bool {
	_, ok := m.data[key]
	return ok
}

func (m *LinkedMap[K, V]) IsEmpty() bool {
	return m.Count() == 0
}

func (m *LinkedMap[K, V]) Count() int {
	return len(m.data)
}

//
// ForEach
// @Description: 遍历输出
// @receiver s
// @param consumer
func (m *LinkedMap[K, V]) ForEach(consumer func(k K, v V)) {
	for current := m.head.next; current != nil; current = current.next {
		consumer(current.key, current.val)
	}
}

//
// KeySlice
// @Description: 返回 key 切片
// @receiver m
// @return []K
//
func (m *LinkedMap[K, V]) KeySlice() []K {
	var result = make([]K, 0, m.Count())

	m.ForEach(func(k K, _ V) { result = append(result, k) })

	return result
}

//
// ValueSlice
// @Description: 返回值切片
func (m *LinkedMap[K, V]) ValueSlice() []V {
	var result = make([]V, 0, m.Count())

	m.ForEach(func(_ K, v V) { result = append(result, v) })

	return result
}
