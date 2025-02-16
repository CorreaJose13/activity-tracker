resource "aws_iam_user" "dummy" {
  name = "dummy"
  path = "/"
}

resource "aws_iam_user_group_membership" "dummy_groups" {
  user   = aws_iam_user.dummy.name
  groups = [var.tracker_group]
}
