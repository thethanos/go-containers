package algorithms

import (
	"golang.org/x/exp/constraints"
)

// counts all occurrences of the given value in the given slice
func Count[T comparable](data []T, value T) int {

	res := 0
	for i := range data {
		if data[i] == value {
			res++
		}
	}

	return res
}

// counts all occurrences of values satisfying the given predicate in the given slice
func CountIf[T comparable](data []T, predicate func(value T) bool) int {

	res := 0
	for i := range data {
		if predicate(data[i]) {
			res++
		}
	}

	return res
}

// finds the given value in the given slice
func Find[T comparable](data []T, value T) int {

	for i := range data {
		if data[i] == value {
			return i
		}
	}

	return -1
}

// finds a value satisfying the given predicate in the given slice
func FindIf[T any](data []T, predicate func(value T) bool) int {

	for i := range data {
		if predicate(data[i]) {
			return i
		}
	}

	return -1
}

// finds a maximal value in the given slice
func MaxElement[T constraints.Ordered](data []T) T {

	var max T
	if len(data) == 0 {
		return max
	}

	max = data[0]
	for i := range data {
		if data[i] > max {
			max = data[i]
		}
	}

	return max

}

// finds a minimal value in the given slice
func MinElement[T constraints.Ordered](data []T) T {

	var min T
	if len(data) == 0 {
		return min
	}

	min = data[0]
	for i := range data {
		if data[i] < min {
			min = data[i]
		}
	}

	return min
}
