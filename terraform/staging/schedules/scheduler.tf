resource "aws_scheduler_schedule" "drink_water" {
  name = var.scheduler_name

  flexible_time_window {
    mode = "OFF"
  }

  schedule_expression          = "rate(1 minutes)"
  schedule_expression_timezone = "America/Bogota"

  target {
    arn      = "arn:aws:scheduler:::aws-sdk:lambda:invoke"
    role_arn = aws_iam_role.scheduler_eventbridge_role.arn

    input = jsonencode({
      FunctionName   = aws_lambda_function.scheduler_lambda_function.function_name
      InvocationType = "Event"
      Payload = jsonencode({
        message = "drink-water"
      })
    })
  }
}
