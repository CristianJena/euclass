// Package handlers manages the different versions of the API.
package handlers

import (
	"expvar"
	"net/http"
	"net/http/pprof"

	"github.com/ardanlabs/service/app/services/sales-api/handlers/testgrp"
	"github.com/ardanlabs/service/business/web/mid"
	"github.com/ardanlabs/service/foundation/web"
	"go.uber.org/zap"
)

// APIMux constructs a http.Handler with all application routes defined.
func APIMux(log *zap.SugaredLogger) *web.App {
	app := web.NewApp(mid.Logger(log))

	app.Handle(http.MethodGet, "/test", testgrp.Handler)

	return app
}

// DebugStandardLibraryMux registers all the debug routes from the standard library
// into a new mux bypassing the use of the DefaultServerMux. Using the
// DefaultServerMux would be a security risk since a dependency could inject a
// handler into our service without us knowing it.
func DebugStandardLibraryMux() *http.ServeMux {
	mux := http.NewServeMux()

	// Register all the standard library debug endpoints.
	mux.HandleFunc("/debug/pprof/", pprof.Index)
	mux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	mux.HandleFunc("/debug/pprof/trace", pprof.Trace)
	mux.Handle("/debug/vars", expvar.Handler())

	return mux
}
