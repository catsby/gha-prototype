package goapp

import (
	"fmt"
	"net/http"
	"testing"
)

func TestHandlerHandleCheck(t *testing.T) {
	// Create the handler on our own so we can override the test flag
	ln := TestServer(t, Handler(&HandlerOpts{
		Test: false,
	}))
	defer ln.Close()

	// Make the request
	resp, err := http.Get(fmt.Sprintf("http://%s/_check", ln.Addr()))
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	if resp.StatusCode != 204 {
		t.Fatalf("bad: %d", resp.StatusCode)
	}
	if resp.Body != nil {
		resp.Body.Close()
	}
}

func TestHandlerHandleCheck_testMode(t *testing.T) {
	// Create the handler on our own so we can override the test flag
	ln := TestServer(t, Handler(&HandlerOpts{
		Test: true,
	}))
	defer ln.Close()

	// Make the request
	resp, err := http.Get(fmt.Sprintf("http://%s/_check", ln.Addr()))
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	if resp.StatusCode != 500 {
		t.Fatalf("bad: %d", resp.StatusCode)
	}
	if resp.Body != nil {
		resp.Body.Close()
	}
}
