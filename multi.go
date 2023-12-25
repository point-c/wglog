package wglog

// Multi will broadcast a logged message on all given loggers
func Multi(loggers ...*Logger) *Logger {
	verbosef, errorf := func(string, ...any) {}, func(string, ...any) {}
	appendFmtFn := func(fn1, fn2 func(string, ...any)) func(string, ...any) {
		return func(s string, a ...any) { fn1(s, a...); fn2(s, a...) }
	}
	for _, l := range loggers {
		if l.Verbosef != nil {
			verbosef = appendFmtFn(l.Verbosef, verbosef)
		}
		if l.Errorf != nil {
			errorf = appendFmtFn(l.Errorf, errorf)
		}
	}
	return &Logger{Verbosef: verbosef, Errorf: errorf}
}
