package display

import "github.com/wailsapp/wails/v3/pkg/application"

// WindowConfig holds the configuration for a window.
type WindowConfig struct {
	Name                string
	Title               string
	Width               int
	Height              int
	URL                 string
	AlwaysOnTop         bool
	Hidden              bool
	MinimiseButtonState application.ButtonState
	MaximiseButtonState application.ButtonState
	CloseButtonState    application.ButtonState
	Frameless           bool
}

// WindowOption is a function that applies a configuration option to a WindowConfig.
type WindowOption interface {
	Apply(*WindowConfig)
}

// WindowOptionFunc is a function that implements the WindowOption interface.
type WindowOptionFunc func(*WindowConfig)

func (f WindowOptionFunc) Apply(c *WindowConfig) {
	f(c)
}

func WithName(name string) WindowOption {
	return WindowOptionFunc(func(c *WindowConfig) {
		c.Name = name
	})
}

func WithTitle(title string) WindowOption {
	return WindowOptionFunc(func(c *WindowConfig) {
		c.Title = title
	})
}

func WithWidth(width int) WindowOption {
	return WindowOptionFunc(func(c *WindowConfig) {
		c.Width = width
	})
}

func WithHeight(height int) WindowOption {
	return WindowOptionFunc(func(c *WindowConfig) {
		c.Height = height
	})
}

func WithURL(url string) WindowOption {
	return WindowOptionFunc(func(c *WindowConfig) {
		c.URL = url
	})
}

func WithAlwaysOnTop(alwaysOnTop bool) WindowOption {
	return WindowOptionFunc(func(c *WindowConfig) {
		c.AlwaysOnTop = alwaysOnTop
	})
}

func WithHidden(hidden bool) WindowOption {
	return WindowOptionFunc(func(c *WindowConfig) {
		c.Hidden = hidden
	})
}

func WithMinimiseButtonState(state application.ButtonState) WindowOption {
	return WindowOptionFunc(func(c *WindowConfig) {
		c.MinimiseButtonState = state
	})
}

func WithMaximiseButtonState(state application.ButtonState) WindowOption {
	return WindowOptionFunc(func(c *WindowConfig) {
		c.MaximiseButtonState = state
	})
}

func WithCloseButtonState(state application.ButtonState) WindowOption {
	return WindowOptionFunc(func(c *WindowConfig) {
		c.CloseButtonState = state
	})
}

func WithFrameless(frameless bool) WindowOption {
	return WindowOptionFunc(func(c *WindowConfig) {
		c.Frameless = frameless
	})
}
