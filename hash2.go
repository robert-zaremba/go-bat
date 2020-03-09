package bat

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	hashPkg "hash"

	"github.com/spaolacci/murmur3"
)

const hashLength = 16

func hash(i interface{}, h hashPkg.Hash) error {
	listener := func(i interface{}) error {
		value := fmt.Sprint(i)
		_, err := h.Write(UnsafeStrToByteArray(value))
		return err
	}
	return TraverseObj(i, listener, false)
}

// MurmurHash computes murmur hash of the provided object
func MurmurHash(i interface{}) ([]byte, error) {
	h := murmur3.New128()
	err := hash(i, h)
	if err != nil {
		return nil, err
	}
	return h.Sum(nil), nil
}

// MurmurHashSet computes murmur hash of the provided set of objects. Hash will have the save value
// regardless of objects order.
func MurmurHashSet(elems ...interface{}) ([]byte, error) {
	result := make([]byte, hashLength)
	for _, elem := range elems {
		h, err := MurmurHash(elem)
		if err != nil {
			return result, err
		}
		if len(h) != hashLength {
			return result, fmt.Errorf("Hash has unexpected length (%d). Expected: %d", len(h), hashLength)
		}
		for i := range result {
			result[i] = result[i] ^ h[i]
		}
	}
	return result, nil
}

// MurmurHashBase64 computes murmur string hash of the provided object
func MurmurHashBase64(i interface{}) (string, error) {
	bs, err := MurmurHash(i)
	return base64.StdEncoding.EncodeToString(bs), err
}

// MurmurHashSetBase64 computes murmur string hash of the provided set of objects
func MurmurHashSetBase64(elems ...interface{}) (string, error) {
	bs, err := MurmurHashSet(elems...)
	return base64.StdEncoding.EncodeToString(bs), err
}

// MurmurHashHex takes murmur hash and encodes hash into hex
func MurmurHashHex(content []byte) (string, error) {
	hash, err := MurmurHash(content)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(hash), nil
}
