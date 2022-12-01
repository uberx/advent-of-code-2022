package util

import (
	"fmt"
	"strconv"
)

func ToInt(arg string) int {
	i, err := strconv.Atoi(arg)
	if err != nil {
		panic(fmt.Sprintf("Cannot convert %s to int: %v", arg, err))
	}
	return i
}
