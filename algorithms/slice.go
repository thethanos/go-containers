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
// returns -1 if nothing found or the slice is empty
func Find[T comparable](data []T, value T) int {

	for i := range data {
		if data[i] == value {
			return i
		}
	}

	return -1
}

// finds a value satisfying the given predicate in the given slice
// returns -1 if nothing found or the slice is empty
func FindIf[T any](data []T, predicate func(value T) bool) int {

	for i := range data {
		if predicate(data[i]) {
			return i
		}
	}

	return -1
}

// finds a maximal value in the given slice and returns its index
// returns -1 if the slice is empty
func MaxElement[T constraints.Ordered](data []T) int {

	if len(data) == 0 {
		return -1
	}

	maxIndex := -1
	maxValue := data[0]
	for i := range data {
		if data[i] > maxValue {
			maxIndex = i
			maxValue = data[i]
		}
	}

	return maxIndex
}

// finds a minimal value in the given slice and returns its index
// returns -1 if the slice is empty
func MinElement[T constraints.Ordered](data []T) int {

	if len(data) == 0 {
		return -1
	}

	minIndex := -1
	minValue := data[0]
	for i := range data {
		if data[i] < minValue {
			minIndex = i
			minValue = data[i]
		}
	}

	return minIndex
}
