package bat

import (
	"os"

	"github.com/robert-zaremba/errstack"
)

// IsFile check if `filename` is a valid path to a file
func IsFile(filename string) errstack.E {
	stat, err := os.Stat(filename)
	if err != nil {
		return errstack.WrapAsReqF(err, "")
	} else if stat.IsDir() {
		return errstack.NewReq("Expected path to a file. Got path to a directory.")
	}
	return nil
}

// AssertIsFile calls IsFile and panics if error is returned.
// `key` is the log argument to name the path entity.
func AssertIsFile(filename, key string, logger Logger) {
	if err := IsFile(filename); err != nil {
		logger.Fatal("Expected a valid file path.", key, filename, err)
	}
}

// IsDir check if `dir` is a valid path to a directory
func IsDir(dir string) errstack.E {
	stat, err := os.Stat(dir)
	if err != nil {
		return errstack.WrapAsReqF(err, "")
	} else if stat.IsDir() {
		return errstack.NewReq("Expected path to a directory. Got path to a file.")
	}
	return nil
}

// AssertIsDir calls IsDir and panics if error is returned.
// `key` is the log argument to name the path entity.
func AssertIsDir(dir, key string, logger Logger) {
	if err := IsDir(dir); err != nil {
		logger.Fatal("Expected a valid directory path.", key, dir, err)
	}
}
