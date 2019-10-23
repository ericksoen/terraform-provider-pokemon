package main

import (
	"log"
	"strconv"

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

	log.Printf("[INFO] Random Pokemon name is: %s", pokemon.Name)

	if err != nil {
		return err
	}

	d.SetId(strconv.Itoa(pokemon.ID))
	d.Set("name", pokemon.Name)

	return nil
}
