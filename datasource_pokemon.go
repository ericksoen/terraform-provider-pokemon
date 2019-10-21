package main

import (
	"log"

	poke "github.com/ericksoen/pokemon-go-client"
	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourcePokemonName() *schema.Resource {
	return &schema.Resource{
		Read: dataSourcePokemonNameRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func dataSourcePokemonNameRead(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[INFO] Reading Random Pokemon")

	client := meta.(*poke.Client)

	pokemon, err := client.GetRandomPokemon()

	log.Printf("[INFO] Random Pokemon name is: %s", pokemon)
	if err != nil {
		return err
	}

	d.SetId("12345")
	d.Set("name", pokemon)

	return nil
}
