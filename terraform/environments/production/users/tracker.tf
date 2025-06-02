resource "aws_iam_group" "tracker" {
  name = "tracker"
  path = "/"
}

resource "aws_iam_group_policy_attachment" "tracker_dynamodb" {
  group      = aws_iam_group.tracker.name
  policy_arn = "arn:aws:iam::aws:policy/AmazonDynamoDBFullAccess"
}

resource "aws_iam_group_policy_attachment" "tracker_eventbridge" {
  group      = aws_iam_group.tracker.name
  policy_arn = "arn:aws:iam::aws:policy/AmazonEventBridgeSchedulerFullAccess"
}

resource "aws_iam_group_policy_attachment" "tracker_s3" {
  group      = aws_iam_group.tracker.name
  policy_arn = "arn:aws:iam::aws:policy/AmazonS3FullAccess"
}

resource "aws_iam_group_policy_attachment" "tracker_kms" {
  group      = aws_iam_group.tracker.name
  policy_arn = "arn:aws:iam::aws:policy/AWSKeyManagementServicePowerUser"
}

resource "aws_iam_group_policy_attachment" "tracker_iam" {
  group      = aws_iam_group.tracker.name
  policy_arn = "arn:aws:iam::aws:policy/IAMFullAccess"
}

resource "aws_iam_group_policy_attachment" "tracker_change_password" {
  group      = aws_iam_group.tracker.name
  policy_arn = "arn:aws:iam::aws:policy/IAMUserChangePassword"
}
