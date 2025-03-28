terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.75.0"
    }
  }
  backend "s3" {
    bucket         = "tf-state-activity-tracker-api"
    key            = "state/terraform.tfstate"
    region         = "us-east-1"
    encrypt        = true
    dynamodb_table = "terraform_api_state_db"
  }
}

module "api_gateway" {
  source      = "../../../modules/services/api_gateway/"
  name        = "Activity Tracker API"
  description = "Este es el api gateway para el activity tracker, sigan viendo"
}
