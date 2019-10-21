data "pokemon_name" "pokemon" {}

output "name" {
  value = "${data.pokemon_name.pokemon.name}"
}
