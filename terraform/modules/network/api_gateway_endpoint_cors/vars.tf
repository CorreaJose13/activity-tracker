variable "rest_api_id" {
  description = "Name of the api gateway"
  type        = string
}

variable "rest_api_exec_arn" {
  description = "Execution ARN of the api gateway"
  type        = string
}

variable "parent_id" {
  description = "ID of the parent resource"
  type        = string
}

variable "path" {
  description = "Path of the api gateway resource"
  type        = string
}

variable "method" {
  description = "HTTP method"
  type        = string
}

variable "stage" {
  description = "Name of the stage"
  type        = string
}

variable "lambda_name" {
  description = "Name of the lambda function"
  type        = string
}

variable "lambda_invoke_arn" {
  description = "Lambda function ARN"
  type        = string
}
