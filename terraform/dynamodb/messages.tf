resource "aws_dynamodb_table" "messages" {
  name = "messages"
  hash_key = "UserId"
  range_key = "CreatedAt"
  billing_mode = "PROVISIONED"
  read_capacity = 2
  write_capacity = 2

  attribute {
    name = "UserId"
    type = "S"
  }

  attribute {
    name = "ToUserId"
    type = "S"
  }

  attribute {
    name = "CreatedAt"
    type = "S"
  }

  global_secondary_index {
    name = "ToUserIdIndex"
    hash_key = "ToUserId"
    projection_type = "ALL"

    # TODO: optimize it!
    read_capacity = 2
    write_capacity = 2
  }
}