package types

// RunMode determines how the application should run
type RunMode int

const (
	// ModeLambda indicates the app should run as an AWS Lambda function
	RunModeLambda RunMode = iota
	// ModeCLI indicates the app should run as a CLI application
	RunModeCLI
)
