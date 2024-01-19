package wglog

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"golang.zx2c4.com/wireguard/device"
	"testing"
)

func TestMulti(t *testing.T) {
	for _, i := range []int{1, 5, 10, 123} {
		t.Run(fmt.Sprintf("testing with %d loggers", i), func(t *testing.T) {
			loggers, check := getMultipleLoggers(t, i)
			logger := Multi(loggers...)
			logger.Verbosef("")
			logger.Errorf("")
			check()
		})
	}
}

func getMultipleLoggers(t *testing.T, n int) (loggers []*device.Logger, checkAllCalled func()) {
	t.Helper()
	checkAllCalled = func() {}
	for i := 0; i < n; i++ {
		var v, e bool
		l := device.Logger{
			Verbosef: func(string, ...any) { v = true },
			Errorf:   func(string, ...any) { e = true },
		}
		loggers = append(loggers, &l)
		fn := checkAllCalled
		checkAllCalled = func() {
			fn()
			require.True(t, v)
			require.True(t, e)
		}
	}
	return
}
