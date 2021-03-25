// +build linux darwin

package lexfloatclient

//#include <stdlib.h>
import "C"
import "unsafe"

const (
	maxCArrayLength  C.uint = 256
	maxGoArrayLength C.int  = 256
)

func goToCString(data string) *C.char {
	cString := C.CString(data)
	return cString
}

func ctoGoString(cString *C.char) string {
	goString := C.GoStringN(cString, maxGoArrayLength)
	return goString
}

func getCArray() [maxCArrayLength]C.char {
	var cArray [maxCArrayLength]C.char
	return cArray
}

func freeCString(cString *C.char) {
	defer C.free(unsafe.Pointer(cString))
}
