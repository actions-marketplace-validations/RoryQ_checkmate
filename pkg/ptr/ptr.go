package ptr

func To[T any](t T) *T {
	return &t
}
