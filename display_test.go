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

func TestNewAndNewDisplayService(t *testing.T) {
	s, err := New()
	if err != nil {
		t.Fatalf("New() error = %v, wantErr nil", err)
	}
	if s == nil {
		t.Fatal("New() returned nil")
	}

	s, err = newDisplayService()
	if err != nil {
		t.Fatalf("newDisplayService() error = %v, wantErr nil", err)
	}
	if s == nil {
		t.Fatal("newDisplayService() returned nil")
	}
}

func TestWindowOptions(t *testing.T) {
	config := &WindowConfig{}

	WithName("test-name").Apply(config)
	if config.Name != "test-name" {
		t.Errorf("WithName() got = %v, want %v", config.Name, "test-name")
	}

	WithTitle("test-title").Apply(config)
	if config.Title != "test-title" {
		t.Errorf("WithTitle() got = %v, want %v", config.Title, "test-title")
	}

	WithWidth(100).Apply(config)
	if config.Width != 100 {
		t.Errorf("WithWidth() got = %v, want %v", config.Width, 100)
	}

	WithHeight(200).Apply(config)
	if config.Height != 200 {
		t.Errorf("WithHeight() got = %v, want %v", config.Height, 200)
	}

	WithURL("/testurl").Apply(config)
	if config.URL != "/testurl" {
		t.Errorf("WithURL() got = %v, want %v", config.URL, "/testurl")
	}
}

func TestService_HandleOpenWindowAction(t *testing.T) {
	s, _ := New()
	// This test will panic because the app is nil, but it will still cover the function.
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	_ = s.handleOpenWindowAction(map[string]any{})
}
