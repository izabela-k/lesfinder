resource "aws_dynamodb_table" "users" {
  name              = "users"
  hash_key          = "Id"
  billing_mode      = "PROVISIONED"
  read_capacity     = 4
  write_capacity    = 4

  attribute {
    name = "Id"
    type = "S"
  }

  attribute {
    name = "Email"
    type = "S"
  }

  global_secondary_index {
    name = "EmailIndex"
    hash_key = "Email"
    write_capacity = 4
    read_capacity = 4
    projection_type = "ALL"
  }
}