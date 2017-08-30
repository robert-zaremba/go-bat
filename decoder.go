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
		return nil
	}
	return errstack.WrapAsInf(err, "Can't decode JSON data into given structure")
}

// DecodeJSONFile as `DecodeJSON` but reads data from the file.
func DecodeJSONFile(fname string, dest interface{}, logger errstack.Logger) errstack.E {
	f, err := os.Open(fname)
	defer errstack.CallAndLog(logger, f.Close)
	if err != nil {
		return errstack.WrapAsInf(err, "Can't open contract schema file: "+fname)
	}
	return DecodeJSON(f, dest)
}
