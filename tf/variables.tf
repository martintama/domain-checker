variable "aws_region" {
  default = "eu-west-1"
  type    = string
}

variable "notification_email" {
  description = "Email address to receive domain availability notifications"
  type        = string
}

variable "domain" {
  description = "Domain to automatically check for availability"
  type        = string
}