package bat

import (
	"math/rand"
	"reflect"
	"regexp"
	"strings"
	"unsafe"

	"github.com/mozillazg/go-unidecode"
)

var stripper = regexp.MustCompile("  +")

// StrJoin joins all tail arguments using the first argument
func StrJoin(sep string, args ...string) string {
	return strings.Join(args, sep)
}

// StrTruncateUpTo cut a part of a string `s` after n occurrence of `r`
func StrTruncateUpTo(s string, r rune, n int) string {
	var c rune
	var i, seen int
	for i, c = range s {
		if c == r {
			seen++
		}
		if seen > n {
			return s[:i]
		}
	}
	if c == '/' {
		return s[:i]
	}
	return s
}

// StrIsNum check if a string is a number (contains only digits)
func StrIsNum(s string) bool {
	for _, c := range s {
		if c < '0' || c > '9' {
			return false
		}
	}
	return true
}

// StrTrimMultiSpace replaces multiple spaces with one space and
// also trims space from both ends
func StrTrimMultiSpace(s string) string {
	return strings.Trim(stripper.ReplaceAllString(s, " "), " ")
}

// TrimSuffixes returns s without any of the provided trailing suffixes strings.
func TrimSuffixes(s string, suffixes ...string) string {
	for _, suffix := range suffixes {
		if strings.HasSuffix(s, suffix) {
			return strings.TrimSuffix(s, suffix)
		}
	}
	return s
}

// RmDiacritics removes diacritics from a string. If non-alphanumeric character is encountered
// diacritics are removed from it. If removing diacritics is not possible, character is removed.
func RmDiacritics(s string) string {
	return unidecode.Unidecode(s)
}

// RandomCode generates a random code
// where random long integer is serialized into hex string
func RandomCode() string {
	return I64tox(rand.Int63())
}

// UnsafeByteArrayToStr uses unsafe to convert byte array into string. Supplied array cannot be
// altered after this functions is called
func UnsafeByteArrayToStr(b []byte) string {
	if b == nil {
		return ""
	}
	return *(*string)(unsafe.Pointer(&b))
}

// UnsafeStrToByteArray uses unsafe to convert string into byte array. Returned array cannot be
// altered after this functions is called
func UnsafeStrToByteArray(s string) []byte {
	sh := *(*reflect.SliceHeader)(unsafe.Pointer(&s))
	sh.Cap = sh.Len
	bs := *(*[]byte)(unsafe.Pointer(&sh))
	return bs
}

// MatchesPrefixes checks if given string has a prefix from given prefix list
func MatchesPrefixes(s string, prefixes []string) bool {
	for _, prefix := range prefixes {
		if strings.HasPrefix(s, prefix) {
			return true
		}
	}
	return false
}

// StrIterator provides a generator of names / strings
type StrIterator interface {
	Get() string
	Next() bool
}
