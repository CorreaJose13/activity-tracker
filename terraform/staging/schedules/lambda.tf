resource "aws_iam_role" "schedules_lambda" {
  name               = "schedules-lambda"
  assume_role_policy = jsonencode({
    Version   = "2012-10-17",
    Statement = [
      {
        Effect    = "Allow",
        Principal = {
          Service = "lambda.amazonaws.com"
        },
        Action    = "sts:AssumeRole"
      }
    ]
  })
}

resource "aws_iam_policy_attachment" "lambda_basic_logs" {
  name       = "lambda_basic_logs"
  roles      = [aws_iam_role.schedules_lambda.name]
  policy_arn = "arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole"
}

resource "aws_lambda_function" "schedules_route" {
  filename      = "example.zip"
  function_name = "schedules-route"
  role          = aws_iam_role.schedules_lambda.arn
  runtime       = "provided.al2"
  handler       = "main"
  architectures = ["arm64"]
  source_code_hash = filebase64sha256("example.zip")
}
