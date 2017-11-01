package bat

import (
	"fmt"
	"reflect"
	"unsafe"
)

// UnsafeToString converts string or []byte to string
func UnsafeToString(source interface{}) (string, error) {
	switch t := source.(type) {
	default:
		return "", fmt.Errorf("Unable to parse %T as string. ", source)
	case []byte:
		return UnsafeByteArrayToStr(t), nil
	case string:
		return t, nil
	}
}

// UnsafeToBytes converts string or []byte to []byte
func UnsafeToBytes(source interface{}) ([]byte, error) {
	switch t := source.(type) {
	default:
		return nil, fmt.Errorf("Unable to parse %T as string. ", source)
	case []byte:
		return t, nil
	case string:
		return UnsafeStrToByteArray(t), nil
	}
}

// UnsafeByteArrayToStr uses unsafe to convert byte array into string. Supplied array cannot be
// altered after this functions is called
func UnsafeByteArrayToStr(b []byte) string {
	if b == nil {
		return ""
	}
	return *(*string)(unsafe.Pointer(&b))
}

// UnsafeStrToByteArray uses unsafe to convert string into byte array. Returned array cannot be
// altered after this functions is called
func UnsafeStrToByteArray(s string) []byte {
	sh := *(*reflect.SliceHeader)(unsafe.Pointer(&s))
	sh.Cap = sh.Len
	bs := *(*[]byte)(unsafe.Pointer(&sh))
	return bs
}
