# Domain checker

## How to run this locally?

### App

The application code is located in the `app` directory.

This is a standard Go application following conventional Go project structure.

A Makefile is provided with common operations:

`make test`: Run tests
`make build`: Build app

View all available commands with with `make help`.

#### Running the app

First, build the application from the `app` directory:

`make build`

Then run it with:

`./bin/domain-checker check --domain domaintocheck.com` 

### Terraform code

Terraform code is located in the `tf` directory.

#### Prerequisites

Terraform configuration expects a local AWS SSO session named `tf-admin` is configured. Set this up with:

`aws configure sso`

Follow the instructions. When asked about the **Session name**, enter `tf-admin` (Recommended). If you want to use a different name, please ensure scripts and files are changed accordingly.

>![IMPORTANT]
> Terraform code needs to create new IAM roles for Lambda functions. The AWS `PowerUserAccess` role is insufficient - you'll need administrative privileges.

#### Deployment

1. Run `make login` to authenticate with SSO
1. Run `source ./set-aws.sh` to configure AWS credentials as environment variables for the S3 backend state.
1. (First time only) `make init`. This calls `terraform init` on the directory.
1. `make plan`. This will trigger the `terraform plan` command.
1. `make apply`. This will run the `terraform apply`
