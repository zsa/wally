package usb

import (
	"sync"
)

var lock = &sync.Mutex{}

type callback struct {
	cb func(connected bool)
}

var _callback *callback

func RegisterCallback(cb func(connected bool)) {
	if _callback == nil {
		lock.Lock()
		defer lock.Unlock()
		if _callback == nil {
			_callback = &callback{}
		}
	}

	_callback.cb = cb
}
