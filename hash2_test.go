package bat

import (
	"bytes"
	"testing"
)

// For test input do following steps:
// pip install mmh3
// >>> import mmh3
// >>> mm3.hash_bytes("trueagflow")
// `trueagflow` is the traversion result
func TestMurmurHash(t *testing.T) {
	s := struct {
		A bool
		B string
	}{
		A: true,
		B: "agflow",
	}
	expected := []byte{81, 208, 60, 196, 108, 207, 233, 96, 29, 230, 45, 180, 51, 14, 198, 38}
	res, err := MurmurHash(s)
	if err != nil || !bytes.Equal(res, expected) {
		t.Errorf("Expected %v, got %v with error %v", expected, res, err)
	}
}
