package wglog

import (
	"github.com/stretchr/testify/require"
	"golang.zx2c4.com/wireguard/device"
	"runtime"
	"sync"
	"testing"
)

func TestAsync(t *testing.T) {
	var v, va, e, ea bool
	var wg sync.WaitGroup
	wg.Add(2)
	l := Async(&device.Logger{
		Verbosef: func(string, ...any) {
			t.Helper()
			defer wg.Done()
			v = true
			va = runtime.Callers(3, make([]uintptr, 1)) == 0
		},
		Errorf: func(string, ...any) {
			t.Helper()
			defer wg.Done()
			e = true
			ea = runtime.Callers(3, make([]uintptr, 1)) == 0
		},
	})
	l.Verbosef("")
	l.Errorf("")
	wg.Wait()
	require.True(t, v, "failed to call verbose")
	require.True(t, va, "failed to call verbose in a goroutine")
	require.True(t, e, "failed to call error")
	require.True(t, ea, "failed to call error in a goroutine")
}
