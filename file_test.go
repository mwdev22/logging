package logging

import (
	"bytes"
	"strings"
	"testing"
)

func Test_filelogger_implementsloggerinterface(t *testing.T) {
	var _ Logger = &FileLogger{}
}

func Test_newfilelogger_withwriter(t *testing.T) {
	var buf bytes.Buffer
	logger := newFileLoggerWithWriter(&buf)

	if logger == nil {
		t.Fatal("NewFileLoggerWithWriter returned nil")
	}

	if logger.logger == nil {
		t.Fatal("logger.logger is nil")
	}
}

func Test_filelogger_printf(t *testing.T) {
	var buf bytes.Buffer
	logger := newFileLoggerWithWriter(&buf)

	logger.Printf("test message: %s", "hello")

	output := buf.String()
	if !strings.Contains(output, "test message: hello") {
		t.Errorf("expected output to contain 'test message: hello', got: %s", output)
	}
}

func Test_filelogger_println(t *testing.T) {
	var buf bytes.Buffer
	logger := newFileLoggerWithWriter(&buf)

	logger.Println("simple", "message")

	output := buf.String()
	if !strings.Contains(output, "simple") || !strings.Contains(output, "message") {
		t.Errorf("expected output to contain 'simple' and 'message', got: %s", output)
	}
}

func Test_filelogger_multiplewrites(t *testing.T) {
	var buf bytes.Buffer
	logger := newFileLoggerWithWriter(&buf)

	logger.Printf("first message")
	logger.Println("second message")
	logger.Printf("third: %d", 3)

	output := buf.String()
	if !strings.Contains(output, "first message") {
		t.Errorf("expected output to contain 'first message', got: %s", output)
	}
	if !strings.Contains(output, "second message") {
		t.Errorf("expected output to contain 'second message', got: %s", output)
	}
	if !strings.Contains(output, "third: 3") {
		t.Errorf("expected output to contain 'third: 3', got: %s", output)
	}
}

func Test_filelogger_formattedoutput(t *testing.T) {
	var buf bytes.Buffer
	logger := newFileLoggerWithWriter(&buf)

	logger.Printf("formatted: %d", 42)

	output := buf.String()
	if !strings.Contains(output, "formatted: 42") {
		t.Errorf("expected output to contain 'formatted: 42', got: %s", output)
	}

	if !strings.Contains(output, "level=INFO") {
		t.Errorf("expected output to contain 'level=INFO', got: %s", output)
	}
}
