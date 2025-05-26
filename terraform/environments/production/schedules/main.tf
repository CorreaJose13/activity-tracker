terraform {
  backend "s3" {
    bucket       = "tf-state-activity-tracker-schedules"
    key          = "state/tracker.tfstate"
    region       = "us-east-1"
    encrypt      = true
    use_lockfile = true
  }
}

provider "aws" {
  region = "us-east-1"
}
