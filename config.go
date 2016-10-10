package main

import (
	"encoding/hex"
	"fmt"
	"os"

	log "github.com/Sirupsen/logrus"
)

// Config is the configuration we load from Consul
type Config struct {
	DBAddr               string
	GHAuthEmail          string
	GHAuthToken          string
	SessionSecretKey     string
	SessionEncryptionKey string
}

func (c *Config) SessionSecretBytes() []byte {
	v, err := hex.DecodeString(c.SessionSecretKey)
	if err != nil {
		panic(err)
	}

	return v
}

func (c *Config) SessionEncryptionBytes() []byte {
	v, err := hex.DecodeString(c.SessionEncryptionKey)
	if err != nil {
		panic(err)
	}

	return v
}

// initConfig initializes the configuration, returning our initial
// configuration from default values in the environment
func initConfig() (*Config, error) {
	ghEmail := os.Getenv("GH_EMAIL")
	if ghEmail == "" {
		log.Fatalf("GitHub user email is empty")
	}

	ghToken := os.Getenv("GH_TOKEN")
	if ghToken == "" {
		log.Fatalf("GitHub auth token is empty")
	}

	if ghEmail != "" || ghToken != "" {
		return &Config{
			GHAuthEmail: ghEmail,
			GHAuthToken: ghToken,
		}, nil
	}
	return nil, fmt.Errorf("[ERR] No config loaded")
}
