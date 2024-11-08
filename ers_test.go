package ers

/* 100% coverage */
import (
	"errors"
	"strings"
	"testing"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name       string
		message    string
		formatTags []interface{}
		wantErr    bool
		contains   string
	}{
		{
			name:     "simple error",
			message:  "basic error",
			wantErr:  true,
			contains: "basic error",
		},
		{
			name:       "formatted error",
			message:    "count: %d",
			formatTags: []interface{}{42},
			wantErr:    true,
			contains:   "count: 42",
		},
		{
			name:       "multiple format params",
			message:    "%s: %d, %f",
			formatTags: []interface{}{"test", 1, 3.14},
			wantErr:    true,
			contains:   "test: 1, 3.14",
		},
		{
			name:    "empty message",
			message: "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := New(tt.message, tt.formatTags...)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.contains != "" && !strings.Contains(err.Error(), tt.contains) {
				t.Errorf("New() error = %v, should contain %v", err, tt.contains)
			}
		})
	}
}

func TestWrap(t *testing.T) {
	baseErr := errors.New("base error")
	ersErr := New("ers error")

	tests := []struct {
		name     string
		err      error
		details  []string
		wantNil  bool
		contains []string
	}{
		{
			name:     "wrap nil error",
			err:      nil,
			details:  []string{"context"},
			wantNil:  true,
			contains: nil,
		},
		{
			name:     "wrap standard error",
			err:      baseErr,
			details:  []string{"context1", "context2"},
			contains: []string{"base error", "context1", "context2"},
		},
		{
			name:     "wrap ers error",
			err:      ersErr,
			details:  []string{"additional"},
			contains: []string{"ers error", "additional"},
		},
		{
			name:     "wrap with empty details",
			err:      baseErr,
			details:  []string{},
			contains: []string{"base error"},
		},
		{
			name:     "wrap with empty string detail",
			err:      baseErr,
			details:  []string{""},
			contains: []string{"base error"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wrapped := Wrap(tt.err, tt.details...)
			if tt.wantNil {
				if wrapped != nil {
					t.Errorf("Wrap() = %v, want nil", wrapped)
				}
				return
			}

			if wrapped == nil {
				t.Error("Wrap() = nil, want error")
				return
			}

			errStr := wrapped.Error()
			for _, contain := range tt.contains {
				if !strings.Contains(errStr, contain) {
					t.Errorf("Wrap() = %v, should contain %v", errStr, contain)
				}
			}
		})
	}
}

func TestWrapf(t *testing.T) {
	baseErr := errors.New("base error")
	ersErr := New("ers error")

	tests := []struct {
		name       string
		err        error
		format     string
		formatTags []interface{}
		wantNil    bool
		contains   []string
	}{
		{
			name:       "wrapf nil error",
			err:        nil,
			format:     "context %d",
			formatTags: []interface{}{42},
			wantNil:    true,
		},
		{
			name:       "wrapf standard error",
			err:        baseErr,
			format:     "context %d: %s",
			formatTags: []interface{}{1, "test"},
			contains:   []string{"base error", "context 1: test"},
		},
		{
			name:       "wrapf ers error",
			err:        ersErr,
			format:     "value: %f",
			formatTags: []interface{}{3.14},
			contains:   []string{"ers error", "value: 3.14"},
		},
		{
			name:     "wrapf without format args",
			err:      baseErr,
			format:   "plain message",
			contains: []string{"base error", "plain message"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wrapped := Wrapf(tt.err, tt.format, tt.formatTags...)
			if tt.wantNil {
				if wrapped != nil {
					t.Errorf("Wrapf() = %v, want nil", wrapped)
				}
				return
			}

			if wrapped == nil {
				t.Error("Wrapf() = nil, want error")
				return
			}

			errStr := wrapped.Error()
			for _, contain := range tt.contains {
				if !strings.Contains(errStr, contain) {
					t.Errorf("Wrapf() = %v, should contain %v", errStr, contain)
				}
			}
		})
	}
}

func TestError_Unwrap(t *testing.T) {
	baseErr := errors.New("base error")
	wrapped := Wrap(baseErr, "context")

	if !errors.Is(wrapped, baseErr) {
		t.Error("Error_Unwrap() failed, wrapped error should contain base error")
	}
}

func TestStackLine(t *testing.T) {
	err := New("test error")
	ersErr := err.(*Error)

	if len(ersErr.Stack()) == 0 {
		t.Error("StackLine collection failed, stack should not be empty")
	}

	stack := ersErr.Stack()
	if !strings.Contains(stack[0].String(), ".go") {
		t.Errorf("StackLine String() = %v, should contain .go file extension", stack[0])
	}
}

