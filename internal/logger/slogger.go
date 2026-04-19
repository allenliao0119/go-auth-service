package logger

import (
	"context"
	"log/slog"
	"os"

	"github.com/allenliao0119/go-auth-service/internal/config"
)

type slogger struct {
	logger *slog.Logger
}

func NewSlogger(cfg config.LoggerConfig) Logger {
	return &slogger{
		logger: slog.New(handler(cfg)),
	}
}

func (l *slogger) With(ctx context.Context) Logger {
	return &slogger{
		logger: l.logger.With(attrsFromLoggedCtxKeys(ctx)...),
	}
}

func (l *slogger) WithAdditionalField(fields map[string]any) Logger {
	return &slogger{
		logger: l.logger.With(attrsFromFields(fields)...),
	}
}

func (l *slogger) Info(ctx context.Context, msg string, args ...any) {
	l.logger.InfoContext(ctx, msg, args...)
}

func (l *slogger) Debug(ctx context.Context, msg string, args ...any) {
	l.logger.DebugContext(ctx, msg, args...)
}

func (l *slogger) Warn(ctx context.Context, msg string, args ...any) {
	l.logger.WarnContext(ctx, msg, args...)
}

func (l *slogger) Error(ctx context.Context, msg string, args ...any) {
	l.logger.ErrorContext(ctx, msg, args...)
}

func level(level config.LogLevel) slog.Level {
	switch level {
	case config.LogLevelDebug:
		return slog.LevelDebug
	case config.LogLevelWarn:
		return slog.LevelWarn
	case config.LogLevelError:
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}

func handler(cfg config.LoggerConfig) slog.Handler {
	if cfg.Format == config.LogFormatJSON {
		return slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: level(cfg.Level)})
	}
	return slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: level(cfg.Level)})
}

func attrsFromFields(fields map[string]any) []any {
	attrs := make([]any, 0, len(fields))
	for k, v := range fields {
		attrs = append(attrs, slog.Any(k, v))
	}
	return attrs
}

func attrsFromLoggedCtxKeys(ctx context.Context) []any {
	var attrs []any
	for _, key := range loggedContextKey {
		if val := ctx.Value(key); val != nil {
			attrs = append(attrs, slog.Any(string(key), val))
		}
	}
	return attrs
}
