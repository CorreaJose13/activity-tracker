terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.46.0"
    }
  }
}

provider "aws" {
  region = "us-east-1"
}

resource "aws_scheduler_schedule" "drink_water" {
  name = "drink-water"

  flexible_time_window {
    mode = "OFF"
  }

  schedule_expression = "cron(30 22 * * ? *)"
  schedule_expression_timezone = "America/Bogota"

  target {
    arn      = "arn:aws:scheduler:::aws-sdk:lambda:invoke"
    role_arn = aws_iam_role.schedules_eventbridge.arn

    input = jsonencode({
      FunctionName = aws_lambda_function.schedules_route.function_name
      InvocationType = "Event"
      Payload = jsonencode({
        message = "drink-water"
      })
    })
  }
}