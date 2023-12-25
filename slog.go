package wglog

import (
	"fmt"
	"golang.zx2c4.com/wireguard/device"
	"log/slog"
)

// Slog creates a [Logger] instance that is backed by a specified [slog.Logger].
// No args are passed to the slog logger, instead the message is created from the format string and args passed to the [Logger.Verbosef] and [Logger.Errorf] funcs.
// Verbose messages are logged at the Debug level, while errors are logged at the Error level.
func Slog(logger *slog.Logger) *device.Logger {
	logger = loggerElseDefault(logger, slog.Default())
	return &device.Logger{
		Verbosef: func(format string, args ...any) { logger.Debug(fmt.Sprintf(format, args...)) },
		Errorf:   func(format string, args ...any) { logger.Error(fmt.Sprintf(format, args...)) },
	}
}

func loggerElseDefault(logger, def *slog.Logger) *slog.Logger {
	if logger == nil {
		return def
	}
	return logger
}
