output "api_gateway_id" {
  description = "Id of the api gateway"
  value       = aws_api_gateway_rest_api.api_gateway.id
}

output "api_gateway_root_resource_id" {
  description = "Root resource id of the api gateway"
  value       = aws_api_gateway_rest_api.api_gateway.root_resource_id
}

output "api_gateway_execution_arn" {
  description = "Execution ARN of the api gateway"
  value       = aws_api_gateway_rest_api.api_gateway.execution_arn
}

