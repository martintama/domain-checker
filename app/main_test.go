package main

import (
	"os"
	"testing"

	"github.com/martintama/domain-checker/internal/types"
)

func TestMode(t *testing.T) {
	tests := []struct {
		name         string
		envSetup     func()
		expectedMode types.RunMode
		envTeardown  func()
	}{
		{
			name: "Lambda mode",
			envSetup: func() {
				os.Setenv("AWS_LAMBDA_FUNCTION_NAME", "test-function")
			},
			expectedMode: types.RunModeLambda,
			envTeardown: func() {
				os.Unsetenv("AWS_LAMBDA_FUNCTION_NAME")
			},
		},
		{
			name:         "CLI mode",
			envSetup:     func() {},
			expectedMode: types.RunModeCLI,
			envTeardown:  func() {},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.envSetup()

			mode := DetermineRunMode()

			if mode != tc.expectedMode {
				t.Errorf("Expected mode %v, got %v", tc.expectedMode, mode)
			}

			tc.envTeardown()
		})
	}
}
