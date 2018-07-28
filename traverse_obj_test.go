package bat

import (
	"testing"
)

func mustReceive(t *testing.T, underTest interface{}, expected ...interface{}) {
	numberOfCals := 0
	failed := false
	listener := func(i interface{}) error {
		if failed {
			return nil
		}
		if numberOfCals >= len(expected) {
			t.Errorf("Received too many values. Last received: %v", i)
			t.Fail()
			failed = true
			return nil
		}
		if i != expected[numberOfCals] {
			t.Errorf("Received %v expected %v", i, expected[numberOfCals])
			t.Fail()
			failed = true
			return nil
		}
		numberOfCals++
		return nil
	}
	err := TraverseObj(underTest, listener, true)
	if failed {
		return
	}
	if err != nil {
		t.Error(err)
		t.Fail()
	} else if numberOfCals != len(expected) {
		t.Errorf("Missing elements: %v", expected[numberOfCals:])
		t.Fail()
	}
}

func TestTraverseObjMap(t *testing.T) {
	m := map[int]interface{}{
		1: "1",
		2: "2",
		3: "3",
		4: nil,
	}
	mustReceive(t, m, 1, "1", 2, "2", 3, "3", 4, nil)
}

type compositeKey struct {
	a, b string
}

func TestTraverseObjMapWithComplexKey(t *testing.T) {
	m := map[compositeKey]interface{}{
		{"a", "b"}: "c",
		{"d", "e"}: "f",
		{"g", "h"}: "i",
	}
	mustReceive(t, m, "a", "b", "c", "d", "e", "f", "g", "h", "i")
}

func TestTraverseObjStruct(t *testing.T) {
	s := struct {
		A string
		B int
		C bool
		d int
	}{
		C: true,
		A: "a",
		B: 1,
		d: 5,
	}
	mustReceive(t, s, "a", 1, true, 5)
}

func TestTraverseObjPrimitives(t *testing.T) {
	s := []interface{}{
		int(1),
		int64(1),
		int32(1),
		int16(1),
		int8(1),
		uint(1),
		uint64(1),
		uint32(1),
		uint16(1),
		uint8(1),
		uintptr(1),
		float64(1),
		float32(1),
		complex64(complex(1, 1)),
		complex128(complex(1, 1)),
		"1",
		true,
	}

	mustReceive(t, s, s...)
}

func TestTraverseObjStructWithNils(t *testing.T) {
	s := struct {
		A string
		B int
		C bool
		D *string
		E map[string]string
		F []byte
	}{
		C: true,
		A: "a",
		B: 1,
		F: nil,
		E: nil,
	}

	mustReceive(t, s, "a", 1, true, nil)
}

func TestTraverseObjStructWithChanel(t *testing.T) {
	c := make(chan string)
	s := struct {
		A chan string
	}{
		A: c,
	}

	mustReceive(t, s, c)
}
