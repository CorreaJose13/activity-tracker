variable "name" {
  description = "The name of the IAM policy"
  type        = string
}

variable "description" {
  description = "The description of the IAM policy"
  type        = string
}

variable "action" {
  description = "The action that the policy will allow"
  type        = list(string)
}

variable "resource" {
  description = "The resource that the policy will grant permissions to access or invoke"
  type        = string
}

variable "role_name" {
  description = "The name of the IAM role to which the policy will be attached"
  type        = string
}

variable "path" {
  description = "The path for the IAM policy"
  type        = string
  default     = "/"
}
