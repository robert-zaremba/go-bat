package bat

import (
	"reflect"
	"testing"
)

func TestPages(t *testing.T) {
	tests := []struct {
		step     int
		length   int
		expected []Page
	}{
		{4, 10, []Page{{0, 4}, {4, 8}, {8, 10}}},
		{1, 10, []Page{{0, 1}, {1, 2}, {2, 3}, {3, 4}, {4, 5}, {5, 6}, {6, 7}, {7, 8}, {8, 9}, {9, 10}}},
		{10, 10, []Page{{0, 10}}},
		{20, 10, []Page{{0, 10}}},
		{2, 0, []Page{}},
	}

	for _, tt := range tests {
		a := Pages(tt.step, tt.length)
		if !reflect.DeepEqual(a, tt.expected) {
			t.Errorf("Pages(%d,%d) => %v, want %v", tt.step, tt.length, a, tt.expected)
		}
	}
}

func TestPagesNegativeStep(t *testing.T) {

	shouldPanic := func(f func(), testCase string) {
		defer func() {
			if recover() == nil {
				t.Error(testCase, "expected panic")
			}
		}()
		f()
	}

	shouldPanic(func() { Pages(-1, 4) }, "Negative step")
	shouldPanic(func() { Pages(0, 4) }, "Zero step")
	shouldPanic(func() { Pages(4, -1) }, "Negative length")
}
