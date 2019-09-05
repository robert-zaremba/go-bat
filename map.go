package bat

import "sort"

// MapKeysStrBool returns a list of keys from a map[string]bool
func MapKeysStrBool(m map[string]bool) []string {
	keys := make([]string, len(m))
	i := 0
	for s := range m {
		keys[i] = s
		i++
	}
	return keys
}

// MapSortedKeysStrBool returns a list of sorted keys from a map[string]bool
func MapSortedKeysStrBool(m map[string]bool) []string {
	ls := MapKeysStrBool(m)
	sort.Strings(ls)
	return ls
}

// MapKeysStrStruct returns a sorted list of keys from a map[string]struct{}
func MapKeysStrStruct(m map[string]struct{}) []string {
	keys := make([]string, len(m))
	i := 0
	for s := range m {
		keys[i] = s
		i++
	}
	sort.Strings(keys)
	return keys
}

// MapKeysStrInterface returns a sorted list of keys from a map[string]nterface{}
func MapKeysStrInterface(m map[string]interface{}) []string {
	keys := make([]string, len(m))
	i := 0
	for s := range m {
		keys[i] = s
		i++
	}
	sort.Strings(keys)
	return keys
}

// MapKeysStrStr returns a sorted list of keys from map[string]string
func MapKeysStrStr(m map[string]string) []string {
	keys := make([]string, len(m))
	i := 0
	for s := range m {
		keys[i] = s
		i++
	}
	sort.Strings(keys)
	return keys
}

// MapKeysInt64Bool returns a sorted list of keys from a map[int64]bool
func MapKeysInt64Bool(m map[int64]bool) []int64 {
	keys := make([]int64, len(m))
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

// MapUpdateStrStr loads destination map into source map.
// If dest is empty it will create a new map.
func MapUpdateStrStr(source map[string]string, dest *map[string]string) {
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

// MapUpdateStrInterface  loads destination map into source map.
func MapUpdateStrInterface(source map[string]interface{}, dest *map[string]interface{}) {
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

// MapCloneStrInterface will clone a given map[str]interface{} to a new one
func MapCloneStrInterface(source map[string]interface{}) map[string]interface{} {
	d := make(map[string]interface{}, len(source))
	for k, v := range source {
		d[k] = v
	}
	return d
}

// MapCloneStrStr will clone a given map[string]string{} to a new one
func MapCloneStrStr(source map[string]string) map[string]string {
	d := make(map[string]string, len(source))
	for k, v := range source {
		d[k] = v
	}
	return d
}

// SliceStrToMap creates a map from string slice for frequent search in slice
func SliceStrToMap(ls []string) map[string]struct{} {
	m := make(map[string]struct{}, len(ls))
	for _, v := range ls {
		m[v] = struct{}{}
	}
	return m
}

// MapReverseStrStr reverses keys and values of given map
// If same value exists multiple times, the one with bigger key wins
// to give deterministic ordering and to make it testable.
// {"a": "1", "b": 1} -> {"1": "b"}
func MapReverseStrStr(m map[string]string) map[string]string {
	r := make(map[string]string, len(m))
	for k, v := range m {
		if val, ok := r[v]; !ok || k > val {
			r[v] = k
		}
	}
	return r
}

// MapStrStrGetAnyValue returns any values from the map or empty string if map is empty
func MapStrStrGetAnyValue(m map[string]string) string {
	for _, v := range m {
		return v
	}
	return ""
}
