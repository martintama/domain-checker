package internal

import (
	"embed"
	"testing"

	"github.com/stretchr/testify/assert"
)

//go:embed testdata/*.txt
var testDataFiles embed.FS

func TestAvailability(t *testing.T) {
	tests := []struct {
		name     string
		datafile string
		expected DomainStatus
	}{
		{
			name:     "1-not available",
			datafile: "testdata/1-not-available.txt",
			expected: DomainStatusUnavailable,
		},
		{
			name:     "2-available",
			datafile: "testdata/2-available.txt",
			expected: DomainStatusAvailable,
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
