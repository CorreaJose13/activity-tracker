variable "lambda_function_name" {
  description = "Name of the lambda function"
  type        = string
  default     = "activity-tracker-bot"
}
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
