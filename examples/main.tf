terraform {
  required_providers {
    cause = {
      version = "1.0"
      source  = "github.com/puni9869/cause"
    }
  }
}


provider "cause" {}

resource "cause" "create" {
  number = "10"
}
