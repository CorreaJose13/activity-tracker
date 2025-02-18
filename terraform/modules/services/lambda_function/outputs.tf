output "lambda_function_name" {
  description = "Name of lambda function"
  value       = aws_lambda_function.lambda_function.function_name
}

output "lambda_function_arn" {
  description = "ARN of lambda function"
  value       = aws_lambda_function.lambda_function.arn
}
