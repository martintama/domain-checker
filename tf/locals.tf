locals {
  function_name = "domain-checker"
  binary_name   = "bootstrap"

  binary_path  = "${path.module}/../app/bin/${local.binary_name}"
  archive_path = "${path.module}/tf_generated/${local.function_name}.zip"

  alarm_name = "${var.domain}-available"
}