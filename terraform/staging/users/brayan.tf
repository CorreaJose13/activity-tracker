resource "aws_iam_user" "brayan" {
  name = "brayan"
  path = "/"
}

resource "aws_iam_user_group_membership" "brayan_groups" {
  user   = aws_iam_user.brayan.name
  groups = [var.tracker_group]
}
