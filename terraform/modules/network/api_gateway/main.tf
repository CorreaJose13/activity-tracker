resource "aws_api_gateway_rest_api" "this" {
  name        = var.name
  description = var.description
}

resource "aws_cloudwatch_log_group" "api_gateway" {
  name              = "/aws/apigateway/${var.name}"
  retention_in_days = var.log_retention_days
}

// TO DO: For improvement stage
# resource "aws_api_gateway_method_settings" "this" {
#   rest_api_id = aws_api_gateway_rest_api.this.id
#   stage_name  = aws_api_gateway_stage.this.stage_name
#   method_path = "*/*"

#   settings {
#     metrics_enabled        = var.enable_cloudwatch_metrics
#     logging_level          = var.logging_level
#     data_trace_enabled     = var.data_trace_enabled
#     throttling_burst_limit = var.throttling_burst_limit
#     throttling_rate_limit  = var.throttling_rate_limit
#   }
# }
