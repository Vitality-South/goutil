// Package slice implements helper functions for working with slices
package slice

import (
	"sort"

	"golang.org/x/exp/constraints"
)

// PointerSlice returns a slice of pointers that point to the items in the
// original slice ([]T ==> []*T)
//
// This is convenient for libraries that often require []*T.
func PointerSlice[T any](in []T) []*T {
	out := make([]*T, len(in))

	for i := range in {
		out[i] = &in[i]
	}

	return out
}

// Contains returns true if the slice in contains an element equal to the value
// val, and false otherwise. The comparison is based on the == operator, so
// elements and the value being searched must be comparable.
//
// # Consider the search methods from the standard library sort package
//
//	instead of this if the slice is already sorted.
//
// # Empty or nil slice always returns false.
func Contains[T comparable](in []T, val T) bool {
	for _, v := range in {
		if v == val {
			return true
		}
	}

	return false
}

// Filter returns a new slice holding only the elements of in that match the
// predicate fn.
// The order of elements in the input slice is preserved in the output slice.
func Filter[T any](in []T, fn func(T) bool) []T {
	var out []T

	for _, v := range in {
		if fn(v) {
			out = append(out, v)
		}
	}

	return out
}

// Map (apply-to-all) applies the provided function fn to each item in the
// input slice in and returns a new slice containing the results.
//
// The predicate function fn must accept an element of the same type as the
// slice and returns a single value. The type of the returned slice elements
// will match the return type of fn. Map does not modify the input slice.
func Map[T any, V any](in []T, fn func(T) V) []V {
	out := make([]V, len(in))

	for i, v := range in {
		out[i] = fn(v)
	}

	return out
}

// Index returns the index of the first occurrence of an element that satisfies
// the predicate fn, or -1 if no such element is found.
//
// The predicate function fn must accept an element of the same type as the
// slice and return a bool indicating whether the element satisfies the
// condition.
//
// Example:
//
//	greaterThanTwo := func(n int) bool { return n > 2 }
//	index := Index[int]([]int{1, 2, 3, 4, 5}, greaterThanTwo) // Returns: 2
func Index[T any](in []T, fn func(T) bool) int {
	for i, v := range in {
		if fn(v) {
			return i
		}
	}

	return -1
}

// All returns true if all elements in the slice in satisfy the predicate fn,
// and false otherwise. If the slice is empty, All returns true.
//
// The predicate function fn must accept an element of the same type as the
// slice and return a bool indicating whether the element satisfies the
// condition.
//
// Example:
//
//	greaterThanTwo := func(n int) bool { return n > 2 }
//	All[int]([]int{3, 4, 5, 6, 7}, greaterThanTwo) // Returns: true
func All[T any](in []T, fn func(T) bool) bool {
	for _, v := range in {
		if !fn(v) {
			return false
		}
	}

	return true
}

// Any returns true if any element of the slice in satisfies the predicate fn,
// and false otherwise. If the slice is empty, Any returns false.
//
// The predicate function fn must accept an element of the same type as the
// slice and return a bool indicating whether the element satisfies the
// condition.
//
// Example:
//
//	hasNegative := func(n int) bool { return n < 0 }
//	Any[int]([]int{1, -2, 3, 0}, hasNegative) // Returns: true
func Any[T any](in []T, fn func(T) bool) bool {
	for _, v := range in {
		if fn(v) {
			return true
		}
	}

	return false
}

// Equal returns true if the slices a and b contain the same elements in the
// same order. Both slices must be of the same length, and their elements must
// be comparable. A nil argument is equivalent to an empty slice.
//
// Example:
//
//	a := []int{1, 2, 3, 3}
//	b := []int{1, 2, 3, 3}
//	Equal(a, b) // Returns: true
func Equal[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if b[i] != v {
			return false
		}
	}

	return true
}

