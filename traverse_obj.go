package bat

import (
	"fmt"
	"reflect"
	"sort"
)

// TraverseObj traverses `v` and sends all encountered values to the listener.
// Traversal is stable, so values will always come in the same order.
// Only simple values (like `int` or `string`) will be send to listener.
// Error is returned only if cycle is detected in the traversed object and
// `errorOnCyclic` is true. Otherwise cyclic references are ignored.
func TraverseObj(v interface{}, listener func(interface{}) error, errorOnCyclic bool) error {
	return traverse(reflect.ValueOf(v), listener, map[visit]bool{}, errorOnCyclic)
}

// visited obj
type visit struct {
	a   uintptr
	typ reflect.Type
}

// TODO: this function is too complex, should be splitted.
func traverse(v reflect.Value, listener func(interface{}) error, visited map[visit]bool, errorOnCyclic bool) error {
	if !v.IsValid() {
		return listener(nil)
	}
	if p, ok := castToPrimitive(v); ok {
		return listener(p)
	}
	if v.CanAddr() && isComplexType(v.Kind()) {
		addr := v.UnsafeAddr()
		typ := v.Type()
		v := visit{addr, typ}
		if visited[v] {
			if errorOnCyclic {
				return fmt.Errorf("Cyclic reference on type %s", typ)
			}
			return nil
		}

		// Remember for later.
		visited[v] = true
	}

	switch v.Kind() {
	case reflect.Array:
		for i := 0; i < v.Len(); i++ {
			if err := traverse(v.Index(i), listener, visited, errorOnCyclic); err != nil {
				return err
			}
		}
		return nil
	case reflect.Slice:
		for i := 0; i < v.Len(); i++ {
			if err := traverse(v.Index(i), listener, visited, errorOnCyclic); err != nil {
				return err
			}
		}
		return nil
	case reflect.Interface:
		return traverse(v.Elem(), listener, visited, errorOnCyclic)
	case reflect.Ptr:
		if v.IsNil() {
			return listener(nil)
		}
		return traverse(v.Elem(), listener, visited, errorOnCyclic)
	case reflect.Struct:
		for i, n := 0, v.NumField(); i < n; i++ {
			if err := traverse(v.Field(i), listener, visited, errorOnCyclic); err != nil {
				return err
			}
		}
		return nil
	case reflect.Map:
		keys := valueSorter{
			values: v.MapKeys(),
		}
		sort.Sort(&keys)
		if keys.err != nil {
			return keys.err
		}
		for _, k := range keys.values {
			if err := traverse(k, listener, visited, errorOnCyclic); err != nil {
				return err
			}
			if err := traverse(v.MapIndex(k), listener, visited, errorOnCyclic); err != nil {
				return err
			}
		}
		return nil
	case reflect.Func:
		if v.CanInterface() {
			return listener(v.Interface())
		}
		return nil
	default:
		if !v.CanInterface() {
			//v is an unexported struct field
			return nil
		}
		return listener(v.Interface())
	}
}

func isComplexType(k reflect.Kind) bool {
	switch k {
	case reflect.Array, reflect.Map, reflect.Slice, reflect.Struct:
		return true
	}
	return false
}

func castToPrimitive(v reflect.Value) (interface{}, bool) {
	switch v.Kind() {
	case reflect.Int:
		return int(v.Int()), true
	case reflect.Int64:
		return v.Int(), true
	case reflect.Int32:
		return int32(v.Int()), true
	case reflect.Int16:
		return int16(v.Int()), true
	case reflect.Int8:
		return int8(v.Int()), true
	case reflect.Uint:
		return uint(v.Uint()), true
	case reflect.Uint64:
		return v.Uint(), true
	case reflect.Uint32:
		return uint32(v.Uint()), true
	case reflect.Uint16:
		return uint16(v.Uint()), true
	case reflect.Uint8:
		return uint8(v.Uint()), true
	case reflect.Uintptr:
		return uintptr(v.Uint()), true
	case reflect.Float64:
		return v.Float(), true
	case reflect.Float32:
		return float32(v.Float()), true
	case reflect.Complex64:
		return complex64(v.Complex()), true
	case reflect.Complex128:
		return v.Complex(), true
	case reflect.Bool:
		return v.Bool(), true
	case reflect.String:
		return v.String(), true
	default:
		return nil, false
	}
}

type valueSorter struct {
	values []reflect.Value
	err    error
}

func (s *valueSorter) toString(i int) string {
	return fmt.Sprint(s.values[i])
}

// Len is part of sort.Interface.
func (s *valueSorter) Len() int {
	return len(s.values)
}

// Swap is part of sort.Interface.
func (s *valueSorter) Swap(i, j int) {
	s.values[i], s.values[j] = s.values[j], s.values[i]
}

// Less is part of sort.Interface
func (s *valueSorter) Less(i, j int) bool {
	iStr := s.toString(i)
	jStr := s.toString(j)
	return iStr <= jStr
}
