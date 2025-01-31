locals {
  src_path     = "${path.module}/../../../schedules/functions/keratine/main.go"
  binary_path  = "${path.module}/tf_generated/${var.binary_name}"
  archive_path = "${path.module}/tf_generated/${var.scheduler_lambda_function_name}.zip"

  all_reports_src_path     = "${path.module}/../../../schedules/functions/all-reports/main.go"
  all_reports_binary_path  = "${path.module}/tf_generated/${var.binary_name}"
  all_reports_archive_path = "${path.module}/tf_generated/${var.all_reports_lambda_function_name}.zip"
}
