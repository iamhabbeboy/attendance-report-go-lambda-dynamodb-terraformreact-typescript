#lambda function
provider "aws" {
  region = "us-east-1" # Change to your desired AWS region
}

resource "aws_lambda_function" "example" {
  function_name = "attendance-report"
  handler      = "example.handler"
  runtime      = "go1.x"
  role         = aws_iam_role.lambda_role.arn
  filename     = "example.zip"
}

resource "aws_iam_role" "lambda_role" {
  name = "example-lambda-role"
  
  assume_role_policy = jsonencode({
    Version = "2012-10-17",
    Statement = [{
      Action = "sts:AssumeRole",
      Effect = "Allow",
      Principal = {
        Service = "lambda.amazonaws.com"
      }
    }]
  })
}

resource "aws_iam_policy" "lambda_policy" {
  name        = "example-lambda-policy"
  description = "Policy for the example Lambda function"

  policy = jsonencode({
    Version = "2012-10-17",
    Statement = [{
      Action   = "logs:CreateLogGroup",
      Effect   = "Allow",
      Resource = "*",
    }, {
      Action   = "logs:CreateLogStream",
      Effect   = "Allow",
      Resource = "*",
    }, {
      Action   = "logs:PutLogEvents",
      Effect   = "Allow",
      Resource = "*",
    }]
  })
}

resource "aws_iam_role_policy_attachment" "lambda_policy_attachment" {
  policy_arn = aws_iam_policy.lambda_policy.arn
  role       = aws_iam_role.lambda_role.name
}

#DynamoDB