// Package numbers implements helper functions for working with numbers
package numbers

import (
	"math"
	"sort"

	"golang.org/x/exp/constraints"
)

// AlmostEqual returns true if the difference of a and b are within epsilon
//
// It is important to choose a correct epsilon value (error tolerance) for a
// given application.
//
// This may or may not be appropriate for a given application and there
// are possible edge cases to consider.
// See https://floating-point-gui.de/errors/comparison/ and
// https://randomascii.wordpress.com/2012/02/25/comparing-floating-point-numbers-2012-edition/
func AlmostEqual(a, b, epsilon float64) bool {
	return math.Abs(a-b) <= epsilon
}

// Abs returns the absolute value of the given integer.
//
// Warning: math.MinInt as input returns a negative number which matches the
// behavior of C and Java. That is to say, math.MinInt as input returns
// math.MinInt as output. Other valid integer values work as expected.
func Abs[T constraints.Integer](n T) T {
	if n < 0 {
		return -n
	}

	return n
}

// IsEven returns true if the given integer is even.
func IsEven[T constraints.Integer](n T) bool {
	return n%2 == 0
}

// IsOdd returns true if the given integer is odd.
func IsOdd[T constraints.Integer](n T) bool {
	return n%2 != 0
}

// OddNumbers returns all odd numbers between "from" and "to" inclusive.
func OddNumbers[T constraints.Integer](from T, to T) []T {
	odds := []T{}

	for i := from; i <= to; i++ {
		if IsOdd(i) {
			odds = append(odds, i)
		}
	}

	return odds
}

// EvenNumbers returns all even numbers between "from" and "to" inclusive.
func EvenNumbers[T constraints.Integer](from T, to T) []T {
	evens := []T{}

	for i := from; i <= to; i++ {
		if IsEven(i) {
			evens = append(evens, i)
		}
	}

	return evens
}

// EvenNumbers returns a slice containing only the even numbers from the input
// slice.
func EvenNumbersInSlice[T constraints.Integer](in []T) []T {
	var evens []T

	for _, n := range in {
		if IsEven(n) {
			evens = append(evens, n)
		}
	}

	return evens
}

// OddNumbers returns a slice containing only the odd numbers from the input
// slice.
func OddNumbersInSlice[T constraints.Integer](in []T) []T {
	var odds []T

	for _, n := range in {
		if IsOdd(n) {
			odds = append(odds, n)
		}
	}

	return odds
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

// Sum calculates the sum of a slice of numbers.
func Sum[T constraints.Integer | constraints.Float](in []T) T {
	var sum T

	for _, v := range in {
		sum += v
	}

	return sum
}

// Average calculates the average of a slice of numbers, returning the result
// as type R. When R is an integer type, the result is truncated towards zero.
// A nil or empty slice as input will return the zero value for R as output.
func Average[T constraints.Integer | constraints.Float, R constraints.Integer | constraints.Float](in []T) R {
	n := len(in)

	if n == 0 {
		return R(0)
	}

	var sum T

	for _, v := range in {
		sum += v
	}

	return R(sum) / R(n)
}

// Median calculates the median value of a slice of numbers and returns the
// result as type R. For slices with an even number of elements, it calculates
// the average of the two middle elements. When R is an integer type, the
// result is truncated towards zero. A nil or empty slice as input will return
// the zero value for R as output.
func Median[T constraints.Integer | constraints.Float, R constraints.Integer | constraints.Float](in []T) R {
	n := len(in)

	if n == 0 {
		return R(0)
	}

	out := make([]T, n)

	copy(out, in)

	sort.Slice(out, func(i, j int) bool { return out[i] < out[j] })

	mid := n / 2

	if n%2 == 0 {
		return R((out[mid-1] + out[mid]) / T(2))
	}

	return R(out[mid])
}

// Product calculates the product of all elements in the input slice.
// A nil or empty slice returns the multiplicative identity, 1.
func Product[T constraints.Integer | constraints.Float](in []T) T {
	product := T(1)

	for _, v := range in {
		product *= v
	}

	return product
}
