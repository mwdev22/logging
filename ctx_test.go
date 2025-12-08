package logging

import (
	"context"
	"testing"
)

func Test_fromcontext_withlogger(t *testing.T) {
	logger := DefaultLogger()

	result := FromContext(logger)

	if result == nil {
		t.Fatal("FromContext returned nil")
	}

	if result != logger {
		t.Error("FromContext did not return the expected logger")
	}
}

func Test_fromcontext_withoutlogger(t *testing.T) {
	ctx := context.Background()

	result := FromContext(ctx)

	if result == nil {
		t.Fatal("FromContext returned nil")
	}

	_, ok := result.(stdlogger)
	if !ok {
		t.Error("FromContext did not return default logger")
	}
}

func Test_tocontext(t *testing.T) {
	ctx := context.Background()
	logger := DefaultLogger()

	newctx := ToContext(ctx, logger)

	if newctx == nil {
		t.Fatal("ToContext returned nil context")
	}

	storedlogger := newctx.Value(CtxKey)
	if storedlogger == nil {
		t.Fatal("logger not stored in context")
	}

	if storedlogger != logger {
		t.Error("stored logger does not match original logger")
	}
}

func Test_tocontext_andretrieve(t *testing.T) {
	ctx := context.Background()
	log := NewJsonLogger()

	newctx := ToContext(ctx, log)
	retrieved := newctx.Value(CtxKey)

	if retrieved == nil {
		t.Fatal("retrieved logger is nil")
	}

	retrievedlogger, ok := retrieved.(Logger)
	if !ok {
		t.Fatal("retrieved value is not a Logger")
	}

	if retrievedlogger != log {
		t.Error("retrieved logger does not match stored logger")
	}
}

func Test_contextkey(t *testing.T) {
	if CtxKey != "logger" {
		t.Errorf("expected CtxKey to be 'logger', got %s", CtxKey)
	}
}