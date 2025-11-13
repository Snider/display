package display

import (
	"reflect"
	"testing"

	"github.com/wailsapp/wails/v3/pkg/application"
)

func TestParseWindowOptions(t *testing.T) {
	tests := []struct {
		name string
		msg  map[string]any
		want application.WebviewWindowOptions
	}{
		{
			name: "Valid options",
			msg: map[string]any{
				"name": "main",
				"options": map[string]any{
					"Title":  "My App",
					"Width":  1024.0,
					"Height": 768.0,
				},
			},
			want: application.WebviewWindowOptions{
				Name:   "main",
				Title:  "My App",
				Width:  1024,
				Height: 768,
			},
		},
		{
			name: "Missing options",
			msg: map[string]any{
				"name": "main",
			},
			want: application.WebviewWindowOptions{
				Name: "main",
			},
		},
		{
			name: "Empty message",
			msg:  map[string]any{},
			want: application.WebviewWindowOptions{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseWindowOptions(tt.msg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseWindowOptions() = %v, want %v", got, tt.want)
			}
		})
	}
}

// mockWindowOption is a mock implementation of the WindowOption interface for testing.
type mockWindowOption struct {
	applyFunc func(*WindowConfig)
}

func (m *mockWindowOption) Apply(opts *WindowConfig) {
	m.applyFunc(opts)
}

func TestBuildWailsWindowOptions(t *testing.T) {
	tests := []struct {
		name string
		opts []WindowOption
		want application.WebviewWindowOptions
	}{
		{
			name: "Default options",
			opts: []WindowOption{},
			want: application.WebviewWindowOptions{
				Name:   "main",
				Title:  "Core",
				Width:  1280,
				Height: 800,
				URL:    "/",
			},
		},
		{
			name: "Override options",
			opts: []WindowOption{
				&mockWindowOption{
					applyFunc: func(opts *WindowConfig) {
						opts.Name = "test"
						opts.Title = "Test Window"
						opts.Width = 1920
						opts.Height = 1080
						opts.URL = "/test"
					},
				},
			},
			want: application.WebviewWindowOptions{
				Name:   "test",
				Title:  "Test Window",
				Width:  1920,
				Height: 1080,
				URL:    "/test",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildWailsWindowOptions(tt.opts...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildWailsWindowOptions() = %v, want %v", got, tt.want)
			}
		})
	}
}
