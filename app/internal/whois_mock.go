package internal

import "fmt"

type WhoisMock struct{}

const (
	domainAvailable    string = "available.com"
	domainNotAvailable string = "notavailable.com"
	domainUnknown      string = ""
)

func NewWhoisMock() WhoisClient {
	return &WhoisMock{}
}

func (m WhoisMock) CheckDomainAvailability(domain string, verbose bool) (DomainStatus, error) {
	if verbose {
		fmt.Printf("Checking domain availability for: %s", domain)
	}

	switch domain {
	case domainAvailable:
		return DomainStatusAvailable, nil
	case domainNotAvailable:
		return DomainStatusUnavailable, nil
	}

	return DomainStatusUnknown, fmt.Errorf("error getting domain information")
}
