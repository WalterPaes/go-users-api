package errors

import (
	"errors"
	"testing"
)

func TestCustomError_Error(t *testing.T) {
	type fields struct {
		err error
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Should create a custom error",
			fields: fields{
				err: errors.New("error"),
			},
			want: "error",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := NewCustomError(tt.fields.err)
			if got := e.Error(); got != tt.want {
				t.Errorf("CustomError.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}
