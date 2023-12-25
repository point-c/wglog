package wglog

import (
	"context"
	"github.com/stretchr/testify/require"
	"log/slog"
	"testing"
)

func TestSlog(t *testing.T) {
	t.Run("default", func(t *testing.T) {
		def := slog.Default()
		defer slog.SetDefault(def)
		handler := &testSlogHandler{t: t}
		slog.SetDefault(slog.New(handler))
		l := Slog(nil)
		l.Verbosef("")
		l.Errorf("")
		handler.Check()
	})

	t.Run("", func(t *testing.T) {
		handler := &testSlogHandler{t: t}
		l := Slog(slog.New(handler))
		l.Verbosef("")
		l.Errorf("")
		handler.Check()
	})
}

type testSlogHandler struct {
	t            *testing.T
	debug, error bool
	invalid      *slog.Level
}

func (t *testSlogHandler) Check() {
	t.t.Helper()
	if t.invalid != nil {
		require.NotNil(t.t, t.invalid, "level(%s)", t.invalid.String())
	}
	require.True(t.t, t.debug, "debug was not called")
	require.True(t.t, t.error, "error was not called")
}

func (t *testSlogHandler) Handle(_ context.Context, record slog.Record) error {
	t.t.Helper()
	switch record.Level {
	case slog.LevelDebug:
		t.debug = true
	case slog.LevelError:
		t.error = true
	default:
		t.invalid = &record.Level
	}
	return nil
}

func (t *testSlogHandler) Enabled(context.Context, slog.Level) bool { t.t.Helper(); return true }
func (t *testSlogHandler) WithAttrs([]slog.Attr) slog.Handler       { t.t.Helper(); return t }
func (t *testSlogHandler) WithGroup(string) slog.Handler            { t.t.Helper(); return t }
