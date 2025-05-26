resource "aws_iam_user" "juan" {
  name = "juan"
  path = "/"
}

resource "aws_iam_user_group_membership" "juan_groups" {
  user   = aws_iam_user.juan.name
  groups = [var.developer_group]
}
