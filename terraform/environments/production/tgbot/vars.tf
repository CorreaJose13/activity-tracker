variable "bot_api_token" {
  description = "Telegram bot key"
  type        = string
  sensitive   = true
}

variable "mongo_token" {
  description = "MongoDB connection token"
  type        = string
  sensitive   = true
}

variable "gemini_key" {
  description = "Gemini API key"
  type        = string
  sensitive   = true
}

variable "stage" {
  description = "Stage of the API Gateway"
  type        = string
}

variable "endpoint_path" {
  description = "Path of the API Gateway endpoint"
  type        = string
}
