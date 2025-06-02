resource "aws_iam_user" "johan" {
  name = "johan"
  path = "/"
}

resource "aws_iam_user_group_membership" "johan_groups" {
  user   = aws_iam_user.johan.name
  groups = [var.developer_group]
}
