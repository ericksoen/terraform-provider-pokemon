data "pokemon_name" "pokemon" {}

resource "aws_s3_bucket" "b" {
  bucket_prefix = "${data.pokemon_name.pokemon.name}-"

  acl = "private"
}

provider "aws" {
  region = "us-east-1"
}

output "name" {
  value = "${aws_s3_bucket.b.name.id}"
}
