package util

import "github.com/uberx/advent-of-code-2022/mathy"

type Point struct {
	X int
	Y int
}

func ManhattanDistance(a, b Point) int {
	return mathy.Abs(a.X, b.X) + mathy.Abs(a.Y, b.Y)
}
