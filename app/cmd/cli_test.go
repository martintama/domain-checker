package cmd

import (
	"testing"

	"github.com/martintama/domain-checker/internal/types"
)

func TestWhoIs(t *testing.T) {
	tests := []struct {
		name           string
		domain         string
		expectedResult types.DomainStatus
		expectError    bool
		expectedError  error
	}{
		{
			name:           "Existing domain",
			domain:         "google.com",
			expectedResult: types.DomainStatusUnavailable,
			expectError:    false,
			expectedError:  nil,
		},
	}

	for _, ts := range tests {
		t.Run(ts.name, func(t *testing.T) {
			r, err := RunWhois(ts.domain, false)

			if err != nil {
				if err != ts.expectedError {
					t.Errorf("Expected error %v, got %v", ts.expectedError, err)
				}
			}
			if r != ts.expectedResult {
				t.Errorf("Expected %v, got %v", ts.expectedResult, r)
			}

		})
	}
}
