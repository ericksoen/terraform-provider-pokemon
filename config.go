package main

import (
	"log"

	poke "github.com/ericksoen/pokemon-go-client"
)

type Config struct {
	APIKey string
}

func (c *Config) Client() (*poke.Client, error) {
	config := poke.Config{
		APIKey: c.APIKey,
	}

	client := poke.NewClient(config)

	log.Printf("[INFO] PokemonGoClient configured")

	return &client, nil

}
