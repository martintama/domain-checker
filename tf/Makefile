# Main targets
.PHONY: login init plan apply help

# Create env vars to be able to use terraform locally
login:
	@echo "Logging in to AWS. Please check the browser window."
	aws sso login --profile tf-admin

	@echo "Success! To set the environment variables please run: "
	@echo "source ./set-aws.sh"

init:
	terraform init -backend-config=backend-config.hcl

plan:
	terraform plan -out=plan.out

apply:
	terraform apply plan.out

help:
	@echo "Available targets:"
	@echo "  login            : Login to AWS"
	@echo "  init             : Run terraform init"
	@echo "  plan             : Run terraform plan"
	@echo "  apply            : Run terraform apply"
