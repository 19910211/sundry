package collection

import (
	"golang.org/x/exp/constraints"
	"sort"
)

type (
	AnySliceSortOption[T any] func(a *AnySliceSort[T])
	// SortSliceWrap 对象排序使用
	AnySliceSort[T any] struct {
		data        []T
		lenHandler  func(data []T) int
		swapHandler func(data []T, i, j int)
		lessHandler func(data []T, i, j int) bool
	}
)

//
// NewAnySliceSort
// @Description:
// @param data
// @param opts
// @return *AnySliceSort[T]
//
//  var list []*ClassInfo
// 	NewAnySliceSort(List, WithLessHandler(func(data []*ClassInfo, i, j int) bool {
//		return data[i].Id > data[j].Id
//	})).Sort()
//
//
func NewAnySliceSort[T any](data []T, opts ...AnySliceSortOption[T]) *AnySliceSort[T] {
	o := &AnySliceSort[T]{
		data:        data,
		lenHandler:  func(data []T) int { return len(data) },
		swapHandler: func(data []T, i, j int) { data[i], data[j] = data[j], data[i] },
		lessHandler: func(data []T, i, j int) bool { return true }, // default true 请自己实现
	}

	for _, opt := range opts {
		opt(o)
	}

	return o
}

func (s *AnySliceSort[T]) Sort() {
	sort.Sort(s)
}

func (s *AnySliceSort[T]) Len() int {
	return s.lenHandler(s.data)
}

func (s *AnySliceSort[T]) Less(i, j int) bool {
	return s.lessHandler(s.data, i, j)
}

func (s *AnySliceSort[T]) Swap(i, j int) {
	s.swapHandler(s.data, i, j)
}

// 升序还是降序
func WithLessHandler[T any](handler func(data []T, i, j int) bool) AnySliceSortOption[T] {
	return func(o *AnySliceSort[T]) {
		o.lessHandler = handler
	}
}
func WithSwapHandler[T any](handler func(data []T, i, j int)) AnySliceSortOption[T] {
	return func(o *AnySliceSort[T]) {
		o.swapHandler = handler
	}
}

func WithLenHandler[T any](handler func(data []T) int) AnySliceSortOption[T] {
	return func(o *AnySliceSort[T]) {
		o.lenHandler = handler
	}
}

const (
	SortAsc  = false
	SortDesc = true
)

//
// Sort
// @Description: 结构体排序
// @param data   排序的切片数据
// @param orderByVal 获取结构中 需要比较的值
// @param isDesc 是否倒序
//
//  var rooms []*model.Room
// 	Sort(rooms, func(t *model.Room) int64 {return t.RoomId}, SortAsc)
//
func Sort[T any, V constraints.Ordered](data []T, orderByVal func(T) V, isDesc ...bool) {

	if len(data) < 2 {
		return
	}

	// 升序
	var less = func(v1, v2 T) bool {
		return orderByVal(v1) < orderByVal(v2)
	}

	// 降序
	if len(isDesc) > 0 && isDesc[0] {
		less = func(v1, v2 T) bool {
			return orderByVal(v1) > orderByVal(v2)
		}
	}

	// 开始排序

	// 官方api
	//sort.Slice(data, func(i, j int) bool {
	//	return less(data[i], data[j])
	//})

	// 继承 Interface 接口自定义实现
	NewAnySliceSort(data, WithLessHandler(func(data []T, i, j int) bool {
		return less(data[i], data[j])
	})).Sort()

}
