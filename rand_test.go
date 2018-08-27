package bat

import (
	"strings"
	"testing"
)

func TestRandomASCII(t *testing.T) {
	const length = 60

	// Generate 1000 strings and check that they are unique
	ls := make([]string, 1000)
	for i := range ls {
		ls[i] = RandomASCII(length)
		if len(ls[i]) != length {
			t.Fatalf("wrong length: expected %d, got %d", length, len(ls[i]))
		}
	}
	for i, s1 := range ls {
		for j, s2 := range ls {
			if i != j && s1 == s2 {
				t.Fatalf("not unique: %d:%q and %d:%q", i, s1, j, s2)
			}
		}
	}
}

func checkStrOfSet(str, set string, t *testing.T) {
	for _, c := range str {
		if !strings.ContainsRune(set, c) {
			t.Fatal("Generated string should contain only required characters. Got: ", c)
		}
	}
}

func TestRandomStrChars(t *testing.T) {
	charsStr := "aBk1"
	chars := []byte(charsStr)
	const length = 100
	ls := make([]string, 2)
	for i := range ls {
		s := RandomStrChars(length, chars)
		ls[i] = s
		if len(s) != length {
			t.Fatalf("wrong length: expected %d, got %d", length, len(s))
		}
		checkStrOfSet(s, charsStr, t)
	}
	if ls[0] == ls[1] {
		t.Fatalf("not unique: %q", ls)
	}
}

func TestRandomCode(t *testing.T) {
	n := 10000
	m := map[string]struct{}{}
	for i := 0; i < n; i++ {
		m[RandomCode()] = struct{}{}
	}
	if len(m) != n {
		t.Errorf("Expected %d, got %d", n, len(m))
	}
}

func BenchmarkRandomCode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = RandomCode()
	}
}
