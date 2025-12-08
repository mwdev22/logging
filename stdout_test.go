package logging

import (
	"bytes"
	"log"
	"strings"
	"testing"
)

func Test_stdlogger_implementsloggerinterface(t *testing.T) {
	var _ Logger = stdlogger{}
}

func Test_stdlogger_printf(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer log.SetOutput(log.Writer())

	logger := stdlogger{}
	logger.Printf("test message: %s", "hello")

	output := buf.String()
	if !strings.Contains(output, "test message: hello") {
		t.Errorf("expected output to contain 'test message: hello', got: %s", output)
	}
}

func Test_stdlogger_println(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer log.SetOutput(log.Writer())

	logger := stdlogger{}
	logger.Println("test", "message")

	output := buf.String()
	if !strings.Contains(output, "test") || !strings.Contains(output, "message") {
		t.Errorf("expected output to contain 'test' and 'message', got: %s", output)
	}
}

func Test_defaultlogger(t *testing.T) {
	log := DefaultLogger()

	if log == nil {
		t.Fatal("DefaultLogger returned nil")
	}

	var _ Logger = log
}

func Test_defaultlogger_printf(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer log.SetOutput(log.Writer())

	logger := DefaultLogger()
	logger.Printf("formatted: %d", 42)

	output := buf.String()
	if !strings.Contains(output, "formatted: 42") {
		t.Errorf("expected output to contain 'formatted: 42', got: %s", output)
	}
}

func Test_defaultlogger_println(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer log.SetOutput(log.Writer())

	logger := DefaultLogger()
	logger.Println("simple", "message")

	output := buf.String()
	if !strings.Contains(output, "simple") || !strings.Contains(output, "message") {
		t.Errorf("expected output to contain 'simple' and 'message', got: %s", output)
	}
}
