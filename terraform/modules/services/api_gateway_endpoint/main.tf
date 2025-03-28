resource "aws_api_gateway_resource" "api_resource" {
  rest_api_id = var.api_gateway_id
  parent_id   = var.root_resource_id
  path_part   = var.path
}

resource "aws_api_gateway_method" "api_method" {
  rest_api_id   = var.api_gateway_id
  resource_id   = aws_api_gateway_resource.api_resource.id
  http_method   = var.method
  authorization = "NONE"
}

resource "aws_api_gateway_integration" "api_integration" {
  resource_id             = aws_api_gateway_resource.api_resource.id
  rest_api_id             = var.api_gateway_id
  http_method             = aws_api_gateway_method.api_method.http_method
  integration_http_method = "POST"
  type                    = "AWS_PROXY"
  uri                     = var.lambda_invoke_arn
}

resource "aws_lambda_permission" "api_gateway_lambda" {
  statement_id  = "AllowExecutionFromAPIGateway"
  action        = "lambda:InvokeFunction"
  function_name = var.lambda_name
  principal     = "apigateway.amazonaws.com"
  source_arn    = "${var.api_execution_arn}/*/*/*"
}

// TO DO: Decide whether to use triggers or depends_on
resource "aws_api_gateway_deployment" "api_deployment" {
  rest_api_id = var.api_gateway_id

  triggers = {
    redeployment = sha1(jsonencode([
      aws_api_gateway_resource.api_resource.id,
      aws_api_gateway_method.api_method.id,
      aws_api_gateway_integration.api_integration.id,
    ]))
  }

  lifecycle {
    create_before_destroy = true
  }
}

resource "aws_api_gateway_stage" "api_stage" {
  deployment_id = aws_api_gateway_deployment.api_deployment.id
  rest_api_id   = var.api_gateway_id
  stage_name    = var.stage_name
}
