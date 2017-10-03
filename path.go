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

// IsFileErrp check if `filename` is a valid path to a file and sets the error in the putter.
func IsFileErrp(filename string, errp errstack.Putter) {
	if err := IsFile(filename); err != nil {
		errp.Put(err)
	}
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
	} else if !stat.IsDir() {
		return errstack.NewReq("Expected path to a directory. Got path to a file.")
	}
	return nil
}

// IsDirErrp check if `dir` is a valid path to a directory  and sets the error in the putter.
func IsDirErrp(dir string, errp errstack.Putter) {
	if err := IsFile(dir); err != nil {
		errp.Put(err)
	}
}

// AssertIsDir calls IsDir and panics if error is returned.
// `key` is the log argument to name the path entity.
func AssertIsDir(dir, key string, logger Logger) {
	if err := IsDir(dir); err != nil {
		logger.Fatal("Expected a valid directory path.", key, dir, err)
	}
}
