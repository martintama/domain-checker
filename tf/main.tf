// Get current AWS region and account ID
data "aws_region" "current" {}

data "aws_caller_identity" "current" {}

// allow lambda service to assume (use) the role with such policy
data "aws_iam_policy_document" "assume_lambda_role" {
  statement {
    actions = ["sts:AssumeRole"]

    principals {
      type        = "Service"
      identifiers = ["lambda.amazonaws.com"]
    }
  }
}

// create lambda role, that lambda function can assume (use)
resource "aws_iam_role" "lambda" {
  name               = "${local.function_name}-execution-role"
  description        = "Execution role for ${local.function_name} lambda function"
  assume_role_policy = data.aws_iam_policy_document.assume_lambda_role.json
}

data "aws_iam_policy_document" "allow_lambda_logging" {
  statement {
    effect = "Allow"
    actions = [
      "logs:CreateLogStream",
      "logs:PutLogEvents",
    ]

    resources = [
      "arn:aws:logs:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:log-group:/aws/lambda/${local.function_name}:*",
    ]
  }
}

// create a policy to allow writing into logs and create logs stream
resource "aws_iam_policy" "function_logging_policy" {
  name        = "AllowLambdaLoggingPolicy"
  description = "Policy for lambda cloudwatch logging"
  policy      = data.aws_iam_policy_document.allow_lambda_logging.json
}

// attach policy to out created lambda role
resource "aws_iam_role_policy_attachment" "lambda_logging_policy_attachment" {
  role       = aws_iam_role.lambda.id
  policy_arn = aws_iam_policy.function_logging_policy.arn
}

// build the binary for the lambda function in a specified path
resource "null_resource" "function_binary" {
  triggers = {
    // Track changes in all Go files in the app directory
    source_code_hash = join("", [for f in fileset("${path.module}/../app", "**/*.go") : filemd5("${path.module}/../app/${f}")])
  }

  provisioner "local-exec" {
    command     = "make buildlambda"
    working_dir = "${path.module}/../app"
  }
}

// zip the binary with the required "bootstrap" name for provided.al2 runtime
data "archive_file" "function_archive" {
  depends_on = [null_resource.function_binary]

  type        = "zip"
  // For AWS Lambda with provided.al2 runtime, the binary must be named "bootstrap"
  source_file = local.binary_path
  output_path = local.archive_path
  
  // This is the critical change - rename the file to "bootstrap" within the ZIP
  output_file_mode = "0755"
}


// create the lambda function from zip file
resource "aws_lambda_function" "function" {
  function_name = local.function_name
  description   = "Periodically checks the availability of the given domain"
  role          = aws_iam_role.lambda.arn
  handler       = local.binary_name
  memory_size   = 128
  timeout       = 10

  filename         = local.archive_path
  source_code_hash = data.archive_file.function_archive.output_base64sha256

  runtime = "provided.al2"
}


// create log group in cloudwatch to gather logs of our lambda function
resource "aws_cloudwatch_log_group" "log_group" {
  name              = "/aws/lambda/${aws_lambda_function.function.function_name}"
  retention_in_days = 7
}

// Add EventBridge scheduler for periodic invocation
resource "aws_cloudwatch_event_rule" "schedule" {
  name                = "${local.function_name}-schedule"
  description         = "Schedule for ${local.function_name} Lambda function"
  schedule_expression = "cron(0 0/12 * * ? *)"
}

resource "aws_cloudwatch_event_target" "lambda_target" {
  rule      = aws_cloudwatch_event_rule.schedule.name
  target_id = "${local.function_name}-target"
  arn       = aws_lambda_function.function.arn
}

resource "aws_lambda_permission" "allow_cloudwatch" {
  statement_id  = "AllowExecutionFromCloudWatch"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.function.function_name
  principal     = "events.amazonaws.com"
  source_arn    = aws_cloudwatch_event_rule.schedule.arn
}
