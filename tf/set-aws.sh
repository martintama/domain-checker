#!/bin/bash

eval "$(aws configure export-credentials --profile tf --format env)"
echo "AWS credentials have been set"
