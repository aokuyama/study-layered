package util

func PanicOr[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}
