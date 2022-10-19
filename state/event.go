package state

import (
	"context"
	"sync"

	"github.com/wailsapp/wails/v2/pkg/runtime"
	"github.com/zsa/wally/usb"
)

type DeviceConnectionEvent struct {
	Device usb.USBDevice `json:"device"`
}

type DeviceSelectedEvent struct {
	Device usb.USBDevice `json:"device"`
}

type DeviceDisconnectionEvent struct {
	Fingerprint int `json:"fingerprint"`
}

type FirmwareVersionEvent struct {
	LayoutId   string `json:"layoutId"`
	RevisionId string `json:"revisionId"`
}

type LayerChangeEvent struct {
	LayerNum byte `json:"layerNum"`
}

type KeyPressEvent struct {
	Column  byte `json:"column"`
	Row     byte `json:"row"`
	Pressed bool `json:"pressed"`
}

type StepChangeEvent struct {
	Step Step `json:"step"`
}

type LogEvent struct {
	Log log `json:"log"`
}

type ProgressEvent struct {
	Current int `json:"current"`
	Total   int `json:"total"`
}

var lock = &sync.Mutex{}

type uiEventEmitter struct {
	ctx context.Context
}

// The event emitter is a singleton that allows sending event to the UI Thread
// It is created right after the DOM Ready event
var uiEvent *uiEventEmitter

func InitUIEventEmitter(ctx context.Context) {
	if uiEvent == nil {
		lock.Lock()
		defer lock.Unlock()
		if uiEvent == nil {
			uiEvent = &uiEventEmitter{ctx: ctx}
		}
	}
}

func (e *uiEventEmitter) Emit(eventName string, eventData ...interface{}) {
	runtime.EventsEmit(e.ctx, eventName, eventData...)
}
