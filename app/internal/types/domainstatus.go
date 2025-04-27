package types

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
