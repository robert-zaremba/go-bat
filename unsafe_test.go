package bat

import "testing"

func TestUnsafeByteArrayToStr(t *testing.T) {
	tests := [][]byte{nil, []byte(""), []byte("lola"), []byte("test")}
	expects := []string{"", "", "lola", "test"}

	for i, tt := range tests {
		res := UnsafeByteArrayToStr(tt)
		if res != expects[i] {
			t.Errorf("%d: Expected %s, got %s", i, expects[i], res)
		}
	}
}
