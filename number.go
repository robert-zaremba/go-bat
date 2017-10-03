package bat

import (
	"fmt"
	"math"
	"strconv"

	"github.com/robert-zaremba/errstack"
)

// I64toa converts int64 value to 10-based string
func I64toa(x int64) string {
	return strconv.FormatInt(x, 10)
}

// I64tox converts int64 value to 16-based string
func I64tox(x int64) string {
	return strconv.FormatInt(x, 16)
}

// F64toa converts float64 value to 10-based string.
// Function takes optional argument - precision - which is described in strconv.FormatFloat
func F64toa(x float64, precision ...int) string {
	p := -1
	if len(precision) > 0 {
		p = precision[0]
	}
	return strconv.FormatFloat(x, 'f', p, 64)
}

// Atoi64 converts 10-based string into int64 value.
func Atoi64(s string) (int64, error) {
	return strconv.ParseInt(s, 10, 64)
}

// Atoi64Errp converts 10-based string into int64 value and sets the error in the putter.
func Atoi64Errp(s string, errp errstack.Putter) int64 {
	v, err := Atoi64(s)
	if err != nil {
		errp.Put(err)
	}
	return v
}

// Atof64 converts 10-based string into float64 value.
func Atof64(s string) (float64, error) {
	return strconv.ParseFloat(s, 64)
}

// Atof64Errp converts 10-based string into float64 value and sets the error in the putter.
func Atof64Errp(s string, errp errstack.Putter) float64 {
	v, err := Atof64(s)
	if err != nil {
		errp.Put(err)
	}
	return v
}

// HumanizeInt produces a human readable representation of an SI size.
func HumanizeInt(n uint64) string {
	if n < 10 {
		return fmt.Sprintf("%d B", n)
	}
	const base = 1000
	var sizes = []string{"B", "kB", "MB", "GB", "TB", "PB", "EB"}
	e := math.Floor(math.Log(float64(n)) / math.Log(base))
	suffix := sizes[int(e)]
	val := math.Floor(float64(n)/math.Pow(base, e)*10+0.5) / 10
	f := "%.0f %s"
	if val < 10 {
		f = "%.1f %s"
	}
	return fmt.Sprintf(f, val, suffix)
}

// CmpInt64Pairs compares pairs of int64 for sorting
func CmpInt64Pairs(ls [][2]int64) int {
	for i := range ls {
		if x := ls[i][0] - ls[i][1]; x != 0 {
			return int(x)
		}
	}
	return 0
}

// Min3Int returns the minimum of 3 integeres.
func Min3Int(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		}
	} else {
		if b < c {
			return b
		}
	}
	return c
}
