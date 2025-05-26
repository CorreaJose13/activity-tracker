module "update_scheduler" {
  source             = "../../../modules/lambda_scheduler_integration/"
  lambda_source_path = "${path.module}/../../../../functions/schedules/keratine/main.go"
  function_name      = "keratin-scheduler-lambda"
  s3_bucket          = "activity-tracker-bot-lambdas"
  lambda_role        = data.aws_iam_role.this.arn
  timeout            = 10
  memory_size        = 128
  log_retention_days = 7

  env_vars = {
    "BOT_TOKEN" = var.bot_api_token
  }

  schedule_expression     = "cron(5 9,21 * * ? *)"
  scheduler_name          = "keratin_scheduler"
  message                 = "acordate de tomar la creatina 💪 sapa asquerosa"
  lambda_scheduler_role   = "role_for_lambda_keratine_scheduler"
  lambda_scheduler_policy = "invoke_lambda_policy_for_keratine_scheduler"
}

module "report_all_scheduler" {
  source             = "../../../modules/lambda_scheduler_integration/"
  lambda_source_path = "${path.module}/../../../../functions/schedules/all-reports/main.go"
  function_name      = "report-all-scheduler-lambda"
  s3_bucket          = "activity-tracker-bot-lambdas"
  lambda_role        = data.aws_iam_role.this.arn
  timeout            = 10
  memory_size        = 128
  log_retention_days = 7

  env_vars = {
    "BOT_TOKEN" = var.bot_api_token
  }

  schedule_expression     = "cron(0 21 ? * 7 *)"
  scheduler_name          = "report_all_scheduler"
  lambda_scheduler_role   = "role_for_lambda_report_scheduler"
  lambda_scheduler_policy = "invoke_lambda_policy_for_report_scheduler"
}
