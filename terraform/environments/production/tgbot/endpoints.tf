module "bot_lambda" {
  source             = "../../../modules/lambda_api_integration/"
  lambda_source_path = "${path.module}/../../../../functions/api/telegram/main.go"
  s3_bucket          = module.lambda_bucket.bucket
  lambda_role        = module.lambda_role.arn
  timeout            = 10
  memory_size        = 128
  log_retention_days = 7

  env_vars = {
    "BOT_TOKEN"   = var.bot_api_token
    "MONGO_TOKEN" = var.mongo_token
    "GEM_API_KEY" = var.gemini_key
  }

  endpoint_name     = "bot"
  rest_api_id       = module.api_gateway.id
  rest_api_exec_arn = module.api_gateway.execution_arn
  parent_id         = module.api_gateway.root_resource_id
  endpoint_path     = var.endpoint_path
  http_method       = "POST"
  stage             = var.stage
}

module "tracks_lambda" {
  source             = "../../../modules/lambda_api_integration/"
  lambda_source_path = "${path.module}/../../../../functions/api/get-available-activities/main.go"
  s3_bucket          = module.lambda_bucket.bucket
  lambda_role        = module.lambda_role.arn
  timeout            = 10
  memory_size        = 128
  log_retention_days = 7

  env_vars = {
    "MONGO_TOKEN" = var.mongo_token
  }

  endpoint_name     = "available-tracks"
  rest_api_id       = module.api_gateway.id
  rest_api_exec_arn = module.api_gateway.execution_arn
  parent_id         = module.api_gateway.root_resource_id
  endpoint_path     = "tracks"
  http_method       = "GET"
  stage             = var.stage
}

module "tracker_lambda" {
  source             = "../../../modules/lambda_api_integration/"
  lambda_source_path = "${path.module}/../../../../functions/api/tracker/main.go"
  s3_bucket          = module.lambda_bucket.bucket
  lambda_role        = module.lambda_role.arn
  timeout            = 10
  memory_size        = 128
  log_retention_days = 7

  env_vars = {
    "MONGO_TOKEN" = var.mongo_token
  }

  endpoint_name     = "tracker"
  rest_api_id       = module.api_gateway.id
  rest_api_exec_arn = module.api_gateway.execution_arn
  parent_id         = module.api_gateway.root_resource_id
  endpoint_path     = "tracker"
  http_method       = "POST"
  stage             = var.stage
}

resource "aws_api_gateway_deployment" "deployment" {
  rest_api_id = module.api_gateway.id

  depends_on = [module.api_gateway, module.bot_lambda]

  lifecycle {
    create_before_destroy = true
  }
}

resource "aws_api_gateway_stage" "stage" {
  deployment_id = aws_api_gateway_deployment.deployment.id
  rest_api_id   = module.api_gateway.id
  stage_name    = var.stage
}

resource "terraform_data" "this" {
  triggers_replace = {
    always_run = "${timestamp()}"
  }

  provisioner "local-exec" {
    command = "curl -X POST 'https://api.telegram.org/bot${var.bot_api_token}/setWebhook?url=${aws_api_gateway_stage.stage.invoke_url}/${var.endpoint_path}'"
  }

  depends_on = [aws_api_gateway_stage.stage]
}
