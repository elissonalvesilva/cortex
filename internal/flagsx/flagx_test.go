package flagsx

import (
	"testing"

	"github.com/elissonalvesilva/cortex/utils"
)

func TestConfigLoader_GetString(t *testing.T) {
	tests := []struct {
		name    string
		key     string
		want    string
		wantErr bool
	}{
		{
			name:    "existing key",
			key:     "cortex.short",
			want:    "Cortex CLI tool for search operations",
			wantErr: false,
		},
		{
			name:    "non-existing key",
			key:     "cortex.nonexistent",
			want:    "",
			wantErr: false,
		},
		{
			name:    "nested key",
			key:     "cortex.search.google.short",
			want:    "Google Search",
			wantErr: false,
		},
		{
			name:    "nested key with deeper structure",
			key:     "cortex.search.bing.long",
			want:    "Performs a search using Microsoft's Bing search engine. Similar to Google, you can use search terms to get results from Bing.",
			wantErr: false,
		},
	}

	flags := MustLoad()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := flags.loader.GetString(tt.key)
			if got != tt.want {
				t.Errorf("ConfigLoader.GetString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlags_Get(t *testing.T) {
	tests := []struct {
		name    string
		key     string
		want    string
		wantErr bool
	}{
		{
			name:    "existing key",
			key:     "cortex.short",
			want:    "Cortex CLI tool for search operations",
			wantErr: false,
		},
		{
			name:    "non-existing key",
			key:     "nonexistent.key",
			want:    "",
			wantErr: false,
		},
	}

	flags := MustLoad()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := flags.Get(tt.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Flags.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Flags.Get() = %v, want %v", got, tt.want)
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
			name: "template variable with nested keys",
			args: "{{.cortex.search.google.short}}",
			want: "cortex.search.google.short",
		},
		{
			name: "template variable with deeper nested keys",
			args: "{{.cortex.search.bing.long}}",
			want: "cortex.search.bing.long",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := utils.TemplateVariable(tt.args)
			if got != tt.want {
				t.Errorf("TemplateVariable() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsTemplateVariable(t *testing.T) {
	tests := []struct {
		name string
		args string
		want bool
	}{
		{
			name: "is template variable",
			args: "{{.cortex.search.google.short}}",
			want: true,
		},
		{
			name: "not a template variable",
			args: "cortex.search.google.short",
			want: false,
		},
		{
			name: "invalid template format",
			args: "{{cortex.search.google.short}}",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := utils.IsTemplateVariable(tt.args)
			if got != tt.want {
				t.Errorf("IsTemplateVariable() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetTestData(t *testing.T) {
	testData := []byte(`{
		"cortex": {
			"short": "Cortex CLI tool for search operations",
			"long": "Cortex is a powerful command-line tool designed for performing various search operations, including searching through Google and Bing.",
			"search": {
				"google": {
					"use": "searches on Google",
					"short": "Google Search",
					"long": "Performs a search using Google search engine. You can specify search terms and get results from Google's search engine."
				},
				"bing": {
					"use": "searches on Bing",
					"short": "Bing Search",
					"long": "Performs a search using Microsoft's Bing search engine. Similar to Google, you can use search terms to get results from Bing."
				}
			}
		}
	}`)
	setTestData(t, testData)

	flags := MustLoad()

	got := flags.loader.GetString("cortex.search.google.short")
	expected := "Google Search"
	if got != expected {
		t.Errorf("Expected value = %v, got = %v", expected, got)
	}

	gotTemplate := flags.loader.GetString("cortex.short")
	expectedTemplate := "Cortex CLI tool for search operations"
	if gotTemplate != expectedTemplate {
		t.Errorf("Expected value = %v, got = %v", expectedTemplate, gotTemplate)
	}
}
