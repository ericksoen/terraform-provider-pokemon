package main

import (
	"log"

	poke "github.com/ericksoen/pokemon-go-client"
)

// Config defines the values needed to create
// the PokemonGo Client
type Config struct {
	APIKey string
}

// Client creates an instance of the PokemonGo client
// using the API Key from the provider config
func (c *Config) Client() (*poke.Client, error) {
	config := poke.Config{
		APIKey: c.APIKey,
	}

	client := poke.NewClient(config)

	log.Printf("[INFO] PokemonGoClient configured")

	return &client, nil

}
