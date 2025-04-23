package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
		expectedStatus         DomainStatus
		expectError            bool
		expectedErrorSubstring string
	}{
		{
			name:           "available",
			domain:         "available.com",
			expectedStatus: DomainStatusAvailable,
			expectError:    false,
		},
		{
			name:           "not-available",
			domain:         "notavailable.com",
			expectedStatus: DomainStatusUnavailable,
			expectError:    false,
		},
		{
			name:                   "invalid",
			domain:                 "test",
			expectedStatus:         DomainStatusUnknown,
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
