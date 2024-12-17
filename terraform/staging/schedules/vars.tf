variable "scheduler_eventbridge_role_name" {
  description = "Name for the EventBridge schedule resource used to trigger events"
  type        = string
  default     = "scheduler_eventbridge_role"
}

variable "scheduler_lambda_execution_role_name" {
  description = "Name for the scheduler Lambda execution IAM role"
  type        = string
  default     = "scheduler_lambda_execution_role"
}

variable "scheduler_invoke_lambda_policy_name" {
  description = "Name for the policy to invoke Lambda function"
  type        = string
  default     = "scheduler_invoke_lambda_policy"
}

variable "scheduler_lambda_function_name" {
  description = "Name for the scheduler Lambda function"
  type        = string
  default     = "tg_bot_scheduler_lambda"
}

variable "scheduler_name" {
  description = "Name for the scheduler"
  type        = string
  default     = "tg_bot_scheduler"
}

variable "binary_name" {
  description = "Name for the binary file for the scheduler Lambda deployment"
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

variable "chat_id" {
  description = "Telegram chat id"
  type        = string
  sensitive   = true
}

