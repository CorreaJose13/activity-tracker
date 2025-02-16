resource "aws_iam_user" "valery" {
  name = "valery"
  path = "/"
}

resource "aws_iam_user_group_membership" "valery_groups" {
  user   = aws_iam_user.valery.name
  groups = [var.developer_group]
}
