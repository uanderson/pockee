package utils

func Deref[T any](ptr *T, def T) T {
	if ptr != nil {
		return *ptr
	}

	return def
}
