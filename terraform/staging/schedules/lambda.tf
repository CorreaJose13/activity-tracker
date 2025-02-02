resource "null_resource" "function_binary" {
  triggers = {
    always_run = "${timestamp()}"
  }
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

resource "aws_lambda_function" "scheduler_lambda_function" {
  filename         = local.archive_path
  function_name    = var.scheduler_lambda_function_name
  role             = aws_iam_role.scheduler_lambda_execution_role.arn
  runtime          = "provided.al2023"
  handler          = "main"
  architectures    = ["x86_64"]
  source_code_hash = data.archive_file.function_archive.output_base64sha256

  environment {
    variables = {
      BOT_TOKEN = var.bot_api_token
    }
  }
}
