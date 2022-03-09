package got

type Result[T any] struct {
	data T
	err  error
}

func (r Result[T]) Data() T {
	return r.data
}

func (r Result[T]) Err() error {
	return r.err
}
