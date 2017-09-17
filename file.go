package bat

import (
	"io/ioutil"
	"os"

	"github.com/robert-zaremba/errstack"
)

// ReadFile returns the bytes content of the file
func ReadFile(filename string, logger errstack.Logger) ([]byte, errstack.E) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, errstack.WrapAsReq(err, "Can't open file")
	}
	defer errstack.CallAndLog(logger, f.Close)
	data, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, errstack.WrapAsReq(err, "Can't read file")
	}
	return data, nil
}
