package logger

import "context"

type Logger interface {
	With(ctx context.Context) Logger
	WithAdditionalField(fields map[string]any) Logger

	Info(ctx context.Context, msg string, args ...any)
	Debug(ctx context.Context, msg string, args ...any)
	Warn(ctx context.Context, msg string, args ...any)
	Error(ctx context.Context, msg string, args ...any)
}
