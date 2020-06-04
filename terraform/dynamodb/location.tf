resource "aws_dynamodb_table" "dynamodb_locations" {
  name = "locations"
  hash_key = "LocationId"
  billing_mode = "PROVISIONED"
  read_capacity = 4
  write_capacity = 4

  attribute {
    name = "LocationId"
    type = "S"
  }

  attribute {
    name = "UserId"
    type = "S"
  }

  global_secondary_index {
    name = "UserIdIndex"
    hash_key = "UserId"
    projection_type = "ALL"

    # TODO: optimize it!
    read_capacity = 2
    write_capacity = 2
  }
}
//
//resource "aws_appautoscaling_target" "dynamodb_locations_read_target" {
//  max_capacity       = 10
//  min_capacity       = 4
//  resource_id        = "table/${aws_dynamodb_table.dynamodb_locations.name}"
//  scalable_dimension = "dynamodb:table:ReadCapacityUnits"
//  service_namespace  = "dynamodb"
//}
//
//resource "aws_appautoscaling_policy" "dynamodb_locations_table_read_policy" {
//  name               = "DynamoDBReadCapacityUtilization:${aws_appautoscaling_target.dynamodb_locations_read_target.resource_id}"
//  policy_type        = "TargetTrackingScaling"
//  resource_id        = aws_appautoscaling_target.dynamodb_locations_read_target.resource_id
//  scalable_dimension = aws_appautoscaling_target.dynamodb_locations_read_target.scalable_dimension
//  service_namespace  = aws_appautoscaling_target.dynamodb_locations_read_target.service_namespace
//
//  target_tracking_scaling_policy_configuration {
//    predefined_metric_specification {
//      predefined_metric_type = "DynamoDBReadCapacityUtilization"
//    }
//
//    target_value = 70
//  }
//}
//
//resource "aws_appautoscaling_target" "dynamodb_locations_write_target" {
//  max_capacity       = 10
//  min_capacity       = 4
//  resource_id        = "table/${aws_dynamodb_table.dynamodb_locations.name}"
//  scalable_dimension = "dynamodb:table:ReadCapacityUnits"
//  service_namespace  = "dynamodb"
//}
//
//resource "aws_appautoscaling_policy" "dynamodb_locations_write_policy" {
//  name               = "DynamoDBReadCapacityUtilization:${aws_appautoscaling_target.dynamodb_locations_write_target.resource_id}"
//  policy_type        = "TargetTrackingScaling"
//  resource_id        = aws_appautoscaling_target.dynamodb_locations_write_target.resource_id
//  scalable_dimension = aws_appautoscaling_target.dynamodb_locations_write_target.scalable_dimension
//  service_namespace  = aws_appautoscaling_target.dynamodb_locations_write_target.service_namespace
//
//  target_tracking_scaling_policy_configuration {
//    predefined_metric_specification {
//      predefined_metric_type = "DynamoDBReadCapacityUtilization"
//    }
//
//    target_value = 70
//  }
//}