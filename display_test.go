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
						opts.AlwaysOnTop = true
						opts.Hidden = true
						opts.MinimiseButtonState = application.ButtonHidden
						opts.MaximiseButtonState = application.ButtonDisabled
						opts.CloseButtonState = application.ButtonEnabled
						opts.Frameless = true
					},
				},
			},
			want: application.WebviewWindowOptions{
				Name:                "test",
				Title:               "Test Window",
				Width:               1920,
				Height:              1080,
				URL:                 "/test",
				AlwaysOnTop:         true,
				Hidden:              true,
				MinimiseButtonState: application.ButtonHidden,
				MaximiseButtonState: application.ButtonDisabled,
				CloseButtonState:    application.ButtonEnabled,
				Frameless:           true,
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

	WithAlwaysOnTop(true).Apply(config)
	if !config.AlwaysOnTop {
		t.Errorf("WithAlwaysOnTop() got = %v, want %v", config.AlwaysOnTop, true)
	}

	WithHidden(true).Apply(config)
	if !config.Hidden {
		t.Errorf("WithHidden() got = %v, want %v", config.Hidden, true)
	}

	WithMinimiseButtonState(application.ButtonHidden).Apply(config)
	if config.MinimiseButtonState != application.ButtonHidden {
		t.Errorf("WithMinimiseButtonState() got = %v, want %v", config.MinimiseButtonState, application.ButtonHidden)
	}

	WithMaximiseButtonState(application.ButtonDisabled).Apply(config)
	if config.MaximiseButtonState != application.ButtonDisabled {
		t.Errorf("WithMaximiseButtonState() got = %v, want %v", config.MaximiseButtonState, application.ButtonDisabled)
	}

	WithCloseButtonState(application.ButtonEnabled).Apply(config)
	if config.CloseButtonState != application.ButtonEnabled {
		t.Errorf("WithCloseButtonState() got = %v, want %v", config.CloseButtonState, application.ButtonEnabled)
	}

	WithFrameless(true).Apply(config)
	if !config.Frameless {
		t.Errorf("WithFrameless() got = %v, want %v", config.Frameless, true)
	}
}

func TestService_HandleOpenWindowAction(t *testing.T) {
	t.Skip("Skipping test that requires a running Wails application.")
	s, _ := New()
	_ = s.handleOpenWindowAction(map[string]any{})
}

func TestService_ShowEnvironmentDialog(t *testing.T) {
	t.Skip("Skipping test that requires a running Wails application.")
	s, _ := New()
	s.ShowEnvironmentDialog()
}

func TestService_OpenWindow(t *testing.T) {
	t.Skip("Skipping test that requires a running Wails application.")
	s, _ := New()
	_ = s.OpenWindow()
}

func TestService_MonitorScreenChanges(t *testing.T) {
	t.Skip("Skipping test that requires a running Wails application.")
	s, _ := New()
	s.monitorScreenChanges()
}

func TestService_BuildMenu(t *testing.T) {
	t.Skip("Skipping test that requires a running Wails application.")
	s, _ := New()
	s.buildMenu()
}

func TestService_SystemTray(t *testing.T) {
	t.Skip("Skipping test that requires a running Wails application.")
	s, _ := New()
	s.systemTray()
}

func TestService_Startup(t *testing.T) {
	t.Skip("Skipping test that requires a running Wails application.")
	s, _ := New()
	_ = s.Startup(nil)
}
