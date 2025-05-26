module "lambda_bucket" {
  source = "../../../modules/storage/s3"

  name          = "activity-tracker-bot-lambdas"
  force_destroy = true
}
