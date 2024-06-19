locals {
  bucket_name   = "terraform-tg-lambda"
  bucket_key    = "dev/terraform.tfstate"
  function_name = "tg_bot_lambda"
  src_path      = "${path.module}/../../../lambda/"

  binary_name  = "bootstrap"
  binary_path  = "${path.module}/tf_generated/${local.binary_name}"
  archive_path = "${path.module}/tf_generated/${local.function_name}.zip"
}
