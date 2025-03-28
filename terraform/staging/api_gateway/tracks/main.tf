terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.75.0"
    }
  }
  backend "s3" {
    bucket         = "tf-state-activity-tracker-api"
    key            = "state/tracks/terraform.tfstate"
    region         = "us-east-1"
    encrypt        = true
    dynamodb_table = "terraform_api_state_db"
  }
}

provider "aws" {
  region = "us-east-1"
}

// Change src_path for specific location of the lambda function
locals {
  src_path     = "${path.module}/../../../../functions/api/get-available-activities/main.go"
  binary_path  = "${path.module}/tf_generated/bootstrap"
  archive_path = "${path.module}/tf_generated/lambda_function.zip"
}

module "iam_role" {
  source = "../../../modules/iam_role/"

  role_name               = "iam_for_get_available_activities_endpoint"
  assume_role_identifiers = ["lambda.amazonaws.com"]
}

module "iam_policy_attachment_logs" {
  source = "../../../modules/iam_policy_attachment/"

  policy_name = "logs_lambda_get_available_activities"
  action      = ["logs:CreateLogStream", "logs:PutLogEvents"]
  resource    = "arn:aws:logs:*:*:*"
  role_name   = module.iam_role.role_name
}

module "lambda_function" {
  source = "../../../modules/services/lambda_function/"

  src_path             = local.src_path
  binary_path          = local.binary_path
  archive_path         = local.archive_path
  lambda_function_name = "lambda_get_available_activities"
  role_arn             = module.iam_role.role_arn

  environment_variables = {
    "MONGO_TOKEN" = var.mongo_token
  }
}

data "terraform_remote_state" "api_gateway" {
  backend = "s3"
  config = {
    bucket         = "tf-state-activity-tracker-api"
    key            = "state/terraform.tfstate"
    region         = "us-east-1"
    encrypt        = true
    dynamodb_table = "terraform_api_state_db"
  }
}

module "api_gateway_endpoint" {
  source            = "../../../modules/services/api_gateway_endpoint/"
  api_gateway_id    = data.terraform_remote_state.api_gateway.outputs.api_gateway_id
  root_resource_id  = data.terraform_remote_state.api_gateway.outputs.api_gateway_root_resource_id
  path              = "tracks"
  method            = "GET"
  api_execution_arn = data.terraform_remote_state.api_gateway.outputs.api_gateway_execution_arn
  stage_name        = "dev"
  lambda_invoke_arn = module.lambda_function.lambda_function_invoke_arn
  lambda_name       = module.lambda_function.lambda_function_name
}

output "endpoint_url" {
  value       = module.api_gateway_endpoint.endpoint_url
  description = "The invoke URL for the API Gateway endpoint"
}
