package wglog

import "golang.zx2c4.com/wireguard/device"

// Async will call the underlying logger methods in a goroutine.
// If a field in logger is nil, it will not spawn a goroutine.
func Async(logger *device.Logger) *device.Logger {
	return &device.Logger{
		Verbosef: async(logger.Verbosef),
		Errorf:   async(logger.Errorf),
	}
}

func async(fn FmtFn) (r func(string, ...any)) {
	if fn == nil {
		return noopFn
	}
	return func(s string, a ...any) { go fn(s, a...) }
}
