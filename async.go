package wglog

// Async will call the underlying logger methods in a new goroutine
func Async(logger *Logger) *Logger {
	return &Logger{
		Verbosef: func(format string, args ...any) { go logger.Verbosef(format, args...) },
		Errorf:   func(format string, args ...any) { go logger.Errorf(format, args...) },
	}
}
