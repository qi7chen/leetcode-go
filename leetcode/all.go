package leetcode

import (
	"unsafe"
)

// bytes to string
func btos(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
