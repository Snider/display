package display

import (
	"context"
	"fmt"

	"github.com/wailsapp/wails/v3/pkg/application"
	"github.com/wailsapp/wails/v3/pkg/events"
)

// Options holds configuration for the display service.
type Options struct{}

// Service manages windowing, dialogs, and other visual elements.
type Service struct {
	app    *application.App
	config Options
}

// newDisplayService contains the common logic for initializing a Service struct.
func newDisplayService() (*Service, error) {
	return &Service{}, nil
}

// New is the constructor for static dependency injection.
func New() (*Service, error) {
	s, err := newDisplayService()
	if err != nil {
		return nil, err
	}
	return s, nil
}

// Startup initializes the display service and sets up the main application window and system tray.
func (s *Service) Startup(ctx context.Context) error {
	s.app = application.Get()
	s.app.Logger.Info("Display service started")
	s.buildMenu()
	s.systemTray()
	return s.OpenWindow()
}

// handleOpenWindowAction processes a message to configure and create a new window using specified name and options.
func (s *Service) handleOpenWindowAction(msg map[string]any) error {
	opts := parseWindowOptions(msg)
	s.app.Window.NewWithOptions(opts)
	return nil
}

// parseWindowOptions extracts window configuration from a map and returns it as WebviewWindowOptions.
func parseWindowOptions(msg map[string]any) application.WebviewWindowOptions {
	opts := application.WebviewWindowOptions{}
	if name, ok := msg["name"].(string); ok {
		opts.Name = name
	}
	if optsMap, ok := msg["options"].(map[string]any); ok {
		if title, ok := optsMap["Title"].(string); ok {
			opts.Title = title
		}
		if width, ok := optsMap["Width"].(float64); ok {
			opts.Width = int(width)
		}
		if height, ok := optsMap["Height"].(float64); ok {
			opts.Height = int(height)
		}
	}
	return opts
}

// ShowEnvironmentDialog displays a dialog containing detailed information about the application's runtime environment.
func (s *Service) ShowEnvironmentDialog() {
	envInfo := s.app.Env.Info()

	details := "Environment Information:\n\n"
	details += fmt.Sprintf("Operating System: %s\n", envInfo.OS)
	details += fmt.Sprintf("Architecture: %s\n", envInfo.Arch)
	details += fmt.Sprintf("Debug Mode: %t\n\n", envInfo.Debug)
	details += fmt.Sprintf("Dark Mode: %t\n\n", s.app.Env.IsDarkMode())
	details += "Platform Information:"

	// Add platform-specific details
	for key, value := range envInfo.PlatformInfo {
		details += fmt.Sprintf("\n%s: %v", key, value)
	}

	if envInfo.OSInfo != nil {
		details += fmt.Sprintf("\n\nOS Details:\nName: %s\nVersion: %s",
			envInfo.OSInfo.Name,
			envInfo.OSInfo.Version)
	}

	dialog := s.app.Dialog.Info()
	dialog.SetTitle("Environment Information")
	dialog.SetMessage(details)
	dialog.Show()
}

// OpenWindow creates a new window with the default options.
func (s *Service) OpenWindow(opts ...WindowOption) error {
	wailsOpts := buildWailsWindowOptions(opts...)
	s.app.Window.NewWithOptions(wailsOpts)
	return nil
}

// buildWailsWindowOptions creates Wails window options from core window options.
func buildWailsWindowOptions(opts ...WindowOption) application.WebviewWindowOptions {
	// Default options
	winOpts := &WindowConfig{
		Name:   "main",
		Title:  "Core",
		Width:  1280,
		Height: 800,
		URL:    "/",
	}

	// Apply options
	for _, opt := range opts {
		opt.Apply(winOpts)
	}

	// Create Wails window options
	return application.WebviewWindowOptions{
		Name:   winOpts.Name,
		Title:  winOpts.Title,
		Width:  winOpts.Width,
		Height: winOpts.Height,
		URL:    winOpts.URL,
	}
}

// monitorScreenChanges listens for theme change events and logs when screen configuration changes occur.
func (s *Service) monitorScreenChanges() {
	s.app.Event.OnApplicationEvent(events.Common.ThemeChanged, func(event *application.ApplicationEvent) {
		s.app.Logger.Info("Screen configuration changed")
	})
}
