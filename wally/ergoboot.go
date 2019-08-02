// +build darwin

package wally

/*

#cgo darwin CFLAGS: -O2 -Wall
#cgo darwin LDFLAGS: -framework IOKit -framework CoreFoundation
#include "ergoboot.h"

*/
import "C"

func Ergoboot() {
	C.ergoboot()
}
