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
