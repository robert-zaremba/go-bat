package bat

import (
	"encoding/binary"
	"encoding/hex"
	"io"
	"os"

	"github.com/spaolacci/murmur3"
)

// HashOfFile computes murmur3 hash of the given file via its path
func HashOfFile(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	hash := murmur3.New64()
	if _, err := io.Copy(hash, file); err != nil {
		_ = file.Close()
		return "", err
	}
	res := make([]byte, 8)
	binary.LittleEndian.PutUint64(res, hash.Sum64())
	return hex.EncodeToString(res), file.Close()
}
