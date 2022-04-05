package algorithms

import (
	"golang.org/x/exp/constraints"
)

// Counts all occurrences of the given value in the given slice.
func Count[T comparable](data []T, value T) int {

	res := 0
	for i := range data {
		if data[i] == value {
			res++
		}
	}

	return res
}

// Counts all occurrences of values satisfying the given predicate in the given slice.
func CountIf[T comparable](data []T, predicate func(value T) bool) int {

	res := 0
	for i := range data {
		if predicate(data[i]) {
			res++
		}
	}

	return res
}

// Finds the given value in the given slice, returns true if it is found.
// Returns false if nothing found or the slice is empty.
func Find[T comparable](data []T, value T) bool {

	for i := range data {
		if data[i] == value {
			return true
		}
	}

	return false
}

// Finds a value satisfying the given predicate in the given slice and returns it.
// Returns defaultValue if nothing found or the slice is empty.
func FindIf[T any](data []T, defaultValue T, predicate func(value T) bool) T {

	for _, value := range data {
		if predicate(value) {
			return value
		}
	}

	return defaultValue
}

// Fills the given slice with the given value
func Fill[T any](data []T, value T) {

	for i := range data {
		data[i] = value
	}
}

// Finds a maximal value in the given slice and returns it.
// Returns defaultValue if the slice is empty
func MaxElement[T constraints.Ordered](data []T, defaultValue T) T {

	if len(data) == 0 {
		return defaultValue
	}

	maxValue := data[0]
	for i := range data {
		if data[i] > maxValue {
			maxValue = data[i]
		}
	}

	return maxValue
}

// Finds a minimal value in the given slice and returns it.
// Returns defaultValue if the slice is empty
func MinElement[T constraints.Ordered](data []T, defaultValue T) T {

	if len(data) == 0 {
		return defaultValue
	}

	minValue := data[0]
	for i := range data {
		if data[i] < minValue {
			minValue = data[i]
		}
	}

	return minValue
}
