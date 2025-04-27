package aws

import (
	"context"
	"fmt"
	"time"

	"github.com/martintama/domain-checker/internal/client"
	"github.com/martintama/domain-checker/internal/logger"
	"github.com/martintama/domain-checker/internal/types"
)

const (
	domainKey = "domain"
)

// Lambda handler
func HandleRequest(ctx context.Context, event map[string]string) (string, error) {
	// Extract input from Lambda event
	domain, ok := event[domainKey]
	if !ok {
		return "", fmt.Errorf("default input missing: %s", domainKey)
	}

	// Create request-scoped logger
	reqLogger := logger.WithField(domainKey, domain)

	w := client.NewWhoIsClient()
	w.Timeout = 2 * time.Second

	result, err := w.CheckDomainAvailability(domain, reqLogger)
	if err != nil {
		reqLogger.Error("error checking domain availability")
		return types.DomainStatusUnknown, err
	}

	r := fmt.Sprintf("%s: %s", domain, result)
	reqLogger.Info(r)
	return r, nil
}
