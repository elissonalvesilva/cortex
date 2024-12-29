package utils

import "testing"

func TestIsTemplateVariable(t *testing.T) {
	tests := []struct {
		name string
		args string
		want bool
	}{
		{
			name: "is template variable",
			args: "{{.variable}}",
			want: true,
		},
		{
			name: "is not template variable",
			args: "variable",
			want: false,
		},
		{
			name: "is not template variable",
			args: "{{variable}}",
			want: false,
		},
		{
			name: "is template variable",
			args: "{{.v1.v2}}",
			want: true,
		},
		{
			name: "is template variable",
			args: "{{.v1.v2.v3}}",
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsTemplateVariable(tt.args); got != tt.want {
				t.Errorf("IsTemplateVariable() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTemplateVariable(t *testing.T) {
	tests := []struct {
		name string
		args string
		want string
	}{
		{
			name: "template variable",
			args: "{{.variable}}",
			want: "variable",
		},
		{
			name: "template variable",
			args: "{{.v1.v2}}",
			want: "v1.v2",
		},
		{
			name: "template variable",
			args: "{{.v1.v2.v3}}",
			want: "v1.v2.v3",
		},
		{
			name: "template variable",
			args: "{{.v1.v2.v3.v4}}",
			want: "v1.v2.v3.v4",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TemplateVariable(tt.args); got != tt.want {
				t.Errorf("TemplateVariable() = %v, want %v", got, tt.want)
			}
		})
	}
}
