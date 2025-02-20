resource "aws_scheduler_schedule" "drink_keratine_scheduler" {
  name = var.scheduler_name

  flexible_time_window {
    mode = "OFF"
  }

  schedule_expression          = var.schedule_expression
  schedule_expression_timezone = "America/Bogota"

  target {
    arn      = "arn:aws:scheduler:::aws-sdk:lambda:invoke"
    role_arn = var.role_arn

    input = jsonencode({
      FunctionName   = var.lambda_function_name
      InvocationType = "Event"
      Payload = jsonencode({
        message = var.message
      })
    })
  }
}
