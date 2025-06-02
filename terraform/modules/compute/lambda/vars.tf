variable "function_name" {
  description = "Name for the Lambda function"
  type        = string
}

variable "role" {
  description = "The ARN of the IAM role for the Lambda function"
  type        = string
}

variable "s3_bucket" {
  description = "The S3 bucket where the Lambda function code is stored"
  type        = string
}

variable "s3_key" {
  description = "The S3 key for the Lambda function code"
  type        = string
}

variable "description" {
  description = "Description for the Lambda function"
  type        = string
  default     = "Lambda function created by Terraform"
}

variable "runtime" {
  description = "The runtime for the Lambda function"
  type        = string
}

variable "handler" {
  description = "The handler for the Lambda function"
  type        = string
}

variable "timeout" {
  description = "Lambda function timeout in seconds"
  type        = number
}

variable "memory_size" {
  description = "Lambda function memory size in MB"
  type        = number
}

variable "env_vars" {
  description = "Environment variables for the Lambda function"
  type        = map(string)
}

variable "log_retention_days" {
  description = "CloudWatch log retention period in days"
  type        = number
}
