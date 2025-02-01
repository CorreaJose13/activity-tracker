locals {
  src_path     = "${path.module}/../../../schedules/functions/keratine/main.go"
  binary_path  = "${path.module}/tf_generated/${var.binary_name}"
  archive_path = "${path.module}/tf_generated/${var.scheduler_lambda_function_name}.zip"
}
