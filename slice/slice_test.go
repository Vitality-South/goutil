package slice

import (
	"reflect"
	"testing"
)

func TestPointerSlice(t *testing.T) {
	test := func(s []string) {
		sl := PointerSlice(s)

		// length of new slice and original slice should be the same
		if len(sl) != len(s) {
			t.Errorf("len(sl) != len(s); expected %d, got %d", len(s), len(sl))
		}

		// both values and addresses of elements in both slices should be the same
		for i, v := range sl {
			if *v != s[i] {
				t.Errorf("expected %s, got %s", s[i], *v)
			}

			if v != &s[i] {
				t.Errorf("expected %v, got %v", &s[i], v)
			}
		}

		// changing values in original slice should be reflected in the slice of pointers
		for so := range s {
			s[so] = "changed"
		}

		for _, ss := range sl {
			if *ss != "changed" {
				t.Errorf("expected changed, got %s", *ss)
			}
		}
	}

	// test a string slice
	s := []string{"bob", "is", "awesome"}
	test(s)

	// test an empty slice
	var empty []string

	test(empty)

	// test nil
	test(nil)
}

func TestContains(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}

	if !Contains(nums, 3) {
		t.Error("nums contains 3")
	}

	if Contains(nums, 6) {
		t.Error("nums does not contain 3")
	}
}

func TestContainsFunc(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}

	if !ContainsFunc(nums, func(n int) bool { return n == 3 }) {
		t.Error("nums contains 3")
	}

	if ContainsFunc(nums, func(n int) bool { return n == 6 }) {
		t.Error("nums does not contain 6")
	}
}

func TestRemoveElement(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	expected := []int{1, 2, 4, 5}

	removed := RemoveElement(nums, 3)

	if len(removed) != 4 {
		t.Errorf("expected len 4 for removed; got %d", len(removed))
	}

	if !reflect.DeepEqual(removed, expected) {
		t.Error("removed and expected should match")
	}

	notremoved := RemoveElement(nums, 6)

	if len(notremoved) != 5 {
		t.Errorf("expected len 5 for notremoved; got %d", len(notremoved))
	}

	if !reflect.DeepEqual(nums, notremoved) {
		t.Error("nums and notremoved should match")
	}
}

func TestRemoveElementFunc(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	expected := []int{1, 2, 4, 5}

	removed := RemoveElementFunc(nums, func(n int) bool { return n == 3 })

	if len(removed) != 4 {
		t.Errorf("expected len 4 for removed; got %d, [%v]", len(removed), removed)
	}

	if !reflect.DeepEqual(removed, expected) {
		t.Error("removed and expected should match")
	}

	notremoved := RemoveElementFunc(nums, func(n int) bool { return n == 6 })

	if len(notremoved) != 5 {
		t.Errorf("expected len 5 for removed; got %d", len(notremoved))
	}

	if !reflect.DeepEqual(nums, notremoved) {
		t.Error("nums and notremoved should match")
	}
}
