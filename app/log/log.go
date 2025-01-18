package log

import (
	"context"
	"log/slog"
)

type loggerKey struct{}

func FromCtx(ctx context.Context) *slog.Logger {
	value := ctx.Value(loggerKey{})
	if value == nil {
		return nil
	}
	return value.(*slog.Logger)
}

func doWithLogger(ctx context.Context, action func(*slog.Logger)) {
	logger := FromCtx(ctx)
	if logger == nil {
		return
	}
	action(logger)
}
