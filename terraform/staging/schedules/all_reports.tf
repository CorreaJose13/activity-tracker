resource "aws_scheduler_schedule" "all_reports_scheduler" {
  name = var.all_reports_scheduler_name

  flexible_time_window {
    mode = "OFF"
  }

  schedule_expression          = "cron(7 9,21 * * ? *)"
  schedule_expression_timezone = "America/Bogota"

  target {
    arn      = "arn:aws:scheduler:::aws-sdk:lambda:invoke"
    role_arn = aws_iam_role.scheduler_eventbridge_role.arn

    input = jsonencode({
      FunctionName   = aws_lambda_function.all_reports_lambda_function.function_name
      InvocationType = "Event"
      Payload = jsonencode({
        message = "all-reports"
      })
    })
  }
}
