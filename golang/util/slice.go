package util

import "golang.org/x/exp/constraints"

func Reverse[T constraints.Ordered](input []T) {
	for i, j := 0, len(input)-1; i < j; i, j = i+1, j-1 {
		input[i], input[j] = input[j], input[i]
	}
}
