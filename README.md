# wglog

[![Go Reference](https://img.shields.io/badge/godoc-reference-%23007d9c.svg)](https://point-c.github.io/wglog)
[![Go Coverage](https://github.com/point-c/wglog/wiki/coverage.svg)](https://raw.githack.com/wiki/point-c/wglog/coverage.html)

This is a set of extensions for `Wireguard-go`'s logger.

## Installation

To use wglog in your Go project, install it using `go get`:

```bash
go get github.com/point-c/wglog
```

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

## Testing

The package includes tests that demonstrate its functionality. Use Go's testing tools to run the tests:

```bash
go test
```

## Godocs

To regenerate godocs:

```bash
go generate -tags docs ./...
```
