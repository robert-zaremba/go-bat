package bat

import (
	"fmt"
	"reflect"
)

// SliceStrIdx returns first index of `x` in `slice` and -1 if `x` is not present.
func SliceStrIdx(slice []string, x string) int {
	for i, v := range slice {
		if v == x {
			return i
		}
	}
	return -1
}

// SliceIntIdx returns first index of `x` in `slice` and -1 if `x` is not present.
func SliceIntIdx(slice []int, x int) int {
	for i, v := range slice {
		if v == x {
			return i
		}
	}
	return -1
}

// SliceStrSliceIdx returns first index of `x` in `slices` and -1 if `x` is not present.
func SliceStrSliceIdx(slices [][]string, x []string) int {
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

// SliceStrConcat concats all given string slices into a new slice
func SliceStrConcat(slices ...[]string) []string {
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
func StrsOnlyWhitespace(ss []string) bool {
	for _, s := range ss {
		if !StrIsWhitespace(s) {
			return false
		}
	}
	return true
}

// StrsSame behaves like StringsEq but order doesn't matter
func StrsSame(a, b []string) bool {
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

// SliceStrUniqueAppend appends `strs` strings into source without excluding elements which
// alread exist in `src`.
func SliceStrUniqueAppend(src []string, strs ...string) []string {
	for _, s := range strs {
		if SliceStrIdx(src, s) < 0 {
			src = append(src, s)
		}
	}
	return src
}

// SliceStrAsNil returns initialized empty slice as nil
func SliceStrAsNil(ls []string) []string {
	if len(ls) == 0 && ls != nil {
		return nil
	}
	return ls
}
