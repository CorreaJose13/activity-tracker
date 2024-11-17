terraform {
  backend "s3" {
    bucket         = "tf-state-activity-tracker-bot"
    key            = "state/terraform.tfstate"
    region         = "us-east-1"
    encrypt        = true
    dynamodb_table = "terraform_state_db"
  }
}
