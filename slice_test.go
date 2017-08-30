package bat

import (
	"reflect"
	"testing"
)

func TestStrSliceIdx(t *testing.T) {
	testStringSlices := [][]string{
		nil,
		{},
		{"a", "b"},
		{"a", "b"},
		{"a", "b", "c", "d", "e"},
	}
	testStrs := []string{"", "a", "a", "c", "e"}
	expects := []int{-1, -1, 0, -1, 4}
	for i, str := range testStringSlices {
		j := StrSliceIdx(str, testStrs[i])
		if expects[i] != j {
			t.Errorf("%d: Expected %d, got %d", i, expects[i], j)
		}
	}
}

func TestIntSliceIdx(t *testing.T) {
	testIntSlices := [][]int{
		nil,
		{},
		{0, 1},
		{1, 2, 3},
		{1, 2, 3, 4, 0},
	}
	testInts := []int{0, 0, 0, 0, 0}
	expects := []int{-1, -1, 0, -1, 4}
	for i, is := range testIntSlices {
		j := IntSliceIdx(is, testInts[i])
		if expects[i] != j {
			t.Errorf("%d: Expected %d, got %d", i, expects[i], j)
		}
	}
}

func TestStrSliceSliceIdx(t *testing.T) {
	testStringSlice := [][][]string{
		nil,
		{},
		{nil},
		{{}},
		{nil, {}},
		{nil, {}, {"a", "b", "c"}},
	}
	testStrs := [][]string{
		nil,
		nil,
		nil,
		{},
		{"a", "b", "c"},
		{"a", "b", "c"},
	}
	expects := []int{-1, -1, 0, 0, -1, 2}

	for i, s := range testStringSlice {
		j := StrSliceSliceIdx(s, testStrs[i])
		if expects[i] != j {
			t.Errorf("%d: Expected %d, got %d", i, expects[i], j)
		}
	}
}

func TestSliceIdx(t *testing.T) {
	t.Skip("What if we want to search for uncomparable types")
	testInterfaces := [][]interface{}{
		{nil, true, -1, struct{}{}, [][][]string{}},
		{nil, true, -1, struct{}{}, [][][]string{}},
		{nil, true, -1, struct{}{}, [][][]string{}},
		{nil, true, -1, struct{}{}, [][][]string{}},
		{nil, true, -1, struct{}{}, [][][]string{}},
		{nil, true, -1, struct{}{}, [][][]string{}},
	}
	tests := []interface{}{
		nil,
		true,
		-1,
		struct{}{},
		[][][]string{},
		0,
	}
	expects := []int{0, 1, 2, 3, 4, -1}

	for i, in := range testInterfaces {
		j := SliceIdx(in, tests[i])
		if expects[i] != j {
			t.Errorf("%d: Expected %d, got %d", i, expects[i], j)
		}
	}
}

func TestStrSliceConcat(t *testing.T) {
	tests := [][][]string{
		{},
		{nil, {}},
		{nil, {}, {"a"}},
		{nil, {}, {"a"}, {"a"}},
	}
	expects := [][]string{
		nil,
		nil,
		{"a"},
		{"a", "a"},
	}

	for i, s := range tests {
		result := StrSliceConcat(s...)
		if !reflect.DeepEqual(result, expects[i]) {
			t.Errorf("%d: Expected %v, got %v", i, expects[i], result)
		}
	}
}

func TestStrsEq(t *testing.T) {
	tests := [][]string{
		nil,
		{},
		{"a"},
		{"", "a", "1"},
	}
	for i, s := range tests {
		if !StrsEq(s, s) {
			t.Errorf("%d: Equality failed for %v", i, s)
		}
	}
}

func TestStrsSimilar(t *testing.T) {
	tests := [][]string{
		nil,
		{},
		{"a", "c", "b"},
		{"1", "2", "3", "4"},
	}
	expects := [][]string{
		nil,
		{},
		{"b", "c", "a"},
		{"4", "3", "2", "1"},
	}
	for i, s := range tests {
		if !StrsSimilar(s, expects[i]) {
			t.Errorf("%d: %v is expected to be similar to %v", i, s, expects[i])
		}
	}
}

func TestStrSliceUniqueAppend(t *testing.T) {
	sources := [][]string{
		nil,
		nil,
		{},
		{"a"},
	}
	adds := [][]string{
		{"a"},
		{"a", "a"},
		{"a"},
		{"a", "a", "a", "a"},
	}
	expects := [][]string{
		{"a"},
		{"a"},
		{"a"},
		{"a"},
	}
	for i, s := range sources {
		result := StrSliceUniqueAppend(s, adds[i]...)
		if !reflect.DeepEqual(result, expects[i]) {
			t.Errorf("%d: Expected %v, got %v", i, expects[i], result)
		}
	}
}

func TestEmptyStrSliceAsNil(t *testing.T) {
	if EmptyStrSliceAsNil(nil) != nil {
		t.Errorf("nil shouldn't be changed")
	}
	if EmptyStrSliceAsNil([]string{}) != nil {
		t.Errorf("empty slice should be nil")
	}
	if EmptyStrSliceAsNil([]string{"a"}) == nil {
		t.Errorf("Not empty slice shouldn't be changed")
	}
}
