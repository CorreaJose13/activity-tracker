resource "aws_iam_group" "developer" {
  name = "developer"
  path = "/"
}

resource "aws_iam_group_policy_attachment" "developer_dynamodb" {
  group      = aws_iam_group.developer.name
  policy_arn = "arn:aws:iam::aws:policy/AmazonDynamoDBFullAccess"
}

resource "aws_iam_group_policy_attachment" "developer_lambda" {
  group      = aws_iam_group.developer.name
  policy_arn = "arn:aws:iam::aws:policy/AWSLambda_FullAccess"
}

resource "aws_iam_group_policy_attachment" "developer_eventbridge" {
  group      = aws_iam_group.developer.name
  policy_arn = "arn:aws:iam::aws:policy/AmazonEventBridgeSchedulerFullAccess"
}

resource "aws_iam_group_policy_attachment" "developer_s3" {
  group      = aws_iam_group.developer.name
  policy_arn = "arn:aws:iam::aws:policy/AmazonS3FullAccess"
}

resource "aws_iam_group_policy_attachment" "developer_kms" {
  group      = aws_iam_group.developer.name
  policy_arn = "arn:aws:iam::aws:policy/AWSKeyManagementServicePowerUser"
}

resource "aws_iam_group_policy_attachment" "developer_iam" {
  group      = aws_iam_group.developer.name
  policy_arn = "arn:aws:iam::aws:policy/IAMFullAccess"
}

resource "aws_iam_group_policy_attachment" "developer_change_password" {
  group      = aws_iam_group.developer.name
  policy_arn = "arn:aws:iam::aws:policy/IAMUserChangePassword"
}

resource "aws_iam_group_policy_attachment" "developer_cloudwatch_logs" {
  group      = aws_iam_group.developer.name
  policy_arn = "arn:aws:iam::aws:policy/CloudWatchLogsFullAccess"
}
