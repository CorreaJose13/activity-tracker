resource "aws_api_gateway_resource" "this" {
  rest_api_id = var.rest_api_id
  parent_id   = var.parent_id
  path_part   = var.path
}

resource "aws_api_gateway_method" "this" {
  rest_api_id   = var.rest_api_id
  resource_id   = aws_api_gateway_resource.this.id
  http_method   = var.method
  authorization = "NONE"
}

resource "aws_api_gateway_integration" "this" {
  rest_api_id = var.rest_api_id
  resource_id = aws_api_gateway_resource.this.id

  http_method             = aws_api_gateway_method.this.http_method
  integration_http_method = "POST"
  type                    = "AWS"
  uri                     = var.lambda_invoke_arn
}

resource "aws_lambda_permission" "this" {
  statement_id  = "AllowExecutionFromAPIGateway"
  action        = "lambda:InvokeFunction"
  function_name = var.lambda_name
  principal     = "apigateway.amazonaws.com"
  source_arn    = "${var.rest_api_exec_arn}/*/*/*"
}

resource "aws_api_gateway_method_response" "this" {
  rest_api_id = var.rest_api_id
  resource_id = aws_api_gateway_resource.this.id
  http_method = aws_api_gateway_method.this.http_method
  status_code = "200"
}

resource "aws_api_gateway_integration_response" "this" {
  rest_api_id = var.rest_api_id
  resource_id = aws_api_gateway_resource.this.id
  http_method = aws_api_gateway_method.this.http_method
  status_code = aws_api_gateway_method_response.this.status_code

  depends_on = [
    aws_api_gateway_method.this,
    aws_api_gateway_integration.this
  ]
}
