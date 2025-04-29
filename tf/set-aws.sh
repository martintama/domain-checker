#!/bin/bash

eval "$(aws configure export-credentials --profile tf-admin --format env)"
echo "AWS credentials have been set"
