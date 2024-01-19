package wglog

import "golang.zx2c4.com/wireguard/device"

// FmtFn is a function capable of outputting a printf style string.
type FmtFn func(string, ...any)

// Multi will emit a logged message on all given loggers.
func Multi(loggers ...*device.Logger) *device.Logger {
	verbosef, errorf := collectFns(loggers)
	return &device.Logger{Verbosef: multi(verbosef), Errorf: multi(errorf)}
}

func collectFns(loggers []*device.Logger) (verbosef, errorf []FmtFn) {
	for _, l := range loggers {
		verbosef = append(verbosef, elseDefault(l.Verbosef, noopFn))
		errorf = append(errorf, elseDefault(l.Errorf, noopFn))
	}
	return
}

func multi(fns []FmtFn) func(string, ...any) {
	return func(s string, a ...any) {
		for _, fn := range fns {
			fn(s, a...)
		}
	}
}
