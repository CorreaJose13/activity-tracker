terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.46.0"
    }
    archive = {
      source  = "hashicorp/archive"
      version = "~> 2.4.2"
    }
    null = {
      source  = "hashicorp/null"
      version = "~> 3.2.2"
    }
  }
  backend "s3" {
    bucket = "terraform-tg-lambda"
    key    = "dev/terraform.tfstate"
    region = "us-east-1"
  }
}

provider "aws" {
  region = "us-east-1"
}

//Mejorar el manejo de estos secrets luego
data "aws_secretsmanager_secret" "bot_token" {
  name = "BOT_TOKEN"
}

data "aws_secretsmanager_secret" "mongo_token" {
  name = "MONGO_TOKEN"
}

data "aws_secretsmanager_secret_version" "bot_token_version" {
  secret_id = data.aws_secretsmanager_secret.bot_token.id
}

data "aws_secretsmanager_secret_version" "mongo_token_version" {
  secret_id = data.aws_secretsmanager_secret.mongo_token.id
}
