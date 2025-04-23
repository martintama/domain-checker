package internal

import (
	"fmt"
	"strings"
	"time"

	"github.com/likexian/whois"
)

// DomainStatus represents the availability status of a domain
type DomainStatus string

const (
	// DomainStatusAvailable indicates the domain is available for registration
	DomainStatusAvailable DomainStatus = "DomainAvailable"
	// DomainStatusUnavailable indicates the domain is already registered
	DomainStatusUnavailable = "DomainUnavailable"
	// DomainStatusUnknown indicates there was an error and domain status cannot be determined
	DomainStatusUnknown = ""
)

// Map of TLDs to their authoritative WHOIS servers
var tldServerMap = map[string]string{
	".com":    "whois.verisign-grs.com",
	".net":    "whois.verisign-grs.com",
	".org":    "whois.pir.org",
	".info":   "whois.afilias.net",
	".com.ar": "whois.nic.ar",
	".ar":     "whois.nic.ar",
	".co.jp":  "whois.jprs.jp",
	".jp":     "whois.jprs.jp",
	// Add more as needed
}

type WhoisClient interface {
	// CheckDomainAvailability queries the appropriate WHOIS server for the given TLD
	CheckDomainAvailability(domain string, verbose bool) (DomainStatus, error)
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

func (c *DefaultWhoisClient) CheckDomainAvailability(domain string, verbose bool) (DomainStatus, error) {
	var raw string
	var err error

	if verbose {
		fmt.Printf("Checking availability of %s\n", domain)
	}

	tld, err := extractTld(domain)
	if err != nil {
		return DomainStatusUnknown, err
	}

	whoisClient := whois.NewClient()
	whoisClient.SetTimeout(5 * time.Second) // Increased timeout for international servers

	// Check if we have a specific server for this TLD
	server, found := tldServerMap[tld]

	if verbose {
		fmt.Printf("Using server: %s\n", server)
	}
	if found {
		// Query the specific server for this TLD
		raw, err = whoisClient.Whois(domain, server)
	} else {
		// If we don't know the specific server, use the default behavior
		raw, err = whoisClient.Whois(domain)
	}

	if err != nil {
		fmt.Printf("Error querying whois for %s: %s\n", domain, err)
		return DomainStatusUnknown, err
	}

	if verbose {
		fmt.Println(raw)
	}

	result, err := analyzeResult(raw, verbose)
	if err != nil {
		fmt.Printf("Error analyzing lookup result for %s: %s\n", domain, err)
		return DomainStatusUnknown, err
	}

	return result, nil
}
