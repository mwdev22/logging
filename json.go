package logging

import (
	"fmt"
	"io"
	"log/slog"
	"os"
)

type JsonLogger struct {
	logger *slog.Logger
}

func NewJsonLogger() *JsonLogger {
	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	})
	return &JsonLogger{
		logger: slog.New(handler),
	}
}

func NewJsonLoggerWithWriter(writer io.Writer) *JsonLogger {
	handler := slog.NewJSONHandler(writer, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	})
	return &JsonLogger{
		logger: slog.New(handler),
	}
}

func (l *JsonLogger) Printf(format string, v ...interface{}) {
	l.logger.Info(fmt.Sprintf(format, v...))
}

func (l *JsonLogger) Println(v ...interface{}) {
	l.logger.Info(fmt.Sprint(v...))
}
