resource "aws_dynamodb_table" "terraform-state-db" {
  name         = var.terraform_state_db_name
  billing_mode = "PAY_PER_REQUEST"
  hash_key     = "LockID"

  attribute {
    name = "LockID"
    type = "S"
  }

  deletion_protection_enabled = true
}
