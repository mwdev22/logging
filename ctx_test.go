package logging

import (
	"context"
	"testing"
)

func Test_fromcontext(t *testing.T) {
	stored := NewJsonLogger()

	tests := []struct {
		name string
		ctx  interface{}
		want Logger
	}{
		{name: "logger stored via ToContext", ctx: ToContext(context.Background(), stored), want: stored},
		{name: "logger passed directly", ctx: stored, want: stored},
		{name: "nil logger stored falls back to default", ctx: ToContext(context.Background(), nil), want: DefaultLogger()},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FromContext(tt.ctx)
			if got != tt.want {
				t.Errorf("FromContext() = %T, want %T", got, tt.want)
			}
		})
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
