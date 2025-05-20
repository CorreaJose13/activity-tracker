locals {
  zip_file = "${var.endpoint_name}.zip"
}

resource "terraform_data" "this" {
  triggers_replace = {
    always_run = "${timestamp()}"
  }
  provisioner "local-exec" {
    working_dir = dirname(var.lambda_source_path)
    command     = "make publish"
    environment = {
      BUCKET_NAME = var.s3_bucket
      BUILD_NAME  = local.zip_file
    }
  }

  depends_on = [var.s3_bucket]
}

module "lambda_function" {
  source = "../compute/lambda/"

  s3_bucket = var.s3_bucket
  s3_key    = local.zip_file

  function_name = "stock-api-${var.endpoint_name}"
  description   = "Lambda function for ${var.endpoint_name}"
  role          = var.lambda_role

  runtime            = var.runtime
  handler            = var.handler
  timeout            = var.timeout
  memory_size        = var.memory_size
  log_retention_days = var.log_retention_days
  env_vars           = var.env_vars

  depends_on = [terraform_data.this]
}

module "api_endpoint" {
  source            = "../network/api_gateway_endpoint/"
  rest_api_id       = var.rest_api_id
  rest_api_exec_arn = var.rest_api_exec_arn
  parent_id         = var.parent_id
  path              = var.endpoint_path
  method            = var.http_method
  stage             = var.stage

  lambda_name       = module.lambda_function.name
  lambda_invoke_arn = module.lambda_function.invoke_arn
}
