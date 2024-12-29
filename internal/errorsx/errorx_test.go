package errorsx

import (
	"errors"
	"testing"
)

func TestCortexError_Error(t *testing.T) {
	tests := []struct {
		name    string
		err     error
		message []string
		want    string
	}{
		{
			name:    "simple error",
			err:     errors.New("original error"),
			message: []string{"Something went wrong"},
			want:    "cortex error: original error\nmessage: Something went wrong",
		},
		{
			name:    "multiple messages",
			err:     errors.New("original error"),
			message: []string{"Step 1 failed", "Step 2 failed"},
			want:    "cortex error: original error\nmessage: Step 1 failed\nStep 2 failed",
		},
		{
			name:    "empty message",
			err:     errors.New("original error"),
			message: []string{},
			want:    "cortex error: original error\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := &CortexError{
				message: tt.message,
				err:     tt.err,
			}

			got := err.Error()
			if got != tt.want {
				t.Errorf("CortexError.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWrap(t *testing.T) {
	tests := []struct {
		name    string
		err     error
		message []string
		want    string
	}{
		{
			name:    "wrap simple error",
			err:     errors.New("original error"),
			message: []string{"Additional context"},
			want:    "cortex error: original error\nmessage: Additional context",
		},
		{
			name:    "wrap already CortexError",
			err:     &CortexError{message: []string{"Context 1"}, err: errors.New("wrapped error")},
			message: []string{"Context 2"},
			want:    "cortex error: wrapped error\nmessage: Context 1\nContext 2",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Wrap(tt.err, tt.message...)
			if got.Error() != tt.want {
				t.Errorf("Wrap() = %v, want %v", got.Error(), tt.want)
			}
		})
	}
}

func TestWrapNilError(t *testing.T) {
	err := Wrap(nil, "Nil error test")
	if err != nil {
		t.Errorf("Expected nil error, got: %v", err)
	}
}
