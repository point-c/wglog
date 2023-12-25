package wglog

import "golang.zx2c4.com/wireguard/device"

// Multi will emit a logged message on all given loggers.
// If the logger or any of its funcs are nil they will be skipped.
func Multi(loggers ...*device.Logger) *device.Logger {
	verbosef, errorf := func(string, ...any) {}, func(string, ...any) {}
	appendFmtFn := func(fn1, fn2 func(string, ...any)) func(string, ...any) {
		return func(s string, a ...any) { fn1(s, a...); fn2(s, a...) }
	}
	for _, l := range loggers {
		if l != nil {
			if l.Verbosef != nil {
				verbosef = appendFmtFn(l.Verbosef, verbosef)
			}
			if l.Errorf != nil {
				errorf = appendFmtFn(l.Errorf, errorf)
			}
		}
	}
	return &device.Logger{Verbosef: verbosef, Errorf: errorf}
}
