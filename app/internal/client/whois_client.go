package client

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/likexian/whois"
	"github.com/martintama/domain-checker/internal/logger"
	"github.com/martintama/domain-checker/internal/types"
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

// Map of TLDs to their authoritative WHOIS servers
var tldServerMap = map[string]string{
	".com":   "whois.verisign-grs.com",
	".co.jp": "whois.jprs.jp",
	// Add more as needed, but be mindful of editing the availability patterns too.
}

type WhoisClient interface {
	// CheckDomainAvailability queries the appropriate WHOIS server for the given TLD
	CheckDomainAvailability(domain string, verbose bool) (types.DomainStatus, error)
}

type DefaultWhoisClient struct {
	Timeout time.Duration
}

func NewWhoIsClient() *DefaultWhoisClient {
	return &DefaultWhoisClient{
		Timeout: 5 * time.Second,
	}
}

func extractTld(domain string) (string, error) {
	// Get the TLD by finding the first dot and taking everything after it
	dotIndex := strings.Index(domain, ".")
	if dotIndex == -1 {
		return "", fmt.Errorf("invalid domain format for '%s'. Please include TLD (e.g., domain.com)", domain)
	}

	tld := domain[dotIndex:] // Everything from the first dot onwards

	return tld, nil
}

func (c *DefaultWhoisClient) CheckDomainAvailability(domain string, log *logger.AppLogger) (types.DomainStatus, error) {
	var raw string
	var err error

	log.Debugf("Checking availability of %s\n", domain)

	tld, err := extractTld(domain)
	if err != nil {
		return types.DomainStatusUnknown, err
	}

	whoisClient := whois.NewClient()
	whoisClient.SetTimeout(5 * time.Second) // Increased timeout for international servers

	// Check if we have a specific server for this TLD
	server, found := tldServerMap[tld]

	log.Debugf("Using server: %s\n", server)

	if found {
		// Query the specific server for this TLD
		raw, err = whoisClient.Whois(domain, server)
	} else {
		// If we don't know the specific server, use the default behavior
		raw, err = whoisClient.Whois(domain)
	}

	if err != nil {
		log.Errorf("Error querying whois for %s: %s\n", domain, err)
		return types.DomainStatusUnknown, err
	}

	log.Debug(raw)

	result, err := analyzeResult(raw, log)
	if err != nil {
		log.Errorf("Error analyzing lookup result for %s: %s\n", domain, err)
		return types.DomainStatusUnknown, err
	}

	return result, nil
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

func analyzeResult(lookupResult string, log *logger.AppLogger) (types.DomainStatus, error) {

	patterns, err := prepareRegexPatterns()
	if err != nil {
		return types.DomainStatusUnknown, err
	}

	// Check if any availability pattern matches the whois response
	for _, pattern := range patterns {
		if pattern.MatchString(lookupResult) {
			log.Debugf("Found match for string: %v\n", pattern.String())

			return types.DomainStatusAvailable, nil
		}
	}

	return types.DomainStatusUnavailable, nil
}
