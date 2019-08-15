// +build darwin

package wally

/*
#cgo darwin CFLAGS: -O2 -Wall -mmacosx-version-min=10.8
#cgo darwin LDFLAGS: -framework IOKit -framework CoreFoundation -mmacosx-version-min=10.8
#include "ergoboot.h"

*/
import "C"

func Ergoboot() {
	C.ergoboot()
}
