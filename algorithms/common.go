package algorithms

import "golang.org/x/exp/constraints"

// Returns the minimum of two values
func Min[T constraints.Ordered](first, second T) T {

	if first <= second {
		return first
	}

	return second
}

// Returns the maxmimum of two values
func Max[T constraints.Ordered](first, second T) T {

	if first >= second {
		return first
	}

	return second
}