func TestZeroValueStackLine(t *testing.T) {
	var s StackLine
	if s.String() != "(:0)" {
		t.Errorf("Zero value StackLine should format as (:0), got %s", s.String())
	}
	if s.File() != "" {
		t.Error("Zero value StackLine.File() should return empty string")
	}
	if s.Line() != 0 {
		t.Error("Zero value StackLine.Line() should return 0")
	}
}

func TestStackLine_FileAndLine(t *testing.T) {
	err := New("test")
	ersErr := err.(*Error)
	stack := ersErr.Stack()

	if stack[0].File() == "" {
		t.Error("StackLine File() should return non-empty string")
	}

	if stack[0].Line() <= 0 {
		t.Error("StackLine Line() should return positive line number")
	}
}

func TestError_AddContext(t *testing.T) {
	err := New("base")
	ersErr := err.(*Error)

	ersErr.AddContext("context1")
	if !strings.Contains(err.Error(), "context1") {
		t.Error("AddContext failed to add single context")
	}

	ersErr.AddContext("context2", "context3")
	errStr := err.Error()
	for _, ctx := range []string{"context2", "context3"} {
		if !strings.Contains(errStr, ctx) {
			t.Errorf("AddContext failed to add context: %s", ctx)
		}
	}
}

func TestDeepWrapping(t *testing.T) {
	base := New("level0")
	level1 := Wrap(base, "level1")
	level2 := Wrapf(level1, "level%d", 2)
	level3 := Wrap(level2, "level3", "extra")
	level4 := Wrapf(level3, "level%d-%s", 4, "final")

	result := level4.Error()
	expected := []string{"level0", "level1", "level2", "level3", "extra", "level4-final"}

	for _, exp := range expected {
		if !strings.Contains(result, exp) {
			t.Errorf("Deep wrapping failed, missing: %s", exp)
		}
	}

	ersErr := level4.(*Error)
	if len(ersErr.Stack()) != 5 {
		t.Errorf("Expected 5 stack frames, got %d", len(ersErr.Stack()))
	}
}

func TestStackTraceOrder(t *testing.T) {
	base := New("first")
	wrapped := Wrap(base, "second")
	ersErr := wrapped.(*Error)

	if len(ersErr.Stack()) != 2 {
		t.Fatal("Expected exactly 2 stack frames")
	}

	firstFrame := ersErr.Stack()[0].String()
	secondFrame := ersErr.Stack()[1].String()

	if firstFrame == secondFrame {
		t.Error("Stack frames should be different")
	}
}

func TestNewStackLine(t *testing.T) {
	stack := NewStackLine()

	if stack.File() == "" {
		t.Error("NewStackLine should return non-empty file")
	}
	if stack.Line() <= 0 {
		t.Error("NewStackLine should return positive line number")
	}
	if !strings.Contains(stack.String(), ".go:") {
		t.Errorf("NewStackLine string format should contain .go:, got %s", stack.String())
	}
}

func TestError_Contexts(t *testing.T) {
	err := New("base")
	ersErr := err.(*Error)

	// Test empty contexts
	if len(ersErr.Contexts()) != 0 {
		t.Error("New error should have empty contexts")
	}

	// Test single context
	ersErr.AddContext("context1")
	contexts := ersErr.Contexts()
	if len(contexts) != 1 || contexts[0] != "context1" {
		t.Error("Contexts() not returning correct single context")
	}

	// Test multiple contexts
	ersErr.AddContext("context2", "context3")
	contexts = ersErr.Contexts()
	if len(contexts) != 3 {
		t.Error("Contexts() not returning all contexts")
	}
	expected := []string{"context1", "context2", "context3"}
	for i, ctx := range expected {
		if contexts[i] != ctx {
			t.Errorf("Context at position %d expected %s, got %s", i, ctx, contexts[i])
		}
	}
}

func TestGetCallerUnknown(t *testing.T) {
	// Call with an impossibly high skip value to force failure
	stack := getCaller(9999)

	if stack.File() != "unknown" {
		t.Errorf("Expected unknown file for invalid caller, got %s", stack.File())
	}
	if stack.Line() != 0 {
		t.Errorf("Expected 0 line for invalid caller, got %d", stack.Line())
	}
	if stack.String() != "(unknown:0)" {
		t.Errorf("Expected (unknown:0) string format, got %s", stack.String())
	}
}
