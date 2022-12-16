package util

func CopyMap[T comparable, U any](m map[T]U) map[T]U {
	r := map[T]U{}
	for k, v := range m {
		r[k] = v
	}
	return r
}
