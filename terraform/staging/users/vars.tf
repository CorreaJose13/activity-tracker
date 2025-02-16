variable "region" {
  description = "The AWS region"
  type        = string
  default     = "us-east-1"
}

variable "developer_group" {
  description = "The developer group name"
  type        = string
  default     = "developer"
}

variable "tracker_group" {
  description = "The tracker group name"
  type        = string
  default     = "tracker"
}
