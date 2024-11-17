resource "aws_iam_role" "lambda_execution_role" {
  name        = var.lambda_execution_role_name
  description = "IAM Role for Lambda function execution"

  assume_role_policy = jsonencode({
    Version = "2012-10-17",
    Statement = [
      {
        Effect    = "Allow",
        Principal = { Service = "lambda.amazonaws.com" },
        Action    = "sts:AssumeRole"
      },
    ],
  })
}

resource "aws_iam_policy" "restrict_lambda_modifications" {
  name        = "RestrictLambdaModifications"
  description = "Restrict updates to Lambda function to the organization"

  policy = jsonencode({
    Version = "2012-10-17",
    Statement = [
      {
        Sid    = "AllowOrgToModifyLambda",
        Effect = "Allow",
        Action = [
          "lambda:UpdateFunctionCode",
          "lambda:UpdateFunctionConfiguration",
          "lambda:AddPermission",
          "lambda:RemovePermission"
        ],
        Resource = aws_lambda_function.lambda_function.arn,
        Condition = {
          StringEquals = {
            "aws:PrincipalOrgID" = "o-dsvxwengs8"
          }
        }
      },
      {
        Sid    = "DenyAllOthersFromModifyingLambda",
        Effect = "Deny",
        Action = [
          "lambda:UpdateFunctionCode",
          "lambda:UpdateFunctionConfiguration",
          "lambda:AddPermission",
          "lambda:RemovePermission"
        ],
        Resource  = aws_lambda_function.lambda_function.arn,
        Principal = "*"
      }
    ]
  })
}

resource "aws_iam_role_policy_attachment" "lambda_policy_attachment" {
  role       = aws_iam_role.lambda_execution_role.name
  policy_arn = "arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole"
}

resource "aws_iam_role_policy_attachment" "lambda_modifications_policy_attachment" {
  role       = aws_iam_role.lambda_execution_role.name
  policy_arn = aws_iam_policy.restrict_lambda_modifications.arn
}
