locals {
  bucket_name   = "terraform-tgbot-lambda"
  function_name = "tg_bot_lambda"
  src_path      = "${path.module}/../../../lambda/main.go"

  binary_name  = "bootstrap"
  binary_path  = "${path.module}/tf_generated/${local.binary_name}"
  archive_path = "${path.module}/tf_generated/${local.function_name}.zip"

  bot_key = "7048395318:AAGUGeQM-wjZymlZujdBQEZZD0EEHj-qB64"
  db_key  = "activitytracker"
}
