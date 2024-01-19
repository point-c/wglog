package wglog

import (
	"fmt"
	"golang.zx2c4.com/wireguard/device"
	"log/slog"
	"reflect"
	"slices"
)

// Slog creates a [device.Logger] instance that is backed by a specified [slog.Logger].
// No args are passed to the slog logger, instead the message is created from the format string and args passed to the [device.Logger].Verbosef and [device.Logger].Errorf funcs.
// Verbose messages are logged at the Debug level, while errors are logged at the Error level.
func Slog(logger *slog.Logger) *device.Logger {
	logger = elseDefault(logger, slog.Default())
	return &device.Logger{
		Verbosef: func(format string, args ...any) { logger.Debug(fmt.Sprintf(format, args...)) },
		Errorf:   func(format string, args ...any) { logger.Error(fmt.Sprintf(format, args...)) },
	}
}

// validNil is a set of [reflect.Kind] that will not panic on [reflect.Value.IsNil] if [reflect.Value] is valid.
var validNil = []reflect.Kind{reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Pointer, reflect.Slice}

func elseDefault[T any](v, def T) T {
	if vv := reflect.ValueOf(v); slices.Contains(validNil, vv.Kind()) && (vv.IsZero() || vv.IsNil()) {
		return def
	}
	return v
}
