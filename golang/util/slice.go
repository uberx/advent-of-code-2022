package util

import "golang.org/x/exp/constraints"

func Reverse[T constraints.Ordered](input []T) {
	for i, j := 0, len(input)-1; i < j; i, j = i+1, j-1 {
		input[i], input[j] = input[j], input[i]
	}
}

func Sum[T constraints.Ordered](input []T) T {
	var sum T
	for i := 0; i < len(input); i++ {
		sum += input[i]
	}
	return sum
}

func IndexOf[T any](input []T, match T, equalityFunc func(t1, t2 T) bool) int {
	for k, v := range input {
		if equalityFunc(match, v) {
			return k
		}
	}
	return -1
}