// EqualUnordered returns true if the slices a and b contain the same elements,
// regardless of the order. Both slices must be of the same length, and their
// elements must be comparable. A nil argument is equivalent to an empty slice.
//
// Each distinct element must appear the same number of times in each slice or
// else false is returned.
//
// Example:
//
//	a := []int{3, 2, 1, 1}
//	b := []int{1, 2, 1, 3}
//	EqualUnordered(a, b) // Returns: true
func EqualUnordered[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}

	elementCount := make(map[T]int)

	for _, item := range a {
		elementCount[item]++
	}

	for _, item := range b {
		if count, exists := elementCount[item]; !exists || count == 0 {
			return false
		}

		elementCount[item]--
	}

	return true
}

// Deduplicate returns a new slice with duplicates removed from the input slice
// in. The order of unique elements in the output slice matches their first
// occurrence in the input slice. Elements in the slice must be comparable.
//
// Example usage:
//
//	numbers := []int{1, 2, 2, 3, 4, 3, 5}
//	Deduplicate(numbers) // Returns: []int{1, 2, 3, 4, 5}
func Deduplicate[T comparable](in []T) []T {
	seen := make(map[T]bool)

	var out []T

	for _, v := range in {
		if _, ok := seen[v]; !ok {
			seen[v] = true
			out = append(out, v)
		}
	}

	return out
}

// RemoveElement returns a copy of slice in with the provided element removed.
// Elements in the slice must be comparable.
//
// Example usage:
//
//	numbers := []int{1, 2, 3, 4, 5}
//	RemoveElement(numbers, 3) // Returns: []int{1, 2, 4, 5}
func RemoveElement[T comparable](in []T, elem T) []T {
	var out []T

	for _, v := range in {
		if v != elem {
			out = append(out, v)
		}
	}

	return out
}

// ContainsDuplicates returns true if the slice contains duplicate elements.
func ContainsDuplicates[T comparable](in []T) bool {
	seen := make(map[T]bool)

	for _, v := range in {
		if _, ok := seen[v]; ok {
			return true
		}

		seen[v] = true
	}

	return false
}

// SortedCopy takes a slice in, makes a copy of it, sorts the copy in ascending
// order, and returns the sorted copy. The function leverages the
// constraints.Ordered constraint to allow usage with any type that supports
// ordering operations.
//
// Example usage:
//
//	original := []int{3, 1, 4, 1, 5, 9}
//	sorted := SortedCopy(original)
//	fmt.Println("Original:", original) // Original: [3 1 4 1 5 9]
//	fmt.Println("Sorted:", sorted)     // Sorted: [1 1 3 4 5 9]
func SortedCopy[T constraints.Ordered](in []T) []T {
	tmp := make([]T, len(in))
	copy(tmp, in)

	sort.Slice(tmp, func(i, j int) bool {
		return tmp[i] < tmp[j]
	})

	return tmp
}

// StableSortedCopy takes a slice in, makes a copy of it, sorts the copy in
// ascending order using a stable sort algorithm, and returns the sorted copy.
// The function leverages the constraints.Ordered constraint to allow usage
// with any type that supports ordering operations.
//
// Example usage:
//
//	original := []int{3, 1, 4, 1, 5, 9}
//	sorted := StableSortedCopy(original)
//	fmt.Println("Original:", original) // Original: [3 1 4 1 5 9]
//	fmt.Println("Sorted:", sorted)     // Sorted: [1 1 3 4 5 9]
func StableSortedCopy[T constraints.Ordered](in []T) []T {
	tmp := make([]T, len(in))
	copy(tmp, in)

	sort.SliceStable(tmp, func(i, j int) bool {
		return tmp[i] < tmp[j]
	})

	return tmp
}

// Min returns the minimum value in a slice. If the slice is nil or
// empty, the zero value for the type is returned.
func Min[T constraints.Ordered](in []T) T {
	if len(in) == 0 {
		var zero T
		return zero
	}

	min := in[0]

	for _, v := range in[1:] {
		if v < min {
			min = v
		}
	}

	return min
}

// Max returns the maximum value in a slice. If the slice is nil or
// empty, the zero value for the type is returned.
func Max[T constraints.Ordered](in []T) T {
	if len(in) == 0 {
		var zero T
		return zero
	}

	max := in[0]

	for _, v := range in[1:] {
		if v > max {
			max = v
		}
	}

	return max
}
