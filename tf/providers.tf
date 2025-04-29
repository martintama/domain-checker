provider "aws" {
  region  = var.aws_region
  profile = "tf-admin"

  default_tags {
    tags = {
      Application = "domain-checker"
      ManagedBy   = "terraform"
    }
  }
}
