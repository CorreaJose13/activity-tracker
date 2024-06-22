//Mejorar el manejo de estos secrets luego
data "aws_secretsmanager_secret" "bot_token" {
  name = "BOT_TOKEN"
}

data "aws_secretsmanager_secret" "mongo_token" {
  name = "MONGO_TOKEN"
}

data "aws_secretsmanager_secret_version" "bot_token_version" {
  secret_id = data.aws_secretsmanager_secret.bot_token.id
}

data "aws_secretsmanager_secret_version" "mongo_token_version" {
  secret_id = data.aws_secretsmanager_secret.mongo_token.id
}
