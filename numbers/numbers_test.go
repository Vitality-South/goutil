package numbers

import (
	"math"
	"reflect"
	"testing"
)

func TestIsEven(t *testing.T) {
	t.Parallel()

	testStructs := []struct {
		Input    int
		Expected bool
	}{
		{
			Input:    0,
			Expected: true,
		},
		{
			Input:    1,
			Expected: false,
		},
		{
			Input:    -1,
			Expected: false,
		},
		{
			Input:    2,
			Expected: true,
		},
		{
			Input:    -2,
			Expected: true,
		},
		{
			Input:    -98,
			Expected: true,
		},
		{
			Input:    98,
			Expected: true,
		},
		{
			Input:    97,
			Expected: false,
		},
		{
			Input:    -97,
			Expected: false,
		},
	}

	for i, testStruct := range testStructs {
		got := IsEven(testStruct.Input)
		if got != testStruct.Expected {
			t.Errorf("Expected %v, got %v on iteration %d", testStruct.Expected, got, i)
		}
	}
}

func TestIsOdd(t *testing.T) {
	t.Parallel()

	testStructs := []struct {
		Input    int
		Expected bool
	}{
		{
			Input:    0,
			Expected: false,
		},
		{
			Input:    1,
			Expected: true,
		},
		{
			Input:    -1,
			Expected: true,
		},
		{
			Input:    2,
			Expected: false,
		},
		{
			Input:    -2,
			Expected: false,
		},
		{
			Input:    -98,
			Expected: false,
		},
		{
			Input:    98,
			Expected: false,
		},
		{
			Input:    97,
			Expected: true,
		},
		{
			Input:    -97,
			Expected: true,
		},
	}

	for i, testStruct := range testStructs {
		got := IsOdd(testStruct.Input)
		if got != testStruct.Expected {
			t.Errorf("Expected %v, got %v on iteration %d", testStruct.Expected, got, i)
		}
	}
}

func TestAbs(t *testing.T) {
	t.Parallel()

	testStructs := []struct {
		Input    int
		Expected int
	}{
		{
			Input:    0,
			Expected: 0,
		},
		{
			Input:    1,
			Expected: 1,
		},
		{
			Input:    -1,
			Expected: 1,
		},
		{
			Input:    2,
			Expected: 2,
		},
		{
			Input:    -2,
			Expected: 2,
		},
		{
			Input:    -98,
			Expected: 98,
		},
		{
			Input:    98,
			Expected: 98,
		},
		{
			Input:    97,
			Expected: 97,
		},
		{
			Input:    -97,
			Expected: 97,
		},
		{
			Input:    math.MaxInt,
			Expected: math.MaxInt,
		},
		{
			Input:    math.MinInt,
			Expected: math.MinInt,
		},
	}

	for i, testStruct := range testStructs {
		got := Abs(testStruct.Input)
		if got != testStruct.Expected {
			t.Errorf("Expected %d, got %d on iteration %d", testStruct.Expected, got, i)
		}
	}
}

func TestOddNumbers(t *testing.T) {
	t.Parallel()

	testStructs := []struct {
		A        int
		B        int
		Expected []int
	}{
		{
			A:        11,
			B:        10,
			Expected: []int{},
		},
		{
			A:        0,
			B:        0,
			Expected: []int{},
		},
		{
			A:        1,
			B:        1,
			Expected: []int{1},
		},
		{
			A:        2,
			B:        2,
			Expected: []int{},
		},
		{
			A:        3,
			B:        3,
			Expected: []int{3},
		},
		{
			A:        -1,
			B:        -1,
			Expected: []int{-1},
		},
		{
			A:        -2,
			B:        -2,
			Expected: []int{},
		},
		{
			A:        -1,
			B:        1,
			Expected: []int{-1, 1},
		},
		{
			A:        1,
			B:        10,
			Expected: []int{1, 3, 5, 7, 9},
		},
		{
			A:        -12,
			B:        11,
			Expected: []int{-11, -9, -7, -5, -3, -1, 1, 3, 5, 7, 9, 11},
		},
	}

	for i, testStruct := range testStructs {
		got := OddNumbers(testStruct.A, testStruct.B)

		if !reflect.DeepEqual(got, testStruct.Expected) {
			t.Errorf("Expected %v, got %v for %d and %d on interation %d", testStruct.Expected, got, testStruct.A, testStruct.B, i)
		}
	}
}

func TestEvenNumbers(t *testing.T) {
	t.Parallel()

	testStructs := []struct {
		A        int
		B        int
		Expected []int
	}{
		{
			A:        11,
			B:        10,
			Expected: []int{},
		},
		{
			A:        0,
			B:        0,
			Expected: []int{0},
		},
		{
			A:        1,
			B:        1,
			Expected: []int{},
		},
		{
			A:        2,
			B:        2,
			Expected: []int{2},
		},
		{
			A:        3,
			B:        3,
			Expected: []int{},
		},
		{
			A:        -1,
			B:        -1,
			Expected: []int{},
		},
		{
			A:        -1,
			B:        1,
			Expected: []int{0},
		},
		{
			A:        -2,
			B:        -2,
			Expected: []int{-2},
		},
		{
			A:        1,
			B:        10,
			Expected: []int{2, 4, 6, 8, 10},
		},
		{
			A:        -12,
			B:        11,
			Expected: []int{-12, -10, -8, -6, -4, -2, 0, 2, 4, 6, 8, 10},
		},
	}

	for i, testStruct := range testStructs {
		got := EvenNumbers(testStruct.A, testStruct.B)

		if !reflect.DeepEqual(got, testStruct.Expected) {
			t.Errorf("Expected %v, got %v for %d and %d on interation %d", testStruct.Expected, got, testStruct.A, testStruct.B, i)
		}
	}
}
