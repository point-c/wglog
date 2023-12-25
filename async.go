package wglog

import "golang.zx2c4.com/wireguard/device"

// Async will call the underlying logger methods in a goroutine.
func Async(logger *device.Logger) *device.Logger {
	return &device.Logger{
		Verbosef: func(format string, args ...any) { go logger.Verbosef(format, args...) },
		Errorf:   func(format string, args ...any) { go logger.Errorf(format, args...) },
	}
}
