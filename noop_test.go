package wglog

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNoop(t *testing.T) {
	l := Noop()
	require.NotNil(t, l)
	require.NotNil(t, noopFn)
	noopFn("")
	require.NotNil(t, l.Verbosef)
	require.NotNil(t, l.Errorf)
	l.Errorf("")
	l.Verbosef("")
}
