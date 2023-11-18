package test

func PanicOr[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}

func P[T any](v T) *T {
	p := v
	return &p
}
