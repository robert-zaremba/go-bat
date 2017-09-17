package bat

import "sort"

// StrBoolMapKeys returns a list of keys from a map[string]bool
func StrBoolMapKeys(m map[string]bool) []string {
	keys := make([]string, len(m))
	i := 0
	for s := range m {
		keys[i] = s
		i++
	}
	return keys
}

// StrBoolMapSortedKeys returns a list of sorted keys from a map[string]bool
func StrBoolMapSortedKeys(m map[string]bool) []string {
	ls := StrBoolMapKeys(m)
	sort.Strings(ls)
	return ls
}

// StrStructMapKeys returns a sorted list of keys from a map[string]struct{}
func StrStructMapKeys(m map[string]struct{}) []string {
	keys := make([]string, len(m))
	i := 0
	for s := range m {
		keys[i] = s
		i++
	}
	sort.Strings(keys)
	return keys
}

// StrInterfaceMapKeys returns a sorted list of keys from a map[string]nterface{}
func StrInterfaceMapKeys(m map[string]interface{}) []string {
	keys := make([]string, len(m))
	i := 0
	for s := range m {
		keys[i] = s
		i++
	}
	sort.Strings(keys)
	return keys
}

// StrStrMapKeys returns a sorted list of keys from map[string]string
func StrStrMapKeys(m map[string]string) []string {
	keys := make([]string, len(m))
	i := 0
	for s := range m {
		keys[i] = s
		i++
	}
	sort.Strings(keys)
	return keys
}

// Int64BoolMapKeys returns a sorted list of keys from a map[int64]bool
func Int64BoolMapKeys(m map[int64]bool) []int64 {
	keys := make([]int64, len(m), len(m))
	i := 0
	for s := range m {
		keys[i] = s
		i++
	}
	less := func(i, j int) bool {
		return keys[i] <= keys[j]
	}
	sort.Slice(keys, less)
	return keys
}

// UpdateStrStrMap loads destination map into source map.
// If dest is empty it will create a new map.
func UpdateStrStrMap(source map[string]string, dest *map[string]string) {
	if len(source) == 0 {
		return
	}
	d := *dest
	if d == nil {
		d = make(map[string]string, len(source))
	}
	for k, v := range source {
		d[k] = v
	}
	*dest = d
}

// UpdateStrInterfaceMap  loads destination map into source map.
func UpdateStrInterfaceMap(source map[string]interface{}, dest *map[string]interface{}) {
	if len(source) == 0 {
		return
	}
	d := *dest
	if d == nil {
		d = make(map[string]interface{}, len(source))
	}
	for k, v := range source {
		d[k] = v
	}
	*dest = d
}

// CloneMapStrInterface will clone a given map[str]interface{} to a new one
func CloneMapStrInterface(source map[string]interface{}) map[string]interface{} {
	d := make(map[string]interface{}, len(source))
	for k, v := range source {
		d[k] = v
	}
	return d
}

// CloneMapStrStr will clone a given map[string]string{} to a new one
func CloneMapStrStr(source map[string]string) map[string]string {
	d := make(map[string]string, len(source))
	for k, v := range source {
		d[k] = v
	}
	return d
}

// StrSliceToMap creates a map from string slice for frequent search in slice
func StrSliceToMap(ls []string) map[string]bool {
	m := make(map[string]bool, len(ls))
	for _, v := range ls {
		m[v] = true
	}
	return m
}

// ReverseStrStrMap reverses keys and values of given map
// If same value exists multiple times, the one with bigger key wins
// to give deterministic ordering and to make it testable.
// {"a": "1", "b": 1} -> {"1": "b"}
func ReverseStrStrMap(m map[string]string) map[string]string {
	r := make(map[string]string, len(m))
	for k, v := range m {
		if val, ok := r[v]; !ok || k > val {
			r[v] = k
		}
	}
	return r
}

// GetAnyValueFromMapStrStr returns any values from the map or empty string if map is empty
func GetAnyValueFromMapStrStr(m map[string]string) string {
	for _, v := range m {
		return v
	}
	return ""
}
