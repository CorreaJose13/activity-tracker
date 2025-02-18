variable "policy_name" {
  description = "The name of the IAM policy"
  type        = string
}

variable "action" {
  description = "The action that the policy will allow"
  type        = string
}

variable "lambda_function_arn" {
  description = "The ARN of the Lambda function that the policy will allow to invoke"
  type        = string
}

variable "role_name" {
  description = "The name of the IAM role to which the policy will be attached"
  type        = string
}
