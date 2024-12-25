package uitls

import "unsafe"

type Int interface {
	~int | ~int64 | ~int32 | ~int16 | ~int8
}

type Number interface {
	Int | ~float32 | ~float64
}

func UnsafeString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
