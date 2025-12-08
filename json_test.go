package logging

import (
	"bytes"
	"encoding/json"
	"log/slog"
	"strings"
	"testing"
)

func Test_jsonlogger_implementsloggerinterface(t *testing.T) {
	var _ Logger = &JsonLogger{}
}

func Test_newjsonlogger(t *testing.T) {
	logger := NewJsonLogger()

	if logger == nil {
		t.Fatal("json logger init")
	}

	if logger.logger == nil {
		t.Fatal("json logger is nil")
	}
}

func Test_jsonlogger_withwriter(t *testing.T) {
	var buf bytes.Buffer
	logger := NewJsonLoggerWithWriter(&buf)

	if logger == nil {
		t.Fatal("NewJsonLoggerWithWriter returned nil")
	}

	if logger.logger == nil {
		t.Fatal("logger.logger is nil")
	}
}

func Test_jsonlogger_printf(t *testing.T) {
	var buf bytes.Buffer
	logger := NewJsonLoggerWithWriter(&buf)

	logger.Printf("test message: %s", "hello")

	output := buf.String()

	var logentry map[string]interface{}
	if err := json.Unmarshal([]byte(output), &logentry); err != nil {
		t.Fatalf("output is not valid json: %v", err)
	}

	if msg, ok := logentry["msg"].(string); !ok || !strings.Contains(msg, "test message: hello") {
		t.Errorf("expected message to contain 'test message: hello', got: %v", logentry["msg"])
	}

	if level, ok := logentry["level"].(string); !ok || !strings.EqualFold(level, "INFO") {
		t.Errorf("expected level to be info, got: %v", logentry["level"])
	}
}

func Test_jsonlogger_println(t *testing.T) {
	var buf bytes.Buffer
	logger := NewJsonLoggerWithWriter(&buf)

	logger.Println("simple", "message")

	output := buf.String()

	var logentry map[string]interface{}
	if err := json.Unmarshal([]byte(output), &logentry); err != nil {
		t.Fatalf("output is not valid json: %v", err)
	}

	if msg, ok := logentry["msg"].(string); !ok || !strings.Contains(msg, "simple") || !strings.Contains(msg, "message") {
		t.Errorf("expected message to contain 'simple' and 'message', got: %v", logentry["msg"])
	}
}

func Test_jsonlogger_formattedoutput(t *testing.T) {
	var buf bytes.Buffer
	logger := NewJsonLoggerWithWriter(&buf)

	logger.Printf("formatted: %d", 42)

	output := buf.String()

	var logentry map[string]interface{}
	if err := json.Unmarshal([]byte(output), &logentry); err != nil {
		t.Fatalf("output is not valid json: %v", err)
	}

	if msg, ok := logentry["msg"].(string); !ok || !strings.Contains(msg, "formatted: 42") {
		t.Errorf("expected message to contain 'formatted: 42', got: %v", logentry["msg"])
	}

	if _, ok := logentry["time"]; !ok {
		t.Error("expected log entry to have 'time' field")
	}
}

func Test_jsonlogger_withcustomhandler(t *testing.T) {
	var buf bytes.Buffer
	handler := slog.NewJSONHandler(&buf, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})

	logger := &JsonLogger{
		logger: slog.New(handler),
	}

	logger.Printf("custom handler test")

	output := buf.String()

	var logentry map[string]interface{}
	if err := json.Unmarshal([]byte(output), &logentry); err != nil {
		t.Fatalf("output is not valid json: %v", err)
	}

	if msg, ok := logentry["msg"].(string); !ok || !strings.Contains(msg, "custom handler test") {
		t.Errorf("expected message to contain 'custom handler test', got: %v", logentry["msg"])
	}
}
