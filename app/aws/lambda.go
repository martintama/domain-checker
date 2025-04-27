package aws

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/martintama/domain-checker/internal/client"
	"github.com/martintama/domain-checker/internal/logger"
	"github.com/martintama/domain-checker/internal/types"
)

// Lambda handler
func HandleRequest(ctx context.Context, event map[string]string) (string, error) {

	verbose := false
	if os.Getenv("APP_LOG_LEVEL") != "" {
		verbose = true
	}

	// Extract input from Lambda event
	d, ok := event["domain"]
	if !ok {
		return "", fmt.Errorf("default input missing: domain")
	}

	w := client.NewWhoIsClient()

	w.Timeout = 2 * time.Second

	result, err := w.CheckDomainAvailability(d, verbose)
	if err != nil {
		logger.WithField("domain", d).Error("error checking domain")
		return types.DomainStatusUnknown, err
	}

	r := fmt.Sprintf("%s: %s", d, result)
	return r, nil
}
