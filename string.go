package bat

import (
	"regexp"
	"strings"
	"unicode"

	"github.com/mozillazg/go-unidecode"
)

var stripper = regexp.MustCompile("  +")

// StrIsWhitespace check if string contains only whitespace characters
func StrIsWhitespace(s string) bool {
	for _, r := range s {
		if !unicode.IsSpace(r) {
			return false
		}
	}
	return true
}

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

// MatchesPrefixes checks if given string has a prefix from given prefix list
func MatchesPrefixes(s string, prefixes []string) bool {
	for _, prefix := range prefixes {
		if strings.HasPrefix(s, prefix) {
			return true
		}
	}
	return false
}

// Levenshtein computes the Levenshtein distance. It's a measure of the similarity between
// two strings, which we will refer to as the source string (s) and the target string (t).
// The distance is the number of deletions, insertions, or substitutions required to transform
// s into t. For example, the Levenshtein distance between "Asheville" and "Arizona" is 8.
func Levenshtein(str1, str2 []rune) int {
	s1len := len(str1)
	s2len := len(str2)
	column := make([]int, len(str1)+1)

	for y := 1; y <= s1len; y++ {
		column[y] = y
	}
	for x := 1; x <= s2len; x++ {
		column[0] = x
		lastkey := x - 1
		for y := 1; y <= s1len; y++ {
			oldkey := column[y]
			var incr int
			if str1[y-1] != str2[x-1] {
				incr = 1
			}
			column[y] = Min3Int(column[y]+1, column[y-1]+1, lastkey+incr)
			lastkey = oldkey
		}
	}
	return column[s1len]
}

// StrIterator provides a generator of names / strings
type StrIterator interface {
	Get() string
	Next() bool
}
