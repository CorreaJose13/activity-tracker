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
    bucket         = "tf-state-tracker-lambda"
    key            = "state/tracker.tfstate"
    region         = "us-east-1"
    encrypt        = true
    dynamodb_table = "tracker_terraform_state_db"
  }
}

provider "aws" {
  region = "us-east-1"
}

locals {
  src_path     = "${path.module}/../../../functions/tracker/main.go"
  binary_path  = "${path.module}/tf_generated/bootstrap"
  archive_path = "${path.module}/tf_generated/lambda_function.zip"
}

module "iam_role_lambda_tracker" {
  source = "../../modules/iam_role/"

  role_name               = "iam_for_lambda_tracker"
  assume_role_identifiers = ["lambda.amazonaws.com"]
}

module "iam_policy_attachment_logs" {
  source = "../../modules/iam_policy_attachment/"

  policy_name = "logs_lambda_tracker"
  action      = ["logs:CreateLogStream", "logs:PutLogEvents"]
  resource    = "arn:aws:logs:*:*:*"
  role_name   = module.iam_role_lambda_tracker.role_name
}

module "lambda_function" {
  source = "../../modules/services/lambda_function/"

  src_path             = local.src_path
  binary_path          = local.binary_path
  archive_path         = local.archive_path
  lambda_function_name = "tracker_lambda"
  role_arn             = module.iam_role_lambda_tracker.role_arn
  environment_variables = {
    "MONGO_TOKEN" = var.mongo_token
  }
}

module "iam_role_tracker" {
  source = "../../modules/iam_role/"

  role_name               = "iam_for_tracker"
  assume_role_identifiers = ["scheduler.amazonaws.com"]
}

module "iam_policy_attachment_tracker" {
  source = "../../modules/iam_policy_attachment/"

  policy_name = "invoke_lambda_policy_tracker"
  action      = ["lambda:InvokeFunction"]
  resource    = module.lambda_function.lambda_function_arn
  role_name   = module.iam_role_tracker.role_name
}
