variable "scheduler_name" {
  description = "The name of the scheduler"
  type        = string
}

variable "schedule_expression" {
  description = "The cron expression that defines when the scheduler should run"
  type        = string
}

variable "lambda_scheduler_role" {
  description = "The name of the IAM role for the Lambda scheduler"
  type        = string
}

variable "lambda_scheduler_policy" {
  description = "The name of the IAM policy for the Lambda scheduler"
  type        = string
}

#Lambda variables
variable "function_name" {
  description = "The name of the Lambda function"
  type        = string
}

variable "lambda_source_path" {
  description = "The path to the source code of the Lambda function"
  type        = string
}

variable "s3_bucket" {
  description = "The S3 bucket where the Lambda function code is stored"
  type        = string
}

variable "lambda_role" {
  description = "The IAM role ARN for the Lambda function"
  type        = string
}

variable "runtime" {
  description = "The runtime for the Lambda function"
  type        = string
  default     = "provided.al2"
}

variable "handler" {
  description = "The handler for the Lambda function"
  type        = string
  default     = "bootstrap"
}

variable "timeout" {
  description = "The timeout for the Lambda function in seconds"
  type        = number
}

variable "memory_size" {
  description = "The memory size for the Lambda function in MB"
  type        = number
}

variable "log_retention_days" {
  description = "The number of days to retain logs for the Lambda function"
  type        = number
}

variable "env_vars" {
  description = "Environment variables for the Lambda function"
  type        = map(string)
  default     = {}
}
