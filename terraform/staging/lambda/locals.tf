locals {
  function_name = "my_lambda_function"
  src_path      = "${path.module}/../../../lambda/"

  binary_name  = "bootstrap"
  binary_path  = "${path.module}/tf_generated/${local.binary_name}"
  archive_path = "${path.module}/tf_generated/${local.function_name}.zip"
}