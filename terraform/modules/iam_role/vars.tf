variable "role_name" {
  description = "The name of the IAM role"
  type        = string
}

variable "assume_role_identifiers" {
  description = "List of identifiers that are allowed to assume the role"
  type        = list(string)
}
