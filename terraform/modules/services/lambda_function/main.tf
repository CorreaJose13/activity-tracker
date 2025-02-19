resource "null_resource" "function_binary" {
  triggers = {
    always_run = "${timestamp()}"
  }
  provisioner "local-exec" {
    command = "GOOS=linux GOARCH=amd64 go build -o ${var.binary_path} ${var.src_path}"
  }
}

data "archive_file" "function_archive" {
  depends_on = [null_resource.function_binary]

  type        = "zip"
  source_file = var.binary_path
  output_path = var.archive_path
}

resource "aws_lambda_function" "lambda_function" {
  filename         = var.archive_path
  function_name    = var.lambda_function_name
  role             = var.role_arn
  runtime          = "provided.al2023"
  handler          = "main"
  architectures    = ["x86_64"]
  source_code_hash = data.archive_file.function_archive.output_base64sha256

  environment {
    variables = var.environment_variables
  }
}

resource "aws_cloudwatch_log_group" "function_log_group" {
  name              = "/aws/lambda/${aws_lambda_function.lambda_function.function_name}"
  retention_in_days = 7
  lifecycle {
    prevent_destroy = false
  }
}
