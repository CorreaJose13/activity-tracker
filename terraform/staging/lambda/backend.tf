terraform {
  backend "s3" {
    bucket = "terraform-tgbot-lambda"
    key    = "dev/terraform.tfstate"
    region = "us-east-1"
  }
}
