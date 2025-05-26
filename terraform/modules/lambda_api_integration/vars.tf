#Lambda variables
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

#API Gateway variables
variable "endpoint_name" {
  description = "The name of the endpoint"
  type        = string
}

variable "rest_api_id" {
  description = "The name of the API Gateway"
  type        = string
}

variable "rest_api_exec_arn" {
  description = "The execution ARN of the API Gateway"
  type        = string
}

variable "parent_id" {
  description = "The ID of the parent resource in the API Gateway"
  type        = string
}

variable "http_method" {
  description = "The HTTP method for the API Gateway endpoint"
  type        = string
}

variable "endpoint_path" {
  description = "The path for the API Gateway endpoint"
  type        = string
}

variable "stage" {
  description = "The stage for the API Gateway"
  type        = string
}
