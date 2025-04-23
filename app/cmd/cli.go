package cmd

import (
	"fmt"
	"time"

	"github.com/martintama/domain-checker/internal"
	"github.com/spf13/cobra"
)

const domainFlag = "domain"
const verboseFlag = "verbose"

var WhoisCmd = &cobra.Command{
	Use:   "check",
	Short: "Check whois records for a given domain name",
	Run: func(cmd *cobra.Command, args []string) {
		domain, _ := cmd.Flags().GetString(domainFlag)
		verbose, _ := cmd.Flags().GetBool(verboseFlag)
		RunWhois(domain, verbose)
	},
}

func RunWhois(domain string, verbose bool) (internal.DomainStatus, error) {
	w := internal.NewWhoIsClient()
	w.Timeout = 2 * time.Second

	result, err := w.CheckDomainAvailability(domain, verbose)
	if err != nil {
		fmt.Println("Error checking domain")
		return internal.DomainStatusUnknown, err
	}
	fmt.Println(result)

	return result, nil
}

func init() {
	// Add the WhoisCmd as a subcommand of RootCmd
	RootCmd.AddCommand(WhoisCmd)

	// Add flags to WhoisCmd
	WhoisCmd.Flags().StringP(domainFlag, "d", "", "Domain name to check (required)")
	WhoisCmd.Flags().BoolP(verboseFlag, "v", false, "Enable verbose mode")
	WhoisCmd.MarkFlagRequired(domainFlag)
}
