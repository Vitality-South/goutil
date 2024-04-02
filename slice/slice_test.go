package slice

import (
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
