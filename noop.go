package wglog

import (
	"golang.zx2c4.com/wireguard/device"
)

var noopFn = func(string, ...any) {}

// Noop is a logger that does not output anything.
// The logger and both of its funcs are not nil.
func Noop() *device.Logger {
	return &device.Logger{
		Verbosef: noopFn,
		Errorf:   noopFn,
	}
}
