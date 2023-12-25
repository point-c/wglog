package wglog

import (
	"golang.zx2c4.com/wireguard/device"
)

type Logger = device.Logger

var noopFn = func(string, ...any) {}

// Noop is a logger that does not output anything.
func Noop() *Logger {
	return &Logger{
		Verbosef: noopFn,
		Errorf:   noopFn,
	}
}
