package utils

import "golang.org/x/exp/constraints"

func Clamp[T constraints.Integer](value, high, low T) T {
	if value >= high {
		return high
	}

	if value <= low {
		return low
	}

	return value
}
