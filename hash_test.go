package bat

import "testing"

func TestHashOfFile(t *testing.T) {
	expected := "78793ff0e7f4de77"
	h, err := HashOfFile("testdata/randomdata")
	if err != nil || h != expected {
		t.Errorf("Expected %q, got %q, error = %v", expected, h, err)
	}
}
