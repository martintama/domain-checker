# Domain checker

## How to run this locally?

### App

### Terraform code

This expects a local AWS SSO session named `tf` is configured. You can do so by running:

`aws configure sso`

And follow the instructions. When asked about the **Session name**, enter `tf` (Recommended). If you want to use a different name,
please ensure scripts and files are changed to reflect the new naming.

With that out of the way:

1. Run `make login`
2. Run `source ./set-aws.sh`. This will set the AWS creds as environment variables so they can be used by the s3 backend for the state.
3. (Change whatever needs to be changed)
4. `make plan`. This will trigger the `terraform plan` command.
5. `make apply`. This will run the `terraform apply`