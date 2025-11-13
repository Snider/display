package display

// WindowConfig holds the configuration for a window.
type WindowConfig struct {
	Name   string
	Title  string
	Width  int
	Height int
	URL    string
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
