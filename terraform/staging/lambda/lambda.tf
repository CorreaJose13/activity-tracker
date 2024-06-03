resource "null_resource" "function_binary" {
  provisioner "local-exec" {
    command = "GOOS=linux GOARCH=amd64 go build -o ${local.binary_path} ${local.src_path}"
  }
}

data "archive_file" "function_archive" {
  depends_on = [null_resource.function_binary]

  type        = "zip"
  source_file = local.binary_path
  output_path = local.archive_path
}

resource "aws_lambda_function" "lambda_function" {
  filename         = local.archive_path
  function_name    = "my_lambda_function"
  role             = aws_iam_role.lambda_exec.arn
  runtime          = "provided.al2023"
  handler          = "main"
  architectures    = ["x86_64"]
  source_code_hash = data.archive_file.function_archive.output_base64sha256
  environment {
    variables = {
      BOT_TOKEN   = data.aws_secretsmanager_secret_version.bot_token_version.secret_string
      MONGO_TOKEN = data.aws_secretsmanager_secret_version.mongo_token_version.secret_string
    }
  }
}

resource "aws_lambda_permission" "apigw_lambda" {
  statement_id  = "AllowAPIGatewayInvoke"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.lambda_function.function_name
  principal     = "apigateway.amazonaws.com"
  source_arn    = "${aws_api_gateway_rest_api.my_api.execution_arn}/*/*"
}

