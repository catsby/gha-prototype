package main

import (
	"os"

	log "github.com/Sirupsen/logrus"

	"github.com/catsby/gha-prototype/goapp"
	"github.com/gorilla/sessions"
)

func main() {
	// Grab the port
	addr := ""
	port := os.Getenv("PORT")
	if port == "" {
		log.Printf("[ERR] PORT is empty, using default 8001")
		port = "8001"
	}
	addr = ":" + port

	config, err := initConfig()
	if err != nil {
		log.Fatalf("Error loading config: %s", err)
	}

	// Create the app handler
	goapp.Serve(&goapp.ServeOpts{
		Addr: addr,
		Handler: goapp.Handler(&goapp.HandlerOpts{
			GHAuthEmail: config.GHAuthEmail,
			GHAuthToken: config.GHAuthToken,

			Session: sessions.NewCookieStore(
				config.SessionSecretBytes(),
				config.SessionEncryptionBytes()),
		}),
	})
}
