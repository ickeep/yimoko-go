package utils

import "testing"

func TestFirstUpper(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want string
	}{
		{
			name: "empty string",
			arg:  "",
			want: "",
		},
		{
			name: "single character",
			arg:  "a",
			want: "A",
		},
		{
			name: "multiple characters",
			arg:  "hello",
			want: "Hello",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FirstUpper(tt.arg); got != tt.want {
				t.Errorf("FirstUpper() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFirstLower(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want string
	}{
		{
			name: "empty string",
			arg:  "",
			want: "",
		},
		{
			name: "single character",
			arg:  "A",
			want: "a",
		},
		{
			name: "multiple characters",
			arg:  "HELLO",
			want: "hELLO",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FirstLower(tt.arg); got != tt.want {
				t.Errorf("FirstLower() = %v, want %v", got, tt.want)
			}
		})
	}
}
