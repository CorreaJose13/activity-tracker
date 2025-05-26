module "api_gateway" {
  source             = "../../../modules/network/api_gateway/"
  name               = "activity-tracker-api"
  description        = "API gateway for the activity tracker bot"
  log_retention_days = 7
}
