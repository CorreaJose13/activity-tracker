terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.75.0"
    }
  }
  backend "s3" {
    bucket         = "tf-state-activity-tracker-api"
    key            = "state/terraform.tfstate"
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
  src_path     = "${path.module}/../../../functions/telegram/main.go"
  binary_path  = "${path.module}/tf_generated/bootstrap"
  archive_path = "${path.module}/tf_generated/lambda_function.zip"
}

// Change role name
module "iam_role" {
  source = "../../modules/iam_role/"

  role_name               = "test_role"
  assume_role_identifiers = ["lambda.amazonaws.com"]
}

// Change policy name
module "iam_policy_attachment_logs" {
  source = "../../modules/iam_policy_attachment/"

  policy_name = "test_policy"
  action      = ["logs:CreateLogStream", "logs:PutLogEvents"]
  resource    = "arn:aws:logs:*:*:*"
  role_name   = module.iam_role.role_name
}

// Change lambda_function_name for specific name of the lambda function
module "lambda_function" {
  source = "../../modules/services/lambda_function/"

  src_path             = local.src_path
  binary_path          = local.binary_path
  archive_path         = local.archive_path
  lambda_function_name = "test_lambda"
  role_arn             = module.iam_role.role_arn
}

module "api_gateway" {
  source      = "../../modules/services/api_gateway/"
  name        = "Activity Tracker API"
  description = "Este es el api gateway para el activity tracker, sigan viendo"
}

// Change path and method for specific values
module "api_gateway_endpoint" {
  source            = "../../modules/services/api_gateway_endpoint/"
  api_gateway_id    = module.api_gateway.api_gateway_id
  root_resource_id  = module.api_gateway.api_gateway_root_resource_id
  path              = "tracks"
  method            = "GET"
  api_execution_arn = module.api_gateway.api_gateway_execution_arn
  stage_name        = "dev"
  lambda_invoke_arn = module.lambda_function.lambda_function_invoke_arn
  lambda_name       = module.lambda_function.lambda_function_name
}

output "endpoint_url" {
  value       = module.api_gateway_endpoint.endpoint_url
  description = "The invoke URL for the API Gateway endpoint"
}
