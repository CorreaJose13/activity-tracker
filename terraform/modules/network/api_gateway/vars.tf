variable "name" {
  description = "Name for the API gateway"
  type        = string
}

variable "description" {
  description = "Description for the API gateway"
  type        = string
}

variable "log_retention_days" {
  description = "Number of days to retain logs in CloudWatch"
  type        = number
}
