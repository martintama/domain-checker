package client

import (
	"embed"
	"testing"

	"github.com/martintama/domain-checker/internal/types"
	"github.com/stretchr/testify/assert"
)

//go:embed testdata/*.txt
var testDataFiles embed.FS

func TestAvailability(t *testing.T) {
	tests := []struct {
		name     string
		datafile string
		expected types.DomainStatus
	}{
		{
			name:     "1-not available",
			datafile: "testdata/1-not-available.txt",
			expected: types.DomainStatusUnavailable,
		},
		{
			name:     "2-available",
			datafile: "testdata/2-available.txt",
			expected: types.DomainStatusAvailable,
		},
	}

	for _, ts := range tests {
		t.Run(ts.name, func(t *testing.T) {
			r, err := testDataFiles.ReadFile(ts.datafile)
			if err != nil {
				t.Errorf("Error testing %s: %v", ts.name, err)
			}

			status, _ := analyzeResult(string(r), false)

			assert.Equal(t, ts.expected, status)
		})
	}
}

func TestExtractTld(t *testing.T) {
	tests := []struct {
		name                   string
		domain                 string
		expectedTld            string
		expectError            bool
		expectedErrorSubstring string
	}{
		{
			name:        "valid .com",
			domain:      "test.com",
			expectedTld: ".com",
			expectError: false,
		},
		{
			name:        "valid .co.jp",
			domain:      "test.co.jp",
			expectedTld: ".co.jp",
			expectError: false,
		},
		{
			name:                   "invalid",
			domain:                 "test",
			expectedTld:            "",
			expectError:            true,
			expectedErrorSubstring: "invalid domain format",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tld, err := extractTld(tc.domain)

			if tc.expectError {
				assert.Error(t, err)
				assert.ErrorContains(t, err, tc.expectedErrorSubstring)
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, tc.expectedTld, tld)
		})
	}
}

func TestCheckDomain(t *testing.T) {
	tests := []struct {
		name                   string
		domain                 string
		expectedStatus         types.DomainStatus
		expectError            bool
		expectedErrorSubstring string
	}{
		{
			name:           "available",
			domain:         "available.com",
			expectedStatus: types.DomainStatusAvailable,
			expectError:    false,
		},
		{
			name:           "not-available",
			domain:         "notavailable.com",
			expectedStatus: types.DomainStatusUnavailable,
			expectError:    false,
		},
		{
			name:                   "invalid",
			domain:                 "test",
			expectedStatus:         types.DomainStatusUnknown,
			expectError:            true,
			expectedErrorSubstring: "error getting domain information",
		},
	}

	w := NewWhoisMock()

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			status, err := w.CheckDomainAvailability(tc.domain, false)
			if tc.expectError {
				assert.Error(t, err)
				assert.ErrorContains(t, err, tc.expectedErrorSubstring)
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, tc.expectedStatus, status)
		})
	}
}
