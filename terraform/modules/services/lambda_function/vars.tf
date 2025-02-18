#Required variables

variable "src_path" {
  description = "Path to the source code"
  type        = string
}

variable "binary_path" {
  description = "Path to the binary"
  type        = string
}

variable "archive_path" {
  description = "Path to the archive"
  type        = string
}

variable "lambda_function_name" {
  description = "Name for the Lambda function"
  type        = string
}

variable "role_arn" {
  description = "The ARN of the IAM role for the Lambda function"
  type        = string
}

#Optional variables

variable "environment_variables" {
  description = "Environment variables for the Lambda function"
  type        = map(string)
  default     = {}
}
