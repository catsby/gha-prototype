package goapp

import (
	"errors"
	"fmt"
	"net/http"
	"testing"
)

func TestHandler_root(t *testing.T) {
	ln := TestServer(t, nil)
	defer ln.Close()

	// fn is the CheckRedirect function
	var success bool
	fn := func(req *http.Request, via []*http.Request) error {
		success = true
		return errors.New("")
	}

	// Client
	client := &http.Client{CheckRedirect: fn}

	// Make the request
	client.Get(fmt.Sprintf("http://%s/", ln.Addr()))
	if !success {
		t.Fatal("should've redirected")
	}
}
