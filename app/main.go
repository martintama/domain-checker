package main

import (
	"fmt"
	"os"

	"github.com/martintama/domain-checker/aws"
	"github.com/martintama/domain-checker/cmd"
	"github.com/martintama/domain-checker/internal/logger"
	"github.com/martintama/domain-checker/internal/types"

	"github.com/aws/aws-lambda-go/lambda"
)

// DetermineRunMode checks the environment to determine how the app should run
func DetermineRunMode() types.RunMode {
	if os.Getenv("AWS_LAMBDA_FUNCTION_NAME") != "" {
		return types.RunModeLambda
	}
	return types.RunModeCLI
}

func init() {

}

func main() {

	runMode := DetermineRunMode()

	logger.Initialize(logger.Config{
		LogLevel: logger.ParseLogLevel(os.Getenv("LOG_LEVEL")),
		RunMode:  runMode,
	})

	switch runMode {
	case types.RunModeLambda:
		logger.Info("Starting as Lambda")
		// Run as Lambda
		lambda.Start(aws.HandleRequest)
	case types.RunModeCLI:
		// Run as CLI
		if err := cmd.RootCmd.Execute(); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}
}
