package bat

import (
	"fmt"
	"reflect"
	"strings"
)

// StrSliceIdx returns first index of `x` in `slice` and -1 if `x` is not present.
func StrSliceIdx(slice []string, x string) int {
	for i, v := range slice {
		if v == x {
			return i
		}
	}
	return -1
}

// IntSliceIdx returns first index of `x` in `slice` and -1 if `x` is not present.
func IntSliceIdx(slice []int, x int) int {
	for i, v := range slice {
		if v == x {
			return i
		}
	}
	return -1
}

// StrSliceSliceIdx returns first index of `x` in `slices` and -1 if `x` is not present.
func StrSliceSliceIdx(slices [][]string, x []string) int {
	for k, s := range slices {
		if len(s) != len(x) {
			goto out
		}
		for i := range s {
			if s[i] != x[i] {
				goto out
			}
		}
		return k
	out:
	}
	return -1
}

// SliceIdx returns first index of `x` in `slice` and -1 if `x` is not present.
func SliceIdx(slice []interface{}, x interface{}) int {
	for i, v := range slice {
		if v == x {
			return i
		}
	}
	return -1
}

// StrSliceConcat concats all given string slices into a new slice
func StrSliceConcat(slices ...[]string) []string {
	var r []string
	for _, s := range slices {
		r = append(r, s...)
	}
	return r
}

// SliceConcat concats all given slices into a new slice
func SliceConcat(slices ...[]interface{}) []interface{} {
	var r []interface{}
	for _, s := range slices {
		r = append(r, s...)
	}
	return r
}

// StrsEq checks if two string slices has the same content
func StrsEq(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// StrsOnlyWhitespace checks if input consists of only whitespace
func StrsOnlyWhitespace(s []string) bool {
	for _, elem := range s {
		if strings.TrimSpace(elem) != "" {
			return false
		}
	}
	return true
}

// StrsSimilar behaves like StringsEq but order doesn't matter
func StrsSimilar(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	c := map[string]struct{}{}
	for _, v := range a {
		c[v] = struct{}{}
	}
	for _, v := range b {
		if _, ok := c[v]; !ok {
			return false
		}
	}
	return true
}

// ToInterfaceSlice converts given slice into empty interface slice.
// Mostly for batch DB operations
func ToInterfaceSlice(i interface{}) ([]interface{}, error) {
	switch reflect.TypeOf(i).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(i)
		result := make([]interface{}, s.Len())
		for j := 0; j < s.Len(); j++ {
			result[j] = s.Index(j).Interface()
		}
		return result, nil
	default:
		return nil, fmt.Errorf("Cannot convert %T to []interface{}", i)
	}
}

// StrSliceUniqueAppend appends given strings into source if not exist in source
func StrSliceUniqueAppend(src []string, strs ...string) []string {
	for _, s := range strs {
		if StrSliceIdx(src, s) < 0 {
			src = append(src, s)
		}
	}
	return src
}

// EmptyStrSliceAsNil returns initialized empty slice as nil
func EmptyStrSliceAsNil(ls []string) []string {
	if len(ls) == 0 && ls != nil {
		return nil
	}
	return ls
}
