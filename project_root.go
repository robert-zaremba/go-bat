package bat

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/kardianos/osext"
	"github.com/robert-zaremba/errstack"
)

// FindRoot returns an absolute path to the project root, defined by the existance of
// .git directory or .hg directory based on the caller stack trace or path to
// the executable file.
// It takes one optional parameter: number of directories to travers up to search for the
// project root. Be default it's 6.
func FindRoot(levels ...uint) (string, errstack.E) {
	var levelsUp uint = 6
	if len(levels) >= 1 {
		levelsUp = levels[0]
	}
	root, callerErr := findRootFromRuntimeCaller(2, levelsUp)
	if callerErr == nil {
		return root, nil
	}
	root, executableErr := findRootFromExecutable(levelsUp)
	if executableErr == nil {
		return root, nil
	}
	return root, errstack.WrapAsInf(
		errstack.Join(callerErr, executableErr), "Can't locate project root")
}

// findRootFromRuntimeCaller - as `FindRoot`, but limits the search strategy to looking
// up based on the stack trace.
// The argument `callerSkip` is the number of stack frames to ascend, as specified
// in `runtimeCaller`
func findRootFromRuntimeCaller(callerSkip int, levels uint) (string, error) {
	_, filename, _, ok := runtime.Caller(1)
	if !ok {
		return "", fmt.Errorf("Can't retrive runtime caller info")
	}
	root, ok := traverseUpForRoot(filename, levels)
	if !ok {
		return "", fmt.Errorf("Unable to find project root for runtime caller position %s", filename)
	}
	return root, nil
}

// findRootFromExecutable - as `FindRoot`, but limits the search strategy to looking
// up based on the path to Executable file.
func findRootFromExecutable(levels uint) (string, error) {
	filename, err := osext.Executable()
	if err != nil {
		return "", fmt.Errorf("Can't retrive runtime path to executable: %s", err.Error())
	}
	root, ok := traverseUpForRoot(filename, levels)
	if !ok {
		return "", fmt.Errorf("Unable to find project root for runtime executable position %s", filename)
	}
	return root, nil
}

func traverseUpForRoot(p string, levels uint) (string, bool) {
	p = filepath.Dir(p)
	for i := uint(0); i < levels; i++ {
		if isRootPath(p) {
			return p, true
		}
		p = filepath.Dir(p)
	}
	return "", false
}

func isRootPath(p string) bool {
	return exists(filepath.Join(p, ".git")) || exists(filepath.Join(p, ".hg"))
}

func exists(name string) bool {
	_, err := os.Stat(name)
	if err != nil && !os.IsNotExist(err) {
		panic(err)
	}
	return err == nil || !os.IsNotExist(err)
}
