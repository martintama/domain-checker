package client

import (
	"fmt"

	"github.com/martintama/domain-checker/internal/types"
)

type WhoisMock struct{}

const (
	domainAvailable    string = "available.com"
	domainNotAvailable string = "notavailable.com"
	domainUnknown      string = ""
)

func NewWhoisMock() WhoisClient {
	return &WhoisMock{}
}

func (m WhoisMock) CheckDomainAvailability(domain string, verbose bool) (types.DomainStatus, error) {
	if verbose {
		fmt.Printf("Checking domain availability for: %s", domain)
	}

	switch domain {
	case domainAvailable:
		return types.DomainStatusAvailable, nil
	case domainNotAvailable:
		return types.DomainStatusUnavailable, nil
	}

	return types.DomainStatusUnknown, fmt.Errorf("error getting domain information")
}
