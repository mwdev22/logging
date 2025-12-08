package logging

import "context"

type contextkey string

const (
	CtxKey contextkey = "logger"
)

func FromContext(ctx interface{}) Logger {
	if logger, ok := ctx.(Logger); ok {
		return logger
	}
	return DefaultLogger()
}

func ToContext(ctx context.Context, logger Logger) context.Context {
	return context.WithValue(ctx, CtxKey, logger)
}
