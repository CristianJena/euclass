// Package web contains a small web framework extension.
package web

import (
	"github.com/dimfeld/httptreemux/v5"
)

// App is the entrypoint into our application and what configures our context
// object for each of our http handlers. Feel free to add any configuration
// data/logic on this App struct.
type App struct {
	*httptreemux.ContextMux
}

// NewApp creates an App value that handle a set of routes for the application.
func NewApp() *App {
	return &App{
		ContextMux: httptreemux.NewContextMux(),
	}
}
