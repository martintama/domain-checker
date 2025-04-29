# Domain checker

This is a super basic application that runs on AWS Lambda and checks the availability status of a given domain name. It sends an email notification when (or IF) it gets released.

Checks are made using a simple whois lookup, trying to stay under the radar, as any other action like opening the URL in the browser could signal there's interest to the owners.

## Components

This is composed of two parts:
- App: Written in go, can be run locally to confirm the domain in question is effectively detected as not available to start with.
- TF code: Builds and deploys the app to an AWS Lambda function, that is configured to run hourly and log the results to Clodwatch Logs. From there, a Cloudwatch Metric Filter looks for the text "DomainAvailable". If found, it sends an email to the pre-configured email.

Currently it sends an email every time it runs (i.e. hourly) if the domain is found available. It would have been better to have a "cooldown period", and only send emails every 2 days, for example. Unfortunately, it would have meant adding more innecessary components to the mix, increasing the overall solution complexity. 

On the other hand, worst case scenario is that we get 8-10 email while we sleep, but we won't miss the news!

## How to run this locally?

### App

The application code is located in the `app` directory.

This is a standard Go application following conventional Go project structure.

A Makefile is provided with common operations:

`make test`: Run tests
`make build`: Build app

View all available commands with `make help`.

#### Running the app

First, build the application from the `app` directory:

`make build`

Then run it with:

`./bin/domain-checker check --domain domaintocheck.com` 

### Terraform code

> [!INFORMATION]
> Minimum Terraform version required: 1.11.4

Terraform code is located in the `tf` directory.

#### Prerequisites

##### AWS
Terraform configuration expects a local AWS SSO session named `tf-admin` is configured. Set this up with:

`aws configure sso`

Follow the instructions. When asked about the **Session name**, enter `tf-admin` (Recommended). If you want to use a different name, please ensure scripts and files are changed accordingly.

> [!IMPORTANT]
> Terraform code needs to create new IAM roles for Lambda functions. Thus, AWS `PowerUserAccess` role is insufficient - you'll need administrative privileges.

##### Statefile
Terraform is configured to store the state in a S3 bucket. To do so, rename the `tf/backend-config.hcl.template` file to `tf/backend-config.hcl` (remove the `.template` suffix), and point to a bucket that already exists.

The role for AWS must have the following IAM permissions on the bucket:
- `s3:ListBucket` on `arn:aws:s3:::mybucket`. At a minimum, this must be able to list the path where the state is stored.
- `s3:GetObject` on `arn:aws:s3:::mybucket/path/to/my/key` and `arn:aws:s3:::mybucket/path/to/my/key.tflock` (lockfile)
- `s3:PutObject` on `arn:aws:s3:::mybucket/path/to/my/key` and `arn:aws:s3:::mybucket/path/to/my/key.tflock` (lockfile)

Please refer to the S3 backend [documentation](https://developer.hashicorp.com/terraform/language/backend/s3#s3-bucket-permissions) for more details.

#### Deployment

1. Run `make login` to authenticate with SSO
1. Run `source ./set-aws.sh` to configure AWS credentials as environment variables for the S3 backend state.
1. (First time only) `make init`. This calls `terraform init` on the directory.
1. `make plan`. This will trigger the `terraform plan` command.
1. `make apply`. This will run the `terraform apply`
