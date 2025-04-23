package main

import (
	"fmt"
	"os"

	"github.com/martintama/domain-checker/aws"
	"github.com/martintama/domain-checker/cmd"

	"github.com/aws/aws-lambda-go/lambda"
)

// RunMode determines how the application should run
type RunMode int

const (
	// ModeLambda indicates the app should run as an AWS Lambda function
	ModeLambda RunMode = iota
	// ModeCLI indicates the app should run as a CLI application
	ModeCLI
)

// DetermineRunMode checks the environment to determine how the app should run
func DetermineRunMode() RunMode {
	if os.Getenv("AWS_LAMBDA_FUNCTION_NAME") != "" {
		return ModeLambda
	}
	return ModeCLI
}

func main() {

	runMode := DetermineRunMode()

	switch runMode {
	case ModeLambda:
		fmt.Println("Starting as Lambda")
		// Run as Lambda
		lambda.Start(aws.HandleRequest)
	case ModeCLI:
		// Run as CLI
		if err := cmd.RootCmd.Execute(); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}
}
