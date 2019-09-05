package bat

import (
	"reflect"
	"sort"
	"testing"
)

func TestStrBoolMapKeys(t *testing.T) {
	tests := []map[string]bool{
		nil,
		{},
		{"bool": false},
		{"a": true, "b": false},
	}
	expects := [][]string{
		{},
		{},
		{"bool"},
		{"a", "b"},
	}

	for i, test := range tests {
		result := MapKeysStrBool(test)
		sort.Strings(result)
		if !reflect.DeepEqual(result, expects[i]) {
			t.Errorf("%d: Expected %v, got %v", i, expects[i], result)
		}
	}
}

func TestStrBoolMapSortedKeys(t *testing.T) {
	tests := []map[string]bool{
		nil,
		{},
		{"bool": false},
		{"a": true, "b": false},
		{"1": false, "10": false, "2": true, "a": true},
	}
	expects := [][]string{
		{},
		{},
		{"bool"},
		{"a", "b"},
		{"1", "10", "2", "a"},
	}

	for i, test := range tests {
		result := MapSortedKeysStrBool(test)
		if !reflect.DeepEqual(result, expects[i]) {
			t.Errorf("%d: Expected %v, got %v", i, expects[i], result)
		}
	}
}

func TestInt64BoolMapKeys(t *testing.T) {
	tests := []map[int64]bool{
		nil,
		{},
		{-1: false},
		{0: true, 15129: false},
	}
	expects := [][]int64{
		{},
		{},
		{-1},
		{0, 15129},
	}

	for i, test := range tests {
		result := MapKeysInt64Bool(test)
		sort.Slice(result, func(i, j int) bool { return result[i] < result[j] })
		if !reflect.DeepEqual(result, expects[i]) {
			t.Errorf("%d: Expected %v, got %v", i, expects[i], result)
		}
	}
}

func TestUpdateStrStrMap(t *testing.T) {
	var dest map[string]string
	MapUpdateStrStr(nil, &dest)
	if dest != nil {
		t.Error("dest should be nil since source is nil", dest)
	}
	MapUpdateStrStr(map[string]string{"a": "A"}, &dest)
	if _, ok := dest["a"]; !ok {
		t.Fatal("'a' should be in dest", dest)
	}
}

func TestStrInterfaceMap(t *testing.T) {
	var dest map[string]interface{}
	MapUpdateStrInterface(nil, &dest)
	if dest != nil {
		t.Error("dest should be nil since source is nil", dest)
	}
	src := map[string]interface{}{
		"bytes":    []byte("agflow"),
		"intSlice": []int{1, 13},
		"nil":      nil,
	}
	MapUpdateStrInterface(src, &dest)
	if !reflect.DeepEqual(dest["bytes"], []byte("agflow")) ||
		!reflect.DeepEqual(dest["intSlice"], []int{1, 13}) ||
		dest["nil"] != nil {
		t.Error("dest doesn't contain expected values", dest)
	}
}

func TestCloneMapStrInterface(t *testing.T) {
	t.Skip("What if map contains reference types? i.e. Should pointers be followed?")
	var m1 map[string]interface{}
	m2 := MapCloneStrInterface(m1)
	if m2 == nil {
		t.Fatal("m2 must be empty map", m1, m2)
	}

	m1 = map[string]interface{}{
		"t1": struct{ test bool }{false},
		"t2": []int{3, 2, 1},
	}
	m2 = MapCloneStrInterface(m1)
	if !reflect.DeepEqual(m1, m2) {
		t.Error("m2 is expected to be same with m1", m1, m2)
	}
	m2["t2"].([]int)[1] = 4
	if m1["t2"].([]int)[1] != 2 {
		t.Error("m2 updates shouldn't change m1", m1, m2)
	}
}

func TestStrSliceToMap(t *testing.T) {
	tests := [][]string{
		nil,
		{},
		{"a"},
		{"b", "a"},
	}
	emptyS := struct{}{}
	expects := []map[string]struct{}{
		{},
		{},
		{"a": emptyS},
		{"a": emptyS, "b": emptyS},
	}
	for i, test := range tests {
		result := SliceStrToMap(test)
		if !reflect.DeepEqual(result, expects[i]) {
			t.Errorf("%d: Expected %v, got %v", i, expects[i], result)
		}
	}
}

func TestReverseStrStrMap(t *testing.T) {
	tests := []map[string]string{
		nil,
		{},
		{"a": "1", "b": "2", "c": "3"},
		{"a": "1", "b": "1"},
	}
	expects := []map[string]string{
		{},
		{},
		{"1": "a", "2": "b", "3": "c"},
		{"1": "b"},
	}

	for i, test := range tests {
		result := MapReverseStrStr(test)
		if !reflect.DeepEqual(result, expects[i]) {
			t.Errorf("%d: Expected %v, got %v", i, expects[i], result)
		}
	}
}
