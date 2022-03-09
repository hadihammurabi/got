package got

func Ok[T any](data T) Result[T] {
	return Result[T]{data: data}
}

func Err(err error) Result[interface{}] {
	return Result[interface{}]{err: err}
}
