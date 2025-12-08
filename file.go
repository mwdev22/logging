package logging

import (
	"fmt"
	"io"
	"log/slog"
	"os"
)

type FileLogger struct {
	filepath string
	logger   *slog.Logger
	file     io.WriteCloser
}

func NewFileLogger(filepath string) *FileLogger {
	file, err := os.OpenFile(filepath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		file = os.Stdout
	}

	handler := slog.NewTextHandler(file, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	})

	return &FileLogger{
		filepath: filepath,
		logger:   slog.New(handler),
		file:     file,
	}
}

func newFileLoggerWithWriter(writer io.Writer) *FileLogger {
	handler := slog.NewTextHandler(writer, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	})

	return &FileLogger{
		logger: slog.New(handler),
	}
}

func (l *FileLogger) Printf(format string, v ...interface{}) {
	l.logger.Info(fmt.Sprintf(format, v...))
}

func (l *FileLogger) Println(v ...interface{}) {
	l.logger.Info(fmt.Sprint(v...))
}

func (l *FileLogger) Close() error {
	if l.file != nil {
		if closer, ok := l.file.(io.Closer); ok && l.file != os.Stdout {
			return closer.Close()
		}
	}
	return nil
}
