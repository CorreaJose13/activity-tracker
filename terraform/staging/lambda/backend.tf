terraform {
  backend "s3" {
    bucket = "terraform-tgbot-lambda"
    key    = "state/terraform.tfstate"
    region = var.region
  }
}
