package goapp

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/hashicorp/errwrap"
)

type HandlerOpts struct {
	// DB *gorm.DB        // DB is the DB instance
	// ES *elastic.Client // ES is the ElasticSearch client

	// Client ID and secret for the GitHub authentication
	GHAuthEmail string
	GHAuthToken string

	// Session is the session store
	Session sessions.Store

	// Test is true if we're in test mode. Be careful to set this since
	// it will break a lot of things.
	Test bool
}

// Handler returns the http.Handler implementation for the HTTP API.
func Handler(opts *HandlerOpts) http.Handler {
	// Build our router
	r := mux.NewRouter()

	// Stable
	// r.Handle("/", http.RedirectHandler(
	// 	"https://hearthcloud.com", http.StatusFound))
	r.Handle("/_check", handleCheck(opts))

	// V1 API
	r.Handle("/v1/issues", handleIssues(opts))
	// r.Handle("/v1/games/latest", handleGamesLatest(opts))
	// r.Handle("/v1/games/win-rates", handleGamesWinRates(opts))
	// r.Handle("/v1/games/{gameId}", handleGamesSingle(opts))
	// r.Handle("/v1/games/{gameId}/events", handleGameEvents(opts))
	// r.Handle("/v1/search", handleSearch(opts))
	// r.Handle("/v1/user/resolve/{userId}", handleUserResolve(opts))
	// r.Handle("/v1/account/complete", handleAccountComplete(opts))
	// r.Handle("/v1/auth/logout", handleAuthLogout(opts))
	// r.Handle("/v1/auth/status", handleAuthStatus(opts))
	// r.Handle("/v1/auth-bnet", handleAuthGH(opts))
	// r.Handle("/v1/auth-bnet/callback", handleAuthGHCallback(opts))

	return r
}

func respondOk(w http.ResponseWriter, body interface{}) {
	w.Header().Add("Content-Type", "application/json")

	if body == nil {
		w.WriteHeader(http.StatusNoContent)
	} else {
		w.WriteHeader(http.StatusOK)
		enc := json.NewEncoder(w)
		enc.Encode(body)
	}
}

func respondError(w http.ResponseWriter, status int, err error) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	resp := &ErrorResponse{Errors: make([]string, 0, 1)}
	if err != nil {
		if w, ok := err.(errwrap.Wrapper); ok {
			for _, err := range w.WrappedErrors() {
				resp.Errors = append(resp.Errors, err.Error())
			}
		} else {
			resp.Errors = []string{err.Error()}
		}
	}

	enc := json.NewEncoder(w)
	enc.Encode(resp)
}

// ErrorResponse is the structure for error values.
type ErrorResponse struct {
	Errors []string `json:"errors"`
}
