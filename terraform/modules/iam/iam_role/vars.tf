variable "name" {
  description = "The name of the IAM role"
  type        = string
}

variable "services" {
  description = "List of services that are allowed to assume the role"
  type        = list(string)
}
