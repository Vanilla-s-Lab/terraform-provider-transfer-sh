terraform {
  required_providers {
    transfer-sh = {
      source = "hashicorp.com/edu/transfer-sh"
    }
  }
}

provider "transfer-sh" {}

resource "transfer-sh_file" "test" {
    file_path = "./hello.txt"
}
