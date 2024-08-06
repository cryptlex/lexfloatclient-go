// +build windows

package lexfloatclient

import "C"
import (
	"bytes"
	"encoding/binary"
	"unicode/utf16"
	"unsafe"
)

const (
	maxCArrayLength  C.uint = 4096
	maxGoArrayLength C.int  = 4096
)

func goToCString(goString string) *C.ushort {
	bytes := []rune(goString)
	encodedBytes := utf16.Encode(bytes)
	// Ensure the slice is null-terminated
	encodedBytes = append(encodedBytes, 0)
	cString := (*C.ushort)(unsafe.Pointer(&encodedBytes[0]))
	return cString
}

func ctoGoString(cString *C.ushort) string {
	encodedBytes := C.GoBytes(unsafe.Pointer(cString), maxGoArrayLength)
	goString, _ := decodeUtf16(encodedBytes, binary.LittleEndian)
	return goString
}

func getCArray() [maxCArrayLength]C.ushort {
	var cArray [maxCArrayLength]C.ushort
	return cArray
}

func freeCString(cString *C.ushort) {
	// do nothing
}

func decodeUtf16(b []byte, order binary.ByteOrder) (string, error) {
	ints := make([]uint16, len(b)/2)
	if err := binary.Read(bytes.NewReader(b), order, &ints); err != nil {
		return "", err
	}
	return string(utf16.Decode(ints)), nil
}
