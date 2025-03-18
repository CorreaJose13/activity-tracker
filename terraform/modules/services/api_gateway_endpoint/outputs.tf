output "endpoint_url" {
  description = "The complete endpoint URL for the API resource"
  value       = "${aws_api_gateway_stage.api_stage.invoke_url}/${var.path}"
}
