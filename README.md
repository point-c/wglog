# wglog
[![Go Reference](https://pkg.go.dev/badge/github.com/point-c/wglog@v0.0.2.svg)](https://godocs.io/github.com/point-c/wglog@v0.0.3)

This is a set of extensions for `Wireguard-go`'s logger.

## Loggers

- `noop`: Non-nil logger that does nothing.
- `slog`: Converts a `slog.Logger` to a `device.Logger`.
  - Logging is done on the following levels:
    - `Verbosef`
      - `slog.LevelDebug`
    - `Errorf`
      - `slog.LevelError`
- `multi`: Emits log messages on multiple loggers. 
- `async`: Runs the logger funcs in a goroutine.