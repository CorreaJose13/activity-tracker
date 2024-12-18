variable "lambda_execution_role_name" {
  description = "Name for the Lambda execution IAM role"
  type        = string
  default     = "lambda_execution_role"
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

variable "gemini_key" {
  description = "Gemini API key"
  type        = string
  sensitive   = true
}
