package bat

import (
	"testing"
)

func TestStrJoin(t *testing.T) {
	sep, args := "-", [][]string{
		nil,
		{},
		{"a", "b", "c"},
	}
	expects := []string{"", "", "a-b-c"}

	for i, as := range args {
		res := StrJoin(sep, as...)
		if res != expects[i] {
			t.Errorf("%d: Expected %s, got %s", i, expects[i], res)
		}
	}
}

func TestStrTruncateUpTo(t *testing.T) {
	cut, str, ns := 'a', "aaabbcdefghaaaa", []int{-1, 0, 1, 2, 3, 4, 5}
	expects := []string{"", "", "a", "aa", "aaabbcdefgh", "aaabbcdefgha", "aaabbcdefghaa"}

	for i, c := range ns {
		res := StrTruncateUpTo(str, cut, c)
		if res != expects[i] {
			t.Errorf("%d: Expected %s, got %s", i, expects[i], res)
		}
	}
}

func TestStrIsNum(t *testing.T) {
	tests := []string{"", "1", "15129", "729", "1f85"}
	expects := []bool{true, true, true, true, false}

	for i, tt := range tests {
		res := StrIsNum(tt)
		if res != expects[i] {
			t.Errorf("%d %s: Expected %v, got %v", i, tt, expects[i], res)
		}
	}
}

func TestStrTrimMultiSpace(t *testing.T) {
	tests := []string{"   a   b   c d    ef  ", "     123   "}
	expects := []string{"a b c d ef", "123"}

	for i, tt := range tests {
		res := StrTrimMultiSpace(tt)
		if res != expects[i] {
			t.Errorf("%d '%s': Expected %s, got %s", i, tt, expects[i], res)
		}
	}
}

func TestRmDiacritics(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"ąćęłńóśźż ĄĆĘŁŃÓŚŹŻ", "acelnoszz ACELNOSZZ"},                         // Polish
		{"čďěňřšťž ČĎĚŇŘŠŤŽ", "cdenrstz CDENRSTZ"},                             // Czech
		{"éàèùâêîôûëïöüÿç ÉÀÈÙÂÊÎÔÛËÏÖÜŸÇ", "eaeuaeioueiouyc EAEUAEIOUEIOUYC"}, // French
		{"ğöüçş ĞÖÜÇŞ", "goucs GOUCS"},                                         // Turkish
	}

	for _, tt := range tests {
		if res := RmDiacritics(tt.input); res != tt.expected {
			t.Errorf("For input %v, expected %s, got %s", tt.input, tt.expected, res)
		}
	}
}
