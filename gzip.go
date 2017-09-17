package bat

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
)

// GzipJSON encodes as JSON and compress with Gzip in default compression level
func GzipJSON(i interface{}) ([]byte, error) {
	return gzipJSON(i, gzip.DefaultCompression)
}

// GzipJSONFast encodes as JSON and compress with Gzip in fastest compression level
func GzipJSONFast(i interface{}) ([]byte, error) {
	return gzipJSON(i, gzip.BestSpeed)
}

// GzipJSONBest encodes as JSON and compress with Gzip in best compression level
func GzipJSONBest(i interface{}) ([]byte, error) {
	return gzipJSON(i, gzip.BestCompression)
}

// gzipJSON encodes given data into JSON and returns its Gzipped version.
func gzipJSON(i interface{}, level int) ([]byte, error) {
	var b bytes.Buffer
	w, err := gzip.NewWriterLevel(&b, level)
	if err != nil {
		return nil, err
	}
	if err := json.NewEncoder(w).Encode(i); err != nil {
		return nil, err
	}
	if err := w.Close(); err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
