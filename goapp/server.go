package goapp

import (
	"net/http"
	"os"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/handlers"
)

type ServeOpts struct {
	Addr    string
	Handler http.Handler
}

// Serve serves the application. The parameter should be the handler for the app.
func Serve(opts *ServeOpts) {
	h := opts.Handler

	// When deployed, force TLS
	// if hcruntime.Nomad() {
	// 	log.Printf("[INFO] www: Nomad detected")
	// 	h = hchandlers.ForceTLS(h)
	// }

	// Record requests
	// h = hchandlers.Metrics(h)

	// Compress
	h = handlers.CompressHandler(h)

	// CORS
	cors := handlers.CORS(
		handlers.AllowedMethods([]string{"GET", "HEAD", "PUT", "POST"}),
		handlers.AllowedHeaders([]string{"Content-Type"}))
	// handlers.AllowedHeaders([]string{"Content-Type"}),
	// handlers.AllowCredentials())
	h = cors(h)

	// Create the server
	server := &http.Server{
		Addr:         opts.Addr,
		Handler:      handlers.CombinedLoggingHandler(os.Stderr, h),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// Kick it off
	log.Printf("[INFO] api: api server listening on %s", opts.Addr)
	server.ListenAndServe()
}
