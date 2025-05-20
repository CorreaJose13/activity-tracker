data "aws_s3_object" "this" {
  bucket = var.s3_bucket
  key    = var.s3_key
}

resource "aws_lambda_function" "this" {
  function_name = var.function_name
  description   = var.description
  role          = var.role

  s3_bucket        = data.aws_s3_object.this.bucket
  s3_key           = data.aws_s3_object.this.key
  source_code_hash = data.aws_s3_object.this.version_id

  runtime = var.runtime
  handler = var.handler

  timeout     = var.timeout
  memory_size = var.memory_size

  environment {
    variables = var.env_vars
  }
}

resource "aws_cloudwatch_log_group" "this" {
  name              = "/aws/lambda/${aws_lambda_function.this.function_name}"
  retention_in_days = var.log_retention_days
  lifecycle {
    prevent_destroy = false
  }
}
