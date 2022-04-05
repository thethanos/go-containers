package algorithms

import "golang.org/x/exp/constraints"

// Returns minimal value
func Min[T constraints.Ordered](first, second T) T {

	if first <= second {
		return first
	}

	return second
}

// Returns maxmimal value
func Max[T constraints.Ordered](first, second T) T {

	if first >= second {
		return first
	}

	return second
}
