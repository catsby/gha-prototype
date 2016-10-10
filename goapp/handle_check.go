package goapp

import (
	"net/http"
)

func handleCheck(opts *HandlerOpts) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// If we're in test mode, we always fail the health check so that
		// we NEVER enter production traffic.
		if opts.Test {
			respondError(w, 500, nil)
			return
		}

		// In the future we might do fancier checks, but for now we just
		// respond we're okay which is a good check that the HTTP server
		// is up and running.
		respondOk(w, nil)
	})
}
