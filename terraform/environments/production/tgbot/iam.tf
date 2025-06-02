module "lambda_role" {
  source = "../../../modules/iam/iam_role/"

  name     = "role_for_tg_bot_lambda"
  services = ["lambda.amazonaws.com"]
}

module "lambda_logs_policy" {
  source = "../../../modules/iam/iam_policy_attachment/"

  name      = "logs_policy_for_tg_bot_lambda"
  action    = ["logs:CreateLogStream", "logs:PutLogEvents"]
  resource  = "arn:aws:logs:*:*:*"
  role_name = module.lambda_role.name
}
