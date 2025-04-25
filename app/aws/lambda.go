package aws

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/martintama/domain-checker/internal"
)

// Lambda handler
func HandleRequest(ctx context.Context, event map[string]string) (string, error) {

	verbose := false
	if os.Getenv("APP_VERBOSE") != "" {
		verbose = true
	}

	// Extract input from Lambda event
	domain, ok := event["domain"]
	if !ok {
		return "", fmt.Errorf("default input missing: domain")
	}

	w := internal.NewWhoIsClient()
	w.Timeout = 2 * time.Second

	result, err := w.CheckDomainAvailability(domain, verbose)
	if err != nil {
		fmt.Println("Error checking domain")
		return internal.DomainStatusUnknown, err
	}

	r := fmt.Sprintf("%s: %s", domain, result)
	return r, nil
}
