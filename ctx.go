package logging

import "context"

type contextkey string

const (
	CtxKey contextkey = "logger"
)

func FromContext(ctx interface{}) Logger {
	switch v := ctx.(type) {
	case context.Context:
		if logger, ok := v.Value(CtxKey).(Logger); ok && logger != nil {
			return logger
		}
	// kept for callers that passed a Logger directly before ctx lookup worked
	case Logger:
		return v
	}
	return DefaultLogger()
}

func ToContext(ctx context.Context, logger Logger) context.Context {
	return context.WithValue(ctx, CtxKey, logger)
}
