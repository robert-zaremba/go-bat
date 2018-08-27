package bat

import (
	"crypto/rand"
	mrand "math/rand"
)

// RandomCode generates a random code
// where random long integer is serialized into hex string
func RandomCode() string {
	return I64tox(mrand.Int63())
}

// StdChars is a set of standard characters allowed in uniuri string.
var stdChars = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")

// RandomASCII uses RandomStrChars to returns a new random string consisting of
// ASCII alphabet.
func RandomASCII(length int) string {
	return RandomStrChars(length, stdChars)
}

// RandomStrChars returns a new random string of the provided length, consisting
// of the provided byte slice of allowed characters (maximum 256).
// The generated result is a cryptographically secure uniform (unbiased) strings.
// This code is based on the Go unri package and is a subject of public domain:
//     http://creativecommons.org/publicdomain/zero/1.0/
func RandomStrChars(length int, chars []byte) string {
	if length == 0 {
		return ""
	}
	clen := len(chars)

	if clen < 2 || clen > 256 {
		panic("uniuri: wrong charset length for NewLenChars")
	}
	maxrb := 255 - (256 % clen)
	b := make([]byte, length)
	r := make([]byte, length+(length/4)) // storage for random bytes.
	for i := 0; ; {
		if _, err := rand.Read(r); err != nil {
			panic("bat: error reading random bytes: " + err.Error())
		}
		for _, rb := range r {
			c := int(rb)
			if c > maxrb {
				// Skip this number to avoid modulo bias.
				continue
			}
			b[i] = chars[c%clen]
			i++
			if i == length {
				return string(b)
			}
		}
	}
}
