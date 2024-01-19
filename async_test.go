package wglog

import (
	"github.com/stretchr/testify/require"
	"golang.zx2c4.com/wireguard/device"
	"runtime"
	"sync"
	"testing"
)

func TestAsync(t *testing.T) {
	t.Run("normal", func(t *testing.T) {
		var v, va, e, ea bool
		var w sync.WaitGroup
		w.Add(2)
		l := Async(&device.Logger{
			Verbosef: func(string, ...any) {
				defer w.Done()
				v = true
				// Check if func was called with `go`
				// https://stackoverflow.com/a/56702614
				va = runtime.Callers(3, make([]uintptr, 1)) == 0
			},
			Errorf: func(string, ...any) {
				defer w.Done()
				e = true
				// Check if func was called with `go`
				// https://stackoverflow.com/a/56702614
				ea = runtime.Callers(3, make([]uintptr, 1)) == 0
			},
		})
		l.Verbosef("")
		l.Errorf("")
		w.Wait()
		require.True(t, v, "failed to call verbose")
		require.True(t, va, "failed to call verbose in a goroutine")
		require.True(t, e, "failed to call error")
		require.True(t, ea, "failed to call error in a goroutine")
	})
	t.Run("noop", func(t *testing.T) {
		l := Async(new(device.Logger))
		require.NotNil(t, l.Verbosef)
		require.NotNil(t, l.Errorf)
		l.Verbosef("")
		l.Errorf("")
	})
}
