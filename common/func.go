package common

type (
	Supplier[T any] func() T
)

// 三目表达式
func IfElse[T any](condition bool, tureVal T, falseVal T) T {
	if condition {
		return tureVal
	}

	return falseVal
}

// 延迟执行求值
func IfElseGetFunc[T any](condition bool, tureSupplier Supplier[T], falseSupplier Supplier[T]) T {
	if condition {
		return tureSupplier()
	}

	return falseSupplier()
}

func IfElseGet[T any](condition bool, t T, falseSupplier Supplier[T]) T {
	if condition {
		return t
	}

	return falseSupplier()
}
