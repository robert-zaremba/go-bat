package bat

import (
	"encoding/json"
	"io"
	"os"

	"github.com/robert-zaremba/errstack"
)

// DecodeJSON decodes JSON data from `r` into `dest`.
// `dest` has to be a pointer!
func DecodeJSON(r io.Reader, dest interface{}) errstack.E {
	err := json.NewDecoder(r).Decode(dest)
	if err == io.EOF {
		if rc, ok := r.(io.ReadCloser); ok {
			return errstack.WrapAsInf(rc.Close(), "Can't close Reader")
		}
		return nil
	}
	return errstack.WrapAsInf(err, "Can't decode JSON data into given structure")
}

// DecodeJSONFile as `DecodeJSON` but reads data from the file.
func DecodeJSONFile(fname string, dest interface{}, logger errstack.Logger) errstack.E {
	f, err := os.Open(fname)
	if err != nil {
		return errstack.WrapAsInf(err, "Can't open contract schema file: "+fname)
	}
	defer errstack.CallAndLog(logger, f.Close)
	return DecodeJSON(f, dest)
}
