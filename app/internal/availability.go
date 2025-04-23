package internal

import (
	"fmt"
	"regexp"
)

// patterns that indicate a domain is available
var availabilityPatterns = []string{
	"(?i)No match",
	"(?i)NOT FOUND",
	"(?i)Not fo",
	"(?i)No Data Fou",
	"(?i)has not been regi",
	"(?i)No entri",
	"(?i)Domain not found",
	"(?i)Status: free",
	"(?i)Status: AVAILABLE",
	"(?i)No Object Found",
	"(?i)Domain Status: free",
	"(?i)The domain has not been registered",
}

func prepareRegexPatterns() ([]*regexp.Regexp, error) {
	// Compile all regex patterns once
	var regexPatterns []*regexp.Regexp
	for _, pattern := range availabilityPatterns {
		regex, err := regexp.Compile(pattern)
		if err != nil {
			fmt.Printf("Error compiling regex pattern '%s': %v\n", pattern, err)
			return nil, err
		}
		regexPatterns = append(regexPatterns, regex)
	}

	return regexPatterns, nil
}

func analyzeResult(lookupResult string, verbose bool) (DomainStatus, error) {

	patterns, err := prepareRegexPatterns()
	if err != nil {
		return DomainStatusUnknown, err
	}

	// Check if any availability pattern matches the whois response
	for _, pattern := range patterns {
		if pattern.MatchString(lookupResult) {
			if verbose {
				fmt.Printf("Found match for string: %v\n", pattern.String())
			}

			return DomainStatusAvailable, nil
		}
	}

	return DomainStatusUnavailable, nil
}
