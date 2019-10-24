## Overview

Create a custom Terraform provider and register a data source that returns a random Pokemon character from the [RapidAPI PokemonGo API](https://rapidapi.com/brianiswu/api/pokemon-go1?endpoint=apiendpoint_af8fc313-bc8c-4b86-8671-83a2212bdefc). The PokemonGo API requires an API key (free) to use, so you'll need to sign up for one before you can use the client.

## Implementation

Create a `Provider` that registers the provider `Schema` and creates an instance of the [3rd party PokemonGo client](https://github.com/ericksoen/pokemon-go-client). The provider also registers a single data source that retrieves a random Pokemon character and stores it in the state file.

## Build the provider

To build the provider for local testing, run `go build -o terraform-provider-pokemon` (you may need to run `go build -o terraform-provider-pokemon.exe` on Windows machines).

## Use the provider

Run `terraform init` to confirm you can correctly initialize your custom provider.

Create a new `*.tf` file in the same folder as the your custom provider code since Terraform provider discovery will scan the current working directory. Add the following code snippet to the new file:

```hcl
data "pokemon_name" "pokemon" {}

output "name" {
  value = "${data.pokemon_name.pokemon.name}"
}
```

The custom provider schema requires an API key to operate correctly&mdash;you can either inject this as an environment variable `export POKE_API_KEY=value` (or the equivalent for your operating system shell) or supply it when prompted during the next step.

You should now be able to use our custom provider to generate a random Pokemon character.

```bash
$ terraform apply
> data.pokemon_name.pokemon: Refreshing state...
> Apply complete! Resources: 0 added, 0 changed, 0 destroyed.

> Outputs:
> name = Pikachu
```

You can also run `cat terraform.tfstate` to confirm that the output has correctly updated our state file.

## Additional Reading

The [Terraform developer docs](https://www.terraform.io/docs/extend/writing-custom-providers.html) provide some clear instructions on how to write and test custom providers.
