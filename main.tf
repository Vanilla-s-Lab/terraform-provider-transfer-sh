terraform {
  required_providers {
    transfer-sh = {
      source = "hashicorp.com/edu/transfer-sh"
    }
  }
}

provider "transfer-sh" {}

resource "transfer-sh_file" "hello" {
  file_path = "./hello.txt"
}

output "file_hash" {
  // noinspection HILUnresolvedReference
  value = transfer-sh_file.hello.file_hash
}

output "link" {
  // noinspection HILUnresolvedReference
  value = transfer-sh_file.hello.link
}
