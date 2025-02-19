terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.75.0"
    }
    archive = {
      source  = "hashicorp/archive"
      version = "~> 2.6.0"
    }
    null = {
      source  = "hashicorp/null"
      version = "~> 3.2.3"
    }
  }
  backend "s3" {
    bucket         = "tf-state-activity-tracker-bot"
    key            = "state/terraform.tfstate"
    region         = "us-east-1"
    encrypt        = true
    dynamodb_table = "terraform_state_db"
  }
}

provider "aws" {
  region = "us-east-1"
}

locals {
  src_path     = "${path.module}/../../../functions/telegram/main.go"
  binary_path  = "${path.module}/tf_generated/bootstrap"
  archive_path = "${path.module}/tf_generated/lambda_function.zip"
}

module "iam_role" {
  source = "../../modules/iam_role/"

  role_name               = "iam_for_lambda"
  assume_role_identifiers = ["lambda.amazonaws.com"]
}

module "iam_policy_attachment_logs" {
  source = "../../modules/iam_policy_attachment/"

  policy_name = "logs_lambda_scheduler"
  action      = ["logs:CreateLogStream", "logs:PutLogEvents"]
  resource    = "arn:aws:logs:*:*:*"
  role_name   = module.iam_role.role_name
}

module "lambda_function" {
  source = "../../modules/services/lambda_function/"

  src_path             = local.src_path
  binary_path          = local.binary_path
  archive_path         = local.archive_path
  lambda_function_name = "activity-tracker-bot"
  role_arn             = module.iam_role.role_arn
  environment_variables = {
    "BOT_TOKEN"   = var.bot_api_token
    "MONGO_TOKEN" = var.mongo_token
    "GEM_API_KEY" = var.gemini_key
  }
}
