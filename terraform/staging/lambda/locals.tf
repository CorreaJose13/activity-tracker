locals {
  src_path     = "${path.module}/../../../lambda/main.go"
  binary_path  = "${path.module}/tf_generated/${var.binary_name}"
  archive_path = "${path.module}/tf_generated/${var.lambda_function_name}.zip"
}
