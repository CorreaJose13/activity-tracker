resource "aws_iam_role" "scheduler_eventbridge_role" {
  name        = var.scheduler_eventbridge_role_name
  description = "IAM role for EventBridge Scheduler to assume and execute scheduled tasks"

  assume_role_policy = jsonencode({
    Version : "2012-10-17",
    Statement : [
      {
        Effect : "Allow",
        Principal : { Service : "scheduler.amazonaws.com" },
        Action : "sts:AssumeRole"
      }
    ]
  })
}

resource "aws_iam_role" "scheduler_lambda_execution_role" {
  name = var.scheduler_lambda_execution_role_name
  assume_role_policy = jsonencode({
    Version = "2012-10-17",
    Statement = [
      {
        Effect    = "Allow",
        Principal = { Service = "lambda.amazonaws.com" },
        Action    = "sts:AssumeRole"
      }
    ]
  })
}

resource "aws_iam_policy" "scheduler_invoke_lambda_policy" {
  name        = var.scheduler_invoke_lambda_policy_name
  description = "Policy to invoke Lambda function"

  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Effect   = "Allow",
        Action   = "lambda:InvokeFunction",
        Resource = aws_lambda_function.scheduler_lambda_function.arn
      },
    ],
  })
}

resource "aws_iam_role_policy_attachment" "scheduler_lambda_policy_attachment" {
  role       = aws_iam_role.scheduler_lambda_execution_role.name
  policy_arn = "arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole"
}

resource "aws_iam_role_policy_attachment" "scheduler_lambda_invoke_policy_attachment" {
  role       = aws_iam_role.scheduler_eventbridge_role.name
  policy_arn = aws_iam_policy.scheduler_invoke_lambda_policy.arn
}
