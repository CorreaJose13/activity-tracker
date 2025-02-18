variable "scheduler_name" {
  description = "The name of the scheduler"
  type        = string
}

variable "schedule_expression" {
  description = "The cron expression that defines when the scheduler should run"
  type        = string
}

variable "role_arn" {
  description = "The ARN of the IAM role that the scheduler will use to invoke the Lambda function"
  type        = string
}

variable "lambda_function_name" {
  description = "The name of the Lambda function to be invoked by the scheduler"
  type        = string
}

variable "message" {
  description = "The message to be passed as input to the Lambda function"
  type        = string
}
