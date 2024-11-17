variable "lambda_execution_role_name" {
  description = "Name for the Lambda execution IAM role"
  type        = string
  default     = "lambda_execution_role"
}

variable "s3_bucket_name" {
  description = "Name for the S3 bucket to store Terraform backend"
  type        = string
  default     = "tf-state-activity-tracker-bot"
}

variable "lambda_function_name" {
  description = "Name for the Lambda function"
  type        = string
  default     = "tg_bot_lambda"
}

variable "binary_name" {
  description = "Name for the binary file for Lambda deployment"
  type        = string
  default     = "bootstrap"
}

variable "terraform_state_db_name" {
  description = "Name for the DynamoDB table for Terraform backend"
  type        = string
  default     = "terraform_state_db"
}

variable "region" {
  description = "The AWS region"
  type        = string
  default     = "us-east-1"
}

variable "bot_api_token" {
  description = "Telegram bot key"
  type        = string
  sensitive   = true
}

variable "mongo_token" {
  description = "MongoDB connection token"
  type        = string
  sensitive   = true
}
