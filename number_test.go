package bat

import "testing"

func TestI64toa(t *testing.T) {
	if I64toa(12) != "12" {
		t.Errorf("Expected 12, got %s", I64toa(12))
	}
}

func TestI64tox(t *testing.T) {
	if I64tox(12) != "c" {
		t.Errorf("Expected 12, got %s", I64tox(12))
	}
	if I64tox(31) != "1f" {
		t.Errorf("Expected 1F, got %s", I64tox(31))
	}
}

func TestF64toa(t *testing.T) {
	f := 12.56789
	e := "12.56789"
	if F64toa(f) != "12.56789" {
		t.Errorf("Expected 12.56, got %s", F64toa(f))
	}
	e = "12.57"
	if F64toa(f, 2) != "12.57" {
		t.Errorf("Expected %s, got %s", e, F64toa(f, 2))
	}
	e = "12.5678900000"
	if F64toa(f, 10) != e {
		t.Errorf("Expected %s, got %s", e, F64toa(f, 10))
	}
}

func TestAtoi64(t *testing.T) {
	var e int64 = 12
	i, err := Atoi64("12")
	if err != nil || i != e {
		t.Errorf("Expected %d, got %d", e, i)
	}
	i, err = Atoi64("12.0")
	if err == nil {
		t.Errorf("Expected %d, got %d", e, i)
	}
}

func TestAtof64(t *testing.T) {
	var e = 12.0
	i, err := Atof64("12")
	if err != nil || i != e {
		t.Errorf("Expected %f, got %f", e, i)
	}
	i, err = Atof64("12.0")
	if err != nil || i != e {
		t.Errorf("Expected %f, got %f", e, i)
	}
}

func TestHumanizeInt(t *testing.T) {
	tests := []struct {
		val      uint64
		expected string
	}{
		{0, "0 B"},
		{8, "8 B"},
		{20, "20 B"},
		{2000, "2.0 kB"},
		{8000762, "8.0 MB"},
		{9854271223, "9.9 GB"},
		{7200067891245, "7.2 TB"},
	}
	for _, tt := range tests {
		if s := HumanizeInt(tt.val); tt.expected != s {
			t.Errorf("Expected %s for %d but got %s", tt.expected, tt.val, s)
		}
	}
}
