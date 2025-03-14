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
    bucket         = "tf-state-scheduler-activity-tracker-bot"
    key            = "state/terraform.tfstate"
    region         = "us-east-1"
    encrypt        = true
    dynamodb_table = "scheduler_terraform_state_db"
  }
}

provider "aws" {
  region = "us-east-1"
}

locals {
  src_path     = "${path.module}/../../../../functions/schedules/keratine/main.go"
  binary_path  = "${path.module}/tf_generated/bootstrap"
  archive_path = "${path.module}/tf_generated/lambda_function.zip"
}

module "iam_role_lambda_scheduler" {
  source = "../../../modules/iam_role/"

  role_name               = "iam_for_lambda_scheduler"
  assume_role_identifiers = ["lambda.amazonaws.com"]
}

module "iam_policy_attachment_logs" {
  source = "../../../modules/iam_policy_attachment/"

  policy_name = "logs_lambda_keratine_scheduler"
  action      = ["logs:CreateLogStream", "logs:PutLogEvents"]
  resource    = "arn:aws:logs:*:*:*"
  role_name   = module.iam_role_lambda_scheduler.role_name
}

module "lambda_function" {
  source = "../../../modules/services/lambda_function/"

  src_path             = local.src_path
  binary_path          = local.binary_path
  archive_path         = local.archive_path
  lambda_function_name = "tg_bot_keratine_scheduler_lambda"
  role_arn             = module.iam_role_lambda_scheduler.role_arn
  environment_variables = {
    "BOT_TOKEN" = var.bot_api_token
  }
}

module "iam_role_scheduler" {
  source = "../../../modules/iam_role/"

  role_name               = "iam_for_scheduler"
  assume_role_identifiers = ["scheduler.amazonaws.com"]
}

module "iam_policy_attachment" {
  source = "../../../modules/iam_policy_attachment/"

  policy_name = "scheduler_invoke_lambda_policy"
  action      = ["lambda:InvokeFunction"]
  resource    = module.lambda_function.lambda_function_arn
  role_name   = module.iam_role_scheduler.role_name
}

module "scheduler" {
  source = "../../../modules/services/scheduler/"

  scheduler_name       = "tg_bot_scheduler"
  schedule_expression  = "cron(5 9,21 * * ? *)"
  role_arn             = module.iam_role_scheduler.role_arn
  lambda_function_name = module.lambda_function.lambda_function_name
  message              = "acordate de tomar la creatina 💪 sapa asquerosa"
}
