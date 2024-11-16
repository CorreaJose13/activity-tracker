variable "lambda_exec_role_name" {
  description = "Name for the Lambda execution IAM role"
  type        = string
  default     = "lambda_exec_role"
}

variable "organizations_access_policy_name" {
  description = "Name for the Lambda policy to allow access to AWS Organizations"
  type        = string
  default     = "organizations_access_policy"
}

variable "s3_bucket_name" {
  description = "Name for the S3 bucket to store Terraform backend"
  type        = string
  default     = "terraform-tgbot-lambda"
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
