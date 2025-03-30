variable "api_gateway_id" {
  description = "Id of the api gateway"
  type        = string
}

variable "root_resource_id" {
  description = "Root resource id"
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

variable "stage_name" {
  description = "Name of the stage"
  type        = string
}

variable "lambda_invoke_arn" {
  description = "ARN of the lambda function to be invoked"
  type        = string
}

variable "lambda_name" {
  description = "Name of the lambda function"
  type        = string
}

variable "api_execution_arn" {
  description = "ARN of the api gateway execution"
  type        = string
}


