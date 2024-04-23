resource "aws_iam_role" "schedules_eventbridge" {
  name               = "schedules-eventbridge"
  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Principal": {
        "Service": "scheduler.amazonaws.com"
      },
      "Action": "sts:AssumeRole"
    }
  ]
}
EOF
}

resource "aws_iam_policy" "schedules_invoke_lambda" {
  name        = "schedules-invoke-lambda"
  description = "Policy to invoke Lambda function"
  
  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Effect   = "Allow",
        Action   = "lambda:InvokeFunction",
        Resource = aws_lambda_function.schedules_route.arn
      }
    ]
  })
}

resource "aws_iam_role_policy_attachment" "schedules_lambda_policy_attachment" {
  role       = aws_iam_role.schedules_eventbridge.name
  policy_arn = aws_iam_policy.schedules_invoke_lambda.arn
}